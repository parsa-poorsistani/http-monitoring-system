package healthcheck


import (

  "fmt"
  "net/http"
  "time"
  "github.com/parsa-poorsistani/http-monitoring-system/pkg/config"
  "github.com/parsa-poorsistani/http-monitoring-system/pkg/metric"
  "github.com/parsa-poorsistani/http-monitoring-system/pkg/database"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/sirupsen/logrus"
)
type HealthCheck struct {
  db *database.Database
  cfg *config.Config
  log *logrus.Logger
}

func NewHealthChecker(db *database.Database, cfg *config.Config, log *logrus.Logger) *HealthCheck {
  return &HealthCheck{
    db: db,
    cfg: cfg,
    log: log,
  }
}

func (hc *HealthCheck) Start() {
  ticker := time.NewTicker(time.Duration(hc.cfg.HealthChecker.Interval))

  for {
    select {
    case <- ticker.C:
    hc.checkServers()
  }
  }
}

func (hc *HealthCheck) checkServers() {
  servers, err := hc.db.GetAllServers()
  if err != nil {
    hc.log.WithError(err).Error("Failed to retrieve servers for health check")
    return
  }

  for _, server := range servers {
    hc.checkServerHealth(&server)
  } 
}

func (hc *HealthCheck) checkServerHealth(server *database.Server) {
  timer := prometheus.NewTimer(metric.HealthCheckDuration.WithLabelValues(fmt.Sprintf("%d", server.ID)))
  defer timer.ObserveDuration()

  resp, err := http.Get(server.Address)
  success := err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300

  if err := hc.db.UpdateServerStatus(server.ID, success); err != nil {
    hc.log.WithFields(logrus.Fields{"server_id: ": server.ID, "address: ": server.Address}).WithError(err).Error("Failed to update the server")
  }

}

package metric

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    HttpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Number of HTTP requests",
        },
        []string{"path"},
    )

    HttpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Duration of HTTP requests",
            Buckets: prometheus.DefBuckets,
        },
        []string{"path"},
    )

    HttpRequestsErrorsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_errors_total",
            Help: "Total number of HTTP requests with errors",
        },
        []string{"path"},
    )

    DatabaseErrorsTotal = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "database_errors_total",
            Help: "Total number of database errors",
        },
    )

    DatabaseQueryDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "database_query_duration_seconds",
            Help: "Duration of database queries",
            Buckets: prometheus.DefBuckets,
        },
        []string{"operation"},
    )

    HealthCheckDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "health_check_duration_seconds",
            Help: "Duration of health check operations",
            Buckets: prometheus.DefBuckets,
        },
        []string{"server_id"},
    )
)

func init() {
    // Register the metrics
    prometheus.MustRegister(HttpRequestsTotal)
    prometheus.MustRegister(HttpRequestDuration)
    prometheus.MustRegister(HttpRequestsErrorsTotal)
    prometheus.MustRegister(DatabaseErrorsTotal)
    prometheus.MustRegister(DatabaseQueryDuration)
    prometheus.MustRegister(HealthCheckDuration)
}

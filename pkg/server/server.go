package server

import (
	"encoding/json"
	"net/http"
	"strconv"

  "github.com/sirupsen/logrus"
	"github.com/parsa-poorsistani/http-monitoring-system/pkg/config"
	"github.com/parsa-poorsistani/http-monitoring-system/pkg/database"
)

type Server struct {
	db  *database.Database
	cfg *config.Config
  log *logrus.Logger
}

func NewServer(db *database.Database, cfg *config.Config, logger *logrus.Logger) *Server {
	return &Server{
		db:  db,
		cfg: cfg,
    log: logger,
	}
}

func (s *Server) Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/server", s.handleServer)
	mux.HandleFunc("/api/server/all", s.handleAllServers)

	return mux
}

func (s *Server) handleServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createServerModel(w, r)
	case http.MethodGet:
		s.getServerModel(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleAllServers(w http.ResponseWriter, r *http.Request) {

  s.log.WithFields(logrus.Fields{
        "method": r.Method,
        "endpoint": "/api/server/all",
    }).Info("handleAllServers called")


	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

  servers, err := s.db.GetAllServers()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(servers)
}

func (s *Server) createServerModel(w http.ResponseWriter, r *http.Request) {
  s.log.WithFields(logrus.Fields{
    "method": r.Method,
    "endpoint": "/api/server",
  }).Info("createServerModel called")

  type Request struct {
    Address string `json:"address"`
  }

  var req Request
  decoder := json.NewDecoder(r.Body)  
  if err := decoder.Decode(&req); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  } 

  id, err := s.db.AddServer(req.Address)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  response := map[string]int64{"id": id}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (s *Server) getServerModel(w http.ResponseWriter, r *http.Request) {
  s.log.WithFields(logrus.Fields{
        "method": r.Method,
        "endpoint": "/api/server",
        "id": r.URL.Query().Get("id"),
    }).Info("getServerModel called")


  idStr := r.URL.Query().Get("id")
  
  if idStr == "" {
    http.Error(w, "Server ID is required", http.StatusBadRequest)
    return
  } 
  
  id, err := strconv.ParseInt(idStr, 10, 64)
  if err != nil {
    http.Error(w, "Invalid Server ID", http.StatusBadRequest)
    return
  }

  server, err := s.db.GetServer(id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(server)
}











package database

import (
	"database/sql"
	"fmt"
	"time"

  _ "github.com/lib/pq"
  "github.com/parsa-poorsistani/http-monitoring-system/pkg/config"
)

type Server struct {
	ID          int64
	Address     string
	Success     int64
	Failure     int64
	LastFailure time.Time
	CreatedAt   time.Time
}

type Database struct {
	conn *sql.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password ,cfg.Database.Name)
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    fmt.Println("Error connecting to PostgreSQL:", err)
    return nil, err
  }

  if err = db.Ping(); err != nil {
    fmt.Println("Error pinging PostgreSQL database:", err)
    return nil, err
  }

  return &Database{conn: db}, nil
}

func (db *Database) Close() {
  db.conn.Close()
}

func (db *Database) AddServer(addr string) (int64, error) {
  var id int64
  query := `INSERT INTO servers (address, success, failure, last_failure, created_at)
              VALUES ($1, 0, 0, NULL, NOW()) RETURNING id;`
  err := db.conn.QueryRow(query, addr).Scan(&id)
  if err != nil {
    return 0, err
  }
  return id, nil
} 


func (db *Database) GetServer(id int64) (*Server, error) {
  var s Server
  query := `SELECT id, address, success, failure, last_failure, created_at 
              FROM servers WHERE id = $1;`
  row := db.conn.QueryRow(query, id)
  err := row.Scan(&s.ID, &s.Address, &s.Success, &s.Failure, &s.LastFailure, &s.CreatedAt)
  if err != nil {
    return nil, err
  }
  return &s, nil
}

func (db *Database) GetAllServers() ([]Server, error) {
  var s []Server
  query := `SELECT id, address, success, failure, last_failure, created_at FROM servers;`

    rows, err := db.conn.Query(query)
    if err != nil {
        return nil, err
    }
  defer rows.Close()

  for rows.Next() {
        var server Server
        if err = rows.Scan(&server.ID, &server.Address, &server.Success, &server.Failure, &server.LastFailure, &server.CreatedAt); err != nil {
            return nil, err
        }
        s = append(s, server)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return s, nil
}

func (db *Database) UpdateServerStatus(id int64, success bool) error {
  if success {
        query := `UPDATE servers SET success = success + 1 WHERE id = $1;`
        _, err := db.conn.Exec(query, id)
        return err
    } else {
        query := `UPDATE servers SET failure = failure + 1, last_failure = NOW() WHERE id = $1;`
        _, err := db.conn.Exec(query, id)
        return err
    }
}

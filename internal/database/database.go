package database

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  log "github.com/sirupsen/logrus"
)

type Database struct {
  dbType string
  url    string
  logger *log.Logger
}

func NewDatabase() *Database {
  return &Database{
    "sqllite3",
    "./pa.db",
    log.New(),
  }
}

func (d Database) CreateConn() *sql.DB {
  conn, err := sql.Open(d.dbType, d.url)
  if err != nil {
    d.logger.Error("SQL", "createConn")
  }
  return conn
}

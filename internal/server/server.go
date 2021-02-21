package server

import (
	"database/sql"
	"github.com/benmcmahon100/PA/internal/database"
)

type Server struct {
	dbDefinition *database.Database
	dbConn       *sql.DB
}

func NewServer(dbDefinition *database.Database) *Server {
	return &Server{
		dbDefinition,
		dbDefinition.CreateConn(),
	}
}

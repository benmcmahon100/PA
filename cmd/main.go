package main

import (
	"fmt"
	"github.com/benmcmahon100/PA/internal/database"
	"github.com/benmcmahon100/PA/internal/server"
)

func main() {
	fmt.Println("Test")
	db := database.NewDatabase()
	server := server.NewServer(db)

}



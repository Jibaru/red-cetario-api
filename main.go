package main

import (
	"log"
	"redcetarioapi/config"
	"redcetarioapi/server"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(mysql.Open(cfg.DBURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	s := server.New(db, cfg)
	s.Run(":8080")
}

package handler

import (
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"redcetarioapi/config"
	"redcetarioapi/server"
)

var cfg = config.Load()

func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(cfg.DBURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	dbconn, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}
	defer dbconn.Close()

	server.New(db).ServeHTTP(w, r)
}

package db

import (
	"fmt"
	"log"
	"golang-todo-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Falha ao conectar ao banco de dados:", err)
	}

	// Migração do modelo
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("❌ Falha ao migrar o banco de dados:", err)
	}

	fmt.Println("✅ Banco de dados conectado e migrado com sucesso!")
}

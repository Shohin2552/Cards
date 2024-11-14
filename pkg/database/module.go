package database

import (
	"cards/pkg/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewDatabase() {
	// Строка подключения
	cfg := *config.NewConfig()
	// Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), cfg.PostgresURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v\n", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Успешное подключение к базе данных!")
}

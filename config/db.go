package config

import (
	"context"
	"fmt"
	"log"
	"os"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase() (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://myapp:myapp@localhost:26001/myapp?sslmode=disable"
	}

	dbConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Erro ao parsear a URL do banco de dados: %v", err)
		return nil, err
	}
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	fmt.Println("✅ Conectado ao PostgreSQL!")
	return pool, nil
}

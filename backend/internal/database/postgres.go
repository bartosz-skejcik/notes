package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"notes/config"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func CreatePostgresConnection(env config.EnvVars, lc fx.Lifecycle) *sql.DB {
    var err error

    connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%s", env.POSTGRES_USER, env.POSTGRES_DATABASE, env.POSTGRES_PASSWORD, env.POSTGRES_HOST, env.POSTGRES_PORT)
    conn, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err := conn.Ping(); err != nil {
        log.Fatal(err)
    }

    lc.Append(fx.Hook{
        OnStart: func(context.Context) error {
			if err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			conn.Close()
			return nil
		},

    })

    return conn
}
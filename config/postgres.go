package config

import (
	"context"
	"fmt"
	"log"
	"github.com/go-pg/pg/v10"
	"os"
)

func ConnectPGDB() *pg.DB  {
	fmt.Println("Connecting to Postgres SQL")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER_POSTGRES"),
		os.Getenv("DB_PASSWORD_POSTGRES"),
		os.Getenv("DB_HOST_POSTGRES"),
		os.Getenv("DB_PORT_POSTGRES"),
		os.Getenv("DB_NAME_POSTGRES"))

	opt, err := pg.ParseURL(url)
	if err != nil {
		log.Println("Invalid database URL!")
		log.Fatal(err)
	}

	client := pg.Connect(opt)

	err = client.Ping(context.TODO())
	if err != nil {
		log.Println("Koneksi database terputus!")
		log.Fatal(err)
		defer client.Close()
	}

	fmt.Println("Connected to Postgres SQL")
	return client
}
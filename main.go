package main

import (
	"context"
	"fmt"
	"lab/productLab/internal/database"
	"lab/productLab/internal/store"
	"lab/productLab/internal/transport"
	"lab/productLab/internal/usecase"

	"log"
	"net/http"
	"os"
)

func main() {
	user := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	dbname := os.Getenv("APP_DB_NAME")
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	ctx := context.Background()

	// start db
	db, err := database.NewClientDB(ctx, "postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	m := database.NewMigration(db, "./internal/database/migrations")
	err = m.StartMigration()
	if err != nil {
		log.Fatalf("Migration failed. Error: [%s]", err)
	}

	// usecases, transport
	store := store.NewStore(db)
	ucProduct := usecase.NewProductUC(store)

	middProduct := transport.NewProductTransport(ucProduct)
	router := transport.NewRouter(middProduct)

	log.Fatal(http.ListenAndServe(":8010", router))
}

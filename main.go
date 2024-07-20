package main

import (
	"log"

	"github.com/MohammedAbdulJabbar23/book-management-api/config"
	"github.com/MohammedAbdulJabbar23/book-management-api/routes"
)

func main() {
    config.ConnectDatabase();
    r := routes.SetupRouter();
    log.Fatal(r.Run(":8080"));
}

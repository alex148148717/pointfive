package main

import (
	_ "github.com/lib/pq"
	"pointfive/internal/server"
)

func main() {
	server.NewApp().Run()

}

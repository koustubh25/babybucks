package main

import (
	s "github.com/koustubh25/babybucks/internal/server"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic(err)
	// }
	s.RunServer()
}

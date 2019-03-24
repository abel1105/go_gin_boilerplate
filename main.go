package main

import (
	"flag"
	"fmt"
	"github.com/abel1105/go_gin_boilerplate/db"
	"os"

	"github.com/abel1105/go_gin_boilerplate/config"
	"github.com/abel1105/go_gin_boilerplate/server"
)

func help() {
	fmt.Println("")
	fmt.Println("Usage: ")
	fmt.Println("")
	fmt.Println(" -env {mode}, default: development")
	fmt.Println(" -port {port}, default: 8080")
	os.Exit(1)
}

func main() {
	env := flag.String("env", "development", "")
	port := flag.String("port", "8080", "")
	flag.Usage = help
	flag.Parse()
	config.Init(*env, *port)
	DB:= db.Init()
	server.Init()
	defer DB.Close()
}
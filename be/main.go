package main

import (
	"server/databases"
	"server/servers"
)

func main() {
	databases.Seed()
	servers.RunServer()
}

package main

import (
	_ "github.com/lib/pq"
	"github.com/luuisavelino/crud-golang/pkg/routes"
)

func main() {

	routes.HandlerRequest()
}

package main

import (
	"github.com/gocraft/web"
	"github.com/minivac/lebroc/app/core"
	"github.com/minivac/lebroc/app/gql"
	"net/http"
)

func main() {
	router := web.New(core.Context{})

	core.Configure(router)
	gql.Configure(router)

	http.ListenAndServe("localhost:3000", router)
}


package core

import (
	"github.com/gocraft/web"
	"fmt"
)

type Context struct {
}

var BookService struct {
	Repo *BookRepository
}

func (c *Context) rootPage(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "Welcome!")
}

func Configure(router *web.Router) {
	BookService.Repo = &BookRepository{
		NewMongoDataSource("mongodb://admin:admin@ds147905.mlab.com:47905/ebooks"),
	}

	router.Get("/", (*Context).rootPage)
	router.Middleware(web.LoggerMiddleware)
	router.Middleware(web.ShowErrorsMiddleware)
}
package main

import (
	"github.com/koreset/noredd-app/config/application"
	"github.com/koreset/noredd-app/app/home"
	"github.com/koreset/noredd-app/app/posts"
	"github.com/koreset/noredd-app/app/admin"
)

// main Main entrypoint for the application
func main() {
	Application := application.New()

	Application.Use(admin.New(&admin.Config{}))
	Application.Use(home.New(&home.Config{}))
	Application.Use(posts.New(&posts.Config{}))
	Application.Router.Static("/public", "./public")
	Application.Router.Static("/assets", "./assets")
	Application.Serve()
}

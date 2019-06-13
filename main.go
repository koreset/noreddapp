package main

import (
	"flag"

	"github.com/koreset/noredd-app/app/admin"
	"github.com/koreset/noredd-app/app/home"
	"github.com/koreset/noredd-app/app/posts"
	"github.com/koreset/noredd-app/config/application"
	"github.com/koreset/noredd-app/config/bindatafs"
)

// main Main entrypoint for the application
func main() {
	compileTemplate := flag.Bool("compile-templates", false, "Set this to true to compile templates to binary")
	flag.Parse()

	if *compileTemplate {
		bindatafs.AssetFS.Compile()
	} else {
		Application := application.New()

		Application.Use(admin.New(&admin.Config{}))
		Application.Use(home.New(&home.Config{}))
		Application.Use(posts.New(&posts.Config{}))
		Application.Router.Static("/public", "./public")
		Application.Router.Static("/assets", "./assets")
		Application.Serve()

	}
}

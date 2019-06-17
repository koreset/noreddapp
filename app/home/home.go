package home

import "github.com/koreset/noredd-app/config/application"

// Config this is the config for this module
type Config struct {
}

// App
type App struct {
	Config *Config
}

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

func (app App) ConfigureApplication(application *application.Application) {
	controller := &Controller{}
	application.Router.GET("/", controller.Index)
	application.Router.GET("/show", controller.Show)
	application.Router.GET("/about", controller.About)

}

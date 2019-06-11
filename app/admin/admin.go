package admin

import (
	"github.com/koreset/noredd-app/config/application"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
	"github.com/qor/admin"
)

type Config struct {
	Prefix string
}

type App struct {
	Config *Config
}

func New(cfg *Config) *App {
	if cfg == nil {
		cfg = &Config{Prefix: "/admin"}
	}

	return &App{Config: cfg}
}

func (app App) ConfigureApplication(application *application.Application) {
	Admin := application.Admin
	mux := http.NewServeMux()

	Admin.MountTo("/admin", mux)
	// Add Media Library
	application.DB.AutoMigrate(&media_library.MediaLibrary{})
	media.RegisterCallbacks(application.DB)
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	application.Router.Any("/admin/*resources", gin.WrapH(mux))

}

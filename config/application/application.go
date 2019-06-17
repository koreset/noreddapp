package application

import (
	"fmt"
	"html/template"
	"os"

	"github.com/qor/media/asset_manager"
	"github.com/qor/media/media_library"

	"github.com/Masterminds/sprig"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/koreset/gtf"
	"github.com/koreset/noredd-app/config/bindatafs"
	"github.com/koreset/noredd-app/config/db"
	"github.com/koreset/noredd-app/utils"
	"github.com/koreset/noredd-app/utils/eztemplate"
	"github.com/qor/admin"
	"github.com/qor/publish2"
)

type AppModuleInterface interface {
	ConfigureApplication(application *Application)
}

type Application struct {
	*Config
}

type Config struct {
	DB     *gorm.DB
	Admin  *admin.Admin
	Router *gin.Engine
}

func New() *Application {
	config := &Config{
		DB:     db.DB,
		Router: configureRouter(),
		Admin: admin.New(&admin.AdminConfig{
			SiteName: "No REDD in Africa Network",
			DB:       db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff),
		}),
	}
	config.Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
	config.Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
	config.Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	application := &Application{
		config,
	}

	return application
}

func (application *Application) Serve() {
	application.Router.Run("0.0.0.0:5000")
}

func (application *Application) Use(app AppModuleInterface) {
	app.ConfigureApplication(application)
}

func configureRouter() *gin.Engine {
	router := gin.Default()
	dir, _ := os.Getwd()
	fmt.Println(">>>>>>>>>>> ", dir)

	render := eztemplate.New()
	render.Ext = ".html"
	render.Layout = "layouts/application"
	render.TemplatesDir = "templates/"
	render.TemplateFuncMap = setupTemplateFuncs()

	router.HTMLRender = render.Init()
	return router

}

func setupTemplateFuncs() template.FuncMap {
	funcMaps := sprig.FuncMap()
	funcMaps["unsafeHtml"] = utils.UnsafeHtml
	funcMaps["stripSummaryTags"] = utils.StripSummaryTags
	funcMaps["displayDateString"] = utils.DisplayDateString
	funcMaps["displayDate"] = utils.DisplayDateV2
	funcMaps["truncateBody"] = utils.TruncateBody

	gtf.Inject(funcMaps)
	return funcMaps
}

package application

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/gin-gonic/gin"
	"github.com/koreset/noredd-app/config/db"
	"github.com/qor/publish2"
	"github.com/koreset/noredd-app/utils/eztemplate"
	"html/template"
	"os"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/koreset/noredd-app/utils"
	"github.com/koreset/gtf"
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
	application := &Application{
		config,
	}

	return application
}

func (application *Application) Serve() {
	application.Router.Run(":5000")
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

package posts

import (
	"github.com/koreset/noredd-app/config/application"
	"github.com/koreset/noredd-app/models/posts"
	"github.com/qor/admin"
	"github.com/qor/media/asset_manager"
)

type Config struct {
}

type App struct {
	Config *Config
}

func New(cfg *Config) *App {
	return &App{Config: cfg}
}

func (app App) ConfigureApplication(application *application.Application) {
	controller := &Controller{}
	application.DB.AutoMigrate(&posts.Post{})
	app.ConfigureAdmin(application.Admin)

	application.Router.GET("/posts", controller.Index)
	application.Router.GET("/posts/:slug", controller.GetPost)

}

func (app App) ConfigureAdmin(Admin *admin.Admin) {
	assetManager := Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
	post := Admin.AddResource(&posts.Post{}, &admin.Config{Menu: []string{"Posts Management"}})
	post.IndexAttrs("Title", "Published")
	post.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: assetManager}})
	//post.Meta(&admin.Meta{Name:"Body", Config:&admin.RichEditorConfig{}, Type:"rich_editor"})
	//post.AddProcessor(&resource.Processor{
	//	Name:"posts-processor",
	//	Handler: func(value interface{}, values *resource.MetaValues, context *qor.Context) error {
	//		if post, ok := value.(*posts.Post); ok{
	//			newId, _ := uuid.NewV4()
	//			post.ID = newId
	//		}
	//		return nil
	//	},
	//})

}

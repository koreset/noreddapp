package posts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/koreset/noredd-app/config/db"
	posts2 "github.com/koreset/noredd-app/models/posts"
	"fmt"
)

type Controller struct {
}

func (ctrl Controller) Index(c *gin.Context) {
	c.HTML(200, "posts/index", nil)
}

func (ctrl Controller) GetPost(c *gin.Context) {
	slug, found := c.Params.Get("slug")
	if !found {
		//Do something here
	}

	var post posts2.Post
	err := db.GetDB().Where("slug = ?", slug).Find(&post).Error
	if err != nil {
		//Redirect to 404 page
	}
	fmt.Println("MainImage: ", post.MainImage.URL("post_preview"))
	c.HTML(http.StatusOK, "posts/getpost", gin.H{"post": post})
}

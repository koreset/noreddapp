package posts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koreset/noredd-app/config/db"
	posts2 "github.com/koreset/noredd-app/models/posts"
)

type Controller struct {
}

func (ctrl Controller) Index(c *gin.Context) {
	var posts []posts2.Post
	err := db.GetDB().Order("publish_date DESC").Find(&posts).Error
	if err != nil {
		//Redirect to 404 page
	}

	c.HTML(200, "posts/index", gin.H{"posts": posts})
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

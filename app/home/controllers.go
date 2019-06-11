package home

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koreset/noredd-app/config/db"
	posts2 "github.com/koreset/noredd-app/models/posts"
	"github.com/pkg/errors"
)

type Controller struct {
}

type CarouselItems struct {
	Url         string
	Title       string
	Description string
	Active      bool
}

func (ctrl Controller) Index(c *gin.Context) {
	carouselItems := []CarouselItems{
		{
			Url:         "images/landgrab01.jpeg",
			Title:       "Oil company Eni plans 8.1 million hectare land grab in Africa",
			Description: "And the land grab continues...",
			Active:      true,
		},
		{
			Url:         "images/deforestation01.jpg",
			Title:       "ENI and REDD",
			Description: "And the land grab continues...",
		},
		{
			Url:         "images/deforestation02.jpg",
			Title:       "Something About NRAN",
			Description: "And the land grab continues again...",
		},
	}
	var posts []posts2.Post
	err := db.GetDB().Find(&posts).Error
	if err != nil {
		//TODO: Do something there
		fmt.Println(errors.WithStack(err))
	}

	c.HTML(http.StatusOK, "home/index", gin.H{"message": "My Message", "posts": posts, "carouselItems": carouselItems})
}

func (ctrl Controller) Show(c *gin.Context) {
	c.HTML(200, "home/show", nil)
}

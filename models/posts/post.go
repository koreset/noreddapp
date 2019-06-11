package posts

import (
	"time"

	"github.com/gosimple/slug"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
)

type Post struct {
	ID           uint                              `json:"id" db:"id" gorm:"primary_key"`
	CreatedAt    time.Time                         `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time                         `json:"updated_at" db:"updated_at"`
	Title        string                            `json:"title" db:"title"`
	Slug         string                            `json:"slug" db:"slug"`
	Body         string                            `json:"body" db:"body"`
	Published    bool                              `json:"published" db:"published"`
	PublishDate  time.Time                         `json:"publish_date" db:"publish_date"`
	MainImageURL string                            `json:"main_image_url" db:"main_image_url"`
	MainImage    media_library.MediaLibraryStorage `gorm:"type:text" sql:"size:4294967295;" media_library:"url:/content/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./assets"`
	Type         string                            `json:"post_type" db:"post_type"`
}

func (p *Post) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":        {Width: 320, Height: 320},
		"middle":       {Width: 640, Height: 640},
		"big":          {Width: 1024, Height: 720},
		"post_preview": {Width: 350, Height: 240},
		"preview":      {Width: 200, Height: 200},
	}
}

func (p *Post) BeforeCreate() (err error) {

	p.Slug = createUniqueSlug(p.Title)
	p.MainImage.Sizes = p.GetSizes()
	file, _ := p.MainImage.Base.FileHeader.Open()
	p.MainImage.Scan(file)

	//for i := range p.Images {
	//	p.Images[i].File.Sizes = p.Images[i].GetSizes()
	//	file, _ := p.Images[i].File.Base.FileHeader.Open()
	//	p.Images[i].File.Scan(file)
	//}

	return nil
}

func createUniqueSlug(title string) string {
	slugTitle := slug.Make(title)
	if len(slugTitle) > 50 {
		slugTitle = slugTitle[:50]
		if slugTitle[len(slugTitle)-1:] == "-" {
			slugTitle = slugTitle[:len(slugTitle)-1]
		}
	}
	return slugTitle
}

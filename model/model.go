package model

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// The struct to handle Album data.
type Album struct {
	Id   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name string    `json:"albumName" gorm:"references:album_name"`
}

// The struct to handle Image data.
type Image struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;"`
	AlbumId   uuid.UUID
	//Album     Album
	ImageName string `json:"imageName" gorm:"references:image_name"`
	ImageData []byte `json:"imageData" gorm:"references:image_data"`
}

func (album *Album) BeforeCreate(scope *gorm.Scope) error {
	id, _ := uuid.NewV4()
	return scope.SetColumn("Id", id)
}
func (image *Image) BeforeCreate(scope *gorm.Scope) error {
	id, _ := uuid.NewV4()
	return scope.SetColumn("Id", id)
}

package service

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"image-store-service/constants"
	"image-store-service/kafka"
	. "image-store-service/logger"
	"image-store-service/model"
	sqlite_orm "image-store-service/sqlite"
)

func AddAlbum(body string) error {
	Logger.Trace(constants.LogBeginFunc)
	albumRequest := model.Album{}
	json.Unmarshal([]byte(body), &albumRequest)
	Logger.Debug("Album Request", albumRequest)

	err := sqlite_orm.Post(&albumRequest)
	if err != nil {
		Logger.Warn(err.Error(), err)
		return err
	}
	kafka.Publish("An Album added with the name " + albumRequest.Name)

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

func AddImage(id string, body string) error {
	Logger.Trace(constants.LogBeginFunc)
	imageRequest := model.Image{}
	json.Unmarshal([]byte(body), &imageRequest)
	imageRequest.AlbumId, _ = uuid.FromString(id)

	err := sqlite_orm.Post(&imageRequest)
	if err != nil {
		Logger.Warn(err.Error(), err)
		return err
	}
	kafka.Publish("An Image added in album " + imageRequest.AlbumId.String() + " with name " + imageRequest.ImageName)

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Gets an image by id, returns nil if not exists
func GetImageByID(id string) (*model.Image, error) {
	Logger.Trace(constants.LogBeginFunc)

	image := model.Image{}

	Logger.Debug("Getting Image with id : ", id)
	err := sqlite_orm.Get(id, &image)
	if err != nil {
		Logger.Warn(err.Error(), err)
		return nil, err
	}

	Logger.Trace(constants.LogFinishFunc)
	return &image, nil
}

// Lists all albums
func ListAlbums() ([]model.Album, error) {
	Logger.Trace(constants.LogBeginFunc)

	albums := []model.Album{}

	Logger.Debug("List all albums ")
	err := sqlite_orm.List(&albums)
	if err != nil {
		Logger.Warn(err.Error(), err)
		return nil, err
	}

	Logger.Trace(constants.LogFinishFunc)
	return albums, nil
}

// Deletes Image by id
func DeleteImageById(id string) error {
	Logger.Trace(constants.LogBeginFunc)

	image := model.Image{}

	Logger.Debug("Deleting image with id : ", id)
	err := sqlite_orm.Delete(id, &image)
	if err != nil {
		Logger.Warn(err.Error(), err)
		return err
	}
	kafka.Publish("An Image deleted with id " + id)

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Deletes Album by id------------------------TBD
func DeleteAlbumById(id string) error {
	Logger.Trace(constants.LogBeginFunc)

	album := model.Album{}

	Logger.Debug("Deleting image with id : ", id)
	err := sqlite_orm.Delete(id, &album)
	err = DeleteImagesByAlbum(id)
	if err != nil {
		Logger.Warn(err.Error(), err)
		return err
	}
	kafka.Publish("An Album deleted with id " + id)
	Logger.Trace(constants.LogFinishFunc)
	return nil
}

func DeleteImagesByAlbum(id string) error {
	Logger.Trace(constants.LogBeginFunc)
	var images []model.Image
	Logger.Debug("Deleting all images in an album with id : ", id)
	convertedId, _ := uuid.FromString(id)
	err := sqlite_orm.DeleteBasedOnCondition(&images, &model.Image{AlbumId: convertedId})
	if err != nil {
		Logger.Warn(err.Error(), err)
		return err
	}
	kafka.Publish("All images deleted associated with the album id " + id)
	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Lists all images in an album by albumID
func ListImagesByAlbumID(id string) ([]model.Image, error) {
	Logger.Trace(constants.LogBeginFunc)

	var images []model.Image

	Logger.Debug("Listing all images in an album with id : ", id)
	convertedId, _ := uuid.FromString(id)
	err := sqlite_orm.ListBasedOnCondition(&images, &model.Image{AlbumId: convertedId})
	if err != nil {
		Logger.Warn(err.Error(), err)
		return nil, err
	}

	Logger.Trace(constants.LogFinishFunc)
	return images, nil
}

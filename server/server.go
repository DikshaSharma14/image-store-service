package server

import (
	"context"
	"github.com/gorilla/mux"
	"image-store-service/logger"
	"image-store-service/openapi"
	"net/http"
)

var Router *mux.Router
var ctx context.Context
var server *http.Server

func init() {
	logger.Logger.Info("Configuring Router")
	AddAlbumApiService := openapi.NewAddAlbumApiService()
	AddAlbumApiController := openapi.NewAddAlbumApiController(AddAlbumApiService)

	AddImageApiService := openapi.NewAddImageApiService()
	AddImageApiController := openapi.NewAddImageApiController(AddImageApiService)

	DeleteAlbumApiService := openapi.NewDeleteAlbumApiService()
	DeleteAlbumApiController := openapi.NewDeleteAlbumApiController(DeleteAlbumApiService)

	DeleteImageByIdApiService := openapi.NewDeleteImageByIdApiService()
	DeleteImageByIdApiController := openapi.NewDeleteImageByIdApiController(DeleteImageByIdApiService)

	DeleteImagesByAlbumApiService := openapi.NewDeleteImagesByAlbumApiService()
	DeleteImagesByAlbumApiController := openapi.NewDeleteImagesByAlbumApiController(DeleteImagesByAlbumApiService)

	GetAlbumsApiService := openapi.NewGetAlbumsApiService()
	GetAlbumsApiController := openapi.NewGetAlbumsApiController(GetAlbumsApiService)

	GetImageByIdApiService := openapi.NewGetImageByIdApiService()
	GetImageByIdApiController := openapi.NewGetImageByIdApiController(GetImageByIdApiService)

	GetImagesByAlbumApiService := openapi.NewGetImagesByAlbumApiService()
	GetImagesByAlbumApiController := openapi.NewGetImagesByAlbumApiController(GetImagesByAlbumApiService)

	Router = openapi.NewRouter(AddAlbumApiController, AddImageApiController, DeleteAlbumApiController, DeleteImageByIdApiController, DeleteImagesByAlbumApiController, GetAlbumsApiController, GetImageByIdApiController, GetImagesByAlbumApiController)

	logger.Logger.Info("Router configured")

}

func Start() {
	ctx, _ = context.WithCancel(context.Background())
	port := "9000"
	server = &http.Server{Addr: ":" + port}
	http.Handle("/", Router)
	logger.Logger.Info("Starting HTTP Server on port : ", port)
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			logger.Logger.Critical(err)
			panic(err)
		}
	}
}

func Stop() {
	logger.Logger.Info("Server shutdown")
	server.Shutdown(ctx)
}

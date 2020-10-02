/*
 * Image Store
 *
 * Image store service provides Albums and Image management operaions
 *
 * API version: 1.0.0
 * Contact: diksha.may14@outlook.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"image-store-service/constants"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A GetImagesByAlbumApiController binds http requests to an api service and writes the service results to the http response
type GetImagesByAlbumApiController struct {
	service GetImagesByAlbumApiServicer
}

// NewGetImagesByAlbumApiController creates a default api controller
func NewGetImagesByAlbumApiController(s GetImagesByAlbumApiServicer) Router {
	return &GetImagesByAlbumApiController{service: s}
}

// Routes returns all of the api route for the GetImagesByAlbumApiController
func (c *GetImagesByAlbumApiController) Routes() Routes {
	return Routes{
		{
			"GetImagesByAlbum",
			strings.ToUpper("Get"),
			"/albums/{id}/images",
			c.GetImagesByAlbum,
		},
	}
}

// GetImagesByAlbum - Get all the images associated with an album of a given id
func (c *GetImagesByAlbumApiController) GetImagesByAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := c.service.GetImagesByAlbum(id)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrorDBRecordNotFound) || strings.Contains(err.Error(), constants.ErrorDBNoSuchTable) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}
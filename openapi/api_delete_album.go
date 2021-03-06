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
	"github.com/gorilla/mux"
	"net/http"
)

// A DeleteAlbumApiController binds http requests to an api service and writes the service results to the http response
type DeleteAlbumApiController struct {
	service DeleteAlbumApiServicer
}

// NewDeleteAlbumApiController creates a default api controller
func NewDeleteAlbumApiController(s DeleteAlbumApiServicer) Router {
	return &DeleteAlbumApiController{service: s}
}

// Routes returns all of the api route for the DeleteAlbumApiController
func (c *DeleteAlbumApiController) Routes() Routes {
	return Routes{
		{
			"DeleteAlbum",
			http.MethodDelete,
			"/albums/{id}",
			c.DeleteAlbum,
		},
	}
}

// DeleteAlbum - delete an album with the the given id
func (c *DeleteAlbumApiController) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := c.service.DeleteAlbum(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	EncodeJSONResponse(result, http.StatusNoContent, w)
}

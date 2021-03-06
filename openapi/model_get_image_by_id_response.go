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

type GetImageByIdResponse struct {
	ImageId string `json:"imageId,omitempty"`

	AlbumId string `json:"albumId,omitempty"`

	AlbumName string `json:"albumName,omitempty"`

	ImageName *string `json:"imageName,omitempty"`

	ImageData string `json:"imageData,omitempty"`
}

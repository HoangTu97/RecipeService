package controller

import (
	"Food/dto/response"
	"Food/helpers/converter"
	fileUtil "Food/helpers/file"
	"Food/service"

	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

type Image interface {
	Upload(c *gin.Context)
	Download(c *gin.Context)
	FileDisplay(c *gin.Context)
}

type image struct {
}

func NewImage() Image {
	return &image{}
}

// Upload upload image
// @Summary Upload
// @Tags PublicImage
// @Accept mpfd
// @Param file formData file true "Body with file"
// @Success 200 {object} response.APIResponseDTO{data=string} "desc"
// @Router /api/public/image/upload [post]
func (r *image) Upload(c *gin.Context) {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		response.CreateErrorResponse(c, err.Error())
		return
	}

	// baseFilename := filepath.Base(file.Filename)
	ext := filepath.Ext(file.Filename)

	filename := service.GenFileBaseFileName(ext)
	path := service.GetFilePath(filename)

	_ = fileUtil.MkDir(service.GetFilePathDir(filename))
	if err := c.SaveUploadedFile(file, path); err != nil {
		response.CreateErrorResponse(c, err.Error())
		return
	}

	response.CreateSuccesResponse(c, filename)
}

// Download download image
// @Summary Download
// @Tags PublicImage
// @Produce octet-stream
// @Param id path string true "Image id"
// @Router /api/public/image/{id}/download [get]
func (r *image) Download(c *gin.Context) {
	filename := converter.MustString(c.Param("id"))

	filePath := service.GetFilePath(filename)

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	// fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filePath)
}

// Display display image
// @Summary Display
// @Tags PublicImage
// @Produce octet-stream
// @Param id path string true "Image id"
// @Router /api/public/image/{id} [get]
func (r *image) FileDisplay(c *gin.Context) {
	filename := converter.MustString(c.Param("id"))

	filePath := service.GetFilePath(filename)

	b, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		response.CreateErrorResponse(c, err.Error())
		return
	}

	mime := mimetype.Detect(b)

	c.Data(200, mime.String(), b)
}

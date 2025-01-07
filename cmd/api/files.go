package main

import (
	"MyDrive/internal/utils"
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type DownloadFileRequestBody struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

const maxFileSize = 10 << 20

// downloadFileHandler godoc
//
//	@Summary		Downloads a file at the specified path.
//	@Description	Downloads a file at the specified path.
//	@Tags			myfiles
//	@Accept			json
//	@Produce		json
//
//	@Param			downloadFileRequestBody	body		DownloadFileRequestBody	true	"Metadata about the file to download"
//
//	@Success		200						{string}	string					"Downloaded file successfully!"
//	@Failure		400						{object}	error					"Bad request"
//	@Failure		404						{object}	error					"File not found/Internal server error"
//	@Security		ApiKeyAuth
//	@Router			/mydrive/myfiles [get]
func (app *application) downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	var downloadRequest DownloadFileRequestBody
	if err := readJSON(w, r, &downloadRequest); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	filePath := app.config.drive.root + downloadRequest.Path

	app.logger.Infof("Downloading file at path: %s", filePath)

	// validate request body
	// check if file exist
	fileStats, err := utils.FileStats(filePath)
	if err != nil {
		app.fileNotFoundResponse(w, r, err)
		return
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			app.logger.Warnf("file close error: %v", err)
			return
		}
	}(file)

	// write header
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename:"+downloadRequest.Name)
	w.Header().Set("Content-Length", strconv.FormatInt(fileStats.Size(), 10))

	// send file through SFT protocol
	http.ServeContent(w, r, downloadRequest.Name, fileStats.ModTime(), file)
}

// uploadFileHandler godoc
//
//	@Summary		Uploads a file at the specified path.
//	@Description	Uploads a file at the specified path.
//	@Tags			myfiles
//	@Accept			json
//	@Produce		json
//
// @Param			path	path		string	true	"File path"
// @Param			name	path		string	true	"File name"
// @Param			size	path		int	true	"File size"
// @Success		200						{string}	string					"Uploaded file successfully!"
// @Failure		400						{object}	error					"Bad request"
// @Failure		404						{object}	error					"Internal Server Error"
// @Failure		413						{object}	error					"File too large"
// @Security		ApiKeyAuth
// @Router			/mydrive/myfiles [post]
func (app *application) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request URL
	path := r.URL.Query().Get("path")
	_ = r.URL.Query().Get("name")
	sizeStr := r.URL.Query().Get("size")

	if path == "" || sizeStr == "" {
		app.badRequestResponse(w, r, errors.New("missing path or size"))
		return
	}

	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if size > maxFileSize {
		app.fileTooLargeResponse(w, r, errors.New("file too large"))
		return
	}

	// Limit file size to 10MB
	err = r.ParseMultipartForm(maxFileSize)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Parse the form
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			app.logger.Warnf("file close error: %v", err)
			return
		}
	}(file)

	// File size is bigger than 10 MB
	if handler.Size > maxFileSize {
		app.fileTooLargeResponse(w, r, errors.New("file too large"))
		return
	}

	// TODO() Implement file validation - #9

	filePath := app.config.drive.root + path

	// Create the file
	err = app.repo.FilesSystem.CreateFile(filePath)

	// Save the file
	fileContents, err := utils.ReadFile(file)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.repo.FilesSystem.UpdateFile(filePath, fileContents, 0)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Return a success message
	if err := jsonResponse(w, http.StatusOK, "File uploaded successfully!"); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

package main

import (
	"MyDrive/internal/repo"
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUploadFileHandler_InMemoryFileSystem(t *testing.T) {
	mockInMemoryRepoBuilder := repo.NewMockInMemoryRepositoryBuilder()
	mockInMemoryRepoBuilder.SetFileSystem()
	mockInMemoryRepoBuilder.SetUsers()

	app := newTestApplication(t, config{}, mockInMemoryRepoBuilder)
	mux := app.mount()

	t.Run("should upload file successfully", func(t *testing.T) {
		testingFileContent := "File uploaded successfully!"
		writer, body := prepareRequestForUpload(t, []byte(testingFileContent), "/testfile.txt", "1024", "testfile.txt")

		// Create the request with the multipart body
		req, err := http.NewRequest(http.MethodPost, "/v1/myfiles/upload?path=/testfile.txt&size=1024&name=testfile.txt", body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Perform the request
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		// Validate response
		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}

func TestUploadFileHandler_OnDiskFileSystem(t *testing.T) {
	mockInMemoryRepoBuilder := repo.NewMockOnDiskRepositoryBuilder()
	mockInMemoryRepoBuilder.SetFileSystem()
	mockInMemoryRepoBuilder.SetUsers()

	app := newTestApplication(t, config{}, mockInMemoryRepoBuilder)
	_ = app.mount()

	t.Run("should upload file successfully", func(t *testing.T) {
		// TODO(): Implement this more complex text
	})
}

func TestDownloadFileHandler_Success(t *testing.T) {

}

func prepareRequestForUpload(t *testing.T, content []byte, path, size, name string) (*multipart.Writer, *bytes.Buffer) {
	// Create a multipart request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add form fields
	err := writer.WriteField("path", "/testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteField("size", "1024")
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteField("name", "testfile.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Add a file
	fileWriter, err := writer.CreateFormFile("myFile", "testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = fileWriter.Write(content)
	if err != nil {
		t.Fatal(err)
	}

	// Close the writer to finalize the body
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}
	return writer, body
}

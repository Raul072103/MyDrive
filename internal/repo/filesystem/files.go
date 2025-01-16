package filesystem

import (
	"MyDrive/internal/utils"
	"errors"
	"os"
)

var (
	ErrFileAlreadyExists = errors.New("file already exists")
	ErrFileNotFound      = errors.New("file not found")
)

type FileRepo struct {
}

func New() *FileRepo {
	return &FileRepo{}
}

type File struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
	Size  int64  `json:"size"`
}

// CreateFile creates a file at the specified path.
// It doesn't create a new file if there is already an existing one.
func (fr *FileRepo) CreateFile(path string) (err error) {
	fileExists, err := utils.FileExists(path)
	if fileExists {
		return ErrFileAlreadyExists
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(err, file.Close())
	}()

	return err
}

// DeleteFile deletes a file at the specified path.
func (fr *FileRepo) DeleteFile(path string) (err error) {
	fileExists, err := utils.FileExists(path)
	if err != nil {
		return err
	}

	if fileExists {
		err = os.Remove(path)
		return err
	} else {
		return ErrFileNotFound
	}
}

// UpdateFile updates the contents of the file.
func (fr *FileRepo) UpdateFile(path string, content []byte, updateAt int64) (err error) {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(err, file.Close())
	}()

	_, err = file.WriteAt(content, updateAt)
	if err != nil {
		return err
	}

	return err
}

// ReadFile reads the contents of the file, if it exists, at the given path and returns the content of that file.
// TODO() handle bigger filesystem.
func (fr *FileRepo) ReadFile(path string) ([]byte, error) {
	fileExists, err := utils.FileExists(path)
	if err != nil {
		return nil, err
	}

	if fileExists {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		return data, nil
	} else {
		return nil, ErrFileNotFound
	}
}

func (fr *FileRepo) ListFiles(path string) ([]File, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var fileNames []File
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, err
		}

		fileNames = append(fileNames, File{
			Name:  file.Name(),
			IsDir: file.IsDir(),
			Size:  fileInfo.Size(),
		})
	}

	return fileNames, nil
}

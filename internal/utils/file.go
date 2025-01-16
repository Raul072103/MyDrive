package utils

import (
	"io"
	"os"
)

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func FileStats(path string) (os.FileInfo, error) {
	stats, err := os.Stat(path)
	if err == nil {
		return stats, nil
	}

	return stats, err
}

func ReadFile(file io.Reader) ([]byte, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//func IsValidFileType(file []byte) bool {
//	fileType := http.DetectContentType(file)
//	return strings.HasPrefix(fileType, "image/") // Only allow images
//}

package filesystem

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

type testFile struct {
	path      string
	operation string
	args      []any
}

// TestFileRepo_CreateFile_TestingDirectory tests if the file is successfully created in that directory
func TestFileRepo_CreateFile_TestingDirectory(t *testing.T) {
	path := createTestsDirectory + "\\create_test1.txt"

	err := FR.CreateFile(path)
	if err != nil {
		t.Errorf("File wasn't created successfully!")
	}

	err = testIfFileExists(path)
	if err != nil {
		t.Errorf("File doesn't exist after it was created!")
	}
}

// TestFileRepo_CreateFile_PathInWrongFormat tests if an error is thrown when the path doesn't exist
func TestFileRepo_CreateFile_PathInWrongFormat(t *testing.T) {
	paths := []string{"\\\\invalid_path", "invalid_path////\\", "\\\nilfasf@@...124124.txtxt"}

	for _, path := range paths {
		err := FR.CreateFile(path)
		if err == nil {
			t.Errorf("File %s was created even though it shouldn't have!", path)
		}
	}
}

// TestFileRepo_CreateFile_AlreadyExistingFile tests if an error is thrown when the file already exist
func TestFileRepo_CreateFile_AlreadyExistingFile(t *testing.T) {
	path := createTestsDirectory + "\\create_test2.txt"

	err := FR.CreateFile(path)
	if err == nil {
		t.Errorf("File was created when it shouldn't have been!")
	}
}

func TestFileRepo_ReadFile_ReadsCorrectly(t *testing.T) {
	path := readTestsDirectory + "\\read_test1.txt"

	content, err := FR.ReadFile(path)
	if err != nil {
		t.Errorf("File %s couldn't be read!", path)
	}

	if !bytes.Equal(content, fileContentReadTest1) {
		t.Errorf("File contents don't match!")
	}
}

func TestFileRepo_ReadFile_FileDoesntExist(t *testing.T) {
	path := readTestsDirectory + "\\nonexistent_file.txt"

	_, err := FR.ReadFile(path)
	if err == nil {
		t.Errorf("Tried to read even if the file doesn't exist")
	}
}

// TODO(): tests for UNIX systems
//func TestFileRepo_ReadFile_NoPermissionsToAccessTheDirectory(t *testing.T) {
//	path := readTestsDirectory + "\\no_E_permissions" + "\\read_test2.txt"
//
//	_, err := FR.ReadFile(path)
//	if err == nil {
//		t.Errorf("Tried to access a file from a directory with no execute rights")
//	}
//}
//
//func TestFileRepo_ReadFile_NoPermissionsToRead(t *testing.T) {
//	path := readTestsDirectory + "\\no_R_permissions" + "\\read_test3.txt"
//
//	_, err := FR.ReadFile(path)
//	if err == nil {
//		t.Errorf("Tried to access a file from a directory with no read rights")
//	}
//}

func TestFileRepo_DeleteFile_DeletesCorrectly(t *testing.T) {
	path := deleteTestsDirectory + "\\delete_test1.txt"

	err := FR.DeleteFile(path)
	if err != nil {
		t.Errorf("Failed to delete file!")
	}
}

func TestFileRepo_DeleteFile_FileDoesntExist(t *testing.T) {
	path := deleteTestsDirectory + "\\delete_nonexistent.txt"

	err := FR.DeleteFile(path)
	if err == nil {
		t.Errorf("Tried to delete non-existent file!")
	}
}

// TODO(): tests for UNIX systems
//func TestFileRepo_DeleteFile_NoPermissionsToAccessTheDirectory(t *testing.T) {}

const updateContentSize = 100

type updateContentOffset struct {
	Content []byte
	Offset  int64
}

var updatesCorrectlyTestCases = []updateContentOffset{
	{generateRandomByteArray(updateContentSize), 0},
	{generateRandomByteArray(updateContentSize), 1000},
	{generateRandomByteArray(updateContentSize), 2000},
	{generateRandomByteArray(updateContentSize), 10000},
}

func TestFileRepo_UpdateFile_UpdatesCorrectly(t *testing.T) {
	path := updateTestsDirectory + "\\update_test1.txt"

	for _, updateTestCase := range updatesCorrectlyTestCases {
		content := updateTestCase.Content
		offset := updateTestCase.Offset

		err := FR.UpdateFile(path, content, offset)
		if err == nil {
			t.Errorf("Tried to delete non-existent file!")
		}
	}
}

func TestFileRepo_UpdateFile_OffsetBiggerThanFileSize(t *testing.T) {

}

func TestFileRepo_UpdateFile_OffsetIsNegative(t *testing.T) {

}

func TestFileRepo_UpdateFile_FileDoesntExist(t *testing.T) {

}

// TODO(): tests for UNIX systems
//func TestFileRepo_UpdateFile_NoPermissionsToWrite(t *testing.T) {}

func testIfFileExists(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(err, file.Close())
	}()
	return err
}

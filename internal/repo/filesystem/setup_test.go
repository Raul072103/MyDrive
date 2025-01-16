package filesystem

import (
	"MyDrive/internal/utils"
	"crypto/rand"
	"os"
	"testing"
)

var FR = New()

func TestMain(m *testing.M) {
	// Setup code here (runs before all tests)
	setup()

	// Run the tests
	code := m.Run()

	// Teardown code here (runs after all tests)
	teardown()

	// Exit with the code from m.Run()
	os.Exit(code)
}

// TODO() Make this independent on the system the tests are run
const testDirPath = "C:\\Users\\raula\\Desktop\\facultate\\anul 3 sem 1\\Software Engineering\\Golang\\MyDrive\\testfiles"

func setup() {
	createTestsSetup()
	readTestsSetup()
	deleteTestsSetup()
}

func teardown() {
	createTestsTeardown()
	readTestsTeardown()
	deleteTestsTeardown()
}

const createTestsDirectory = testDirPath + "\\create_operations_tests"

func createTestsSetup() {
	createTestDirectoryIfNotExistent(createTestsDirectory)

	// Setup for TestFileRepo_CreateFile_AlreadyExistingFile
	createExistingFileTest := createTestsDirectory + "\\create_test2.txt"
	createTestFileIfNonexistent(createExistingFileTest, []byte("1 2 3"))

}

func createTestsTeardown() {
	// Teardown for TestFileRepo_CreateFile_TestingDirectory
	createCorrectFileTest := createTestsDirectory + "\\create_test1.txt"
	deleteTestFile(createCorrectFileTest)
}

const readTestsDirectory = testDirPath + "\\read_operations_tests"
const fileSize = 16_000

var fileContentReadTest1 = generateRandomByteArray(fileSize)

func readTestsSetup() {
	createTestDirectoryIfNotExistent(readTestsDirectory)

	// Setup for TestFileRepo_ReadFile_ReadsCorrectly
	correctFile := readTestsDirectory + "\\read_test1.txt"
	createTestFile(correctFile, fileContentReadTest1)

	// Setup for TestFileRepo_ReadFile_NoPermissionsToAccessTheDirectory
	dirPermissionsNoExecute := readTestsDirectory + "\\no_E_permissions"
	noExecPermFile := dirPermissionsNoExecute + "\\read_test2.txt"
	createTestDirectoryWithPermissions(dirPermissionsNoExecute, 0222)
	createTestFileIfNonexistent(noExecPermFile, nil)

	// Setup for TestFileRepo_ReadFile_NoPermissionsToRead
	dirPermissionsNoRead := readTestsDirectory + "\\no_R_permissions"
	noReadPermFile := dirPermissionsNoRead + "\\read_test3.txt"
	createTestDirectoryWithPermissions(dirPermissionsNoRead, 0222)
	createTestFileIfNonexistent(noReadPermFile, nil)
}

func readTestsTeardown() {
	// TODO(): Add any teardown operations if necessary for READ tests
}

const deleteTestsDirectory = testDirPath + "\\delete_operations_tests"

func deleteTestsSetup() {
	createTestDirectoryIfNotExistent(deleteTestsDirectory)

	// Setup for TestFileRepo_DeleteFile_DeletesCorrectly
	deleteFile := deleteTestsDirectory + "\\delete_test1.txt"
	createTestFileIfNonexistent(deleteFile, nil)
}

func deleteTestsTeardown() {
	// TODO(): Add any teardown operations if necessary for DELETE tests
}

const updateTestsDirectory = testDirPath + "\\update_operations_tests"
const updateFileTest = 16_000

var updateTestFile1Before = generateRandomByteArray(updateFileTest)

func updateTestsSetup() {
	createTestDirectoryIfNotExistent(updateTestsDirectory)

	// Setup for TestFileRepo_UpdateFile_UpdatesCorrectly
	updateFile := updateTestsDirectory + "\\update_test1.txt"
	createTestFile(updateFile, updateTestFile1Before)

	// Setup for TestFileRepo_UpdateFile_OffsetBiggerThanFileSize
	updateFile2 := updateTestsDirectory + "\\update_test2.txt"
	createTestFileIfNonexistent(updateFile2, updateTestFile1Before)

	// Setup for TestFileRepo_UpdateFile_OffsetIsNegative
	updateFile3 := updateTestsDirectory + "\\update_test3.txt"
	createTestFileIfNonexistent(updateFile3, updateTestFile1Before)
}

func updateTestsTeardown() {
	// TODO(): Add any teardown operations if necessary for DELETE tests
}

func createTestFileIfNonexistent(path string, content []byte) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		createTestFile(path, content)
	}
}

func createTestFile(path string, content []byte) {
	file, err := os.Create(path)
	if err != nil {
		panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
	}

	if content != nil {
		_, err = file.Write(content)
		if err != nil {
			panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
		}
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
		}
	}()
}

func deleteTestFile(path string) {
	fileExists, err := utils.FileExists(path)
	if err != nil {
		panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
	}

	if fileExists {
		err := os.Remove(path)
		if err != nil {
			panic("SETUP OF THE TESTS FAILED: \n" + err.Error())
		}
	}
}

func createTestDirectoryIfNotExistent(path string) {
	createTestDirectoryWithPermissions(path, 0777)
}

func createTestDirectoryWithPermissions(path string, perm os.FileMode) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, perm)

		if err != nil {
			panic("SETUP FOR TESTS FAILED: \n" + err.Error())
		}
	}
}

// createRandomByteArray generates an array of the given size with random values.
// Helper function for setting up the tests for the files.
func generateRandomByteArray(size int) []byte {
	byteArray := make([]byte, size)

	_, err := rand.Read(byteArray)
	if err != nil {
		panic("SETUP FOR TESTS FAILED: \n" + err.Error())
	}

	return byteArray
}

package repo

import "MyDrive/internal/repo/filesystem"

type IMockRepositoryBuilder interface {
	SetUsers()
	SetFileSystem()
	GetRepository() *Repository
}

type MockInMemoryRepositoryBuilder struct {
	Users       Users
	FilesSystem FileSystem
}

func NewMockInMemoryRepositoryBuilder() *MockInMemoryRepositoryBuilder {
	return &MockInMemoryRepositoryBuilder{}
}

func (m *MockInMemoryRepositoryBuilder) SetUsers() {
	m.Users = NewUsersMock()
}

func (m *MockInMemoryRepositoryBuilder) SetFileSystem() {
	m.FilesSystem = filesystem.NewInMemoryFileSystemMock()
}

func (m *MockInMemoryRepositoryBuilder) GetRepository() *Repository {
	return &Repository{
		Users:       m.Users,
		FilesSystem: m.FilesSystem,
	}
}

type MockOnDiskRepositoryBuilder struct {
	Users       Users
	FilesSystem FileSystem
}

func NewMockOnDiskRepositoryBuilder() *MockOnDiskRepositoryBuilder {
	return &MockOnDiskRepositoryBuilder{}
}

func (m *MockOnDiskRepositoryBuilder) SetUsers() {
	m.Users = NewUsersMock()
}

func (m *MockOnDiskRepositoryBuilder) SetFileSystem() {
	m.FilesSystem = filesystem.NewOnDiskFileSystemMock()
}

func (m *MockOnDiskRepositoryBuilder) GetRepository() *Repository {
	return &Repository{
		Users:       m.Users,
		FilesSystem: m.FilesSystem,
	}
}

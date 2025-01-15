package filesystem

type InMemoryFileSystemMock struct {
	fileMap map[string][]byte
}

func (u *InMemoryFileSystemMock) CreateFile(path string) (err error) {
	val, ok := u.fileMap[path]
	if ok {
		return ErrFileAlreadyExists
	}

	u.fileMap[path] = val
	return nil
}

func (u *InMemoryFileSystemMock) DeleteFile(path string) (err error) {
	_, ok := u.fileMap[path]
	if !ok {
		return ErrFileNotFound
	}

	delete(u.fileMap, path)
	return nil
}

func (u *InMemoryFileSystemMock) UpdateFile(path string, content []byte, updateAt int64) (err error) {
	u.fileMap[path] = content
	return nil
}

func (u *InMemoryFileSystemMock) ReadFile(path string) ([]byte, error) {
	_, ok := u.fileMap[path]
	if !ok {
		return nil, ErrFileNotFound
	}

	return u.fileMap[path], nil
}

func (u *InMemoryFileSystemMock) ListFiles(path string) ([]File, error) {
	return nil, nil
}

type OnDiskFileSystemMock struct {
}

func (u *OnDiskFileSystemMock) CreateFile(path string) (err error) {
	return nil
}

func (u *OnDiskFileSystemMock) DeleteFile(path string) (err error) {
	return nil
}

func (u *OnDiskFileSystemMock) UpdateFile(path string, content []byte, updateAt int64) (err error) {
	return nil
}

func (u *OnDiskFileSystemMock) ReadFile(path string) ([]byte, error) {
	return nil, nil
}

func (u *OnDiskFileSystemMock) ListFiles(path string) ([]File, error) {
	return nil, nil
}

func NewInMemoryFileSystemMock() *InMemoryFileSystemMock {
	return &InMemoryFileSystemMock{make(map[string][]byte)}
}

func NewOnDiskFileSystemMock() *OnDiskFileSystemMock {
	return &OnDiskFileSystemMock{}
}

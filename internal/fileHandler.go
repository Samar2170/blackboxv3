package internal

import (
	"bytes"
	"os"
	"sync"

	"github.com/google/uuid"
)

type FileInfo struct {
	FileName string
	FileSize uint32
	FilePath string
	FileType string
}
type LocalFileStore struct {
	mutex  sync.Mutex
	folder string
	files  map[string]FileInfo
}

func NewLocalFileStore(folder string) *LocalFileStore {
	return &LocalFileStore{
		folder: folder,
		files:  make(map[string]FileInfo),
	}
}

func (s *LocalFileStore) SaveFile(channelID uint, fileName string, file bytes.Buffer) (string, error) {
	fileId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	filePath := s.folder + "/" + fileId.String()
	newFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	_, err = newFile.Write(file.Bytes())
	if err != nil {
		return "", err
	}
	s.mutex.Lock()
	s.files[fileId.String()] = FileInfo{
		FileName: fileName,
		FileSize: uint32(file.Len()),
		FilePath: filePath,
		FileType: "image",
	}
	s.mutex.Unlock()
	return fileId.String(), nil
}

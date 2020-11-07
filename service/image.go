package service

import (
	"fmt"
	"hash/fnv"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

type Image interface {
	GenFileBaseFileName(extension string) string
	GetFilePath(fileName string) string
	GetFilePathDir(fileName string) string
}

type image struct {
	location string
}

func NewImage() Image {
	return &image{
		location: "./data/Images",
	}
}

func (s *image) GenFileBaseFileName(extension string) string {
	return uuid.NewV4().String() + extension
}

func (s *image) GetFilePath(fileName string) string {
	hash := s.hash(fileName)
	var mask uint32 = 255
	firstDir := hash & mask
	secondFir := (hash >> 8) & mask
	return filepath.Join(s.location, fmt.Sprintf("%02x", firstDir), fmt.Sprintf("%02x", secondFir), fileName)
}

func (s *image) GetFilePathDir(fileName string) string {
	return filepath.Dir(s.GetFilePath(fileName))
}

func (s *image) hash(str string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(str))
	return h.Sum32()
}

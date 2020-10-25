package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"hash/fnv"
	"path/filepath"
)

var (
	location = "./data/Images"
)

func GenFileBaseFileName(extension string) string {
	return uuid.NewV4().String() + extension
}

func GetFilePath(fileName string) string {
	hash := hash(fileName)
	var mask uint32 = 255
	firstDir := hash & mask
	secondFir := (hash >> 8) & mask
	return filepath.Join(location, fmt.Sprintf("%02x", firstDir), fmt.Sprintf("%02x", secondFir), fileName)
}

func GetFilePathDir(fileName string) string {
	return filepath.Dir(GetFilePath(fileName))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

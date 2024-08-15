package storage_manager

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func InitPath(pathParts ...string) (path string, err error) {

	path = filepath.Join(pathParts...)
	if err = MkDir(path); err != nil {
		err = errors.Wrapf(err, "directory %s was not maded", path)
	}

	return
}

func MkDir(path string) (err error) {

	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		err = errors.Wrap(err, "failed to make directory")
	}

	return
}

func Create(pathParts ...string) (dst *os.File, err error) {

	path := filepath.Join(pathParts...)
	dst, err = os.Create(path)
	if err != nil {
		err = errors.Wrap(err, "failed to create directory")
	}

	return
}

func DeletePath(pathParts ...string) (path string, err error) {

	path = filepath.Join(pathParts...)
	if err = os.RemoveAll(path); err != nil {
		err = errors.Wrap(err, "failed to delete directory")
	}

	return
}

func Find(pathParts ...string) (path string, err error) {

	path = filepath.Join(pathParts...)
	// Проверка существования файла
	if _, err = os.Stat(path); err == nil {
		return
	} else if os.IsNotExist(err) {
		err = errors.Wrapf(err, "path %s was not found", path)
	} else {
		err = errors.Wrap(err, "error occurred while checking the path")
	}

	return
}

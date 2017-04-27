package loggo_writer

import (
	"os"
	"path"
)

func EnsureDir(p string) error {
	path := path.Dir(p)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModeDir|0700); err != nil {
			return err
		}
	}
	return nil
}

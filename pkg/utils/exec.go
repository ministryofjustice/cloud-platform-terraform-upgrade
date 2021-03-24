package utils

import (
	"io/fs"
	"os/exec"
	"path/filepath"
)

// Execute a command in each directory in passed argument dir.
func Execute(dir, command string) error {
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			command := exec.Command("/bin/sh", "-c", command)
			command.Dir = filepath.Dir(path) + "/" + info.Name()
			err := command.Run()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

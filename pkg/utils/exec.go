package utils

import (
	"io/fs"
	"path/filepath"

	execute "github.com/alexellis/go-execute/pkg/v1"
)

func Execute(dir, command string) error {
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			cmd := execute.ExecTask{
				Command:     command,
				StreamStdio: true,
			}

			_, err := cmd.Execute()
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

package fileutil

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"wiiki_server/common/wiikierr"
)

func GetBytes(path string) ([]byte, error) {

	fp, err := os.Open(path)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedOpenFile, "path=%s", path)
	}
	defer fp.Close()

	b, err := io.ReadAll(fp)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedReadFile, "path=%s", path)
	}

	return b, nil
}

func WalkDir(rootPath string) ([]string, error) {

	var fileList []string
	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return wiikierr.Bind(err, wiikierr.FailedWalkDir, "root=%s, path=%s", rootPath, path)
		}
		if d.IsDir() {
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileList, nil
}

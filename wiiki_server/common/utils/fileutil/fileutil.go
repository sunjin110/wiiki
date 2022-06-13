package fileutil

import (
	"io"
	"os"
	"wiiki_server/common/wiikierr"
)

func GetBytes(path string) ([]byte, error) {

	fp, err := os.Open(path)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedOpenFile, "path=%s", path)
	}
	defer fp.Close()

	// buf := &bufio
	b, err := io.ReadAll(fp)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedReadFile, "path=%s", path)
	}

	return b, nil
}

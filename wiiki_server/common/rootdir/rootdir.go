package rootdir

import (
	"path"
	"path/filepath"
	"runtime"
)

func Dir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Join(filepath.Dir(d), "..")
}

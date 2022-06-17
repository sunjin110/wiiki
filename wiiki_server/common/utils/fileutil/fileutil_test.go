package fileutil_test

import (
	"testing"
	"wiiki_server/common/utils/fileutil"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 -timeout 30s -run ^TestFileUtil$ wiiki_server/common/utils/fileutil

func TestFileUtil(t *testing.T) {
	t.Run("GetBytes", func(t *testing.T) {
		t.Run("When specified correct path. then return file data.", func(t *testing.T) {
			b, err := fileutil.GetBytes("./testdata/test_file.text")
			assert.Nil(t, err)
			assert.Equal(t, string(b), "hello_world\nline2")
		})
		t.Run("When specified incorrect path. then return error.", func(t *testing.T) {
			_, err := fileutil.GetBytes("./incorrect/file/path.text")
			assert.EqualError(t, err, "open ./incorrect/file/path.text: no such file or directory")
		})
	})
}

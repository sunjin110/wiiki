package testtool_test

import (
	"testing"
	"wiiki_server/common/testtool"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 -timeout 30s -run ^TestTestTool$ wiiki_server/common/testtool

func TestTestTool(t *testing.T) {

	t.Run("TestTool", func(t *testing.T) {
		t.Run("TestRootDir", func(t *testing.T) {
			rootDir := testtool.TestRootDir()
			assert.NotEmpty(t, rootDir)
		})
		t.Run("Config", func(t *testing.T) {
			conf := testtool.Config()
			assert.NotNil(t, conf)
		})
	})

}

package rootdir_test

import (
	"log"
	"testing"
	"wiiki_server/common/rootdir"
)

// go test -v -count=1 -timeout 30s -run ^TestRootDir$ wiiki_server/common/rootdir

func TestRootDir(t *testing.T) {

	t.Run("RootDir:test", func(t *testing.T) {

		str := rootdir.Dir()
		log.Println("str is ", str)
	})

}

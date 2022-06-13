package tomlutil_test

import (
	"testing"
	"wiiki_server/common/utils/tomlutil"

	"github.com/stretchr/testify/assert"
)

type TestTomlStruct struct {
	Str        string
	Num        int
	Num32      int32
	Num64      int64
	Map        map[string]string
	List       []string
	NameChange string `toml:"changed_name"`
}

// go test -v -count=1 -timeout 30s -run ^TestTomlUtil$ wiiki_server/common/utils/tomlutil

func TestTomlUtil(t *testing.T) {

	t.Run("TomlUtil", func(t *testing.T) {
		t.Run("Marshal", func(t *testing.T) {

			obj := &TestTomlStruct{
				Str:   "test_str",
				Num:   1,
				Num32: 2,
				Num64: 3,
				Map: map[string]string{
					"hello":  "world",
					"myname": "sunjin",
				},
				List:       []string{"a", "b", "c", "d"},
				NameChange: "change_change",
			}

			b, err := tomlutil.Marshal(obj)
			assert.Nil(t, err)
			assert.Equal(t, string(b), "Str = 'test_str'\nNum = 1\nNum32 = 2\nNum64 = 3\nList = ['a', 'b', 'c', 'd']\nchanged_name = 'change_change'\n[Map]\nhello = 'world'\nmyname = 'sunjin'\n\n")
		})

		t.Run("Unmarshal", func(t *testing.T) {

			str := "Str = 'test_str'\nNum = 1\nNum32 = 2\nNum64 = 3\nList = ['a', 'b', 'c', 'd']\nchanged_name = 'change_change'\n[Map]\nhello = 'world'\nmyname = 'sunjin'\n\n"

			obj := &TestTomlStruct{}
			err := tomlutil.Unmarshal([]byte(str), obj)
			assert.Nil(t, err)
			assert.EqualValues(t, obj, &TestTomlStruct{
				Str:   "test_str",
				Num:   1,
				Num32: 2,
				Num64: 3,
				Map: map[string]string{
					"hello":  "world",
					"myname": "sunjin",
				},
				List:       []string{"a", "b", "c", "d"},
				NameChange: "change_change",
			})

		})
	})

}

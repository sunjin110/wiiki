package tomlutil

import (
	"wiiki_server/common/wiikierr"

	"github.com/pelletier/go-toml/v2"
)

// Unmarshal obj is must pointer
func Unmarshal(b []byte, obj interface{}) error {
	err := toml.Unmarshal(b, obj)
	if err != nil {
		return wiikierr.Bind(err, wiikierr.FailedUnmarshalToml, "data=%s", string(b))
	}
	return nil
}

func Marshal(obj interface{}) ([]byte, error) {
	b, err := toml.Marshal(obj)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedMarshalToml, "obj=%s", obj)
	}
	return b, nil
}

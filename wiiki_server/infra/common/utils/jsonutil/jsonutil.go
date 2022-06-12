package jsonutil

import (
	"encoding/json"
	"wiiki_server/infra/common/wiikierr"
)

func Marshal(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", wiikierr.Bind(err, wiikierr.JsonMarshalFailed, "obj is %v", obj)
	}
	return string(b), nil
}

func MustMarshal(obj interface{}) string {
	str, err := Marshal(obj)
	if err != nil {
		panic(err)
	}
	return str
}

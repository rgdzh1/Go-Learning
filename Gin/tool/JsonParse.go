package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

/**
参数解析
io 用户请求接口传过来的参数
*/
func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}

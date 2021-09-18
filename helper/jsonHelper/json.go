package jsonHelper

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"reflect"
	"unsafe"
)

func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	marshal, err := Marshal(v)
	return Byte2Str(marshal), err

}

func MarshalToStringNoError(v interface{}) string {
	str, _ := MarshalToString(v)
	return str
}

func Unmarshal(data string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(Str2ByteByReflect(data), v)
}

func UnmarshalByte(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}

// 解析时容忍空数组作为对象
func UnmarshalStringFuzzyDecoders(data string, v interface{}) error {
	extra.RegisterFuzzyDecoders()
	return jsoniter.UnmarshalFromString(data, v)
}

func Str2Byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2ByteByReflect(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

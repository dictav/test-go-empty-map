package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"

	"github.com/dictav/test-go-empty-map/model"
)

func main() {
	println("nil")
	data := &model.Data{KV: nil}
	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buf, err = easyjson.Marshal(data)
	if err != nil {
		panic(err)
	}
	println("encoding/json", string(buf))
	println("mailru/easyjson", string(buf))

	println("")
	println("nil map")
	var kv *map[string]float64
	data = &model.Data{KV: kv}
	buf, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buf, err = easyjson.Marshal(data)
	if err != nil {
		panic(err)
	}
	println("encoding/json", string(buf))
	println("mailru/easyjson", string(buf))

	println("")
	println("map of nil")
	var src map[string]float64
	src = nil
	kv = &src
	data = &model.Data{KV: kv}
	buf, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buf, err = easyjson.Marshal(data)
	if err != nil {
		panic(err)
	}
	println("encoding/json", string(buf))
	println("mailru/easyjson", string(buf))

	println("")
	println("empty map")
	kv = &map[string]float64{}
	data = &model.Data{KV: kv}
	buf, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buf, err = easyjson.Marshal(data)
	if err != nil {
		panic(err)
	}
	println("encoding/json", string(buf))
	println("mailru/easyjson", string(buf))

	println("")
	println("not nil map")
	kv = &map[string]float64{"2": 1}
	data = &model.Data{KV: kv}
	buf, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buf, err = easyjson.Marshal(data)
	if err != nil {
		panic(err)
	}
	println("encoding/json", string(buf))
	println("mailru/easyjson", string(buf))
}

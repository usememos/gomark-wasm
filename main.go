package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/yourselfhosted/gomark/parser"
	"github.com/yourselfhosted/gomark/parser/tokenizer"
	"github.com/yourselfhosted/gomark/restore"
)

// Parse converts markdown to nodes.
func Parse(this js.Value, inputs []js.Value) any {
	markdown := inputs[0].String()
	tokens := tokenizer.Tokenize(markdown)
	astNodes, err := parser.Parse(tokens)
	if err != nil {
		panic(err)
	}

	nodes := convertFromASTNodes(astNodes)
	bytes, _ := json.Marshal(nodes)
	data := []interface{}{}
	json.Unmarshal(bytes, &data)
	return data
}

// Restore converts nodes to markdown.
func Restore(this js.Value, inputs []js.Value) any {
	astNodes, ok := convertJSValueToInterface(inputs[0]).([]interface{})
	if !ok {
		return nil
	}

	nodes := []*Node{}
	bytes, _ := json.Marshal(astNodes)
	json.Unmarshal(bytes, &nodes)
	content := restore.Restore(convertToASTNodes(nodes))
	return content
}

// convertJSValueToInterface converts a js.Value to a Go interface{}.
func convertJSValueToInterface(value js.Value) interface{} {
	switch value.Type() {
	case js.TypeString:
		return value.String()
	case js.TypeNumber:
		return value.Float()
	case js.TypeBoolean:
		return value.Bool()
	case js.TypeObject:
		if value.InstanceOf(js.Global().Get("Array")) {
			length := value.Length()
			array := make([]interface{}, length)
			for i := 0; i < length; i++ {
				array[i] = convertJSValueToInterface(value.Index(i))
			}
			return array
		} else {
			obj := make(map[string]interface{})
			keys := js.Global().Get("Object").Call("keys", value)
			for i := 0; i < keys.Length(); i++ {
				key := keys.Index(i).String()
				obj[key] = convertJSValueToInterface(value.Get(key))
			}
			return obj
		}
	default:
		return nil
	}
}

func registerCallbacks() {
	js.Global().Set("parse", js.FuncOf(Parse))
	js.Global().Set("restore", js.FuncOf(Restore))
}

func main() {
	registerCallbacks()

	select {} // block forever
}

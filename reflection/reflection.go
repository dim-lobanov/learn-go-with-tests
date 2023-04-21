package reflection

import (
	"reflect"
)

// Reflection in computing is the ability of a program to examine its own structure, particularly through types;
// it's a form of metaprogramming. It's also a great source of confusion.

// any <=> interface{}

func Walk(x any, fn func(input string)) {
	val := getValue(x)
	walkFunc := func(v reflect.Value) {
		Walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkFunc(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkFunc(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkFunc(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkFunc(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkFunc(res)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}

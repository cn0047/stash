package main

import (
	"fmt"
	"io"
	"reflect"
)

func main() {
	//one()
	two()
}

func one() {
	var r io.Reader
	var w io.Writer
	w = r.(io.Writer)
	fmt.Printf("%#v", w)
}

func two() {
	a := 1
	t := reflect.TypeOf(a)
	fmt.Printf("\n%+v", t)              // int
	fmt.Printf("\nKind: %+v", t.Kind()) // int
	v := reflect.ValueOf(a)
	fmt.Printf("\n%+v", v)        // 1
	fmt.Printf("\n%+v", v.Type()) // int
}

func three() {
}

//type ConfigPrototype struct {
//	Type  string
//	Value interface{}
//}
//
//func (cp ConfigPrototype) validateValue() error {
//	switch cp.Type {
//	case TypeInt:
//		switch cp.Value.(type) {
//		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
//		default:
//			return fmt.Errorf("inappropriate config value for type %s, actual type: %T", cp.Type, cp.Value)
//		}
//	case TypeFloat:
//		switch cp.Value.(type) {
//		case float32, float64:
//		default:
//			return fmt.Errorf("inappropriate config value for type %s, actual type: %T", cp.Type, cp.Value)
//		}
//	case TypeString:
//		switch cp.Value.(type) {
//		case string:
//		default:
//			return fmt.Errorf("inappropriate config value for type %s, actual type: %T", cp.Type, cp.Value)
//		}
//	case TypeBool:
//		switch cp.Value.(type) {
//		case bool:
//		default:
//			return fmt.Errorf("inappropriate config value for type %s, actual type: %T", cp.Type, cp.Value)
//		}
//	case TypeArray:
//		typeOf := reflect.TypeOf(cp.Value)
//		switch typeOf.Kind() {
//		case reflect.Array, reflect.Slice:
//		default:
//			return fmt.Errorf("inappropriate config value for type array, actual type: %s", typeOf)
//		}
//	case TypeObject:
//		typeOf := reflect.TypeOf(cp.Value)
//		switch typeOf.Kind() {
//		case reflect.Map, reflect.Struct:
//		default:
//			return fmt.Errorf("inappropriate config value for type object, actual type: %T", typeOf)
//		}
//	}
//
//	return nil
//}

package gomock

import (
	//"math/rand"
	"reflect"
	//"strings"
	//"time"
	"log"
)

func Struct(v interface{}) {
	//r(reflect.TypeOf(v), reflect.ValueOf(v), "")
	// 不经过prt(not isNil)->struct，一步直达
	rStruct(reflect.TypeOf(v).Elem(), reflect.ValueOf(v).Elem())
}

func rStruct(t reflect.Type, v reflect.Value) {
	n := t.NumField()
	for i := 0; i < n; i++ {
		elementT := t.Field(i)
		elementV := v.Field(i)
		//mock := true
		t, ok := elementT.Tag.Lookup("mock")
		if ok && t != "skip" {
			log.Println("Should Mock:", t)
			switch t {
			case "name":
				elementV.SetString(Name())
			case "age":
				elementV.SetInt(int64(Age()))
			}
		}

		/*if mock && elementV.CanSet() {
			r(elementT.Type, elementV, t)
		}*/
	}
}

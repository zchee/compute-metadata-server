package fakemetadata

import (
	"reflect"
	"unsafe"
)

var offset uintptr

func init() {
	field, ok := reflect.ValueOf(reflect.Value{}).Type().FieldByName("flag")
	if !ok {
		panic("unable to find the flag field of reflect.Value, contact dark gophers")
	}
	offset = field.Offset
}

func toExport(v reflect.Value) reflect.Value {
	// copied from reflect package. hopefully it says in sync!
	const flagRO = 1<<5 | 1<<6

	ptr := unsafe.Pointer(&v)
	fptr := (*uintptr)(unsafe.Pointer(uintptr(ptr) + offset))
	*fptr = *fptr &^ flagRO

	return v
}

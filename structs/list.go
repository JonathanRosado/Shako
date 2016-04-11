package structs

import (
	. "reflect"
)

type List struct {
	elems []interface{}
}

func (l *List) Add(elem interface{}) {
	append(l.elems, elem)
}

func (l *List) Get(pos int) func() {
	elem := ValueOf(l.elems[pos])

	switch elem.Kind() {
	case Int, Int8, Int16, Int32, Int64:
		return elem.Int()
	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		return elem.Uint()
	case Float32, Float64:
		return elem.Float()
	case Complex64, Complex128:
		return elem.Complex()
	case String:
		return elem.String()
	case Bool:
		return elem.Bool()
	case Ptr:
		return elem.Pointer()
	//case Array, Slice:
	//	return elem.sli()
	//case Map:
	//	return elem.map()
	//case Chan:
	//	return elem.C()
	//case Struct:
	//	return elem.Uint()
	//case Interface:
	//	return elem.Uint()
	//case Func:
	//	return elem.Uint()
	default:
		panic("Can't determine element's type in List")
	}
}

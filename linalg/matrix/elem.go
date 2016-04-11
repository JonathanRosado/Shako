package matrix

import "reflect"

func (m *Matrix) Elem(args ...interface{}) float64 {

	switch len(args) {
	case 0:
		return m.Elem(0)
	case 1:
		arg := reflect.ValueOf(args[0])
		switch arg.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return m.vectorIndex(int(arg.Int()))
		default:
			panic("Argument not supported by Matrix.Elem")
		}
	case 2:
		ijType := []reflect.Kind{}
		ijVal := []reflect.Value{}

		ijVal = append(ijVal, reflect.ValueOf(args[0]))
		ijVal = append(ijVal, reflect.ValueOf(args[1]))

		ijType = append(ijType, ijVal[0].Kind())
		ijType = append(ijType, ijVal[1].Kind())

		switch ijType[0] {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch ijType[1] {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return m.matrixIndex(int(ijVal[0].Int()), int(ijVal[1].Int()))
			default:
				panic("Argument not supported by Matrix.Elem")
			}
		default:
			panic("Argument not supported by Matrix.Elem")
		}
	default:
		panic("Too many arguments given to Matrix.Elem")
	}

	return -1
}
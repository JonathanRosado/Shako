package matrix

import "reflect"

func (m *Matrix) Index(args ...interface{}) *Matrix {

	switch len(args) {
	case 1:
		arg := reflect.ValueOf(args[0])
		switch arg.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return (&Matrix{}).Insert(
				Row{m.vectorIndex(int(arg.Int()))},
			)
		case reflect.String:
			return m.colVector(-1)
		default:
			panic("Argument not supported by Matrix.Index")
		}
	case 2:
		ijType := []reflect.Kind{}
		ijVal := []reflect.Value{}

		ijVal = append(ijVal, reflect.ValueOf(args[0]))
		ijVal = append(ijVal, reflect.ValueOf(args[1]))

		ijType = append(ijType, ijVal[0].Kind())
		ijType = append(ijType, ijVal[1].Kind())

		switch ijType[0] {
		// Index(Int, [String, Int])
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch ijType[1] {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return m.i_i_index(int(ijVal[0].Int()), int(ijVal[1].Int()))
			case reflect.String:
				return m.i_s_index(int(ijVal[0].Int()), ijVal[1].String())
			default:
				panic("Argument not supported by Matrix.Index")
			}
		// Index(String, [String, Int])
		case reflect.String:
			switch ijType[1] {
			//Index(String, Int)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return m.s_i_index(ijVal[0].String(), int(ijVal[1].Int()))
			//Index(String, String)
			case reflect.String:
				return m.s_s_index(ijVal[0].String(), ijVal[1].String())
			default:
				panic("Argument not supported by Matrix.Index")
			}
		default:
			panic("Argument not supported by Matrix.Index")
		}
	case 3:
		returnMatrix := &Matrix{}

		ijzType := []reflect.Kind{}
		ijzVal := []reflect.Value{}

		ijzVal = append(ijzVal, reflect.ValueOf(args[0]))
		ijzVal = append(ijzVal, reflect.ValueOf(args[1]))
		ijzVal = append(ijzVal, reflect.ValueOf(args[2]))

		ijzType = append(ijzType, ijzVal[0].Kind())
		ijzType = append(ijzType, ijzVal[1].Kind())
		ijzType = append(ijzType, ijzVal[2].Kind())

		if ijzType[2] != reflect.Func {
			panic("Argument not supported by Matrix.Index")
		}

		switch ijzType[0] {
		// Index(Int, [String, Int])
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch ijzType[1] {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				returnMatrix = m.i_i_index(int(ijzVal[0].Int()), int(ijzVal[1].Int()))
			case reflect.String:
				returnMatrix = m.i_s_index(int(ijzVal[0].Int()), ijzVal[1].String())
			default:
				panic("Argument not supported by Matrix.Index")
			}
		// Index(String, [String, Int])
		case reflect.String:
			switch ijzType[1] {
			//Index(String, Int)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				returnMatrix = m.s_i_index(ijzVal[0].String(), int(ijzVal[1].Int()))
			//Index(String, String)
			case reflect.String:
				returnMatrix = m.s_s_index(ijzVal[0].String(), ijzVal[1].String())
			default:
				panic("Argument not supported by Matrix.Index")
			}
		default:
			panic("Argument not supported by Matrix.Index")
		}

		returnMatrix.ElemOpInterface(ijzVal[2].Interface())

		return returnMatrix
	default:
		panic("Too many arguments given to Matrix.Index")
	}

	return (&Matrix{}).Insert(
		Row{0},
	)
}

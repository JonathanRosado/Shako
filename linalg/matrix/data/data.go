package data

import (
	log "github.com/Sirupsen/logrus"
	"fmt"
)

type Mode int
const (
	NORMAL Mode = iota
	TRANSPOSE
	COLUMN_VECTOR
	ROW_VECTOR
)

type Data struct {
	Table  [][]float64
	Cols, Rows int
	cols, rows int
	mode Mode
}

func (d *Data) Print() {
	for i := 0; i < d.Rows; i++ {
		fmt.Print("[ ")
		for j := 0; j < d.Cols; j++ {
			fmt.Printf("%f ",d.Get(i, j))
		}
		fmt.Print("]\n")
	}
}

func (d *Data) Create(rows [][]float64) *Data {
	d.rows = len(rows)
	d.cols = len(rows[0])
	d.Rows = d.rows
	d.Cols = d.cols

	d.mode = NORMAL

	d.Table = rows

	return d
}

func (d *Data) Get(i, j int) float64 {
	switch d.mode {
	case ROW_VECTOR:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		pos := j

		// i(vector_index) = vector_index/number_of_columns
		i = pos/d.cols

		// j(vector_index) = vector_index - (number_of_columns*i)
		j = pos - (d.cols * i)

		return d.Table[i][j]
	case COLUMN_VECTOR:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		pos := i

		// i(vector_index) = vector_index/number_of_columns
		i = pos/d.cols

		// j(vector_index) = vector_index - (number_of_columns*i)
		j = pos - (d.cols * i)

		return d.Table[i][j]
	case TRANSPOSE:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		return d.Table[j][i]
	case NORMAL:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		return d.Table[i][j]
	default:
		panic("No Data mode selected")
	}
}

func (d *Data) GetRow(row int) []float64 {
	if row >= d.Rows {
		log.WithFields(log.Fields{
			"d.Rows": d.Rows,
			"row pos": row,
		}).Warn("GetRow out of bounds:")
		panic("GetRow out of bounds:")
	}

	return d.Table[row]
}

func (d *Data) GetColumn(col int) []float64 {
	if col >= d.Cols {
		panic("GetColumn out of bounds")
	}

	column := []float64{}

	for i := range d.Table {
		column = append(column, d.Table[i][col])
	}

	return column
}

func (d *Data) Set(i, j int, val float64) *Data {
	switch d.mode {
	case ROW_VECTOR:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		pos := j

		// i(vector_index) = vector_index/number_of_columns
		i = pos/d.cols

		// j(vector_index) = vector_index - (number_of_columns*i)
		j = pos - (d.cols * i)

		d.Table[i][j] = val
	case COLUMN_VECTOR:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		pos := i

		// i(vector_index) = vector_index/number_of_columns
		i = pos/d.cols

		// j(vector_index) = vector_index - (number_of_columns*i)
		j = pos - (d.cols * i)

		d.Table[i][j] = val
	case TRANSPOSE:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		d.Table[j][i] = val
	case NORMAL:
		if j >= d.Cols || i >= d.Rows {
			panic("Get out of bounds")
		}

		d.Table[i][j] = val
	default:
		panic("No Data mode selected")
	}

	return d
}

// TODO:Fix row col vectors SEE octave
func (d *Data) ColumnVector() *Data {
	switch d.mode {
	case NORMAL:
		d.Rows = d.cols * d.rows
		d.Cols = 1
		d.mode = COLUMN_VECTOR
	case ROW_VECTOR:
		d.Rows = d.cols * d.rows
		d.Cols = 1
		d.mode = COLUMN_VECTOR
	case TRANSPOSE:
		// Switch dimensions
		rowsTemp := d.Rows
		d.Rows = d.Cols
		d.Cols = rowsTemp
		d.mode = NORMAL

		d.RowVector()
	case COLUMN_VECTOR:
		// Do nothing
	default:
		panic("No Data mode selected")
	}

	return d
}

func (d *Data) RowVector() *Data {
	switch d.mode {
	case NORMAL:
		d.Cols = d.cols * d.rows
		d.Rows = 1
		d.mode = ROW_VECTOR
	case ROW_VECTOR:
	// Do nothing
	case TRANSPOSE:
		// Switch dimensions
		rowsTemp := d.Rows
		d.Rows = d.Cols
		d.Cols = rowsTemp
		d.mode = NORMAL

		d.ColumnVector()
	case COLUMN_VECTOR:
		d.Cols = d.cols * d.rows
		d.Rows = 1
		d.mode = ROW_VECTOR
	default:
		panic("No Data mode selected")
	}

	return d
}

func (d *Data) Transpose() *Data {
	switch d.mode {
	case NORMAL:
		// Switch dimensions
		rowsTemp := d.Rows
		d.Rows = d.Cols
		d.Cols = rowsTemp
		d.mode = TRANSPOSE
	case ROW_VECTOR:
		d.ColumnVector()
	case TRANSPOSE:
		// Switch dimensions
		rowsTemp := d.Rows
		d.Rows = d.Cols
		d.Cols = rowsTemp
		d.mode = NORMAL
	case COLUMN_VECTOR:
		d.RowVector()
	default:
		panic("No Data mode selected")
	}

	return d
}

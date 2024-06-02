package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/jedib0t/go-pretty/table"
)

// Function to get comma-separated values from a slice of elements that implement String() string
func getCommaSeparated[T interface{ String() string }](values []T) (result string) {
	for _, v := range values {
		if result != "" {
			result += ", "
		}
		result += v.String()
	}
	return result
}

// Função para imprimir uma tabela de valores
func Print[T any](values []T) {
	if len(values) == 0 {
		fmt.Println("No data to display")
		return
	}

	v := reflect.ValueOf(values[0])
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	headers := createHeaders(v)
	rows := createRows(values)
	footer := createFooter(values, len(headers))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(headers)
	t.AppendRows(rows)
	t.AppendFooter(footer)
	t.Render()
}

// Função auxiliar para criar cabeçalhos da tabela
func createHeaders(v reflect.Value) table.Row {
	headers := make(table.Row, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		headers[i] = v.Type().Field(i).Name
	}
	return headers
}

// Função auxiliar para criar as linhas da tabela
func createRows[T any](values []T) []table.Row {
	rows := make([]table.Row, len(values))
	for i, value := range values {
		rows[i] = createRow(value)
	}
	return rows
}

// Função auxiliar para criar rodapés da tabela
func createFooter[T any](values []T, size int) table.Row {
	footer := make(table.Row, size)
	footer[size-2] = "Total"
	footer[size-1] = strconv.Itoa(len(values))
	return footer
}

// Função auxiliar para criar uma linha da tabela
func createRow[T any](value T) table.Row {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	row := make(table.Row, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		row[i] = formatField(field)
	}
	return row
}

// Função auxiliar para formatar o campo da tabela
func formatField(field reflect.Value) interface{} {
	if field.Kind() == reflect.Array || field.Kind() == reflect.Slice {
		return formatSliceField(field)
	}
	return field.Interface()
}

// Função auxiliar para formatar campos que são slices ou arrays
func formatSliceField(field reflect.Value) string {
	items := []interface{ String() string }{}
	for i := 0; i < field.Len(); i++ {
		item := field.Index(i).Interface()
		if strItem, ok := item.(interface{ String() string }); ok {
			items = append(items, strItem)
		}
	}
	return getCommaSeparated(items)
}

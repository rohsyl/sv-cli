package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
	"sv-cli/metrics"
)

// OutputResult handles JSON and table output formats correctly
func OutputResult(result metrics.MetricResult, format string) {
	if format == "json" {
		output, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(output))
		return
	}

	if !result.Success {
		fmt.Println("❌ Error:", result.Error)
		return
	}

	fmt.Println("✅ Success")

	// Dynamically handle different types of results
	printDynamicTable(result.Data)
}

func printDynamicTable(data interface{}) {
	v := reflect.ValueOf(data)

	// Unwrap pointer if necessary
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice:
		if v.Len() == 0 {
			fmt.Println("No data available.")
			return
		}
		printStructSlice(v.Interface())

	case reflect.Struct:
		printStructSlice([]interface{}{data})

	case reflect.Map:
		printMapTable(data) // ✅ Now works for both map[string]string and map[string]interface{}

	default:
		fmt.Printf("⚠️ Unsupported data type: %s\n", v.Kind())
	}
}

// Print a table for both map[string]string and map[string]interface{}
func printMapTable(data interface{}) {
	switch mapData := data.(type) {
	case map[string]string:
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Key", "Value"})

		for key, value := range mapData {
			table.Append([]string{key, value})
		}

		table.Render()

	case map[string]interface{}:
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Key", "Value"})

		for key, value := range mapData {
			table.Append([]string{key, fmt.Sprintf("%v", value)})
		}

		table.Render()

	default:
		fmt.Println("❌ Error: Expected a map[string]string or map[string]interface{}, but got", reflect.TypeOf(data))
	}
}

// Print a table from a slice of structs
func printStructSlice(slice interface{}) {
	v := reflect.ValueOf(slice)

	if v.Len() == 0 {
		fmt.Println("No data available.")
		return
	}

	headers, rows := extractStructData(v)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, row := range rows {
		table.Append(row)
	}
	table.Render()
}

// Extracts struct data dynamically for use in tables
func extractStructData(v reflect.Value) ([]string, [][]string) {
	var headers []string
	var rows [][]string

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)

		// Unwrap interface and pointer if necessary
		if elem.Kind() == reflect.Interface {
			elem = elem.Elem()
		}
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		// Ensure we have a struct
		if elem.Kind() != reflect.Struct {
			fmt.Println("❌ Error: Expected a struct but got", elem.Kind())
			return nil, nil
		}

		elemType := elem.Type()
		var row []string

		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			value := elem.Field(j)

			if i == 0 {
				headers = append(headers, field.Name)
			}

			row = append(row, fmt.Sprintf("%v", value.Interface()))
		}
		rows = append(rows, row)
	}

	return headers, rows
}

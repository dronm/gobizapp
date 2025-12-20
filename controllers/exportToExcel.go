package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	md "github.com/dronm/crudifier/metadata"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

const (
	defaultExcelFileName         = "export"
	excelFileExt                 = ".xlsx"
	defaultExcelHeaderBackground = "#d6d6d6"
)

var excelHeaderBackground string

type Stringer interface {
	String() string
}

// BuildRowsFromStructs converts a slice of structs into [][]any for Excel
func BuildRowsFromStructs[T any](list []T, fieldNames []string) [][]any {
	rows := make([][]any, len(list))

	for i, item := range list {
		val := reflect.ValueOf(item)

		// Dereference pointer to struct if needed
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				continue
			}
			val = val.Elem()
		}

		row := make([]any, len(fieldNames))
		for j, field := range fieldNames {
			f := val.FieldByName(field)

			if !f.IsValid() {
				row[j] = "" // field not found

				continue
			}

			// Handle pointers
			if f.Kind() == reflect.Ptr {
				if f.IsNil() {
					row[j] = ""
					continue
				}
				f = f.Elem()
			}

			// Handle structs (skip value, leave empty)
			if f.Kind() == reflect.Struct {
				// check of gobizap field types
				if fIntf, ok := f.Interface().(md.ModelFieldInt); ok {
					row[j] = fIntf.GetValue()
				} else if fIntf, ok := f.Interface().(md.ModelFieldFloat); ok {
					row[j] = fIntf.GetValue()
				} else if fIntf, ok := f.Interface().(md.ModelFieldBool); ok {
					row[j] = fIntf.GetValue()
				} else if fIntf, ok := f.Interface().(md.ModelFieldDate); ok {
					row[j] = fIntf.GetValue()
				} else if fIntf, ok := f.Interface().(md.ModelFieldText); ok {
					row[j] = fIntf.GetValue()
				} else if fIntf, ok := f.Interface().(Stringer); ok {
					fmt.Println(field)
					row[j] = fIntf.String()
				} else {
					row[j] = ""
				}
				continue
			}

			// Handle nil interface values
			if f.Kind() == reflect.Interface && f.IsNil() {
				row[j] = ""
				continue
			}

			row[j] = f.Interface()
		}

		rows[i] = row
	}

	return rows
}

func WriteExcel(c *gin.Context, filename string, headers []string, rows [][]any) error {
	file := excelize.NewFile()
	sheet := "Sheet1"

	headerBg := excelHeaderBackground
	if headerBg == "" {
		headerBg = defaultExcelHeaderBackground
	}

	// Create header style: centered text + background color
	headerStyle, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{headerBg}, 
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true,
		},
	})

	// headers
	for i, h := range headers {
		col := string(rune('A' + i))
		cell := fmt.Sprintf("%s1", col)
		file.SetCellValue(sheet, cell, h)
		file.SetCellStyle(sheet, cell, cell, headerStyle)

		width := float64(len(h)) + 2
		file.SetColWidth(sheet, col, col, width)
	}

	// data rows
	for r, row := range rows {
		for i, val := range row {
			col := string(rune('A' + i))
			file.SetCellValue(sheet, fmt.Sprintf("%s%d", col, r+2), val)
		}
	}
	encodedFilename := url.PathEscape(filename) // URL-encode UTF-8 filename

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition",
		fmt.Sprintf(`attachment; filename="%s"; filename*=UTF-8''%s`, filename, encodedFilename))
	c.Header("Access-Control-Expose-Headers", "Content-Disposition") // to enable javascript to access the header

	return file.Write(c.Writer)
}

type ModelExporter interface {
	ExportFileName() string
}

func ExportToExcel[T any](resList []T, model any, c *gin.Context) {
	modelMD, err := md.NewModelMetadata(model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ExportToExcel md.NewModelMetadata", err)
		return
	}

	var excelFields []string
	var headers []string                     // number of fields is  unknown
	for i, f := range modelMD.FieldTagList { // sql fields
		fld, ok := modelMD.Fields[f]
		if !ok {
			ServeError(c, http.StatusInternalServerError, "ExportToExcel", fmt.Errorf("modelMD.Fields for %s not found", f))
			return
		}
		fieldName := modelMD.FieldList[i]

		// check export tag, if - OR false - skip field.
		if fieldTags, ok := modelMD.Tags[fieldName]; ok {
			if expTagVal, ok := fieldTags["export"]; ok {
				if expTagVal == "-" || expTagVal == "false" || expTagVal == "FALSE" {
					continue
				}
			}
		}

		headers = append(headers, fld.Descr())
		excelFields = append(excelFields, fieldName)
	}

	rows := BuildRowsFromStructs(resList, excelFields)

	fileName := defaultExcelFileName // common name if undefined
	if exporter, ok := model.(ModelExporter); ok {
		fileName = exporter.ExportFileName()
	}

	if err := WriteExcel(c, fileName+excelFileExt, headers, rows); err != nil {
		ServeError(c, http.StatusInternalServerError, "ExportToExcel WriteExcel", err)
		return
	}
}

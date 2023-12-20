package converter

import (
	"fmt"

	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

type enumData struct {
	enum        *protogen.Enum
	enumName    string
	description string
	values      map[string]*enumValue
	valuesOrder []string
}

func (data *enumData) GetTableData() (*enumTableData, error) {
	tableData := &enumTableData{
		colLengths: []int{0, 0},
		rows:       []*enumTableFieldRow{},
	}

	for _, valueKey := range data.valuesOrder {
		value, ok := data.values[valueKey]
		if !ok {
			return nil, fmt.Errorf("value %s not found", valueKey)
		}

		row := &enumTableFieldRow{
			valueName:   fmt.Sprintf("`%s`", valueKey),
			description: value.description,
		}

		tableData.rows = append(tableData.rows, row)

		utils.StringGTLengthHelper(&tableData.colLengths[0], row.valueName)
		utils.StringGTLengthHelper(&tableData.colLengths[1], row.description)
	}

	return tableData, nil
}

type enumValue struct {
	valueName   string
	description string
}

type enumTableData struct {
	colLengths []int
	rows       []*enumTableFieldRow
}

type enumTableFieldRow struct {
	valueName   string
	description string
}

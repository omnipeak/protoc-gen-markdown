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

func (data *enumData) GetTableData() (*utils.TableData, error) {
	tableData := &utils.TableData{
		Headers: []string{"Value", "Description"},
		Rows:    [][]string{},
	}

	if len(data.values) != len(data.valuesOrder) {
		return nil, fmt.Errorf(
			"values and valuesOrder slice lengths do not match: %d != %d",
			len(data.values),
			len(data.valuesOrder),
		)
	}

	for _, valueKey := range data.valuesOrder {
		value, ok := data.values[valueKey]
		if !ok {
			return nil, fmt.Errorf("value '%s' not found in enum data", valueKey)
		}

		row := []string{
			fmt.Sprintf("`%s`", valueKey),
			value.description,
		}

		tableData.Rows = append(tableData.Rows, row)
	}

	return tableData, nil
}

type enumValue struct {
	valueName   string
	description string
}

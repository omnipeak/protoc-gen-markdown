package converter

import (
	"fmt"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

type messageData struct {
	message     *protogen.Message
	messageName string
	description string
	fields      map[string]*messageField
	fieldsOrder []string
}

func (data *messageData) GetTableData() (*messageTableData, error) {
	res := &messageTableData{
		colLengths: []int{4, 4, 9, 11},
		rows:       []*messageTableFieldRow{},
	}

	for _, fieldKey := range data.fieldsOrder {
		field, ok := data.fields[fieldKey]
		if !ok {
			return nil, fmt.Errorf("field %s not found", fieldKey)
		}

		row := &messageTableFieldRow{
			fieldName:   fmt.Sprintf("`%s`", string(field.fieldName)),
			description: field.description,
			required:    field.required,
		}

		switch field.fieldType {
		case "message", "enum":
			typeName := field.fieldMessage
			if field.isList {
				typeName += "[]"
			}

			row.fieldType = fmt.Sprintf(
				"[`%s`](#%s)",
				typeName,
				strings.ToLower(field.fieldMessage),
			)

		default:
			typeName := field.fieldType
			if field.isList {
				typeName += "[]"
			}

			row.fieldType = fmt.Sprintf("`%s`", string(typeName))
		}

		utils.StringGTLengthHelper(&res.colLengths[0], row.fieldName)
		utils.StringGTLengthHelper(&res.colLengths[1], row.fieldType)
		utils.StringGTLengthHelper(&res.colLengths[3], row.description)

		res.rows = append(res.rows, row)
	}

	return res, nil
}

type messageField struct {
	fieldName    string
	fieldType    string
	fieldMessage string
	isList       bool
	required     bool
	description  string
	options      *validate.FieldConstraints
}

type messageTableData struct {
	colLengths []int
	rows       []*messageTableFieldRow
}

type messageTableFieldRow struct {
	fieldName   string
	fieldType   string
	required    bool
	description string
}

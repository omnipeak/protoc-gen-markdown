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

		pathPrefix := strings.Repeat(
			"../",
			strings.Count(
				data.message.Location.SourceFile,
				"/",
			),
		)

		var typeName string
		var typeLink string
		switch field.fieldType {
		case "enum":
			typeName = string(field.fieldEnum.Desc.Name())
			if field.fieldEnum.Location.SourceFile != data.message.Location.SourceFile {
				typeLink += pathPrefix + strings.ReplaceAll(
					string(field.fieldEnum.Location.SourceFile),
					".proto",
					".md",
				)
			}
			typeLink += "#" + strings.ToLower(typeName) + "-enum"

			if field.isList {
				typeName += "[]"
			}

			row.fieldType = fmt.Sprintf(
				"[`%s`](%s)",
				typeName,
				typeLink,
			)

		case "message":
			if field.fieldMessage.Desc.IsMapEntry() {
				row.fieldType = fmt.Sprintf(
					"`map<%s, %s>`",
					field.fieldMessage.Fields[0].Desc.Kind().String(),
					field.fieldMessage.Fields[1].Desc.Kind().String(),
				)
			} else {
				typeName = string(field.fieldMessage.Desc.Name())
				if field.fieldMessage.Location.SourceFile != data.message.Location.SourceFile {
					typeLink += pathPrefix + strings.ReplaceAll(
						string(field.fieldMessage.Location.SourceFile),
						".proto",
						".md",
					)
				}
				typeLink += "#" + strings.ToLower(typeName) + "-message"

				if field.isList {
					typeName += "[]"
				}

				row.fieldType = fmt.Sprintf(
					"[`%s`](%s)",
					typeName,
					typeLink,
				)
			}

		default:
			typeName = field.fieldType
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
	fieldEnum    *protogen.Enum
	fieldMessage *protogen.Message
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

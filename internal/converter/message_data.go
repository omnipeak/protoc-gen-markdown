package converter

import (
	"fmt"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type messageData struct {
	message     *protogen.Message
	messageName string
	description string
	fields      map[string]*messageField
	fieldsOrder []string
}

func (data *messageData) GetTableData() (*utils.TableData, error) {
	tableData := &utils.TableData{
		Headers: []string{
			"Name",
			"Type",
			"Required?",
			"Description",
		},
		Rows: [][]string{},
	}

	if len(data.fields) != len(data.fieldsOrder) {
		return nil, errors.Errorf(
			"fields and fieldsOrder lengths do not match: %d != %d",
			len(data.fields),
			len(data.fieldsOrder),
		)
	}

	for _, fieldKey := range data.fieldsOrder {
		field, ok := data.fields[fieldKey]
		if !ok {
			return nil, errors.Errorf("field %s not found", fieldKey)
		}

		pathPrefix := strings.Repeat(
			"../",
			strings.Count(
				data.message.Location.SourceFile,
				"/",
			),
		)

		fieldType := data.getFieldType(field, pathPrefix)

		row := []string{
			fmt.Sprintf("`%s`", string(field.fieldName)),
			fieldType,
			utils.BoolToTickOrCross(field.required),
			field.description,
		}

		tableData.Rows = append(tableData.Rows, row)
	}

	return tableData, nil
}

func (data *messageData) getFieldType(field *messageField, pathPrefix string) string {
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

		return fmt.Sprintf(
			"[`%s`](%s)",
			typeName,
			typeLink,
		)

	case "message":
		if field.fieldMessage == nil {
			return "message"
		}

		if field.fieldMessage.Desc.IsMapEntry() {
			return fmt.Sprintf(
				"`map<%s, %s>`",
				field.fieldMessage.Fields[0].Desc.Kind().String(),
				field.fieldMessage.Fields[1].Desc.Kind().String(),
			)
		}

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

		return fmt.Sprintf(
			"[`%s`](%s)",
			typeName,
			typeLink,
		)

	default:
		typeName = field.fieldType
		if field.isList {
			typeName += "[]"
		}

		if strings.Contains(typeName, "`") {
			return typeName
		}

		return fmt.Sprintf("`%s`", typeName)
	}
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

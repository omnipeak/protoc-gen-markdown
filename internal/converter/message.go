package converter

import (
	"fmt"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
)

func (c *converter) processMessage(g *protogen.GeneratedFile, message *protogen.Message) error {
	var err error

	data := &messageData{
		message:     message,
		messageName: string(message.Desc.Name()),
		description: utils.FlattenComment(
			message.Comments.Leading.String() + "\n\n" + message.Comments.Trailing.String(),
		),
		fields:      map[string]*messageField{},
		fieldsOrder: []string{},
	}

	for _, f := range message.Fields {
		field := &messageField{
			fieldName: string(f.Desc.Name()),
			description: utils.FlattenComment(
				f.Comments.Leading.String() + " " + f.Comments.Trailing.String(),
			),
		}

		switch f.Desc.Kind() {
		case protoreflect.EnumKind:
			field.fieldType = "enum"
			field.fieldEnum = f.Enum

		case protoreflect.MessageKind:
			field.fieldType = "message"
			field.fieldMessage = f.Message

		default:
			field.fieldType = f.Desc.Kind().String()
		}

		field.isList = f.Desc.IsList()
		field.options = proto.GetExtension(f.Desc.Options(), validate.E_Field).(*validate.FieldConstraints)

		err = c.processFieldOptions(g, data, field, message)
		if err != nil {
			return err
		}

		data.fields[string(field.fieldName)] = field
		data.fieldsOrder = append(data.fieldsOrder, field.fieldName)
	}

	err = c.writeMessageFieldsTable(g, data)
	if err != nil {
		return err
	}

	err = c.writeMessageFieldValidationTable(g, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *converter) processFieldOptions(
	f *protogen.GeneratedFile,
	data *messageData,
	row *messageField,
	message *protogen.Message,
) error {
	if row.options == nil {
		return nil
	}

	row.required = row.options.GetRequired()

	return nil
}

func (c *converter) writeMessageFieldsTable(f *protogen.GeneratedFile, messageData *messageData) error {
	f.P()
	f.P("### ", messageData.messageName, " message")
	f.P()

	if messageData.description != "" {
		f.P(messageData.description)
		f.P()
	}

	tableData, err := messageData.GetTableData()
	if err != nil {
		return err
	}

	tpl := "| %s | %s | %s | %s |"

	f.P(fmt.Sprintf(
		tpl,
		utils.PadRight("Name", " ", tableData.colLengths[0]),
		utils.PadRight("Type", " ", tableData.colLengths[1]),
		utils.PadRight("Required?", " ", tableData.colLengths[2]),
		utils.PadRight("Description", " ", tableData.colLengths[3]),
	))

	f.P(fmt.Sprintf(
		tpl,
		utils.PadRight("", "-", tableData.colLengths[0]),
		utils.PadRight("", "-", tableData.colLengths[1]),
		utils.PadRight("", "-", tableData.colLengths[2]),
		utils.PadRight("", "-", tableData.colLengths[3]),
	))

	for _, row := range tableData.rows {
		f.P(fmt.Sprintf(
			tpl,
			utils.PadRight(row.fieldName, " ", tableData.colLengths[0]),
			utils.PadRight(row.fieldType, " ", tableData.colLengths[1]),
			utils.PadRight(utils.BoolToTickOrCross(row.required), " ", tableData.colLengths[2]),
			utils.PadRight(row.description, " ", tableData.colLengths[3]),
		))
	}

	return nil
}

func (c *converter) writeMessageFieldValidationTable(f *protogen.GeneratedFile, messageData *messageData) error {
	for _, fieldName := range messageData.fieldsOrder {
		field, ok := messageData.fields[fieldName]
		if !ok {
			return fmt.Errorf("field %s not found", fieldName)
		}

		if field.options == nil {
			continue
		}

		ignoreEmpty := field.options.GetIgnoreEmpty()

		items := []string{}

		items = append(items, getStringValidationRules(field.options.GetString_(), ignoreEmpty)...)
		items = append(items, getInt32ValidationRules(field.options.GetInt32(), ignoreEmpty)...)
		items = append(items, getUInt32ValidationRules(field.options.GetUint32(), ignoreEmpty)...)
		items = append(items, getTimestampValidationRules(field.options.GetTimestamp(), ignoreEmpty)...)

		if len(items) == 0 {
			continue
		}

		f.P()
		f.P("#### `", messageData.messageName, ".", field.fieldName, "` validation")
		f.P()
		f.P("The following validation rules apply to the `", field.fieldName, "` field:")
		f.P()

		for _, item := range items {
			f.P("- ", item)
		}

	}

	return nil
}

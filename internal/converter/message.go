package converter

import (
	"fmt"
	"strings"

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
			field.fieldMessage = string(f.Enum.Desc.Name())

		case protoreflect.MessageKind:
			field.fieldType = "message"
			field.fieldMessage = string(f.Message.Desc.Name())

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
	f.P("## ", messageData.messageName)
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

	f.P()

	return nil
}

func (c *converter) writeMessageFieldValidationTable(f *protogen.GeneratedFile, messageData *messageData) error {
	for _, field := range messageData.fields {
		if field.options == nil {
			continue
		}

		ignoreEmpty := field.options.GetIgnoreEmpty()

		items := []string{}

		// TODO: There's got to be a better way to do this...
		if s := field.options.GetString_(); s != nil {
			if s.Const != nil {
				items = append(items, fmt.Sprintf("Must be `%s`", *s.Const))
			}

			if s.Contains != nil {
				items = append(items, fmt.Sprintf("Must contain `%s`", *s.Contains))
			}

			if s.In != nil {
				items = append(items, fmt.Sprintf("Must be one of: `%s`", strings.Join(s.In, "`, `")))
			}

			if s.Len != nil {
				msg := "Exactly %d character%s long"
				if ignoreEmpty {
					msg = "Must be empty, or exactly %d character%s long"
				}

				items = append(items, fmt.Sprintf(msg, *s.Len, utils.PluralSuffix(int(*s.Len), "s", "")))
			}

			if s.MinLen != nil {
				msg := "Must be at least %d character%s long"
				if ignoreEmpty {
					msg = "Must be empty, or at least %d character%s long"
				}

				items = append(items, fmt.Sprintf(msg, *s.MinLen, utils.PluralSuffix(int(*s.MinLen), "s", "")))
			}

			if s.MaxLen != nil {
				items = append(items, fmt.Sprintf("Must be %d or fewer character%s long", *s.MaxLen, utils.PluralSuffix(int(*s.MaxLen), "s", "")))
			}

			if s.LenBytes != nil {
				msg := "Exactly %d byte%s long"
				if ignoreEmpty {
					msg = "Must be empty, or exactly %d byte%s long"
				}

				items = append(items, fmt.Sprintf(msg, *s.LenBytes, utils.PluralSuffix(int(*s.LenBytes), "s", "")))
			}

			if s.MinBytes != nil {
				msg := "Must be at least %d byte%s long"
				if ignoreEmpty {
					msg = "Must be empty, or at least %d byte%s long"
				}

				items = append(items, fmt.Sprintf(msg, *s.MinBytes, utils.PluralSuffix(int(*s.MinBytes), "s", "")))
			}

			if s.NotContains != nil {
				items = append(items, fmt.Sprintf("Must not contain `%s`", *s.NotContains))
			}

			if s.NotIn != nil {
				items = append(items, fmt.Sprintf("Must not be one of: `%s`", strings.Join(s.NotIn, "`, `")))
			}

			if s.Pattern != nil {
				items = append(items, fmt.Sprintf("Must match the regex pattern `%s`", *s.Pattern))
			}

			if s.Prefix != nil {
				items = append(items, fmt.Sprintf("Must start with `%s`", *s.Prefix))
			}

			if s.Suffix != nil {
				items = append(items, fmt.Sprintf("Must end with `%s`", *s.Suffix))
			}

			if s.GetAddress() {
				items = append(items, "Must be a valid hostname or IP address")
			}

			if s.GetEmail() {
				items = append(items, "Must be a valid email address")
			}

			if s.GetHostname() {
				items = append(items, "Must be a valid hostname")
			}

			if s.GetIp() {
				items = append(items, "Must be a valid IP address")
			}

			if s.GetIpPrefix() {
				items = append(items, "Must be a valid IP prefix (eg, 20.0.0.0/16)")
			}

			if s.GetUri() {
				items = append(items, "Must be a valid URI")
			}

			if s.GetUriRef() {
				items = append(items, "Must be a valid URI reference")
			}

			if s.GetUuid() {
				items = append(items, "Must be a valid UUID")
			}
		}

		if t := field.options.GetTimestamp(); t != nil {
			if t.Const != nil {
				items = append(items, fmt.Sprintf("Must be `%s`", t.Const.String()))
			}

			if t.Within != nil {
				items = append(items, fmt.Sprintf("Must be within %s of now", t.Within.String()))
			}

			if t.GetGt() != nil {
				items = append(items, fmt.Sprintf("Must be after %s", t.GetGt().String()))
			}

			if t.GetGtNow() {
				items = append(items, "Must be after now")
			}

			if t.GetGte() != nil {
				items = append(items, fmt.Sprintf("Must be equal to or after %s", t.GetGte().String()))
			}

			if t.GetLt() != nil {
				items = append(items, fmt.Sprintf("Must be before %s", t.GetLt().String()))
			}

			if t.GetLtNow() {
				items = append(items, "Must be before now")
			}

			if t.GetLte() != nil {
				items = append(items, fmt.Sprintf("Must be equal to or before %s", t.GetLte().String()))
			}
		}

		if len(items) == 0 {
			continue
		}

		f.P("### `", messageData.messageName, ".", field.fieldName, "` validation")
		f.P()
		f.P("The following validation rules apply to the `", field.fieldName, "` field:")
		f.P()

		for _, item := range items {
			f.P("- ", item)
		}

		f.P()
	}

	return nil
}

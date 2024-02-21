package converter

import (
	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

func (c *converter) processEnum(g *protogen.GeneratedFile, enum *protogen.Enum) error {
	var err error

	data := &enumData{
		enum:     enum,
		enumName: string(enum.Desc.Name()),
		description: utils.FlattenComment(
			enum.Comments.Leading.String() + "\n\n" + enum.Comments.Trailing.String(),
		),
		values:      map[string]*enumValue{},
		valuesOrder: []string{},
	}

	for _, f := range enum.Values {
		value := &enumValue{
			valueName: string(f.Desc.Name()),
			description: utils.FlattenComment(
				f.Comments.Leading.String() + " " + f.Comments.Trailing.String(),
			),
		}

		data.values[value.valueName] = value
		data.valuesOrder = append(data.valuesOrder, value.valueName)
	}

	err = c.writeEnumFieldsTable(g, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *converter) writeEnumFieldsTable(g *protogen.GeneratedFile, data *enumData) error {
	g.P()
	g.P("### ", data.enumName, " enum")
	g.P()

	if data.description != "" {
		g.P(data.description)
		g.P()
	}

	tableData, err := data.GetTableData()
	if err != nil {
		return err
	}

	return utils.WriteTable(g, tableData)
}

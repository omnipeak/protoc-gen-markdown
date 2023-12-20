package converter

import (
	"fmt"

	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

func (c *converter) processService(g *protogen.GeneratedFile, svc *protogen.Service) error {
	data := &serviceData{
		svc:          svc,
		serviceName:  string(svc.Desc.Name()),
		description:  utils.FlattenComment(svc.Comments.Leading.String() + "\n\n" + svc.Comments.Trailing.String()),
		methods:      make(map[string]*serviceMethodData),
		methodsOrder: []string{},
	}

	for _, m := range svc.Methods {
		method := &serviceMethodData{
			methodName:  string(m.Desc.Name()),
			params:      []string{},
			description: utils.FlattenComment(m.Comments.Leading.String() + "\n\n" + m.Comments.Trailing.String()),
		}

		method.params = append(method.params, string(m.Input.Desc.Name()))
		method.response = string(m.Output.Desc.Name())

		data.methods[method.methodName] = method
		data.methodsOrder = append(data.methodsOrder, method.methodName)
	}

	c.writeServiceFieldsTable(g, data)

	return nil
}

func (c *converter) writeServiceFieldsTable(g *protogen.GeneratedFile, data *serviceData) error {
	g.P("## ", data.serviceName)
	g.P()

	if data.description != "" {
		g.P(data.description)
		g.P()
	}

	g.P("### Methods")
	g.P()

	tableData, err := data.GetTableData()
	if err != nil {
		return err
	}

	tpl := "| %s | %s | %s | %s |"

	g.P(fmt.Sprintf(
		tpl,
		utils.PadRight("Method", " ", tableData.colLengths[0]),
		utils.PadRight("Inputs", " ", tableData.colLengths[1]),
		utils.PadRight("Response", " ", tableData.colLengths[2]),
		utils.PadRight("Description", " ", tableData.colLengths[3]),
	))

	g.P(fmt.Sprintf(
		tpl,
		utils.PadRight("", "-", tableData.colLengths[0]),
		utils.PadRight("", "-", tableData.colLengths[1]),
		utils.PadRight("", "-", tableData.colLengths[2]),
		utils.PadRight("", "-", tableData.colLengths[3]),
	))

	for _, row := range tableData.rows {
		g.P(fmt.Sprintf(
			tpl,
			utils.PadRight(row.methodName, " ", tableData.colLengths[0]),
			utils.PadRight(row.inputs, " ", tableData.colLengths[1]),
			utils.PadRight(row.response, " ", tableData.colLengths[2]),
			utils.PadRight(row.description, " ", tableData.colLengths[3]),
		))
	}

	g.P()

	return nil
}

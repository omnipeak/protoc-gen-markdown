package converter

import (
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
			description: utils.FlattenComment(m.Comments.Leading.String() + "\n\n" + m.Comments.Trailing.String()),
		}

		method.request = string(m.Input.Desc.Name())
		method.requestMessage = m.Input

		method.response = string(m.Output.Desc.Name())
		method.responseMessage = m.Output

		data.methods[method.methodName] = method
		data.methodsOrder = append(data.methodsOrder, method.methodName)
	}

	c.writeServiceFieldsTable(g, data)

	return nil
}

func (c *converter) writeServiceFieldsTable(g *protogen.GeneratedFile, data *serviceData) error {
	g.P()
	g.P("### ", data.serviceName, " service")
	g.P()

	if data.description != "" {
		g.P(data.description)
		g.P()
	}

	g.P("#### Methods")
	g.P()

	tableData, err := data.GetTableData()
	if err != nil {
		return err
	}

	return utils.WriteTable(g, tableData)
}

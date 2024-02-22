package converter

import (
	"fmt"
	"strings"

	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

type serviceData struct {
	svc          *protogen.Service
	serviceName  string
	description  string
	methods      map[string]*serviceMethodData
	methodsOrder []string
}

func (data *serviceData) GetTableData() (*utils.TableData, error) {
	tableData := &utils.TableData{
		Headers: []string{
			"Method",
			"Request",
			"Response",
			"Description",
		},
		Rows: [][]string{},
	}

	pathPrefix := strings.Repeat(
		"../",
		strings.Count(
			data.svc.Location.SourceFile,
			"/",
		),
	)

	if len(data.methods) != len(data.methodsOrder) {
		return nil, fmt.Errorf("methods and methodsOrder length mismatch")
	}

	for _, methodName := range data.methodsOrder {
		method, ok := data.methods[methodName]
		if !ok {
			return nil, fmt.Errorf("method %s not found", methodName)
		}

		link := ""

		if method.requestMessage.Location.SourceFile != data.svc.Location.SourceFile {
			link += pathPrefix + strings.ReplaceAll(
				string(method.requestMessage.Location.SourceFile),
				".proto",
				".md",
			)
		}
		link += "#" + strings.ToLower(method.request) + "-message"

		request := fmt.Sprintf("[`%s`](%s)", method.request, link)

		link = ""
		if method.responseMessage.Location.SourceFile != data.svc.Location.SourceFile {
			link += pathPrefix + strings.ReplaceAll(
				string(method.responseMessage.Location.SourceFile),
				".proto",
				".md",
			)
		}
		link += "#" + strings.ToLower(method.response) + "-message"
		response := fmt.Sprintf("[`%s`](%s)", method.response, link)

		row := []string{
			fmt.Sprintf("`%s`", method.methodName),
			request,
			response,
			method.description,
		}

		tableData.Rows = append(tableData.Rows, row)
	}

	return tableData, nil
}

type serviceMethodData struct {
	methodName      string
	request         string
	requestMessage  *protogen.Message
	response        string
	responseMessage *protogen.Message
	description     string
}

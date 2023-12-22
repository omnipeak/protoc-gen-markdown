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

func (sd *serviceData) GetTableData() (*serviceTableData, error) {
	data := &serviceTableData{
		colLengths: []int{9, 40, 25, 35},
		rows:       []*serviceTableMethodRow{},
	}

	pathPrefix := strings.Repeat(
		"../",
		strings.Count(
			sd.svc.Location.SourceFile,
			"/",
		),
	)

	for _, methodName := range sd.methodsOrder {
		method, ok := sd.methods[methodName]
		if !ok {
			return nil, fmt.Errorf("method %s not found", methodName)
		}

		row := &serviceTableMethodRow{
			methodName:  fmt.Sprintf("`%s`", method.methodName),
			description: method.description,
		}

		link := ""

		if method.requestMessage.Location.SourceFile != sd.svc.Location.SourceFile {
			link += pathPrefix + strings.ReplaceAll(
				string(method.requestMessage.Location.SourceFile),
				".proto",
				".md",
			)
		}
		link += "#" + strings.ToLower(method.request) + "-message"

		row.request = fmt.Sprintf("[`%s`](%s)", method.request, link)

		link = ""
		if method.responseMessage.Location.SourceFile != sd.svc.Location.SourceFile {
			link += pathPrefix + strings.ReplaceAll(
				string(method.responseMessage.Location.SourceFile),
				".proto",
				".md",
			)
		}
		link += "#" + strings.ToLower(method.response) + "-message"
		row.response = fmt.Sprintf("[`%s`](%s)", method.response, link)

		utils.StringGTLengthHelper(&data.colLengths[0], row.methodName)
		utils.StringGTLengthHelper(&data.colLengths[1], row.request)
		utils.StringGTLengthHelper(&data.colLengths[2], row.response)
		utils.StringGTLengthHelper(&data.colLengths[3], row.description)

		data.rows = append(data.rows, row)
	}

	return data, nil
}

type serviceMethodData struct {
	methodName      string
	request         string
	requestMessage  *protogen.Message
	response        string
	responseMessage *protogen.Message
	description     string
}

type serviceTableData struct {
	colLengths []int
	rows       []*serviceTableMethodRow
}

type serviceTableMethodRow struct {
	methodName  string
	request     string
	response    string
	description string
}

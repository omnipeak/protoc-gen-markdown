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

	for _, methodName := range sd.methodsOrder {
		method, ok := sd.methods[methodName]
		if !ok {
			return nil, fmt.Errorf("method %s not found", methodName)
		}

		row := &serviceTableMethodRow{
			methodName:  fmt.Sprintf("`%s`", method.methodName),
			description: method.description,
		}

		inputBits := []string{}
		for _, p := range method.params {
			inputBits = append(inputBits, fmt.Sprintf("[`%s`](#%s)", p, strings.ToLower(p)))
		}

		row.inputs = strings.Join(inputBits, ", ")
		row.response = fmt.Sprintf("[`%s`](#%s)", method.response, strings.ToLower(method.response))

		utils.StringGTLengthHelper(&data.colLengths[0], row.methodName)
		utils.StringGTLengthHelper(&data.colLengths[1], row.inputs)
		utils.StringGTLengthHelper(&data.colLengths[2], row.response)
		utils.StringGTLengthHelper(&data.colLengths[3], row.description)

		data.rows = append(data.rows, row)
	}

	return data, nil
}

type serviceMethodData struct {
	methodName  string
	params      []string
	response    string
	description string
}

type serviceTableData struct {
	colLengths []int
	rows       []*serviceTableMethodRow
}

type serviceTableMethodRow struct {
	methodName  string
	inputs      string
	response    string
	description string
}

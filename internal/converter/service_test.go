package converter

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func getTestServiceData() *serviceData {
	return &serviceData{
		svc:         &protogen.Service{},
		serviceName: "TestService",
		description: "This is a test service",
		methods: map[string]*serviceMethodData{
			"Method1": {
				methodName:  "Method1",
				request:     "Request1",
				response:    "Response1",
				description: "This is test method 1's description",
			},
			"Method2": {
				methodName:  "Method2",
				request:     "Request2",
				response:    "Response2",
				description: "This is test method 2's description",
			},
		},
		methodsOrder: []string{"Method1", "Method2"},
	}
}

func getTestServiceTableData() *serviceTableData {
	return &serviceTableData{
		colLengths: []int{9, 40, 25, 35},
		rows: []*serviceTableMethodRow{
			{
				methodName:  "`Method1`",
				request:     "[`Request1`](#param1)",
				response:    "[`Response1`](#response1)",
				description: "This is test method 1's description",
			},
			{
				methodName:  "`Method2`",
				request:     "[`Request2`](#param2)",
				response:    "[`Response2`](#response2)",
				description: "This is test method 2's description",
			},
		},
	}
}

func getTestServiceMarkdownResult() string {
	return "## TestService\n\n" +
		"This is a test service\n\n" +
		"### Methods\n\n" +
		"| Method    | Inputs                | Response                  | Description                         |\n" +
		"| --------- | --------------------- | ------------------------- | ----------------------------------- |\n" +
		"| `Method1` | [`Request1`](#param1) | [`Response1`](#response1) | This is test method 1's description |\n" +
		"| `Method2` | [`Request2`](#param2) | [`Response2`](#response2) | This is test method 2's description |\n\n"
}

func (ts *ConverterTestSuite) TestGetServiceTableData() {
	data := getTestServiceData()

	actualTableData, err := data.GetTableData()
	ts.NoError(err)

	expectedTableData := getTestServiceTableData()

	ts.Equal(expectedTableData, actualTableData)
}

func (ts *ConverterTestSuite) TestWriteServiceFieldsTable() {
	// Create a mock GeneratedFile
	g := &protogen.GeneratedFile{}

	data := getTestServiceData()

	c := New()

	err := c.writeServiceFieldsTable(g, data)
	ts.NoError(err)

	actualContent, err := g.Content()
	ts.NoError(err)

	ts.Equal(getTestServiceMarkdownResult(), string(actualContent))
}

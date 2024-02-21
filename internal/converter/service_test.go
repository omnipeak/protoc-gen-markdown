package converter

import (
	"github.com/omnipeak/protoc-gen-markdown/internal/utils"
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
				requestMessage: &protogen.Message{
					Location: protogen.Location{
						SourceFile: "test.proto",
					},
				},
				responseMessage: &protogen.Message{
					Location: protogen.Location{
						SourceFile: "test2.proto",
					},
				},
			},
			"Method2": {
				methodName:  "Method2",
				request:     "Request2",
				response:    "Response2",
				description: "This is test method 2's description",
				requestMessage: &protogen.Message{
					Location: protogen.Location{
						SourceFile: "test3.proto",
					},
				},
				responseMessage: &protogen.Message{
					Location: protogen.Location{
						SourceFile: "test4.proto",
					},
				},
			},
		},
		methodsOrder: []string{"Method1", "Method2"},
	}
}

func getTestServiceTableData() *utils.TableData {
	return &utils.TableData{
		Headers: []string{
			"Method",
			"Request",
			"Response",
			"Description",
		},
		Rows: [][]string{
			{
				"`Method1`",
				"[`Request1`](test.md#request1-message)",
				"[`Response1`](test2.md#response1-message)",
				"This is test method 1's description",
			},
			{
				"`Method2`",
				"[`Request2`](test3.md#request2-message)",
				"[`Response2`](test4.md#response2-message)",
				"This is test method 2's description",
			},
		},
	}
}

func getTestServiceMarkdownResult() string {
	return "\n" +
		"### TestService service\n\n" +
		"This is a test service\n\n" +
		"#### Methods\n\n" +
		"| Method    | Request                                 | Response                                  | Description                         |\n" +
		"| --------- | --------------------------------------- | ----------------------------------------- | ----------------------------------- |\n" +
		"| `Method1` | [`Request1`](test.md#request1-message)  | [`Response1`](test2.md#response1-message) | This is test method 1's description |\n" +
		"| `Method2` | [`Request2`](test3.md#request2-message) | [`Response2`](test4.md#response2-message) | This is test method 2's description |\n"
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

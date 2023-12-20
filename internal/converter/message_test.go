package converter

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/compiler/protogen"
)

func getTestMessageData() *messageData {
	minLen1 := uint64(1)
	maxLen1 := uint64(10)

	return &messageData{
		message:     &protogen.Message{},
		messageName: "TestMessage",
		description: "This is a test message",
		fields: map[string]*messageField{
			"field1": {
				fieldName:    "field1",
				fieldType:    "string",
				fieldMessage: "",
				description:  "This is a test field",
				required:     true,
				isList:       true,
				options: &validate.FieldConstraints{
					Required: true,
					Type: &validate.FieldConstraints_String_{
						String_: &validate.StringRules{
							MinLen: &minLen1,
							MaxLen: &maxLen1,
						},
					},
				},
			},
			"field2": {
				fieldName:    "field2",
				fieldType:    "message",
				fieldMessage: "TestMessage2",
				description:  "This is another test field",
				required:     false,
				isList:       false,
			},
		},
		fieldsOrder: []string{"field1", "field2"},
	}
}

func getTestMessageTableData() *messageTableData {
	return &messageTableData{
		colLengths: []int{8, 31, 9, 26},
		rows: []*messageTableFieldRow{
			{
				fieldName:   "`field1`",
				fieldType:   "`string[]`",
				description: "This is a test field",
				required:    true,
			},
			{
				fieldName:   "`field2`",
				fieldType:   "[`TestMessage2`](#testmessage2)",
				description: "This is another test field",
				required:    false,
			},
		},
	}
}

func getTestMessageMarkdownResult() string {
	return "## TestMessage\n\n" +
		"This is a test message\n\n" +
		"| Name     | Type                            | Required? | Description                |\n" +
		"| -------- | ------------------------------- | --------- | -------------------------- |\n" +
		"| `field1` | `string[]`                      | ✅         | This is a test field       |\n" +
		"| `field2` | [`TestMessage2`](#testmessage2) | ❌         | This is another test field |\n\n"
}

func (ts *ConverterTestSuite) TestGetMessageTableData() {
	data := getTestMessageData()

	actualTableData, err := data.GetTableData()
	ts.NoError(err)

	expectedTableData := getTestMessageTableData()

	ts.Equal(expectedTableData, actualTableData)
}

func (ts *ConverterTestSuite) TestWriteMessageFieldsTable() {
	data := getTestMessageData()

	g := &protogen.GeneratedFile{}

	c := New()

	err := c.writeMessageFieldsTable(g, data)
	ts.NoError(err)

	actualContent, err := g.Content()
	ts.NoError(err)

	ts.Equal(getTestMessageMarkdownResult(), string(actualContent))
}

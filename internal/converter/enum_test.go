package converter

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func getTestEnumData() *enumData {
	return &enumData{
		enum:        &protogen.Enum{},
		enumName:    "TestEnum",
		description: "Test enum description",
		values: map[string]*enumValue{
			"TestValue1": {
				valueName:   "TestValue1",
				description: "Test value 1 description",
			},
			"TestValue2": {
				valueName:   "TestValue2",
				description: "Test value 2 description",
			},
		},
		valuesOrder: []string{"TestValue1", "TestValue2"},
	}
}

func getTestEnumTableData() *enumTableData {
	return &enumTableData{
		colLengths: []int{12, 24},
		rows: []*enumTableFieldRow{
			{
				valueName:   "`TestValue1`",
				description: "Test value 1 description",
			},
			{
				valueName:   "`TestValue2`",
				description: "Test value 2 description",
			},
		},
	}
}

func getTestEnumMarkdownResult() string {
	return "## TestEnum\n\n" +
		"Test enum description\n\n" +
		"| Value        | Description              |\n" +
		"| ------------ | ------------------------ |\n" +
		"| `TestValue1` | Test value 1 description |\n" +
		"| `TestValue2` | Test value 2 description |\n\n"
}

func (ts *ConverterTestSuite) TestGetEnumTableData() {
	data := getTestEnumData()

	tableData, err := data.GetTableData()
	ts.NoError(err)

	ts.Equal(getTestEnumTableData(), tableData)
}

func (ts *ConverterTestSuite) TestWriteEnumFieldsTable() {
	data := getTestEnumData()

	g := &protogen.GeneratedFile{}

	c := New()

	err := c.writeEnumFieldsTable(g, data)
	ts.NoError(err)

	actualContent, err := g.Content()
	ts.NoError(err)

	ts.Equal(getTestEnumMarkdownResult(), string(actualContent))
}

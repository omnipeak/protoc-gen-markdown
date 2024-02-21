package utils

import (
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type TableData struct {
	Headers []string
	Rows    [][]string

	colLengths []int
}

func stringSliceToRowString(s []string) string {
	return "| " + strings.Join(s, " | ") + " |"
}

func separatorString(lengths []int) string {
	sep := make([]string, len(lengths))
	for i := range sep {
		sep[i] = strings.Repeat("-", lengths[i])
	}

	return stringSliceToRowString(sep)
}

func WriteTable(g *protogen.GeneratedFile, inputData *TableData) error {
	cleanData := &TableData{
		Headers:    make([]string, len(inputData.Headers)),
		Rows:       make([][]string, len(inputData.Rows)),
		colLengths: make([]int, len(inputData.Headers)),
	}

	expectedLen := len(cleanData.Headers)
	for i, row := range inputData.Rows {
		rowLen := len(row)
		if expectedLen != rowLen {
			return errors.Errorf(
				"row %d length does not match headers length: %d != %d",
				i,
				rowLen,
				len(cleanData.Headers),
			)
		}
	}

	for i, header := range inputData.Headers {
		cleanData.Headers[i] = strings.TrimSpace(header)
		StringGTLengthHelper(&cleanData.colLengths[i], header)
	}

	for i, row := range inputData.Rows {
		cleanData.Rows[i] = make([]string, len(row))
		for j, cell := range row {
			cleanData.Rows[i][j] = strings.TrimSpace(cell)
			StringGTLengthHelper(&cleanData.colLengths[j], cell)
		}
	}

	headers, err := PadRightSlice(cleanData.Headers, " ", cleanData.colLengths)
	if err != nil {
		return errors.Wrap(err, "failed to pad headers")
	}

	g.P(stringSliceToRowString(headers))
	g.P(separatorString(cleanData.colLengths))

	for _, row := range cleanData.Rows {
		padded, err := PadRightSlice(row, " ", cleanData.colLengths)
		if err != nil {
			return errors.Wrap(err, "failed to pad row")
		}
		g.P(stringSliceToRowString(padded))
	}

	return nil
}

package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StringUtilsTestSuite struct {
	suite.Suite
}

func (ts *StringUtilsTestSuite) TestStringGTLengthHelper() {
	v := 0

	StringGTLengthHelper(&v, "")
	ts.Equal(0, v)

	StringGTLengthHelper(&v, "test")
	ts.Equal(4, v)

	StringGTLengthHelper(&v, "test2")
	ts.Equal(5, v)

	StringGTLengthHelper(&v, "test")
	ts.Equal(5, v)

	StringGTLengthHelper(&v, "")
	ts.Equal(5, v)
}

func (ts *StringUtilsTestSuite) TestPadRight() {
	ts.Equal("", PadRight("test", "", 0))
	ts.Equal("test", PadRight("test", "", 4))
	ts.Equal("test ", PadRight("test", "", 5))
	ts.Equal("test  ", PadRight("test", "", 6))
	ts.Equal("test   ", PadRight("test", "", 7))

	ts.Equal("", PadRight("test", " ", 0))
	ts.Equal("test", PadRight("test", " ", 4))
	ts.Equal("test ", PadRight("test", " ", 5))
	ts.Equal("test  ", PadRight("test", " ", 6))
	ts.Equal("test   ", PadRight("test", " ", 7))

	ts.Equal("", PadRight("test", "-", 0))
	ts.Equal("test", PadRight("test", "-", 4))
	ts.Equal("test-", PadRight("test", "-", 5))
	ts.Equal("test--", PadRight("test", "-", 6))
	ts.Equal("test---", PadRight("test", "-", 7))

	ts.Equal("", PadRight("test", "#", 0))
	ts.Equal("test", PadRight("test", "#", 4))
	ts.Equal("test#", PadRight("test", "#", 5))
	ts.Equal("test##", PadRight("test", "#", 6))
	ts.Equal("test###", PadRight("test", "#", 7))
}

func (ts *StringUtilsTestSuite) TestBoolToTickOrCross() {
	ts.Equal("✅", BoolToTickOrCross(true))
	ts.Equal("❌", BoolToTickOrCross(false))
}

func (ts *StringUtilsTestSuite) TestFlattenComment() {
	ts.Equal("", FlattenComment(""))
	ts.Equal("", FlattenComment(" "))
	ts.Equal("", FlattenComment("\t"))
	ts.Equal("", FlattenComment("\n"))
	ts.Equal("", FlattenComment("\r"))
	ts.Equal("", FlattenComment("\r\n"))
	ts.Equal("", FlattenComment("\n\r"))
	ts.Equal("", FlattenComment("\n\n"))
	ts.Equal("", FlattenComment("\r\r"))
	ts.Equal("", FlattenComment("\r\n\r\n"))

	expected := "Testing Testing"

	ts.Equal(expected, FlattenComment("Testing\n\nTesting"))
	ts.Equal(expected, FlattenComment("\n\nTesting\n\nTesting\n\n"))
	ts.Equal(expected, FlattenComment(" \n\n  Testing\n\nTesting\t\n\n "))
	ts.Equal(expected, FlattenComment("  \t \n \n  Testing  \t\n \r\n\r\t Testing   \t\n \n "))
}

func (ts *StringUtilsTestSuite) TestPluralSuffix() {
	ts.Equal("s", PluralSuffix(0, "s", ""))
	ts.Equal("", PluralSuffix(1, "s", ""))
	ts.Equal("s", PluralSuffix(2, "s", ""))
	ts.Equal("s", PluralSuffix(3, "s", ""))
}

func (ts *StringUtilsTestSuite) TestPluralSuffixWithDifferentSingularSuffix() {
	ts.Equal("s", PluralSuffix(0, "s", "x"))
	ts.Equal("x", PluralSuffix(1, "s", "x"))
	ts.Equal("s", PluralSuffix(2, "s", "x"))
	ts.Equal("s", PluralSuffix(3, "s", "x"))
}

func (ts *StringUtilsTestSuite) TestPluralSuffixWithDifferentPluralSuffix() {
	ts.Equal("x", PluralSuffix(0, "x", "s"))
	ts.Equal("s", PluralSuffix(1, "x", "s"))
	ts.Equal("x", PluralSuffix(2, "x", "s"))
	ts.Equal("x", PluralSuffix(3, "x", "s"))
}

func TestStringUtils(t *testing.T) {
	suite.Run(t, new(StringUtilsTestSuite))
}

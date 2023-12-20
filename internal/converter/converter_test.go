package converter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConverterTestSuite struct {
	suite.Suite
}

func TestConverter(t *testing.T) {
	suite.Run(t, new(ConverterTestSuite))
}

package converter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type ConverterTestSuite struct {
	suite.Suite
}

func getTestPlugin() (*protogen.Plugin, error) {
	return (&protogen.Options{}).New(&pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{
			{
				Name:    proto.String("example/protos/models.proto"),
				Package: proto.String("example.protos"),
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("example/protos"),
				},
			},
			{
				Name:    proto.String("example/protos/service.proto"),
				Package: proto.String("example.protos"),
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("example/protos"),
				},
			},
		},
		FileToGenerate: []string{
			"example/protos/models.proto",
			"example/protos/service.proto",
		},
	})
}

func getTestConverter() (*converter, error) {
	pl, err := getTestPlugin()
	if err != nil {
		return nil, err
	}

	c := New()
	pl.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, file := range pl.Files {
		if err := registerAllExtensions(c.extTypes, file.Desc); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (ts *ConverterTestSuite) TestConverter() {
	testConverter, err := getTestConverter()
	ts.NoError(err)

	testPlugin, err := getTestPlugin()
	ts.NoError(err)

	err = testConverter.generate(testPlugin)
	ts.NoError(err)

	for _, name := range []string{"models", "service"} {
		f, ok := testConverter.files["example/protos/"+name+".md"]
		if !ok {
			ts.FailNow("example/protos/" + name + ".md not found in generated files")
		}

		bytes, err := f.Content()
		ts.NoError(err)

		filePath := filepath.Join("..", "..", "tests", "data", name+".md")
		expectedContent, err := os.ReadFile(filePath)
		ts.NoError(err)

		ts.Equal(string(expectedContent), string(bytes))
	}
}

func TestConverter(t *testing.T) {
	suite.Run(t, new(ConverterTestSuite))
}

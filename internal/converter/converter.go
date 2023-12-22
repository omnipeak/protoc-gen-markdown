package converter

import (
	"flag"
	"log"
	"path"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type converter struct {
	outputMermaidDiagrams bool

	extTypes *protoregistry.Types
	packages map[string]*packageData
}

func New() *converter {
	return &converter{
		extTypes: new(protoregistry.Types),
		packages: map[string]*packageData{},
	}
}

func Run() {
	c := New()

	flags := flag.FlagSet{
		Usage: func() {
			log.Fatal(
				"Usage:\n" +
					"  protoc --markdown_out=mermaid=true:./path/to/output/dir/ foo.proto\n" +
					"  protoc --markdown_out=mermaid=false:./path/to/output/dir/ foo.proto\n\n" +
					"markdown_out params:\n" +
					"  mermaid: true / false\n",
			)
		},
	}

	flags.BoolVar(&c.outputMermaidDiagrams, "mermaid", false, "Generate mermaid diagrams")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(c.generate)
}

func (c *converter) generate(plugin *protogen.Plugin) error {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, file := range plugin.Files {
		if err := registerAllExtensions(c.extTypes, file.Desc); err != nil {
			return err
		}

		if file.Generate {
			if err := c.processFile(plugin, file); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *converter) addFile(file *protogen.File) {
	pkgName := string(file.Desc.Package())
	if _, ok := c.packages[pkgName]; !ok {
		c.packages[pkgName] = &packageData{
			enums:      map[string]*protogen.Enum{},
			extensions: map[string]*protogen.Extension{},
			files:      map[string]*protogen.File{},
			messages:   map[string]*protogen.Message{},
			services:   map[string]*protogen.Service{},
		}
	}

	c.packages[pkgName].files[file.Desc.Path()] = file

	for _, enm := range file.Enums {
		c.packages[pkgName].enums[string(enm.Desc.Name())] = enm
	}

	for _, ext := range file.Extensions {
		c.packages[pkgName].extensions[string(ext.Desc.Name())] = ext
	}

	for _, msg := range file.Messages {
		c.packages[pkgName].messages[string(msg.Desc.Name())] = msg
	}

	for _, svc := range file.Services {
		c.packages[pkgName].services[string(svc.Desc.Name())] = svc
	}
}

func (c *converter) processFile(plugin *protogen.Plugin, file *protogen.File) error {
	c.addFile(file)

	g := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".md", file.GoImportPath)

	fileName := cases.Title(language.Und, cases.NoLower).String(
		strings.ReplaceAll(
			strings.Replace(path.Base(file.Desc.Path()), ".proto", "", 1),
			"_",
			" ",
		),
	)

	writeServices := file.Services != nil && len(file.Services) > 0
	writeMessages := file.Messages != nil && len(file.Messages) > 0
	writeEnums := file.Enums != nil && len(file.Enums) > 0

	if !writeServices && !writeMessages && !writeEnums {
		return nil
	}

	g.P("# ", fileName)
	g.P()
	g.P("## Table of Contents")

	if writeServices {
		g.P()
		g.P("- [Services](#services)")

		for _, svc := range file.Services {
			g.P("  - [", svc.Desc.Name(), "](#", strings.ToLower(string(svc.Desc.Name())), "-service)")
		}
	}

	if writeMessages {
		g.P()
		g.P("- [Messages](#messages)")

		for _, msg := range file.Messages {
			g.P("  - [", msg.Desc.Name(), "](#", strings.ToLower(string(msg.Desc.Name())), "-message)")
		}
	}

	if writeEnums {
		g.P()
		g.P("- [Enums](#enums)")

		for _, enm := range file.Enums {
			g.P("  - [", enm.Desc.Name(), "](#", strings.ToLower(string(enm.Desc.Name())), "-enum)")
		}
	}

	if writeServices {
		g.P()
		g.P("## Services")

		for _, svc := range file.Services {
			if err := c.processService(g, svc); err != nil {
				return err
			}
		}
	}

	if writeMessages {
		g.P()
		g.P("## Messages")

		for _, msg := range file.Messages {
			if err := c.processMessage(g, msg); err != nil {
				return err
			}
		}
	}

	if writeEnums {
		g.P()
		g.P("## Enums")

		for _, enm := range file.Enums {
			if err := c.processEnum(g, enm); err != nil {
				return err
			}
		}
	}

	return nil
}

// func (c *converter) processField(g *protogen.GeneratedFile, field *protogen.Field) error {
// 	g.P("### ", field.Desc.Name())
// 	g.P()

// 	return nil
// }

// Recursively register all extensions into the provided protoregistry.Types,
// starting with the protoreflect.FileDescriptor and recursing into its MessageDescriptors,
// their nested MessageDescriptors, and so on.
//
// This leverages the fact that both protoreflect.FileDescriptor and protoreflect.MessageDescriptor
// have identical Messages() and Extensions() functions in order to recurse through a single function
func registerAllExtensions(extTypes *protoregistry.Types, descriptions interface {
	Messages() protoreflect.MessageDescriptors
	Extensions() protoreflect.ExtensionDescriptors
}) error {
	mds := descriptions.Messages()
	for i := 0; i < mds.Len(); i++ {
		registerAllExtensions(extTypes, mds.Get(i))
	}
	xds := descriptions.Extensions()
	for i := 0; i < xds.Len(); i++ {
		if err := extTypes.RegisterExtension(dynamicpb.NewExtensionType(xds.Get(i))); err != nil {
			return err
		}
	}
	return nil
}

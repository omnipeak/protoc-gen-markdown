package converter

import "google.golang.org/protobuf/compiler/protogen"

type packageData struct {
	enums      map[string]*protogen.Enum
	extensions map[string]*protogen.Extension
	files      map[string]*protogen.File
	messages   map[string]*protogen.Message
	services   map[string]*protogen.Service
}

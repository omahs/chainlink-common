package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"unicode"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
)

type GeneratedInfo struct {
	Package        string
	Config         Struct
	Input          *Struct
	Types          map[string]Struct
	CapabilityType capabilities.CapabilityType
	BaseName       string
	RootOutput     string
}

func (g GeneratedInfo) RootType() Struct {
	return g.Types[g.RootOutput]
}

type Struct struct {
	Name    string
	Outputs map[string]Field
}

type Field struct {
	Type        string
	NumSlice    int
	IsPrimitive bool
}

func StructsFromSrc(src, baseName string, tpe capabilities.CapabilityType) (GeneratedInfo, error) {
	fset := token.NewFileSet()

	// Parse the source code string
	node, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		return GeneratedInfo{}, err
	}
	pkg := node.Name.Name

	rawInfo := map[string]Struct{}

	ast.Inspect(node, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			s := Struct{
				Name:    ts.Name.Name,
				Outputs: map[string]Field{},
			}

			nestedPlain := false
			if structType, ok := ts.Type.(*ast.StructType); ok {
				for _, field := range structType.Fields.List {
					start := fset.Position(field.Type.Pos()).Offset
					end := fset.Position(field.Type.End()).Offset
					typeStr := src[start:end]
					if typeStr == "interface{}" {
						typeStr = "any"
					}
					f := Field{}

					f.Type = typeStr
					for strings.HasPrefix(f.Type, "[]") {
						f.NumSlice++
						f.Type = f.Type[2:]
					}
					f.IsPrimitive = unicode.IsLower(rune(f.Type[0]))
					s.Outputs[field.Names[0].Name] = f
				}

			} else {
				start := fset.Position(ts.Type.Pos()).Offset
				end := fset.Position(ts.Type.End()).Offset
				typeStr := src[start:end]
				nestedPlain = typeStr == "Plain"
			}

			// artifact used for deserializing
			if !nestedPlain {
				rawInfo[ts.Name.Name] = s
			}
		}
		return true
	})

	root := rawInfo[baseName]
	delete(rawInfo, baseName)
	configType := root.Outputs["Config"].Type
	config := rawInfo[configType]
	delete(rawInfo, configType)
	inputField, ok := root.Outputs["Inputs"]
	var input *Struct
	if ok {
		inputType := inputField.Type
		inputS := rawInfo[inputType]
		input = &inputS
		delete(rawInfo, inputType)
	}

	for k, _ := range rawInfo {
		if strings.HasPrefix(k, configType) || (input != nil && strings.HasPrefix(k, input.Name)) {
			delete(rawInfo, k)
		}
	}

	return GeneratedInfo{
		Package:        pkg,
		Config:         config,
		Types:          rawInfo,
		RootOutput:     root.Outputs["Outputs"].Type,
		BaseName:       baseName,
		CapabilityType: tpe,
		Input:          input,
	}, nil
}
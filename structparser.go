package structparser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"strings"
)

type Struct struct {
	Name   string
	Fields []Field
	Docs   []string
}

type Field struct {
	Name    string
	Type    string
	Tag     string
	Pointer bool
	Slice   bool
	Docs    []string
	Comment string
}

func ParseDirectory(fileOrDirectory string) ([]Struct, error) {
	return ParseDirectoryWithFilter(fileOrDirectory, nil)
}

func ParseDirectoryWithFilter(fileOrDirectory string, filter func(fs.FileInfo) bool) ([]Struct, error) {
	structs := make([]Struct, 0)

	fi, err := os.Stat(fileOrDirectory)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	dir := make(map[string]*ast.Package)
	switch mode := fi.Mode(); {
	case mode.IsDir():
		dir, err = parser.ParseDir(token.NewFileSet(), fileOrDirectory, filter, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
		if err != nil {
			return nil, err
		}
	case mode.IsRegular():
		tmp, err := parser.ParseFile(token.NewFileSet(), fileOrDirectory, nil, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
		if err != nil {
			return nil, err
		}
		dir[fileOrDirectory] = &ast.Package{
			Name:  tmp.Name.Name,
			Files: make(map[string]*ast.File),
		}

		dir[fileOrDirectory].Files[fileOrDirectory] = tmp
	}

	for _, pkg := range dir {
		for _, t := range doc.New(pkg, "", 0).Types {
			// safety
			if t == nil || t.Decl == nil {
				return nil, errors.New("t or t.Decl is nil")
			}
			for _, spec := range t.Decl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					return nil, errors.New("not a *ast.TypeSpec")
				}
				structType, ok := typeSpec.Type.(*ast.StructType)
				if ok {
					parsedStruct := Struct{
						Name:   t.Name,
						Fields: make([]Field, 0, len(structType.Fields.List)),
						Docs:   getDocsForStruct(t.Doc),
					}
					for _, fvalue := range structType.Fields.List {
						field := Field{
							Name:    fvalue.Names[0].Obj.Name,
							Type:    "",
							Tag:     "",
							Pointer: false,
							Slice:   false,
						}
						if fvalue.Doc != nil {
							field.Docs = getDocsForField(fvalue.Doc)
						}
						if fvalue.Comment != nil {
							field.Comment = cleanDocText(fvalue.Comment.Text())
						}
						if fvalue.Tag != nil {
							field.Tag = fvalue.Tag.Value
						}
						var err error
						field.Type, field.Slice, field.Pointer, err = getType(fvalue.Type)
						if err != nil {
							return nil, err
						}

						parsedStruct.Fields = append(parsedStruct.Fields, field)
					}

					structs = append(structs, parsedStruct)
				}
			}
		}
	}
	return structs, nil
}

func getDocsForStruct(doc string) []string {
	tmp := strings.Split(strings.Trim(doc, "\n"), "\n")
	docs := make([]string, 0, len(tmp))
	for _, v := range tmp {
		docs = append(docs, cleanDocText(v))
	}
	return docs
}

func getDocsForField(cg *ast.CommentGroup) []string {
	docs := make([]string, 0, len(cg.List))
	for _, v := range cg.List {
		docs = append(docs, cleanDocText(v.Text))
	}
	return docs
}

func cleanDocText(doc string) string {
	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	if strings.HasPrefix(doc, "// ") {
		doc = strings.Replace(doc, "// ", "", 1)
	} else if strings.HasPrefix(doc, "//") {
		doc = strings.Replace(doc, "//", "", 1)
	} else if strings.HasPrefix(doc, "/*") {
		doc = strings.Replace(doc, "/*", "", 1)
	}
	if strings.HasSuffix(doc, "*/") {
		doc = reverseString(strings.Replace(reverseString(doc), "/*", "", 1))
	}
	return strings.Trim(strings.Trim(doc, " "), "\n")
}

func justTypeString(a string, b, c bool, err error) string {
	return a
}

func getType(expr ast.Expr) (typeString string, isSlice, isPointer bool, err error) {
	switch expr.(type) {
	case *ast.Ident:
		x := expr.(*ast.Ident)
		return x.Name, false, false, nil
	case *ast.SelectorExpr:
		x := expr.(*ast.SelectorExpr)
		return x.X.(*ast.Ident).Name + "." + x.Sel.Name, false, false, nil
	case *ast.ArrayType:
		tmp := expr.(*ast.ArrayType)
		if tmp.Len != nil {
			tmpLen, ok := tmp.Len.(*ast.BasicLit)
			if !ok {
				return "", false, false, errors.New("array len has unknown type")
			}
			return "[" + tmpLen.Value + "]" + justTypeString(getType(tmp.Elt)), true, false, nil
		}
		return "[]" + justTypeString(getType(tmp.Elt)), true, false, nil
	case *ast.MapType:
		tmp := expr.(*ast.MapType)
		return "map[" + justTypeString(getType(tmp.Key)) + "]" + justTypeString(getType(tmp.Value)), false, false, nil
	case *ast.StarExpr:
		return "*" + justTypeString(getType(expr.(*ast.StarExpr).X)), false, true, nil
	case *ast.FuncType:
		return "", false, false, fmt.Errorf("unsupported type for %#v", expr)
	case *ast.StructType:
		return "", false, false, fmt.Errorf("unsupported type for %#v", expr)
	case *ast.ChanType:
		tmp := expr.(*ast.ChanType)
		switch tmp.Dir {
		case ast.SEND:
			return "chan<- " + justTypeString(getType(tmp.Value)), false, false, nil
		case ast.RECV:
			return "<-chan " + justTypeString(getType(tmp.Value)), false, false, nil
		}
		return "chan " + justTypeString(getType(tmp.Value)), false, false, nil
	}
	return "", false, false, fmt.Errorf("unknown type for %#v", expr)
}

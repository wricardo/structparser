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

type Output struct {
	Structs   []Struct
	Functions []Function
}

type Struct struct {
	Name    string
	Fields  []Field
	Methods []Method
	Docs    []string
}

type Method struct {
	Receiver  string   // Receiver type (e.g., "*MyStruct" or "MyStruct")
	Name      string   // Method name
	Params    []Param  // List of method parameters
	Returns   []Param  // List of return values
	Docs      []string // Method documentation
	Signature string   // Full signature for easy comparisons
}

type Param struct {
	Name string // Name of the parameter or return value
	Type string // Type (e.g., "int", "*string")
}

type Field struct {
	Name    string
	Type    string
	Tag     string
	Private bool
	Pointer bool
	Slice   bool
	Docs    []string
	Comment string
}

type Function struct {
	Name      string
	Params    []Param
	Returns   []Param
	Docs      []string
	Signature string
}

func ParseDirectory(fileOrDirectory string) (*Output, error) {
	return ParseDirectoryWithFilter(fileOrDirectory, nil)
}

func ParseString(fileContent string) (*Output, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", fileContent, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}

	packages := map[string]*ast.Package{
		"": {
			Name:  file.Name.Name,
			Files: map[string]*ast.File{"": file},
		},
	}

	return extractStructsFromPackages(packages)
}

func ParseDirectoryWithFilter(fileOrDirectory string, filter func(fs.FileInfo) bool) (*Output, error) {
	fi, err := os.Stat(fileOrDirectory)
	if err != nil {
		return nil, err
	}

	var packages map[string]*ast.Package
	fset := token.NewFileSet()

	switch mode := fi.Mode(); {
	case mode.IsDir():
		packages, err = parser.ParseDir(fset, fileOrDirectory, filter, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
		if err != nil {
			return nil, err
		}
	case mode.IsRegular():
		file, err := parser.ParseFile(fset, fileOrDirectory, nil, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
		if err != nil {
			return nil, err
		}
		packages = map[string]*ast.Package{
			fileOrDirectory: {
				Name:  file.Name.Name,
				Files: map[string]*ast.File{fileOrDirectory: file},
			},
		}
	}

	return extractStructsFromPackages(packages)
}

// func ParseDirectoryWithFilter(fileOrDirectory string, filter func(fs.FileInfo) bool) ([]Struct, error) {
// 	structs := make([]Struct, 0)

// 	fi, err := os.Stat(fileOrDirectory)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	dir := make(map[string]*ast.Package)
// 	switch mode := fi.Mode(); {
// 	case mode.IsDir():
// 		dir, err = parser.ParseDir(token.NewFileSet(), fileOrDirectory, filter, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
// 		if err != nil {
// 			return nil, err
// 		}
// 	case mode.IsRegular():
// 		tmp, err := parser.ParseFile(token.NewFileSet(), fileOrDirectory, nil, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
// 		if err != nil {
// 			return nil, err
// 		}
// 		dir[fileOrDirectory] = &ast.Package{
// 			Name:  tmp.Name.Name,
// 			Files: make(map[string]*ast.File),
// 		}

// 		dir[fileOrDirectory].Files[fileOrDirectory] = tmp
// 	}

// 	for _, pkg := range dir {
// 		tmp := doc.New(pkg, "", doc.AllDecls|doc.AllMethods)
// 		for _, t := range tmp.Types {
// 			// safety
// 			if t == nil || t.Decl == nil {
// 				return nil, errors.New("t or t.Decl is nil")
// 			}
// 			for _, spec := range t.Decl.Specs {
// 				typeSpec, ok := spec.(*ast.TypeSpec)
// 				if !ok {
// 					return nil, errors.New("not a *ast.TypeSpec")
// 				}
// 				structType, ok := typeSpec.Type.(*ast.StructType)
// 				if ok {
// 					parsedStruct := Struct{
// 						Name:    t.Name,
// 						Fields:  make([]Field, 0, len(structType.Fields.List)),
// 						Docs:    getDocsForStruct(t.Doc),
// 						Methods: make([]Method, 0),
// 					}
// 					for _, fvalue := range structType.Fields.List {
// 						name := ""
// 						if len(fvalue.Names) > 0 {
// 							name = fvalue.Names[0].Obj.Name
// 						}
// 						field := Field{
// 							Name:    name,
// 							Type:    "",
// 							Tag:     "",
// 							Pointer: false,
// 							Slice:   false,
// 						}
// 						if len(field.Name) > 0 {
// 							field.Private = strings.ToLower(string(field.Name[0])) == string(field.Name[0])
// 						}

// 						if fvalue.Doc != nil {
// 							field.Docs = getDocsForField(fvalue.Doc)
// 						}
// 						if fvalue.Comment != nil {
// 							field.Comment = cleanDocText(fvalue.Comment.Text())
// 						}
// 						if fvalue.Tag != nil {
// 							field.Tag = strings.Trim(fvalue.Tag.Value, "`")
// 						}
// 						var err error
// 						field.Type, field.Slice, field.Pointer, err = getType(fvalue.Type)
// 						if err != nil {
// 							return nil, err
// 						}

// 						parsedStruct.Fields = append(parsedStruct.Fields, field)
// 					}

// 					structs = append(structs, parsedStruct)
// 				}
// 			}
// 			for _, spec := range t.Methods {
// 				funcDecl := spec.Decl
// 				receiver, _, _, _ := getType(funcDecl.Recv.List[0].Type)
// 				method := Method{
// 					Name:     funcDecl.Name.Name,
// 					Receiver: receiver,
// 				}
// 				tmpArgs := []string{}
// 				for _, v := range funcDecl.Type.Params.List {
// 					a, _, _, err := getType(v.Type)
// 					if err != nil {
// 						return nil, err
// 					}

// 					tmpNames := []string{}
// 					for _, n := range v.Names {
// 						tmpNames = append(tmpNames, n.Name)
// 					}
// 					tmpArgs = append(tmpArgs, strings.Join(tmpNames, ", ")+" "+a)
// 				}

// 				tmpReturns := []string{}
// 				if funcDecl != nil && funcDecl.Type != nil && funcDecl.Type.Results != nil && funcDecl.Type.Results.List != nil {
// 					for _, v := range funcDecl.Type.Results.List {
// 						a, _, _, err := getType(v.Type)
// 						if err != nil {
// 							return nil, err
// 						}
// 						tmpNames := []string{}
// 						for _, n := range v.Names {
// 							tmpNames = append(tmpNames, n.Name)
// 						}

// 						tmpReturns = append(tmpReturns, strings.Join(tmpNames, ", ")+" "+a)
// 					}
// 				}
// 				method.Signature = method.Name + "(" + strings.Join(tmpArgs, ", ") + ") (" + strings.Join(tmpReturns, ", ") + ")"

// 				// find struct and add method
// 				for k, v := range structs {
// 					tmp := strings.Trim(method.Receiver, "*")
// 					if v.Name == tmp {
// 						structs[k].Methods = append(structs[k].Methods, method)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return structs, nil
// }

// Extract the logic to parse packages into a common function
func extractStructsFromPackages(packages map[string]*ast.Package) (*Output, error) {
	output := &Output{
		Structs:   make([]Struct, 0),
		Functions: make([]Function, 0), // Initialize functions slice
	}

	for _, pkg := range packages {
		docPkg := doc.New(pkg, "", doc.AllDecls|doc.AllMethods)
		for _, t := range docPkg.Types {
			if t == nil || t.Decl == nil {
				return nil, errors.New("t or t.Decl is nil")
			}

			// Extract structs
			for _, spec := range t.Decl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					return nil, errors.New("not a *ast.TypeSpec")
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if ok {
					parsedStruct := Struct{
						Name:    t.Name,
						Fields:  make([]Field, 0, len(structType.Fields.List)),
						Docs:    getDocsForStruct(t.Doc),
						Methods: make([]Method, 0),
					}

					for _, fvalue := range structType.Fields.List {
						name := ""
						if len(fvalue.Names) > 0 {
							name = fvalue.Names[0].Obj.Name
						}

						field := Field{
							Name:    name,
							Type:    "",
							Tag:     "",
							Pointer: false,
							Slice:   false,
						}

						if len(field.Name) > 0 {
							field.Private = strings.ToLower(string(field.Name[0])) == string(field.Name[0])
						}

						if fvalue.Doc != nil {
							field.Docs = getDocsForFieldAst(fvalue.Doc)
						}

						if fvalue.Comment != nil {
							field.Comment = cleanDocText(fvalue.Comment.Text())
						}

						if fvalue.Tag != nil {
							field.Tag = strings.Trim(fvalue.Tag.Value, "`")
						}

						var err error
						field.Type, field.Slice, field.Pointer, err = getType(fvalue.Type)
						if err != nil {
							return nil, err
						}

						parsedStruct.Fields = append(parsedStruct.Fields, field)
					}

					output.Structs = append(output.Structs, parsedStruct)
				}
			}

			// Extract methods associated with the struct
			for _, spec := range t.Methods {
				funcDecl := spec.Decl
				receiver, _, _, _ := getType(funcDecl.Recv.List[0].Type)

				method := Method{
					Name:     funcDecl.Name.Name,
					Receiver: receiver,
					Docs:     getDocsForField([]string{spec.Doc}),
				}

				// Parse function parameters
				params := []Param{}
				for _, param := range funcDecl.Type.Params.List {
					paramType, _, _, err := getType(param.Type)
					if err != nil {
						return nil, err
					}

					for _, name := range param.Names {
						params = append(params, Param{
							Name: name.Name,
							Type: paramType,
						})
					}
				}
				method.Params = params

				// Parse return types
				returns := []Param{}
				if funcDecl.Type.Results != nil {
					for _, result := range funcDecl.Type.Results.List {
						returnType, _, _, err := getType(result.Type)
						if err != nil {
							return nil, err
						}

						if len(result.Names) > 0 {
							for _, name := range result.Names {
								returns = append(returns, Param{
									Name: name.Name,
									Type: returnType,
								})
							}
						} else {
							returns = append(returns, Param{
								Name: "",
								Type: returnType,
							})
						}
					}
				}
				method.Returns = returns

				// Construct the full method signature for easy comparison
				paramStrings := []string{}
				for _, param := range method.Params {
					if param.Name != "" {
						paramStrings = append(paramStrings, param.Name+" "+param.Type)
					} else {
						paramStrings = append(paramStrings, param.Type)
					}
				}

				returnStrings := []string{}
				for _, ret := range method.Returns {
					if ret.Name != "" {
						returnStrings = append(returnStrings, ret.Name+" "+ret.Type)
					} else {
						returnStrings = append(returnStrings, ret.Type)
					}
				}

				method.Signature = fmt.Sprintf("%s(%s) (%s)",
					method.Name,
					strings.Join(paramStrings, ", "),
					strings.Join(returnStrings, ", "),
				)

				// Find the struct and add the method
				for k, v := range output.Structs {
					if strings.Trim(method.Receiver, "*") == v.Name {
						output.Structs[k].Methods = append(output.Structs[k].Methods, method)
					}
				}
			}
		}

		// Extract top-level functions (not associated with structs)
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					// Skip methods (they have a receiver)
					if funcDecl.Recv != nil {
						continue
					}

					function := Function{
						Name: funcDecl.Name.Name,
						Docs: getDocsForFieldAst(funcDecl.Doc),
					}

					// Parse function parameters
					params := []Param{}
					for _, param := range funcDecl.Type.Params.List {
						paramType, _, _, err := getType(param.Type)
						if err != nil {
							return nil, err
						}

						for _, name := range param.Names {
							params = append(params, Param{
								Name: name.Name,
								Type: paramType,
							})
						}
					}
					function.Params = params

					// Parse return types
					returns := []Param{}
					if funcDecl.Type.Results != nil {
						for _, result := range funcDecl.Type.Results.List {
							returnType, _, _, err := getType(result.Type)
							if err != nil {
								return nil, err
							}

							if len(result.Names) > 0 {
								for _, name := range result.Names {
									returns = append(returns, Param{
										Name: name.Name,
										Type: returnType,
									})
								}
							} else {
								returns = append(returns, Param{
									Name: "",
									Type: returnType,
								})
							}
						}
					}
					function.Returns = returns

					// Construct the full function signature for easy comparison
					paramStrings := []string{}
					for _, param := range function.Params {
						if param.Name != "" {
							paramStrings = append(paramStrings, param.Name+" "+param.Type)
						} else {
							paramStrings = append(paramStrings, param.Type)
						}
					}

					returnStrings := []string{}
					for _, ret := range function.Returns {
						if ret.Name != "" {
							returnStrings = append(returnStrings, ret.Name+" "+ret.Type)
						} else {
							returnStrings = append(returnStrings, ret.Type)
						}
					}

					function.Signature = fmt.Sprintf("%s(%s) (%s)",
						function.Name,
						strings.Join(paramStrings, ", "),
						strings.Join(returnStrings, ", "),
					)

					output.Functions = append(output.Functions, function)
				}
			}
		}
	}

	return output, nil
}

func getDocsForStruct(doc string) []string {
	trimmed := strings.Trim(doc, "\n")
	if trimmed == "" {
		return []string{}
	}
	tmp := strings.Split(trimmed, "\n")

	docs := make([]string, 0, len(tmp))
	for _, v := range tmp {
		docs = append(docs, cleanDocText(v))
	}
	return docs
}

func getDocsForFieldAst(cg *ast.CommentGroup) []string {
	if cg == nil {
		return []string{}
	}
	docs := make([]string, 0, len(cg.List))
	for _, v := range cg.List {
		docs = append(docs, cleanDocText(v.Text))
	}
	return docs
}

func getDocsForField(list []string) []string {
	docs := make([]string, 0, len(list))
	for _, v := range list {
		docs = append(docs, cleanDocText(v))
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
		return "/*func*/", false, false, nil
	case *ast.StructType:
		return "/*struct*/", false, false, nil
	case *ast.ChanType:
		tmp := expr.(*ast.ChanType)
		switch tmp.Dir {
		case ast.SEND:
			return "chan<- " + justTypeString(getType(tmp.Value)), false, false, nil
		case ast.RECV:
			return "<-chan " + justTypeString(getType(tmp.Value)), false, false, nil
		}
		return "chan " + justTypeString(getType(tmp.Value)), false, false, nil
	case *ast.Ellipsis:
		tmp := expr.(*ast.Ellipsis)
		return "..." + justTypeString(getType(tmp.Elt)), false, false, nil

	}
	return "", false, false, fmt.Errorf("unknown type for %#v", expr)
}

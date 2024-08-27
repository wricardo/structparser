package structparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDirectory(t *testing.T) {
	tmp, err := ParseDirectory("./example")
	require.NoError(t, err)

	parsed := newHelper(tmp)

	t.Run("FirstStruct", func(t *testing.T) {
		firstStruct := parsed.Struct("FirstStruct")
		require.Len(t, firstStruct.Docs, 2)
		assert.Equal(t, "FirstStruct this is the comment for the first struct.", firstStruct.Docs[0])
		assert.Equal(t, "This is new line.", firstStruct.Docs[1])

		var f Field

		f = firstStruct.Field("Int")
		assert.Equal(t, "Int", f.Name)
		assert.Equal(t, "int", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)
		assert.Equal(t, `json:"int" bson:"int"`, f.Tag)

		f = firstStruct.Field("Int8")
		assert.Equal(t, "Int8", f.Name)
		assert.Equal(t, "int8", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)
		assert.Equal(t, `bson:"int8"`, f.Tag)

		f = firstStruct.Field("Int16")
		assert.Equal(t, "Int16", f.Name)
		assert.Equal(t, "int16", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Int32")
		assert.Equal(t, "Int32", f.Name)
		assert.Equal(t, "int32", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Int64")
		assert.Equal(t, "Int64", f.Name)
		assert.Equal(t, "int64", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Uint")
		assert.Equal(t, "Uint", f.Name)
		assert.Equal(t, "uint", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Uintptr")
		assert.Equal(t, "Uintptr", f.Name)
		assert.Equal(t, "uintptr", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Uint8")
		assert.Equal(t, "Uint8", f.Name)
		assert.Equal(t, "uint8", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Uint16")
		assert.Equal(t, "Uint16", f.Name)
		assert.Equal(t, "uint16", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Uint32")
		assert.Equal(t, "Uint32", f.Name)
		assert.Equal(t, "uint32", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Uint64")
		assert.Equal(t, "Uint64", f.Name)
		assert.Equal(t, "uint64", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Float32")
		assert.Equal(t, "Float32", f.Name)
		assert.Equal(t, "float32", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Float64")
		assert.Equal(t, "Float64", f.Name)
		assert.Equal(t, "float64", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Complex64")
		assert.Equal(t, "Complex64", f.Name)
		assert.Equal(t, "complex64", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Complex128")
		assert.Equal(t, "Complex128", f.Name)
		assert.Equal(t, "complex128", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Byte")
		assert.Equal(t, "Byte", f.Name)
		assert.Equal(t, "byte", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Rune")
		assert.Equal(t, "Rune", f.Name)
		assert.Equal(t, "rune", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("String")
		assert.Equal(t, "String", f.Name)
		assert.Equal(t, "string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("SecondStruct")
		assert.Equal(t, "SecondStruct", f.Name)
		assert.Equal(t, "SecondStruct", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("ArrayInt")
		assert.Equal(t, "ArrayInt", f.Name)
		assert.Equal(t, "[3]int", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, true, f.Slice)

		f = firstStruct.Field("SliceString")
		assert.Equal(t, "SliceString", f.Name)
		assert.Equal(t, "[]string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, true, f.Slice)

		f = firstStruct.Field("SlicePointerString")
		assert.Equal(t, "SlicePointerString", f.Name)
		assert.Equal(t, "[]*string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, true, f.Slice)

		f = firstStruct.Field("PointerSliceString")
		assert.Equal(t, "PointerSliceString", f.Name)
		assert.Equal(t, "*[]string", f.Type)
		assert.Equal(t, true, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("PointerSlicePointerString")
		assert.Equal(t, "PointerSlicePointerString", f.Name)
		assert.Equal(t, "*[]*string", f.Type)
		assert.Equal(t, true, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("ChanString")
		assert.Equal(t, "ChanString", f.Name)
		assert.Equal(t, "chan string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("RChanString")
		assert.Equal(t, "RChanString", f.Name)
		assert.Equal(t, "<-chan string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("SChanString")
		assert.Equal(t, "SChanString", f.Name)
		assert.Equal(t, "chan<- string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("MapStringString")
		assert.Equal(t, "MapStringString", f.Name)
		assert.Equal(t, "map[string]string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("MapPointerStringString")
		assert.Equal(t, "MapPointerStringString", f.Name)
		assert.Equal(t, "map[*string]string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("MapPointerStringPointerString")
		assert.Equal(t, "MapPointerStringPointerString", f.Name)
		assert.Equal(t, "map[*string]*string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("PointerMapStringString")
		assert.Equal(t, "PointerMapStringString", f.Name)
		assert.Equal(t, "*map[string]string", f.Type)
		assert.Equal(t, true, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("PointerMapPointerStringPointerString")
		assert.Equal(t, "PointerMapPointerStringPointerString", f.Name)
		assert.Equal(t, "*map[*string]*string", f.Type)
		assert.Equal(t, true, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("Func")
		assert.Equal(t, "Func", f.Name)
		assert.Equal(t, "SomeFunc", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("PointerFunc")
		assert.Equal(t, "PointerFunc", f.Name)
		assert.Equal(t, "*SomeFunc", f.Type)
		assert.Equal(t, true, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("MapStringSliceString")
		assert.Equal(t, "MapStringSliceString", f.Name)
		assert.Equal(t, "map[string][]string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("MapStringSlicePointerString")
		assert.Equal(t, "MapStringSlicePointerString", f.Name)
		assert.Equal(t, "map[string][]*string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("MapPointerStringSlicePointerString")
		assert.Equal(t, "MapPointerStringSlicePointerString", f.Name)
		assert.Equal(t, "map[*string][]*string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("MapChanPointerStringStruct")
		assert.Equal(t, "MapChanPointerStringStruct", f.Name)
		assert.Equal(t, "map[chan *string]SecondStruct", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("SpecialString")
		assert.Equal(t, "SpecialString", f.Name)
		assert.Equal(t, "SpecialString", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("PackageStruct")
		assert.Equal(t, "PackageStruct", f.Name)
		assert.Equal(t, "other.Struct", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("PointerPackageStruct")
		assert.Equal(t, "PointerPackageStruct", f.Name)
		assert.Equal(t, "*other.Struct", f.Type)
		assert.Equal(t, true, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("SlicePointerPackageStruct")
		assert.Equal(t, "SlicePointerPackageStruct", f.Name)
		assert.Equal(t, "[]*other.Struct", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, true, f.Slice)

		f = firstStruct.Field("MapStringPackageStruct")
		assert.Equal(t, "MapStringPackageStruct", f.Name)
		assert.Equal(t, "map[string]other.Struct", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = firstStruct.Field("ChanPackagePointerStruct")
		assert.Equal(t, "ChanPackagePointerStruct", f.Name)
		assert.Equal(t, "chan *other.Struct", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

	})

	t.Run("FirstStruct", func(t *testing.T) {
		tmp := parsed.Struct("CommentsAndDocs")

		require.Len(t, tmp.Docs, 1)
		assert.Equal(t, "CommentsAndDocs this is the comment for the CommentsAndDocs struct.", tmp.Docs[0])

		t.Run("fields comments and docs", func(t *testing.T) {
			{
				docs := tmp.Field("SingleDoc").Docs
				assert.Len(t, docs, 1)
				assert.Equal(t, "this is line 1 of comment 001", docs[0])
			}
			{
				docs := tmp.Field("MultiLineDoc").Docs
				assert.Len(t, docs, 2)
				assert.Equal(t, "this is line 1 of comment 001", docs[0])
				assert.Equal(t, "this is line 2 of comment 002", docs[1])
			}
			{
				docs := tmp.Field("MixedSpacesDoc").Docs
				assert.Len(t, docs, 2)
				assert.Equal(t, "this is line 1 of comment 003", docs[0])
				assert.Equal(t, "this is line 2 of comment 004", docs[1])
			}
			{
				docs := tmp.Field("MixedTypesDoc").Docs
				assert.Len(t, docs, 2)
				assert.Equal(t, "this is line 1 of comment 005", docs[0])
				assert.Equal(t, "this is line 2 of comment 006", docs[1])
			}
			{
				docs := tmp.Field("DocAndComment").Docs
				com := tmp.Field("DocAndComment").Comment
				assert.Len(t, docs, 1)
				assert.Equal(t, "this is line 1 of comment 007", docs[0])
				assert.Equal(t, "comment 008", com)
			}
			{
				docs := tmp.Field("CommentNoSpaces").Docs
				com := tmp.Field("CommentNoSpaces").Comment
				assert.Len(t, docs, 0)
				assert.Equal(t, "comment abc", com)
			}
			{
				docs := tmp.Field("StarDoc").Docs
				com := tmp.Field("StarDoc").Comment
				assert.Len(t, docs, 1)
				assert.Equal(t, "this is line 1 of comment 009", docs[0])
				assert.Equal(t, "comment 010", com)
			}
			{
				docs := tmp.Field("CommentWithTag").Docs
				com := tmp.Field("CommentWithTag").Comment
				assert.Len(t, docs, 1)
				assert.Equal(t, "this is line 1 of comment 010", docs[0])
				assert.Equal(t, "comment 11", com)
			}
			{
				docs := tmp.Field("CrazyDoc").Docs
				assert.Len(t, docs, 9)
				assert.Equal(t, "001", docs[0])
				assert.Equal(t, "002", docs[1])
				assert.Equal(t, "003", docs[2])
				assert.Equal(t, "004", docs[3])
				assert.Equal(t, "005", docs[4])
				assert.Equal(t, "006", docs[5])
				assert.Equal(t, "007", docs[6])
				assert.Equal(t, "* 008 *", docs[7])
				assert.Equal(t, "009", docs[8])
			}
		})
	})

	t.Run("SecondStruct", func(t *testing.T) {
		secondStruct := parsed.Struct("SecondStruct")
		require.Len(t, secondStruct.Docs, 0)
	})
}

func TestCleanDocText(t *testing.T) {
	inputs := []struct {
		input  string
		expect string
	}{
		{input: "// something", expect: "something"},
		{input: "//   something  ", expect: "something"},
		{input: "//something", expect: "something"},
		{input: "///something", expect: "/something"},
		{input: "/*something*/", expect: "something"},
		{input: "/* something */", expect: "something"},
		{input: "/*  something  */", expect: "something"},
		{input: "/*something */", expect: "something"},
		{input: "/*  something*/", expect: "something"},
		{input: "/*/ something*/", expect: "/ something"},
		{input: "/*/ something /*/", expect: "/ something /"},
		{input: "\nsomething\n", expect: "something"},
	}
	for _, v := range inputs {
		assert.Equal(t, v.expect, cleanDocText(v.input))
	}

}

func TestFirstStructMethods(t *testing.T) {
	tmp, err := ParseDirectory("./example/")
	require.NoError(t, err)

	parsed := newHelper(tmp)

	t.Run("FirstStruct", func(t *testing.T) {
		firstStruct := parsed.Struct("FirstStruct")
		require.Len(t, firstStruct.Docs, 2)
		assert.Equal(t, "FirstStruct this is the comment for the first struct.", firstStruct.Docs[0])
		assert.Len(t, firstStruct.Methods, 2)
		assert.Equal(t, "MyOtherTestMethod(ctx context.Context, x string) (string, error)", firstStruct.Methods[0].Signature)
		assert.Equal(t, "MyTestMethod(ctx context.Context, x []string, y []string, z int) (a string, b string, c int)", firstStruct.Methods[1].Signature)
	})
}

func TestPrivateStruct(t *testing.T) {
	tmp, err := ParseDirectory("./example/")
	require.NoError(t, err)

	parsed := newHelper(tmp)

	t.Run("privateStruct", func(t *testing.T) {
		privateStruct := parsed.Struct("privateStruct")

		f := privateStruct.Field("String")
		assert.Equal(t, "String", f.Name)
		assert.Equal(t, "string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)
		assert.Empty(t, f.Tag)

		assert.Len(t, privateStruct.Methods, 1)
		assert.Equal(t, "MyPrivateStructMethod(ctx context.Context, x string) (string, error)", privateStruct.Methods[0].Signature)
	})
}

func TestParseString(t *testing.T) {
	// Test simple struct parsing
	t.Run("SimpleStruct", func(t *testing.T) {
		code := `
		package test
		// SimpleStruct represents a simple test case
		type SimpleStruct struct {
			Name  string  // Name is a string field
			Value int     // Value is an integer field
		}
		`
		output, err := ParseString(code)
		require.NoError(t, err)

		parsed := newHelper(output)
		structInfo := parsed.Struct("SimpleStruct")

		require.Len(t, structInfo.Docs, 1)
		assert.Equal(t, "SimpleStruct represents a simple test case", structInfo.Docs[0])

		var f Field
		f = structInfo.Field("Name")
		assert.Equal(t, "Name", f.Name)
		assert.Equal(t, "string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)
		assert.Equal(t, "Name is a string field", f.Comment)

		f = structInfo.Field("Value")
		assert.Equal(t, "Value", f.Name)
		assert.Equal(t, "int", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)
		assert.Equal(t, "Value is an integer field", f.Comment)
	})

	// Test parsing methods of a struct
	t.Run("StructWithMethods", func(t *testing.T) {
		code := `
		package test
		// StructWithMethods represents a struct with methods
		type StructWithMethods struct {}

		// Greet returns a greeting message
		func (s *StructWithMethods) Greet(name string) string {
			return "Hello " + name
		}

		// Sum adds two integers
		func (s *StructWithMethods) Sum(a, b int) int {
			return a + b
		}
		`
		output, err := ParseString(code)
		require.NoError(t, err)

		parsed := newHelper(output)
		structInfo := parsed.Struct("StructWithMethods")

		require.Len(t, structInfo.Methods, 2)

		assert.Equal(t, "Greet(name string) (string)", structInfo.Methods[0].Signature)
		assert.Equal(t, "Sum(a int, b int) (int)", structInfo.Methods[1].Signature)

		// Check method documentation
		assert.Equal(t, "Greet returns a greeting message", structInfo.Methods[0].Docs[0])
		assert.Equal(t, "Sum adds two integers", structInfo.Methods[1].Docs[0])
	})

	// Test struct with complex types
	t.Run("ComplexStruct", func(t *testing.T) {
		code := `
		package test
		// ComplexStruct represents a struct with various field types
		type ComplexStruct struct {
			SliceOfStrings []string
			PointerToInt   *int
			MapOfIntToStr  map[int]string
			FuncField      func(string) error
		}
		`
		output, err := ParseString(code)
		require.NoError(t, err)

		parsed := newHelper(output)
		structInfo := parsed.Struct("ComplexStruct")

		f := structInfo.Field("SliceOfStrings")
		assert.Equal(t, "SliceOfStrings", f.Name)
		assert.Equal(t, "[]string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, true, f.Slice)

		f = structInfo.Field("PointerToInt")
		assert.Equal(t, "PointerToInt", f.Name)
		assert.Equal(t, "*int", f.Type)
		assert.Equal(t, true, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = structInfo.Field("MapOfIntToStr")
		assert.Equal(t, "MapOfIntToStr", f.Name)
		assert.Equal(t, "map[int]string", f.Type)
		assert.Equal(t, false, f.Pointer)
		assert.Equal(t, false, f.Slice)

		f = structInfo.Field("FuncField")
		assert.Equal(t, "FuncField", f.Name)
		assert.Equal(t, "/*func*/", f.Type)
	})
}

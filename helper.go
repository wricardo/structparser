package structparser

type helper struct {
	structs map[string]helperField
}

func newHelper(out *Output) helper {
	tmp := helper{
		structs: make(map[string]helperField),
	}
	for _, v := range out.Structs {
		tmp.structs[v.Name] = newHelperField(v.Fields, v)
	}
	return tmp
}

func (h helper) Struct(name string) helperField {
	return h.structs[name]
}

type helperField struct {
	Struct
	fields map[string]Field
}

func newHelperField(fields []Field, structt Struct) helperField {
	tmp := helperField{
		Struct: structt,
		fields: make(map[string]Field, 0),
	}
	for _, v := range fields {
		tmp.fields[v.Name] = v
	}
	return tmp
}

func (h helperField) Field(name string) Field {
	return h.fields[name]

}

/*
 This file was autogenerated via
 ---------------------------------------------------------------
 gen-builtin --declarable --go-name bool --type Bool --name bool
 ---------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Bool("")

// Bool represents field of type bool
type Bool string

// Name returns field name
func (i Bool) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Bool) TypeName() string {
	return "bool"
}

// Register registers a field
func (i Bool) Register(registrator FieldRegistrator) {
	registrator.AddBool(i.Name())
}

// GoName returns Go's representation of this field's type
func (i Bool) GoName() string {
	return "bool"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["bool"] = func(fieldName string) Field {
		return Bool(fieldName)
	}

	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["bool"] = struct{}{}

	if backedBy == nil {
		backedBy = map[string]string{}
	}
	backedBy["bool"] = "bool"
}

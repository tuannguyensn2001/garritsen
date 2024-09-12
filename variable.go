package garritsen

type Variable struct {
}

func (v *Variable) Bool(defaultValue bool) bool {
	return defaultValue
}

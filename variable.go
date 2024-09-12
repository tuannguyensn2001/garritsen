package garritsen

type Variable struct {
	value interface{}
}

func (v *Variable) Bool(defaultValue bool) bool {
	val, ok := v.value.(bool)
	if ok {
		return val
	}

	return defaultValue
}

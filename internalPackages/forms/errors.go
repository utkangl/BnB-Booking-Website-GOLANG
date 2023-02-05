package forms

type errors map[string][]string

// Adding occured errors to a map with a hint message
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Returns the first error message
func (e errors) Get(field string) string {
	errString := e[field]
	if len(errString) == 0 {
		return " "
	}
	return errString[0]
}

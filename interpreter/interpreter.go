package interpreter

// TODO - finish everything that needs to be inside here
type InterpreterContext struct {
	Numbers map[string]float64
	Strings map[string]string
	Errors  []error
}

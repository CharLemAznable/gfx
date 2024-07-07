package gerrorx

type ErrorString string

func (e ErrorString) Error() string { return string(e) }

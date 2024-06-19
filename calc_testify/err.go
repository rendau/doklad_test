package calc

type calcErr string

func (e calcErr) Error() string {
	return string(e)
}

const (
	ErrDivideByZero = calcErr("division by zero")
)

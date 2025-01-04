package neuron_model

type Constraint int

const (
	Greater Constraint = iota + 1
	GreaterOrEqual
	Equal
	LessEqual
	Less
	In
)

func (d Constraint) String() string {
	return [...]string{">", ">=", "=", "<=", "<", "в"}[d-1]
}

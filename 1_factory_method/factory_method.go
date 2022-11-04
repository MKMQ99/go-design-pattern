package factory_method

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type PlusOperatorFactor struct{}

func (PlusOperatorFactor) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperatorFactor struct{}

func (MinusOperatorFactor) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperator struct {
	*OperatorBase
}

type MinusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

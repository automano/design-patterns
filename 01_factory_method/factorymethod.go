package factorymethod

// Operator interface declares the operations that all concrete operators must implement.
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory interface declares the factory method that returns a new operator instance.
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase is a concrete implementation of the Operator interface.
type OperatorBase struct {
	a int
	b int
}

// SetA sets the first operand.
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB sets the second operand.
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperator is a concrete implementation of the Operator interface.
type PlusOperator struct {
	*OperatorBase
}

// Result returns the result of the plus operation.
func (o *PlusOperator) Result() int {
	return o.a + o.b
}

// PlusOperatorFactory is a concrete implementation of the OperatorFactory interface.
type PlusOperatorFactory struct{}

// Create returns a new instance of the PlusOperator.
func (f *PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator is a concrete implementation of the Operator interface.
type MinusOperator struct {
	*OperatorBase
}

// Result returns the result of the minus operation.
func (o *MinusOperator) Result() int {
	return o.a - o.b
}

// MultiplyOperatorFactory is a concrete implementation of the OperatorFactory interface.
type MinusOperatorFactory struct{}

// Create returns a new instance of the MinusOperator.
func (f *MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

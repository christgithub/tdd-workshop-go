package number

import "fmt"

type MarksInteger interface {
	GetNative() int

	IsEqualTo(integer MarksInteger) bool

	IsGreaterThan(integer MarksInteger) bool

	IsLessThan(integer MarksInteger) bool

	Add(integer MarksInteger) MarksInteger

	Subtract(integer MarksInteger) MarksInteger

	Multiply(integer MarksInteger) MarksInteger

	DivideBy(integer MarksInteger) (MarksInteger, error)
}

type ChrisInteger struct {
	value int
}

func NewChrisInt(i int) *ChrisInteger {
	return &ChrisInteger{
		value: i,
	}
}

func (c *ChrisInteger) GetNative() int {
	return c.value
}

func (c *ChrisInteger) IsEqualTo(integer MarksInteger) bool {
	return true
}

func (c *ChrisInteger) IsGreaterThan(integer MarksInteger) bool {
	return c.value > integer.GetNative()
}

func (c *ChrisInteger) IsLessThan(integer MarksInteger) bool {
	return c.value < integer.GetNative()
}

func (c *ChrisInteger) Add(integer MarksInteger) MarksInteger {
	result := c.value + integer.GetNative()
	return NewChrisInt(result)
}

func (c *ChrisInteger) Subtract(integer MarksInteger) MarksInteger {
	result := c.value - integer.GetNative()
	return NewChrisInt(result)
}

func (c *ChrisInteger) Multiply(integer MarksInteger) MarksInteger {
	result := c.value * integer.GetNative()
	return NewChrisInt(result)
}

func (c *ChrisInteger) DivideBy(integer MarksInteger) (MarksInteger, error) {
	if integer.GetNative() == 0 {
		return nil, fmt.Errorf("Can't divide by %d", integer.GetNative())
	}

	if c.value%integer.GetNative() != 0 {
		return nil, fmt.Errorf("Result not of type int", c.value%integer.GetNative())
	}

	result := c.value / integer.GetNative()
	return NewChrisInt(result), nil
}

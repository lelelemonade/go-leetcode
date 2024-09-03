package problem537

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type complex struct{
	real int
	imaginary int
}

func (c *complex)toString()string{
	return strconv.Itoa(c.real)+"+"+strconv.Itoa(c.imaginary)+"i"
}

func (c *complex)Multiply(toMul complex){
	real:=c.real*toMul.real-c.imaginary*toMul.imaginary
	imaginary:=c.real*toMul.imaginary+c.imaginary*toMul.real
	c.real=real
	c.imaginary=imaginary
}

func parseComplex(complexString string)complex{
	s:=strings.Split(complexString,"+")
	real,_:=strconv.Atoi(s[0])
	imaginary,_:=strconv.Atoi(s[1][:len(s[1])-1])
	return complex{real: real,imaginary: imaginary}
}

func complexNumberMultiply(num1 string, num2 string) string {
    complex1:=parseComplex(num1)
    complex2:=parseComplex(num2)
	complex1.Multiply(complex2)
	return complex1.toString()
}

func Test537(t *testing.T) {
	assert.Equal(t,"0+2i",complexNumberMultiply("1+1i","1+1i"))
}
package problem43

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func multiply(num1 string, num2 string) string {
//     result := make([]int, len(num1)+len(num2))

//     for i := len(num1) - 1; i >= 0; i-- {
//         carry := 0
//         for j := len(num2) - 1; j >= 0; j-- {
//             iNum := int(num1[i] - '0')
//             jNum := int(num2[j] - '0')
//             total := iNum*jNum + result[i+j+1] + carry

//             carry = total / 10
//             result[i+j+1] = total % 10
//         }
//         result[i] += carry
//     }

//     // Convert the result slice to a string
//     var strResult []rune
//     for i := 0; i < len(result); i++ {
//         if !(len(strResult) == 0 && result[i] == 0) { // Skip leading zeros
//             strResult = append(strResult, rune(result[i]+'0'))
//         }
//     }

//     if len(strResult) == 0 {
//         return "0"
//     }

//     return string(strResult)
// }

func multiply(num1 string, num2 string) string {
	result:=make([]int,len(num1)+len(num2))
    for i:=len(num1)-1;i>=0;i--{
		carry:=0
		for j:=len(num2)-1;j>=0;j--{
			iNum:=int(num1[i]-'0')
			jNum:=int(num2[j]-'0')
			total:=iNum*jNum+carry+result[i+j+1]
			carry=total/10
			result[i+j+1]=total%10
		}
		result[i]=carry
	}
	resultString:=""
	for i:=0;i<len(result);i++{
		if !(len(resultString)==0&&result[i]==0){
			resultString+=string(rune(result[i]+'0'))
		}
	}
	if resultString == ""{
		return "0"
	}
	return resultString
}

func Test43(t *testing.T) {
	// assert.Equal(t,"6",multiply("2","3"))
	assert.Equal(t,"56088",multiply("123","456"))
}
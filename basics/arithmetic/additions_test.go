package arithmetic

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5,5)
	expectedResult := 10
	if total == expectedResult {
		t.Log("Test Successful")
	}else {
		t.Log("Test Failed..actual result=",total,"expected=",expectedResult)
	}
}

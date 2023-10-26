package compiler 

import {
	    "monkey/code"
		"testing"
}

type compilerTestcase struct {
	input      string 
	expectedConstants   []interface{}
	expectedInstructions[]code.Instructions
}
func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestcase{
		{
			input :    "1+2"
			expectedConstants: []interface{} {1,2},
			expectedInstructions: []code.Instructions{
				   code.Make(code.OpConstant,0),
				   code.Make(code.OpConstant,1),
			},
		},

		runCompilerTests(t, tests)
	}
}
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

func runCompilerTests (t *testing.T, tests[]compilerTestcase){
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		compiler :=New()
		err:= compiler.compile(program)
		if err != nil {
			t.Fatalf("compile error: %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed: %s", err)
		}

		err = testConstants(t,tt.expectedConstants, bytecode.constants)
		if err != nil {
			t.Fatalf("testConstants failed %s", err)
		}
	}
}
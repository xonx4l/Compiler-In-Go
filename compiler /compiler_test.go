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

		func parse (input string ) *ast.Program {
			l := lexer.New(input)
			p := parser.New(l)
			return p.ParserProgram()
		}

		func testInstructions (
			expected []code.Instructions,
			actual code.instructions,
		)error{
			concatted := concatInstruction(expected)

			if len(actual) != len(concatted){
				return fmt.errorf("wrong instructions length . \nwant=%q\ngot =%q", concatted, actual)
			}

			for i, ins:= range concatted {
				if actual [i] != ins {
					return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot =%q",i,concatted,actual)
				}
			}
            return nil
		}
	}
}
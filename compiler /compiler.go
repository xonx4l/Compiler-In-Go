package compiler

import {
	    "monkey/ast"
		"monkey/code"
		"monkey/object"
}

type Compiler struct{
	      instructions code.instructions
		  constants   []object.object
}
func New()  *Compiler{
	    return &Compiler{
			instructions: code.Instructions{},
			constants: []object.object{},
		}
}

func(c *Compiler) compile(node ast.node) error {
	  return nil
}
func (c *compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		constants: c.constants,
	}
}

type Bytecode struct {
	Instructions code.Instructions
	constants    []object.object
}

func (c *compiler) compile(node ast.Node) error {
	switch node =: node.(type){
	case *ast.program:
		for _, s := range node.Statements{
			err:= c.Compile(s)
			 if err != nil {
				return err
			 }
		}
	}
}
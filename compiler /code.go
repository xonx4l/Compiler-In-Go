import (
	"encoding/binary"
	"fmt"
)

func Make(op Opcode, operands ...int) []byte {
	 def, ok := definitions[op]
	 if !ok {
		 return []byte{}
	 }

	 instructionLen := 1
	 for _, w := range def.OperandsWidths {
		instructionLen += w	
	 }

	 instruction := make[]byte, instructionLen)
	 instruction[0] = byte(op)

	 offset :=1
	 for i,o := range operands{
		width := def.OperandsWidth[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16[offset:], uint16(o))
		}
		offset += width
	 }

	 return instruction
}
  func  (ins Instructions string () string {
	  return ""
  }

  func ReadOperands (def *Definition, ins Instructions) ([]int, int ) {
	operands := make ([]int, len(def.OperandsWidth))
	offset :=0

	for i, width := range def.operandWidths {
		switch width {
		case 2:
			operands[i]= int(ReadUint16(ins[offset:]))
		}

		offset +=width
	}

	return operands, offset
  }
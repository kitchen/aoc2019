package day2

import (
	"fmt"
)

func Run(ops []int, pos int) ([]int, error) {
	op := ops[pos]
	if op == 99 {
		return ops, nil
	}

	pos1 := ops[pos+1]
	pos2 := ops[pos+2]
	pos3 := ops[pos+3]

	if pos1 >= len(ops) || pos2 >= len(ops) || pos3 >= len(ops) {
		return nil, fmt.Errorf("bad index")
	}

	switch ops[pos] {
	case 1:
		ops[ops[pos+3]] = ops[ops[pos+1]] + ops[ops[pos+2]]
	case 2:
		ops[ops[pos+3]] = ops[ops[pos+1]] * ops[ops[pos+2]]
	case 99:
		return ops, nil
	default:
		return nil, fmt.Errorf("bad operation")
	}

	pos += 4
	return Run(ops, pos)
}

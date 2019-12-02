package day2

import (
	"fmt"
)

func Run(ops []int, pos int) ([]int, error) {
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

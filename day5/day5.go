package day5

import (
	"fmt"
	"math"
)

func Run(ops []int, pos int, input int) ([]int, int, error) {
	op := ops[pos] % 100

	switch op {
	case 1:
		ops[ops[pos+3]] = Arg(ops, pos, 1) + Arg(ops, pos, 2)
		pos += 4
	case 2:
		ops[ops[pos+3]] = Arg(ops, pos, 1) * Arg(ops, pos, 2)
		pos += 4
	case 3:
		ops[ops[pos+1]] = input
		pos += 2
	case 4:
		input = Arg(ops, pos, 1)
		pos += 2
	case 5:
		if Arg(ops, pos, 1) != 0 {
			pos = Arg(ops, pos, 2)
		} else {
			pos += 3
		}
	case 6:
		if Arg(ops, pos, 1) == 0 {
			pos = Arg(ops, pos, 2)
		} else {
			pos += 3
		}
	case 7:
		if Arg(ops, pos, 1) < Arg(ops, pos, 2) {
			ops[ops[pos+3]] = 1
		} else {
			ops[ops[pos+3]] = 0
		}
		pos += 4
	case 8:
		if Arg(ops, pos, 1) == Arg(ops, pos, 2) {
			ops[ops[pos+3]] = 1
		} else {
			ops[ops[pos+3]] = 0
		}
		pos += 4
	case 99:
		return ops, input, nil
	default:
		return nil, 0, fmt.Errorf("bad operation")
	}

	return Run(ops, pos, input)
}

func Arg(ops []int, pos int, arg int) int {
	flags := ops[pos] / 100
	argflag := (flags % int(math.Pow10(arg))) / int(math.Pow10(arg-1))
	if argflag == 1 {
		return ops[pos+arg]
	}
	return ops[ops[pos+arg]]
}

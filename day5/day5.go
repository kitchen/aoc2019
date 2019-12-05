package day5

import "fmt"

func Run(ops []int, pos int, input int) ([]int, int, error) {
	rawOp := ops[pos]
	op := rawOp % 100
	mode := rawOp / 100
	mode1 := mode % 10
	mode /= 10
	mode2 := mode % 10
	mode /= 10
	mode3 := mode

	switch op {
	case 1:
		var pos1, pos2, pos3 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		if mode2 == 1 {
			pos2 = pos + 2
		} else {
			pos2 = ops[pos+2]
		}

		pos3 = ops[pos+3]

		ops[pos3] = ops[pos1] + ops[pos2]
		pos += 4
	case 2:
		var pos1, pos2, pos3 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		if mode2 == 1 {
			pos2 = pos + 2
		} else {
			pos2 = ops[pos+2]
		}

		pos3 = ops[pos+3]

		ops[pos3] = ops[pos1] * ops[pos2]
		pos += 4
	case 3:
		var pos1 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		ops[pos1] = input
		pos += 2
	case 4:
		var pos1 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		input = ops[pos1]
		pos += 2
	case 5:
		var pos1, pos2 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		if mode2 == 1 {
			pos2 = pos + 2
		} else {
			pos2 = ops[pos+2]
		}

		if ops[pos1] != 0 {
			pos = ops[pos2]
		} else {
			pos += 3
		}
	case 6:
		var pos1, pos2 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		if mode2 == 1 {
			pos2 = pos + 2
		} else {
			pos2 = ops[pos+2]
		}

		if ops[pos1] == 0 {
			pos = ops[pos2]
		} else {
			pos += 3
		}
	case 7:
		var pos1, pos2, pos3 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		if mode2 == 1 {
			pos2 = pos + 2
		} else {
			pos2 = ops[pos+2]
		}

		if mode3 == 1 {
			pos3 = pos + 3
		} else {
			pos3 = ops[pos+3]
		}

		if ops[pos1] < ops[pos2] {
			ops[pos3] = 1
		} else {
			ops[pos3] = 0
		}
		pos += 4
	case 8:
		var pos1, pos2, pos3 int
		if mode1 == 1 {
			pos1 = pos + 1
		} else {
			pos1 = ops[pos+1]
		}

		if mode2 == 1 {
			pos2 = pos + 2
		} else {
			pos2 = ops[pos+2]
		}

		if mode3 == 1 {
			pos3 = pos + 3
		} else {
			pos3 = ops[pos+3]
		}

		if ops[pos1] == ops[pos2] {
			ops[pos3] = 1
		} else {
			ops[pos3] = 0
		}
		pos += 4

	case 99:
		return ops, input, nil
	default:
		return nil, 0, fmt.Errorf("bad operation")
	}

	return Run(ops, pos, input)
}

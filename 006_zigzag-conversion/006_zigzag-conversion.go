package leetcode

func convert(s string, numRows int) string {
	if numRows < 2 || len(s) <= numRows {
		return s
	}
	str := []byte(s)
	zn := 2*numRows - 2
	offsets := make([]int, numRows)
	rem := len(str) % zn
	for i := 0; i < numRows-1; i++ {
		if i == 0 {
			offsets[i+1] = len(str) / zn
			if rem > i {
				offsets[i+1]++
			}
		} else {
			offsets[i+1] = 2 * (len(str) / zn)
			if rem > i {
				offsets[i+1]++
			}
			if rem > zn-i {
				offsets[i+1]++
			}
		}
		offsets[i+1] += offsets[i]
	}
	var pos int
	res := make([]byte, len(str))
	for i, v := range str {
		if i%zn == 0 || i%zn == numRows-1 {
			pos = offsets[i%zn] + i/zn
		} else if i%zn < numRows-1 {
			pos = offsets[i%zn] + 2*(i/zn)
		} else {
			pos = offsets[zn-i%zn] + 2*(i/zn) + 1
		}
		res[pos] = v
	}
	return string(res)
}

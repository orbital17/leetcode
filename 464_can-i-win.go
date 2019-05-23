package leetcode

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	var chosen uint32
	memo = make(map[uint32]bool)
	return _canIWin(chosen, 0, maxChoosableInteger, desiredTotal)
}

var memo map[uint32]bool

func _canIWin(chosen uint32, cur, max, desired int) bool {
	if v, ok := memo[chosen]; ok {
		return v
	}
	var opWins bool
	var mask uint32
	for i := max; i >= 1; i-- {
		mask = 1 << uint(i-1)
		if chosen&mask == 0 {
			chosen |= mask
			if cur+i >= desired {
				chosen &^= mask
				memo[chosen] = true
				return true
			}
			opWins = _canIWin(chosen, cur+i, max, desired)
			if opWins {
				chosen &^= mask
			} else {
				chosen &^= mask
				memo[chosen] = true
				return true
			}
		}
	}
	memo[chosen] = false
	return false
}

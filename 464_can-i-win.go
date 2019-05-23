package leetcode

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	chosen := make([]bool, maxChoosableInteger)
	return _canIWin(chosen, 0, maxChoosableInteger, desiredTotal)
}

func _canIWin(chosen []bool, cur, max, desired int) bool {
	var opWins bool
	for i := max; i >= 1; i-- {
		if !chosen[i-1] {
			chosen[i-1] = true
			if cur+i >= desired {
				chosen[i-1] = false
				return true
			}
			opWins = _canIWin(chosen, cur+i, max, desired)
			if opWins {
				chosen[i-1] = false
			} else {
				chosen[i-1] = false
				return true
			}
		}
	}
	return false
}

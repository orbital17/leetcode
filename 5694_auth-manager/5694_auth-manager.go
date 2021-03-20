package leetcode

type AuthenticationManager struct {
	ttl    int
	tokens map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive,
		map[string]int{},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if currentTime >= this.tokens[tokenId] {
		return
	}
	this.tokens[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var count int
	for _, time := range this.tokens {
		if currentTime < time {
			count++
		}
	}
	return count
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */

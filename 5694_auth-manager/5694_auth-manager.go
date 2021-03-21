package auth_manager

type AuthenticationManager struct {
	ttl int
	hm  map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive,
		map[string]int{},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.hm[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if currentTime >= this.hm[tokenId] {
		return
	}
	this.hm[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var count int
	for _, time := range this.hm {
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

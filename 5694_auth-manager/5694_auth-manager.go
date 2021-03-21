package auth_manager

import (
	"github.com/emirpasic/gods/sets/treeset"
)

type AuthenticationManager struct {
	ttl int
	hm  map[string]int
	tm  *treeset.Set
}

type pair struct {
	time int
	id   string
}

func pairComparator(a, b interface{}) int {
	ak := a.(pair)
	bk := b.(pair)
	switch {
	case ak.time > bk.time:
		return 1
	case ak.time < bk.time:
		return -1
	case ak.id > bk.id:
		return 1
	case ak.id < bk.id:
		return -1
	default:
		return 0
	}
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive,
		map[string]int{},
		treeset.NewWith(pairComparator),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.hm[tokenId] = currentTime + this.ttl
	this.tm.Add(pair{currentTime + this.ttl, tokenId})
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if currentTime >= this.hm[tokenId] {
		return
	}
	this.tm.Remove(pair{this.hm[tokenId], tokenId})
	this.hm[tokenId] = currentTime + this.ttl
	this.tm.Add(pair{this.hm[tokenId], tokenId})
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	it := this.tm.Iterator()
	it.First()
	for !this.tm.Empty() && it.Value().(pair).time <= currentTime {
		v := it.Value().(pair)
		if v.time > currentTime {
			break
		}
		it.Next()
		delete(this.hm, v.id)
		this.tm.Remove(v)
	}
	return this.tm.Size()
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */

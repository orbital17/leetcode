package auth_manager

import "testing"

func TestAuthenticationManager_CountUnexpiredTokens(t *testing.T) {
	var r int
	authenticationManager := Constructor(5)
	authenticationManager.Renew("aaa", 1)             // No token exists with tokenId "aaa" at time 1, so nothing happens.
	authenticationManager.Generate("aaa", 2)          // Generates a new token with tokenId "aaa" at time 2.
	r = authenticationManager.CountUnexpiredTokens(6) // The token with tokenId "aaa" is the only unexpired one at time 6, so return 1.
	if r != 1 {
		t.Errorf("error")
	}
	authenticationManager.Generate("bbb", 7) // Generates a new token with tokenId "bbb" at time 7.
	authenticationManager.Renew("aaa", 8)    // The token with tokenId "aaa" expired at time 7, and 8 >= 7, so at time 8 the renew request is ignored, and nothing happens.
	authenticationManager.Renew("bbb", 10)   // The token with tokenId "bbb" is unexpired at time 10, so the renew request is fulfilled and now the token will expire at time 15.
	r = authenticationManager.CountUnexpiredTokens(15)
	if r != 0 {
		t.Errorf("error")
	}
}

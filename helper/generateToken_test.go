package helper

import "testing"

func TestGenerateToken(t *testing.T) {
	token, _ := GenerateToken("1", "blkcor")
	t.Log(token)
}

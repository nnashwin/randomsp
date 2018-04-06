package randomsp

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestGetRandomInt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	int1 := GetRandomInt(1, 1000000)
	int2 := GetRandomInt(1, 1000000)
	if int1 == int2 {
		t.Error("the GetRandomInt ints should not equal each other")
	}
}

func TestGetRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ss := strings.Split(str, "")

	if GetRandomString(ss) == GetRandomString(ss) {
		t.Error("The Get Rand String function should get different strings from the slice")
	}
}

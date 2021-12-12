package main

type UserInfo struct {
	Username string
	Password string
}

var userInfo = map[string]UserInfo{
	"lemma": {"lemma", "@dm1n"},
}

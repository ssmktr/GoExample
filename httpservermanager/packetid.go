package httpservermanager

import (
	"GoExample/gamedata"
)

// 회원 가입 요청
type req_SignupPacket struct {
	Uid string
	Id string
	Pw string
	Nickname string
	LoginType gamedata.LT_LoginType
}

// 회원 가입 응답
type rsp_SignupPacket struct {
	Error    int
	Uid      string
	Id       string
	Pw       string
	Nickname string
}

// 로그인 요청
type req_LoginPacket struct {
	Id string
	Pw string
	LoginType gamedata.LT_LoginType
}

// 로그인 응답
type rsp_LoginPacket struct {
	Error    int
	Uid      string
	Id       string
	Pw       string
	Nickname string
}

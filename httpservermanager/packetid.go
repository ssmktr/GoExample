package httpservermanager

import (
	"GoExample/gamedata"
)

// 계정 인증 요청
type req_AuthPacket struct {
	Uid       string
	Id        string
	NickName  string
	LoginType gamedata.LT_LoginType
}

// 계정 인증 응답
type rsp_AuthPacket struct {
	Error         gamedata.EC_ErrorCode
	Uid           string
	Id            string
	LoginType     gamedata.LT_LoginType
	Lastlogindate string
}

// 로그인 요청
type req_LoginPacket struct {
	Uid       string
	Id        string
	LoginType gamedata.LT_LoginType
}

// 로그인 응답
type rsp_LoginPacket struct {
	Error         gamedata.EC_ErrorCode
	Uid           string
	Id            string
	LoginType     gamedata.LT_LoginType
	Lastlogindate string
}

// 유저 정보 요청
type req_GetUserInfoPacket struct {
	Uid string
}

// 유저 정보 응답
type rsp_GetUserInfoPacket struct {
	Error    gamedata.EC_ErrorCode
	NickName string
	Energy   int
	Gold     int
	Heart    int
}

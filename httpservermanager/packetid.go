package httpservermanager

import (
	"GoExample/gamedata"
)

// EnglishWord 데이터 요청
type Req_EnglishWordData struct {
	Uid string
	Idx int
}

// EnglishWord 데이터 응답
type Rsp_EnglishWordData struct {
	Error gamedata.EC_ErrorCode
	Datas string
}

// Localization 데이터 요청
type Req_LocalizationData struct {
	Uid string
}

// Localization 데이터 응답
type Rsp_LocalizationData struct {
	Error gamedata.EC_ErrorCode
	Datas string
}

// 계정 인증 요청
type Req_AuthPacket struct {
	Uid       string
	Id        string
	NickName  string
	LoginType gamedata.LT_LoginType
}

// 계정 인증 응답
type Rsp_AuthPacket struct {
	Error         gamedata.EC_ErrorCode
	Uid           string
	Id            string
	LoginType     gamedata.LT_LoginType
	Lastlogindate string
}

// 로그인 요청
type Req_LoginPacket struct {
	Uid       string
	Id        string
	LoginType gamedata.LT_LoginType
}

// 로그인 응답
type Rsp_LoginPacket struct {
	Error         gamedata.EC_ErrorCode
	Uid           string
	Id            string
	LoginType     gamedata.LT_LoginType
	Lastlogindate string
}

// 유저 정보 요청
type Req_GetUserInfoPacket struct {
	Uid string
}

// 유저 정보 응답
type Rsp_GetUserInfoPacket struct {
	Error        gamedata.EC_ErrorCode
	NickName     string
	AlphabatType int8
	Energy       int
	Gold         int
	Heart        int
}

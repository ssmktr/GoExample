package gamedata

type RTT_ReadTableType int

const (
	RTT_None RTT_ReadTableType = iota
	RTT_ReadStart
	RTT_ReadFinish
)

type LT_LoginType int

const (
	LT_Guest    LT_LoginType = 1
	LT_Android  LT_LoginType = 2
	LT_FaceBook LT_LoginType = 3
)

type EC_ErrorCode int

const (
	EC_Success           = 0 // 성공
	EC_UnknownError      = 1 // 알수 없는 에러
	EC_TimeOut           = 2 // 타임 아웃
	EC_AlreadyAccount    = 3 // 이미 있는 계정
	EC_MysqlConnectFail  = 4 // MySQL 연결 실패
	EC_NotFoundAccount   = 5 // 계정을 찾을수 없습니다
	EC_NotFoundTableInfo = 6 // 테이블에서 정보를 찾을수 없습니다
	EC_Table_Insert      = 7 // 테이블에 Insert 에러
	EC_Table_Update      = 8 // 테이블에 Update 에러
)


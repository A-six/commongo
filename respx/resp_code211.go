package respx

import "encoding/json"

// errcode 211_00~99
const (
	ERR211_00     errcode = "21100"  //物联网关错误
	ERR211_00_MSG errmsg  = "物联网关错误" //物联网关错误

	ERR211_01     errcode = "21101" //编号不存在
	ERR211_01_MSG errmsg  = "编号不存在" //编号不存在

	ERR211_02     errcode = "21102" //编号已激活
	ERR211_02_MSG errmsg  = "编号已激活" //编号已激活

	ERR211_03     errcode = "21103" //编号已停用
	ERR211_03_MSG errmsg  = "编号已停用" //编号已停用
)

func Resp21100() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR211_00, Errmsg: ERR211_00_MSG})
	return string(respjson)
}

func Resp21101() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR211_01, Errmsg: ERR211_01_MSG})
	return string(respjson)
}

func Resp21102() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR211_02, Errmsg: ERR211_02_MSG})
	return string(respjson)
}

func Resp21103() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR211_03, Errmsg: ERR211_03_MSG})
	return string(respjson)
}

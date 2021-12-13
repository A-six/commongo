package respx

import "encoding/json"

// errcode 100_00~99
const (
	ERR100_00 errcode = "10000" //弹出提醒用户

	ERR100_01     errcode = "10001" //数据不存在
	ERR100_01_MSG errmsg  = "数据不存在" //数据不存在

	ERR100_02 errcode = "10002" //提示用户，不是弹出，前端自定义提示
)

func Resp10000(msg string) string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR100_00, Errmsg: Errmsg(msg)})
	return string(respjson)
}

func Resp10001() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR100_01, Errmsg: ERR100_01_MSG})
	return string(respjson)
}

func Resp10002(msg string) string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR100_02, Errmsg: Errmsg(msg)})
	return string(respjson)
}

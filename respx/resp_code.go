package respx

import (
	"encoding/json"
	"fmt"
)

type errcode string
type errmsg string

const (
	SUCCESS     errcode = "00000" //成功
	SUCCESS_MSG errmsg  = "成功"    //成功
)

func Errmsg(msg string) errmsg {
	return errmsg(msg)
}

func SErrmsg(f string, msgs ...interface{}) errmsg {
	return errmsg(fmt.Sprintf(f, msgs...))
}

func (p errmsg) Str() string {
	return fmt.Sprint(p)
}

func (p errcode) Str() string {
	return fmt.Sprint(p)
}

func RespSuccessByte() []byte {
	respjson, _ := json.Marshal(RespCode{Errcode: SUCCESS, Errmsg: SUCCESS_MSG})
	return respjson
}

func RespSuccess() string {
	respjson, _ := json.Marshal(RespCode{Errcode: SUCCESS, Errmsg: SUCCESS_MSG})
	return string(respjson)
}

func RespSuccessMsgByte(msg string) string {
	respjson, _ := json.Marshal(RespCode{Errcode: SUCCESS, Errmsg: Errmsg(msg)})
	return string(respjson)
}

func RespSuccessMsg(msg string) string {
	respjson, _ := json.Marshal(RespCode{Errcode: SUCCESS, Errmsg: Errmsg(msg)})
	return string(respjson)
}

func RespData(data interface{}) string {
	respjson, _ := json.Marshal(RespCode{Errcode: SUCCESS, Data: data})
	return string(respjson)
}

func RespDatas(datas interface{}) string {
	respjson, _ := json.Marshal(RespCode{Errcode: SUCCESS, Datas: datas})
	return string(respjson)
}

func RespPageDatas(datas, page interface{}) string {
	respjson, _ := json.Marshal(RespCode{Errcode: SUCCESS, Datas: datas, Page: page})
	return string(respjson)
}

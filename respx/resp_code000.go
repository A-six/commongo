package respx

import "encoding/json"

// errcode 000_00~99
const (
	ERR000_01     errcode = "00001"       //通用错误
	ERR000_01_MSG errmsg  = "网络错误，请稍后重试！" //通用错误

	ERR000_02     errcode = "00002" //无操作权限
	ERR000_02_MSG errmsg  = "无操作权限" //无操作权限

	ERR000_03     errcode = "00003"   //请求格式不正确
	ERR000_03_MSG errmsg  = "请求格式不正确" //请求格式不正确

	ERR000_04     errcode = "00004" //无访问权限
	ERR000_04_MSG errmsg  = "无访问权限" //无访问权限

	ERR000_05     errcode = "00005"  //请登录后操作
	ERR000_05_MSG errmsg  = "请登录后操作" //请登录后操作

	ERR000_06     errcode = "00006"   //请求类型未定义
	ERR000_06_MSG errmsg  = "请求类型未定义" //请求类型未定义

	ERR000_07     errcode = "00007" //令牌过期
	ERR000_07_MSG errmsg  = "令牌过期"  //令牌过期

	ERR000_08     errcode = "00008"    //帐号已在别处登陆
	ERR000_08_MSG errmsg  = "帐号已在别处登陆" //帐号已在别处登陆

	ERR000_09     errcode = "00009"      //权限被修改请重新登录
	ERR000_09_MSG errmsg  = "权限被修改请重新登录" //权限被修改请重新登录
)

//00001
func Resp00001Byte() []byte {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_01, Errmsg: ERR000_01_MSG})
	return respjson
}

func Resp00001() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_01, Errmsg: ERR000_01_MSG})
	return string(respjson)
}

func Resp00001MsgByte(msg string) string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_01, Errmsg: Errmsg(msg)})
	return string(respjson)
}

func Resp00001Msg(msg string) string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_01, Errmsg: Errmsg(msg)})
	return string(respjson)
}

//00002
func Resp00002() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_02, Errmsg: ERR000_02_MSG})
	return string(respjson)
}

//00003
func Resp00003() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_03, Errmsg: ERR000_03_MSG})
	return string(respjson)
}

//00004
func Resp00004Byte() []byte {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_04, Errmsg: ERR000_04_MSG})
	return respjson
}

func Resp00004() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_04, Errmsg: ERR000_04_MSG})
	return string(respjson)
}

//00005
func Resp00005Byte(msgs ...string) []byte {
	p := RespCode{Errcode: ERR000_05, Errmsg: ERR000_05_MSG}
	if len(msgs) > 0 {
		pErrmsg := ""
		for _, v := range msgs {
			if len(v) == 0 {
				continue
			}
			pErrmsg += v + ","
		}
		pErrmsg += ERR000_05_MSG.Str()
		p.Errmsg = Errmsg(pErrmsg)
	}
	respjson, _ := json.Marshal(p)
	return respjson
}

func Resp00005() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_05, Errmsg: ERR000_05_MSG})
	return string(respjson)
}

//00006
func Resp00006() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_06, Errmsg: ERR000_06_MSG})
	return string(respjson)
}

//00007
func Resp00007Byte() []byte {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_07, Errmsg: ERR000_07_MSG})
	return respjson
}

func Resp00007() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_07, Errmsg: ERR000_07_MSG})
	return string(respjson)
}

//00008
func Resp00008() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_08, Errmsg: ERR000_08_MSG})
	return string(respjson)
}

//00009
func Resp00009() string {
	respjson, _ := json.Marshal(RespCode{Errcode: ERR000_09, Errmsg: ERR000_09_MSG})
	return string(respjson)
}

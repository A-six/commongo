package respx

import (
	"encoding/json"
)

type RespCode struct {
	Errcode errcode `json:"errcode"`          //错误代码
	Errmsg  errmsg  `json:"errmsg,omitempty"` //错误信息

	Page  interface{} `json:"page,omitempty"`  //分页参数
	Data  interface{} `json:"data,omitempty"`  //数据内容
	Datas interface{} `json:"datas,omitempty"` //数据数组
}

func (p RespCode) Resp() string {
	respjson, _ := json.Marshal(p)
	return string(respjson)
}

func Resp(p *RespCode) string {
	respjson, _ := json.Marshal(p)
	return string(respjson)
}

func RespByte(p *RespCode) []byte {
	respjson, _ := json.Marshal(p)
	return respjson
}

package respx

import (
	"fmt"
	"testing"
)

func TestResp(t *testing.T) {
	fmt.Println(Resp(&RespCode{Errcode: SUCCESS, Data: RespCode{Errcode: SUCCESS}}))
	fmt.Println(Resp(&RespCode{Errcode: SUCCESS, Datas: []RespCode{RespCode{Errcode: SUCCESS}}}))
	fmt.Println(Resp(&RespCode{Errcode: SUCCESS, Page: RespCode{Errcode: SUCCESS}}))

	fmt.Println(Resp(&RespCode{Errcode: ERR000_01}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_02}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_03}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_04}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_05}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_06}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_07}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_08}))
	fmt.Println(Resp(&RespCode{Errcode: ERR000_09}))

	fmt.Println(Resp(&RespCode{Errcode: ERR100_00}))
	fmt.Println(Resp(&RespCode{Errcode: ERR100_01}))
	fmt.Println(Resp(&RespCode{Errcode: ERR100_02}))

	fmt.Println(Resp(&RespCode{Errcode: ERR211_00}))
	fmt.Println(Resp(&RespCode{Errcode: ERR211_01}))
	fmt.Println(Resp(&RespCode{Errcode: ERR211_02}))
	fmt.Println(Resp(&RespCode{Errcode: ERR211_03}))

}

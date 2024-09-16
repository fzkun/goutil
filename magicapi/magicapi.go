package magicapi

import (
	"errors"
	"fmt"
	"github.com/fzkun/goutil/jsonutil"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"reflect"
)

type MagicSdk struct {
	ctx *Context //上下文
}

func NewMagicSdk(ctx *Context) *MagicSdk {
	return &MagicSdk{ctx: ctx}
}

// NewMagicRequest 创建magic请求
func (s *MagicSdk) NewMagicRequest(uri string, body any, obj any) (err error) {
	var (
		httpResp *resty.Response
		respJson string
	)
	req := resty.New().R()
	kind := reflect.TypeOf(body).Kind()
	if kind == reflect.Map ||
		kind == reflect.String {
		req.SetBody(body)
	} else {
		var bodyMap map[string]interface{}
		if bodyMap, err = jsonutil.StructToMap(body); err != nil {
			return
		}
		req.SetBody(bodyMap)
	}
	if httpResp, err = req.Post(s.ctx.Config.MagicUrl + uri); err != nil {
		return
	}
	respJson = httpResp.String()
	if gjson.Get(respJson, "code").Int() != 0 {
		err = errors.New(gjson.Get(respJson, "message").String())
		return
	}
	if err = jsonutil.JsonStrToStruct(gjson.Get(respJson, "data").String(), obj); err != nil {
		err = errors.New(fmt.Sprintf("magic解析参数失败,err=%s,url=%s,输入内容=%s,返回内容=%s", err.Error(), uri, jsonutil.StructToJsonString(body), respJson))
		return
	}
	return
}

// Package src /**
package mian

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Express struct {
	Customer string
	Key      string
}

func NewExpress(customer, key string) (express *Express, err error) {
	if customer == NULL {
		return nil, newErr("授权码不能为空")
	}

	if key == NULL {
		return nil, newErr("key不能为空")
	}

	express = &Express{
		Customer: customer,
		Key:      key,
	}
	return
}

// SyncQuery  实时快递单号查询
func (e *Express) SyncQuery(p Param) (resp QueryResp, err error) {
	if p == nil {
		return
	}

	if num := p.Get("num"); num == NULL {
		return resp, newErr("快递单号不能为空")
	}

	com := p.Get("com")
	if com == NULL {
		return resp, newErr("快递公司不能为空")
	}

	if com == ComSF {
		if phone := p.Get("phone"); phone == NULL {
			return resp, newErr("顺丰快递手机号不能为空")
		}
	}

	sign := e.Sign(p)
	if sign == NULL {
		return resp, newErr("签名失败")
	}

	reqParam := e.buildReq(sign, p)
	res, err := NewClient().Post(methodSyncQuery, reqParam)
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal(res, &resp); err != nil {
		fmt.Println(err)
		return resp, newErr("数据解析失败")
	}

	if resp.BaseResp.ReturnCode != NULL {
		return resp, newErr(resp.BaseResp.Message)
	}

	return
}

func (e *Express) buildReq(sign string, p Param) url.Values {
	b, _ := json.Marshal(p)
	return url.Values{"customer": {e.Customer}, "sign": {sign}, "param": {string(b)}}

}

// Package kd100   /**
package kd100

import (
	"encoding/json"
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

// SyncQuery  实时快递单号查询 https://api.kuaidi100.com/document/shishichaxunchanpinjieshao
func (e *Express) SyncQuery(p Param) (resp QueryResp, err error) {
	// 参数校验
	if err = p.queryVerify(); err != nil {
		return resp, err
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
		return resp, newErr("数据解析失败")
	}

	if resp.BaseResp.ReturnCode != NULL {
		return resp, newErr(resp.BaseResp.Message)
	}

	return
}

// MapTrack 快递查询地图轨迹 https://api.kuaidi100.com/document/5ff2c3e7ba1bf00302f5612e
func (e *Express) MapTrack(p Param) (resp MapTrackResp, err error) {
	// 参数校验
	if err = p.queryVerify(); err != nil {
		return resp, err
	}

	if from := p.Get("from"); from == NULL {
		return resp, newErr("出发地信息不能为空")
	}

	if to := p.Get("to"); to == NULL {
		return resp, newErr("目的地信息不能为空")
	}

	// 签名
	sign := e.Sign(p)
	if sign == NULL {
		return resp, newErr("签名失败")
	}

	reqParam := e.buildReq(sign, p)
	res, err := NewClient().Post(methodMapTrack, reqParam)
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal(res, &resp); err != nil {
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

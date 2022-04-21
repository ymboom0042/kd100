// Package kd100   /**
package kd100

type Param map[string]string

// Set 设置参数
func (p Param) Set(key string, value string) Param {
	p[key] = value
	return p
}

// Get  获取参数
func (p Param) Get(key string) string {
	if p == nil {
		return NULL
	}
	v, ok := p[key]
	if !ok {
		return NULL
	}

	return v
}

// queryVerify 查询参数校验
func (p Param) queryVerify() error {
	if p == nil {
		return newErr("参数错误")
	}
	if num := p.Get("num"); num == NULL {
		return newErr("快递单号不能为空")
	}

	com := p.Get("com")
	if com == NULL {
		return newErr("快递公司不能为空")
	}

	if com == ComSF {
		if phone := p.Get("phone"); phone == NULL {
			return newErr("顺丰快递手机号不能为空")
		}
	}

	return nil
}

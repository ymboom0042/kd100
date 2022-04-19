// Package src /**
package src

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

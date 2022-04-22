// Package kd100   /**
package kd100

const (
	NULL    = ""
	gateway = "https://poll.kuaidi100.com"
)

const (
	// 实时快递查询
	methodSyncQuery = "poll/query.do"
	// 快递查询地图轨迹
	methodMapTrack = "poll/maptrack.do"
)

const (
	ComYT = "yuantong"
	ComYD = "yunda"
	ComZT = "zhongtong"
	ComST = "shentong"
	ComYZ = "youzhengguonei"
	ComSF = "shunfeng"
	ComJT = "jtexpress"
)

var ComValue = map[string]string{
	ComYT: "圆通快递",
	ComYD: "韵达快递",
	ComZT: "中通快递",
	ComST: "申通快递",
	ComYZ: "邮政快递",
	ComSF: "顺丰快递",
	ComJT: "极兔速递",
}

const (
	StateOnPassage     = "0"
	StateOnPassage1001 = "1001"
	StateOnPassage1002 = "1002"
	StateOnPassage1003 = "1003"

	StateCollect    = "1"
	StateCollect101 = "101"
	StateCollect102 = "102"
	StateCollect103 = "103"

	StateKnotty    = "2"
	StateKnotty201 = "201"
	StateKnotty202 = "202"
	StateKnotty203 = "203"
	StateKnotty204 = "204"
	StateKnotty205 = "205"
	StateKnotty206 = "206"
	StateKnotty207 = "207"
	StateKnotty208 = "208"
	StateKnotty209 = "209"

	StateChecked    = "3"
	StateChecked301 = "301"
	StateChecked302 = "302"
	StateChecked303 = "303"
	StateChecked304 = "304"

	StateChargeBack    = "4"
	StateChargeBack401 = "401"

	StateDeliver    = "5"
	StateDeliver501 = "501"

	StateCustomsClearance   = "8"
	StateCustomsClearance10 = "10"
	StateCustomsClearance11 = "11"
	StateCustomsClearance12 = "12"
	StateCustomsClearance13 = "13"

	StateRejection = "14"
)

var StateValue = map[string]string{
	StateOnPassage:     "在途",
	StateOnPassage1001: "到达派件城市",
	StateOnPassage1002: "干线",
	StateOnPassage1003: "转递",

	StateCollect:    "揽收",
	StateCollect101: "已下单",
	StateCollect102: "待揽收",
	StateCollect103: "已揽收",

	StateKnotty:    "疑难",
	StateKnotty201: "超时未签收",
	StateKnotty202: "超时未更新",
	StateKnotty203: "拒收",
	StateKnotty204: "派件异常",
	StateKnotty205: "柜或驿站超时未取",
	StateKnotty206: "无法联系",
	StateKnotty207: "超区",
	StateKnotty208: "滞留",
	StateKnotty209: "破损",

	StateChecked:    "签收",
	StateChecked301: "本人签收",
	StateChecked302: "派件异常后签收",
	StateChecked303: "代签",
	StateChecked304: "投柜或驿站签收",

	StateChargeBack:    "退签",
	StateChargeBack401: "已销单",

	StateDeliver:    "派件",
	StateDeliver501: "投柜或驿站",

	StateCustomsClearance:   "清关",
	StateCustomsClearance10: "待清关",
	StateCustomsClearance11: "清关中",
	StateCustomsClearance12: "已清关",
	StateCustomsClearance13: "清关异常",

	StateRejection: "拒签",
}

package model

// user profile
type Profile struct {
	Name string // 用户名
	//Homepage      string // 个人主页地址（一般为用户详情页面） // this field is just Item's Url
	Age           int    // 年龄
	Gender        string // 性别
	Height        int    // 身高: cm
	Weight        int    // 体重: kg
	Salary        string // 收入
	Marriage      string // 婚姻状况
	Education     string // 学历
	Occupation    string // 职业
	NativePlace   string // 籍贯
	Workplace     string // 工作地
	Constellation string // 星座
	Zodiac        string // 生肖
	House         string // 购房
	Car           string // 购车
	Pic           string // 个人生活照片
}

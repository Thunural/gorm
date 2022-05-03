package model
import "time"

// CommonRes 通用api响应
type CommonRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限  -1  失败
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	Btype BunissType  `json:"otype"` //业务类型
}

type TimeStruct struct {
	CreateTime time.Time `gorm:"column:createTime;type:datetime(6);index:IDX_b9af0e100be034924b270aab31;" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime;type:datetime(6);index:IDX_8857d8d43d38bebd7159af1fa6;" json:"updateTime"`
}

// CaptchaRes 验证码响应
type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}

// TableDataInfo 通用分页表格响应
type TableDataInfo struct {
	Total int         `json:"total"` //总数
	Rows  interface{} `json:"rows"`  //数据
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
}

// Ztree 通用的树形结构
type Ztree struct {
	Id      int64  `json:"id"`      //节点ID
	Pid     int64  `json:"pId"`     //节点父ID
	Name    string `json:"name"`    //节点名称
	Title   string `json:"title"`   //节点标题
	Checked bool   `json:"checked"` //是否勾选
	Open    bool   `json:"open"`    //是否展开
	Nocheck bool   `json:"nocheck"` //是否能勾选
}

// RemoveReq 通用的删除请求
type RemoveReq struct {
	Ids []int `form:"ids"  binding:"required"`
}

// PageReq 通用页码请求参数
type PageReq struct {
	Page int `p:"page"` //当前页码
	Size int `p:"size"` //每页数
}

// SelectPageReq 通用分页请求
type SelectPageReq struct {
	Keyword       string `form:"keyWord" p:"keyWord"` // 搜索
	DepartmentIds []int  `form:"departmentIds" p:"departmentIds"`
	Sort          string `p:"sort"`
	Order         string `p:"order"`
	ClassifyId    int    `p:"classifyId"` //分类ID
	// 继承通用分页
	PageReq
}

// DetailReq 通用详情请求
type DetailReq struct {
	Id int64 `json:"id" binding:"required"` //主键ID
}

// EditReq 通用修改请求
type EditReq struct {
	Id int64 `json:"id"` //主键ID
}

// LoginReq 登陆请求结构体
type LoginReq struct {
	VerifyCode string `form:"verifyCode" json:"verifyCode" binding:"required"`
	UserName   string `form:"username"   json:"username"   binding:"required"`
	Password   string `form:"password"   json:"password"   binding:"required"`
	CaptchaId  string `form:"captchaId" json:"captchaId" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

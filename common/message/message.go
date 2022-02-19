package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterResMesType = "RegisterResMes"
	RegisterMesType = "RegisterMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	Useroffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //数据
}

type LoginMes struct {
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //500未注册  200登录成功
	UserIds []string `json:"userids"`  //保存用户id
	Error string `json:"error"` //错误信息
}

//用户结构体
type User struct {
	//确定字段信息

	//为了序列化和反序列化成功，我们必须保证用户信息的json字符串的key和结构体字段对应的tag名字一致
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"`
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code int `json:"code"` // 返回状态码 400 表示该用户已经占用 200注册成功
	Error string `json:"error"` //返回错误信息
}

// 为了配合服务器端推送用户状态变化的消息

type NotifyUserStatusMes struct {
	UserId string `json:userId` //用户id
	Status int `json:status` //用户状态
}


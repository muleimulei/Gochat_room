package process

var (
	userMgr *UserMgr
)

type UserMgr struct{
	onlineUsers map[int] *UserProcessor
}

//完成对UserMgr初始化工作

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int] *UserProcessor),
	}
}

//完成对onlineUser添加

func (um *UserMgr) AddOnlineUser(up *UserProcessor) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (um *UserMgr) DeOnlineUser(userid string){
	delete(this.onlineUsers, userid)
}

// 返回当前所有在线用户
func (um *UserMgr) GetAllOnlineUser() map[int] *UserProcessor {
	return this.onlineUsers
}

//根据id返回对应的值
func (um *UserMgr) GetOnlineUserById (userId string) (up *UserProcessor, err error){
	// 如何从 map取出一个值， 带检测模式
	up, ok := this.onlineUsers[userId]
	if ok == false {
		err = fmt.Errorf("用户 %d 不存在", userId)
	}
	return
}


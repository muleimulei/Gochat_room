package model

import (
	"encoding/json"
	"fmt"
	"message"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	Pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) *UserDao {
	return &UserDao{
		Pool: pool,
	}
}

//思考一下在UserDao 应该提供哪些方法

//1. 根据用户id 返回一个User实例+err

func (ud *UserDao) getUserById(conn redis.Conn, id string) (user User, err error) {

	//通过给定id去查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXIST
		}
		return
	}

	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	}
	return
}

/*
完成登录的效验
1. login 完成对用户的效验
2. 如果用户的id和pwd都正确，则返回一个user实例
3. 如果用户的id和pwd有错误，则返回对应的错误信息
*/

func (ud *UserDao) Login(userId string, userPwd string) (user User, err error) {
	//先从UserDao 的连接池中取出一根连接
	conn := ud.Pool.Get()
	defer conn.Close()

	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 比较密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
	}
	return
}

func (ud *UserDao) Register(user *message.User) (err error) {
	//先从UserDao 的连接池中取出一根连接
	conn := ud.Pool.Get()
	defer conn.Close()

	_, err = ud.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXIST
		return err
	}else if err == ERROR_USER_NOTEXIST {
		//这时，说明id在redis还没有，可以注册
		data, err := json.Marshal(*user)
		_, err = conn.Do("HSet", "users", user.UserId, string(data))
		if err != nil {
			fmt.Println("保存注册用户错误 err = ", err)
			return err
		}
	}
	return nil
}

package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go_code/chatroom/common/message"
)

//定义一个UserDao结构体
//完成对User结构体的各种操作

//在服务器启动后就初始化一个UserDao的实例
//做成全局变量，在需要和redis操作时，直接使用即可

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao的实例

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//1.根据用户id 返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))

	if err != nil {
		if err == redis.ErrNil {
			//表示在users 哈希中，没有找到对应id
			err = ERROR_USER_NOTEISTS
		}
		return
	}

	user = &User{}

	//这里我们需要把res反序列化成User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

//完成登录校验
//1.Login  完成对用户的验证
//2.如果用户的id和pwd都正确，则返回一个user实例
//3.如果用户的id或pwd有错误，则返回相应的错误信息

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao连接池中取出一个连接

	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//证明用户已经获取到
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao连接池中取出一个连接

	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//说明该用户还没有注册，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 ,err=", err)
	}
	return
}

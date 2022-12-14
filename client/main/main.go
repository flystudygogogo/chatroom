package main

import (
	"fmt"
	"go_code/chatroom/client/process"
	"os"
)

//定义两个全局变量，一个表示用户的id，一个表示用户密码
var userId int
var userPwd string
var UserName string

func main() {
	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	//var loop = true

	for true {
		fmt.Println("---------------欢迎登录多人聊天系统---------------")
		fmt.Println("\t\t1 登录聊天室")
		fmt.Println("\t\t2 注册用户")
		fmt.Println("\t\t3 退出系统")
		fmt.Println("请选择(1-3):")
		fmt.Println("----------------")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			//1，创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的名字(nickname):")
			fmt.Scanf("%s\n", &UserName)

			//2.调用UserProcess实例完成注册的请求
			up := &process.UserProcess{}
			_ = up.Register(userId, userPwd, UserName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)

		default:
			fmt.Println("你的输入有误，请重新输入：")
		}
	}

	//更多用户的输入，显示新的提示信息
	//if key == 1 {
	//	fmt.Println("请输入用户的id")
	//	fmt.Scan("%d\n", &userId)
	//	fmt.Println("请输入用户的密码")
	//	fmt.Scan("%s\n", &userPwd)
	//
	//	//因为使用了新的结构，因此我们创建一个
	//	//先把登录的函数，写到另一个文件
	//	//这里重新调用
	//	//login(userId, userPwd)
	//	//if err != nil {
	//	//	fmt.Println("登录失败")
	//	//} else {
	//	//	fmt.Println("登录成功")
	//	//}
	//
	//} else {
	//	if key == 2 {
	//		fmt.Println("进行用户注册的逻辑")
	//	}
	//}

}

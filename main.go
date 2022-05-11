package main

import (
	"fmt"
	"http-server/pkg"
)

func home(ctx *pkg.Context) {
	fmt.Fprintf(ctx.W, "这是主页")
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int    `json:"biz_code"`
	Msg     string `json:"msg"`
	//Data    interface{} `json:"data"`
}

func SignUp(ctx *pkg.Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		// 上述处理失败了
		err = ctx.BadRequestJson(nil)
		return
	}
	// 返回一个虚拟的 user id 表示注册成功了
	res := &commonResponse{
		BizCode: 934655044,
		Msg:     "success",
	}
	err = ctx.OKJson(res)
	if err != nil {
		fmt.Printf("写入相应失败：%v", err)
	}
}

func main() {
	server1 := pkg.NewServer("server-1")
	server1.Route("/", home)
	server1.Route("/sign", SignUp)
	server1.Start(":8080")
	/*
		http.HandleFunc("/", home)
		http.ListenAndServe(":8080", nil)
	*/
}

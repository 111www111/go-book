package util

import (
	"books/model"
)

func GetTrueDataResult(data interface{}) model.Result {
	return model.Result{
		Status: true,
		Msg:    "操作成功",
		Data:   data,
	}
}

func GetTrueMsgResult(msg string) model.Result {
	return model.Result{
		Status: true,
		Msg:    msg,
		Data:   nil,
	}
}

func GetFalseMsgResult(msg string) model.Result {
	return model.Result{
		Status: false,
		Msg:    msg,
		Data:   nil,
	}
}

/*
GetErrorResult
这个不是让你返回业务错误的
这个是捕获全局异常捕获的！！
*/
func GetErrorResult() model.Result {
	return model.Result{
		Status: false,
		Msg:    "服务异常",
		Data:   nil,
	}
}

/*
GetErrorMsgResult
这个不是让你全局异常捕获的！
这个是返回业务错误的！
*/
func GetErrorMsgResult(msg string) model.Result {
	return model.Result{
		Status: false,
		Msg:    msg,
		Data:   nil,
	}
}

/*
GetCheckErrorResult
参数为空的异常
*/
func GetCheckErrorResult() model.Result {
	return model.Result{
		Status: false,
		Msg:    "关键参数为空",
		Data:   nil,
	}
}

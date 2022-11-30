package tasks

import "errors"

var ErrStopTask error = errors.New("终止定时任务") // 停止定时任务的错误  方法抛出此异常  终止定时任务继续执行

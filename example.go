package main

import (
	"fmt"

	"github.com/solstice-gao/chatglm-go/chat"
)

func main() {
	// 获取方式见下方常见问题
	authorization := "Bearer eyxxxx"
	cookie := "acw_tc=xxx;"
	prompt := "请写一篇2000字的申请书"

	// 创建 ChatService 实例
	chatService := chat.NewChatService(authorization, cookie)

	// 获取任务ID
	task := chatService.GetTaskId(prompt)
	fmt.Println("task id-------")
	fmt.Println(task.Result.TaskID)
	fmt.Println("-------")

	// 获取上下文ID
	context := chatService.GetContextId(prompt, task.Result.TaskID)

	fmt.Println("context id-------")
	fmt.Println(context.Result.ContextID)
	fmt.Println("-------")

	// 调用获取聊天数据的方法
	ans, err := chatService.GetChat(context.Result.ContextID)

	if err != nil {
		fmt.Println(err)
		return
	}

	//var lastLine string
	fmt.Println(ans)
	r := []rune(ans)
	fmt.Println(len(r))

	//// 处理聊天数据
	//// TODO: 根据自己的需求进行处理
	//
	//// 输出聊天数据
	//fmt.Println(chatData)
}

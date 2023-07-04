# chatglm-go

ChatGLM Go SDK

## 项目简介

ChatGLM Go SDK 是一个将网页版 ChatGLM 转换为 Go 语言的软件开发工具包。它提供了一组方便易用的函数和方法，使开发人员能够在 Go 项目中快速集成 ChatGLM 的功能。

## 特性

- 与 ChatGLM 网页版一致的功能和交互体验
- 支持从 ChatGLM 后端 API 获取实时聊天数据
- 提供简洁的 API 接口，方便在 Go 项目中使用
- 支持自定义输出方式，以满足不同的需求

## 安装

使用 Go 的包管理工具 `go get` 可以快速获取并安装 ChatGLM Go SDK：

```
go get github.com/solstice-gao/chatglm-go
```

## 使用示例

下面是一个简单的示例代码，展示了如何在 Go 项目中使用 ChatGLM Go SDK：

```go
package main

import (
	"fmt"
	"github.com/solstice-gao/chatglm-go/chat"
)

func main() {
  // 获取方式见下方常见问题
	authorization: "your-authorization-tokrn"
	cookie: "your-cookie"
	// 创建 ChatService 实例
	chatService := chatglm.NewChatService(authorization, cookie)

	// 获取任务ID
	task := chatservice.GetTaskId(prompt)

	// 获取上下文ID
	context := chatservice.GetContextId(prompt, task.Result.TaskID)

	// 调用获取聊天数据的方法
	sacnner,err := chatservice.GetChatStream(context.Result.ContextID, true)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	var lastLine string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// 处理聊天数据
	// TODO: 根据自己的需求进行处理

	// 输出聊天数据
	fmt.Println(chatData)
}
```

## 贡献

欢迎对 ChatGLM Go SDK 进行贡献！如果你发现了 bug，或者有改进建议，请提出 issue 或提交 pull request。在参与贡献之前，请阅读项目的贡献指南（CONTRIBUTING.md）。

## 许可证

ChatGLM Go SDK 使用 MIT 许可证。详细信息请参阅许可证文件（LICENSE）。

## 联系方式

如果你有任何问题或建议，可以通过以下方式联系我：

- 邮箱：17604515707@163.com
- GitHub 项目主页：https://github.com/solstice-gao

请随时提出问题或反馈，我会尽快回复。

## 常见问题

在这里列出一些常见问题和解答，以帮助其他用户更好地使用 ChatGLM Go SDK。

1. **怎么获取authorization/cookie？**
   - 首先需要有 https://chatglm.cn 的账号，目前还在内测中，不过很好申请，申请后很快就能通过
   - 登录账号，浏览器F12/开发者模式，选择Network
   - 输入问题进行提问
   - <img width="1432" alt="image" src="https://github.com/solstice-gao/chatglm-go/assets/71240666/8f7d86f4-ef06-48b0-9ed6-bf7955f537ed">
   - <img width="1432" alt="image" src="https://github.com/solstice-gao/chatglm-go/assets/71240666/33653ed8-8bca-40a8-b20c-169255ddde4a">

## 致谢

特别感谢以下人员对 ChatGLM Go SDK 的贡献和支持：

- Gaoxu

感谢你们的支持！

## 更新

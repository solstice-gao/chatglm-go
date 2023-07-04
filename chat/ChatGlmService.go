package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/solstice-gao/chatglm-go/entity"
)

type ChatService struct {
	// ... 可能的依赖或其他属性
	authorization string
	cookie        string
}

func NewChatService(authorization string, cookie string) *ChatService {
	return &ChatService{
		// ... 初始化依赖或其他属性
		authorization: authorization,
		cookie:        cookie,
	}
}

func (s *ChatService) GetChatStream(contextId string) (bufio.Scanner, error) {
	scanner := bufio.NewScanner(nil)
	url := "https://chatglm.cn/chatglm/backend-api/v1/stream?context_id=" + contextId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return *scanner, err
	}
	req.Header.Add("authority", "chatglm.cn")
	req.Header.Add("accept", "text/event-stream")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("cookie", s.cookie)
	req.Header.Add("referer", "https://chatglm.cn/detail")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "macOS")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return *scanner, err
	}
	defer res.Body.Close()

	scanner = bufio.NewScanner(res.Body)
	return *scanner, nil
}

func (s *ChatService) GetTaskId(prompt string) *entity.TaskResponse {
	url := "https://chatglm.cn/chatglm/backend-api/v1/conversation"
	method := "POST"

	payload := strings.NewReader(`{"prompt":"` + prompt + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("authority", "chatglm.cn")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("authorization", s.authorization)
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("cookie", s.cookie)
	req.Header.Add("origin", "https://chatglm.cn")
	req.Header.Add("referer", "https://chatglm.cn/detail")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "macOS")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var response *entity.TaskResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return response
}

func (s *ChatService) GetContextId(prompt string, taskid string) *entity.ContextResponse {

	url := "https://chatglm.cn/chatglm/backend-api/v1/stream_context"
	method := "POST"

	payload := strings.NewReader(`{"prompt":"` + prompt + `","seed":76049,"max_tokens":512,"conversation_task_id":"` + taskid + `","retry":false,"retry_history_task_id":null}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", s.authorization)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var response *entity.ContextResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return response
}

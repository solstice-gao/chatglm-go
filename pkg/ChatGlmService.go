package pkg

import (
	"bufio"
	"chatglm/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ChatService struct {
	// ... 可能的依赖或其他属性
	authorization string
	cookie        string
}

func NewChatService() *ChatService {
	return &ChatService{
		// ... 初始化依赖或其他属性
		authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTY4ODM2ODQ2NywianRpIjoiY2RkMjU4ZjUtOTkzZi00ZDg5LWFmZTUtMWNhODJlOTAwOWU1IiwidHlwZSI6ImFjY2VzcyIsInN1YiI6Ijk0MWEzZDdkNmQ4YjQ0YjFiMWMxNmMzMmY0YTY3ZmY3IiwibmJmIjoxNjg4MzY4NDY3LCJleHAiOjE2ODg0NTQ4NjcsInJvbGVzIjpbInVuYXV0aGVkX3VzZXIiXX0.-Xgua7nSmBhFiCihwi1WwXhRJtTl0o5SIXbfDUmPvho",
		cookie:        "_ga=GA1.1.1779331165.1685785881; chatglm_refresh_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTY4NTgwMjE1MSwianRpIjoiYzRiZjJjOTktYTgwOS00MWJhLWIwNTktMTVlM2U4YTJmMTJiIiwidHlwZSI6InJlZnJlc2giLCJzdWIiOiI5NDFhM2Q3ZDZkOGI0NGIxYjFjMTZjMzJmNGE2N2ZmNyIsIm5iZiI6MTY4NTgwMjE1MSwiZXhwIjoxNzAxMzU0MTUxLCJyb2xlcyI6WyJ1bmF1dGhlZF91c2VyIl19.L9p18l4agffQAQd5SGbBPNexYi79py_IlBKeGTXQbuo; chatglm_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTY4ODM2ODQ2NywianRpIjoiY2RkMjU4ZjUtOTkzZi00ZDg5LWFmZTUtMWNhODJlOTAwOWU1IiwidHlwZSI6ImFjY2VzcyIsInN1YiI6Ijk0MWEzZDdkNmQ4YjQ0YjFiMWMxNmMzMmY0YTY3ZmY3IiwibmJmIjoxNjg4MzY4NDY3LCJleHAiOjE2ODg0NTQ4NjcsInJvbGVzIjpbInVuYXV0aGVkX3VzZXIiXX0.-Xgua7nSmBhFiCihwi1WwXhRJtTl0o5SIXbfDUmPvho; chatglm_token_expires=2023-07-03%2017:14:27; SL_G_WPT_TO=zh; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; _ga_PMD05MS2V9=GS1.1.1688376755.18.1.1688376759.0.0.0",
	}
}

func (s *ChatService) GetChatStream(contextId string, outputAll bool) {

	url := "https://chatglm.cn/chatglm/backend-api/v1/stream?context_id=" + contextId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
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
		return
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	var lastLine string
	for scanner.Scan() {
		line := scanner.Text()
		if outputAll {
			fmt.Println(line)
		}
		lastLine = line
	}

	if !outputAll && lastLine != "" {
		fmt.Println(lastLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *ChatService) GetTaskId() *entity.TaskResponse {
	url := "https://chatglm.cn/chatglm/backend-api/v1/conversation"
	method := "POST"

	payload := strings.NewReader(`{"prompt":"我想研究的论文方向是chat大模型对未来的影响，请帮我生成论文目录大纲，使用markdown格式生成"}`)

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

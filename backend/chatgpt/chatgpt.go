package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	// "os"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

var messages []Message

func getUserRequest() {
	apiKey := "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"

	var theme = "企業"
	var elements = []string{
    "エンジニア",
    "投資",
	}
	jointedElements := strings.Join(elements, "と")

	question := fmt.Sprintf(`
		日本語で答えてください。
		%sを組み合わせて、%sをテーマにしたアイデアを一つあげてください。
		フォーマットは以下のようにjson形式にしてください
		{ idea:（アイデア), description:（内容説明）内容は50字以内でお願いします。}
`, jointedElements, theme)

	fmt.Println(question)

	messages = append(messages, Message{
		Role:    "user",
		Content: question,
	})

	response := getOpenAIResponse(apiKey)

	fmt.Println(response.Choices[0].Messages.Content)
}

func getOpenAIResponse(apiKey string) OpenaiResponse {
	requestBody := OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	requestJSON, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		println("Error: ", err.Error())
		return OpenaiResponse{}
	}

	messages = append(messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Messages.Content,
	})

	return response
}

func main () {
	getUserRequest()
}
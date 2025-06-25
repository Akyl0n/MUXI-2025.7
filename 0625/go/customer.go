package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Comment struct {
	User    string `json:"user"`
	Comment string `json:"comment"`
}

func main() {
	fmt.Println("发送get/book请求")
	resp1, err := http.Get("http://localhost:8080/book?title=三体")
	if err != nil {
		fmt.Println("GET请求失败", err)
	} else {
		body, _ := io.ReadAll(resp1.Body)
		fmt.Println("GET响应", string(body))
		resp1.Body.Close()
	}

	fmt.Println("发送post请求")
	comment := Comment{
		User:    "小李",
		Comment: "这本书真棒",
	}

	jsonData, _ := json.Marshal(comment)
	if err != nil {
		fmt.Println("error", err)
	}
	resp2, _ := http.Post("http://localhost:8080/comment", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("post error", err)
	} else {
		body, _ := io.ReadAll(resp2.Body)
		fmt.Println("post response : ", string(body))
		resp2.Body.Close()
	}
}

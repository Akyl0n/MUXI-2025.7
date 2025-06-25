// ⽤go编写⼀个简单的服务端并⽀持两类请求，同时模拟客户端操作
// 要求：
// 1. 使⽤net/http编写⼀个web服务程序，监听8080端
// 实现两个结构
// GET /book 通过URL查询参数传⼊书名
// ⽰例：
// 请求
// GET http://localhost:8080/book?title=三体
// 响应
// 您正在查询图书：《三体》
// POST /comment 客户端发送json格式的评论，服务端解析后以json返回
// 结果
// 请求
// {
//  "user": "⼩李",
//  "comment": "这本书真棒！"
// }
// 响应
// {
//  "message": "评论提交成功",
//  "user": "⼩李",
//  "comment": "这本书真棒！"
// }
// 2. 利⽤net/http库模拟客户端对你编写的服务端发送请求



package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Comment struct{
	User string `json:"user"`
	Comment string `json:"comment"`
}

type CommentResponse struct{
	Message string `json:"message"`
	User string `json:"user"`
	Comment string `json:"comment"`
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "缺少title信息", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "你正在查询图书：《%s》", title)
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "非post请求，出错", http.StatusBadRequest)
		return
	}

	var c Comment

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "json解析失败", http.StatusBadRequest)
		return
	}

	response := CommentResponse{
		Message: "评论提交成功",
		User:    c.User,
		Comment: c.Comment,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/book", bookHandler)
	http.HandleFunc("/comment", CommentHandler)
	fmt.Println("服务器已经启动 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
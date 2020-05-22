package main

import "encoding/json"

// 角色/身份
const (
	Student		=	"Student"
	Teacher		= 	"Teacher"
	Reviewer	= 	"Reviewer"
	Admin		=	"Admin"
)

//需审核操作
const (
	Delete 		=	"delete"
	Put			= 	"put"
)

// 用户
type User struct {
	Name string `json:"name"`
	//Id string `json:"id"`
	Role string `json:"role"`
	Metadata []string `json:"metadata"`	//特殊属性，若是教师记录题目id，若是学生记录成绩id
}

// 试题
type Question struct {
	Owner string `json:"owner"` 	//先增加拥有者属性，防止以后需要
	Id string `json:"id"`
	Data string `json:"data"`		//试题题目
	Answer string `json:"answer"`	//试题答案
}

// 试题,不含答案
type QuestionN struct {
	Owner string `json:"owner"` 	//先增加拥有者属性，防止以后需要
	Id string `json:"id"`
	Data string `json:"data"`		//试题题目
}

// 成绩
type ExamScore struct {
	Exam string `json:"exam"`
	Score map[string]string `json:"score"`	//key:name value:score
}

// 日志
type Log struct {
	Id string `json:"id"`
	Time string `json:"time"`		//产生日志的时间
	Name string `json:"name"`		//产生该条日志的用户
	Op string `json:"op"`			//操作
	QuestionId string `json:"question_id"`	//操作目标
}

// 日志打包
type Logs struct {
	Data map[string]Log	`json:"data"`
}

// 打包待审核事件ID
type  ListCache struct {
	PutCache []json.RawMessage `json:"putCache"`
	DelCache []json.RawMessage `json:"delCache"`
}

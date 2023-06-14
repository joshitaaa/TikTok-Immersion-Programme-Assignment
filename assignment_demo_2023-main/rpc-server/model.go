package main

type Message struct{
	sender string `json: "sender"`
	bodytext string `json: "bodytext"`
	time int64 `json: "time"`
}
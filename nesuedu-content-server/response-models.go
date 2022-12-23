package main

import "time"

type User struct {
	Id uint `json:"id"`
}

type Task struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Solved      bool      `json:"solved"`
	Audience    string    `json:"audience"`
	Deadline    time.Time `json:"deadline"`
	Price       uint      `json:"price"`
	CustomerId  uint      `json:"customerId"`
	WorkerId    uint      `json:"workerId"`
}

type SubscribeTask struct {
	WorkerId uint `json:"workerId"`
	TaskId   uint `json:"taskId"`
}

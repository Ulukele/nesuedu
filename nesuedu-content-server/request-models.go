package main

import "time"

// RequestFrom all requests contains user
// that requires that data
type RequestFrom struct {
	UserId uint `json:"userId" validate:"required"`
}

type RequestTasks struct {
	RequestFrom
}

type RequestTask struct {
	RequestFrom
	Id uint `json:"taskId" validate:"required"`
}

type RequestCreateTask struct {
	RequestFrom
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Audience    string    `json:"audience"`
	Deadline    time.Time `json:"deadline"`
	Price       uint      `json:"price"`
}

type RequestUpdateTask struct {
	RequestFrom
	TaskId      uint      `json:"taskId" validate:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Audience    string    `json:"audience"`
	Deadline    time.Time `json:"deadline"`
	Price       uint      `json:"price"`
	Solved      bool      `json:"solved"`
}

type RequestApproveTaskWorker struct {
	RequestFrom
	TaskId   uint `json:"taskId" validate:"required"`
	WorkerId uint `json:"workerId" validate:"required"`
}

type RequestSubscribeTask struct {
	RequestFrom
	TaskId uint `json:"taskId" validate:"required"`
}

type RequestSubsForTask struct {
	RequestFrom
	TaskId uint `json:"taskId" validate:"required"`
}

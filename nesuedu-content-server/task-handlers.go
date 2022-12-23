package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (s *Server) HandleGetTasks(c *fiber.Ctx) error {
	log.Printf("handle get tasks at %s", c.Path())

	req := RequestTasks{}
	req.UserId = c.Locals("userId").(uint)

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	tasks, err := s.contentDBEngine.GetTasks()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can't get tasks")
	}

	resp := make([]Task, 0)

	for _, taskModel := range tasks {
		res := s.SerializeTask(&taskModel)
		resp = append(resp, res)
	}

	return c.JSON(resp)
}

func (s *Server) HandleCreateTask(c *fiber.Ctx) error {
	log.Printf("handle create task at %s", c.Path())

	req := RequestCreateTask{}
	req.UserId = c.Locals("userId").(uint)

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "expect title")
	}

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	taskEntity := Task{
		Title:       req.Title,
		Description: req.Description,
		Audience:    req.Audience,
		Deadline:    req.Deadline,
		Price:       req.Price,
	}
	task, err := s.contentDBEngine.CreateTask(req.UserId, taskEntity)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can't create task")
	}

	return c.JSON(s.SerializeTask(task))
}

func (s *Server) HandleUpdateTask(c *fiber.Ctx) error {
	log.Printf("handle update task at %s", c.Path())

	req := RequestUpdateTask{}
	req.UserId = c.Locals("userId").(uint)
	req.TaskId = c.Locals("taskId").(uint)

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "expect taskId")
	}

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	taskEntity := Task{
		Id:          req.TaskId,
		Title:       req.Title,
		Description: req.Description,
		Audience:    req.Audience,
		Deadline:    req.Deadline,
		Price:       req.Price,
		Solved:      req.Solved,
	}
	task, err := s.contentDBEngine.UpdateTask(req.UserId, taskEntity)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("can't update task: %s", err.Error()))
	}

	return c.JSON(s.SerializeTask(task))
}

func (s *Server) HandleApproveTaskWorker(c *fiber.Ctx) error {
	log.Printf("handle approve task worker at %s", c.Path())

	req := RequestApproveTaskWorker{}
	req.UserId = c.Locals("userId").(uint)
	req.TaskId = c.Locals("taskId").(uint)

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "expect workerId")
	}

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	taskModel, err := s.contentDBEngine.GetTask(req.TaskId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("no such task"))
	}
	taskEntity := s.SerializeTask(taskModel)

	taskEntity.WorkerId = req.WorkerId
	task, err := s.contentDBEngine.UpdateTask(req.UserId, taskEntity)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("can't update task: %s", err.Error()))
	}

	return c.JSON(s.SerializeTask(task))
}

func (s *Server) HandleGetTask(c *fiber.Ctx) error {
	log.Printf("handle get task at %s", c.Path())

	req := RequestTask{}
	req.UserId = c.Locals("userId").(uint)
	req.Id = c.Locals("taskId").(uint)

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	task, err := s.contentDBEngine.GetTask(req.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can't get task")
	}

	return c.JSON(s.SerializeTask(task))
}

func (s *Server) HandleDeleteTask(c *fiber.Ctx) error {
	log.Printf("handle delete task at %s", c.Path())

	req := RequestTask{}
	req.UserId = c.Locals("userId").(uint)
	req.Id = c.Locals("taskId").(uint)

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	_, err := s.contentDBEngine.DeleteTask(req.UserId, req.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can't delete task")
	}

	return c.JSON("")
}

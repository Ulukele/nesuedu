package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (s *Server) HandleGetSubscribeTasks(c *fiber.Ctx) error {
	log.Printf("handle get subscriptions on tasks at %s", c.Path())

	req := RequestFrom{}
	req.UserId = c.Locals("userId").(uint)

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	subscriptions, err := s.contentDBEngine.GetSubscribeTasks(req.UserId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can't get subscriptions")
	}

	resp := make([]SubscribeTask, 0)

	for _, subModel := range subscriptions {
		res := s.SerializeSubscription(&subModel)
		resp = append(resp, res)
	}

	return c.JSON(resp)
}

func (s *Server) HandleGetSubscribeTasksForCustomer(c *fiber.Ctx) error {
	log.Printf("handle get subscriptions on tasks at %s", c.Path())

	req := RequestSubsForTask{}
	req.UserId = c.Locals("userId").(uint)
	req.TaskId = c.Locals("taskId").(uint)

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	subscriptions, err := s.contentDBEngine.GetSubscribeTasksByTaskId(req.TaskId, req.UserId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can't get subscriptions")
	}

	resp := make([]SubscribeTask, 0)

	for _, subModel := range subscriptions {
		res := s.SerializeSubscription(&subModel)
		resp = append(resp, res)
	}

	return c.JSON(resp)
}

func (s *Server) HandleSubscribeTask(c *fiber.Ctx) error {
	log.Printf("handle create subscription task at %s", c.Path())

	req := RequestSubscribeTask{}
	req.UserId = c.Locals("userId").(uint)
	req.TaskId = c.Locals("taskId").(uint)

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	subscription := SubscribeTask{
		WorkerId: req.UserId,
		TaskId:   req.TaskId,
	}

	subscriptionModel, err := s.contentDBEngine.CreateSubscribeTasks(subscription)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("can't subscribe: %s", err.Error()))
	}

	return c.JSON(s.SerializeSubscription(subscriptionModel))
}

func (s *Server) HandleUnSubscribeTask(c *fiber.Ctx) error {
	log.Printf("handle delete subscription task at %s", c.Path())

	req := RequestSubscribeTask{}
	req.UserId = c.Locals("userId").(uint)
	req.TaskId = c.Locals("taskId").(uint)

	if err := validate.Struct(req); err != nil {
		log.Printf("validation error: %s", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "validation error")
	}

	_, err := s.contentDBEngine.DeleteSubscribeTask(req.UserId, req.TaskId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can't unsubscribe task")
	}

	return c.JSON("")
}

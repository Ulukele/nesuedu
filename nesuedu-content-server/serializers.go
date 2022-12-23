package main

func (s *Server) SerializeTask(model *TaskModel) Task {
	return Task{
		Id:          model.Id,
		Title:       model.Title,
		Description: model.Description,
		Solved:      model.Solved,
		Audience:    model.Audience,
		Deadline:    model.Deadline,
		Price:       model.Price,
		CustomerId:  model.CustomerId,
		WorkerId:    model.WorkerId,
	}
}

func (s *Server) SerializeSubscription(model *SubscribeTaskModel) SubscribeTask {
	return SubscribeTask{
		WorkerId: model.WorkerId,
		TaskId:   model.TaskId,
	}
}

package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBEngine struct {
	DB *gorm.DB
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
	Tz       string
}

func NewDBEngine(dbc DBConfig) (*DBEngine, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbc.Host,
		dbc.User,
		dbc.Password,
		dbc.Name,
		dbc.Port,
		dbc.SSLMode,
		dbc.Tz)
	log.Printf("Use config: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbe := &DBEngine{}
	dbe.DB = db
	return dbe, nil
}

func (dbe *DBEngine) GetTasks() ([]TaskModel, error) {

	var tasks []TaskModel
	if err := dbe.DB.
		Where("solved = ?", false).
		Find(&tasks).
		Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (dbe *DBEngine) CreateTask(userId uint, task Task) (*TaskModel, error) {

	taskModel := &TaskModel{
		Title:       task.Title,
		Description: task.Description,
		CustomerId:  userId,
		Audience:    task.Audience,
		Deadline:    task.Deadline,
		Price:       task.Price,
		Solved:      false,
	}

	if err := dbe.DB.Create(&taskModel).Error; err != nil {
		return nil, err
	}

	return taskModel, nil
}

func (dbe *DBEngine) UpdateTask(userId uint, task Task) (*TaskModel, error) {
	taskModel, err := dbe.GetTask(task.Id)
	if err != nil {
		return nil, err
	}

	if taskModel.CustomerId != userId {
		return nil, fmt.Errorf("user not customer")
	}
	if taskModel.Solved && !task.Solved {
		return nil, fmt.Errorf("unable to set completed task as uncompleted")
	}
	if taskModel.WorkerId != 0 {
		return nil, fmt.Errorf("unable to change task with worker")
	}

	taskModel.Title = task.Title
	taskModel.Description = task.Description
	taskModel.Audience = task.Audience
	taskModel.Deadline = task.Deadline
	taskModel.Price = task.Price
	taskModel.Solved = task.Solved
	taskModel.WorkerId = task.WorkerId

	if err := dbe.DB.Save(taskModel).Error; err != nil {
		return nil, err
	}
	return taskModel, nil
}

func (dbe *DBEngine) GetTask(taskId uint) (*TaskModel, error) {

	var task TaskModel
	if err := dbe.DB.
		Where("Id = ?", taskId).
		Take(&task).
		Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (dbe *DBEngine) DeleteTask(userId uint, taskId uint) (*TaskModel, error) {

	task := &TaskModel{}
	if err := dbe.DB.
		Where("Id = ? AND Customer_id = ?", taskId, userId).
		Delete(&task).
		Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (dbe *DBEngine) GetSubscribeTasksByTaskId(taskId uint, userId uint) ([]SubscribeTaskModel, error) {
	task, err := dbe.GetTask(taskId)
	if err != nil {
		return nil, err
	}
	var subscriptions []SubscribeTaskModel
	if err := dbe.DB.
		Model(task).
		Association("Subscribers").
		Find(&subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (dbe *DBEngine) GetSubscribeTasks(workerId uint) ([]SubscribeTaskModel, error) {
	var subscriptions []SubscribeTaskModel
	if err := dbe.DB.
		Where("worker_id = ?", workerId).
		Find(&subscriptions).
		Error; err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (dbe *DBEngine) CheckExistSubscribeTask(taskId uint, workerId uint) (bool, error) {
	var exist bool
	if err := dbe.DB.
		Model(&SubscribeTaskModel{}).
		Select("count(*) > 0").
		Where("task_id = ? AND worker_id = ?", taskId, workerId).
		Find(&exist).
		Error; err != nil {
		return false, err
	}

	return exist, nil
}

func (dbe *DBEngine) CreateSubscribeTasks(subscription SubscribeTask) (*SubscribeTaskModel, error) {
	exist, err := dbe.CheckExistSubscribeTask(subscription.TaskId, subscription.WorkerId)
	if err != nil {
		return nil, fmt.Errorf("unable to check if exist subscription: %s", err.Error())
	}

	if exist {
		return nil, fmt.Errorf("already exist")
	}

	task, err := dbe.GetTask(subscription.TaskId)
	if err != nil {
		return nil, err
	}

	subscriptionModel := &SubscribeTaskModel{WorkerId: subscription.WorkerId, TaskId: subscription.TaskId}
	if err := dbe.DB.
		Model(&task).
		Association("Subscribers").
		Append(subscriptionModel); err != nil {
		return nil, err
	}

	return subscriptionModel, nil
}

func (dbe *DBEngine) DeleteSubscribeTask(userId uint, taskId uint) (*SubscribeTaskModel, error) {

	subscriptionModel := &SubscribeTaskModel{}
	if err := dbe.DB.
		Where("Worker_id = ? AND Task_id = ?", userId, taskId).
		Delete(&subscriptionModel).
		Error; err != nil {
		return nil, err
	}

	return subscriptionModel, nil
}

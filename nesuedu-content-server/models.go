package main

import (
	"gorm.io/gorm"
	"time"
)

// UserModel shadow model -- don't exist in db
type UserModel struct {
	Id uint
}

type TaskModel struct {
	gorm.Model
	Id          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Solved      bool
	Audience    string
	Deadline    time.Time
	Price       uint
	CustomerId  uint
	WorkerId    uint
	Subscribers []SubscribeTaskModel `gorm:"foreignKey:TaskId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type SubscribeTaskModel struct {
	gorm.Model
	WorkerId uint
	TaskId   uint
}

func (dbe *DBEngine) initTables() error {

	if err := dbe.DB.AutoMigrate(&SubscribeTaskModel{}); err != nil {
		return err
	} else if err := dbe.DB.AutoMigrate(&TaskModel{}); err != nil {
		return err
	}

	return nil
}

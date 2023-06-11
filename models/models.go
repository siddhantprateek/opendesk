package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	TaskTitle   string             `json:"task_title,omitempty" validate:"required"`
	TaskTag     string             `json:"task_tag,omitempty" validate:"required"`
	TimeLeft    string             `json:"time_left,omitempty" validate:"required"`
	DateAdded   string             `json:"date_added,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	Created     string             `json:"created,omitempty" validate:"required"`
	TaskStatus  bool               `json:"task_status,omitempty"`
}

type Quote struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Quote  string             `json:"quote,omitempty" validate:"required"`
	Author string             `json:"author,omitempty" validate:"required"`
	Slug   string             `json:"slug,omitempty" validate:"required"`
}

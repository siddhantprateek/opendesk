package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	connection "github.com/nikhitabarat/opendesk/connection"
	models "github.com/nikhitabarat/opendesk/models"
	responses "github.com/nikhitabarat/opendesk/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var employeeCollection *mongo.Collection = connection.GetCollection(connection.DB, "opendeskdata")
var validate = validator.New()

// @desp: Add Employee Task
func AddTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var task models.Task
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.ResponseType{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	// validate the request body
	if validationErr := validate.Struct(&task); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.ResponseType{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": validationErr.Error()}})
	}

	newTask := models.Task{
		Id:          primitive.NewObjectID(),
		TaskTitle:   task.TaskTitle,
		TaskTag:     task.TaskTag,
		TimeLeft:    task.TaskTag,
		DateAdded:   task.DateAdded,
		Description: task.Description,
		Created:     task.Created,
		TaskStatus:  task.TaskStatus,
	}

	// fmt.Println(task.TaskTitle, task.TaskTag, task.TaskTag, task.DateAdded, task.Description, task.Created, task.TaskStatus)

	result, err := employeeCollection.InsertOne(ctx, newTask)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.ResponseType{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(
		responses.ResponseType{
			Status:  http.StatusCreated,
			Message: "error",
			Data:    &fiber.Map{"data": result}},
	)
}

// @desp: Delete Employee Task by Id
func DeleteTaskbyId(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	taskId := c.Params("taskId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(taskId)

	result, err := employeeCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.ResponseType{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.ResponseType{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    &fiber.Map{"data": "Task with specified Id not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ResponseType{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": "Task successfully deleted!"}},
	)
}

// @desp: Update Employee Task
func UpdateTask(c *fiber.Ctx) error {
	return nil
}

// @desp: Get Employee Tasks
func GetEmployeeTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var tasks []models.Task
	defer cancel()

	results, err := employeeCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.ResponseType{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	// Reading from the database in an optimal proccess
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTask models.Task
		if err = results.Decode(&singleTask); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.ResponseType{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()}},
			)
		}
		tasks = append(tasks, singleTask)
	}
	return c.Status(http.StatusOK).JSON(
		responses.ResponseType{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": tasks}},
	)
}

// @desp: Get Employee Task by Id
func GetTaskbyId(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	taskId := c.Params("taskId")
	var task models.Task
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(taskId)

	err := employeeCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.ResponseType{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ResponseType{
			Status:  http.StatusOK,
			Message: "error",
			Data:    &fiber.Map{"data": task}},
	)
}

// @desp: Add Assignee to Task
func AddAssigneeToTask(c *fiber.Ctx) error {
	return nil
}

// @desp: Remove Assignee from Task
func RemoveAssigneeTask(c *fiber.Ctx) error {
	return nil
}

// @desp: Track Employee Progress
func GetEmployeeProgress(c *fiber.Ctx) error {
	return nil
}

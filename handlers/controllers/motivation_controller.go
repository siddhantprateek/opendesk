package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	connection "github.com/nikhitabarat/opendesk/connection"
	models "github.com/nikhitabarat/opendesk/models"
	responses "github.com/nikhitabarat/opendesk/res"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var QuoteCollection *mongo.Collection = connection.GetCollection(connection.DB, "quotesdata")

// @desp: GET motivation quotes
func GetAllQuotes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var quotes []models.Quote
	defer cancel()

	results, err := QuoteCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.ResponseType{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}},
		)
	}

	// Reading from the database
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuote models.Quote
		if err = results.Decode(&singleQuote); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.ResponseType{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()}},
			)
		}
		quotes = append(quotes, singleQuote)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ResponseType{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": quotes}},
	)
}

// @desp: POST motivation quotes
func CreateMotivationQuotes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var quote models.Quote
	defer cancel()

	// request body
	if err := c.BodyParser(&quote); err != nil {
		return c.Status(http.StatusBadGateway).JSON(
			responses.ResponseType{
				Status:  http.StatusBadGateway,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	// validate the request body
	if validationErr := validate.Struct(&quote); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.ResponseType{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": validationErr.Error()}})
	}

	newQuote := models.Quote{
		Id:     primitive.NewObjectID(),
		Quote:  quote.Quote,
		Author: quote.Author,
		Slug:   quote.Quote,
	}

	result, err := QuoteCollection.InsertOne(ctx, newQuote)
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

// @desp: PATCH motivation quotes
func UpdateMotivationQuotes(c *fiber.Ctx) error {
	return nil
}

// @desp: DELETE motivation quotes
func DeleteQuotes(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	quoteId := c.Params("quoteId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(quoteId)

	result, err := QuoteCollection.DeleteOne(ctx, bson.M{"id": objId})
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
				Data:    &fiber.Map{"data": "Quote with specified Id not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ResponseType{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": "Quote deleted successfully"}},
	)
}

// @desp: GET motivation quotes by Id
func GetQuotesbyId(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	quoteId := c.Params("quoteId")
	var quote models.Quote
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(quoteId)

	err := QuoteCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&quote)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.ResponseType{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(
		responses.ResponseType{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": quote}})
}

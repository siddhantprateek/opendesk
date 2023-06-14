package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	connection "github.com/siddhantprateek/opendesk/storage/connection"
	models "github.com/siddhantprateek/opendesk/storage/models"
	responses "github.com/siddhantprateek/opendesk/storage/res"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	quoteID := c.Params("quoteId")

	// Parse the quote ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(quoteID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseType{
			Status:  http.StatusBadRequest,
			Message: "Invalid quote ID",
			Data:    nil,
		})
	}

	// Retrieve the existing quote from the database
	filter := bson.M{"_id": objID}
	var existingQuote models.Quote
	err = QuoteCollection.FindOne(ctx, filter).Decode(&existingQuote)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(responses.ResponseType{
				Status:  http.StatusNotFound,
				Message: "Quote not found",
				Data:    nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseType{
			Status:  http.StatusInternalServerError,
			Message: "Error retrieving quote",
			Data:    nil,
		})
	}

	// Get the updated quote data from the request body
	var updatedQuote models.Quote
	err = c.BodyParser(&updatedQuote)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseType{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	// Update only the necessary fields
	existingQuote.Quote = updatedQuote.Quote
	existingQuote.Author = updatedQuote.Author

	// Update the quote in the database
	update := bson.M{
		"$set": existingQuote,
	}
	_, err = QuoteCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseType{
			Status:  http.StatusInternalServerError,
			Message: "Error updating quote",
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responses.ResponseType{
		Status:  http.StatusOK,
		Message: "Quote updated successfully",
		Data:    &fiber.Map{"data": existingQuote},
	})
}

// @desp: DELETE motivation quotes
func DeleteQuotes(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	quoteId := c.Params("quoteId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(quoteId)

	result, err := QuoteCollection.DeleteOne(ctx, bson.M{"_id": objId})
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

	err := QuoteCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&quote)
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

// @desc: GET all quotes by author's name
func GetAllQuotesByAuthor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	authorName := c.Query("author")

	filter := bson.M{"author": authorName}
	var quotes []models.Quote

	results, err := QuoteCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseType{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuote models.Quote
		if err = results.Decode(&singleQuote); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ResponseType{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}
		quotes = append(quotes, singleQuote)
	}

	return c.Status(http.StatusOK).JSON(responses.ResponseType{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"data": quotes},
	})
}

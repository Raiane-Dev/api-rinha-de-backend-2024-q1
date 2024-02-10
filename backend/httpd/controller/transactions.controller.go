package controller

import (
	"log"
	"net/http"
	"rinha_api/backend/entity/io/input"
	"rinha_api/backend/entity/io/output"
	"rinha_api/backend/model"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func SendTransaction(c *fiber.Ctx) (_ error) {

	var (
		transaction_input input.TransactionInput

		client model.ClientEntity
	)

	if err := c.BodyParser(&transaction_input); err != nil {
		log.Println(err)
		c.
			Status(http.StatusBadRequest).
			WriteString("Invalid json body")
		return
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.
			Status(http.StatusBadRequest).
			WriteString("ID not number")
	}

	if err := validate.Struct(&transaction_input); err != nil {
		log.Println(err)
		c.
			Status(http.StatusBadRequest).
			WriteString("Invalid data")
		return
	}

	client, err = client.FindBy("id = ?", id)
	if err != nil {
		c.
			Status(http.StatusNotFound).
			WriteString("Not found client")
		return
	}

	if transaction_input.Type == "d" {
		if (client.Balance + transaction_input.Value) > client.Limit {
			c.
				Status(http.StatusUnprocessableEntity).
				WriteString("Limit exceeded")
			return
		}
	}

	{
		new_transaction := model.TransactionEntity{
			ClientID:    id,
			Type:        transaction_input.Type,
			Value:       transaction_input.Value,
			Description: transaction_input.Description,
		}
		if err := new_transaction.Create(); err != nil {
			c.
				Status(http.StatusBadRequest).
				WriteString("Unable to complete the transaction")
			return
		}

		client.Balance = client.Balance + transaction_input.Value
		if err := client.Update("id = ?", id); err != nil {
			log.Println(err)
			c.
				Status(http.StatusBadRequest).
				WriteString("Unable to complete the transaction")
			return
		}
	}

	return c.Status(http.StatusOK).JSON(output.TransactionOutput{
		Limit:   client.Limit,
		Balance: client.Balance,
	})
}

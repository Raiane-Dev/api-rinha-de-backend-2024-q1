package controller

import (
	"log"
	"net/http"
	"rinha_api/backend/entity/io/output"
	"rinha_api/backend/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ConsultTransaction(c *fiber.Ctx) (_ error) {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.
			Status(http.StatusBadRequest).
			WriteString("ID not number")
	}

	var (
		client      model.ClientEntity
		transaction model.TransactionEntity
	)

	client, err = client.FindBy("id = ?", id)
	if err != nil {
		c.
			Status(http.StatusNotFound).
			WriteString("Not found client")
		return
	}

	transactions, err := transaction.FindMany("cliente_id = ?", id)
	if err != nil {
		log.Println("err in request", err)
		c.
			Status(http.StatusNotFound).
			WriteString("Not found transactions")
		return
	}

	out := output.ExtractOutput[model.TransactionEntity]{
		Balance: output.Balance{
			Total:       client.Balance,
			ExtractDate: time.Now().Format(time.RFC3339),
			Limit:       client.Limit,
		},
		LatestTransaction: transactions,
	}

	return c.Status(http.StatusOK).JSON(out)
}

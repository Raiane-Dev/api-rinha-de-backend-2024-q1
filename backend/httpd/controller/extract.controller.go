package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rinha_api/backend/entity/io/output"
	"rinha_api/backend/model"
	"strings"

	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/valyala/fasthttp"
)

func ConsultTransaction(ctx *fasthttp.RequestCtx) {

	id := ctx.UserValue("id").(string)
	client, err := model.FindByClient("id = ?", id)
	if err != nil {
		// config.ErrReaderDB <- err
		log.Println("not found client", err)
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	tsct := model.FindManyTransactions("cliente_id = ?", id)

	resp_json, _ := jsoniter.Marshal(
		output.ExtractOutput{
			Balance: output.Balance{
				Total:       client.Balance,
				ExtractDate: time.Now().Format(time.RFC3339),
				Limit:       client.Limit,
			},
			LatestTransaction: json.RawMessage(strings.Join(tsct, ",")),
		},
	)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(resp_json)

}

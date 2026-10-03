package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"booking-app/api"
	"booking-app/internal/handlers/tools"

	"github.com/gorilla/schema"

	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)

	if tokenDetails == nil {
		log.Error("No Coin Details Found")
		api.RequestErrorHandler(w, errors.New("No Coin Details Found"))
		return
	}

	var response = api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: (*tokenDetails).Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}

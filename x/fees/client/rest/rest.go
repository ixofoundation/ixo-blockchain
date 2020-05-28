package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	RestFeeId         = "fee_id"
	RestFeeContractId = "fee_contract_id"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/fees/params",
		queryFeeParamsHandler(cliCtx)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/fees/{%s}", RestFeeId),
		queryFeeHandler(cliCtx)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/fee/contracts/{%s}", RestFeeContractId),
		queryFeeContractHandler(cliCtx)).Methods("GET")
}

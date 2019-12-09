package rest

import (
	"encoding/json"
	"net/http"

	cTypes "github.com/cosmos/cosmos-sdk/types"
)

const DefaultCodeSpace = "ixo"

type ErrorResponse struct {
	Success   bool                 `json:"success"`
	Error     interface{}          `json:"error"`
	Code      cTypes.CodeType      `json:"code"`
	CodeSpace cTypes.CodespaceType `json:"codeSpace"`
}

type Log struct {
	CodeType  cTypes.CodeType
	CodeSpace cTypes.CodespaceType
	Message   string
}

func NewErrorResponse(success bool, err interface{}, code cTypes.CodeType, codeSpace cTypes.CodespaceType) ErrorResponse {
	return ErrorResponse{
		Success:   success,
		Error:     err,
		Code:      code,
		CodeSpace: codeSpace,
	}
}

func WriteErrorResponse(w http.ResponseWriter, err cTypes.Error) {
	var log Log
	_err := json.Unmarshal([]byte(err.Result().Log), &log)
	if _err != nil {
		panic(_err)
	}
	errResponse := NewErrorResponse(false, log.Message, err.Code(), err.Codespace())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	bz, _ := json.Marshal(errResponse)
	w.Write(bz)
	return
}

func WriteError(w http.ResponseWriter, err error) {
	var abci []cTypes.ABCIMessageLog
	_err := json.Unmarshal([]byte(err.Error()), &abci)
	if _err != nil {
		panic(_err)
	}

	var log Log
	_err = json.Unmarshal([]byte(abci[0].Log), &log)
	if _err != nil {
		panic(_err)
	}

	errResponse := NewErrorResponse(false, log.Message, log.CodeType, log.CodeSpace)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	bz, _ := json.Marshal(errResponse)
	w.Write(bz)
	return
}

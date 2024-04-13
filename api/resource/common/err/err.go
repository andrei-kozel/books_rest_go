package err

import "net/http"

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}

var (
	RespDBDataInsertFailure = []byte(`{"error":"Failed to insert data into database"}`)
	RespDBDataAccessFailure = []byte(`{"error":"Failed to access data from database"}`)
	RespDBDataUpdateFailure = []byte(`{"error":"Failed to update data in database"}`)
	RespDBDataRemoveFailure = []byte(`{"error":"Failed to remove data from database"}`)

	RespJSONEncodeFailure = []byte(`{"error":"Failed to encode data into JSON"}`)
	RespJSONDecodeFailure = []byte(`{"error":"Failed to decode JSON data"}`)

	RespInvalidURLParamID = []byte(`{"error":"Invalid URL parameter ID"}`)
)

func ServerError(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(resp)
}

func BadRequest(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp)
}

func ValidationError(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(resp)
}

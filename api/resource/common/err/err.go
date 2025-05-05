package err

import (
	"net/http"
)

var (
	RespClusterCreateFailure = []byte(`{"error": "Failed to create a cluster"}`)
	RespClusterAccessFailure = []byte(`{"error": "Failed to access the cluster"}`)
	RespClusterUpdateFailure = []byte(`{"error": "Failed to update the cluster"}`)
	RespClusterRemoveFailure = []byte(`{"error": "Failed to remove the cluster"}`)

	RespJSONEncodeFailure = []byte(`{"error": "json encode failure"}`)
	RespJSONDecodeFailure = []byte(`{"error": "json decode failure"}`)

	RespInvalidURLParamID = []byte(`{"error": "invalid url param-id"}`)
)

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

func ServerError(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(error)
}

func BadRequest(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(error)
}

func ValidationErrors(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(reps)
}

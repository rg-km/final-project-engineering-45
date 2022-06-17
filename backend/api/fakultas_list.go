package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type FakultasListErrorResponse struct {
	Error string `json:"error"`
}

type Fakultas struct {
	ID           int64     `json:"id"`
	FakultasName string    `json:"fakultas_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type FakultasListSuccessResponse struct {
	Fakultas []Fakultas `json:"fakultas"`
}

func (api *API) fakultasList(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	encoder := json.NewEncoder(w)

	response := FakultasListSuccessResponse{}
	response.Fakultas = make([]Fakultas, 0)

	fakultas, err := api.fakultasRepo.FetchFakultas()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(FakultasListErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, fakultas := range fakultas {
		response.Fakultas = append(response.Fakultas, Fakultas{
			ID:           fakultas.ID,
			FakultasName: fakultas.FakultasName,
			CreatedAt:    fakultas.CreatedAt,
		})
	}

	encoder.Encode(response)
}

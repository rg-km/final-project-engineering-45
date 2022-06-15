package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type ProdiListErrorResponse struct {
	Error string `json:"error"`
}

type Prodi struct {
	ID           int64     `json:"id"`
	ProdiName    string    `json:"prodi_name"`
	FakultasID   int64     `json:"fakultas_id"`
	FakultasName string    `json:"fakultas_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type ProdiListSuccessResponse struct {
	Prodi []Prodi `json:"program_studi"`
}

func (api *API) prodiList(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	encoder := json.NewEncoder(w)

	response := ProdiListSuccessResponse{}
	response.Prodi = make([]Prodi, 0)

	prodi, err := api.prodiRepo.FetchProdi()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, prodi := range prodi {
		response.Prodi = append(response.Prodi, Prodi{
			ID:           prodi.ID,
			ProdiName:    prodi.ProdiName,
			FakultasID:   prodi.FakultasID,
			FakultasName: prodi.FakultasName,
			CreatedAt:    prodi.CreatedAt,
		})
	}

	encoder.Encode(response)
}

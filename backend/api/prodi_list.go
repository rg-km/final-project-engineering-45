package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rg-km/final-project-engineering-45/backend/repository"
)

type ProdiListErrorResponse struct {
	Error string `json:"error"`
}

type AddToProdiRequest struct {
	ProdiName  string `json:"prodi_name"`
	FakultasID int64  `json:"fakultas_ID"`
}

type AddToProdiSuccessResponse struct {
	Prodi []repository.ProgramStudi `json:"program_studi"`
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

// func (api *API) addProdi(w http.ResponseWriter, r *http.Request) {
// 	api.AllowOrigin(w, r)
// 	encoder := json.NewEncoder(w)

// 	var prodi Prodi
// 	err := json.NewDecoder(r.Body).Decode(&prodi)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	err = api.prodiRepo.InsertProdi(prodi.ProdiName, prodi.FakultasName, time.Now())
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	encoder.Encode(ProdiListSuccessResponse{Prodi: []Prodi{prodi}})
// }

// func (api *API) addProdi(w http.ResponseWriter, r *http.Request) {
// 	api.AllowOrigin(w, r)

// 	var request AddToProdiRequest
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	encoder := json.NewEncoder(w)

// 	response := AddToProdiSuccessResponse{}
// 	response.Prodi = make([]repository.ProgramStudi, 0)

// 	prodi, err := api.prodiRepo.FetchProdiByName(request.ProdiName)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	if prodi.ProdiName == request.ProdiName {
// 		w.WriteHeader(http.StatusBadRequest)
// 		encoder.Encode(ProdiListErrorResponse{Error: "Prodi already exists"})
// 		return
// 	} else {
// 		err = api.prodiRepo.InsertProdi(request.ProdiName, request.FakultasID)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	fakultas, err := api.fakultasRepo.FetchFakultasByID(request.FakultasID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	if fakultas.ID == request.FakultasID {
// 		w.WriteHeader(http.StatusOK)
// 		encoder.Encode(response)
// 	} else {
// 		w.WriteHeader(http.StatusBadRequest)
// 		encoder.Encode(ProdiListErrorResponse{Error: "Fakultas not found"})
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	encoder.Encode(AddToProdiSuccessResponse{Prodi: []repository.ProgramStudi{prodi}})
// 	w.Write([]byte("Prodi added"))
// }

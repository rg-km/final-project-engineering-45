package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ProdiListErrorResponse struct {
	Error string `json:"error"`
}

type AddToProdiRequest struct {
	ProdiName    string    `json:"prodi_name"`
	FakultasName string    `json:"fakultas_name"`
	Deskripsi    string    `json:"deskripsi"`
	Karakter     string    `json:"karakter"`
	MataKuliah   string    `json:"mata_kuliah"`
	Prospek      string    `json:"prospek"`
	CreatedAt    time.Time `json:"created_at"`
}

type Prodi struct {
	ID           int64     `json:"id"`
	ProdiName    string    `json:"prodi_name"`
	FakultasID   int64     `json:"fakultas_id"`
	FakultasName string    `json:"fakultas_name"`
	Deskripsi    string    `json:"deskripsi"`
	Karakter     string    `json:"karakter"`
	MataKuliah   string    `json:"mata_kuliah"`
	Prospek      string    `json:"prospek"`
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
			Deskripsi:    prodi.Deskripsi,
			Karakter:     prodi.Karakter,
			MataKuliah:   prodi.MataKuliah,
			Prospek:      prodi.Prospek,
			CreatedAt:    prodi.CreatedAt,
		})
	}

	encoder.Encode(response)
}

func (api *API) selectProdiByName(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	encoder := json.NewEncoder(w)

	prodiName := r.URL.Query().Get("prodi_name")
	prodiName = strings.Replace(prodiName, "%20", " ", -1)

	prodi, err := api.prodiRepo.FetchProdiByName(prodiName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
		return
	}

	response := ProdiListSuccessResponse{}
	response.Prodi = make([]Prodi, 0)

	//dont need to loop prodi
	response.Prodi = append(response.Prodi, Prodi{
		ID:           prodi.ID,
		ProdiName:    prodi.ProdiName,
		FakultasID:   prodi.FakultasID,
		FakultasName: prodi.FakultasName,
		Deskripsi:    prodi.Deskripsi,
		Karakter:     prodi.Karakter,
		MataKuliah:   prodi.MataKuliah,
		Prospek:      prodi.Prospek,
		CreatedAt:    prodi.CreatedAt,
	})

	encoder.Encode(response)
}

func (api *API) selectProdi(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	encoder := json.NewEncoder(w)

	prodiName := r.URL.Query().Get("prodi_name")
	prodiName = strings.Replace(prodiName, "%20", " ", -1)

	fakultasName := r.URL.Query().Get("fakultas_name")
	fakultasName = strings.Replace(fakultasName, "%20", " ", -1)

	prodi, err := api.prodiRepo.FetchProdiAndFakultasName(fakultasName, prodiName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
		return
	}

	response := ProdiListSuccessResponse{}
	response.Prodi = make([]Prodi, 0)

	response.Prodi = append(response.Prodi, Prodi{
		ID:           prodi.ID,
		ProdiName:    prodi.ProdiName,
		FakultasID:   prodi.FakultasID,
		FakultasName: prodi.FakultasName,
		Deskripsi:    prodi.Deskripsi,
		Karakter:     prodi.Karakter,
		MataKuliah:   prodi.MataKuliah,
		Prospek:      prodi.Prospek,
		CreatedAt:    prodi.CreatedAt,
	})

	if prodi.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(ProdiListErrorResponse{Error: "Data not found"})
		return
	} else if prodi.ProdiName != prodiName {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(ProdiListErrorResponse{Error: "Program Studi not found"})
		return
	} else if prodi.FakultasName != fakultasName {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(ProdiListErrorResponse{Error: "Fakultas not found"})
		return
	}

	encoder.Encode(response)
	w.WriteHeader(http.StatusOK)
}

func (api *API) prodiListByFakultasName(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	encoder := json.NewEncoder(w)

	fakultasName := r.URL.Query().Get("fakultas_name")
	fakultasName = strings.Replace(fakultasName, "%20", " ", -1)

	prodi, err := api.prodiRepo.FetchProdiByFakultasName(fakultasName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
		return
	}

	response := ProdiListSuccessResponse{}
	response.Prodi = make([]Prodi, 0)

	//loop prodi
	for _, prodi := range prodi {
		response.Prodi = append(response.Prodi, Prodi{
			ID:           prodi.ID,
			ProdiName:    prodi.ProdiName,
			FakultasID:   prodi.FakultasID,
			FakultasName: prodi.FakultasName,
			Deskripsi:    prodi.Deskripsi,
			Karakter:     prodi.Karakter,
			MataKuliah:   prodi.MataKuliah,
			Prospek:      prodi.Prospek,
			CreatedAt:    prodi.CreatedAt,
		})
		if prodi.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			encoder.Encode(ProdiListErrorResponse{Error: "Program Studi not found"})
			return
		}
	}

	encoder.Encode(response)
	w.WriteHeader(http.StatusOK)
}

func (api *API) addProdi(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)

	var request AddToProdiRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)

	prodi, _ := api.prodiRepo.FetchProdiByName(request.ProdiName)

	if prodi.ProdiName == request.ProdiName {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ProdiListErrorResponse{Error: "Prodi already exists"})
		return
	} else if prodi.ProdiName != request.ProdiName {
		fakultas, _ := api.fakultasRepo.FetchFakultasByName(request.FakultasName)

		if fakultas.FakultasName != request.FakultasName {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ProdiListErrorResponse{Error: "Fakultas not found"})
			return
		} else if fakultas.FakultasName == request.FakultasName {
			w.WriteHeader(http.StatusOK)
			//encoder.Encode(response)
			err = api.prodiRepo.InsertProdi(request.ProdiName, request.FakultasName, request.Deskripsi, request.Karakter, request.MataKuliah, request.Prospek)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				encoder.Encode(ProdiListErrorResponse{Error: err.Error()})
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("New Program Studi has been added"))
	fmt.Fprintf(w, "Program Studi %s has been added", request.ProdiName)
}

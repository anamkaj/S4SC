package client

import (
	"calibri/internal/models"
	"calibri/internal/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	store models.ClientInterface
}

func NewHandler(store models.ClientInterface) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("GET /api/call-tracker/client-list", http.HandlerFunc(h.handleClientList))
	router.Handle("POST /api/call-tracker/client-data", http.HandlerFunc(h.handleClientsData))
	router.Handle("POST /api/call-tracker/single-client", http.HandlerFunc(h.handleSingleClient))

}

func (h *Handler) handleClientList(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	if status == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Ошибка: отсутствует параметр status")
		return
	}

	if status != "true" && status != "false" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Ошибка: некорректное значение параметра status")
		return
	}

	if status == "true" {
		data, err := h.store.GetClientList(true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		utils.ResJson(w, http.StatusOK, data)
		return
	}

	if status == "false" {
		data, err := h.store.GetClientList(false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		utils.ResJson(w, http.StatusOK, data)
		return
	}

}

func (h *Handler) handleClientsData(w http.ResponseWriter, r *http.Request) {

	dateStart := r.URL.Query().Get("date_start")
	dateEnd := r.URL.Query().Get("date_end")

	if dateStart == "" || dateEnd == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Ошибка: отсутствует параметр date_start или date_end")
		return
	}

	_, err := time.Parse("2006-01-02", dateStart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Не правильный формат даты: %s", err.Error())
		return
	}

	_, err = time.Parse("2006-01-02", dateEnd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Не правильный формат даты: %s", err.Error())
		return
	}

	data, err := h.store.GetFullDataAllClients(dateStart, dateEnd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	utils.ResJson(w, http.StatusOK, data)

}

func (h *Handler) handleSingleClient(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	dateStart := r.URL.Query().Get("date_start")
	dateEnd := r.URL.Query().Get("date_end")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Ошибка: отсутствует параметр id")
		return
	}

	if dateStart == "" || dateEnd == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Ошибка: отсутствует параметр date_start или date_end")
		return
	}

	_, err := time.Parse("2006-01-02", dateStart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Не правильный формат даты: %s", err.Error())
		return
	}

	_, err = time.Parse("2006-01-02", dateEnd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Не правильный формат даты: %s", err.Error())
		return
	}

	num, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error converting id: %s", err.Error())
		return
	}

	data, err := h.store.GetSingleClient(num, dateStart, dateEnd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	utils.ResJson(w, http.StatusOK, data)

}

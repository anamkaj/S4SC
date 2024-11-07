package client

import (
	"calibri/internal/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockClientStore struct{}

func TestRegisterRoutes(t *testing.T) {

	store := &mockClientStore{}
	handler := NewHandler(store)

	t.Run("handleClientList", func(t *testing.T) {

		test := []struct {
			status string
		}{
			{"true"},
			{"false"},
			{"wrong"},
			{""},
		}

		for _, tt := range test {
			t.Run(tt.status, func(t *testing.T) {

				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/call-tracker/client-list?status="+tt.status, nil)

				handler.handleClientList(w, r)
				if w.Code != http.StatusOK {
					t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
				}

				if w.Body == nil {
					t.Errorf("Expected response body, got nil")
				}
			})
		}

	})

	t.Run("handleSingleClient", func(t *testing.T) {

		test := []struct {
			id        string
			dateStart string
			dateEnd   string
		}{
			{"1", "2022-01-01", "2022-01-02"},
			{"wrong", "2022-01-01", "2022-01-02"},
			{"", "2022-01-01", "2022-01-02"},
			{"1", "wrong", "2022-01-02"},
			{"1", "2022-01-01", "wrong"},
			{"1", "", "2022-01-02"},
			{"1", "2022-01-01", ""},
			{"1", "01-01-2023", "2022-01-02"},
		}

		for _, tt := range test {
			t.Run(tt.id, func(t *testing.T) {

				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/call-tracker/single-client?id="+tt.id+"&date_start="+tt.dateStart+"&date_end="+tt.dateEnd, nil)

				handler.handleSingleClient(w, r)
				if w.Code != http.StatusOK {
					t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
				}

				if w.Body == nil {
					t.Errorf("Expected response body, got nil")
				}

				fmt.Println(w.Body)
			})
		}
	})

	t.Run("handleClientsData", func(t *testing.T) {

		test := []struct {
			dateStart string
			dateEnd   string
		}{
			{"2022-01-01", "2022-01-02"},
			{"wrong", "2022-01-02"},
			{"2022-01-01", "wrong"},
			{"01-01-2023", "2022-01-02"},
		}

		for _, tt := range test {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/call-tracker/clients-data?date_start="+tt.dateStart+"&date_end="+tt.dateEnd, nil)

			handler.handleClientsData(w, r)
			if w.Code != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
			}

			if w.Body == nil {
				t.Errorf("Expected response body, got nil")
			}

			fmt.Println(w.Body)
		}

	})
}

func (m *mockClientStore) GetClientList(bool) (*[]models.ClientCalibri, error) {
	return &[]models.ClientCalibri{}, nil
}

func (m *mockClientStore) GetFullDataAllClients(string, string) (*[]models.CallAndEmail, error) {
	return &[]models.CallAndEmail{}, nil
}

func (m *mockClientStore) GetSingleClient(int, string, string) (*models.CallAndEmail, error) {
	return &models.CallAndEmail{}, nil
}

func GetClientList(bool) (*[]models.ClientCalibri, error) {
	return &[]models.ClientCalibri{}, nil
}

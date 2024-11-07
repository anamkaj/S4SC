package client

import (
	"calibri/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetClientList(status bool) (*[]models.ClientCalibri, error) {

	qStatusTrue := `SELECT 
                    cc.id,cc.site_id,cc.sitename,cc.domains,cc.active,cc.license_start,cc.license_end,cc.not_enough_money,ph.number 
                    FROM client_calibri cc JOIN phone ph ON cc.site_id=ph.client_calibri_site_id_fk 
                    WHERE active=$1;`

	qStatusFalse := `SELECT 
                cc.id, 
                cc.site_id, 
                cc.sitename, 
                cc.domains, 
                cc.active, 
                cc.license_start, 
                cc.license_end, 
                cc.not_enough_money, 
                ph.number 
            FROM 
                client_calibri cc 
            LEFT JOIN 
                phone ph ON cc.site_id = ph.client_calibri_site_id_fk 
            WHERE 
                cc.active = $1;`

	data := []models.ClientCalibri{}

	if status {
		err := s.db.Select(&data, qStatusTrue, status)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return nil, err
		}
	} else {
		err := s.db.Select(&data, qStatusFalse, status)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return nil, err
		}
	}

	return &data, nil

}

func (s *Store) GetFullDataAllClients(start string, end string) (*[]models.CallAndEmail, error) {

	activeID := `SELECT site_id FROM client_calibri WHERE active=$1;`

	ids := []int{}
	err := s.db.Select(&ids, activeID, true)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil, err
	}

	calls := []models.Calls{}
	email := []models.Email{}
	data := []models.CallAndEmail{}

	for _, id := range ids {
		qCalls := fmt.Sprintf(
			`SELECT id, call_id, date, channel_id, source, is_lid, name_type, traffic_type, landing_page, conversations_number, call_status
			FROM calls
			WHERE calls.client_calibri_site_id_fk = %d
			AND 
			TO_TIMESTAMP(calls.date, 'YYYY-MM-DD"T"HH24:MI:SS.USZ') 
			BETWEEN 
			TO_TIMESTAMP('%sT00:00:00.000Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ')
			AND 
			TO_TIMESTAMP('%sT23:59:59.999Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ');`,
			id, start, end,
		)

		qEmails := fmt.Sprintf(
			`SELECT id, email_id, date, source, is_lid, traffic_type, landing_page, conversations_number FROM email 
			WHERE email.client_calibri_site_id_fk = %d
			AND 
			TO_TIMESTAMP(email.date, 'YYYY-MM-DD"T"HH24:MI:SS.USZ') 
			BETWEEN 
			TO_TIMESTAMP('%sT00:00:00.000Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ')
			AND 
			TO_TIMESTAMP('%sT23:59:59.999Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ');`,
			id, start, end,
		)
		err := s.db.Select(&calls, qCalls)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return nil, err
		}
		err = s.db.Select(&email, qEmails)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return nil, err
		}

		data = append(data, models.CallAndEmail{Calls: calls, Emails: email, SiteID: id})

	}

	return &data, nil

}

func (s *Store) GetSingleClient(id int, start string, end string) (*models.CallAndEmail, error) {

	calls := []models.Calls{}
	email := []models.Email{}
	data := models.CallAndEmail{}

	qCalls := fmt.Sprintf(
		`SELECT id, call_id, date, channel_id, source, is_lid, name_type, traffic_type, landing_page, conversations_number, call_status
		FROM calls
		WHERE calls.client_calibri_site_id_fk = %d
		AND 
		TO_TIMESTAMP(calls.date, 'YYYY-MM-DD"T"HH24:MI:SS.USZ') 
		BETWEEN 
		TO_TIMESTAMP('%sT00:00:00.000Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ')
		AND 
		TO_TIMESTAMP('%sT23:59:59.999Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ');`,
		id, start, end,
	)

	qEmails := fmt.Sprintf(
		`SELECT id, email_id, date, source, is_lid, traffic_type, landing_page, conversations_number FROM email 
		WHERE email.client_calibri_site_id_fk = %d
		AND 
		TO_TIMESTAMP(email.date, 'YYYY-MM-DD"T"HH24:MI:SS.USZ') 
		BETWEEN 
		TO_TIMESTAMP('%sT00:00:00.000Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ')
		AND 
		TO_TIMESTAMP('%sT23:59:59.999Z', 'YYYY-MM-DD"T"HH24:MI:SS.USZ');`,
		id, start, end,
	)
	err := s.db.Select(&calls, qCalls)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil, err
	}
	err = s.db.Select(&email, qEmails)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil, err
	}

	data = models.CallAndEmail{Calls: calls, Emails: email, SiteID: id}

	return &data, nil
}

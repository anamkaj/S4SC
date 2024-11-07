package models

import (
	"github.com/lib/pq"
)

type ClientInterface interface {
	GetClientList(status bool) (*[]ClientCalibri, error)
	GetFullDataAllClients(start string, end string) (*[]CallAndEmail, error)
	GetSingleClient(id int, start string, end string) (*CallAndEmail, error)
}

type ClientCalibri struct {
	ID             int            `json:"id" db:"id"`
	SiteID         int            `json:"site_id" db:"site_id"`
	SiteName       string         `json:"sitename" db:"sitename"`
	Domains        string         `json:"domains" db:"domains"`
	Active         string         `json:"active" db:"active"`
	LicenseStart   *string        `json:"license_start,omitempty" db:"license_start"`
	LicenseEnd     *string        `json:"license_end,omitempty" db:"license_end"`
	NotEnoughMoney *bool          `json:"not_enough_money,omitempty" db:"not_enough_money"`
	Number         pq.StringArray `json:"number,omitempty" db:"number"`
}
type Calls struct {
	ID                  int    `json:"id" db:"id"`
	CallID              int    `json:"call_id" db:"call_id"`
	Date                string `json:"date" db:"date"`
	ChannelID           int    `json:"channel_id" db:"channel_id"`
	Source              string `json:"source" db:"source"`
	IsLid               bool   `json:"is_lid" db:"is_lid"`
	NameType            string `json:"name_type" db:"name_type"`
	TrafficType         string `json:"traffic_type" db:"traffic_type"`
	LandingPage         string `json:"landing_page" db:"landing_page"`
	ConversationsNumber int    `json:"conversations_number" db:"conversations_number"`
	CallStatus          string `json:"call_status" db:"call_status"`
}

type Email struct {
	ID                  int    `json:"id" db:"id"`
	EmailID             int    `json:"email_id" db:"email_id"`
	Date                string `json:"date" db:"date"`
	Source              string `json:"source" db:"source"`
	IsLid               bool   `json:"is_lid" db:"is_lid"`
	TrafficType         string `json:"traffic_type" db:"traffic_type"`
	LandingPage         string `json:"landing_page" db:"landing_page"`
	LidLanding          string `json:"lid_landing" db:"lid_landing"`
	ConversationsNumber int    `json:"conversations_number" db:"conversations_number"`
}

type CallAndEmail struct {
	Calls  []Calls `json:"calls"`
	Emails []Email `json:"emails"`
	SiteID int     `json:"site_id"`
}

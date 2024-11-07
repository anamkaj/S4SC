package database

import (
	"calibri/internal/utils"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

const createTableClient = `CREATE TABLE IF NOT EXISTS public.client_calibri (
            id serial4 NOT NULL,
            site_id int8 NOT NULL,
            sitename varchar(255) NULL,
            domains varchar(255) NULL,
            active varchar(255) NOT NULL,
            license_start text NULL,
            license_end text NULL,
            not_enough_money bool NULL,
            CONSTRAINT client_calibri_id_key UNIQUE (id),
            CONSTRAINT client_calibri_pkey PRIMARY KEY (id, site_id),
            CONSTRAINT client_calibri_site_id_key UNIQUE (site_id));`

const createTableCalls = `CREATE TABLE IF NOT EXISTS public.calls (
            id serial4 NOT NULL,
            client_calibri_site_id_fk int8 NOT NULL,
            call_id int8 NOT NULL,
            date varchar(255) NOT NULL,
            channel_id int8 NULL,
            is_lid bool NULL,
            name_type text NULL,
            traffic_type text NULL,
            landing_page text NULL,
            conversations_number int8 NULL,
            call_status varchar(255) NULL,
            source text NULL,
            CONSTRAINT calls_pkey PRIMARY KEY (id),
            CONSTRAINT unique_call_id UNIQUE (call_id),
            CONSTRAINT calls_client_calibri_site_id_fk_fkey FOREIGN KEY (client_calibri_site_id_fk) REFERENCES public.client_calibri(site_id));`

const createTableEmails = `CREATE TABLE IF NOT EXISTS public.email (
            id serial4 NOT NULL,
            client_calibri_site_id_fk int8 NOT NULL,
            email_id int8 NOT NULL,
            date varchar(255) NULL,
            source text NULL,
            is_lid bool NULL,
            traffic_type text NULL,
            landing_page text NULL,
            lid_landing text NULL,
            conversations_number int8 NULL,
            CONSTRAINT email_pkey PRIMARY KEY (id),
            CONSTRAINT unique_email_id UNIQUE (email_id),
            CONSTRAINT email_client_calibri_site_id_fk_fkey FOREIGN KEY (client_calibri_site_id_fk) REFERENCES public.client_calibri(site_id) );`

const createTablePhone = `CREATE TABLE IF NOT EXISTS public.phone (
                id serial4 NOT NULL,
                client_calibri_site_id_fk int8 NOT NULL,
                number _text NULL,
                CONSTRAINT phone_pkey PRIMARY KEY (id),
                CONSTRAINT unique_client_calibri_site_id_fk UNIQUE (client_calibri_site_id_fk),
                CONSTRAINT phone_client_calibri_site_id_fk_fkey FOREIGN KEY (client_calibri_site_id_fk) REFERENCES public.client_calibri(site_id));`

func PostgresConnect() (*sqlx.DB, error) {
	token, err := utils.GetToken()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	pool, err := sqlx.Connect("postgres", token.ClientTable)
	if err != nil {
		log.Fatalln(err)
	}

	for _, query := range []string{
		createTableClient,
		createTableCalls,
		createTableEmails,
		createTablePhone,
	} {
		_, err = pool.Exec(query)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Postgres connected")

	return pool, nil
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupietracker/database"
)

func FetchAPI(url string, s any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return err
	}

	return nil
}

func GetForeignData(artist *database.Artists) error {
	var locations database.Locations
	var dates database.Dates
	var relations database.Relation

	err := FetchAPI(artist.Locations, &locations)
	if err != nil {
		return err
	}
	err = FetchAPI(artist.CongertDates, &dates)
	if err != nil {
		return err
	}
	err = FetchAPI(artist.Relations, &relations)
	if err != nil {
		return err
	}

	artist.Loca = locations
	artist.ConDT = dates
	artist.Rela = relations

	return nil
}

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

	errArtist := make(chan error, 3)

	go func() {
		errArtist <- FetchAPI(artist.Locations, &locations)
	}()

	go func() {
		errArtist <- FetchAPI(artist.CongertDates, &dates)
	}()

	go func() {
		errArtist <- FetchAPI(artist.Relations, &relations)
	}()

	for i := 0; i < 3; i++ {
		if err := <-errArtist; err != nil {
			return err
		}
	}

	artist.Loca = locations
	artist.ConDT = dates
	artist.Rela = relations

	return nil
}

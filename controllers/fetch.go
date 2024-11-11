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
	cha := make(chan error, 3)

	go func() {
		cha <- FetchAPI(artist.Locations, &artist.Loca)
	}()

	go func() {
		cha <- FetchAPI(artist.CongertDates, &artist.ConDT)
	}()

	go func() {
		cha <- FetchAPI(artist.Relations, &artist.Rela)
	}()

	for i := 0; i < 3; i++ {
		if err := <-cha; err != nil {
			return err
		}
	}

	return nil
}

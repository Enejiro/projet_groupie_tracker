package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Members          []string `json:"members"`
	Creation         int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Relation         []string `json:"datesLocations"`
	ConcertDates     []string `json:"dates"`
	ConcertLocations string   `json:"locations"`
}

func SearchArtist(id int) (*Artist, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artist Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

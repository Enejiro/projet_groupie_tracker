package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	Creation     int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	RelationsURL string   `json:"relations"`
	Relations    map[string][]string
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func getOneArtist(id int) (Artist, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return Artist{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Artist{}, err
	}

	var artist Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		return Artist{}, err
	}

	relResp, err := http.Get(artist.RelationsURL)
	if err != nil {
		return Artist{}, err
	}
	defer relResp.Body.Close()

	relBody, err := io.ReadAll(relResp.Body)
	if err != nil {
		return Artist{}, err
	}

	var relation Relation
	err = json.Unmarshal(relBody, &relation)
	if err != nil {
		return Artist{}, err
	}

	// Construire la page
	page := Artist{
		Image:      artist.Image,
		Name:       artist.Name,
		Members:    artist.Members,
		Creation:   artist.Creation,
		FirstAlbum: artist.FirstAlbum,
		Relations:  relation.DatesLocations,
	}

	return page, nil
}

func SearchArtist() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artist []Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}


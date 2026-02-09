package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Artist represents an artist/band with their information
//Relationships contains a map of concert venues (key) and their dates (value)
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

//Relation represents an artist's concert dates and locations
// DatesLocations is a map where the key is the location and the value is the list of dates
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

//getOneArtist retrieves complete information about an artist by their ID
// Makes 2 API requests: one for the artist, another for their relationships (dates/locations)
// Returns the artist with all their data or an error if unsuccessful
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

// Build the complete page with relationships
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

//SearchArtist retrieves the complete list of all artists from the API
// Used to display the home page with the artist grid
// Returns an array of artists or an error if unsuccessful

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


package handlers

import "projet_groupie_tracker/models"

func GetTestArtists() []models.Artist {
    return []models.Artist{
        {
            ID:           1,
            Name:         "Queen",
            Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
            Members:      []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon"},
            CreationDate: 1970,
            FirstAlbum:   "14-12-1973",
        },
        {
            ID:           3,
            Name:         "Nirvana",
            Image:        "https://groupietrackers.herokuapp.com/api/images/nirvana.jpeg",
            Members:      []string{"Kurt Cobain", "Krist Novoselic", "Dave Grohl"},
            CreationDate: 1987,
            FirstAlbum:   "15-06-1989",
        },
    }
}

func GetTestArtistByID(id int) *models.Artist {
    artists := GetTestArtists()
    for _, artist := range artists {
        if artist.ID == id {
            return &artist
        }
    }
    return nil
}
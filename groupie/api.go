package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ArtisteElement struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func Api() Artiste {
	api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err)
	}

	defer api.Body.Close()

	body, err := io.ReadAll(api.Body)
	if err != nil {
		fmt.Print(err)
	}
	artiste, err := UnmarshalArtiste(body)
	if err != nil {
		fmt.Print(err)
	}

	return artiste
}

type Artiste []ArtisteElement

func UnmarshalArtiste(data []byte) (Artiste, error) {
	var r Artiste
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Artiste) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

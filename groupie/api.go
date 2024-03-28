package groupie

import (
	"encoding/json"
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

type ApiLocation struct {
	Index []struct {
		ID        int64    `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type ApiDates struct {
	Index []struct {
		ID    int64    `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

func Api() ([]ArtisteElement, error) {
	api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer api.Body.Close()

	body, err := io.ReadAll(api.Body)
	if err != nil {
		return nil, err
	}
	artiste, err := UnmarshalArtiste(body)
	if err != nil {
		return nil, err
	}

	return artiste, nil
}

type Artiste []ArtisteElement

func UnmarshalArtiste(data []byte) ([]ArtisteElement, error) {
	var r []ArtisteElement
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Artiste) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func LocationApi() (ApiLocation, error) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return ApiLocation{}, err
	}
	defer response.Body.Close()

	var locationData ApiLocation
	err = json.NewDecoder(response.Body).Decode(&locationData)
	if err != nil {
		return ApiLocation{}, err
	}

	return locationData, nil
}

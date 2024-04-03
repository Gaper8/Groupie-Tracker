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

// Here we define the structure of artist elements, including information such as ID, image, and more.

type ApiLocation struct {
	Index []struct {
		ID        int64    `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

// Here we define the structure for the rentals and we store Index all the information from this api.

type ApiDates struct {
	Index []struct {
		ID    int64    `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

// Here we define the structure for the concert dates and we store in Index all the information from this API.

type ApiRelation struct {
	Index []struct {
		ID             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// Here we define the structure for the link between the concert dates and locations of each concert artist and we store all the information in this API in Index.

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

// the Api func allows us to retrieve artist data from a url.
//Then my code reads the raw data then decodes it and converts it into a slice. We implemented error messages for each step.

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

// LocationApi retrieves concert venue data using the URL. I check if my query has no errors.
//Then the data is traversed and converts the data into my struct. My structure contains an index of all concert venues.
// If recovery or decoding fails, the function returns an error.

func DatesApi() (ApiDates, error) {
	response2, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return ApiDates{}, err
	}
	defer response2.Body.Close()

	var datesData ApiDates
	err = json.NewDecoder(response2.Body).Decode(&datesData)
	if err != nil {
		return ApiDates{}, err
	}
	return datesData, nil
}

// DatesApi retrieves concert venue data using the URL.
//I check if my query has no errors.
//Then the data is traversed and converts the data into my struct.
//My structure contains an index of all concert dates. If recovery or decoding fails, the function returns an error.

func RelationApi() (ApiRelation, error) {
	response3, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return ApiRelation{}, err
	}
	defer response3.Body.Close()

	var relationData ApiRelation
	err = json.NewDecoder(response3.Body).Decode(&relationData)
	if err != nil {
		return ApiRelation{}, err
	}
	return relationData, nil
}

// RelationsApi retrieves concert venue data using the URL.
//I check if my query has no errors.
//Then the data is traversed and converts the data into my struct.
//My structure contains an index of the relationships between concert dates and concert locations.
// If recovery or decoding fails, the function returns an error.

package groupie

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

type ApiLocation struct {
    Index []struct {
        ID        int64    `json:"id"`
        Locations []string `json:"locations"`
        Dates     string   `json:"dates"`
    } `json:"index"`
}

func GetCoordinatesFromAPI(address string) (string, error) {
 
    response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations?address=" + url.QueryEscape(address))
    if err != nil {
        return "", err
    }
    defer response.Body.Close()

    
    var locationData ApiLocation
    err = json.NewDecoder(response.Body).Decode(&locationData)
    if err != nil {
        return "", err
    }


    coordinates := ""
    if len(locationData.Index) > 0 && len(locationData.Index[0].Locations) > 0 {
        coordinates = locationData.Index[0].Locations[0]
    }

    return coordinates, nil
}

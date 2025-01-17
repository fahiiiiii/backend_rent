package services

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    beego "github.com/beego/beego/v2/server/web"
    "backend_rental/models"
    "path/filepath"
	"os"
)

type BookingService struct{}

type APIResponse struct {
    Data []struct {
        CityName string `json:"city_name"`
        DestID   string `json:"dest_id"`
        Country  string `json:"country"`
    } `json:"data"`
}

func (s *BookingService) FetchCities(query string) ([]models.City, error) {
    // Get configuration
    apiURL, _ := beego.AppConfig.String("rapidapi.url")
    apiKey, _ := beego.AppConfig.String("rapidapi.key")
    apiHost, _ := beego.AppConfig.String("rapidapi.host")

    client := &http.Client{}
    req, err := http.NewRequest("GET", fmt.Sprintf("%s?query=%s", apiURL, query), nil)
    if err != nil {
        return nil, err
    }

    req.Header.Add("x-rapidapi-key", apiKey)
    req.Header.Add("x-rapidapi-host", apiHost)

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var apiResp APIResponse
    if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
        return nil, err
    }

    cities := make([]models.City, len(apiResp.Data))
    for i, item := range apiResp.Data {
        cities[i] = models.City{
            CityName: item.CityName,
            DestID:   item.DestID,
            Country:  item.Country,
        }
    }

    // Store in JSON file
    if err := s.storeCitiesInFile(cities); err != nil {
        return nil, err
    }

    return cities, nil
}

func (s *BookingService) storeCitiesInFile(cities []models.City) error {
    storagePath, _ := beego.AppConfig.String("storage.path")
    
    // Ensure directory exists
    dir := filepath.Dir(storagePath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }

    data, err := json.MarshalIndent(cities, "", "    ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(storagePath, data, 0644)
}

package bgateway

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "time"
    "log"
    "fmt"
    "net/url"

    "brawlstarsclub.com/m/v2/entity"
)

const (
    clubTag = "2GYU8RQYQ"
    getClubURL = "https://api.brawlstars.com/v1/clubs/%23"
    getPlayerURL = "https://api.brawlstars.com/v1/players/"
)

// GetClub ...
func GetClub() (*entity.GetClubResponse, error) {

    timeout := time.Duration(10 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }

    URL := fmt.Sprintf("%v%v", getClubURL, clubTag)
    request, err := http.NewRequest("GET", URL, nil)
    request.Header.Add("Authorization", entity.BrawlToken)

    if err != nil {
        log.Fatal("[GetClub] Failed to create request", err)
        return nil, err
    }

    resp, err := client.Do(request)
    if err != nil {
        log.Fatal("[GetClub] Failed to execute request", err)
        return nil, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("[GetClub] Failed to read response", err)
        return nil, err
    }

    var clubResp entity.GetClubResponse
    if err := json.Unmarshal(body, &clubResp); err != nil {
        log.Fatal("[GetClub] Failed to unmarshal response", err)
        return nil, err
    }

    return &clubResp, nil
}

func GetPlayer(tag string) (*entity.GetPlayerResponse, error) {
    
    timeout := time.Duration(10 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }

    URL := fmt.Sprintf("%v%v", getPlayerURL, url.QueryEscape(tag))
    request, err := http.NewRequest("GET", URL, nil)
    request.Header.Add("Authorization", entity.BrawlToken)

    if err != nil {
        log.Fatal("[GetPlayer] Failed to create request", err)
        return nil, err
    }

    resp, err := client.Do(request)
    if err != nil {
        log.Fatal("[GetPlayer] Failed to execute request", err)
        return nil, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("[GetPlayer] Failed to read response", err)
        return nil, err
    }

    var player entity.GetPlayerResponse
    if err := json.Unmarshal(body, &player); err != nil {
        log.Fatal("[GetPlayer] Failed to unmarshal response", err)
        return nil, err
    }

    return &player, nil
}
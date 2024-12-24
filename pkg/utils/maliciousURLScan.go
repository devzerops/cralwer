package utils

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

const googleSafeBrowsingAPI = "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=YOUR_API_KEY"

// SafeBrowsingRequest represents the request payload for Google Safe Browsing API
type SafeBrowsingRequest struct {
    Client      ClientInfo      `json:"client"`
    ThreatInfo  ThreatInfo      `json:"threatInfo"`
}

// ClientInfo represents the client information for Google Safe Browsing API
type ClientInfo struct {
    ClientId      string `json:"clientId"`
    ClientVersion string `json:"clientVersion"`
}

// ThreatInfo represents the threat information for Google Safe Browsing API
type ThreatInfo struct {
    ThreatTypes      []string `json:"threatTypes"`
    PlatformTypes    []string `json:"platformTypes"`
    ThreatEntryTypes []string `json:"threatEntryTypes"`
    ThreatEntries    []URL    `json:"threatEntries"`
}

// URL represents a URL entry for Google Safe Browsing API
type URL struct {
    URL string `json:"url"`
}

// SafeBrowsingResponse represents the response from Google Safe Browsing API
type SafeBrowsingResponse struct {
    Matches []Match `json:"matches"`
}

// Match represents a match in the response from Google Safe Browsing API
type Match struct {
    ThreatType string `json:"threatType"`
    PlatformType string `json:"platformType"`
    ThreatEntryType string `json:"threatEntryType"`
    Threat URL `json:"threat"`
}

// IsMaliciousURL checks if the URL is malicious using Google Safe Browsing API
func IsMaliciousURL(url string) (bool, error) {
    requestBody := SafeBrowsingRequest{
        Client: ClientInfo{
            ClientId:      "your-client-id",
            ClientVersion: "1.0",
        },
        ThreatInfo: ThreatInfo{
            ThreatTypes:      []string{"MALWARE", "SOCIAL_ENGINEERING"},
            PlatformTypes:    []string{"ANY_PLATFORM"},
            ThreatEntryTypes: []string{"URL"},
            ThreatEntries:    []URL{{URL: url}},
        },
    }

    requestBytes, err := json.Marshal(requestBody)
    if err != nil {
        return false, fmt.Errorf("failed to marshal request: %v", err)
    }

    resp, err := http.Post(googleSafeBrowsingAPI, "application/json", strings.NewReader(string(requestBytes)))
    if err != nil {
        return false, fmt.Errorf("failed to send request: %v", err)
    }
    defer resp.Body.Close()

    var response SafeBrowsingResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return false, fmt.Errorf("failed to decode response: %v", err)
    }

    return len(response.Matches) > 0, nil
}
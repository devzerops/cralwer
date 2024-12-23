package handlers

import (
    "encoding/json"
    "net"
    "net/http"
    "github.com/gocolly/colly"
)

func CollectHandler(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    if url == "" {
        http.Error(w, "URL parameter is missing", http.StatusBadRequest)
        return
    }

    c := colly.NewCollector()
    var result []byte
    c.OnResponse(func(response *colly.Response) {
        result = response.Body
    })

    err := c.Visit(url)
    if err != nil {
        http.Error(w, "Failed to collect the page", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(result)
}

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
    host := r.URL.Query().Get("host")
    if host == "" {
        http.Error(w, "Host parameter is missing", http.StatusBadRequest)
        return
    }

    ips, err := net.LookupIP(host)
    if err != nil {
        http.Error(w, "Failed to resolve the host", http.StatusInternalServerError)
        return
    }

    ipStrings := make([]string, len(ips))
    for i, ip := range ips {
        ipStrings[i] = ip.String()
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ipStrings)
}
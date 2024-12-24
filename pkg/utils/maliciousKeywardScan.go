package utils

import (
    "strings"
)

// IsMaliciousURLByKeyword checks if the URL contains any malicious keywords
func IsMaliciousURLByKeyword(url string, keywords []string) bool {
    for _, keyword := range keywords {
        if strings.Contains(url, keyword) {
            return true
        }
    }
    return false
}
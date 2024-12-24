package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "runtime"
    "distributed-crawler/m/pkg/utils"
    "github.com/joho/godotenv"
)

func InfoHandler(envPath string) http.Handler {
    return utils.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // .env 파일 로드
        err := godotenv.Load(envPath)
        if (err != nil) {
            http.Error(w, "Failed to load .env file", http.StatusInternalServerError)
            return
        }

        // .env 파일에서 platform 값 읽기
        platform := os.Getenv("platform")
        if platform == "" {
            http.Error(w, "Platform not specified in .env file", http.StatusBadRequest)
            return
        }

        var clientInfo interface{}
        hostname, _ := os.Hostname()
        ip := utils.GetLocalIP()
        osType := runtime.GOOS

        switch platform {
        case "k8s":
            podName := os.Getenv("POD_NAME")
            namespace := os.Getenv("NAMESPACE")
            nodeName := os.Getenv("NODE_NAME")
            clusterName := os.Getenv("CLUSTER_NAME")
            clientInfo = utils.NewK8SClientInfo(ip, osType, hostname, podName, namespace, nodeName, clusterName)
        case "docker":
            containerID := os.Getenv("HOSTNAME")
            clientInfo = utils.NewDockerClientInfo(ip, osType, hostname, containerID)
        case "local":
            clientInfo = utils.NewClientInfo(ip, osType, hostname)
        default:
            http.Error(w, "Unsupported platform", http.StatusBadRequest)
            return
        }

        clientInfoJSON, err := json.Marshal(clientInfo)
        if err != nil {
            http.Error(w, "Failed to marshal client info", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(clientInfoJSON)
    }))
}

func InfoIPHandler() http.Handler {
    return utils.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := utils.GetLocalIP()
        ipJSON, err := json.Marshal(map[string]string{"ip": ip})
        if err != nil {
            http.Error(w, "Failed to marshal IP", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(ipJSON)
    }))
}
package utils

import (
	"distributed-crawler/m/pkg/models"
	"net"
)

func NewClientInfo(ip, os, hostname string) *models.ClientInfo {
	return &models.ClientInfo{
		IP:       ip,
		OS:       os,
		Hostname: hostname,
	}
}

func NewK8SClientInfo(ip, os, hostname, podName, namespace, nodeName, clusterName string) *models.K8SClientInfo {
	return &models.K8SClientInfo{
		ClientInfo: models.ClientInfo{
			IP:       ip,
			OS:       os,
			Hostname: hostname,
		},
		PodName:     podName,
		Namespace:   namespace,
		NodeName:    nodeName,
		ClusterName: clusterName,
	}
}

func NewDockerClientInfo(ip, os, hostname, containerID string) *models.DockerClientInfo {
	return &models.DockerClientInfo{
		ClientInfo: models.ClientInfo{
			IP:       ip,
			OS:       os,
			Hostname: hostname,
		},
		ContainerID: containerID,
	}
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "unknown"
}
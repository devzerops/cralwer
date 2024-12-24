package models

type ClientInfo struct {
	IP       string `json:"ip"`
	OS       string `json:"os"`
	Hostname string `json:"hostname"`
}

type K8SClientInfo struct {
	ClientInfo
	PodName     string `json:"pod_name"`
	Namespace   string `json:"namespace"`
	NodeName    string `json:"node_name"`
	ClusterName string `json:"cluster_name"`
}

type DockerClientInfo struct {
	ClientInfo
	ContainerID string `json:"container_id"`
}
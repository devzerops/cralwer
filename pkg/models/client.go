package models

type ClientInfo struct {
	IP       string `json:"ip"`
	OS       string `json:"os"`
	Hostname string `json:"hostname"`
}

type Worker struct {
	IP        string
	Port      string
	//stopHeartbeat chan bool
}

type WorkerConfig struct {
	Worker
	CassandraIP      string
	CassandraKeyspace string
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

type ClientInfoProvider interface {
	GetClientInfo() ClientInfo
}

func (c *ClientInfo) GetClientInfo() ClientInfo {
	return *c
}

func (k *K8SClientInfo) GetClientInfo() ClientInfo {
	return k.ClientInfo
}

func (d *DockerClientInfo) GetClientInfo() ClientInfo {
	return d.ClientInfo
}

package models

type SystemEmail struct{
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	SSL bool `json:"ssl,omitempty" yaml:"ssl,omitempty"`
	TLS bool `json:"tls,omitempty" yaml:"tls,omitempty"`
}
func (e SystemEmail) ToMap() map[string]interface{}{
	m := make(map[string]interface{})
	m["type"] = e.Type
	m["host"] = e.Host
	m["port"] = e.Port
	m["username"] = e.Username
	m["password"] = e.Password
	m["ssl"] = e.SSL
	m["tls"] = e.TLS
	return m
}

type SystemCluster struct {
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

func (c SystemCluster) ToMap() map[string]interface{}{
	m := make(map[string]interface{})
	m["address"] = c.Address
	m["port"] = c.Port
	m["scheme"] = c.Scheme
	return m
}

type SystemConfig struct {
	Remote bool `json:"remote,omitempty" yaml:"remote,omitempty"`
	Email SystemEmail `json:"email,omitempty" yaml:"email,omitempty"`
	Cluster SystemCluster `json:"cluster,omitempty" yaml:"cluster,omitempty"`
}

func (c SystemConfig) ToMap() map[string]interface{}{
	m := make(map[string]interface{})
	m["remote"] = c.Remote
	m["email"] = c.Email.ToMap()
	m["cluster"] = c.Cluster.ToMap()
	return m
}

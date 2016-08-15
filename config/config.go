package config

var Rancher RancherConfiguration

type RancherConfiguration struct {
	Server    string
	AccessKey string
	SecretKey string
}

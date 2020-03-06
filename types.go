package main

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

// TODO : SWAPAN Add YAML tags.
type NamespaceCrp struct {
	Kind       string   `yaml:"kind", json:"kind"`
	APIVersion string   `yaml:"apiVersion", json:"apiVersion`
	Spec       Spec     `json:"spec"`
	Metadata   Metadata `json:"metadata"`
}
type Metadata struct {
	ClusterID string 		`yaml:"clusterId", json:"clusterId"`
	Name      string 		`yaml:"name", json:"name"`
}
type Labels struct {
	Key2 string `json:"key2"`
	Key1 string `json:"key1"`
}
type Namespace struct {
	Labels Labels `json:"labels"`
	Name   string `json:"name"`
}
type Max struct {
	Memory       string `json:"memory"`
	CPU          int    `json:"cpu"`
	NvidiaComGpu int    `json:"nvidia.com/gpu"`
}
type Min struct {
	Memory       string `json:"memory"`
	CPU          int    `json:"cpu"`
	NvidiaComGpu int    `json:"nvidia.com/gpu"`
}
type Quota struct {
	Max Max `json:"max"`
	Min Min `json:"min"`
}
type Serviceaccounts struct {
	Labels map[string]string  `json:"labels"`
	Name   string  `json:"name"`
	//Roles  []Roles `yaml: "roles", json:"roles"`
	Roles  []rbacv1.Role `yaml: "roles", json:"roles"`
}

type Permissions struct {
	Serviceaccounts []Serviceaccounts `json:"serviceaccounts"`
}
type Spec struct {
	//Namespace   Namespace     `json:"namespace"`
	//Quota       Quota         `json:"quota"`
	//Resources   interface{}   `json:"resources"`
	Permissions []Permissions `json:"permissions"`
}
/*
type Metadata struct {
	ClusterID string `json:"clusterId"`
	Name      string `json:"name"`
}*/
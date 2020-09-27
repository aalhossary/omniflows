package models

import (
	"encoding/json"
	"fmt"
)

type Parameter struct {
	DisplayName string      `json:"displayName,omitempty" yaml:"displayname,omitempty"`
	Name        string      `json:"name" yaml:"name"`
	Type        string      `json:"type,omitempty" yaml:"type,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Value       interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

func (p *Parameter) GetParamValue() string {
	return fmt.Sprintf("%v", p.Value)
}

type Reference struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Website     string `json:"website,omitempty" yaml:"website,omitempty"`
}

type Maintainer struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	FullName string `json:"fullname,omitempty" yaml:"fullname,omitempty"`
	Email    string `json:"email,omitempty" yaml:"email,omitempty"`
}

type Tool struct {
	Type        string   `json:"type,omitempty" yaml:"type,omitempty"`
	URL         string   `json:"url,omitempty" yaml:"url,omitempty"`
	Depends     string   `json:"depends,omitempty" yaml:"depends,omitempty"`
	ID          string   `json:"id,omitempty" yaml:"id,omitempty"`
	Order       int      `json:"order,omitempty" yaml:"order,omitempty"`
	BioflowId   string   `json:"bioflowId,omitempty" yaml:"bioflowId,omitempty"`
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Discussions []string `json:"discussions,omitempty" yaml:"discussions,omitempty"`
	Website     string   `json:"website,omitempty" yaml:"website,omitempty"`
	Version     string   `json:"version,omitempty" yaml:"version,omitempty"`
	Icon        string   `json:"icon,omitempty" yaml:"icon,omitempty"`
	Shadow      bool     `json:"shadow,omitempty" yaml:"shadow,omitempty"`
	Maintainer   *Maintainer  `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`
	References   []Reference  `json:"references,omitempty" yaml:"references,omitempty"`
	Inputs       []Parameter  `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Config       []Parameter  `json:"config,omitempty" yaml:"config,omitempty"`
	Outputs      []Parameter  `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Command      Scriptable   `json:"command" yaml:"command"`
	Dependencies []string     `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Deprecated   bool         `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Conditions   []Scriptable `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Scripts      []Script     `json:"scripts,omitempty" yaml:"scripts,omitempty"`
}

func (t *Tool) ToJson() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		panic("Unable to Convert the current tool into JSON.")
	}
	return string(bytes)
}

package io

import (
  "encoding/json"
  _ "fmt"
  "github.com/v2ray/v2ray-core"
)

type JsonVUser struct {
  id string `json:"id"`
  email string `json:"email"`
}

type JsonVConfig struct {
  RunAs string `json:"runas"`
  Port uint8 `json:"port"`
  Clients []JsonVUser `json:"users"`
  Protocol string `json:"protocol"`
}

type JsonVConfigUnmarshaller struct {

}

func StringToVUser(id string) (u core.VUser, err error) {
  return
}

func (*JsonVConfigUnmarshaller) Unmarshall(data []byte) (*core.VConfig, error) {
  var jsonConfig JsonVConfig
  err := json.Unmarshal(data, &jsonConfig)
  if err != nil {
    return nil, err
  }
  var vconfig = new(core.VConfig)
  vconfig.RunAs = core.VUser{}
  return vconfig, nil
}
package model

type product struct {
 ProductName string `json:"ProductName"`
 Available string `json:"Available"`
 Seen string `json:"Seen"`
 Price int `json:"Price"`
 Dimension string `json:"Dimension"`
 Colours []string `json:"Colours"`
 Material string `json:"Material"`
 Image string `json:"Image"`
}

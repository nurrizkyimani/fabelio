package model

type product struct {
 ProductName string `json:"ProductName"`
 Available string `json:"Available"`
 Dimension string `json:"Dimension"`
 Colours []string `json:"Colours"`
 Material string `json:"Material"`
 Image string `json:"Image"`
}

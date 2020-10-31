package model

type Product struct {
 ProductName string `json:"ProductName"`
 Available bool `json:"Available"`
 Seen bool `json:"Seen"`
 Price int `json:"Price"`
 Dimension string `json:"Dimension"`
 Colours []string `json:"Colours"`
 Material string `json:"Material"`
 Image string `json:"Image"`
}


type ProductSeenUpdate struct  {
	ObjectID string `json:"ObjectID"`
	Seen bool`json:"Seen"`
}

//Response is a xxxx
type Respoonse struct {
	Params string `json:"Params"`
	Hits []Hit `json:"Hits"`
}


type Hit struct {
	ProductName string `json:"ProductName"`
	ObjectID string `json:"ObjectID"`
}
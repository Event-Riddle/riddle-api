package server

/************* Filter Config **************/
type Filter struct {
	Name           string  `json:"name"`
	LowerTrashhold float32 `json:"threshold"`
	TopID          string  `json:"tid"`
	BottomID       string  `json:"bid"`
	UpperTrashhold float32 `json:"threshold-top"`
	Unit           string  `json:"unit"`
	Active         bool    `json:"active"`
}

/************ User Config *******************/
type User struct {
	Toolbar []Items `json:"toolbar"`
	Chain   []Items `json:"chain"`
}

type Items struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

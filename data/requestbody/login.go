package requestbody

type Login struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

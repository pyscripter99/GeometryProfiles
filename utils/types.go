package utils

type Profile struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type DotProfile struct {
	Current  string    `json:"current"`
	Profiles []Profile `json:"profiles"`
}

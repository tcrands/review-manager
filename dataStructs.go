package main

type ReviewData struct {
	Time string `json:"time"`
	Name string `json:"name"`
	Solicited string `json:"solicited"`
    Device string `json:"device"`
    Length int `json:"length"`
	Stars float64 `json:"stars"`
}

type AdvisorData struct {
	Name string `json:"name"`
	Reviews []string `json:"reviews"`
	Score float64 `json:"score"`
	Devices []string `json:"devices"`
	Stars float64 `json:"stars"`
	Error string `json:"error"`
}

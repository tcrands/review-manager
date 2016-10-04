package main

type ReviewData struct {
	Time string `json:"time"`
	Name string `json:"name"`
	Solicited string `json:"solicited"`
    Device string `json:"device"`
    Length int `json:"length"`
	Stars int `json:"stars"`
}
type ReviewDataSet []ReviewData

type AdvisorData struct {
	Name string
	Reviews []string
	Score int
	Devices []string
	Stars int
}
type AdvisorDataSet []AdvisorData

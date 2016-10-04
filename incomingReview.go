package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
)

func incomingReview(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		panic(err)
	}

    reviewArray := strings.Split(string(body[:]), ",")

    if len(reviewArray) == 6 {

        wordCount, _ := strconv.Atoi(strings.Replace(reviewArray[4], "words", "", 1))

        reviewData := &ReviewData{
            Time: reviewArray[0],
            Name: reviewArray[1],
            Solicited: reviewArray[2],
            Device: reviewArray[3],
            Length: wordCount,
            Stars: len(reviewArray[5]),
        }

        advisorData, exists := getAdvisor(reviewData.Name)

        if exists {
            processedAdvisorData := processCurrentAdvisor(advisorData, reviewData)
            updateDatabase(processedAdvisorData)
        } else {
            processedAdvisorData := processNewAdvisor(advisorData, reviewData)
            updateDatabase(processedAdvisorData)
        }

    } else {
        fmt.Println("Could not read review summary data")
    }

}

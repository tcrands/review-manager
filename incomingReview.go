package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
    "encoding/json"
    "reflect"
)

func incomingReview(w http.ResponseWriter, r *http.Request) {
    processedAdvisorData := &AdvisorData{}
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		panic(err)
	}

    reviewArray := strings.Split(string(body[:]), ",")

    if len(reviewArray) == 6 {

        reviewData := processReqBody(reviewArray)

        advisorData, exists := getAdvisor(reviewData.Name)

        if exists {
            processedAdvisorData = processCurrentAdvisor(advisorData, reviewData)
        } else {
            processedAdvisorData = processNewAdvisor(advisorData, reviewData)
        }

        fmt.Println(processedAdvisorData)

        updateDatabase(processedAdvisorData)
        fmt.Println(processOutput(processedAdvisorData))

    } else {
        fmt.Println("Could not read review summary data")
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if reflect.DeepEqual(processedAdvisorData, new(AdvisorData)) == true {
        processedAdvisorData.Error = "Could not read review summary data"
    }

    if err := json.NewEncoder(w).Encode(processedAdvisorData); err != nil {
        	panic(err)
    	}
}

func processReqBody(reviewArray []string) *ReviewData {
    wordCount, _ := strconv.Atoi(strings.Replace(reviewArray[4], "words", "", 1))

    reviewData := &ReviewData{
        Time: reviewArray[0],
        Name: reviewArray[1],
        Solicited: reviewArray[2],
        Device: reviewArray[3],
        Length: wordCount,
        Stars: float64(len(reviewArray[5])),
    }

    return reviewData
}

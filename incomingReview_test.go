package main

import (
    "testing"
    "reflect"
    "net/http"
    "strings"
    "net/http/httptest"
    "encoding/json"
)

func TestAdvisorProcesses_incomingReview_newAdvisor(t *testing.T) {
    dbMap = "testDbMap"
    defer clearTestDatabase("Jon")

    advisorData := "12thJuly12:04,Jon,solicited,LB3‐TYU,50words,*****"

    expectedAdvisorData := &AdvisorData{
        Name: "Jon",
        Score: 100,
        Devices: []string{"LB3‐TYU"},
        Stars: 5,
        Error: "",
    }
    expectedReviewData := &ReviewData{
        Time: "12thJuly12:04",
        Name: "Jon",
        Solicited: "solicited",
        Device: "LB3‐TYU",
        Length: 50,
        Stars: 5,
    }
    review, _ := json.Marshal(expectedReviewData)
    expectedAdvisorData.Reviews = append(expectedAdvisorData.Reviews, string(review))

    reader := strings.NewReader(advisorData)

    req, err := http.NewRequest("POST", "/review", reader)
    if err != nil {
        t.Fatal(err)
    }

    resp := httptest.NewRecorder()
    handler := http.HandlerFunc(incomingReview)
    handler.ServeHTTP(resp, req)

    data := new(AdvisorData)

    json.Unmarshal([]byte(resp.Body.String()), data)

    if err != nil {
        t.Error(err)
    }

    if resp.Code != 200 {
        t.Fatalf("Expected %s, got %s", 200, resp.Code)
    }

    if expectedAdvisorData == data {
        t.Fatalf("Expected %s, got %s", expectedAdvisorData, data)
    }
}


func TestAdvisorProcesses_processReqBody(t *testing.T) {

        expectedReviewData := &ReviewData{
            Time: "12thJuly12:04",
            Name: "Jon",
            Solicited: "solicited",
            Device: "LB3‐TYU",
            Length: 50,
            Stars: 5,
        }
        reviewArray := []string{"12thJuly12:04","Jon","solicited","LB3‐TYU","50words","*****"}

        result := processReqBody(reviewArray)

        if reflect.DeepEqual(expectedReviewData, result) == false {
                t.Fatalf("Expected %s, got %s", expectedReviewData, result)
        }

}

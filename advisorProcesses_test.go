package main

import (
    "testing"
    "encoding/json"
    "reflect"
    "github.com/boltdb/bolt"
)

func TestAdvisorProcesses_getAdvisor_advisorNotPresent(t *testing.T) {
        dbMap = "testDbMap"
        defer clearTestDatabase("Jon")

        _, returnedAdvisor := getAdvisor("Jon")

        if returnedAdvisor {
                t.Fatalf("Expected %s, got %s", false, returnedAdvisor)
        }
}

func TestAdvisorProcesses_getAdvisor_advisorPresent(t *testing.T) {
        dbMap = "testDbMap"
        db, _ := setupDatabase()
        defer clearTestDatabase("Jon")

        advisorData := &AdvisorData{
            Name: "Jon",
            Score: 100,
        }
        advisorName := advisorData.Name
        advisor, _ := json.Marshal(advisorData)

        stringedAdvisor := string(advisor)

        db.Update(func(tx *bolt.Tx) error {
            b, _ := tx.CreateBucketIfNotExists([]byte(advisorName))
            return b.Put([]byte(advisorName), []byte(stringedAdvisor))
        })
        db.Close()
        
        _, returnedAdvisor := getAdvisor("Jon")

        if returnedAdvisor != true{
                t.Fatalf("Expected %s, got %s", true, returnedAdvisor)
        }
}

func TestAdvisorProcess_processNewAdvisor(t *testing.T) {

        advisorData := &AdvisorData{
            Name: "Jon",
            Score: 100,
        }

        reviewData := &ReviewData{
            Time: "2ndSeptember10:04",
            Name: "Jon",
            Solicited: "solicited",
            Device: "AN9‐IPK",
            Length: 50,
            Stars: float64(len("*****")),
        }

        review, _ := json.Marshal(reviewData)
        expectedAdvisorData := &AdvisorData{
            Name: "Jon",
            Score: 100,
            Stars: 5,
            Devices: []string {"AN9‐IPK"},
        }
        expectedAdvisorData.Reviews = append(advisorData.Reviews, string(review))

        result := processNewAdvisor(advisorData,reviewData)

        if reflect.DeepEqual(expectedAdvisorData, result) == false {
                t.Fatalf("Expected %s, got %s", expectedAdvisorData, result)
        }
}

func TestAdvisorProcess_processCurrentAdvisor(t *testing.T) {

        currentReviewData := &ReviewData{
            Time: "2ndSeptember10:04",
            Name: "Jon",
            Solicited: "solicited",
            Device: "AN9‐IPU",
            Length: 50,
            Stars: float64(len("**")),
        }

        review, _ := json.Marshal(currentReviewData)
        currentAdvisorData := &AdvisorData{
            Name: "Jon",
            Score: 100,
            Stars: 5,
            Devices: []string {"AN9‐IPU"},
        }
        currentAdvisorData.Reviews = append(currentAdvisorData.Reviews, string(review))

        newReviewData := &ReviewData{
            Time: "3ndSeptember10:04",
            Name: "Jon",
            Solicited: "unsolicited",
            Device: "AN9‐IPK",
            Length: 50,
            Stars: float64(len("**")),
        }

        result := processCurrentAdvisor(currentAdvisorData,newReviewData)

        newReview, _ := json.Marshal(newReviewData)
        currentAdvisorData.Reviews = append(currentAdvisorData.Reviews, string(newReview))

        if reflect.DeepEqual(currentAdvisorData, result) == false {
                t.Fatalf("Expected %s, got %s", currentAdvisorData, result)
        }
}

func TestAdvisorProcess_processOutput_info(t *testing.T) {

    advisorData := &AdvisorData{
        Name: "Jon",
        Score: 100,
    }

    expectedStr := "Info: Jon has a trusted review score of 100.0"

    result := processOutput(advisorData)

    if expectedStr != result {
        t.Fatalf("Expected %s, got %s", expectedStr, result)
    }
}

func TestAdvisorProcess_processOutput_warning(t *testing.T) {

    advisorData := &AdvisorData{
        Name: "Jon",
        Score: 60,
    }

    expectedStr := "Warning: Jon has a trusted review score of 60.0"

    result := processOutput(advisorData)

    if expectedStr != result {
        t.Fatalf("Expected %s, got %s", expectedStr, result)
    }
}

func TestAdvisorProcess_processOutput_alert(t *testing.T) {

    advisorData := &AdvisorData{
        Name: "Jon",
        Score: 10,
    }

    expectedStr := "Alert: Jon has been deactived due to a low trusted review score"

    result := processOutput(advisorData)

    if expectedStr != result {
        t.Fatalf("Expected %s, got %s", expectedStr, result)
    }
}

func TestAdvisorProcess_processStarRating_singleReview(t *testing.T) {

    advisorData := &AdvisorData{
        Stars: 5,
        Reviews: []string {"1"},
    }

    reviewData := &ReviewData{
        Stars: 4,
    }

    expected := 4.5

    advisorData.processStarRating(reviewData)

    if expected != advisorData.Stars {
        t.Fatalf("Expected %s, got %s", expected, advisorData.Stars)
    }
}

func TestAdvisorProcess_processStarRating_multipleReviews(t *testing.T) {

    advisorData := &AdvisorData{
        Stars: 5,
        Reviews: []string {"1","2","3"},
    }

    reviewData := &ReviewData{
        Stars: 4,
    }

    expected := 4.75

    advisorData.processStarRating(reviewData)

    if expected != advisorData.Stars {
        t.Fatalf("Expected %s, got %s", expected, advisorData.Stars)
    }
}

func TestAdvisorProcess_finaliseData_noName(t *testing.T) {

    advisorData := &AdvisorData{}

    reviewData := &ReviewData{
        Name: "Jon",
    }

    expected := "Jon"

    advisorData.finaliseData(reviewData)

    if expected != advisorData.Name {
        t.Fatalf("Expected %s, got %s", expected, advisorData.Name)
    }
}

func TestAdvisorProcess_finaliseData_overScore(t *testing.T) {

    advisorData := &AdvisorData{
        Name: "Jon",
        Score: 150,
    }

    reviewData := &ReviewData{}

    expected := 100.0

    advisorData.finaliseData(reviewData)

    if expected != advisorData.Score {
        t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
    }
}

func TestAdvisorProcess_finaliseData_underScore(t *testing.T) {

    advisorData := &AdvisorData{
        Name: "Jon",
        Score: -100,
    }

    reviewData := &ReviewData{}

    expected := 0.0

    advisorData.finaliseData(reviewData)

    if expected != advisorData.Score {
        t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
    }
}

// HELPER FUNCTIONS ------------------------

func clearTestDatabase(name string) {
    db, _ := setupDatabase()
    defer db.Close()

    db.Update(func(tx *bolt.Tx) error {
        tx.DeleteBucket([]byte(name))
        return nil
    })
}

package main

import (
    "github.com/boltdb/bolt"
    "encoding/json"
    "strconv"
)
func getAdvisor(name string) (*AdvisorData, bool){
    db, _ := setupDatabase()
    defer db.Close()

    advisor := false

    advisorData := &AdvisorData{
        Name: name,
        Score: 100,
    }

    db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(name))
        if bucket == nil {
            advisor = false
        } else {
            val := bucket.Get([]byte(name))
            json.Unmarshal(val, &advisorData)
            advisor = true
        }
        return nil
    })

    return advisorData, advisor

}

func processNewAdvisor(advisorData *AdvisorData, reviewData *ReviewData) *AdvisorData {
    advisorData.Stars = reviewData.Stars
    reviewData.reviewStarCheck(advisorData)
    reviewData.reviewLengthCheck(advisorData)
    reviewData.reviewSolicitedCheck(advisorData)
    advisorData.Devices = append(advisorData.Devices, reviewData.Device)
    advisorData.finaliseData(reviewData)

    return advisorData
}

func processCurrentAdvisor(advisorData *AdvisorData, reviewData *ReviewData) *AdvisorData {
    reviewData.reviewStarCheck(advisorData)
    reviewData.reviewLengthCheck(advisorData)
    reviewData.reviewSolicitedCheck(advisorData)
    reviewData.reviewBurstCheck(advisorData)
    reviewData.reviewDeviceCheck(advisorData)

    advisorData.processStarRating(reviewData)
    advisorData.finaliseData(reviewData)

    return advisorData
}

func processOutput(advisorData *AdvisorData) string {
    score := strconv.FormatFloat(advisorData.Score, 'f', 1, 64)
    if advisorData.Score >= 70 {
        return "Info: " + advisorData.Name + " has a trusted review score of " + score
    } else if advisorData.Score >= 50 {
        return "Warning: " + advisorData.Name + " has a trusted review score of " + score
    } else {
        return "Alert: " + advisorData.Name + " has been deactived due to a low trusted review score"
    }
}

func (advisorData *AdvisorData) processStarRating(reviewData *ReviewData) {
    reviewCount := len(advisorData.Reviews)
    starsCount := advisorData.Stars * float64(reviewCount)
    advisorData.Stars = (starsCount + reviewData.Stars)/(float64(reviewCount+1))
}

func (advisorData *AdvisorData) finaliseData(reviewData *ReviewData) {
    review, _ := json.Marshal(reviewData)
    advisorData.Reviews = append(advisorData.Reviews, string(review))

    if advisorData.Name == "" {
        advisorData.Name = reviewData.Name
    }
    if advisorData.Score > 100 {
        advisorData.Score = 100
    } else if advisorData.Score < 0 {
        advisorData.Score = 0
    }
}

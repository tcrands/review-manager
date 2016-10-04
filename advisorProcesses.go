package main

import (
    "github.com/boltdb/bolt"
    "encoding/json"
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
    reviewData.reviewStarCheck(advisorData)
    reviewData.reviewLengthCheck(advisorData)
    reviewData.reviewSolicitedCheck(advisorData)

    advisorData.Devices = append(advisorData.Devices, reviewData.Device)
    review, _ := json.Marshal(reviewData)
    advisorData.Reviews = append(advisorData.Reviews, string(review))
    advisorData.Name = reviewData.Name

    return advisorData
}

func processCurrentAdvisor(advisorData *AdvisorData, reviewData *ReviewData) *AdvisorData {
    reviewData.reviewStarCheck(advisorData)
    reviewData.reviewLengthCheck(advisorData)
    reviewData.reviewSolicitedCheck(advisorData)
    // reviewData.reviewBurstCheck(advisorData)
    // reviewData.reviewDeviceCheck(advisorData)

    advisorData.Devices = append(advisorData.Devices, reviewData.Device)
    review, _ := json.Marshal(reviewData)
    advisorData.Reviews = append(advisorData.Reviews, string(review))
    advisorData.Name = reviewData.Name

    return advisorData
}

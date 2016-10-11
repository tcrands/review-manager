package main

import (
    "encoding/json"
    "strings"
    "strconv"
)

func (reviewData ReviewData) reviewLengthCheck(advisorData *AdvisorData) {
    if reviewData.Length >= 100 {
        advisorData.Score = advisorData.Score - 0.5
    }
}

func (reviewData ReviewData) reviewBurstCheck(advisorData *AdvisorData) {
    reviewTime := strings.Replace(reviewData.Time[len(reviewData.Time)-5:], ":", "", 1)
    reviewCount := len(advisorData.Reviews)

    latestReview := advisorData.Reviews[reviewCount -1]
    latestReviewData := &ReviewData{}
    json.Unmarshal([]byte(latestReview), &latestReviewData)
    latestReviewTime := strings.Replace(latestReviewData.Time[len(latestReviewData.Time)-5:], ":", "", 1)

    reviewTimeInt, _ := strconv.Atoi(reviewTime)
    latestReviewTimeInt, _ := strconv.Atoi(latestReviewTime)

    reviewDate := strings.Replace(reviewData.Time, reviewData.Time[len(reviewData.Time)-5:], "", 1)
    latestReviewDate := strings.Replace(latestReviewData.Time, latestReviewData.Time[len(latestReviewData.Time)-5:], "", 1)

    if reviewDate == latestReviewDate {
        if (reviewTimeInt - latestReviewTimeInt) == 0 {
            advisorData.Score = advisorData.Score - 40
        } else if (reviewTimeInt - latestReviewTimeInt) < 100 {
            advisorData.Score = advisorData.Score - 20
        }
    }
}

func (reviewData ReviewData) reviewDeviceCheck(advisorData *AdvisorData) {
    if stringInArray(reviewData.Device, advisorData.Devices) {
        advisorData.Score = advisorData.Score - 30
    } else {
        advisorData.Devices = append(advisorData.Devices, reviewData.Device)
    }
}

func (reviewData ReviewData) reviewStarCheck(advisorData *AdvisorData) {
    if reviewData.Stars == 5 && advisorData.Stars < 3.5 {
        advisorData.Score = advisorData.Score - 8
    } else if reviewData.Stars == 5 {
        advisorData.Score = advisorData.Score - 2
    }
}

func (reviewData ReviewData) reviewSolicitedCheck(advisorData *AdvisorData) {
    if reviewData.Solicited == "solicited" {
        advisorData.Score = advisorData.Score + 3
    }
}

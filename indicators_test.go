package main

import (
    "testing"
    "reflect"
    "encoding/json"
)

func TestIndicators_reviewLengthCheck_notTooLong(t *testing.T) {

        expected := 100.0

        reviewData := &ReviewData{
            Length: 99,
        }
        advisorData := &AdvisorData{
            Score: 100,
        }

        reviewData.reviewLengthCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewLengthCheck_tooLong(t *testing.T) {

        expected := 99.5

        reviewData := &ReviewData{
            Length: 105,
        }
        advisorData := &AdvisorData{
            Score: 100,
        }

        reviewData.reviewLengthCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewSolicitedCheck_solicited(t *testing.T) {

        expected := 100.0

        reviewData := &ReviewData{
            Solicited: "solicited",
        }
        advisorData := &AdvisorData{
            Score: 97,
        }

        reviewData.reviewSolicitedCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewSolicitedCheck_unsolicited(t *testing.T) {

        expected := 97.0

        reviewData := &ReviewData{
            Solicited: "unsolicited",
        }
        advisorData := &AdvisorData{
            Score: 97,
        }

        reviewData.reviewSolicitedCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewStarCheck_fiveStar(t *testing.T) {

        expected := 98.0

        reviewData := &ReviewData{
            Stars: 5,
        }
        advisorData := &AdvisorData{
            Stars: 4,
            Score: 100,
        }

        reviewData.reviewStarCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewStarCheck_fiveStarAndLowAverage(t *testing.T) {

        expected := 92.0

        reviewData := &ReviewData{
            Stars: 5,
        }
        advisorData := &AdvisorData{
            Stars: 3.4,
            Score: 100,
        }

        reviewData.reviewStarCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewStarCheck_fourStar(t *testing.T) {

        expected := 100.0

        reviewData := &ReviewData{
            Stars: 4,
        }
        advisorData := &AdvisorData{
            Stars: 4,
            Score: 100,
        }

        reviewData.reviewStarCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewDeviceCheck_deviceInArray(t *testing.T) {

        expected := 70.0

        reviewData := &ReviewData{
            Device: "one",
        }
        advisorData := &AdvisorData{
            Devices: []string{"one","two"},
            Score: 100,
        }

        reviewData.reviewDeviceCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewDeviceCheck_deviceNotInArray(t *testing.T) {

        expectedScore := 100.0
        expectedDevices := []string{"one","two","three"}

        reviewData := &ReviewData{
            Device: "three",
        }
        advisorData := &AdvisorData{
            Devices: []string{"one","two"},
            Score: 100,
        }

        reviewData.reviewDeviceCheck(advisorData)

        if advisorData.Score != expectedScore {
                t.Fatalf("Expected %s, got %s", expectedScore, advisorData.Score)
        }
        if reflect.DeepEqual(expectedDevices, advisorData.Devices) == false {
                t.Fatalf("Expected %s, got %s", expectedDevices, advisorData.Devices)
        }
}

func TestIndicators_reviewBurstCheck_timesWithinAnMin(t *testing.T) {

        expected := 60.0

        latestReviewData := &ReviewData{
            Time: "2ndSeptember10:03",
        }
        latest, _ := json.Marshal(latestReviewData)

        advisorData := &AdvisorData{
            Score: 100,
        }
        advisorData.Reviews = append(advisorData.Reviews, string(latest))

        reviewData := &ReviewData{
            Time: "2ndSeptember10:03",
        }

        reviewData.reviewBurstCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewBurstCheck_timesWithinAnHour(t *testing.T) {

        expected := 80.0

        latestReviewData := &ReviewData{
            Time: "2ndSeptember09:30",
        }
        latest, _ := json.Marshal(latestReviewData)

        advisorData := &AdvisorData{
            Score: 100,
        }
        advisorData.Reviews = append(advisorData.Reviews, string(latest))

        reviewData := &ReviewData{
            Time: "2ndSeptember10:04",
        }

        reviewData.reviewBurstCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewBurstCheck_timesMoreThanHourArart(t *testing.T) {

        expected := 100.0

        latestReviewData := &ReviewData{
            Time: "2ndSeptember07:30",
        }
        latest, _ := json.Marshal(latestReviewData)

        advisorData := &AdvisorData{
            Score: 100,
        }
        advisorData.Reviews = append(advisorData.Reviews, string(latest))

        reviewData := &ReviewData{
            Time: "2ndSeptember10:04",
        }

        reviewData.reviewBurstCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

func TestIndicators_reviewBurstCheck_DifferentDaysWithinSameHour(t *testing.T) {

        expected := 100.0

        latestReviewData := &ReviewData{
            Time: "3ndSeptember10:00",
        }
        latest, _ := json.Marshal(latestReviewData)

        advisorData := &AdvisorData{
            Score: 100,
        }
        advisorData.Reviews = append(advisorData.Reviews, string(latest))

        reviewData := &ReviewData{
            Time: "2ndSeptember10:04",
        }

        reviewData.reviewBurstCheck(advisorData)

        if advisorData.Score != expected{
                t.Fatalf("Expected %s, got %s", expected, advisorData.Score)
        }
}

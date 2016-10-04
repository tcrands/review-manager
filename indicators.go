package main

import ()

func (reviewData ReviewData) reviewLengthCheck(advisorData *AdvisorData) {
    if reviewData.Length >= 50 {
        advisorData.Score = advisorData.Score - 1
    }
}

func (reviewData ReviewData) reviewBurstCheck(advisorData *AdvisorData) {

}

func (reviewData ReviewData) reviewDeviceCheck(advisorData *AdvisorData) {

}

func (reviewData ReviewData) reviewStarCheck(advisorData *AdvisorData) {
    if reviewData.Stars == 5 {
        advisorData.Score = advisorData.Score - 2
    }
}

func (reviewData ReviewData) reviewSolicitedCheck(advisorData *AdvisorData) {
    if reviewData.Solicited == "solicited" {
        advisorData.Score = advisorData.Score + 3
    }
}

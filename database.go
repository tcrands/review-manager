package main

import (
    "log"
    "github.com/boltdb/bolt"
    // "net/http"
    "fmt"
    "encoding/json"
)

func setupDatabase() (*bolt.DB, error){
    db, err := bolt.Open("advisorMap.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    return db, err
}

func updateDatabase(AdvisorData *AdvisorData) {
    db, _ := setupDatabase()
    defer db.Close()

    advisorName := AdvisorData.Name
    advisor, _ := json.Marshal(AdvisorData)

    stringedAdvisor := string(advisor)

    db.Update(func(tx *bolt.Tx) error {
        b, _ := tx.CreateBucketIfNotExists([]byte(advisorName))
        return b.Put([]byte(advisorName), []byte(stringedAdvisor))
    })

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(advisorName))
        advisor := b.Get([]byte(advisorName))
        fmt.Printf("Value Stored: %s\n", append(advisor, (" With Key: " + advisorName)...))
        return nil
    })
}

// func checkDatabase(depUrl string) (bool, string) {
//     db, _ := setupDatabase()
//     defer db.Close()
//
//     var prUrl = ""
//
//     fmt.Println(depUrl)
//
//     db.View(func(tx *bolt.Tx) error {
//         b := tx.Bucket([]byte("dependencyMapBucket"))
//         val := b.Get([]byte(depUrl))
//         prUrl = string(val[:])
//         fmt.Println(depUrl)
//         return nil
//     })
//
//     fmt.Printf(prUrl)
//
//     if prUrl != "" {
//         return true, prUrl
//     } else {
//         return false, prUrl
//     }
// }
//
// func removeKey(depUrl string) {
//     db, _ := setupDatabase()
//     defer db.Close()
//
//     db.Update(func(tx *bolt.Tx) error {
//         b := tx.Bucket([]byte("dependencyMapBucket"))
//         err := b.Delete([]byte(depUrl))
//         return err
//     })
// }
//
// func flushDatabase(w http.ResponseWriter, r *http.Request) {
//     db, _ := setupDatabase()
//     defer db.Close()
//
//     db.Update(func(tx *bolt.Tx) error {
//         tx.DeleteBucket([]byte("dependencyMapBucket"))
//         return nil
//     })
//
//     db.View(func(tx *bolt.Tx) error {
//         b := tx.Bucket([]byte("dependencyMapBucket"))
//         fmt.Printf("BUCKET: %s\n", b)
//         return nil
//     })
// }

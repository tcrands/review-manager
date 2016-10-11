package main

import (
    "log"
    "github.com/boltdb/bolt"
    "encoding/json"
)

func setupDatabase() (*bolt.DB, error){
    db, err := bolt.Open(dbMap, 0600, nil)
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
}

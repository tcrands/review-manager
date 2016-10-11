# review-manager 

The task can be run but either running `go build` in the directory then runnning the exec file, or by running the exec file sent in the email. There is also a postman collection in the email for the requests I tested against.

I used Godeps for handling the dependencies (In this case BoltDB) I upload the vendor folder direct to github so theres no need to build these.

I used MUX router to handle the api calls. I used BoltDB as a data store. Normally I wuld use some kind of database but for the scope of this task it didn't feel warrented. I considered just storing the data in memory but I wanted to demonstrate some kind of interaction with persistant storage. 


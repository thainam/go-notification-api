# GO Notification API

[![Go Test & Build](https://github.com/thainam/go-notification-api/actions/workflows/main.yaml/badge.svg)](https://github.com/thainam/go-notification-api/actions/workflows/main.yaml)

#### 

**Important:** This GO application was created for learning purpose only. If you're new in GO maybe here you can explore a little bit more about the language

## What is this application?

This application will have 4 endpoints:

* GET /health - Endpoint to check if the webserver is healthy. (Available 🎉️)
* GET /notifications - Endpoint to list the notifications from database. (Available 🎉️)
* GET /notifications/{id} - Endpoint to get details of a specific notification. (Work in Progress 🕙)
* POST /notifications - Endpoint to create a new notification. (Work in Progress 🕙)

## Where does Kafka fit into this application?

There's a Kafka Consumer working to read all notifications that comes into the topic and insert those into database (Sqlite3 -> `./notifications.db` File).

For now, it's only possible to publish the notification int Kafka topic using the Control-Center in `http://localhost:9021/` using the format below:

```json
{
    "Title": "Hello",
    "Message": "World!"
}
```

In the future will be possible to use an endpoint to publish the notifications to kafka. Like said above.


## How can I start the application?

Well, it's quite simple. You'll need to have Docker (and docker-compose) on your machine.

With this all set, just run:

`docker-compose up -d --build`

Enjoy it! 🎉️

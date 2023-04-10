# GO Notification API

[![Go Test & Build](https://github.com/thainam/go-notification-api/actions/workflows/main.yaml/badge.svg)](https://github.com/thainam/go-notification-api/actions/workflows/main.yaml) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thainam/go-notification-api) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/thainam/go-notification-api?label=version) ![GitHub Release Date](https://img.shields.io/github/release-date/thainam/go-notification-api)

#### 

**Important:** This GO application was created for learning purpose only. If you're new in GO maybe here you can explore a little bit more about the language.

## What is this application?

This application will have 4 endpoints:

🎉️ Available:

* GET /health - Endpoint to check if the webserver is healthy.
* GET /notifications - Endpoint to list the notifications from database.

🕙 Work in Progress:

* GET /notifications/{id} - Endpoint to get details of a specific notification.
* POST /notifications - Endpoint to create a new notification.

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

PS: Make sure you have this line not commented out in your hosts file:
`192.168.0.61 host.docker.internal`

Enjoy it! 🎉️

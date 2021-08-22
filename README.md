# Http Monitor 

A service to monitor registered endpoints by sending 10 request in a constant time steps. Updating the status code of each url, http monitor service logs an alert if the number of errors reache the given threshold.

# System Design

![alt text](https://github.com/shirinebadi/http-monitor/design.png?raw=true)

# Run

```
After downloding the project, navigate to http-monitor directory 
$ cd cmd/http-monitor
$ go build
$ ./http-monitor server
$ ./http-monitor worker

```

# Endpoints

## /register

```
$ curl -X POST http://localhost:21345/register -H 'Content-Type: application/json' -d '{"name":"you_username","password":"your_password"}'

```

## /login

```
$ curl -X POST http://localhost:21345/login -H 'Content-Type: application/json' -d '{"name":"you_username","password":"your_password"}'

```
## /request

```
$ curl -X POST http://localhost:21345/request -H 'Content-Type: application/json' -Token 'your token' -d '{"url":"you_url","threshold":"your_threshold"}'

```

## /result

```
$ curl -X GET http://localhost:21345/result -H 'Content-Type: application/json' -Token 'your token'

```

## /result/url

```
$ curl -X POST http://localhost:21345/result/url -H 'Content-Type: application/json' -Token 'your token' -d '{"url":"you_url"}'

```

## /alerts

```
$ curl -X GET http://localhost:21345/login -H 'Content-Type: application/json' -Token 'your token'

```



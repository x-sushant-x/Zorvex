## Zorvex 🚀
Zorvex is a robust framework developed in Golang, designed to simplify and streamline microservice registration, discovery, load balancing, and health checking within distributed systems. 

### Installation 🛠️

`
$ go get -u github.com/sushant/zorvex
`

<br>

### Deployment 🚀
1. Upload all files to a server for example AWS EC2.
2. Power up docker `sudo docker compose up -d`
3. Use command `make run` to start the whole system

<br>
<b>Note: </b> Zorvex agent runs on port 3000 and Zorvex client run on port 3001. So you may need to use nginx for reverse proxy if you are not able to access these ports in your server.

<br> <br>

### Load Balancing ⚖️

Zorver only support Round Robin load balancing at this time. But I'll make sure to add more strategies whenever I get time from my busy life.

<br>

### API Endpoints

### Register Service 📝

POST:  <SERVER_IP>/register

#### Example Body
<pre>
{
    "id": "1",
    "name": "sample_service",
    "description" : "This is the description of the service.",
    "tags": ["tag1", "tag2"],
    "http_method": "GET",
    "http_protocol" : "http",
    "ip_address": "192.168.192.212",
    "port": 5000,
    "register_time": "2023-12-12T12:00:00Z",
    "last_sync_time": "2023-12-12T12:00:00Z",
    "endpoint": "/sample_service",
    "load_balancing_method": "RoundRobin",
    "total_connections": 0,
    "de_register_after": 3600,
    "status": "active",
    "health_config": {
        "health_check_endpoint": "http://192.168.192.212:5000/health",
        "interval": 60,
        "options": {
            "http_headers": [
                {
                    "key": "Content-Type",
                    "value": "application/json"
                }
            ],
            "expected_status_code": 200
        }
    }
}

</pre>

<br>
<br>

### Discover Service 🔍

GET - <SERVER_IP>/discover?service=<SERVICE_NAME>

Response - Redirects to optimal server (round robin load balancing)
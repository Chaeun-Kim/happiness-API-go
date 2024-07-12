# Happiness Index API

This is a toy-size example attempting to build a very simple API in Go

### To Start

#### Prerequisite

- [Docker](https://www.docker.com/)

#### Steps

1. `docker build -t happiness .`
2. `docker run -p 5000:5000 happiness`

These two simple steps should start up a Go API, and you should be able to ping the service and get valid response

```
$ curl localhost:5000/ping
{
    "message": "service is healthy"
}
$ 
```

### Usage

This is a simple READ-ONLY API with two endpoints 

- `GET /happiness/{facet}/{id}`
- `POST /happiness/{facet}`

Currently, only `county` facet is available.

#### GET /happiness/{facet}/{id}; get happiness index by facet id

```
curl localhost:5000/happiness/county/9999
```

```
Example Response
{
    "data": {
        "facet": "county",
        "indices": [
            {
                "id": "9999",
                "value": 101.42
            }
        ]
    }
}
```

#### POST /happiness/{facet}; get happiness index by facet ids along with metrics

```
curl -X POST -d '{"counties":["9999","1001"], "metrics": ["average","max"]}' localhost:5000/happiness/county
```

```
Example Response
{
    "data": {
        "facet": "county",
        "indices": [
            {
                "id": "9999",
                "value": 101.42
            },
            {
                "id": "1001",
                "value": 103
            }
        ],
        "metrics": [
            {
                "name": "average",
                "value": 102.21000000000001
            },
            {
                "name": "max",
                "value": 103
            }
        ]
}
```
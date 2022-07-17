## go-proxy

Test project. HTTP server for proxying HTTP-request.
Receive URL, Method and Headers in json body, sends request to provided endpoint in `URL` param.

### To start
```
# start go-proxy
go run main.go --config=./etc/goproxy.yml

# send request to go-proxy
curl http://localhost:8090/json -d '{"url":"http://google.com", "method":"GET"}'
```


### Description

There are 2 way to communicate with go-proxy
1. json body
2. url params

#### 1. json body
Arguments can be sent as JSON body in request to endpoint (`/json` by default)
```
{
    "method": "GET",
    "url": "http://google.com",
    "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
        ....
    }
}

# for example:
> curl http://localhost:8090/json -d '{"url":"http://google.com", "method":"GET",  "headers": {"Authentication": "Basic bG9naW46cGFzc3dvcmQ="}}'
```

#### 2. url params
Arguments can be sent as url params in request to endpoint (`/url` by default)
headers param names should starts with `header-` prefix
```
method=GET
url=http://google.com
header-Authentication=Basic bG9naW46cGFzc3dvcmQ=
header-Content-Type=application/json

# for example:
> curl "http://localhost:8090/url?method=GET&url=http://google.com&header-Authentication=Basic%20bG9naW46cGFzc3dvcmQ=&header-Content-Type=application/json"
```

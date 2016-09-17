httpserv
=========

Simple HTTP server with headers logging ability written in Go.

Install:

```
go get -u github.com/dddpaul/httpserv
```

Or grab Docker image:

```
docker pull dddpaul/httpserv
```

Usage:

```
httpserv [OPTIONS]
-port string
      Port to listen (prepended by colon), i.e. :8080 (default ":8080")
-verbose
      Print all headers
```

Per request logging may be enabled by setting `X-Logging-Enabled header`, like:

```
http localhost:8080 X-Logging-Enabled:true  

```

or

```
curl -H "X-Logging-Enabled:true" localhost:8080
```

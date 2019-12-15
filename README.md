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
  -prefix string
        Logging prefix (default "httpserv")
  -headers
        Print request headers
  -message string
        Server response (default "HTTP OK")
  -port string
        Port to listen (prepended by colon), i.e. :8080 (default ":8080")
  -verbose
        Print request details
```

Per request logging may be enabled by setting `X-Headers-Logging-Enabled` header, like:

```
http localhost:8080 X-Headers-Logging-Enabled:true  
```

or

```
curl -H "X-Headers-Logging-Enabled:true" localhost:8080
```

# gin-ratelimit
gin-ratelimit, a ratelimit middleware for gin.


example:
```shell
# server
cd example
go build 


# client 
wrk -t4 -c20 -d30s http://localhost:12345/ping
```

# README

Build and start Kong Gateway:

```
docker compose up
```

Wait for Kong Gateway to start.

In a different terminal window, use cURL to send a request:

```
curl -i http://localhost:8000/hello
```

cURL should print the HTTP response header `X-Example-Message` which was inserted by `example-plugin`, demonstrating that the plugin is running.

Observe the log output in the first terminal window. Each request causes Kong Gateway to emit multiple log messages about plugin server instances being closed and started.

Excerpt from container logs:

```
2023/11/21 22:33:08 [debug] 2347#0: *654 [kong] pb_rpc.lua:381 [example-plugin] started plugin server: seq 15, worker 1, instance id 15
2023/11/21 22:33:08 [info] 2346#0: *151 [example-plugin:2350] 2023/11/21 22:33:08 closed instance 15, context: ngx.timer
2023/11/21 22:33:08 [error] 2347#0: *654 [kong] pb_rpc.lua:413 [example-plugin] closed, client: 192.168.65.1, server: kong, request: "GET /hello HTTP/1.1", host: "localhost:8000"
2023/11/21 22:33:08 [info] 2346#0: *151 [example-plugin:2350] 2023/11/21 22:33:08 no plugin instance 15, context: ngx.timer
2023/11/21 22:33:08 [warn] 2347#0: *654 [kong] pb_rpc.lua:417 [example-plugin] closed, client: 192.168.65.1, server: kong, request: "GET /hello HTTP/1.1", host: "localhost:8000"
2023/11/21 22:33:08 [debug] 2347#0: *654 [kong] pb_rpc.lua:381 [example-plugin] started plugin server: seq 15, worker 1, instance id 15
```

`example-plugin` keeps track of how many times `New()` has been called. We can check the count in the container logs:

```
docker compose logs | grep 'New() has been called'
```

I used [hey](https://github.com/rakyll/hey) to send a lot of requests in a single command:

```
hey -n 1000 http://localhost:8000/hello
```

With Kong Gateway 3.4.0.0, after sending 1000 requests, I observed a count of 5.

With Kong Gateway 3.4.1.0, after sending 1000 requests, I observed a count of 1855. I observed similar counts using versions 3.4.1.1, 3.4.2.0, 3.4.2.1, 3.5.0.0, and 3.5.0.1.

You can reproduce the previous behavior under Kong Gateway 3.4.0.0 by editing the Dockerfile, changing the tag on line 7:

```diff
- FROM docker.io/kong/kong-gateway:3.4.1.0
+ FROM docker.io/kong/kong-gateway:3.4.0.0
```

Rebuild and restart the container:

```
docker compose up --build
```

Wait for Kong Gateway to start, then send the same number of requests again, and observe the container logs.

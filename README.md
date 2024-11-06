# GoApp

An application for Go knowledge assessment.

## Description

This is a web application that utilises websockets. A client connects on `localhost:8080` and has three options: a) `open` a websocket connection that reads values from a counter b) `close` the websocket and c) reset the counter to zero. The counter is feeded from a random string generator. On WS session termination, statistics for terminated session are printed.

## Build the application
The application is compiled by running `make` in the root folder and the final binaries are found in the `bin/` folder.

The binaries created are *server* and *client* executables

## Server

### Endpoints

**/goapp** : Returns main html page with controls for websocket.\ User can open and close a websocket connection. Upon connected counter info will be diplayed in page.\
E.g.

```
OPEN
RESPONSE: {"iteration":1,"value":"822876EF10"}
RESPONSE: {"iteration":2,"value":"215100491D"}
RESPONSE: {"iteration":3,"value":"05DCC3B6AB"}
CLOSE
```

Supported Methods: GET\
**/goapp/ws**: Websocket endpoint.\
Supported Methods: GET\
**/goapp/health**: Healthcheck endpoint. Returns success string.\
Supported Methods: GET\
**/goapp/restricted**: Return succcess string. Use for csrf demo. Try sending post request from a rest client to get csrf error.\
Supported Methods: POST\

### Websocket

Upon successful connection, server will notify all clients with a per connection counter and a hexstring.\
Hexstring is always the same for all clients.\
To start the server just run the server executable

`./bin/server`

## Client

Client opens given number of websocket connections to server. It prints the iteration counter and the hexstring received by the server for every connection

```
$ ./bin/client -c 3
[conn #0] iteration: 1, value: 66D53ED788
[conn #1] iteration: 1, value: 66D53ED788
[conn #2] iteration: 1, value: 66D53ED788
...
```

## Profiling

Metrics for heap before and after refactor of websocket for memory usage.
[pprof.server.alloc_objects.alloc_space.inuse_objects.inuse_space.before_refactor.pb.gz](https://github.com/user-attachments/files/17654264/pprof.server.alloc_objects.alloc_space.inuse_objects.inuse_space.before_refactor.pb.gz)
[pprof.server.alloc_objects.alloc_space.inuse_objects.inuse_space.after_refactor.pb.gz](https://github.com/user-attachments/files/17654265/pprof.server.alloc_objects.alloc_space.inuse_objects.inuse_space.after_refactor.pb.gz)



Files can be viewed using pprof tool like:

        go tool pprof -http=:8081 pprof.server.alloc_objects.alloc_space.inuse_objects.inuse_space.before_refactor.pb.gz

or to compare:

        go tool pprof -http=:8081 -diff_base pprof.server.alloc_objects.alloc_space.inuse_objects.inuse_space.before_refactor.pb.gz pprof.server.alloc_objects.alloc_space.inuse_objects.inuse_space.after_refactor.pb.gz 

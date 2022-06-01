# Go Chat Server
Simple Go chat server

## Running 
The `build.sh` script contains the script needed to compile the code into an executable
```code
./build.sh
```

After running this command, you can launch the server like so
```code
./main
```

You can also pass in additional parameters through the command line:

```
optional arguments:
  -h      host server should run on
  -p      port server should run on
  -d      enable for debug mode
  --help  show help message and exit
```

## Server 
`middleware.go` - helpful middleware for incoming requests

`server.go` - `Server` struct represents the server configured with all service routers 



## Services
### Message
`router.go` - router and route handlers for message service

`model.go` - `Message` struct which represents a message produced by a user

`store.go` - in-memory storage for messages



# Frisbee Examples

This repository contains a series of examples for Frisbee. You can learn more about these examples at https://loopholelabs.io/docs/frisbee.

## Echo Example

This example is a simple echo client/server, where the client will repeatedly send messages to the server, and the server will echo them back. Its purpose 
is to describe the flow of messages from Frisbee Client to Server, as well as give an example of how a Frisbee application must be implemented. 

To start the example run:

```shell
go run echo/server/main.go
go run echo/client/main.go
```

Please be advised that these need to run as separate processes.

## PUB/SUB Example

This example is a simple PUB/SUB system, where a broker, publisher, and subscriber have been implemented. 

When the subscriber starts, it connects to the broker and sends a `SUB` message with an embedded topic. This is to tell the broker 
that the subscriber is interested in the topic. In then waits for `PUB` messages from the broker and prints
the contents to the screen.

When the publisher starts, it connects to the broker and repeatedly sends `PUB` messages with varying content.

When the broker starts, it waits for `SUB` message types, and when it receives them it registers the connection that 
sent them as an interested party in the `SUB` topic. It also listens for `PUB` messages and routes those based on
the topic provided. 

To start the example run:

```shell
go run pubsub/broker/main.go
go run pubsub/subscriber/main.go
go run pubsub/publisher/main.go
```

Please be advised that these need to run as separate processes. 

You can also spawn multiple subscribers and publishers and Frisbee will automatically
handle the extra connections.
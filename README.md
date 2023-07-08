# Udp Hole Punching
UDP hole punching POC written in Golang.

## Server
Run the server with listening port as the first argument. For example
```
server.exe 12345
```
This command will run the server on port 12345.

## Client
Punching is done via keys. Each pair of clients which do have same keys, will be connected together.
To run the client, you need a server, a key and another client. Run the client with server address
as the first argument and the key as the second. Key needs to be less than 64 characters. For example:
```
client.exe 127.0.0.1:12345 my_secret_key
```
After two clients are connected to server, server sends the other peer's IP to each peer and punching
happens.
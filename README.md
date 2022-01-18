## Pseudo Auth
Producer is the one sharing their activity
Server is the server relaying activity and handling generating IDs
Client is the one wanting to observe a producers activity

The producer sets a password and sends a request for registration to the server
The server assigns a random UUIDv4 to the hashed password on registration and sends the UUIDv4 to
the producer
The producer saves the UUIDv4 and sends the set password to the server on activity indicating
requests
The server generates a UUIDv5 to share with clients
The client requests the activity of a producer through the UUIDv5

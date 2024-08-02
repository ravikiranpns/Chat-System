# Chat-System
A Go microservice that handles user authentication, message sending, and message retrieval.


# Chat Service

This project is a simplified chat platform built using Go, with Cassandra for data storage and Redis for caching. The chat service allows users to register, login, send messages, and retrieve messages.

## Features

- User Registration
- User Login
- Send Messages
- Retrieve Messages

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```sh
git clone <repository-url>
cd chat-system

##Build and Run the Containers
docker-compose up --build



Verify the Setup

To verify that the chat service is running, you can use curl commands.

curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "user1", "password": "pass1"}'
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "user2", "password": "pass2"}'


	2.	Login Users:

curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "user1", "password": "pass1"}'
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "user2", "password": "pass2"}'

	3.	Send Messages:


curl -X POST http://localhost:8080/send -H "Content-Type: application/json" -d '{"sender": "user1", "recipient": "user2", "content": "Hi user2, this is user1."}'
curl -X POST http://localhost:8080/send -H "Content-Type: application/json" -d '{"sender": "user1", "recipient": "user2", "content": "How are you, user2?"}'

	4.	Retrieve Messages:


curl -X GET 'http://localhost:8080/messages?username=user1'
curl -X GET 'http://localhost:8080/messages?username=user2'



Database Setup

Access Cassandra ShellTo check the data directly in the Cassandra database, access the Cassandra shell using the following command:



docker exec -it <cassandra_container_id> cqlsh

Replace <cassandra_container_id> with the actual container ID or name for Cassandra.


Verify Data in Cassandra

USE chat;
SELECT * FROM users;
SELECT * FROM messages;


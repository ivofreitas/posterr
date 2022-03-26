# Posterr

## Entities and functionality
### post
- Create(userID: uuid, content: string) => Post
- List(limit: int, offset: int) => []Post
- ListByFollower(follower: uuid, limit: int, offset: int) => []Post
- ListByUser(userID: uuid, limit: int, offset: int) => []Post
- Quote(postID: uuid, content: string) => Post
- Repost(postID: uuid) => Post

### follower
- Follow(userID: uuid, following: uuid) => Follower
- UnFollow(userID: uuid, following: uuid) => Follower

## Requirements

- Go v1.17 https://golang.org/dl/
- Docker https://www.docker.com/

## Installation

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

Getting started
```
make setup
```
Run all tests with coverage
```
make coverage
```
Build a binary
```
make build
```

## Documentation

https://documenter.getpostman.com/view/2675416/UVyn2JQY
Requirements:
 - docker
 - go 


Get mongo docker image:

docker pull mongo:latest

Run app: 
- make build
- make run

When done:
- make clean

Tests:
- make test


Note:
I've tried to dockerize the app, yet I didn't succeed. Go app container won't connect to mongo container
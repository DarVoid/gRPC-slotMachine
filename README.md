# Slot machine using gRPC services

the aim of this project was to use gRPC to deal with all the logic and having a frontend facing gateway to deal with http requests.

I'll first start by using an express server and then explore other options

# Install node packages

    > npm install 

# Running the project:
run gRPC server:

windows

    > go run .\services\serverSimple.go 

linux

    > go run services/serverSimple.go 

## run node gateway server:
windows

    > node .\server.js 

linux

    > node server.js 


OR 

## run gorilla gateway server:

windows

    > go run .\gatewayGoserver\gateway.go
    
linux

    > go run gatewayGoserver/gateway.go

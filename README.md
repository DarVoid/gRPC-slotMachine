# Slot machine using gRPC services

the aim of this project was to use gRPC to deal with all the logic and having a frontend facing gateway to deal with http requests.

I'll first start by using an express server. and then a graphQL node server.

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

    > node .resExpress\server.js 

linux

    > node server.js 


OR 

## run gorilla gateway server:

windows

    > go run .\gatewayGoserver\gateway.go
    
linux

    > go run gatewayGoserver/gateway.go

OR 

## run grapqhQL server server

    windows

    > node .\graphQL\index.js

    linux

    > node graphQL/index.js

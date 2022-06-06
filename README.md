# Slot machine using gRPC services

the aim of this project was to use gRPC to deal with all the logic and having a frontend facing gateway to deal with http requests.

I'll first start by using an express server and then explore other options

# Install node packages

    > npm install 

# Running the project:
run gRPC server:

windows

    > go run .\gameService\gameServiceServer.go 

linux

    > go run gameService/gameServiceServer.go 

## run node gateway server:
windows

    > node .\gatewayNodejsServer\server.js 

linux

    > node gatewayNodejsServer/server.js 

OR 

## run negroni gateway server:

windows

    > go run .\gatewayNegroniGoServer\gateway.go
    
linux

    > go run gatewayNegroniGoServer/gateway.go


OR 

## run gorilla gateway server:

windows

    > go run .\gatewayGorillaGoServer\gateway.go
    
linux

    > go run gatewayGorillaGoServer/gateway.go
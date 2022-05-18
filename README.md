# Slot machine using gRPC services

the aim of this project was to use gRPC to deal with all the logic and having a frontend facing gateway to deal with http requests.

I'll first start by using an express server. and then a graphQL node server.

# Install node packages

    > npm install 

# Running the project:
run gRPC server:
 > go run .\services\serverSimple.go 

run node gateway server:

 > node .\server.js 

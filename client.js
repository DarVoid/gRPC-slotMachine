const grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");
const PROTO_PATH = "./src_dir/slot.proto";

const options = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
};

var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

const GameService = grpc.loadPackageDefinition(packageDefinition).GameService;
console.log(GameService.service.CreateGame.requestType)
const client = new GameService(
    "localhost:9000",
    grpc.credentials.createInsecure()
);


module.exports = client;
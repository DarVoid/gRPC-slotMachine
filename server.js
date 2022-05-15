const http = require("http");
const client = require("./client");

const host = "localhost";
const port = 8000;

const requestListener = function (req, res) {
    const url = req.url.split("/");
    const method = req.method;

    switch (method) {
        case "GET":
        case "PUT":
        case "DELETE":
        case "POST":
            client.CreateGame(
                {
                    "WinChance": 20, 
                    "TotalJogadas": 200,
                },
                (error, game) => {
                    if (error) throw error;
                    res.end({
                        data: game,
                        msg: "Successfully created a game.",
                    });
                }
            );
            break;
        default:
            res.end("");
            break;
    }
};

const server = http.createServer(requestListener);
server.listen(port, host, () => {
    console.log(`Server is running on http://${host}:${port}`);
});

const http = require("http");
const client = require("./client");

const host = "localhost";
const port = 8000;

const requestListener = function (req, res) {
    const url = req.url.split("/");
    const method = req.method;
    console.log(url)
    switch (method) {
        case "GET":
        case "PUT":
        case "DELETE":
        case "POST":
            if(url[1] === "create") {

                client.CreateGame(
                    {
                        winChance: 20, 
                        totalJogadas: 200,
                    },
                    (error, game) => {
                        if (error) throw error;
                        res.write(JSON.stringify({
                            data: game,
                            msg: "Successfully created a game.",
                        }));
                        res.end();
                    }
                    );
            }else if(url[1] === "play"){
                client.PlayGame(
                    {
                        gameId: url[2],
                        name: "Jorge",
                        luckyQuote: "socorro, let me win"
                    },
                    (error, response) => {
                        if (error) throw error;
                        res.write(JSON.stringify({
                            data: response,
                            msg: "Successfully played a game.",
                        }));
                        res.end();
                    }
                    );
            }else if(url[1] === "exists"){
                client.GameExists(
                    {
                        gameId: url[2],
                    },
                    (error, response) => {
                        if (error) throw error;
                        res.write(JSON.stringify({
                            data: response,
                            msg: "Successfully checked if a game exists.",
                        }));
                        res.end();
                    }
                    );
            }
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

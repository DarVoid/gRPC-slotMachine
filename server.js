var express = require("express");
var app = express();
//const http = require("http");
const client = require("./client");

const host = "localhost";
const port = 8080;

app.use(function (req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header(
        "Access-Control-Allow-Headers",
        "Origin, X-Requested-With, Content-Type, Accept"
    );
    next();
});

app.use(
    express.json()
)
app.use(    
    express.urlencoded({ extended: true })
)


app.post("/create", (req, res) => {
    console.log(req.body)
    client.CreateGame(
        {
            totalJogadas: req.body.totalJogadas,
            winChance: req.body.winChance,
        },
        (error, game) => {
            if (error) {
                console.error(error);
                JSON.stringify({
                    msg: "Could not create game",
                });
                res.end();
            } else {
                res.write(
                    JSON.stringify({
                        data: game,
                        msg: "Successfully created a game.",
                    })
                );
                res.end();
            }
        }
    );
    //res.send("Hello World!");

});
app.post("/play", (req, res) => {
    console.log({
        gameId: req.body.gameId,
        name: req.body.nameGuy,
        luckyQuote: req.body.luckyQuote,
    })
    client.PlayGame(
        {
            gameId: req.body.gameId,
            name: req.body.nameGuy,
            luckyQuote: req.body.luckyQuote,
        },
        (error, response) => {
            if (error) {
                console.error(error);
                res.write(
                    JSON.stringify({
                        msg: "Could not play game",
                    })
                );
                res.end();
            } else {
                res.write(
                    JSON.stringify({
                        data: response,
                        msg: "Successfully played a game.",
                    })
                );
                res.end();
            }
        }
    );

});

app.post("/exists", (req, res) => {
    client.GameExists(
        {
            gameId: reqreq.body.gameId,
        },
        (error, response) => {
            if (error) {
                console.error(error);
                res.write(
                    JSON.stringify({
                        msg: "Could not check if game exists",
                    })
                );
                res.end();
            } else {
                res.write(
                    JSON.stringify({
                        data: response,
                        msg: "Successfully checked if a game exists.",
                    })
                );
                res.end();
            }
        }
    );

});

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`);
});
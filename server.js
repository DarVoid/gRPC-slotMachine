var express = require("express");
var app = express();
//const http = require("http");
const client = require("./client");

const host = "localhost";
const port = 8000;

app.use(function (req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header(
        "Access-Control-Allow-Headers",
        "Origin, X-Requested-With, Content-Type, Accept"
    );
    next();
});


app.post("/create/:totalJogadas/:winChance", (req, res) => {

    client.CreateGame(
        {
            totalJogadas: req.params['totalJogadas'],
            winChance: req.params['winChance'],
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
app.post("/play/:gameId/:nameGuy/:luckyQuote", (req, res) => {
    console.log({
        gameId: req.params['gameId'],
        name: req.params['nameGuy'],
        luckyQuote: req.params['luckyQuote'],
    })
    client.PlayGame(
        {
            gameId: req.params['gameId'],
            name: req.params['nameGuy'],
            luckyQuote: req.params['luckyQuote'],
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

app.post("/exists/:gameId", (req, res) => {
    client.GameExists(
        {
            gameId: req.params['gameId'],
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
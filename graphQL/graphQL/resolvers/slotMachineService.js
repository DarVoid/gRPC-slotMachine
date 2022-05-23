const { UserInputError } = require("apollo-server");

const client = require("../../gRPCclients/client.js");

module.exports = {
    Mutation: {
        createGame: async (_, { newGameParams }, context) => {
            console.log(newGameParams);
            return await client.CreateGame(
                {
                    totalJogadas: newGameParams.totalJogadas,
                    winChance: newGameParams.winChance,
                },
                (error, game) => {
                    console.log(game);

                    return  new Promise((resolve, reject) => {
                        resolve(game);
                      });
                     //    if (error) {
                    //        console.error(error);
                    //        JSON.stringify({
                    //            msg: "Could not create game",
                    //        });
                    //        res.end();
                    //    } else {
                    //        res.write(
                    //            JSON.stringify({
                    //                data: game,
                    //                msg: "Successfully created a game.",
                    //            })
                    //        );
                    //        res.end();
                    //    }
                }
            );
        },
        playGame: async (_, { play }, context) => {
            console.log(play);
            console.log(context);
            if (post) {
            } else {
                throw new UserInputError("Post not found");
            }
        },
    },
    Query: {
        async gameExists(_, { gameId }) {
            console.log(gameId);
            try {
            } catch (err) {
                throw new Error(err);
            }
        },
    },
};

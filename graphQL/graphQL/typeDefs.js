const gql = require("graphql-tag");

module.exports = gql`
    input ShowGameRequest {
        gameId: String!
    }

    input PlayRequest {
        gameId: String!
        name: String!
        luckyQuote: String!
    }
    input CreateGameRequest {
        winChance: Int!
        totalJogadas: Int!
    }
    type NewGameReply {
        gameId: String
    }
    type ResultPlayReply {
        gameId: String!
        name: String!
        luckyQuote: String!
        reward: Boolean!
    }
    type Query {
        gameExists(gameId: ShowGameRequest!): Boolean!
    }
    type Mutation {
        createGame(newGameParams: CreateGameRequest!): NewGameReply
        playGame(play: PlayRequest!): ResultPlayReply!
    }
`;

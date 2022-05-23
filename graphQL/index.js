const  { ApolloServer, PubSub } = require('apollo-server');
//const mongoose = require('mongoose');

//const { MONGODB } = require('./config.js');
const typeDefs = require('./graphQL/typeDefs');
const resolvers = require('./graphQL/resolvers')



const server = new ApolloServer({
    typeDefs,
    resolvers,
    context: ({req})=>({req})
});
const res = server.listen({port: 5000});
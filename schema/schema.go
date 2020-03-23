package schema

const Schema = `schema {
    query: Query
    mutation: Mutation
}

type Pet {
    id: Int!
    name: String!
    tag: String!
}

input PetInput {
    name: String!
    tag: String!
}

type Query {
    getPets: [Pet!]!
    getPetById(id: Int!): [Pet!]!
} 

type Mutation {
    addPet(input: PetInput!): Pet!
}`

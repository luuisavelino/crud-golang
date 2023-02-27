DROP DATABASE IF EXISTS loja;    

CREATE DATABASE loja;    

\c loja;        

CREATE TABLE produtos (
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
);
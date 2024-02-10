-- +migrate Up
CREATE TABLE IF NOT EXISTS clientes (
    id INTEGER PRIMARY KEY,
    limite INTEGER,
    saldo INTEGER,
    saldo_inicial INTEGER
);
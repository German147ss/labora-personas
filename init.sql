-- Script para crear la tabla "personas"
CREATE TABLE personas IF NOT EXISTS (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    edad INTEGER NOT NULL,
    country_code VARCHAR(10) NOT NULL
);
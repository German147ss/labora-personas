BEGIN;
--1
INSERT INTO personas (nombre, apellido, edad, country_code) VALUES ('Juan2', 'Perez', 30, 'PE');

--2
INSERT INTO personas (nombre, apellido, edad, country_code) VALUES ('Maria2', 'Gomez', 25, 'CLP');

--3
INSERT INTO personas (nombre, apellido, edad) VALUES ('Pedro2', 'Gonzalez', 40);
-- Rollback or Commit
ROLLBACK; -- COMMIT;
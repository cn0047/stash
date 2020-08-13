// Cypher administration

Node types: leader, follower.

SHOW DATABASES;
SHOW DATABASE system;
CREATE DATABASE mydb if not exists; // Enterprise Edition
STOP DATABASE mydb;
START DATABASE mydb;
SHOW DATABASE mydb;
DROP DATABASE mydb;

CALL db.schema.visualization();
CALL db.labels();
CALL dbms.procedures() YIELD name, signature RETURN * ;

// new user
:USE system;
CALL dbms.security.createUser('dbu', 'dbp', false);
SHOW USERS;

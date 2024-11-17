DROP DATABASE IF EXISTS etlap;
CREATE DATABASE etlap;
USE etlap;
CREATE TABLE Vasarlok (
	   VasarloID int NOT NULL AUTO_INCREMENT,
	   Nev varchar(128) NOT NULL,
	   Email varchar(128) NOT NULL,
	   Telefonszam int NOT NULL,
	   PRIMARY KEY(VasarloID)
);
CREATE TABLE Etelek (
	   EtelID int NOT NULL AUTO_INCREMENT,
	   Nev varchar(128) NOT NULL,
	   Leiras varchar(1024),
	   Kep varchar(128),
	   Ar DECIMAL NOT NULL,
	   PRIMARY KEY(EtelID)
);
CREATE TABLE Rendelesek (
	   VasarloID int NOT NULL,
	   EtelID int NOT NULL,
	   FOREIGN KEY (VasarloID) REFERENCES Vasarlok(VasarloID),
	   FOREIGN KEY (EtelID) REFERENCES Etelek(EtelID)
);

INSERT INTO Vasarlok (Nev, Email, Telefonszam) VALUES
("Bob Nagy", "bob@example.com", 36201234),
("Jakab", "jakab@example.com", 12345),
("Cecilia", "cecilia@example.com", 54321);

DROP DATABASE IF EXISTS etlap;
CREATE DATABASE etlap;
USE etlap;
CREATE TABLE Vasarlok (
	   VasarloID int NOT NULL,
	   Nev varchar(128) NOT NULL,
	   Email varchar(128) NOT NULL,
	   Telefonszam int NOT NULL,
	   PRIMARY KEY(VasarloID)
);
CREATE TABLE Etelek (
	   EtelID int NOT NULL,
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

INSERT INTO Vasarlok (VasarloID, Nev, Email, Telefonszam) VALUES
(1, "Bob Nagy", "bob@example.com", 36201234),
(2, "Jakab", "jakab@example.com", 12345),
(3, "Cecilia", "cecilia@example.com", 54321);

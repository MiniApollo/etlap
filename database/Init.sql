DROP DATABASE IF EXISTS etlap;
CREATE DATABASE etlap;
USE etlap;
CREATE TABLE Vasarlok (
	   VasarloID int NOT NULL AUTO_INCREMENT,
	   Nev varchar(128) NOT NULL,
	   Email varchar(128) NOT NULL,
	   Telefonszam varchar(32) NOT NULL,
	   LeadasiIdo DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	   ElkeszultIdo DATETIME,
	   Osszeg DECIMAL,
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
	   Ar DECIMAL NOT NULL,
	   Darab int NOT NULL,
	   FOREIGN KEY (VasarloID) REFERENCES Vasarlok(VasarloID),
	   FOREIGN KEY (EtelID) REFERENCES Etelek(EtelID)
);

CREATE DATABASE etlap;
USE etlap;
CREATE TABLE Vasarlo (
	   RendelesID int NOT NULL,
	   Nev varchar(128) NOT NULL,
	   Email varchar(128) NOT NULL,
	   Telefonszam int,
	   PRIMARY KEY(RendelesID)
);
CREATE TABLE Etelek (
	   EtelID int NOT NULL,
	   Nev varchar(128) NOT NULL,
	   Leiras varchar(1024),
	   Kep varchar(128),
	   Ar DOUBLE NOT NULL,
	   PRIMARY KEY(EtelID)
);
CREATE TABLE Rendelesek (
	   RendelesID int NOT NULL ,
	   EtelID int NOT NULL,
	   FOREIGN KEY (RendelesID) REFERENCES Vasarlo(RendelesID),
	   FOREIGN KEY (EtelID) REFERENCES Etelek(EtelID)
);

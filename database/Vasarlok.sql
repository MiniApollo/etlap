INSERT INTO Vasarlok (Nev, Email, Telefonszam, ElkeszultIdo, Osszeg) VALUES
("Kis Róbert", "kisrobert@example.com", "06201234567", NULL, 1650),
("Kovács Bence", "kovacsbence@gmail.com", "06203202368", NULL, 7200),
("Nagy Attila", "nagyattila@example.com", "06704321441", DATE_ADD(NOW(), INTERVAL 1 MINUTE), 6200),
("Fodor Cecília", "fodorcecilia@hotmail.com", "06201182038", DATE_ADD(NOW(), INTERVAL 1 MINUTE), 10700);

# Tartalomjegyzék

-  [Étlap Weboldal](#Étlap)
-  [Specifikáció](#Specifikáció)
    -  [Funkciók](#Funkciók)
    -  [Adatok](#Adatok)
-  [Fejlesztői Eszközök](#Fejlesztői-Eszközök)
-  [Fejlesztők](#Fejlesztők)

<a id="Étlap"></a>

# Étlap Weboldal

Egy weboldal, ahol ki lehet választani, milyen ételeket szeretne a vevő vásárolni és hozzá tudja adni egy listához, amiket meg tud rendelni. A rendelését a helyszínen át tudja venni.


<a id="Specifikáció"></a>

# Specifikáció

A vevő ki tudja választani egy listából, milyen ételeket akar vásárolni és azokat egy rendelésbe össze tudja gyűjteni, amit el tud küldeni szükséges adatok megadása után az étteremnek.


<a id="Funkciók"></a>

## Funkciók

-   [X] Navigáció (Navbar)
-   [X] Ételek listája (Étlap)
-   [X] Rendelés megtekintése
-   [ ] Rendelés leadása


### Admin Funkciók:

-   [ ] Ételek hozzáadása, módosítása és törlése
-   [ ] Admin felület felhasználók rendelésének megtekintésére


<a id="Adatok"></a>

## Adatok


### Vasarlók adatai

-   Név
-   Email
-   Telefonszám
-   ID: Primary Key


### Rendelések (Kapcsoló tábla)

-   Vasarlo ID: Foreign Key Rendelés
-   Étel ID: Foreign Key Étel


### Ételek adatai

-   Név
-   Leírás
-   Kép
-   Ár
-   ID: Primary Key


<a id="Fejlesztői-Eszközök"></a>

# Fejlesztői Eszközök


## Frontend:

-   HTML CSS JavaScript
-   Tailwind: CSS keretrendszer
-   Vue.js: JavaScript keretrendszer


## Backend:

-   Go programozási nyelv
-   Gin: HTTP web keretrendszer
-   Godotenv: .env változok beolvasása
-   Go-MySQL-Driver

## Adatbázis:

-   MySQL


## Verziókezelés:

-   Git
-   Github


<a id="Fejlesztők"></a>

# Fejlesztők

13.D Szoftver vegyes

-   MiniApollo, SmarkPogany: Surmann Márk
-   BudaLevente: Buda Levente István
-   ocsaimark: Ócsai Márk

![etlapProjektGoVue](https://github.com/user-attachments/assets/7b4b702b-75de-49f6-824d-d1b6d5348dfc)


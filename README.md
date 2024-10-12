# Tartalom Jegyzék

-  [Étlap Weboldal](#org7f7f720)
    -  [Működése](#org953753c)
    -  [Funkciók](#orgfe5793a)
    -  [Adatok](#org6aa2253)
    -  [Fejlesztői Specifikáció](#org1ad515a)
    -  [Fejlesztők](#org5062dff)

<a id="org7f7f720"></a>

# Étlap Weboldal

Egy weboldal amin ki lehet választani milyen ételeket szeretne a vevő vásárolni és hozzá tudja adni egy listához amiket megtud rendelni. A rendelését a helyszínen át tudja venni.


<a id="org953753c"></a>

## Működése

A vevő ki tudja választani egy listából milyen ételeket akar vásárolni és azokat egy rendelésbe össze tudja gyűjteni, amit eltud küldeni szükséges adatok megadása után az étteremnek.


<a id="orgfe5793a"></a>

## Funkciók

-   [ ] Navigáció (Navbar)
-   [ ] Ételek listája (Étlap)
-   [ ] Rendelés megtekintése
-   [ ] Rendelés megvétele


### Admin Funkciók:

-   [ ] Ételek hozzáadása módosítása és törlése
-   [ ] Admin felület felhasználók rendelésének megtekintésére


<a id="org6aa2253"></a>

## Adatok


### Rendelési adatok

-   Név
-   Email
-   Telefonszám
-   ID: Primary Key


### Kapcsoló tábla

-   Rendelés ID: Foreign Key Rendelés
-   Étel ID: Foreign Key Étel


### Ételek adatai

-   Név
-   Leírás
-   Kép
-   Ár
-   ID: Primary Key


<a id="org1ad515a"></a>

## Fejlesztői Specifikáció


### Frontend:

-   HTML CSS JavaScript
-   Tailwind: CSS keretrendszer
-   Vue.js: JavaScript keretrendszer


### Backend:

-   Go programozási nyelv
-   Gin: HTTP web keretrendszer
-   Gorm: ORM könyvtár


### Adatbázis:

-   MySQL


### Verziókezelés:

-   Git
-   Github


<a id="org5062dff"></a>

## Fejlesztők

13.D Szoftver vegyes

-   MiniApollo, SmarkPogany: Surmann Márk
-   BudaLevente: Buda Levente István
-   ocsaimark: Ócsai Márk

![etlapProjektGoVue](https://github.com/user-attachments/assets/7b4b702b-75de-49f6-824d-d1b6d5348dfc)


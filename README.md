# Tartalom Jegyzék

1.  [Étlap Weboldal](#orgacfafa8)
    1.  [Működése](#org0235c9d)
    2.  [Funkciók](#org78b9e72)
    3.  [Adatok](#org3d96be1)
    4.  [Fejlesztői Specifikáció](#org5efdf2d)
    5.  [Fejlesztők](#org3792b4f)



<a id="orgacfafa8"></a>

# Étlap Weboldal

Egy weboldal amin ki lehet választani milyen ételeket szeretne a vevő vásárolni és hozzá tudja adni egy listához amiket megtud rendelni. A rendelését a helyszínen át tudja venni.


<a id="org0235c9d"></a>

## Működése

A vevő ki tudja választani egy listából milyen ételeket akar vásárolni és azokat egy rendelésbe össze tudja gyűjteni, amit eltud küldeni szükséges adatok megadása után az étteremnek.


<a id="org78b9e72"></a>

## Funkciók

-   Ételek listája (Étlap)
-   Rendelés megtekintése


### Admin Funkciók:

-   Ételek hozzáadása módosítása és törlése
-   Admin felület felhasználók rendelésének megtekintésére


<a id="org3d96be1"></a>

## Adatok


### Rendelési adatok

-   Név
-   Email
-   Telefonszám
-   Ételek: Foreign Key Étel
-   ID: Primary Key


### Ételek adatai

-   Név
-   Leírás
-   Kép
-   Ár
-   Elkészítési idő
-   ID: Primary Key


<a id="org5efdf2d"></a>

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

-   SQL


### Verziókezelés:

-   Git
-   Github


<a id="org3792b4f"></a>

## Fejlesztők

13.D Szoftver vegyes

-   MiniApollo, SmarkPogany: Surmann Márk
-   BudaLevente: Buda Levente István
-   ocsaimark: Ócsai Márk


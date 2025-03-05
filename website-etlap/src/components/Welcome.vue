<script setup lang="ts">

</script>

<template>
    <div class="grid grid-cols-2 max-md:grid-cols-1">
        <div class="mx-1 lg:mx-16 text-center">
            <h1 class="text-7xl lg:text-8xl my-8">Üdvözöljük étlapunkon</h1>
            <p class="m-8">Szolgáltatásunkal rendeld meg ételedet online, hogy rögktön átvehesd amint megérkeztél!
                Válassz kedvenc ételeidet közül otthonod kényelméből és mi előre elkészítjük a villámgyors kiszolgálás
                érdekében.
            </p>
            <RouterLink
                class="p-2 bg-slate-300 border-2 rounded-3xl border-black text-black hover:text-xl transition-all duration-500"
                to="/etelek">Étlap megtekíntése</RouterLink>
        </div>
        <div class="mx-2 my-4 lg:mx-6">
            <img class="border-4 rounded-md" src="http://localhost:8080/assets/foods/singaporeScaled.jpg"
                alt="Rustic food">
        </div>
        <div class="mx-1 lg:mx-16 text-center md:my-6">
            <h1 class="text-5xl lg:text-7xl mb-8">Megnyitottunk</h1>
            <p class="m-6">Hosszas várakozás és fejlesztés után végre megnyitott a várva várt szolgáltatásunk!
                Célunk egy olyan webes felület létrehozása, amivel vásárlóink eddiginél könnyebben képesek ételt
                rendelni.
            </p>
        </div>
        <div class="mx-2 my-2 lg:mx-6">
            <img class="border-4 rounded-md" src="http://localhost:8080/assets/foods/platter-2009590_1920.jpg"
                alt="Rustic plate">
        </div>
    </div>

</template>

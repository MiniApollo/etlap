<script setup lang="ts">

</script>

<template>
    <div class="grid grid-cols-2 max-md:grid-cols-1">
        <div class="mx-1 lg:mx-16 text-center">
            <h1 class="text-7xl lg:text-8xl my-12">Üdvözöljük étlapunkon</h1>
            <p class="m-4">Lorem, ipsum dolor sit amet consectetur adipisicing elit. Ad excepturi architecto beatae
                expedita
                nostrum, velit fugit atque sapiente consectetur similique fugiat enim quo dolorem sit? Id quidem tempora
                asperiores dolore?</p>
            <RouterLink
                class="p-2 bg-slate-200 border-2 rounded-3xl border-black text-black hover:text-xl transition-all duration-500"
                to="/etelek">Étlap megtekíntése</RouterLink>
        </div>
        <div class="mx-2 my-2 lg:mx-6">
            <img class="border-4 rounded-md" src="../assets/foods/restaurant-3955372_1920.jpg" alt="Rustic food">
        </div>
        <div class="mx-1 lg:mx-16 text-center">
            <h1 class="text-6xl lg:text-7xl mb-8">Megnyitottunk</h1>
            <p class="m-4">Lorem, ipsum dolor sit amet consectetur adipisicing elit. Ad excepturi architecto beatae
                expedita
                nostrum, velit fugit atque sapiente consectetur similique fugiat enim quo dolorem sit? Id quidem tempora
                asperiores dolore?</p>
        </div>
        <div class="mx-2 my-2 lg:mx-6">
            <img class="border-4 rounded-md" src="../assets/foods/platter-2009590_1920.jpg" alt="Rustic plate">
        </div>
    </div>

</template>

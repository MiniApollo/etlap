<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
    foods?: any[],
    statusMessage?: string
}>()

const isAddNewFood = ref(false);
let requestType: string = "POST";

const newFood = ref({
    EtelID: "",
    Nev: "",
    Leiras: "",
    KepPath: "",
    Ar: "",
});
function setFormToUpdate(food: any) {
    requestType = "PATCH";
    isAddNewFood.value = true;

    newFood.value.EtelID = food.EtelID;
    newFood.value.Nev = food.Nev;
    newFood.value.Leiras = food.Leiras;
    newFood.value.KepPath = food.KepPath;
    newFood.value.Ar = food.Ar;
}

function resetNewFood() {
    newFood.value.EtelID = "";
    newFood.value.Nev = "";
    newFood.value.Leiras = "";
    newFood.value.KepPath = "";
    newFood.value.Ar = "";

    requestType = "POST";
    isAddNewFood.value = false;
}

</script>
<template>
    <div>
        <h1 v-if="!isAddNewFood" class="my-4 px-5 text-5xl font-semibold">Ételek listája</h1>
        <button
            class="mx-4 p-3 bg-green-300 hover:bg-green-400 font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500"
            @click="isAddNewFood = true" v-if="!isAddNewFood">Új Étel hozzáadása</button>
        <button
            class="mx-4 py-3 px-12 max-md:px-16 max-md:w-1/2 my-2 bg-blue-300 hover:bg-blue-400 font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500"
            @click="resetNewFood();" v-else-if="isAddNewFood">Vissza</button>

        <div class="m-4" v-if="!isAddNewFood">
            <h2 class="text-4xl font-semibold bg-gray-300 rounded-3xl p-8"
                v-if="foods === undefined || foods.length == 0">
                Hiba történt <br>
                {{ statusMessage }}
            </h2>
            <ul class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3" v-else>
                <li v-for="food in foods" class="m-3 rounded-md border-black border-2 flex flex-col">
                    <img class="object-cover w-full flex-grow basis-2/3 rounded-md"
                        :src="'/assets/foods/' + food.KepPath" :alt="food.Nev + ' kép'">
                    <div class="m-3">
                        <h3 class="text-4xl font-semibold my-2">{{ food.Nev }}</h3>
                        <p class="lg:min-h-10">{{ food.Leiras }}</p>
                        <p class="my-4 text-2xl font-semibold">{{ food.Ar }} Ft</p>

                        <button @click="setFormToUpdate(food)"
                            class="p-2 bg-blue-300 hover:bg-blue-400 font-semibold border-2 rounded-2xl border-black text-black hover:scale-110 transition-all duration-500">Módosítás</button>
                        <button @click="$emit('deleteFood', food.EtelID)"
                            class="mx-4 p-2 bg-red-300 hover:bg-red-400 font-semibold border-2 rounded-2xl border-black text-black hover:scale-110 transition-all duration-500">Törlés</button>
                    </div>
                </li>
            </ul>
        </div>
        <form class="flex flex-col gap-2 p-16 bg-slate-300 rounded-3xl lg:w-1/2 mx-auto"
            @submit.prevent="$emit('newFoodSubmit', newFood, requestType, newFood.EtelID); resetNewFood()"
            v-else-if="isAddNewFood">

            <h1 class="font-semibold text-4xl mb-4">Étel Szerkesztése</h1>
            <label for="Name">Név: </label>
            <input class="p-1 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none"
                type="text" id="Name" v-model.trim="newFood.Nev" required autocomplete="off" maxlength="128">

            <label for="Leiras">Leírás: </label>
            <input class="p-1 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none"
                type="text" id="Leiras" v-model.trim="newFood.Leiras" autocomplete="off" maxlength="1024">

            <label for="KepPath">Kép útvonal: </label>
            <input class="p-1 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none"
                type="text" id="KepPath" v-model.trim="newFood.KepPath" autocomplete="off" maxlength="128">

            <label for="Ar">Ár: </label>
            <input class="p-1 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none"
                type="number" id="Ar" v-model.trim="newFood.Ar" required autocomplete="off" min="1" step="any">

            <input
                class="p-2 bg-green-300 hover:bg-green-400 cursor-pointer font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500 self-start"
                type="submit" value="Étel feltöltése">
        </form>
    </div>
</template>

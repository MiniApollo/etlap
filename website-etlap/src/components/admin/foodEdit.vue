<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
    foods?: any[],
    statusMessage?: String
}>()

const isAddNewFood = ref(false);
let requestType: String = "POST";

const newFood = ref({
    EtelID: "",
    Nev: "",
    Leiras: "",
    Kep: "",
    Ar: "",
});
function setFormToUpdate(food: any) {
    requestType = "PATCH";
    isAddNewFood.value = true;

    newFood.value.EtelID = food.EtelID;
    newFood.value.Nev = food.Nev;
    newFood.value.Leiras = food.Leiras;
    newFood.value.Kep = food.Kep;
    newFood.value.Ar = food.Ar;
}

function resetNewFood() {
    newFood.value.EtelID = "";
    newFood.value.Nev = "";
    newFood.value.Leiras = "";
    newFood.value.Kep = "";
    newFood.value.Ar = "";

    requestType = "POST";
    isAddNewFood.value = false;
}

</script>
<template>
    <div>
        <button @click="isAddNewFood = !isAddNewFood">
            <p v-if="!isAddNewFood">Új Étel hozzáadása</p>
            <p v-else-if="isAddNewFood">Vissza</p>
        </button>

        <div v-if="!isAddNewFood">
            <h1 class="text-3xl">Ételek Listája</h1>
            <h2 v-if="foods === undefined || foods.length == 0">
                Hiba történt: <br>
                {{ statusMessage }}
            </h2>
            <ul v-else>
                <li v-for="food in foods" class="m-2 p-4 h-96 rounded-md border-black border-2">
                    <h3 class="text-2xl">{{ food.Nev }}</h3>
                    <p>{{ food.Leiras }}</p>
                    <img :src="food.Kep" :alt="food.Nev + ' image'">
                    <p>{{ food.Ar }} Ft</p>
                    <button @click="setFormToUpdate(food)" class="border-2 border-black">Módosítás</button>
                    <button @click="$emit('deleteFood', food.EtelID)" class="border-2 border-black">Törlés</button>
                </li>
            </ul>
        </div>
        <form @submit.prevent="$emit('newFoodSubmit', newFood, requestType, newFood.EtelID); resetNewFood()"
            v-else-if="isAddNewFood">
            <label for="Name">Név: </label>
            <input type="text" id="Name" v-model="newFood.Nev" required autocomplete="off" maxlength="128">

            <label for="Leiras">Leírás: </label>
            <input type="text" id="Leiras" v-model="newFood.Leiras" autocomplete="off" maxlength="1024">

            <label for="Kep">Kép útvonal: </label>
            <input type="text" id="Kep" v-model="newFood.Kep" autocomplete="off" maxlength="128">

            <label for="Ar">Ár: </label>
            <input type="number" id="Ar" v-model="newFood.Ar" required autocomplete="off" min="1" step="any">

            <input type="submit" value="Rendelés Küldése">
        </form>
    </div>
</template>
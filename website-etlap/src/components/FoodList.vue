<script setup lang="ts">
import { onMounted, ref } from 'vue';

const foods: any = ref([]);
const statusHeader = ref();
const status = ref();

async function getAllFood() {
    statusHeader.value = "";
    status.value = "";

    fetch("http://localhost:8080/food")
        .then(res => res.json())
        .then(data => {
            for (let i = 0; i < data.length; i++) {
                let food: any = { EtelID: data[i].EtelID, Nev: data[i].Nev, Leiras: data[i].Leiras, KepPath: data[i].KepPath, Ar: data[i].Ar };
                foods.value.push(food);
            }
        })
        .catch(err => {
            console.log(err);
            statusHeader.value = "Hiba történt:";
            status.value = err;
        });
}

onMounted(() => {
    getAllFood();
});
</script>

<template>
    <!-- Fix warnings:
     https://stackoverflow.com/questions/68803137/vue-3-passing-array-warning-extraneous-non-props-attributes-were-passed-to-comp-->
    <div class="m-auto">
        <h1 class="px-5 text-5xl font-semibold">Kinálatunk</h1>
        <div class="bg-gray-300 rounded-3xl text-center mx-3 my-9 p-16 md:w-3/4 lg:w-1/2 md:mx-auto"
            v-if="foods === undefined || foods.length == 0">
            <h2 class="m-10 text-4xl font-semibold">{{ statusHeader }}</h2>
            <h2 class="text-2xl">{{ status }}</h2>
        </div>
        <ul v-else class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <li v-for="food in foods" class="m-3 rounded-md border-black border-2 flex flex-col">
                <img class="object-cover w-full flex-grow basis-2/3 rounded-md"
                    :src="'http://localhost:8080/assets/foods/' + food.KepPath" :alt="food.Nev + ' kép'">
                <div class="m-3">
                    <h3 class="text-4xl font-semibold my-2">{{ food.Nev }}</h3>
                    <p class="lg:min-h-10">{{ food.Leiras }}</p>
                    <p class="my-4 text-2xl font-semibold">{{ food.Ar }} Ft</p>
                    <button
                        class="p-2 bg-slate-300 active:bg-slate-400 font-semibold border-2 rounded-2xl border-black text-black hover:scale-110 transition-all duration-500"
                        @click="$emit('kosarhoz', food)"> Kosárba</button>
                </div>
            </li>
        </ul>
    </div>
</template>

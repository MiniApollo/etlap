<script setup lang="ts">
import { onMounted, ref } from 'vue';

const foods: any = ref([]);
const error: any = ref();

async function getAllFood() {
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
            error.value = err
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
        <h1 class="px-5 text-3xl">Kinálatunk</h1>
        <h2 v-if="foods === undefined || foods.length == 0">
            Hiba történt: <br>
            {{ error }}
        </h2>
        <ul v-else class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <li v-for="food in foods" class="m-3 rounded-md border-black border-2 flex flex-col">
                <img class="object-cover w-full flex-grow basis-2/3 rounded-md"
                    :src="'http://localhost:8080/assets/foods/' + food.KepPath" :alt="food.Nev + ' kép'">
                <div class="m-3">
                    <h3 class="text-4xl font-semibold my-2">{{ food.Nev }}</h3>
                    <p class="lg:min-h-10">{{ food.Leiras }}</p>
                    <p class="my-4 text-2xl font-semibold">{{ food.Ar }} Ft</p>
                    <button
                        class="p-2 bg-slate-300 font-semibold border-2 rounded-2xl border-black text-black hover:scale-110 transition-all duration-500"
                        @click="$emit('kosarhoz', food)"> Kosárhoz</button>
                </div>
            </li>
        </ul>
    </div>
</template>

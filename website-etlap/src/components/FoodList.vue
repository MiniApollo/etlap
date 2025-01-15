<script setup lang="ts">
import { onMounted, ref } from 'vue';

const foods: any = ref([]);
const error: any = ref();

async function getAllFood() {
    fetch("http://localhost:8080/food")
        .then(res => res.json())
        .then(data => {
            for (let i = 0; i < data.length; i++) {
                let food = { EtelID: data[i].EtelID, Nev: data[i].Nev, Leiras: data[i].Leiras, Kep: data[i].Kep, Ar: data[i].Ar };
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
    <div class="flex justify-evenly items-center flex-col ">
        <h1 class="text-3xl">Ételek Listája</h1>
        <h2 v-if="foods === undefined || foods.length == 0">
            Hiba történt: <br>
            {{ error }}
        </h2>
        <ul v-else>
            <!-- :style="{ 'background-image': 'url(' + food.Kep + ')'}" -->
            <li v-for="food in foods" class="m-2 p-4 h-96 rounded-md border-black border-2">
                <h3 class="text-2xl">{{ food.Nev }}</h3>
                <p>{{ food.Leiras }}</p>
                <img :src="food.Kep" :alt="food.Nev + ' image'">
                <p>{{ food.Ar }} Ft</p>
                <button class="border-2 border-black" @click="$emit('kosarhoz', food)"> Kosárhoz</button>
            </li>
        </ul>
    </div>
</template>

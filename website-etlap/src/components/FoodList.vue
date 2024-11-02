<script setup lang="ts">
import { onMounted, ref } from 'vue';

const foods: any = ref([]);

async function getAllFood() {
    // TODO: Add not found error if web server is not on
    fetch("http://localhost:8080/food")
        .then(res => res.json())
        .then(data => {
            for (let i = 0; i < data.length; i++) {
                let food = { EtelID: data[i].EtelID, Nev: data[i].Nev, Leiras: data[i].Leiras, Kep: data[i].Kep, Ar: data[i].Ar };
                foods.value.push(food);
            }
        })
        .catch(err => console.log(err));
}

onMounted(() => {
    getAllFood();
});

</script>

<template>
    <!-- Fix warnings: 
     https://stackoverflow.com/questions/68803137/vue-3-passing-array-warning-extraneous-non-props-attributes-were-passed-to-comp-->
    <div class="text-center">
        <h1 class="text-3xl">Ételek Listája</h1>
        <ul>
            <li class="my-2" v-for="food in foods">
                <h3 class="text-2xl">{{ food.Nev }}</h3>
                <p>{{ food.Leiras }}</p>
                {{ food.Kep }}
                <p>{{ food.Ar }} Ft</p>
                <button class="border-2 border-black" @click="$emit('kosarhoz', food)"> Kosárhoz</button>
            </li>
        </ul>
    </div>
</template>
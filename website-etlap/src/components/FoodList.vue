<script setup lang="ts">
import { onMounted, ref } from 'vue';

const foods: any = ref([]);

async function getAllFood() {
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
    <h1 class="text-3xl">Ételek Listája</h1>
    <ul>
        <li v-for="food in foods">
            <h3 class="text-2xl">{{ food.Nev }}</h3>
            <p>{{ food.Leiras }}</p>
            {{ food.Kep }}
            <p>{{ food.Ar }} Ft</p>
        </li>
    </ul>

</template>
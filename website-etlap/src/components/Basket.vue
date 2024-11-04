<script setup lang="ts">
import { ref } from 'vue';

// https://vuejs.org/guide/components/props
defineProps<{
    basketContent?: any[]
}>()
const isSubmiting = ref(false);

</script>

<template>
    <!-- Fix warnings: 
     https://stackoverflow.com/questions/68803137/vue-3-passing-array-warning-extraneous-non-props-attributes-were-passed-to-comp-->
    <div>
        <h1 class="text-3xl">Rendelés</h1>
        <h2 v-if="basketContent === undefined || basketContent.length == 0">
            A Kosarad Üres
        </h2>
        <ul v-else-if="!isSubmiting">
            <button @click="isSubmiting = true">Tovább a leadáshoz</button>
            <li class="my-2" v-for="(food, index) in basketContent">
                <h3 class="text-2xl">{{ food.Nev }}</h3>
                <p>{{ food.Leiras }}</p>
                {{ food.Kep }}
                <p>{{ food.Ar }} Ft</p>
                <button @click="$emit('torol', index)">Töröl</button>
            </li>
        </ul>
        <form v-else-if="isSubmiting" action="http://localhost:8080/order" method="post">
            <label for="Name">Név: </label>
            <input type="text" id="Name" required autocomplete="off">

            <label for="Email">Email Cím: </label>
            <input type="email" id="Email" required autocomplete="off">

            <label for="PhoneNumber">Telefonszám: </label>
            <input type="number" id="PhoneNumber" required autocomplete="off">

            <input type="submit" value="Rendelés Küldése">
        </form>
    </div>
</template>

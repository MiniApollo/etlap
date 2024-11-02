<script setup lang="ts">
// https://vuejs.org/guide/components/props
defineProps<{
    basket?: any[]
}>()

</script>

<template>
    <!-- Fix warnings: 
     https://stackoverflow.com/questions/68803137/vue-3-passing-array-warning-extraneous-non-props-attributes-were-passed-to-comp-->
    <div>
        <h1 class="text-3xl">Rendelés</h1>
        <ul>
            <li class="my-2" v-for="(food, index) in basket">
                <h3 class="text-2xl">{{ food.Nev }}</h3>
                <p>{{ food.Leiras }}</p>
                {{ food.Kep }}
                <p>{{ food.Ar }} Ft</p>
                <button @click="$emit('torol', index)">Töröl</button>
            </li>
        </ul>
    </div>
</template>

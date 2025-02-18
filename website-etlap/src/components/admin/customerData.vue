<script setup lang="ts">
const props = defineProps<{
    customer?: any,
    index?: any,
    foodsByCustomer: any[],
    showCurrentOrders?: boolean
}>()

</script>
<template>
    <div>
        <ul class="flex gap-4 max-lg:gap-1 max-lg:flex-col md:justify-between">
            <li class="basis-1/6 text-xl max-md:text-2xl font-semibold">{{ customer.Nev }}</li>
            <li class="basis-1/5 text-xl overflow-auto">{{ customer.Email }}</li>
            <li class="basis-1/5 text-xl overflow-auto">+{{ customer.Telefonszam }}</li>
            <li class="text-xl overflow-auto">ID: {{ customer.VasarloID }}</li>

            <li class="basis-1/5 text-2xl">{{ customer.LeadasiIdo }}</li>
            <button
                class="py-2 px-8 bg-blue-300 hover:bg-blue-400 cursor-pointer font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500 self-start"
                @click="$emit('orderComplete', customer, showCurrentOrders)">
                <p v-if="showCurrentOrders">Kész</p>
                <p v-if="!showCurrentOrders">Visszavon</p>
            </button>
        </ul>
        <ul>
            <h2 class="mb-3 text-2xl font-semibold">Rendelés:</h2>
            <li v-for="food in foodsByCustomer[index]">
                <ul>
                    <li class="my-1 text-2xl inline m-1">{{ food.Nev }}</li>
                    <li class="my-1 text-2xl inline m-1">{{ food.Ar }} Ft</li>
                </ul>
            </li>
            <li class="my-2 text-2xl font-semibold">Összesen: {{ customer.Osszeg.Float64}} Ft</li>
        </ul>
    </div>
</template>
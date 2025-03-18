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
        <ul class="max-lg:mb-4 mb-2 flex gap-4 max-lg:gap-1 max-lg:flex-col md:justify-between">
            <li class="basis-1/6 text-xl max-md:text-2xl font-semibold overflow-auto">Név: <br class="max-lg:hidden">{{
                customer.Nev }}</li>
            <li class="basis-1/5 text-xl overflow-auto">Email: <br class="max-lg:hidden">{{ customer.Email }}</li>
            <li class="basis-1/5 flex-grow-0 text-xl overflow-auto">TelefonSzám: <br class="max-lg:hidden"> +{{
                customer.Telefonszam }}</li>
            <li class="text-xl overflow-auto px-1">ID: <br class="max-lg:hidden">{{ customer.VasarloID }}</li>

            <li class="basis-1/5 text-2xl">Leadás: <br class="max-lg:hidden">{{ customer.LeadasiIdo }}</li>
            <li v-if="customer.ElkeszultIdo.Valid" class="basis-1/5 text-2xl">Elkészült: <br class="max-lg:hidden">{{
                customer.ElkeszultIdo.String }}</li>
            <button
                class="py-2 px-8 bg-blue-300 hover:bg-blue-400 cursor-pointer font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500 self-start"
                @click="$emit('orderComplete', customer, showCurrentOrders)">
                <p v-if="showCurrentOrders">Kész</p>
                <p v-if="!showCurrentOrders">Visszavon</p>
            </button>
        </ul>
        <ul>
            <h2 class="mb-1 text-2xl font-semibold">Rendelés:</h2>
            <li v-for="foodWithVolume in foodsByCustomer[index]">
                <ul>
                    <li class="my-1 text-2xl inline m-1">{{ foodWithVolume.Volume }} {{ foodWithVolume.Food.Nev }}</li>
                    <li class="my-1 text-2xl inline m-1">{{ foodWithVolume.Food.Ar * foodWithVolume.Volume }} Ft</li>
                </ul>
            </li>
            <li class="my-2 text-2xl font-semibold">Összesen: {{ customer.Osszeg.Float64 }} Ft</li>
        </ul>
    </div>
</template>

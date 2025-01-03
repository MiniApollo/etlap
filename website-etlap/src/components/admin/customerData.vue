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
        <ul>
            <li class="inline m-1">{{ customer.Nev }}</li>
            <li class="inline m-1">{{ customer.VasarloID }}</li>
            <li class="inline m-1">{{ customer.Email }}</li>
            <li class="inline m-1">{{ customer.Telefonszam }}</li>
            <li class="inline m-1">{{ customer.LeadasiIdo }}</li>
            <button @click="$emit('orderComplete', customer, showCurrentOrders)">
                <p v-if="showCurrentOrders">Kész</p>
                <p v-if="!showCurrentOrders">Visszavon</p>
            </button>
        </ul>
        <ul>
            <li v-for="food in foodsByCustomer[index]">
                <ul>
                    <li class="inline m-1">{{ food.Nev }}</li>
                    <li class="inline m-1">{{ food.Ar }}</li>
                </ul>
            </li>
        </ul>
    </div>
</template>
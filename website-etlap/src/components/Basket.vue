<script setup lang="ts">
import { ref } from 'vue';

// https://vuejs.org/guide/components/props
const props = defineProps<{
    basketContent?: any[]
}>()
const emit = defineEmits(["torol", "emptyBasket"]);

// Important to use the restapi json names
const newCustomer = ref({
    Nev : "",
    Email: "",
    Telefonszam: "",
});

function sendOrder() {
    // POST
    fetch("http://localhost:8080/order", {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        },
        // Mixed Json order
        // https://stackoverflow.com/questions/3948206/json-order-mixed-up
        // https://stackoverflow.com/questions/17229418/jsonobject-why-jsonobject-changing-the-order-of-attributes
        body: JSON.stringify({
            Customer: newCustomer.value,
            Foods: props.basketContent
        })
    })

    emit("emptyBasket");
    newCustomer.value = {
        Nev: "",
        Email: "",
        Telefonszam: "",
    };
    alert("Sikeres beküldés");
}

</script>

<template>
    <!-- Fix warnings:
     https://stackoverflow.com/questions/68803137/vue-3-passing-array-warning-extraneous-non-props-attributes-were-passed-to-comp-->
    <div>
        <h1 class="px-5 text-3xl">Rendelés</h1>
        <h2 v-if="basketContent === undefined || basketContent.length == 0">
            A Kosarad Üres
        </h2>
        <div v-else class="flex">
            <ul class="flex flex-col basis-3/5">
                <li class="m-3 rounded-md border-black border-2 flex max-sm:flex-col" v-for="(food, index) in basketContent">
                    <img class="w-full basis-1/2 rounded-md"
                        :src="'http://localhost:8080/assets/foods/' + food.KepPath" :alt="food.Nev + ' kép'">
                    <div class="m-3 basis-1/2">
                        <h3 class="text-4xl font-semibold my-2">{{ food.Nev }}</h3>
                        <p class="lg:min-h-10">{{ food.Leiras }}</p>
                        <p class="my-4 text-2xl font-semibold">{{ food.Ar }} Ft</p>
                        <button
                            class="p-2 bg-slate-300 font-semibold border-2 rounded-2xl border-black text-black hover:scale-110 transition-all duration-500"
                            @click="$emit('torol', index)">Törlés</button>
                    </div>
                </li>
            </ul>
            <form class="flex flex-col p-3" @submit.prevent="sendOrder">
                <button
                    class="p-2 bg-slate-300 font-semibold border-2 rounded-2xl border-black text-black hover:scale-110 transition-all duration-500"
                    @click="$emit('emptyBasket')">Kosár kiürítése</button>

                <label for="Name">Név: </label>
                <input type="text" id="Name" v-model="newCustomer.Nev" required autocomplete="off" maxlength="128">

                <label for="Email">Email Cím: </label>
                <input type="email" id="Email" v-model="newCustomer.Email" required autocomplete="off" maxlength="128">

                <label for="PhoneNumber">Telefonszám: </label>
                <input type="text" id="PhoneNumber" v-model="newCustomer.Telefonszam" required autocomplete="off"
                    maxlength="32">

                <input
                    class="p-2 bg-slate-300 font-semibold border-2 rounded-2xl border-black text-black hover:scale-110 transition-all duration-500"
                    type="submit" value="Rendelés Küldése">
            </form>
        </div>
    </div>
</template>

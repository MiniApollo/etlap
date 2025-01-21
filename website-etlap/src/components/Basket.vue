<script setup lang="ts">
import { ref } from 'vue';

// https://vuejs.org/guide/components/props
const props = defineProps<{
    basketContent?: any[]
}>()
const emit = defineEmits(["torol", "emptyBasket"]);

const isSubmiting = ref(false);
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
    isSubmiting.value = false;
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
        <h1 class="text-3xl">Rendelés</h1>
        <h2 v-if="basketContent === undefined || basketContent.length == 0">
            A Kosarad Üres
        </h2>
        <ul v-else-if="!isSubmiting">
            <button @click="isSubmiting = true">Tovább a leadáshoz</button>
            <button @click="$emit('emptyBasket')">Kosár kiürítése</button>
            <li class="my-2" v-for="(food, index) in basketContent">
                <h3 class="text-2xl">{{ food.Nev }}</h3>
                <p>{{ food.Leiras }}</p>
                {{ food.KepPath }}
                <p>{{ food.Ar }} Ft</p>
                <button @click="$emit('torol', index)">Töröl</button>
            </li>
        </ul>
        <form v-else-if="isSubmiting" @submit.prevent="sendOrder">
            <button @click="isSubmiting = false">Vissza</button>

            <label for="Name">Név: </label>
            <input type="text" id="Name" v-model="newCustomer.Nev" required autocomplete="off" maxlength="128">

            <label for="Email">Email Cím: </label>
            <input type="email" id="Email" v-model="newCustomer.Email" required autocomplete="off" maxlength="128">

            <label for="PhoneNumber">Telefonszám: </label>
            <input type="text" id="PhoneNumber" v-model="newCustomer.Telefonszam" required autocomplete="off" maxlength="32">

            <input type="submit" value="Rendelés Küldése">
        </form>
    </div>
</template>

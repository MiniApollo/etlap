<script setup lang="ts">
import { computed, ref } from 'vue';
import { RouterLink } from 'vue-router';

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

const sumPrice = computed(() => {
    if (!props.basketContent) {
        return 0;
    }
    let sum = 0;
    for (let i = 0; i < props.basketContent.length; i++) {
        sum += props.basketContent[i].Ar;
    }
    return sum;
})

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
        <div v-if="basketContent === undefined || basketContent.length == 0">
            <h1 class="px-5 text-5xl font-semibold text-center">Rendelés</h1>
            <div class="bg-gray-300 rounded-3xl text-center mx-3 my-9 p-16 md:w-3/4 lg:w-1/2 md:mx-auto">
                <h1 class="m-10 text-2xl font-semibold">A Kosarad Üres</h1>
                <RouterLink
                    class="m-2 p-4 font-semibold bg-slate-300 border-2 rounded-3xl border-black text-black hover:text-xl transition-all duration-500"
                    to="/etelek">Rendelj valamit</RouterLink>
            </div>
        </div>
        <div v-else>
            <h1 class="px-5 text-5xl font-semibold">Rendelés</h1>
            <div class="flex max-lg:flex-col-reverse mx-3 my-1">
                <ul class="flex flex-col basis-3/5">
                    <button
                        class="mx-3 mt-3 p-2 bg-red-300 hover:bg-red-400 font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500 self-start"
                        @click="$emit('emptyBasket')">Kosár kiürítése</button>
                    <li class="m-3 rounded-md border-black border-2 flex max-sm:flex-col"
                        v-for="(food, index) in basketContent">
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
                <form class="flex flex-col gap-2 p-16 bg-slate-300 rounded-3xl basis-2/5 lg:self-start"
                    @submit.prevent="sendOrder">

                    <h1 class="font-semibold text-4xl mb-4">Rendelési adatok</h1>

                    <label for="Name">Név:</label>
                    <input class="p-1 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none" type="text" id="Name" v-model="newCustomer.Nev" required
                        autocomplete="off" maxlength="128">

                    <label for="Email">Email Cím:</label>
                    <input class="p-1 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none" type="email" id="Email" placeholder="name@example.com"
                        v-model="newCustomer.Email" required autocomplete="off" maxlength="128">

                    <label for="PhoneNumber">Telefonszám:</label>
                    <input class="p-1 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none" type="text" id="PhoneNumber"
                        v-model="newCustomer.Telefonszam" required autocomplete="off" maxlength="32">

                    <h1 class="font-semibold text-lg mb-4">Összesen: {{ sumPrice }} Ft</h1>

                    <input
                        class="p-2 bg-green-300 hover:bg-green-400 cursor-pointer font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500 self-start"
                        type="submit" value="Rendelés Küldése">
                </form>
            </div>
        </div>
    </div>
</template>

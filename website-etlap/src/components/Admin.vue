<script setup lang="ts">
import { onMounted, ref } from 'vue';

const password = ref("");
const isLoggedIn = ref();
const responseMessage = ref("");

const orders: any = ref([]);
const customers: any = ref([]);
const foodsByCustomer: any = ref([]);

const statusMessage = ref("");
const statusError = ref(false);
const showCurrentOrders = ref(true);

async function getData() {
    orders.value = await getWithToken("order");
    customers.value = await getWithToken("order/customer");
    if (!Array.isArray(customers.value)) {
        statusMessage.value = "Nincs új rendelés";
        statusError.value = true;
        return
    }

    for (let i = 0; i < customers.value.length; i++) {
        foodsByCustomer.value.push(await getWithToken("order/food/" + customers.value[i].VasarloID));
    }
    console.log(customers.value);
    console.log(foodsByCustomer.value);
}

async function sendPassword() {
    const response = await fetch("http://localhost:8080/admin", {
        method: "GET",
        headers: {
            'Authorization': 'Bearer ' + btoa(password.value),
            "Content-type": "application/json; charset=UTF-8"
        },
    })
    password.value = "";

    if (!response.ok) {
        console.error(new Error(`Response status: ${response.status}`));
    }

    const accessToken = await response.json();
    if (accessToken.token) {
        console.log("Token: ", accessToken.token);
        sessionStorage.setItem("adminToken", accessToken.token);

        responseMessage.value = "Sikeres bejelentkezés";
        isLoggedIn.value = true;
        getData();
        return;
    }
    console.log("Sikertelen bejelentkezés");
    responseMessage.value = "Sikertelen bejelentkezés";
    isLoggedIn.value = false;
}

async function getWithToken(url: String) {
    // https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
    const response = await fetch("http://localhost:8080/" + url, {
        method: "GET",
        headers: {
            'Authorization': 'Bearer ' + btoa(sessionStorage.getItem("adminToken") || "{}"),
            "Content-type": "application/json; charset=UTF-8"
        },
    })
    if (!response.ok) {
        console.error(new Error(`Response status: ${response.status}`));
    }

    const json = await response.json();
    return json;
}

async function orderDone(customer: any) {
    await fetch("http://localhost:8080/customer/" + customer.VasarloID, {
        method: "PATCH",
        headers: {
            'Authorization': 'Bearer ' + btoa(sessionStorage.getItem("adminToken") || "{}"),
            "Content-type": "application/json; charset=UTF-8"
        },
    })
    getData();
}

function signOut() {
    isLoggedIn.value = false;
    // Remove all saved data from sessionStorage
    sessionStorage.clear();
}

onMounted(() => {
    isLoggedIn.value = false;
    if (sessionStorage.getItem("adminToken")) {
        isLoggedIn.value = true;
        getData();
    }
});
</script>
<template>
    <div>
        <h3>Admin</h3>
        <form v-if="!isLoggedIn" @submit.prevent="sendPassword">
            <label for="adminPassword">Jelszó:</label>
            <input type="password" id="adminPassword" v-model="password" required>
            <input type="submit" value="Bejelentkezés">
            <h3>{{ responseMessage }}</h3>
        </form>
        <div v-else-if="isLoggedIn">
            <button @click="signOut">Kijelentkezés</button>

            <h3>Rendelések:</h3>
            <h3 v-if="statusError">{{ statusMessage }}</h3>

            <div v-else-if="!statusError">
                <ul>
                    <button @click="showCurrentOrders = !showCurrentOrders">Kész Rendelések mutatása</button>

                    <li v-for="(customer, index) in customers" class="m-1 p-1 border border-black">
                        <ul>
                            <li class="inline m-1">{{ customer.Nev }}</li>
                            <li class="inline m-1">{{ customer.VasarloID }}</li>
                            <li class="inline m-1">{{ customer.Email }}</li>
                            <li class="inline m-1">{{ customer.Telefonszam }}</li>
                            <li class="inline m-1">{{ customer.LeadasiIdo }}</li>
                            <button @click="orderDone(customer)">Kész</button>
                        </ul>
                        <ul>
                            <li v-for="food in foodsByCustomer[index]">
                                <ul>
                                    <li class="inline m-1">{{ food.Nev }}</li>
                                    <li class="inline m-1">{{ food.Ar }}</li>
                                </ul>
                            </li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>
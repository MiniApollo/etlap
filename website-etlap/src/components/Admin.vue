<script setup lang="ts">
import { onMounted, ref } from 'vue';

const password: any = ref("");
const isLoggedIn = ref();
const responseMessage = ref("");

const orders: any = ref([]);
const customers: any = ref([]);

function sendPassword() {
    fetch("http://localhost:8080/admin", {
        method: "GET",
        headers: {
            'Authorization': 'Bearer ' + btoa(password.value),
            "Content-type": "application/json; charset=UTF-8"
        },
    })
        .then(res => res.json())
        .then(accessToken => {
            if (accessToken.token) {
                console.log("Token: ", accessToken.token);
                sessionStorage.setItem("adminToken", accessToken.token);

                responseMessage.value = "Sikeres bejelentkezés";
                isLoggedIn.value = true;
                return;
            }
            console.log("Sikertelen bejelentkezés");
            responseMessage.value = "Sikertelen bejelentkezés";
            isLoggedIn.value = false;
        })
}
async function getData() {
    // TODO: Get all customer and food data from database with joins
    orders.value = await getWithToken("order");
    customers.value = await getWithToken("customer");
}

async function getWithToken(url: String, Id: String = "") {
    // https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
    const response = await fetch("http://localhost:8080/" + url + "/" + Id, {
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
            <h3>{{ responseMessage }}</h3>
            <button @click="signOut">Kijelentkezés</button>
            <h3>Rendelések:</h3>
            <ul>
                <li v-for="order in orders">
                    <p>{{ order.VasarloID }}</p>
                </li>
            </ul>
        </div>
    </div>
</template>
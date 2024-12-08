<script setup lang="ts">
import { onMounted, ref } from 'vue';

const password: any = ref("");
const isLoggedIn = ref();

function sendPassword() {
    fetch("http://localhost:8080/admin", {
        method: "GET",
        headers: {
            'Authorization': 'Basic ' + btoa(password.value),
            "Content-type": "application/json; charset=UTF-8"
        },
    })
        .then(res => res.json())
        .then(accessToken => {
            if (accessToken.token) {
                console.log("Token: ", accessToken.token);
                sessionStorage.setItem("adminToken", accessToken.token);
                isLoggedIn.value = true;
                return;
            }
            console.log("Failed to authenticate");
            isLoggedIn.value = false;
        })
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
        </form>
        <div v-else-if="isLoggedIn">
            <h3>Sikeres Bejelentkezés</h3>
            <button @click="signOut">Kijelentkezés</button>
        </div>
    </div>
</template>
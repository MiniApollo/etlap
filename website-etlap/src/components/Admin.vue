<script setup lang="ts">
import { ref } from 'vue';

const password: any = ref("");

function sendPassword() {
    fetch("http://localhost:8080/admin", {
        method: "GET",
        headers: {
            'Authorization': 'Basic ' + btoa(password.value),
            "Content-type": "application/json; charset=UTF-8"
        },
    })
}
</script>
<template>
    <div>
        <h3>Admin</h3>
        <form @submit.prevent="sendPassword">
            <label for="adminPassword">Jelszó:</label>
            <input type="password" id="adminPassword" v-model="password" required>
            <input type="submit" value="Bejelentkezés">
        </form>
    </div>
</template>
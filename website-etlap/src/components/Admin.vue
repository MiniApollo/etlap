<script setup lang="ts">
import { onMounted, ref } from 'vue';
import FoodEdit from './admin/foodEdit.vue';
import CustomerData from './admin/customerData.vue';

const password = ref("");
const isLoggedIn = ref(false);
const responseMessage = ref("");

const orders: any = ref([]);
const customers: any = ref([]);
const foodsByCustomer: any = ref([]);
const foods: any = ref([]);

const statusMessage = ref("");
const statusError = ref(false);
const showCurrentOrders = ref(true);
const isShowOrder = ref(true);

async function getData() {
    orders.value = await getWithToken("order");
    foods.value = await getWithToken("food");

    if (showCurrentOrders.value) {
        customers.value = await getWithToken("order/customer?isDone=false");
    }
    else {
        customers.value = await getWithToken("order/customer?isDone=True");
    }

    if (!Array.isArray(customers.value)) {
        statusMessage.value = "Nincs új rendelés";
        statusError.value = true;
        return
    }
    statusMessage.value = "";
    statusError.value = false;

    for (let i = 0; i < customers.value.length; i++) {
        foodsByCustomer.value.push(await getWithToken("order/food/" + customers.value[i].VasarloID));
    }
    //console.log(customers.value);
    //console.log(foodsByCustomer.value);
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

async function orderDone(customer: any, isDone: boolean) {

    await fetch("http://localhost:8080/customer/" + customer.VasarloID, {
        method: "PATCH",
        headers: {
            'Authorization': 'Bearer ' + btoa(sessionStorage.getItem("adminToken") || "{}"),
            "Content-type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({
            Elkeszult: isDone,
        })
    })
    getData();
}

async function postNewFood(food: any, requestType: string = "POST", etelID: string = "") {

    // /food and /food/ not same api endpoint
    // automatic redirect backend
    // https://stackoverflow.com/questions/61547014/restful-uri-trailing-slash-or-no-trailing-slash 
    if (etelID) {
        etelID = "/" + etelID;
    }
    await fetch("http://localhost:8080/food" + etelID, {
        method: requestType,
        headers: {
            'Authorization': 'Bearer ' + btoa(sessionStorage.getItem("adminToken") || "{}"),
            "Content-type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({
            Nev: food.Nev,
            Leiras: food.Leiras,
            Kep: food.Kep,
            Ar: food.Ar,
        })
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
            <br>
            <button @click="isShowOrder = !isShowOrder; getData()">
                <p v-if="!isShowOrder">Rendelések megjelenítése</p>
                <p v-else-if="isShowOrder">Ételek szerkesztésének megjelenítése</p>
            </button>

            <div v-if="isShowOrder">
                <button @click="showCurrentOrders = !showCurrentOrders; getData()">
                    <p v-if="showCurrentOrders">Korábbi rendelések megjelenítése</p>
                    <p v-else-if="!showCurrentOrders">Jelenlegi rendelések megjelenítése</p>
                </button>

                <h3 v-if="showCurrentOrders">Jelenlegi Rendelések:</h3>
                <h3 v-else-if="!showCurrentOrders">Korrábbi Rendelések:</h3>
                <ul>
                    <h3 v-if="statusError && showCurrentOrders">{{ statusMessage }}</h3>
                    <li v-if="Array.isArray(customers)" v-for="(customer, index) in customers"
                        class="m-1 p-1 border border-black">
                        <CustomerData @orderComplete="orderDone" :customer="customer"
                            :foods-by-customer="foodsByCustomer" :show-current-orders="showCurrentOrders" :index="index" />
                    </li>
                </ul>
            </div>

            <div v-else-if="!isShowOrder">
                <FoodEdit @newFoodSubmit="postNewFood" :foods="foods" :status-message="statusMessage" />
            </div>

        </div>
    </div>
</template>
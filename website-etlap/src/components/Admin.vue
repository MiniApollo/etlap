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
        statusMessage.value = "Nincs rendelés";
        statusError.value = true;
        return
    }
    statusMessage.value = "";
    statusError.value = false;

    // Fix bug, wrong food orders per customer
    // Need to empty foodByCustomerList so foods don't get duplicated
    foodsByCustomer.value = [];

    for (let i = 0; i < customers.value.length; i++) {
        foodsByCustomer.value.push(await getWithToken("order/food/" + customers.value[i].VasarloID));
    }
}

async function sendPassword() {
    const response = await fetch("http://localhost:8080/auth", {
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
    if (!confirm("Biztosan feltöltöd?")) {
        return;
    }
    const response = await fetch("http://localhost:8080/food" + etelID, {
        method: requestType,
        headers: {
            'Authorization': 'Bearer ' + btoa(sessionStorage.getItem("adminToken") || "{}"),
            "Content-type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({
            EtelID: Number(food.EtelID),
            Nev: food.Nev,
            Leiras: food.Leiras,
            KepPath: food.KepPath,
            Ar: food.Ar,
        })
    })
    if (!response.ok) {
        alert("Sikertelen módosítás");
        return;
    }
    getData();
}

async function deleteFood(EtelID: string) {
    if (!confirm("Biztosan törlöd?")) {
        return;
    }
    const response = await fetch("http://localhost:8080/food/" + EtelID, {
        method: "DELETE",
        headers: {
            'Authorization': 'Bearer ' + btoa(sessionStorage.getItem("adminToken") || "{}"),
            "Content-type": "application/json; charset=UTF-8"
        }
    })
    if (!response.ok) {
        alert("Sikertelen törlés");
        return;
    }
    getData();
}

function signOut() {
    if (!confirm("Biztosan kijelenkezel?")) {
        return;
    }
    isLoggedIn.value = false;
    responseMessage.value = ""
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
        <div class="flex flex-col justify-center" v-if="!isLoggedIn" @submit.prevent="sendPassword">
            <h1 class="px-5 text-5xl font-semibold text-center">Admin</h1>
            <form class="bg-gray-300 rounded-3xl mx-3 my-8 p-16 self-center flex flex-col gap-2">
                <h2 class="font-semibold text-3xl mb-4">Bejelentkezés</h2>
                <label class="text-lg" for="adminPassword">Jelszó:</label>
                <input class="p-2 rounded-lg shadow-sm border-2 border-gray-500 focus:border-blue-500 outline-none"
                    type="password" id="adminPassword" v-model="password" required>
                <input
                    class="my-3 p-2 bg-blue-300 hover:bg-blue-400 cursor-pointer font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500 self-center"
                    type="submit" value="Bejelentkezés">
                <p class="text-lg text-center font-semibold">{{ responseMessage }}</p>
            </form>
        </div>

        <div v-else-if="isLoggedIn">
            <h1 class="px-5 text-6xl font-semibold">Admin</h1>
            <nav class="flex max-sm:flex-col mx-4 gap-4">
                <button
                    class="my-3 py-2 px-5 max-md:px-10 bg-red-300 hover:bg-red-400 cursor-pointer font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500 self-center max-md:self-start"
                    @click="signOut">Kijelentkezés</button>
                <button class="max-sm:self-start" @click="isShowOrder = !isShowOrder; getData()">
                    <p class="p-2 bg-blue-300 hover:bg-blue-400 font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500"
                        v-if="!isShowOrder">Rendelések megjelenítése</p>
                    <p class="p-2 bg-blue-300 hover:bg-blue-400 font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500"
                        v-else-if="isShowOrder">Ételek szerkesztésének megjelenítése</p>
                </button>
            </nav>

            <div v-if="isShowOrder">
                <button class="m-4" @click="showCurrentOrders = !showCurrentOrders; getData()">
                    <p class="p-2 bg-blue-300 hover:bg-blue-400 font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500"
                        v-if="showCurrentOrders">Korábbi rendelések megjelenítése</p>
                    <p class="p-2 bg-blue-300 hover:bg-blue-400 font-semibold border-2 rounded-2xl border-black text-black transition-all duration-500"
                        v-else-if="!showCurrentOrders">Jelenlegi rendelések megjelenítése</p>
                </button>
                <h1 class="my-4 px-5 text-5xl font-semibold" v-if="showCurrentOrders">Jelenlegi rendelések: {{
                    customers.length }}</h1>
                <h1 class="my-4 px-5 text-5xl font-semibold" v-else-if="!showCurrentOrders">Korábbi rendelések: {{
                    customers.length }}</h1>
                <ul class="m-4">
                    <h3 class="text-4xl font-semibold bg-gray-300 rounded-3xl p-8" v-if="statusError">{{ statusMessage
                        }}</h3>
                    <li v-if="Array.isArray(customers)" v-for="(customer, index) in customers"
                        class="m-1 p-1 border border-black bg-slate-300 rounded-xl">
                        <CustomerData @orderComplete="orderDone" :customer="customer"
                            :foods-by-customer="foodsByCustomer" :show-current-orders="showCurrentOrders"
                            :index="index" />
                    </li>
                </ul>
            </div>

            <div v-else-if="!isShowOrder">
                <FoodEdit @delete-food="deleteFood" @newFoodSubmit="postNewFood" :foods="foods"
                    :status-message="statusMessage" />
            </div>

        </div>
    </div>
</template>

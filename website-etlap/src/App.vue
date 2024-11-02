<script setup lang="ts">
import { onMounted, ref } from 'vue';

const basket: any = ref([]);

function addToBasket(food: any) {
  basket.value.push(food);
  saveBasket();
}

// https://stackoverflow.com/questions/5767325/how-can-i-remove-a-specific-item-from-an-array-in-javascript
function deleteFromBasket(index: Number) {
  basket.value.splice(index, 1);
  saveBasket();
}

function saveBasket() {
  localStorage.setItem("basket", JSON.stringify(basket.value));
}

onMounted(() => {
  if (localStorage.getItem("basket")) {
    // Check if basket is null
    // https://stackoverflow.com/questions/46915002/argument-of-type-string-null-is-not-assignable-to-parameter-of-type-string
    basket.value = JSON.parse(localStorage.getItem("basket") || "{}");
  }
});


</script>

<template>
  <strong>Current route path:</strong> {{ $route.fullPath }}
  <nav>
    <ul>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/etelek">Ételek</RouterLink>
      <RouterLink to="/rendeles">Rendelés</RouterLink>
    </ul>
  </nav>
  <main>
    <RouterView @kosarhoz="addToBasket" @torol="deleteFromBasket" :basket="basket" />
  </main>

</template>

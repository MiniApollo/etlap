<script setup lang="ts">
import { onMounted, ref } from 'vue';

const basketContent: any = ref([]);

function addToBasket(food: any) {
  basketContent.value.push(food);
  saveBasket();
}

// https://stackoverflow.com/questions/5767325/how-can-i-remove-a-specific-item-from-an-array-in-javascript
function deleteFromBasket(index: Number) {
  basketContent.value.splice(index, 1);
  saveBasket();
}

function saveBasket() {
  localStorage.setItem("basketContent", JSON.stringify(basketContent.value));
}

onMounted(() => {
  if (localStorage.getItem("basketContent")) {
    // Check if basketContent is null
    // https://stackoverflow.com/questions/46915002/argument-of-type-string-null-is-not-assignable-to-parameter-of-type-string
    basketContent.value = JSON.parse(localStorage.getItem("basketContent") || "{}");
  }
});


</script>

<template>
  <strong>Current route path:</strong> {{ $route.fullPath }}
  <nav>
    <ul>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/etelek">Ételek</RouterLink>
      <RouterLink to="/kosar">Kosár</RouterLink>
    </ul>
  </nav>
  <main>
    <RouterView @kosarhoz="addToBasket" @torol="deleteFromBasket" :basketContent="basketContent" />
  </main>

</template>

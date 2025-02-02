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

function emptyBasket() {
  basketContent.value = [];
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
  <div class="bg-gray-200">
    <!-- Need text-base to have smooth font changing if not navbar changes instantly -->
    <nav class="bg-brown fixed top-0 w-full p-3 hover:py-4 transition-all duration-500 text-base">
      <ul>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-black hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/" activeClass="bg-amber-500">Home</RouterLink>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-black hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/etelek" activeClass="bg-amber-500">Ételek</RouterLink>
        <RouterLink
          class="mx-2 my-3 px-1 py-2  text-black hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/kosar" activeClass="bg-amber-500">Kosár</RouterLink>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-black hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/admin" activeClass="bg-amber-500">Admin</RouterLink>
      </ul>
    </nav>
    <main class="pt-16">
      <RouterView @kosarhoz="addToBasket" @torol="deleteFromBasket" @empty-basket="emptyBasket"
        :basketContent="basketContent" />
    </main>
  </div>

</template>

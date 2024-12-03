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
  <div class="bg-slate-300">
    <!-- Need text-base to have smooth font changing if not navbar changes instantly -->
    <nav class="fixed top-0 w-full p-3 bg-orange-200 hover:py-4 transition-all duration-500 text-base">
      <ul>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-slate-800 hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/" activeClass="bg-orange-400 text-black">Home</RouterLink>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-slate-800 hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/etelek" activeClass="bg-orange-400 text-black">Ételek</RouterLink>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-slate-800 hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/kosar" activeClass="bg-orange-400 text-black">Kosár</RouterLink>
      </ul>
    </nav>
    <main class="pt-16">
      <RouterView @kosarhoz="addToBasket" @torol="deleteFromBasket" :basketContent="basketContent" />
    </main>
  </div>

</template>

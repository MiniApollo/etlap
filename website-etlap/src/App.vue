<script setup lang="ts">
import { onMounted, ref } from 'vue';

const basketContent: any = ref([]);
const currentYear = ref(new Date().getFullYear());

function addToBasket(food: any) {
  basketContent.value.push(food);
  saveBasket();
}

// https://stackoverflow.com/questions/5767325/how-can-i-remove-a-specific-item-from-an-array-in-javascript
function deleteFromBasket(EtelID: any) {
  const index: any = basketContent.value.findIndex((element: any) => element.EtelID == EtelID);
  
  if (index == -1) {
    console.error("Item not found in basket");
    return
  }
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
  <div>
    <!-- Need text-base to have smooth font changing if not navbar changes instantly -->
    <nav class="bg-brown fixed top-0 w-full p-3 hover:py-4 transition-all duration-500 text-base">
      <ul>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-black hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/" activeClass="bg-amber-500">Főoldal</RouterLink>
        <RouterLink
          class="mx-2 my-3 px-1 py-2 text-black hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/etelek" activeClass="bg-amber-500">Ételek</RouterLink>
        <RouterLink
          class="mx-2 my-3 px-1 py-2  text-black hover:text-black hover:text-xl transition-all duration-300 rounded-md"
          to="/kosar" activeClass="bg-amber-500">Kosár</RouterLink>
      </ul>
    </nav>
    <main class="pt-16">
      <RouterView @kosarhoz="addToBasket" @torol="deleteFromBasket" @empty-basket="emptyBasket"
        :basketContent="basketContent" />
    </main>
    <footer class="flex justify-center max-sm:flex-col max-sm:items-center gap-2 text-xl px-4 mt-2">
      <p>© {{ currentYear }}. Minden jog fenntartva.</p>
      <p> Vizsga feladat céljából készült.</p>
      <a href="https://github.com/MiniApollo/etlap" class="hover:text-blue-500 hover:underline flex gap-1"
        target="_blank">
        <img class="h-6" src="/assets/github-mark.png" alt="Github logo">Github</a>
    </footer>
  </div>

</template>

<script setup lang="ts">

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
    <RouterView />
  </main>

</template>

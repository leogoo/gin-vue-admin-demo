<template>
  <div>
home
  </div>
</template>

<script setup lang="ts">
import { userStore } from '@/store/user';
const state = userStore();
console.log(state.doubleCount);
state.increment();
console.log(state.doubleCount);
</script>

<style scoped>

</style>
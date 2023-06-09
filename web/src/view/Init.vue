<template>
  <div>
    <el-button type="primary" @click="initDB">init</el-button>
  </div>
</template>

<script setup lang="ts">
import { checkDB } from '@/service/initDb';
const initDB = async () => {
  const res = await checkDB();
  console.log(res);
}
</script>

<style scoped>

</style>
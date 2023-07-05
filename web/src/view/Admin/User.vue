<template>
  <div>
    用户列表

  </div>
</template>

<script setup lang="ts">
import { watchEffect } from 'vue';
import { request } from '@/utils/request';

watchEffect(() => {
  request({
    url: '/userList',
    method: 'GET',
  }).then(res => {
    console.log(res);
  })
})
</script>

<style scoped>

</style>
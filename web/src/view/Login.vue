<template>
  <div class="container">
    <el-form ref="userFormRef" v-model="userFrom">
      <el-form-item prop="username">
        <el-input v-model="userFrom.username" placeholder=""></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input v-model="userFrom.password" placeholder=""></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="clickHandler">login</el-button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { login } from '@/service/common';

const userFormRef = ref();
const userFrom = ref({
  username: "",
  password: ""
});
const clickHandler = async () => {
  const userInfo = { ...userFrom.value };
  const data = await login(userInfo);
  console.log(data);
}
</script>

<style scoped lang="scss">
.container {
  width: 400px;
}
</style>
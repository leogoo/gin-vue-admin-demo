<template>
  <el-container>
    <el-aside width="200px">
      <el-scrollbar>
        <el-menu :default-openeds="['1', '3']">
          <el-sub-menu index="1">
            <template #title>
              <el-icon><message /></el-icon>Navigator One
            </template>
            <el-menu-item-group>
              <template #title>Group 1</template>
              <el-menu-item index="1-1">Option 1</el-menu-item>
              <el-menu-item index="1-2">Option 2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group 2">
              <el-menu-item index="1-3">Option 3</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="1-4">
              <template #title>Option4</template>
              <el-menu-item index="1-4-1">Option 4-1</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
        </el-menu>
      </el-scrollbar>
    </el-aside>
    <el-header>2</el-header>
    <el-main>
      <RouterView></RouterView>
</el-main>
    <el-footer>4</el-footer>
  </el-container>
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
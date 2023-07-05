<template>
  <div class="common-layout">
    <el-container>
      <el-aside class="layout-aside">
        <h5 class="mb-2">gin+vue</h5>
        <el-menu active-text-color="#ffd04b" background-color="#545c64" text-color="#fff" router>
          <el-sub-menu index="/admin">
            <template #title>
              <el-icon>
                <location />
              </el-icon>
              <span>超级管理员</span>
            </template>
            <el-menu-item index="/admin/authority">角色管理</el-menu-item>
            <el-menu-item index="/admin/user">用户管理</el-menu-item>
            <el-menu-item index="/admin/operation">操作历史</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="/media">
            <template #title>
              <el-icon>
                <location />
              </el-icon>
              <span>媒体库</span>
            </template>
            <el-menu-item index="/media/fileList">文件列表</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-aside>
      <el-container>
        <el-main>
          <div>header</div>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import {
  Location,
} from '@element-plus/icons-vue'
</script>

<style scoped lang="scss">
.common-layout {
  height: 100%;
}

.layout-aside {
  width: 200px;
  color: #fff;
  background-color: #545c64;
}
</style>
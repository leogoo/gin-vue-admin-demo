<template>
  <el-upload action="" :http-request="handleRequest" :show-file-list="false">
    Click to upload
  </el-upload>
  <el-button @click="handleDownloadByBlob">二进制流下载</el-button>
  <a href="">文件下载</a>
</template>

<script lang="ts" setup>

import { request } from '@/utils/request';
const handleRequest = (e) => {
  const file = e.file
  const formData = new FormData()
  formData.append('file', file)
  return request({
    url: '/upload',
    method: 'post',
    data: formData
  }).then(res => {
    console.log(res.data)
  })
}
const handleDownloadByBlob = () => {
  request({
    url: '/downloadByBlob?fileName=be1c97d7-4c74-446e-ad39-826ce970e5a1.xlsx',
    method: 'get',
  }).then(res => {
    console.log(res);
    // 用返回二进制数据创建一个Blob实例          
    let blob = new Blob([res.data], {            
      type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", 
    }) // for .xlsx files  

    const url = URL.createObjectURL(blob);

    const a = document.createElement('a');
    a.href = url;
    a.style.display = "none";
    a.download = "自定义文件名.xlsx"

    console.log(url);
    a.click();

    a.remove();

    URL.revokeObjectURL(url); 
  })
}

</script>

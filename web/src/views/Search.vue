<script setup lang="ts">

const bookList = [{
  name: 'zs'
}]

import Header from "@/components/Header.vue";</script>

<template>
  <div>
    <Header title="搜索"/>
    <div class="box">
      <div class="search-box">
        <el-input placeholder="请输入搜索内容"/>
        <el-button type="primary">搜索</el-button>
      </div>
      <div class="search-table">
        <el-table :data="bookList" border>
          <el-table-column prop="name" label="Name"/>
          <el-table-column fixed="right" label="操作" width="100" align="center">
            <template #default>
              <el-button type="text" size="small">添加书架</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.box {
  height: calc(100% - 10px);
  width: calc(100% - 10px);
  margin: 5px;

  .search-box {
    width: 100%;
    display: flex;

    & > :first-child {
      margin-right: 10px;
    }
  }

  .search-table {
    margin-top: 5px;
    height: auto;
    width: 100%;
  }
}
</style>

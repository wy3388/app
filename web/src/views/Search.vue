<script setup lang="ts">

import Header from "@/components/Header.vue";
import {onMounted, ref} from "vue";
import {search, source} from "@/http";
import {Search, Source} from "@/model";
import {ElMessage} from "element-plus";

onMounted(async () => {
  sourceList.value = await source();
})

const sourceList = ref<Array<Source>>([])
const sourceId = ref<number>(1)
const bookList = ref<Array<Search>>([])
const name = ref<string>('')
const isLoading = ref<boolean>(false)


const doSearch = async () => {
  if (name.value === '') {
    ElMessage.warning('请输入搜索内容')
    return
  }
  if (sourceId.value < 1) {
    ElMessage.warning('请选择数据源')
  }
  isLoading.value = true
  bookList.value = await search(name.value, sourceId.value)
  isLoading.value = false
}
</script>

<template>
  <div>
    <Header title="搜索"/>
    <div class="box">
      <div class="search-box">
        <el-select v-model="sourceId" placeholder="请选择数据源">
          <el-option :label="s.name" :value="s.id" v-for="s in sourceList" :key="s.id"></el-option>
        </el-select>
        <el-input placeholder="请输入搜索内容" v-model="name"/>
        <el-button type="primary" @click="doSearch">搜索</el-button>
      </div>
      <div class="search-table">
        <el-table :data="bookList" border empty-text="暂无内容" v-loading="isLoading">
          <el-table-column prop="book_name" label="书名"/>
          <el-table-column prop="author" label="作者"/>
          <el-table-column prop="new_chapter" label="最新章节"/>
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

    & > :nth-child(2) {
      margin-left: 5px;
      margin-right: 5px;
    }
  }

  .search-table {
    margin-top: 5px;
    height: auto;
    width: 100%;
  }

}
</style>

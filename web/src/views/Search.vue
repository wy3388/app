<script setup>

import Body from "@/components/Body.vue";
import {useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import {search, source} from "@/http";
import {ElMessage} from "element-plus";

const router = useRouter()

const source_list = ref()
const source_id = ref()
const book_list = ref()
const name = ref('斗罗大陆')
const loading = ref(false)

onMounted(async () => {
  source_list.value = await source()
  if (source_list.value && source_list.value.length > 0) {
    source_id.value = source_list.value[0].id
  }
})

const doSearch = async () => {
  if (name.value === '') {
    ElMessage.warning('请输入搜搜内容')
    return
  }
  if (!source_id.value) {
    ElMessage.warning('请选择数据源')
    return
  }
  loading.value = true
  book_list.value = await search(name.value, source_id.value)
  loading.value = false
}

const toInfo = (url) => {
  router.push({
    path: '/info',
    query: {
      book_url: url,
      id: source_id.value
    }
  })
}

</script>

<template>
  <Body :show-header="true">
  <div class="box">
    <div class="search-box">
      <el-select v-model="source_id" placeholder="请选择数据源">
        <el-option :label="s.name" :value="s.id" v-for="s in source_list" :key="s.id"></el-option>
      </el-select>
      <el-input placeholder="请输入搜索内容" v-model="name"/>
      <el-button type="primary" @click="doSearch">搜索</el-button>
    </div>
    <div class="search-table">
      <el-table :data="book_list" border empty-text="暂无内容" v-loading="loading">
        <el-table-column prop="book_name" label="书名"/>
        <el-table-column prop="author" label="作者" width="120"/>
        <el-table-column prop="new_chapter" label="最新章节"/>
        <el-table-column fixed="right" label="操作" width="160" align="center">
          <template #default="scope">
            <el-button type="text" size="small" @click.prevent="toInfo(scope.row.book_url)">查看详情</el-button>
            <el-button type="text" size="small">添加书架</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
  </Body>
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

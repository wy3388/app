<script setup>

import Body from "@/components/Body.vue";
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import {info} from "@/http";

const route = useRoute()
const router = useRouter()

const params = {
  id: route.query.id,
  url: route.query.book_url
}

const book = ref()
const loading = ref(true)

onMounted(async () => {
  book.value = await info(params.url, params.id)
  loading.value = false
})

const toRead = (item, index) => {
  localStorage.setItem("chapter", JSON.stringify({
    index: index,
    chapter_list: book.value.chapter_list
  }))
  router.push({
    path: '/read',
    query: {
      url: item.url,
      id: params.id
    }
  })
}

</script>

<template>
  <Body :show-header="true" v-loading="loading">
  <el-card class="card">
    <div class="card-body">
      <img :src="book?.image_url" alt="" class="bs">
      <div class="card-right">
        <div class="card-right-title">
          <h2>{{ book?.book_name }}</h2>
          <el-button type="primary" size="small" class="card-right-title-btn">加入书架</el-button>
        </div>
        <div class="f18">作者：{{ book?.author }}</div>
        <div class="f16">更新时间：{{ book?.update_time }}</div>
        <div class="f14" v-html="book?.info"></div>
      </div>
    </div>
  </el-card>
  <div class="chapter">
    <div class="chapter-box" v-for="(item,index) in book?.chapter_list" :key="index"
         @click="toRead(item, index)">
      <el-card>
        <div class="f14 chapter-div" :title="item.title">{{ item.title }}</div>
      </el-card>
    </div>
  </div>
  </Body>
</template>

<style lang="less" scoped>

@import "../assets/base.css";

.card {
  margin: 5px;


  &-body {
    display: flex;

    img {
      height: 200px;
      width: 150px;
      border-radius: 5px;
      object-fit: cover;
    }
  }

  &-right {
    margin-left: 20px;
    height: 200px;

    &-title {
      display: flex;

      &-btn {
        margin-left: auto;
      }
    }

    & > :nth-child(n+2) {
      margin-top: 5px;
    }
  }
}

.chapter {
  margin: 5px 5px 5px 5px;
  display: flex;
  flex-wrap: wrap;

  &-box {
    width: calc(33.3% - 10px);
    margin: 5px;
    cursor: pointer;
  }

  &-div {
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }
}


</style>
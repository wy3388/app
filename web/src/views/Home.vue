<script setup>
import {useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import {listBookSelf} from "@/http";

const router = useRouter()
const book_self_list = ref([])
const loading = ref(true)

onMounted(async () => {
  book_self_list.value = await listBookSelf()
  loading.value = false
})

const toSearch = () => {
  router.push({path: '/search'})
}

const toRead = (item) => {
  router.push({
    path: '/read',
    query: {
      url: item.chapter_url,
      id: item.source_id,
      type: 'home',
      book_url: item.book_url,
      chapter_index: item.chapter_index,
      book_self_id: item.id
    }
  })
}
</script>

<template>
  <div class="box" v-loading="loading">
    <div class="box-item" v-show="book_self_list.length > 0" @click="toRead(item)"
         v-for="(item, index) in book_self_list"
         :key="index">
      <el-card class="wh100">
        <div class="f16 box-item-book c1 no-select">{{ item.book_name }}</div>
        <div class="f14 box-item-book c3 no-select">{{ item.chapter_name }}</div>
        <span class="f12 box-item-percentage c4 no-select">{{
            (item.chapter_index / item.chapter_total).toFixed(2)
          }}%</span>
      </el-card>
    </div>
    <div class="box-item" @click="toSearch">
      <el-card class="wh100 box-item-empty">
        <img src="../assets/img/add.png" alt="" class="no-select empty-img"/>
      </el-card>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import "../assets/base.css";

.box {
  height: calc(100% - 10px);
  width: calc(100% - 10px);
  position: absolute;
  margin: 5px;
  display: flex;
  flex-wrap: wrap;

  .box-item {
    margin: 10px;
    height: 120px;
    width: 220px;

    &-empty {
      display: flex;
      justify-content: center;
      align-items: center;
    }

    .empty-img {
      height: 100%;
      width: 80px;
      object-fit: cover;
    }

    &-book {
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;

      &:nth-child(2) {
        margin-top: 5px;
      }
    }

    &-percentage {
      margin-top: 5px;
    }
  }
}
</style>

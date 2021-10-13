<script lang="ts" setup>
import {useRouter} from "vue-router";

const bookList = [1]
const router = useRouter()

const toSearch = () => {
  router.push({path: '/search'})
}

const toRead = () => {
  router.push({path: '/read'})
}
</script>

<template>
  <div class="box">
    <div class="box-item" v-show="bookList.length > 0" @click="toRead">
      <el-card class="wh100">
        <div class="f16 box-item-book c1 no-select">你好啊啊啊啊啊啊啊啊</div>
        <div class="f14 box-item-book c3 no-select">第二十章 （爱上发大水发射点发射点）</div>
        <span class="f12 box-item-percentage c4 no-select">%4</span>
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

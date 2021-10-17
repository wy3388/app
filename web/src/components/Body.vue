<script setup>

import Header from "@/components/Header.vue";

defineProps({
  showHeader: {
    type: Boolean,
    required: false,
    default: false
  },
  backFun: {
    type: () => {
    },
    required: false,
    default: null
  },
  title: {
    type: String,
    required: false,
    default: ''
  }
})
</script>

<template>
  <div class="body">
    <Header v-show="showHeader" :back-fun="backFun" :title="title"/>
    <div :class="showHeader?'content':'content100'">
      <el-scrollbar height="100%">
        <slot></slot>
      </el-scrollbar>
    </div>
  </div>
</template>

<style scoped lang="less">
.body {
  height: 100%;
  width: 100%;
  position: absolute;
}

.content {
  height: calc(100% - 60px);
  width: 100%;
}

.content100 {
  height: 100%;
  width: 100%;
}
</style>
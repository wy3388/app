<script setup>

import Body from "@/components/Body.vue";
import {useRoute} from "vue-router";
import {onMounted, ref} from "vue";
import {content} from "@/http";
import {ElMessage} from "element-plus";

const route = useRoute()

const params = {
  id: route.query.id,
  url: route.query.url
}

const body = ref({})
const loading = ref(false)

let chapter = {
  index: -1,
  chapter_list: []
}

onMounted(async () => {
  chapter = JSON.parse(localStorage.getItem("chapter"))
  await load(params.url)
})

const load = async (url) => {
  loading.value = true
  body.value = await content(url, params.id)
  loading.value = false
}

const pre = async () => {
  if (chapter.index > 0 && chapter.chapter_list.length > 0) {
    chapter.index--
    await load(chapter.chapter_list[chapter.index].url)
    return
  }
  ElMessage.warning('没有上一章了')
}

const next = async () => {
  if (chapter.index > -1 && chapter.chapter_list.length > 0 && chapter.index < chapter.chapter_list.length - 1) {
    chapter.index++
    await load(chapter.chapter_list[chapter.index].url)
    return
  }
  ElMessage.warning('没有下一章了')
}

</script>

<template>
  <Body :show-header="true">
  <div v-html="body?.body" class="read bs" v-loading="loading"></div>
  <div class="btn bs">
    <el-button type="primary" @click="pre">上一章</el-button>
    <el-button type="primary" @click="next">下一章</el-button>
  </div>
  </Body>
</template>

<style lang="less" scoped>
.read {
  padding: 20px;
  height: calc(100% - 90px);
  margin-bottom: 90px;
  position: absolute;
  width: calc(100% - 40px);
}

.btn {
  display: flex;
  padding: 20px;
  position: absolute;
  bottom: 0;
  left: 0;
  width: calc(100% - 40px);
  background: white;

  & > :last-child {
    margin-left: auto;
  }
}
</style>
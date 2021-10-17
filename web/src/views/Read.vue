<script setup>

import Body from "@/components/Body.vue";
import {useRoute} from "vue-router";
import {onMounted, ref} from "vue";
import {content, info, updateBookSelf} from "@/http";
import {ElMessage} from "element-plus";

const route = useRoute()

const params = {
  id: route.query.id,
  url: route.query.url,
  type: route.query.type,
  book_url: route.query.book_url,
  chapter_index: parseInt(route.query.chapter_index),
  book_self_id: parseInt(route.query.book_self_id)
}

const body = ref({})
const loading = ref(false)

let chapter = {
  index: -1,
  chapter_list: []
}

onMounted(async () => {
  document.onkeyup = (e) => {
    if (e.key === 'ArrowLeft') {
      pre()
    }
    if (e.key === 'ArrowRight') {
      next()
    }
  }
  if (params.type === 'info') {
    chapter = {
      index: params.chapter_index,
      chapter_list: JSON.parse(localStorage.getItem("chapter_list"))
    }
    await load(params.url)
  } else {
    const book = await info(params.book_url, params.id)
    chapter = {
      index: params.chapter_index,
      chapter_list: book.chapter_list
    }
    await load(params.url)
  }
})

const load = async (url) => {
  loading.value = true
  body.value = await content(url, params.id)
  loading.value = false
}

const save = async () => {
  if (params.type === 'home') {
    let param = {
      id: params.book_self_id,
      chapter_url: chapter.chapter_list[chapter.index].url,
      chapter_name: chapter.chapter_list[chapter.index].title,
      chapter_total: chapter.chapter_list.length,
      chapter_index: chapter.index
    }
    const [resp, err] = await updateBookSelf(param).then(resp => [resp, null]).catch(err => [null, err])
    if (err != null) {
      ElMessage.error(err)
    }
  }
}

const pre = async () => {
  if (chapter.index > 0 && chapter.chapter_list.length > 0) {
    chapter.index--
    await load(chapter.chapter_list[chapter.index].url)
    await save()
    return
  }
  ElMessage.warning('没有上一章了')
}

const next = async () => {
  if (chapter.index > -1 && chapter.chapter_list.length > 0 && chapter.index < chapter.chapter_list.length - 1) {
    chapter.index++
    await load(chapter.chapter_list[chapter.index].url)
    await save()
    return
  }
  ElMessage.warning('没有下一章了')
}

</script>

<template>
  <Body :show-header="true" :title="body.title">
  <div v-html="body.body" class="read" v-loading="loading"></div>
  </Body>
  <div class="btn bs">
    <el-button type="primary" @click="pre">上一章</el-button>
    <el-button type="primary" @click="next">下一章</el-button>
  </div>
</template>

<style lang="less" scoped>

.box {
  height: 100%;
  width: 100%;
}


.read {
  padding: 20px;
  height: calc(100% - 90px);
  margin-bottom: 90px;
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
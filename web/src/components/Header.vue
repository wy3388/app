<script setup>
import {useRouter} from "vue-router";

const props = defineProps({
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

const router = useRouter()
const goBack = () => {
  if (props.backFun != null && typeof props.backFun === 'function') {
    props.backFun()
  } else {
    router.go(-1)
  }
}

</script>

<template>
  <div class="header bs">
    <div class="header-back" @click="goBack">
      <img src="../assets/img/back.png" alt="" class="no-select">
      <span class="f18 no-select c1" v-show="title !== ''">{{ title }}</span>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import "../assets/base.css";

.header {
  height: 60px;
  width: 100%;
  display: flex;
  align-items: center;

  &-back {
    height: 40px;
    margin-left: 1%;
    display: flex;
    align-items: center;

    img {
      height: 100%;
      object-fit: cover;
    }

    span {
      letter-spacing: 3px;
    }
  }
}
</style>
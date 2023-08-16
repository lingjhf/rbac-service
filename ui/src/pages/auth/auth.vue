<template>
  <el-container class="h-screen">
    <el-header class="bg-[var(--header-bg-color)] flex items-center">
      <el-button class="ml-auto" @click="switchSignupOrLogin">
        {{ signupOrLogin }}
      </el-button>
    </el-header>
    <el-main class="bg-[var(--main-bg-color)]">
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { ElContainer, ElHeader, ElMain, ElButton } from 'element-plus'
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { RouteName } from '@/router'
const router = useRouter()
const route = useRoute()
const signupOrLogin = computed(() =>
  route.name === RouteName.login ? '注册' : '登录'
)
function switchSignupOrLogin() {
  router.push({
    name: route.name === RouteName.login ? RouteName.signup : RouteName.login,
  })
}
</script>

<style scoped></style>

<template>
  <div class="flex h-full w-full items-center justify-center">
    <el-form label-position="top" :model="form">
      <el-form-item label="用户名">
        <r-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="密码">
        <r-input type="password" v-model="form.password" />
      </el-form-item>
      <r-button
        :loading="loading"
        type="primary"
        class="w-full"
        @click="onLogin"
      >
        登录
      </r-button>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElForm, ElFormItem } from 'element-plus'
import { RButton, RInput } from '@/components'
import * as api from '@/api'
import cookie from 'js-cookie'
import { useRouter } from 'vue-router'
import { RouteName } from '@/router'

const router = useRouter()
const loading = ref(false)
const form = ref<api.LoginWithUsername>({
  username: '',
  password: '',
})

async function onLogin() {
  loading.value = true
  try {
    const data = await api.Login('username', form.value)
    cookie.set('token', data.token)
    router.push({ name: RouteName.home })
  } catch (e) {
    console.log(e)
  }
  loading.value = false
}
</script>

<template>
  <div class="flex h-full w-full items-center justify-center">
    <el-form label-position="top" :model="form">
      <el-form-item label="用户名">
        <r-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="邮箱">
        <r-input v-model="form.email" />
      </el-form-item>
      <el-form-item label="密码">
        <r-input type="password" v-model="form.password" />
      </el-form-item>
      <r-button type="primary" class="w-full" @click="onSignup">
        注册
      </r-button>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ElForm, ElFormItem } from 'element-plus'
import { RInput, RButton } from '@/components'
import { ref } from 'vue'
import * as api from '@/api'

const form = ref<api.SignupWithEmailData>({
  username: '',
  email: '',
  password: '',
})

async function onSignup() {
  try {
    await api.Signup('email', form.value)
  } catch (e) {}
}
</script>

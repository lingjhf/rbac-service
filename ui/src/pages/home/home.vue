<template>
  <el-container direction="vertical" class="h-full">
    <r-header />
    <el-main class="bg-[var(--main-bg-color)] flex flex-col p-0!">
      <div class="flex h-12 p-2 box-border">
        <r-input
          class="w-[180px]!"
          placeholder="搜索租户"
          :prefix-icon="search"
          v-model="searchValue"
        />
        <r-button class="ml-auto" type="primary" @click="addTenant">
          新建租户
        </r-button>
        <create-tenant-dialog v-model="dialogVisible" />
      </div>
      <div class="h-[calc(100%-48px)]">
        <el-scrollbar>
          <tenant-grid />
        </el-scrollbar>
      </div>
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { defineComponent, h, ref } from 'vue'
import { ElContainer, ElMain, ElScrollbar } from 'element-plus'
import { RButton, RInput } from '@/components'
import RHeader from '@/pages/header'
import { CreateTenantDialog } from '@/pages/components'
import TenantGrid from './tenant-grid.vue'

const search = defineComponent({
  setup: () => () => h('div', { class: 'i-local:search' }),
})

const searchValue = ref('')
const dialogVisible = ref(false)

function addTenant() {
  dialogVisible.value = true
}
</script>

<template>
  <r-loading :loading="initData">
    <el-container direction="vertical" class="bg-[var(--main-bg-color)] h-full">
      <r-header>
        <!-- <r-button class="h-8! w-8!" text @click="showSidebar">
          <span v-if="visibleSidebar" class="h-6 w-6 i-local:expand" />
          <span v-else class="h-6 w-6 i-local:fold" />
        </r-button> -->
        <r-button class="w-8 h-8!" text @click="toHome">
          <div class="h-6 w-6 i-local:home-outline" />
        </r-button>
        <r-button class="h-8! ml-2! px-2!" text>
          <span>{{ tenantStore.currentTenant?.name }}</span>
        </r-button>
      </r-header>
      <el-container>
        <tenant-sidebar v-model="visibleSidebar" />
        <el-main class="flex-col! flex! p-0!">
          <tenant-main />
          <tenant-footer v-model:expand="visibleSidebar" />
        </el-main>
      </el-container>
    </el-container>
  </r-loading>
</template>

<script setup lang="ts">
import { ElContainer, ElMain } from 'element-plus'
import { RLoading, RButton } from '@/components'
import RHeader from '@/pages/header'
import TenantSidebar from './tenant-sidebar.vue'
import TenantMain from './tenant-main.vue'
import TenantFooter from './tenant-footer.vue'
import { useTenantStore } from '@/store'
import { useRoute, useRouter } from 'vue-router'
import { RouteName } from '@/router'
import { ref } from 'vue'

const route = useRoute()
const router = useRouter()
const tenantStore = useTenantStore()

const visibleSidebar = ref(true)

async function initData() {
  const tenantId = route.params.id as string
  tenantStore.setCurrentTenant({ id: tenantId, name: '', parentId: null })
  await Promise.all([
    tenantStore.getTenant(tenantId),
    tenantStore.getTenantChildren(tenantId),
  ])
}

function toHome() {
  router.push({ name: RouteName.home })
}
</script>

<template>
  <div class="flex flex-col h-full">
    <div class="flex h-8 text-sm p-2 items-center">
      <div
        v-for="(tab, index) in tabs"
        :key="index"
        class="border-b-solid border-transparent cursor-pointer h-full border-b-2 px-5 hover:text-[var(--primary-color)]"
        :style="tabStyles(tab.name).value"
        @click="tabChange(tab.name)"
      >
        <span>{{ tab.label }}</span>
      </div>
    </div>
    <div class="h-full">
      <user v-show="currentTab === 'user'" />
      <role v-show="currentTab === 'role'" />
      <permission v-show="currentTab === 'permission'" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import User from '@/pages/user'
import Role from '@/pages/role'
import Permission from '@/pages/permission'

const tabs = ref([
  { label: '用户', name: 'user' },
  { label: '角色', name: 'role' },
  { label: '权限', name: 'permission' },
])
const currentTab = ref('user')

const tabStyles = (name: string) =>
  computed(() => {
    if (name === currentTab.value) {
      return 'color: var(--primary-color);border-color: var(--primary-color)'
    }
    return ''
  })

function tabChange(name: string) {
  currentTab.value = name
}
</script>

<template>
  <div :style="styles" class="h-screen">
    <router-view />
  </div>
</template>

<script setup lang="ts">
import { useThemeStore } from '@/store'
import { computed } from 'vue'

const themeStore = useThemeStore()
const styles = computed(() => ({
  '--header-bg-color': themeStore.theme.headerBgColor,
  '--main-bg-color': themeStore.theme.mainBgColor,
  '--sidebar-bg-color': themeStore.theme.sidebarBgColor,
  '--text-color': themeStore.theme.textColor,
  '--text-btn-hover-bg-color': themeStore.theme.textBtnHoverBgColor,
  '--btn-text-primary-color': themeStore.theme.btnTextPrimaryColor,
  '--primary-color': themeStore.theme.primaryColor,
  '--btn-hover-bg-color': themeStore.theme.btnHoverBgColor,
  '--input-bg-color': themeStore.theme.sideBarInputBgColor,
  '--tree-node-hover-bg-color': themeStore.theme.treeNodeHoverBgColor,
  '--split-border-color':themeStore.theme.splitBorderColor,
  color: 'var(--text-color)',
}))
</script>

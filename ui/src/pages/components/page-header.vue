<template>
  <el-header class="bg-[var(--header-bg-color)] flex items-center h-10!">
    <slot />
    <el-popover
      width="112"
      trigger="click"
      :popper-style="propoverStyles"
      :show-arrow="false"
    >
      <template #reference>
        <el-avatar
          class="ml-auto"
          style="--el-avatar-size: 32px"
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        />
      </template>
    </el-popover>
  </el-header>
</template>

<script setup lang="ts">
import { ElHeader, ElPopover, ElAvatar } from 'element-plus'
import { useThemeStore } from '@/store'
import { computed } from 'vue'

const themeStore = useThemeStore()

const propoverStyles = computed(() => ({
  '--el-popover-bg-color': themeStore.theme.mainBgColor,
  '--el-popover-border-color': themeStore.theme.mainBgColor,
  '--el-popover-padding': '4px',
  'min-width': '112px',
}))
</script>

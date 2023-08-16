<template>
  <el-popover
    v-bind="props"
    v-on="event"
    :popper-style="propoverStyles"
    :show-arrow="false"
  >
    <slot />
    <template #reference>
      <slot name="reference" />
    </template>
  </el-popover>
</template>

<script setup lang="ts">
import { ElPopover, popoverProps, popoverEmits } from 'element-plus'
import { useThemeStore } from '@/store'
import { computed } from 'vue'

const props = defineProps(popoverProps)
const emit = defineEmits(popoverEmits)

const event = {
  'update:visible': (value: boolean) => emit('update:visible', value),
  'before-enter': () => emit('before-enter'),
  'before-leave': () => emit('before-leave'),
  'after-enter': () => emit('after-enter'),
  'after-leave': () => emit('after-leave'),
}

const themeStore = useThemeStore()

const propoverStyles = computed(() => ({
  '--el-popover-bg-color': themeStore.theme.mainBgColor,
  '--el-popover-border-color': themeStore.theme.mainBgColor,
  '--el-popover-padding': '4px 0px 4px 0px',
  'min-width': '112px',
}))
</script>

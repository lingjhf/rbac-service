<template>
  <el-button v-bind="props" v-on="events" :style="styles">
    <template #default>
      <slot />
    </template>
  </el-button>
</template>

<script setup lang="ts">
import { ElButton, buttonProps, buttonEmits } from 'element-plus'
import { computed } from 'vue'

const props = defineProps(buttonProps)
const emit = defineEmits(buttonEmits)

const styles = computed(() => {
  if (props.text) {
    return `
        --el-fill-color-light: var(--text-btn-hover-bg-color);
        --el-fill-color: var(--text-btn-hover-bg-color);
    `
  }

  switch (props.type) {
    case '':
      return `
        --el-button-bg-color: var(--sidebar-bg-color);
        --el-button-border-color: var(--btn-hover-bg-color);

        --el-button-hover-bg-color: var(--sidebar-bg-color);
        --el-button-hover-border-color: var(--primary-color);

        --el-button-active-bg-color: var(--sidebar-bg-color);
        --el-button-active-border-color: var(--btn-hover-bg-color);

        --el-button-text-color: var(--text-color);
        --el-button-hover-text-color: var(--primary-color);
        --el-button-active-text-color: var(--primary-color);
      `

    case 'primary':
      return `
      --el-button-bg-color: var(--primary-color);
      --el-button-border-color: var(--primary-color);

      --el-button-hover-bg-color: var(--btn-hover-bg-color);
      --el-button-hover-border-color: var(--btn-hover-bg-color);

      --el-button-active-bg-color: var(--btn-hover-bg-color);
      --el-button-active-border-color: var(--btn-hover-bg-color);
      
      --el-button-text-color: var(--btn-text-primary-color);
      --el-button-hover-text-color: var(--btn-text-primary-color);
      --el-button-active-text-color: var(--btn-text-primary-color);
      `
  }
  return ''
})

const events = {
  click: (evt: MouseEvent) => emit('click', evt),
}
</script>

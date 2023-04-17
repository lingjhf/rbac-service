<template>
  <el-tabs v-bind="props" v-on="events" class="tab-content">
    <slot />
  </el-tabs>
</template>

<script setup lang="ts">
import { ref, provide, watch } from 'vue'
import {
  ElTabs,
  tabsProps,
  tabsEmits,
  TabPaneName,
  TabsPaneContext,
} from 'element-plus'
const props = defineProps(tabsProps)
const emit = defineEmits(tabsEmits)
const currentTab = ref<TabPaneName | undefined>('')
provide('currentTab', currentTab)
watch(
  () => props.modelValue,
  () => {
    currentTab.value = props.modelValue
  },
  { immediate: true }
)
const events = {
  'update:modelValue': (name: TabPaneName) => emit('update:modelValue', name),
  tabClick: (pane: TabsPaneContext, ev: Event) => emit('tabClick', pane, ev),
  tabChange: (name: TabPaneName) => emit('tabChange', name),
  edit: (paneName: TabPaneName | undefined, action: 'remove' | 'add') =>
    emit('edit', paneName, action),
  tabRemove: (name: TabPaneName) => emit('tabRemove', name),
  tabAdd: () => emit('tabAdd'),
}
</script>

<style scoped>
.tab-content {
  --el-color-primary: var(--primary-color);
  --el-bg-color-overlay: var(--main-bg-color);
  --el-fill-color-light: var(--sidebar-bg-color);
  --el-border-color-light: #272835;
  --el-border-color: #272835;
}
.tab-content:deep() .el-tabs__content {
  --at-apply: h-full p-0 bg-[var(--main-bg-color)];
}

.tab-content:deep() .el-tabs__item {
  padding: 0 !important ;
}
</style>

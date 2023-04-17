<template>
  <el-tab-pane v-bind="props">
    <slot />
    <template #label>
      <slot name="label">
        <div
          class="h-full text-xs w-full"
          @mouseenter="labelEnter"
          @mouseleave="labelLeave"
        >
          <div
            v-if="props.name === currentTab || closeVisible"
            class="flex h-full w-full pl-3 items-center"
          >
            {{ props.label }}
            <span class="mx-2 i-local:close" />
          </div>
          <div v-else class="flex h-full pr-7 pl-3 w-full-2 items-center">
            {{ props.label }}
          </div>
        </div>
      </slot>
    </template>
  </el-tab-pane>
</template>

<script setup lang="ts">
import { Ref, inject, ref } from 'vue'
import { ElTabPane, tabPaneProps, TabPaneName } from 'element-plus'

const props = defineProps(tabPaneProps)
const currentTab = inject<Ref<TabPaneName>>('currentTab')

const closeVisible = ref(false)
function labelEnter() {
  closeVisible.value = true
}
function labelLeave() {
  closeVisible.value = false
}
</script>

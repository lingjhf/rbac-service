<template>
  <el-tab-pane v-bind="props">
    <slot />
    <template #label>
      <slot name="label">
        <div
          class="h-full text-xs w-full"
          @mouseenter="labelEnter"
          @mouseleave="labelLeave"
          @contextmenu.prevent
        >
          <div class="flex h-full pr-7 pl-3 w-full-2 items-center relative">
            <slot name="labelText">
              {{ props.label }}
            </slot>
            <div
              v-if="props.name === currentTab || closeVisible"
              class="rounded flex h-5 mx-1 right-0 w-5 items-center justify-center tab-close absolute"
              @click.stop="close"
            >
              <div class="i-local:close" />
            </div>
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
const emitClose = inject<(name?: TabPaneName) => void>('onClose')
const closeVisible = ref(false)

function labelEnter() {
  closeVisible.value = true
}

function labelLeave() {
  closeVisible.value = false
}

function close() {
  emitClose?.(props.name)
}
</script>

<style scoped>
.tab-close:hover {
  background-color: var(--text-btn-hover-bg-color);
}
</style>

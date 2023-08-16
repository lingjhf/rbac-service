<template>
  <div
    v-if="mediaQuery.md && modelValue"
    class="bg-[var(--sidebar-bg-color)] z-99"
  >
    <r-resizable
      v-model:width="width"
      :min-width="200"
      :max-width="600"
      right
      class="h-full"
    >
      <tenant-tree />
    </r-resizable>
  </div>
  <div v-else class="tenant-drawer">
    <r-drawer
      direction="ltr"
      :show-close="false"
      class="w-[250px]!"
      :model-value="drawerVisible"
      @update:model-value="emitUpdateModelValue"
    >
      <tenant-tree />
    </r-drawer>
  </div>
</template>

<script setup lang="ts">
import { watch, ref } from 'vue'
import { RDrawer, RResizable } from '@/components'
import TenantTree from './tenant-tree.vue'
import { useMediaQuery } from '@/store'

const props = defineProps<{ modelValue?: boolean }>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const mediaQuery = useMediaQuery()

const width = ref(200)
const drawerVisible = ref(false)

watch(
  () => props.modelValue,
  () => {
    drawerVisible.value = props.modelValue ?? false
  },
  {
    immediate: true,
  }
)

watch(
  () => mediaQuery.md,
  () => {
    if (!mediaQuery.md) {
      drawerVisible.value = false
      emitUpdateModelValue(drawerVisible.value)
    } else {
      emitUpdateModelValue(true)
    }
  },
  {
    immediate: true,
  }
)

function emitUpdateModelValue(value: boolean) {
  emit('update:modelValue', value)
}
</script>

<style scoped>
.tenant-drawer:deep(.el-drawer) {
  --el-drawer-padding-primary: 0;
}
.tenant-drawer:deep(.el-drawer__header) {
  margin-bottom: 0px;
}
.tenant-drawer:deep(.el-drawer__footer) {
  padding-top: 0px;
}
</style>

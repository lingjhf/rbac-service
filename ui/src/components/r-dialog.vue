<template>
  <el-dialog
    v-bind="props"
    v-on="events"
    style="
      --el-dialog-bg-color: var(--main-bg-color);
      --el-text-color-primary: var(--text-color);
      --el-color-primary: var(--primary-color);
    "
  >
    <slot />
    <template #header>
      <slot name="header" />
    </template>
    <template #footer>
      <slot name="footer" />
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ElDialog, dialogProps, dialogEmits } from 'element-plus'

const props = defineProps(dialogProps)
const emit = defineEmits(dialogEmits)
const events = {
  open: () => emit('open'),
  opened: () => emit('opened'),
  close: () => emit('close'),
  closed: () => emit('close'),
  'update:modelValue': (value: boolean) => emit('update:modelValue', value),
  openAutoFocus: () => emit('openAutoFocus'),
  closeAutoFocus: () => emit('closeAutoFocus'),
}
</script>

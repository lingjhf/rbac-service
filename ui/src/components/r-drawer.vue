<template>
  <el-drawer v-bind="props" v-on="events">
    <slot />
    <template #header>
      <slot name="header" />
    </template>
    <template #footer>
      <slot name="footer" />
    </template>
  </el-drawer>
</template>

<script setup lang="ts">
import { ElDrawer, drawerProps, drawerEmits } from 'element-plus'

const props = defineProps(drawerProps)
const emit = defineEmits(drawerEmits)

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

<template>
  <el-input
    v-bind="props"
    v-on="events"
    style="
      --el-input-text-color: var(--text-color);
      --el-input-focus-border-color: var(--primary-color);
      --el-input-hover-border-color: var(--primary-color);
      --el-input-bg-color: var(--sidebar-input-bg-color);
      --el-input-border-color: none;
    "
  />
</template>

<script setup lang="ts">
import { ElInput, inputProps, inputEmits } from 'element-plus'

const props = defineProps(inputProps)
const emit = defineEmits(inputEmits)
const events = {
  'update:modelValue': (value: string) => emit('update:modelValue', value),
  input: (value: string) => emit('input', value),
  change: (value: string) => emit('change', value),
  focus: (evt: FocusEvent) => emit('focus', evt),
  blur: (evt: FocusEvent) => emit('blur', evt),
  clear: () => emit('clear'),
  mouseleave: (evt: MouseEvent) => emit('mouseleave', evt),
  mouseenter: (evt: MouseEvent) => emit('mouseenter', evt),
  keydown: (evt: KeyboardEvent | Event) => emit('keydown', evt),
  compositionstart: (evt: CompositionEvent) => emit('compositionstart', evt),
  compositionupdate: (evt: CompositionEvent) => emit('compositionupdate', evt),
  compositionend: (evt: CompositionEvent) => emit('compositionend', evt),
}
</script>

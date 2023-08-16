<template>
  <el-pagination v-bind="props" v-on="events" />
</template>

<script setup lang="ts">
import { ElPagination, paginationProps, paginationEmits } from 'element-plus'

const props = defineProps(paginationProps)
const emit = defineEmits(paginationEmits)

const events = {
  'update:current-page': (val: number) => emit('update:current-page', val),
  'update:page-size': (val: number) => emit('update:page-size', val),
}
</script>

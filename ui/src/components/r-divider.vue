<template>
  <el-divider v-bind="props" style="--el-border-color: #606266">
    <slot />
  </el-divider>
</template>

<script setup lang="ts">
import { ElDivider, dividerProps } from 'element-plus'

const props = defineProps(dividerProps)
</script>

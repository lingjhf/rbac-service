<template>
  <el-tree-v2
    v-bind="$attrs"
    class="bg-[var(--sidebar-bg-color)]!"
    style="
      --el-tree-node-hover-bg-color: var(--tree-node-hover-bg-color);
      --el-color-primary-light-9: var(--tree-node-hover-bg-color);
    "
  >
    <template #default="slotProps">
      <slot v-bind="slotProps" />
    </template>
  </el-tree-v2>
</template>

<script setup lang="ts">
import { ElTreeV2 } from 'element-plus'
</script>

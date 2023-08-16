<template>
  <div>
    <r-splitter>
      <!-- <div class="h-[2000px] w-[100px]" /> -->
      <r-splitter-item :min-height="80" />
      <r-splitter-item :min-height="80" />
      <r-splitter-item :min-height="100" />
    </r-splitter>
  </div>
</template>

<script setup lang="ts">
import { RSplitter, RSplitterItem } from '@/components'
import {} from 'element-plus'
</script>

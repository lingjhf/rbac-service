<template>
  <el-table
    :data="props.data"
    style="
      --el-table-bg-color: var(--main-bg-color);
      --el-table-tr-bg-color: var(--main-bg-color);
      --el-table-header-bg-color: var(--main-bg-color);
      --el-table-border-color: var(--split-border-color);
    "
  >
    <el-table-column
      v-for="column in props.columns"
      :key="column.prop"
      :label="column.label"
      :prop="column.prop"
    >
      <template v-if="slots.column" #default="scope">
        <slot name="column" v-bind="scope" />
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import {
  ElTable,
  ElTableColumn,
  TableProps,
  TableColumnInstance,
} from 'element-plus'
import { useSlots } from 'vue'

interface TableWithColumnProps<T> extends TableProps<T> {
  columns: Partial<TableColumnInstance>[]
}

const props = defineProps<TableWithColumnProps<any>>()
const slots = useSlots()
console.log(slots.column)
</script>

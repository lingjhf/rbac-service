<template>
  <div class=" h-full w-full" ref="containerRef">
    <r-splitter>
      <r-splitter-item />
      <r-splitter-item />
    </r-splitter>
    <!-- <el-scrollbar>
      <div class="flex flex-col h-[50%] w-full">
        <r-table
          class="h-full! w-full!"
          border
          :data="data"
          :columns="columns"
        />
        <div class="flex">
          <r-pagination
            class="ml-auto mt-2"
            background
            :layout="paginationLayout"
            :total="50"
          />
        </div>
      </div>
    </el-scrollbar> -->
    <!-- <r-resizable v-model:height="height" :min-height="130" bottom>
    </r-resizable> -->
    <!-- <div class="border-t border-t-solid border-[var(--split-border-color)]">
      <div class="flex h-9 px-4 items-center">
        <div class="cursor-pointer flex items-center">
          <div class="h-4 mr-2 w-4 i-local:arrow-up" />
          分配角色
        </div>
      </div>
    </div> -->
    <!-- <role-transfer class="h-[50%] mt-2" /> -->
  </div>
</template>

<script setup lang="ts">
import { ref, shallowRef, computed, watch } from 'vue'
import type { TableColumnInstance } from 'element-plus'
import { ElScrollbar } from 'element-plus'
import { RTable, RPagination, RResizable,RSplitter,RSplitterItem } from '@/components'
import RoleTransfer from './role-transfer.vue'
import { useMediaQuery } from '@/store'

const mediaQuery = useMediaQuery()

const containerRef = ref<HTMLElement>()
const userHeight = ref(200)
const roleTansferHeiht = ref(200)

const columns: Partial<TableColumnInstance>[] = [
  { label: 'id', prop: 'id', width: '100' },
  { label: '用户名', prop: 'username', width: '100' },
  { label: '修改时间', prop: 'updateAt', width: '100' },
  { label: '创建时间', prop: 'createAt', width: '100' },
]

const data = [{ username: 'fuck' }]

const paginationLayout = computed(() => {
  return mediaQuery.md ? 'prev, pager, next' : 'prev, jumper, next'
})
</script>

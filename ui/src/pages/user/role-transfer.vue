<template>
  <div class="flex flex-col w-full md:flex-row">
    <div class="flex flex-col md:w-[calc((100%-48px)/2)]">
      <r-table :columns="columns" :data="data" class="h-full! w-full!" />
      <div class="flex">
        <r-pagination
          class="mt-2 lt-md:ml-auto"
          small
          background
          :layout="paginationLayout"
          :total="50"
        />
      </div>
    </div>
    <div class="flex my-2 justify-center md:flex-col md:mx-2">
      <div>
        <r-button type="primary" class="h-8 w-8">
          <span class="i-local:arrow-down md:i-local:arrow-left" />
        </r-button>
      </div>
      <div class="lt-md:ml-2 md:mt-2">
        <r-button type="primary" class="h-8 w-8">
          <span class="i-local:arrow-up md:i-local:arrow-right" />
        </r-button>
      </div>
    </div>
    <div class="flex flex-col md:w-[calc((100%-48px)/2)]">
      <r-table :columns="columns" :data="data" class="h-full! w-full!" />
      <div class="flex">
        <r-pagination
          class="ml-auto mt-2"
          small
          background
          :layout="paginationLayout"
          :total="50"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { RTable, RPagination, RButton } from '@/components'
import type { Column } from 'element-plus'
import { useMediaQuery } from '@/store'

const mediaQuery = useMediaQuery()
const columns: Column[] = [
  {
    label: '角色名',
    dataKey: 'name',
    width: 150,
  },
]
const data:any[] = []

const paginationLayout = computed(() => {
  return mediaQuery.md ? 'prev, pager, next' : 'prev, jumper, next'
})
</script>

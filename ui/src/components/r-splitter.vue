<template>
  <div class="bg-blue-300" :style="splitterStyles">
    <el-scrollbar>
      <slot />
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { provide, ref, Ref, computed, watch } from 'vue'
import { ElScrollbar } from 'element-plus'

export interface SplitterProvider {
  addItem: (item: SplitterItem) => number
  setHeight: (index: number, value: number) => void
  setMinHeight: (index: number, value: number) => void
}

export interface SplitterItem {
  height: Ref<number>
  minHeight: number
}

const props = defineProps<{ height?: number }>()

const items: SplitterItem[] = []
const height = ref(500)
watch(
  () => props.height,
  () => {
    if (props.height && props.height > 0) {
      height.value = props.height
    }
  },
  {
    immediate: true,
  }
)
const splitterStyles = computed(() => {
  return `height:${height.value}px`
})

provide('splitterProvider', { addItem, setHeight, setMinHeight })

function addItem(item: SplitterItem) {
  const index = items.push(item) - 1
  const itemHeight = height.value / items.length
  for (const item of items) {
    if (itemHeight < item.minHeight) {
      item.height.value = item.minHeight
      continue
    }
    item.height.value = itemHeight
  }
  return index
}

function setHeight(index: number, value: number) {
  const prevItem = items[index - 1]
  const currentItem = items[index]
  const currentHeight =
    value < currentItem.minHeight ? currentItem.minHeight : value
  const prevHeight =
    prevItem.height.value + currentItem.height.value - currentHeight
  if (prevHeight < prevItem.minHeight) {
    currentItem.height.value =
      prevItem.height.value + currentItem.height.value - prevItem.minHeight
    prevItem.height.value = prevItem.minHeight
    return
  }
  prevItem.height.value = prevHeight
  currentItem.height.value = currentHeight
}

function setMinHeight(index: number, value: number) {
  items[index].minHeight = value
}
</script>

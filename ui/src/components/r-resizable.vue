<template>
  <div class="relative" :style="styles">
    <slot />
    <div
      v-show="props.right && props.width !== undefined"
      class="cursor-e-resize top-0 -right-[2px] bottom-0 w-1 select-none absolute"
      :style="resizableStyles"
      @mouseenter="enterResizable"
      @mouseleave="leaveResizable"
      @mousedown="startMoveX"
    >
      <slot name="right" />
    </div>
    <div
      v-show="props.bottom && props.height !== undefined"
      class="cursor-s-resize h-1 right-0 -bottom-[2px] left-0 absolute select-none"
      :style="resizableStyles"
      @mouseenter="enterResizable"
      @mouseleave="leaveResizable"
      @mousedown="startMoveY"
    >
      <slot name="bottom" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { onMove } from '@/utils'

const props = defineProps<{
  width?: number
  minWidth?: number
  maxWidth?: number
  height?: number
  minHeight?: number
  maxHeight?: number
  right?: boolean
  bottom?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:width', value: number): void
  (e: 'update:height', value: number): void
}>()

const width = ref(0)
const height = ref(0)
const isResizing = ref(false)
const isEnterResizable = ref(false)
let isLeaveResizable = true

watch(
  () => props.width,
  () => {
    if (props.width !== undefined) {
      width.value = props.width
    }
  },
  { immediate: true }
)

watch(
  () => props.height,
  () => {
    if (props.height !== undefined) {
      height.value = props.height
    }
  },
  { immediate: true }
)

const styles = computed(() => {
  let s = ''
  if (props.width !== undefined) {
    s += `width:${width.value}px;`
  }
  if (props.height !== undefined) {
    s += `height:${props.height}px;`
  }
  return s
})

const resizableStyles = computed(() => {
  return isResizing.value || isEnterResizable.value
    ? 'background-color:var(--primary-color);'
    : ''
})

function enterResizable() {
  isLeaveResizable = false
  setTimeout(() => {
    if (!isLeaveResizable) {
      isEnterResizable.value = true
    }
  }, 300)
}

function leaveResizable() {
  isEnterResizable.value = false
  isLeaveResizable = true
}

function startMoveX(e: MouseEvent) {
  const originWidth = width.value
  const startPageX = e.pageX
  isResizing.value = true
  onMove({
    update(e) {
      const tempWidth = originWidth + e.pageX - startPageX
      if (props.maxWidth !== undefined && tempWidth > props.maxWidth) {
        width.value = props.maxWidth
      } else if (props.minWidth !== undefined && tempWidth < props.minWidth) {
        width.value = props.minWidth
      } else {
        width.value = tempWidth
      }
      emit('update:width', width.value)
    },
    end() {
      isResizing.value = false
    },
  })
}

function startMoveY(e: MouseEvent) {
  const originHeight = height.value
  const startPageY = e.pageY
  isResizing.value = true
  onMove({
    update(e) {
      const tempHeight = originHeight + e.pageY - startPageY
      if (props.maxHeight !== undefined && tempHeight > props.maxHeight) {
        height.value = props.maxHeight
      } else if (
        props.minHeight !== undefined &&
        tempHeight < props.minHeight
      ) {
        height.value = props.minHeight
      } else {
        height.value = tempHeight
      }
      emit('update:height', height.value)
    },
    end() {
      isResizing.value = false
    },
  })
}
</script>

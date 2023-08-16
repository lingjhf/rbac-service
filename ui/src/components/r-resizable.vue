<template>
  <div class="relative" :style="styles">
    <slot />
    <div
      v-show="props.top && props.height !== undefined"
      class="cursor-n-resize h-1 -top-[2px] right-0 left-0 select-none absolute"
      :style="resizableStyles"
      @mouseenter="enterResizable"
      @mouseleave="leaveResizable"
      @mousedown="startMoveTop"
    >
      <slot name="top" />
    </div>
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
  height?: number
  top?: boolean
  right?: boolean
  bottom?: boolean
  left?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:width', value: number): void
  (e: 'update:height', value: number): void
}>()

const isResizing = ref(false)
const isEnterResizable = ref(false)
let isLeaveResizable = true

const styles = computed(() => {
  let s = ''
  if (props.width !== undefined) {
    s += `width:${props.width}px;`
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
  const originWidth = props.width
  if (originWidth === undefined) return
  const startPageX = e.pageX
  isResizing.value = true
  onMove({
    update(e) {
      const width = originWidth + e.pageX - startPageX
      emit('update:width', width)
    },
    end() {
      isResizing.value = false
    },
  })
}

function startMoveY(e: MouseEvent) {
  const originHeight = props.height
  if (originHeight === undefined) return
  const startPageY = e.pageY
  isResizing.value = true
  onMove({
    update(e) {
      const height = originHeight + e.pageY - startPageY
      emit('update:height', height)
    },
    end() {
      isResizing.value = false
    },
  })
}

function startMoveTop(e: MouseEvent) {
  const originHeight = props.height
  if (originHeight === undefined) return
  const startPageY = e.pageY
  isResizing.value = true
  onMove({
    update(e) {
      const height = originHeight - (e.pageY - startPageY)
      emit('update:height', height)
    },
    end() {
      isResizing.value = false
    },
  })
}
</script>

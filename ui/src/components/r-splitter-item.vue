<template>
  <r-resizable
    v-if="index > 0"
    top
    :height="height"
    @update:height="updateHeight"
  >
    <div class="border-t border-t-solid border-[var(--split-border-color)]">
      <div :style="`height:${foldedHeight}px`">
        <div v-if="expand" class="h-6 text-black w-6 i-local:arrow-down" />
        <div v-else class="h-6 text-black w-6 i-local:arrow-up" />
        {{ height }}
      </div>
      <slot />
    </div>
  </r-resizable>
  <div v-else :style="`height:${height}px`">
    <span class="text-black">{{ height }}</span>
  </div>
</template>

<script setup lang="ts">
import { inject, ref, watch } from 'vue'
import RResizable from './r-resizable.vue'
import { SplitterProvider } from './r-splitter.vue'

const splitterProvider = inject<SplitterProvider>('splitterProvider')
const props = defineProps<{
  minHeight?: number
  foldedHeight?: number
  expand?: boolean
}>()

let minHeight = 100
const foldedHeight = ref(36)
const height = ref(minHeight)
const expand = ref(false)
const index = ref(
  splitterProvider?.addItem({
    height,
    minHeight: 100,
  }) ?? 0
)

watch(
  () => props.minHeight,
  () => {
    if (props.minHeight && props.minHeight > 0) {
      minHeight = props.minHeight
      splitterProvider?.setMinHeight(index.value, props.minHeight)
    }
  },
  {
    immediate: true,
  }
)

watch(
  () => props.foldedHeight,
  () => {
    if (props.foldedHeight && props.foldedHeight > 0) {
      foldedHeight.value = props.foldedHeight
    }
  },
  {
    immediate: true,
  }
)

watch(
  () => props.expand,
  () => {
    expand.value = props.expand ?? false
    // if (expand.value) {
    //   splitterProvider?.setMinHeight(index.value, minHeight)
    //   splitterProvider?.setHeight(index.value, minHeight)
    // } else {
    //   if (index.value > 0) {
    //     splitterProvider?.setMinHeight(index.value, foldedHeight.value)
    //     splitterProvider?.setHeight(index.value, foldedHeight.value)
    //   }
    // }
  },
  {
    immediate: true,
  }
)

function updateHeight(value: number) {
  if (!expand.value) {
    if (value > foldedHeight.value * 2) {
      expand.value = true
      splitterProvider?.setMinHeight(index.value, minHeight)
      splitterProvider?.setHeight(index.value, value)
    }
    return
  }
  if (expand.value && value < foldedHeight.value) {
    expand.value = false
    splitterProvider?.setMinHeight(index.value, foldedHeight.value)
    splitterProvider?.setHeight(index.value, foldedHeight.value)
    return
  }
  splitterProvider?.setHeight(index.value, value)
}

</script>

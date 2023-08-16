<template>
  <div
    class="border-t border-t-solid border-[var(--split-border-color)] flex h-9 px-2 items-center"
  >
    <r-button class="h-7! p-0! w-7!" text @click="expand">
      <span v-if="props.expand" class="h-6 w-6 i-local:arrow-left-double" />
      <span v-else class="h-6 w-6 i-local:arrow-right-double" />
    </r-button>
    <r-button class="ml-auto! h-7! p-0! w-7!" text>
      <span class="h-5 w-5 i-local:help" />
    </r-button>
    <r-button class="h-7! p-0! w-7!" text>
      <span class="h-5 w-5 i-local:settings" />
    </r-button>
  </div>
</template>

<script setup lang="ts">
import { RButton } from '@/components'

const props = defineProps<{ expand?: boolean }>()
const emit = defineEmits<{
  (e: 'update:expand', value: boolean): void
}>()

function expand() {
  emit('update:expand', !props.expand)
}
</script>

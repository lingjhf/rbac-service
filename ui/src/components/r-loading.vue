<template>
  <slot />
</template>

<script setup lang="ts">
import { ElLoading } from 'element-plus'
import { watch } from 'vue'

const props = defineProps<{
  loading?: () => Promise<unknown>
}>()

watch(
  () => props.loading,
  async () => {
    if (props.loading) {
      const service = ElLoading.service({
        lock: true,
        text: 'Loading',
        background: '#1b1c21',
      })
      try {
        await props.loading()
      } catch (e) {
        service.close()
        throw new Error()
      }
      service.close()
    }
  },
  {
    immediate: true,
  }
)
</script>

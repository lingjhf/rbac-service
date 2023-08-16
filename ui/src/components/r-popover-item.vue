<template>
  <div>
    <r-divider v-if="props.divider" />
    <div class="rounded-1 h-6 mx-1 mt-[1px] py-1 px-3 hover:bg-[#8276c914]">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import RDivider from './r-divider.vue'
const props = defineProps<{ divider?: boolean }>()
</script>

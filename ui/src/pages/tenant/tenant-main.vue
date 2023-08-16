<template>
  <r-tabs
    v-if="tabsStore.currentTab"
    type="border-card"
    class="flex flex-col h-full border-none!"
    :model-value="tabsStore.currentTab.id"
    @update:model-value="tabChange"
    @tab-remove="tabRemove"
  >
    <r-tab-pane
      v-for="tab in tabsStore.tabs"
      :key="tab.id"
      :label="tab.label"
      :name="tab.id"
    >
      <!-- <router-view /> -->
      <component :is="tab.component" />
    </r-tab-pane>
  </r-tabs>
  <div v-else class="h-full" />
</template>

<script setup lang="ts">
import { TabPaneName } from 'element-plus'
import { RTabs, RTabPane } from '@/components'
import { useTabsStore } from '@/store'

const tabsStore = useTabsStore()

function tabChange(name: TabPaneName) {
  tabsStore.setCurrentTab(name as string)
}

function tabRemove(name: TabPaneName) {
  tabsStore.removeTab(name as string)
}
</script>

<template>
  <div class="bg-[var(--sidebar-bg-color)] flex flex-col h-full">
    <div class="flex mx-2 mt-1 mb-3">
      <r-input v-model="searchValue" />
      <r-button
        type="primary"
        class="flex-shrink-0 ml-2 h-8! p-0! w-8!"
        @click="addTenant"
      >
        <div class="h-5 w-5 i-local:add" />
      </r-button>
      <create-tenant-dialog
        v-model="dialogVisible"
        :parent-id="parentId"
        @submit="submit"
      />
    </div>
    <div class="h-full" ref="treeEl">
      <r-tree
        class="bg-[var(--sidebar-bg-color)]!"
        :data="tenantStore.tenantTree"
        :props="tenantTreeprops"
        :height="treeHeight"
        :item-size="32"
        highlight-current
        :expand-on-click-node="false"
        :current-node-key="currentNodeKey"
        @node-click="nodeClick"
      >
        <template #default="{ node }">
          <span>{{ node.label }}</span>
          <r-popover trigger="click">
            <r-popover-item>重命名</r-popover-item>
            <r-popover-item @click="addTenantChild(node)">
              添加子租户
            </r-popover-item>
            <r-popover-item divider>删除</r-popover-item>
            <template #reference>
              <r-button class="ml-auto mr-2 h-8! p-0! w-8!" text @click.stop>
                <span class="i-local:more" />
              </r-button>
            </template>
          </r-popover>
        </template>
      </r-tree>
    </div>
  </div>
</template>

<script setup lang="ts">
import { shallowRef, ref, watch } from 'vue'
import { RInput, RTree, RButton, RPopover, RPopoverItem } from '@/components'
import { CreateTenantDialog } from '@/pages/components'
import { useTenantStore, useTabsStore } from '@/store'
import TenantContent from './tenant-content.vue'

const tenantStore = useTenantStore()
const tabsStore = useTabsStore()

const treeEl = shallowRef<HTMLElement>()
const treeHeight = ref(0)
const dialogVisible = ref(false)
const parentId = ref<string | undefined | null>()
const searchValue = ref('')
const currentNodeKey = ref('')
const tenantTreeprops = {
  value: 'id',
  label: 'name',
  children: 'children',
}

watch(treeEl, () => {
  treeHeight.value = (treeEl.value?.offsetHeight ?? 0) - 1
})

function addTenant() {
  parentId.value = tenantStore.currentTenant?.id
  dialogVisible.value = true
}

function addTenantChild(node: any) {
  parentId.value = node.key
  dialogVisible.value = true
}

function nodeClick(data: any, node: any) {
  currentNodeKey.value = node.key
  tabsStore.addTab({
    id: node.key as string,
    label: node.label as string,
    component: shallowRef(TenantContent),
  })
}

function submit() {
  if (tenantStore.currentTenant) {
    tenantStore.getTenantChildren(tenantStore.currentTenant.id)
  }
}
</script>

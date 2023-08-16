import { defineStore } from 'pinia'
import { ref } from 'vue'
import * as api from '@/api'

export interface TenantTreeItem {
  id: string
  name: string
  parentId: string | null
  children?: TenantTreeItem[]
}

export const useTenantStore = defineStore('tenant', () => {
  const currentTenant = ref<api.TenantItem>()
  const rootTenantList = ref<api.TenantItem[]>([])
  const tenantTree = ref<TenantTreeItem[]>([])

  function setCurrentTenant(tenant: api.TenantItem) {
    currentTenant.value = tenant
  }

  async function getTenant(id: string) {
    const data = await api.getTenant(id)
    currentTenant.value = data
  }

  async function getRootTenantList() {
    const { list, count } = await api.getRootTenantList()
    rootTenantList.value = list
  }

  async function getTenantChildren(parentId: string) {
    const { list } = await api.getTenantChildren()
    tenantTree.value = generateTenantTree(list, parentId)
  }

  function generateTenantTree(
    data: api.TenantItem[],
    parentId: string | null
  ): TenantTreeItem[] {
    const nextData = [...data]
    const resultData: TenantTreeItem[] = []
    for (let i = 0; i < data.length; i++) {
      const item = data[i]
      if (item.parentId === parentId) {
        const resultItem = {
          id: item.id,
          name: item.name,
          parentId: item.parentId,
          children: generateTenantTree(nextData, item.id),
        }
        nextData.splice(i, 1)
        resultData.push(resultItem)
      }
    }

    return resultData
  }

  return {
    currentTenant,
    tenantTree,
    setCurrentTenant,
    rootTenantList,
    getTenant,
    getRootTenantList,
    getTenantChildren,
  }
})

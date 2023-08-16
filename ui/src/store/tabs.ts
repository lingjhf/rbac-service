import { defineStore } from 'pinia'
import { Component, ref } from 'vue'

export interface Tab {
  id: string
  label: string
  component: Component
  visibleMenu?: boolean
}

export const useTabsStore = defineStore('tabs', () => {
  const tabsMap = new Map<string, Tab>()
  const tabs = ref<Tab[]>([])
  const currentTab = ref<Tab>()

  function addTab(tab: Tab) {
    const index = tabs.value.findIndex((item) => item.id === tab.id)
    if (index === -1) {
      tabsMap.set(tab.id, tab)
      tabs.value.push(tab)
    }
    currentTab.value = tab
  }

  function setCurrentTab(id: string) {
    const tab = tabsMap.get(id)
    if (tab) {
      currentTab.value = tab
    }
  }

  function removeTab(id: string) {
    const tab = tabsMap.get(id)
    if (tab) {
      tabsMap.delete(id)
      if (currentTab.value?.id === id) {
        const index = tabs.value.findIndex((tab) => tab.id === id)
        if (index > -1) {
          setCurrentTab(tabs.value[index + 1]?.id || tabs.value[index - 1]?.id)
        }
      }
      tabs.value = [...tabs.value].filter((tab) => tab.id !== id)
    }
  }

  function visibleMenu(id: string) {
    tabs.value = tabs.value.map((item) => {
      item.visibleMenu = item.id === id
      return item
    })
  }

  return { tabs, currentTab, addTab, removeTab, setCurrentTab, visibleMenu }
})

<template>
  <div class="pt-6">
    <delete-confrim-dialog
      v-if="currentTenantItem"
      v-model="deleteConfirmDialogVisible"
      :tenant="currentTenantItem"
    />
    <update-tenant-dialog
      v-if="currentTenantItem"
      v-model="updateTenantDialogVisible"
      :tenant="currentTenantItem"
    />
    <div
      v-for="item in tenantStore.rootTenantList"
      :key="item.id"
      class="h-[240px] w-[100%] justify-center inline-flex sm:w-[calc(100%/2)] md:w-[calc(100%/2)] lg:w-[calc(100%/3)] xl:w-[calc(100%/4)] 2xl:w-[calc(100%/4)]"
    >
      <div
        class="bg-[var(--input-bg-color)] cursor-pointer rounded-[4px] h-[200px] w-[300px]"
        @click="enterTenant(item)"
      >
        <div class="flex py-4 px-5 items-center">
          {{ item.name }}
          <r-popover>
            <r-button @click="updateTenant(item)"> 修改名称 </r-button>
            <r-button divided @click="deleteTenant(item)"> 删除租户 </r-button>
            <template #reference>
              <r-button class="ml-auto h-7! p-0! w-7!" text>
                <span class="h-6 w-6 i-local:more" />
              </r-button>
            </template>
          </r-popover>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useTenantStore } from '@/store'
import { RButton, RPopover } from '@/components'
import DeleteConfrimDialog from './delete-confirm-dialog.vue'
import UpdateTenantDialog from './update-tenant-dialog.vue'
import { TenantItem } from '@/api'
import { useRouter } from 'vue-router'
import { RouteName } from '@/router'

const router = useRouter()
const tenantStore = useTenantStore()

const currentTenantItem = ref<TenantItem>()
const deleteConfirmDialogVisible = ref(false)
const updateTenantDialogVisible = ref(false)

onMounted(() => {
  tenantStore.getRootTenantList()
})

function enterTenant(item: TenantItem) {
  router.push({ name: RouteName.tenant, params: { id: item.id } })
}

function updateTenant(item: TenantItem) {
  currentTenantItem.value = item
  updateTenantDialogVisible.value = true
}

function deleteTenant(item: TenantItem) {
  currentTenantItem.value = item
  deleteConfirmDialogVisible.value = true
}
</script>

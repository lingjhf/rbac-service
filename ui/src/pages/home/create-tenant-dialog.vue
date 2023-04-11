<template>
  <r-dialog
    :model-value="props.modelValue"
    width="400"
    title="新建租户"
    @update:model-value="updateModelValue"
    @open="open(ruleFormRef)"
  >
    <el-form
      ref="ruleFormRef"
      label-position="top"
      :model="form"
      :rules="rules"
    >
      <el-form-item label="名称" prop="name" :error="formErrorMessage.name">
        <r-input v-model="form.name" />
      </el-form-item>
    </el-form>
    <template #footer>
      <r-button @click="cancel">取消</r-button>
      <r-button
        :loading="confirmLoading"
        type="primary"
        @click="confirm(ruleFormRef)"
      >
        <span v-show="!confirmLoading">确定</span>
      </r-button>
    </template>
  </r-dialog>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { ElForm, ElFormItem, ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { RButton, RInput, RDialog } from '@/components'
import { createTenant, CreateTenantForm } from '@/api'

const props = defineProps<{ modelValue: boolean }>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const confirmLoading = ref(false)
const ruleFormRef = ref<FormInstance>()
const form = ref<CreateTenantForm>({ name: '', parentId: null })
const rules = ref<FormRules>({
  name: [{ required: true, message: '名称不能为空' }],
})
const formErrorMessage = reactive({
  name: '',
})

function emitUpdateModelValue(value: boolean) {
  emit('update:modelValue', value)
}

function updateModelValue(value: boolean) {
  emitUpdateModelValue(value)
}

async function open(formEl: FormInstance | undefined) {
  formEl?.resetFields()
}

function cancel() {
  emitUpdateModelValue(false)
}

async function confirm(formEl: FormInstance | undefined) {
  confirmLoading.value = true
  if (!formEl) return
  try {
    await formEl.validate((valid) => {
      if (!valid) {
        throw new Error('error')
      }
    })
  } catch (e) {
    return
  }
  try {
    await createTenant(form.value)
  } catch (e) {
    if (typeof e === 'object') {
      const err = e as { name: string }
      formErrorMessage.name = err.name
    }
    confirmLoading.value = false
    return
  }
  confirmLoading.value = false
  ElMessage({ message: '新建租户成功', type: 'success' })
  emitUpdateModelValue(false)
}
</script>

<style scoped></style>

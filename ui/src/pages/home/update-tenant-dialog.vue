<template>
  <r-dialog
    :model-value="props.modelValue"
    width="400"
    title="修改名称"
    @update:model-value="updateModelValue"
    @open="open(ruleFormRef)"
  >
    <el-form
      ref="ruleFormRef"
      label-position="top"
      :model="form"
      :rules="rules"
    >
      <el-form-item label="名称" prop="name">
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
        <span v-show="!confirmLoading">保存</span>
      </r-button>
    </template>
  </r-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import {
  ElForm,
  ElFormItem,
  FormInstance,
  FormRules,
  ElMessage,
} from 'element-plus'
import { RDialog, RButton, RInput } from '@/components'
import { TenantItem } from '@/api'
const props = defineProps<{ modelValue: boolean; tenant?: TenantItem }>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const confirmLoading = ref(false)
const form = ref({ name: '' })
const ruleFormRef = ref<FormInstance>()
const rules = ref<FormRules>({
  name: [{ required: true, message: '名称不能为空' }],
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
  confirmLoading.value = true
  try {
  } catch (e) {
    confirmLoading.value = false
    return
  }
  confirmLoading.value = false
  ElMessage({ message: '已保存', type: 'success' })
  emitUpdateModelValue(false)
}
</script>

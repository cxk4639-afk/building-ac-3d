<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { useMessage } from 'naive-ui'
import { createTenantApi, deleteTenantApi, getTenantsApi, updateTenantApi } from '@/api/system'
import type { TenantRecord } from '@/types/system'

const message = useMessage()
const loading = ref(false)
const showModal = ref(false)
const isEdit = ref(false)
const list = ref<TenantRecord[]>([])

const form = reactive<Partial<TenantRecord>>({
  id: undefined,
  name: '',
  code: '',
  status: 1,
  remark: '',
})

function resetForm() {
  form.id = undefined
  form.name = ''
  form.code = ''
  form.status = 1
  form.remark = ''
}

async function loadData() {
  loading.value = true
  try {
    list.value = await getTenantsApi()
  } finally {
    loading.value = false
  }
}

function openCreate() {
  resetForm()
  isEdit.value = false
  showModal.value = true
}

function openEdit(row: TenantRecord) {
  Object.assign(form, row)
  isEdit.value = true
  showModal.value = true
}

async function handleSubmit() {
  if (!form.name || !form.code) {
    message.warning('请填写租户名称和租户编码')
    return
  }

  if (isEdit.value) {
    await updateTenantApi(form)
    message.success('更新成功')
  } else {
    await createTenantApi(form)
    message.success('新增成功')
  }

  showModal.value = false
  await loadData()
}

async function handleDelete(row: TenantRecord) {
  if (row.id === 1) {
    message.warning('默认租户不建议删除')
    return
  }
  await deleteTenantApi(row.id)
  message.success('删除成功')
  await loadData()
}

onMounted(loadData)
</script>

<template>
  <n-card title="租户管理">
    <template #header-extra>
      <n-button type="primary" @click="openCreate">新增租户</n-button>
    </template>

    <n-spin :show="loading">
      <n-table :bordered="false" :single-line="false">
        <thead>
          <tr>
            <th>ID</th>
            <th>租户名称</th>
            <th>租户编码</th>
            <th>状态</th>
            <th>备注</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in list" :key="item.id">
            <td>{{ item.id }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.code }}</td>
            <td>
              <n-tag :type="item.status === 1 ? 'success' : 'warning'">
                {{ item.status === 1 ? '启用' : '停用' }}
              </n-tag>
            </td>
            <td>{{ item.remark }}</td>
            <td>{{ item.createdAt }}</td>
            <td>
              <n-space>
                <n-button size="small" @click="openEdit(item)">编辑</n-button>
                <n-popconfirm @positive-click="handleDelete(item)">
                  <template #trigger>
                    <n-button size="small" type="error">删除</n-button>
                  </template>
                  确认删除该租户吗？
                </n-popconfirm>
              </n-space>
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-spin>

    <n-modal v-model:show="showModal" preset="card" :title="isEdit ? '编辑租户' : '新增租户'" style="width: 520px">
      <n-form label-placement="left" label-width="90">
        <n-form-item label="租户名称">
          <n-input v-model:value="form.name" placeholder="请输入租户名称" />
        </n-form-item>
        <n-form-item label="租户编码">
          <n-input v-model:value="form.code" placeholder="请输入租户编码，如 default" />
        </n-form-item>
        <n-form-item label="状态">
          <n-select v-model:value="form.status" :options="[{ label: '启用', value: 1 }, { label: '停用', value: 2 }]" />
        </n-form-item>
        <n-form-item label="备注">
          <n-input v-model:value="form.remark" type="textarea" placeholder="请输入备注" />
        </n-form-item>
      </n-form>
      <template #footer>
        <n-space justify="end">
          <n-button @click="showModal = false">取消</n-button>
          <n-button type="primary" @click="handleSubmit">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </n-card>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useMessage } from 'naive-ui'
import {
  createUserApi,
  deleteUserApi,
  getRolesApi,
  getTenantsApi,
  getUsersApi,
  updateUserApi,
} from '@/api/system'
import type { RoleRecord, TenantRecord, UserRecord } from '@/types/system'

const message = useMessage()
const loading = ref(false)
const showModal = ref(false)
const isEdit = ref(false)
const list = ref<UserRecord[]>([])
const tenants = ref<TenantRecord[]>([])
const roles = ref<RoleRecord[]>([])

const form = reactive<Partial<UserRecord>>({
  id: undefined,
  tenantId: 1,
  username: '',
  password: '',
  nickname: '',
  roleIds: [],
  status: 1,
})

const tenantOptions = computed(() => tenants.value.map((item) => ({ label: item.name, value: item.id })))
const roleOptions = computed(() => roles.value.map((item) => ({ label: item.name, value: item.id })))

function tenantName(tenantId: number) {
  return tenants.value.find((item) => item.id === tenantId)?.name || '-'
}

function roleNames(roleIds: number[]) {
  return roleIds.map((id) => roles.value.find((item) => item.id === id)?.name).filter(Boolean).join('、') || '-'
}

function resetForm() {
  form.id = undefined
  form.tenantId = tenants.value[0]?.id || 1
  form.username = ''
  form.password = ''
  form.nickname = ''
  form.roleIds = []
  form.status = 1
}

async function loadData() {
  loading.value = true
  try {
    const [tenantData, roleData, userData] = await Promise.all([
      getTenantsApi(),
      getRolesApi(),
      getUsersApi(),
    ])
    tenants.value = tenantData
    roles.value = roleData
    list.value = userData
  } finally {
    loading.value = false
  }
}

function openCreate() {
  resetForm()
  isEdit.value = false
  showModal.value = true
}

function openEdit(row: UserRecord) {
  Object.assign(form, { ...row, password: '' })
  isEdit.value = true
  showModal.value = true
}

async function handleSubmit() {
  if (!form.tenantId || !form.username) {
    message.warning('请选择租户并填写用户名')
    return
  }

  if (isEdit.value) {
    await updateUserApi(form)
    message.success('更新成功')
  } else {
    await createUserApi(form)
    message.success('新增成功')
  }

  showModal.value = false
  await loadData()
}

async function handleDelete(row: UserRecord) {
  if (row.id === 1) {
    message.warning('默认管理员不建议删除')
    return
  }
  await deleteUserApi(row.id)
  message.success('删除成功')
  await loadData()
}

onMounted(loadData)
</script>

<template>
  <n-card title="用户管理">
    <template #header-extra>
      <n-button type="primary" @click="openCreate">新增用户</n-button>
    </template>

    <n-spin :show="loading">
      <n-table :bordered="false" :single-line="false">
        <thead>
          <tr>
            <th>ID</th>
            <th>所属租户</th>
            <th>用户名</th>
            <th>昵称</th>
            <th>角色</th>
            <th>状态</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in list" :key="item.id">
            <td>{{ item.id }}</td>
            <td>{{ tenantName(item.tenantId) }}</td>
            <td>{{ item.username }}</td>
            <td>{{ item.nickname }}</td>
            <td>{{ roleNames(item.roleIds || []) }}</td>
            <td>
              <n-tag :type="item.status === 1 ? 'success' : 'warning'">
                {{ item.status === 1 ? '启用' : '停用' }}
              </n-tag>
            </td>
            <td>{{ item.createdAt }}</td>
            <td>
              <n-space>
                <n-button size="small" @click="openEdit(item)">编辑</n-button>
                <n-popconfirm @positive-click="handleDelete(item)">
                  <template #trigger>
                    <n-button size="small" type="error">删除</n-button>
                  </template>
                  确认删除该用户吗？
                </n-popconfirm>
              </n-space>
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-spin>

    <n-modal v-model:show="showModal" preset="card" :title="isEdit ? '编辑用户' : '新增用户'" style="width: 560px">
      <n-form label-placement="left" label-width="90">
        <n-form-item label="所属租户">
          <n-select v-model:value="form.tenantId" :options="tenantOptions" />
        </n-form-item>
        <n-form-item label="用户名">
          <n-input v-model:value="form.username" placeholder="请输入用户名" />
        </n-form-item>
        <n-form-item label="密码">
          <n-input v-model:value="form.password" type="password" show-password-on="click" :placeholder="isEdit ? '不填则不修改密码' : '默认 123456'" />
        </n-form-item>
        <n-form-item label="昵称">
          <n-input v-model:value="form.nickname" placeholder="请输入昵称" />
        </n-form-item>
        <n-form-item label="角色">
          <n-select v-model:value="form.roleIds" multiple :options="roleOptions" />
        </n-form-item>
        <n-form-item label="状态">
          <n-select v-model:value="form.status" :options="[{ label: '启用', value: 1 }, { label: '停用', value: 2 }]" />
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

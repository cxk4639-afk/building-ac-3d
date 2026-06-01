<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useMessage } from 'naive-ui'
import {
  createRoleApi,
  deleteRoleApi,
  getMenusApi,
  getRolesApi,
  getTenantsApi,
  updateRoleApi,
} from '@/api/system'
import type { MenuRecord, RoleRecord, TenantRecord } from '@/types/system'

const message = useMessage()
const loading = ref(false)
const showModal = ref(false)
const isEdit = ref(false)
const list = ref<RoleRecord[]>([])
const tenants = ref<TenantRecord[]>([])
const menus = ref<MenuRecord[]>([])

const form = reactive<Partial<RoleRecord>>({
  id: undefined,
  tenantId: 1,
  name: '',
  code: '',
  menuIds: [],
  status: 1,
})

const tenantOptions = computed(() => tenants.value.map((item) => ({ label: item.name, value: item.id })))
const menuOptions = computed(() => menus.value.map((item) => ({ label: item.title, value: item.id })))

function tenantName(tenantId: number) {
  return tenants.value.find((item) => item.id === tenantId)?.name || '-'
}

function menuNames(menuIds: number[]) {
  return menuIds.map((id) => menus.value.find((item) => item.id === id)?.title).filter(Boolean).join('、') || '-'
}

function resetForm() {
  form.id = undefined
  form.tenantId = tenants.value[0]?.id || 1
  form.name = ''
  form.code = ''
  form.menuIds = []
  form.status = 1
}

async function loadData() {
  loading.value = true
  try {
    const [tenantData, menuData, roleData] = await Promise.all([
      getTenantsApi(),
      getMenusApi(),
      getRolesApi(),
    ])
    tenants.value = tenantData
    menus.value = menuData
    list.value = roleData
  } finally {
    loading.value = false
  }
}

function openCreate() {
  resetForm()
  isEdit.value = false
  showModal.value = true
}

function openEdit(row: RoleRecord) {
  Object.assign(form, row)
  isEdit.value = true
  showModal.value = true
}

async function handleSubmit() {
  if (!form.tenantId || !form.name || !form.code) {
    message.warning('请选择租户并填写角色名称和角色编码')
    return
  }

  if (isEdit.value) {
    await updateRoleApi(form)
    message.success('更新成功')
  } else {
    await createRoleApi(form)
    message.success('新增成功')
  }

  showModal.value = false
  await loadData()
}

async function handleDelete(row: RoleRecord) {
  if (row.id === 1) {
    message.warning('默认超级管理员不建议删除')
    return
  }
  await deleteRoleApi(row.id)
  message.success('删除成功')
  await loadData()
}

onMounted(loadData)
</script>

<template>
  <n-card title="角色管理">
    <template #header-extra>
      <n-button type="primary" @click="openCreate">新增角色</n-button>
    </template>

    <n-spin :show="loading">
      <n-table :bordered="false" :single-line="false">
        <thead>
          <tr>
            <th>ID</th>
            <th>所属租户</th>
            <th>角色名称</th>
            <th>角色编码</th>
            <th>菜单权限</th>
            <th>状态</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in list" :key="item.id">
            <td>{{ item.id }}</td>
            <td>{{ tenantName(item.tenantId) }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.code }}</td>
            <td>{{ menuNames(item.menuIds || []) }}</td>
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
                  确认删除该角色吗？
                </n-popconfirm>
              </n-space>
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-spin>

    <n-modal v-model:show="showModal" preset="card" :title="isEdit ? '编辑角色' : '新增角色'" style="width: 620px">
      <n-form label-placement="left" label-width="90">
        <n-form-item label="所属租户">
          <n-select v-model:value="form.tenantId" :options="tenantOptions" />
        </n-form-item>
        <n-form-item label="角色名称">
          <n-input v-model:value="form.name" placeholder="请输入角色名称" />
        </n-form-item>
        <n-form-item label="角色编码">
          <n-input v-model:value="form.code" placeholder="如 super_admin" />
        </n-form-item>
        <n-form-item label="菜单权限">
          <n-select v-model:value="form.menuIds" multiple :options="menuOptions" />
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

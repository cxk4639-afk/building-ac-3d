<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useMessage } from 'naive-ui'
import { createMenuApi, deleteMenuApi, getMenusApi, updateMenuApi } from '@/api/system'
import type { MenuRecord } from '@/types/system'

const message = useMessage()
const loading = ref(false)
const showModal = ref(false)
const isEdit = ref(false)
const list = ref<MenuRecord[]>([])

const form = reactive<Partial<MenuRecord>>({
  id: undefined,
  parentId: 0,
  title: '',
  path: '',
  component: '',
  permission: '',
  type: 'menu',
  sort: 0,
  visible: true,
})

const parentOptions = computed(() => [
  { label: '根目录', value: 0 },
  ...list.value.map((item) => ({ label: `${item.title}（ID: ${item.id}）`, value: item.id })),
])

const typeOptions = [
  { label: '目录', value: 'catalog' },
  { label: '菜单', value: 'menu' },
  { label: '按钮', value: 'button' },
]

function parentName(parentId: number) {
  if (parentId === 0) return '根目录'
  return list.value.find((item) => item.id === parentId)?.title || '-'
}

function resetForm() {
  form.id = undefined
  form.parentId = 0
  form.title = ''
  form.path = ''
  form.component = ''
  form.permission = ''
  form.type = 'menu'
  form.sort = 0
  form.visible = true
}

async function loadData() {
  loading.value = true
  try {
    list.value = await getMenusApi()
  } finally {
    loading.value = false
  }
}

function openCreate() {
  resetForm()
  isEdit.value = false
  showModal.value = true
}

function openEdit(row: MenuRecord) {
  Object.assign(form, row)
  isEdit.value = true
  showModal.value = true
}

async function handleSubmit() {
  if (!form.title || !form.type) {
    message.warning('请填写菜单名称和菜单类型')
    return
  }

  if (isEdit.value) {
    await updateMenuApi(form)
    message.success('更新成功')
  } else {
    await createMenuApi(form)
    message.success('新增成功')
  }

  showModal.value = false
  await loadData()
}

async function handleDelete(row: MenuRecord) {
  const hasChildren = list.value.some((item) => item.parentId === row.id)
  if (hasChildren) {
    message.warning('该菜单存在子级，请先删除子级')
    return
  }

  await deleteMenuApi(row.id)
  message.success('删除成功')
  await loadData()
}

onMounted(loadData)
</script>

<template>
  <n-card title="菜单管理">
    <template #header-extra>
      <n-button type="primary" @click="openCreate">新增菜单</n-button>
    </template>

    <n-spin :show="loading">
      <n-table :bordered="false" :single-line="false">
        <thead>
          <tr>
            <th>ID</th>
            <th>上级菜单</th>
            <th>菜单名称</th>
            <th>类型</th>
            <th>路由地址</th>
            <th>组件路径</th>
            <th>权限标识</th>
            <th>排序</th>
            <th>显示</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in list" :key="item.id">
            <td>{{ item.id }}</td>
            <td>{{ parentName(item.parentId) }}</td>
            <td>{{ item.title }}</td>
            <td>
              <n-tag>{{ item.type }}</n-tag>
            </td>
            <td>{{ item.path }}</td>
            <td>{{ item.component }}</td>
            <td>{{ item.permission }}</td>
            <td>{{ item.sort }}</td>
            <td>
              <n-tag :type="item.visible ? 'success' : 'warning'">
                {{ item.visible ? '显示' : '隐藏' }}
              </n-tag>
            </td>
            <td>
              <n-space>
                <n-button size="small" @click="openEdit(item)">编辑</n-button>
                <n-popconfirm @positive-click="handleDelete(item)">
                  <template #trigger>
                    <n-button size="small" type="error">删除</n-button>
                  </template>
                  确认删除该菜单吗？
                </n-popconfirm>
              </n-space>
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-spin>

    <n-modal v-model:show="showModal" preset="card" :title="isEdit ? '编辑菜单' : '新增菜单'" style="width: 640px">
      <n-form label-placement="left" label-width="90">
        <n-form-item label="上级菜单">
          <n-select v-model:value="form.parentId" :options="parentOptions" />
        </n-form-item>
        <n-form-item label="菜单名称">
          <n-input v-model:value="form.title" placeholder="请输入菜单名称" />
        </n-form-item>
        <n-form-item label="菜单类型">
          <n-select v-model:value="form.type" :options="typeOptions" />
        </n-form-item>
        <n-form-item label="路由地址">
          <n-input v-model:value="form.path" placeholder="如 /system/user" />
        </n-form-item>
        <n-form-item label="组件路径">
          <n-input v-model:value="form.component" placeholder="如 system/user/index" />
        </n-form-item>
        <n-form-item label="权限标识">
          <n-input v-model:value="form.permission" placeholder="如 system:user:list" />
        </n-form-item>
        <n-form-item label="排序">
          <n-input-number v-model:value="form.sort" :min="0" />
        </n-form-item>
        <n-form-item label="是否显示">
          <n-switch v-model:value="form.visible" />
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

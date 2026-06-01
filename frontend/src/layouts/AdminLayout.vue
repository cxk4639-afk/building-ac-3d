<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router'
import type { MenuOption } from 'naive-ui'

import { useAuthStore } from '@/stores/modules/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const menuOptions = computed<MenuOption[]>(() => [
  {
    label: () => <RouterLink to="/dashboard">工作台</RouterLink>,
    key: '/dashboard',
  },
  {
    label: '系统管理',
    key: 'system',
    children: [
      { label: () => <RouterLink to="/system/tenant">租户管理</RouterLink>, key: '/system/tenant' },
      { label: () => <RouterLink to="/system/user">用户管理</RouterLink>, key: '/system/user' },
      { label: () => <RouterLink to="/system/role">角色管理</RouterLink>, key: '/system/role' },
      { label: () => <RouterLink to="/system/menu">菜单管理</RouterLink>, key: '/system/menu' },
    ],
  },
  {
    label: () => <RouterLink to="/visual-config">3D配置</RouterLink>,
    key: '/visual-config',
  },
])

async function handleLogout() {
  await authStore.logout()
  await router.replace('/login')
}
</script>

<template>
  <n-layout class="admin-layout" has-sider>
    <n-layout-sider
      bordered
      collapse-mode="width"
      :collapsed-width="64"
      :width="240"
    >
      <div class="logo">building-ac-3d</div>
      <n-menu :value="route.path" :options="menuOptions" />
    </n-layout-sider>

    <n-layout>
      <n-layout-header class="header" bordered>
        <div>
          <strong>{{ route.meta.title || '工作台' }}</strong>
          <span class="tenant">当前租户：{{ authStore.userInfo?.tenant.name || '-' }}</span>
        </div>
        <n-space align="center">
          <span>{{ authStore.userInfo?.nickname || authStore.userInfo?.username }}</span>
          <n-button size="small" @click="handleLogout">退出登录</n-button>
        </n-space>
      </n-layout-header>

      <n-layout-content class="content">
        <RouterView />
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<style scoped>
.admin-layout {
  height: 100vh;
}

.logo {
  height: 56px;
  display: flex;
  align-items: center;
  padding: 0 18px;
  font-weight: 700;
  font-size: 18px;
}

.header {
  height: 56px;
  padding: 0 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
}

.tenant {
  margin-left: 16px;
  color: #6b7280;
  font-size: 13px;
}

.content {
  padding: 20px;
  background: #f5f7fb;
}
</style>

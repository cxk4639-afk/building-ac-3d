<script setup lang="ts">
import { computed } from 'vue'
import { RouterView, useRoute, useRouter } from 'vue-router'
import type { MenuOption } from 'naive-ui'

import { useAuthStore } from '@/stores/modules/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const menuOptions = computed<MenuOption[]>(() => [
  {
    label: '工作台',
    key: '/dashboard',
  },
  {
    label: '系统管理',
    key: 'system',
    children: [
      {
        label: '租户管理',
        key: '/system/tenant',
      },
      {
        label: '用户管理',
        key: '/system/user',
      },
      {
        label: '角色管理',
        key: '/system/role',
      },
      {
        label: '菜单管理',
        key: '/system/menu',
      },
    ],
  },
  {
    label: '3D配置',
    key: '/visual-config',
  },
])

function handleMenuUpdate(key: string | number) {
  const path = String(key)

  if (path.startsWith('/')) {
    router.push(path)
  }
}

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

      <n-menu
        :value="route.path"
        :options="menuOptions"
        @update:value="handleMenuUpdate"
      />
    </n-layout-sider>

    <n-layout>
      <n-layout-header class="header" bordered>
        <div>
          <strong>{{ route.meta.title || '工作台' }}</strong>
          <span class="tenant">
            当前租户：{{ authStore.userInfo?.tenant.name || '-' }}
          </span>
        </div>

        <n-space align="center">
          <span>
            {{ authStore.userInfo?.nickname || authStore.userInfo?.username }}
          </span>
          <n-button size="small" @click="handleLogout">
            退出登录
          </n-button>
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

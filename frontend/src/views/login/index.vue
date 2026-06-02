<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'

import { useAuthStore } from '@/stores/modules/auth'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const authStore = useAuthStore()

const loading = ref(false)
const form = reactive({
  tenantCode: 'default',
  username: 'admin',
  password: '123456',
})

async function handleLogin() {
  loading.value = true
  try {
    await authStore.login(form)
    message.success('登录成功')
    await router.replace((route.query.redirect as string) || '/dashboard')
  } catch (error) {
    message.error(error instanceof Error ? error.message : '登录失败')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <main class="login-page">
    <n-card class="login-card" title="building-ac-3d">
      <p class="subtitle">楼宇空调 3D 可视化配置平台</p>

      <n-form :model="form" label-placement="left" label-width="80">
        <n-form-item label="租户编码">
          <n-input v-model:value="form.tenantCode" placeholder="请输入租户编码" />
        </n-form-item>
        <n-form-item label="用户名">
          <n-input v-model:value="form.username" placeholder="请输入用户名" />
        </n-form-item>
        <n-form-item label="密码">
          <n-input
            v-model:value="form.password"
            type="password"
            show-password-on="click"
            placeholder="请输入密码"
            @keyup.enter="handleLogin"
          />
        </n-form-item>
        <n-button type="primary" block :loading="loading" @click="handleLogin">
          登录
        </n-button>
      </n-form>
    </n-card>
  </main>
</template>

<style scoped>
.login-page {
  min-height: 100vh;
  display: grid;
  place-items: center;
  background:
    radial-gradient(circle at top left, rgba(24, 160, 251, 0.18), transparent 36%),
    linear-gradient(135deg, #edf4ff 0%, #f7f9fc 48%, #eef2ff 100%);
}

.login-card {
  width: 420px;
}

.subtitle {
  margin: -6px 0 24px;
  color: #6b7280;
}
</style>

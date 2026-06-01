<script setup>
import { ref } from 'vue'

const dbHealth = ref(null)
const configs = ref([])
const loading = ref(false)
const errorMessage = ref('')

async function requestJSON(url, options) {
  const response = await fetch(url, options)
  const data = await response.json().catch(() => null)

  if (!response.ok) {
    throw new Error(data?.message || `请求失败：${response.status}`)
  }

  return data
}

async function checkDBHealth() {
  loading.value = true
  errorMessage.value = ''

  try {
    dbHealth.value = await requestJSON('/api/db/health')
  } catch (error) {
    errorMessage.value = error.message
  } finally {
    loading.value = false
  }
}

async function loadConfigs() {
  loading.value = true
  errorMessage.value = ''

  try {
    configs.value = await requestJSON('/api/visual-configs')
  } catch (error) {
    errorMessage.value = error.message
  } finally {
    loading.value = false
  }
}

async function createDemoConfig() {
  loading.value = true
  errorMessage.value = ''

  try {
    await requestJSON('/api/visual-configs', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: `测试配置 ${new Date().toLocaleString()}`,
        description: '用于验证前端、后端、MySQL 是否已经打通',
        configData: {
          building: {
            name: '示例楼宇',
            floors: 3,
          },
          devices: [
            {
              id: 'ac-001',
              type: 'air-conditioner',
              position: { x: 0, y: 1.2, z: 0 },
            },
          ],
        },
      }),
    })

    await loadConfigs()
  } catch (error) {
    errorMessage.value = error.message
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <main class="page">
    <section class="card">
      <p class="eyebrow">building-ac-3d</p>
      <h1>数据库连接</h1>
      <p class="description">
        当前页面用于验证 Vue 前端、Go 后端、MySQL 数据库是否已经连通。
      </p>

      <div class="actions">
        <button :disabled="loading" @click="checkDBHealth">检测数据库连接</button>
        <button :disabled="loading" @click="createDemoConfig">新增测试配置</button>
        <button :disabled="loading" @click="loadConfigs">查询配置列表</button>
      </div>

      <p v-if="loading" class="status">请求中...</p>
      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>

      <div v-if="dbHealth" class="result">
        <h2>连接状态</h2>
        <pre>{{ dbHealth }}</pre>
      </div>
    </section>

    <section class="card">
      <h2>配置列表</h2>
      <p v-if="configs.length === 0" class="empty">暂无配置数据。</p>

      <article v-for="item in configs" :key="item.id" class="config-item">
        <div>
          <strong>{{ item.name }}</strong>
          <p>{{ item.description || '无描述' }}</p>
        </div>
        <small>ID: {{ item.id }}</small>
      </article>
    </section>
  </main>
</template>

<style scoped>
.page {
  min-height: 100vh;
  padding: 48px;
  background: #f5f7fb;
  color: #1f2937;
}

.card {
  max-width: 960px;
  margin: 0 auto 24px;
  padding: 28px;
  border-radius: 18px;
  background: #ffffff;
  box-shadow: 0 16px 40px rgba(15, 23, 42, 0.08);
}

.eyebrow {
  margin: 0 0 8px;
  color: #2563eb;
  font-weight: 700;
}

h1,
h2 {
  margin: 0 0 14px;
}

.description,
.empty {
  color: #6b7280;
}

.actions {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin-top: 24px;
}

button {
  border: 0;
  border-radius: 10px;
  padding: 10px 16px;
  background: #2563eb;
  color: #ffffff;
  cursor: pointer;
}

button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.status {
  color: #2563eb;
}

.error {
  color: #dc2626;
}

.result {
  margin-top: 20px;
}

pre {
  overflow: auto;
  padding: 16px;
  border-radius: 12px;
  background: #111827;
  color: #d1d5db;
}

.config-item {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
  padding: 16px 0;
  border-top: 1px solid #e5e7eb;
}

.config-item p {
  margin: 6px 0 0;
  color: #6b7280;
}

.config-item small {
  color: #9ca3af;
  white-space: nowrap;
}
</style>

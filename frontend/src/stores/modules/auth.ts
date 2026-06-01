import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

import { getUserInfoApi, loginApi, logoutApi } from '@/api/auth'
import type { LoginParams, UserInfo } from '@/types/auth'

const TOKEN_KEY = 'building_ac_3d_token'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem(TOKEN_KEY) || '')
  const userInfo = ref<UserInfo | null>(null)

  const currentTenantId = computed(() => userInfo.value?.tenant.id)

  function setToken(value: string) {
    token.value = value
    localStorage.setItem(TOKEN_KEY, value)
  }

  async function login(params: LoginParams) {
    const result = await loginApi(params)
    setToken(result.token)
    await fetchUserInfo()
  }

  async function fetchUserInfo() {
    userInfo.value = await getUserInfoApi()
  }

  async function logout() {
    if (token.value) {
      await logoutApi().catch(() => null)
    }
    token.value = ''
    userInfo.value = null
    localStorage.removeItem(TOKEN_KEY)
  }

  function hasPermission(permission: string) {
    if (!permission) return true
    return userInfo.value?.permissions.includes('*:*:*') || userInfo.value?.permissions.includes(permission)
  }

  return {
    token,
    userInfo,
    currentTenantId,
    login,
    fetchUserInfo,
    logout,
    hasPermission,
  }
})

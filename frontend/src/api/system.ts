import { request } from './http'
import type { MenuRecord, RoleRecord, TenantRecord, UserRecord } from '@/types/system'

export function getTenantsApi() {
  return request<TenantRecord[]>({ url: '/system/tenants', method: 'GET' })
}

export function getUsersApi() {
  return request<UserRecord[]>({ url: '/system/users', method: 'GET' })
}

export function getRolesApi() {
  return request<RoleRecord[]>({ url: '/system/roles', method: 'GET' })
}

export function getMenusApi() {
  return request<MenuRecord[]>({ url: '/system/menus', method: 'GET' })
}

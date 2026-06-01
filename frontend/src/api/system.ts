import { request } from './http'
import type { MenuRecord, RoleRecord, TenantRecord, UserRecord } from '@/types/system'

export function getTenantsApi() {
  return request<TenantRecord[]>({ url: '/system/tenants', method: 'GET' })
}

export function createTenantApi(data: Partial<TenantRecord>) {
  return request<{ id: number }>({ url: '/system/tenants', method: 'POST', data })
}

export function updateTenantApi(data: Partial<TenantRecord>) {
  return request<boolean>({ url: '/system/tenants', method: 'PUT', data })
}

export function deleteTenantApi(id: number) {
  return request<boolean>({ url: `/system/tenants?id=${id}`, method: 'DELETE' })
}

export function getUsersApi() {
  return request<UserRecord[]>({ url: '/system/users', method: 'GET' })
}

export function createUserApi(data: Partial<UserRecord>) {
  return request<{ id: number }>({ url: '/system/users', method: 'POST', data })
}

export function updateUserApi(data: Partial<UserRecord>) {
  return request<boolean>({ url: '/system/users', method: 'PUT', data })
}

export function deleteUserApi(id: number) {
  return request<boolean>({ url: `/system/users?id=${id}`, method: 'DELETE' })
}

export function getRolesApi() {
  return request<RoleRecord[]>({ url: '/system/roles', method: 'GET' })
}

export function createRoleApi(data: Partial<RoleRecord>) {
  return request<{ id: number }>({ url: '/system/roles', method: 'POST', data })
}

export function updateRoleApi(data: Partial<RoleRecord>) {
  return request<boolean>({ url: '/system/roles', method: 'PUT', data })
}

export function deleteRoleApi(id: number) {
  return request<boolean>({ url: `/system/roles?id=${id}`, method: 'DELETE' })
}

export function getMenusApi() {
  return request<MenuRecord[]>({ url: '/system/menus', method: 'GET' })
}

export function createMenuApi(data: Partial<MenuRecord>) {
  return request<{ id: number }>({ url: '/system/menus', method: 'POST', data })
}

export function updateMenuApi(data: Partial<MenuRecord>) {
  return request<boolean>({ url: '/system/menus', method: 'PUT', data })
}

export function deleteMenuApi(id: number) {
  return request<boolean>({ url: `/system/menus?id=${id}`, method: 'DELETE' })
}

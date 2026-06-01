export interface TenantRecord {
  id: number
  name: string
  code: string
  status: number
  remark: string
  createdAt: string
}

export interface UserRecord {
  id: number
  username: string
  nickname: string
  tenantId: number
  roleIds: number[]
  status: number
  createdAt: string
}

export interface RoleRecord {
  id: number
  name: string
  code: string
  tenantId: number
  menuIds: number[]
  status: number
  createdAt: string
}

export interface MenuRecord {
  id: number
  parentId: number
  title: string
  path: string
  component: string
  permission: string
  type: 'catalog' | 'menu' | 'button'
  sort: number
  visible: boolean
}

export interface LoginParams {
  username: string
  password: string
  tenantCode?: string
}

export interface LoginResult {
  token: string
  refreshToken: string
}

export interface TenantInfo {
  id: number
  name: string
  code: string
}

export interface UserInfo {
  id: number
  username: string
  nickname: string
  tenant: TenantInfo
  roles: string[]
  permissions: string[]
}

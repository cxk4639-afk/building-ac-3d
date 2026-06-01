import { request } from './http'
import type { LoginParams, LoginResult, UserInfo } from '@/types/auth'

export function loginApi(data: LoginParams) {
  return request<LoginResult>({
    url: '/auth/login',
    method: 'POST',
    data,
  })
}

export function getUserInfoApi() {
  return request<UserInfo>({
    url: '/auth/me',
    method: 'GET',
  })
}

export function logoutApi() {
  return request<null>({
    url: '/auth/logout',
    method: 'POST',
  })
}

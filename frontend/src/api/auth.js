/**
 * @fileoverview 管理员认证 API。
 */
import request from './request'

/**
 * 管理员登录
 * @param {{ username: string, password: string }} credentials
 * @returns {Promise<{ access_token: string, user_type: string }>}
 */
export const adminLogin = (credentials) => request.post('/api/auth/loginAdmin', credentials)

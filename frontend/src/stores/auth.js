/**
 * @fileoverview Pinia 管理员认证状态管理。
 */
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { adminLogin } from '../api/auth'
import router from '../router'
import { ElMessage } from 'element-plus'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || '')
  const username = ref(localStorage.getItem('username') || '')

  // 清理旧版本遗留的用户身份标记
  localStorage.removeItem('userType')

  const isLoggedIn = computed(() => !!token.value)

  function _persist() {
    localStorage.setItem('token', token.value)
    localStorage.setItem('username', username.value)
  }

  /**
   * 管理员登录
   * @param {{ username: string, password: string }} credentials
   */
  async function loginAdmin(credentials) {
    const { data } = await adminLogin(credentials)
    token.value = data.access_token
    username.value = credentials.username
    _persist()
    ElMessage.success('登录成功')
    router.push('/admin/nfc')
  }

  function logout() {
    token.value = ''
    username.value = ''
    localStorage.removeItem('token')
    localStorage.removeItem('username')
    localStorage.removeItem('userType')
    router.push('/login')
    ElMessage.success('已退出登录')
  }

  return { token, username, isLoggedIn, loginAdmin, logout }
})

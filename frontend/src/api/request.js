/**
 * @fileoverview Axios 实例封装，统一配置拦截器、错误处理和 Token 注入。
 */
import axios from 'axios'
import { ElMessage } from 'element-plus'
import router from '../router'

const request = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '',
  timeout: 15000
})

request.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

/** @type {Record<number, string>} HTTP 状态码对应的默认错误消息 */
const STATUS_MESSAGES = {
  400: '请求参数错误，请检查输入',
  401: '登录已过期，请重新登录',
  403: '没有权限访问该资源',
  404: '请求的资源不存在',
  422: '数据验证失败，请检查输入格式',
  500: '服务器内部错误，请稍后重试',
  502: '网关错误，请稍后重试',
  503: '服务暂时不可用，请稍后重试'
}

/** @type {Record<string, string>} 后端 detail 特殊值对应的中文提示 */
const DETAIL_MESSAGES = {
  'nfc_id exists': '该 NFC ID 已存在，请使用其他 ID'
}

request.interceptors.response.use(
  response => response,
  error => {
    // 调用方声明自行处理错误时（如公开落地页的停用兜底），不弹全局提示
    if (error.config?.skipErrorToast) return Promise.reject(error)

    if (error.response) {
      const { status, data } = error.response
      if (status === 401) {
        localStorage.removeItem('token')
        ElMessage.error(STATUS_MESSAGES[401])
        router.push('/login')
        return Promise.reject(error)
      }
      const detail = data?.detail
      const message = DETAIL_MESSAGES[detail] || detail || STATUS_MESSAGES[status] || '请求失败，请重试'
      ElMessage.error(message)
    } else if (error.request) {
      ElMessage.error('网络错误，请检查网络连接')
    } else {
      ElMessage.error('请求配置错误')
    }
    return Promise.reject(error)
  }
)

export default request

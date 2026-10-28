/**
 * @fileoverview 用户意见 API：公开提交、已采纳意见、管理员列表与采纳。
 */
import request from './request'

/**
 * 提交意见（无需登录）
 * @param {{ order_no?: string, content: string }} data
 * @returns {Promise}
 */
export const createFeedback = (data) => request.post('/api/feedback', data)

/**
 * 已采纳意见（公开，不返回订单号）
 * @param {{ limit?: number }} [params]
 * @returns {Promise}
 */
export const fetchAdoptedFeedback = (params) => request.get('/api/feedback/adopted', { params })

/**
 * 意见列表（管理员）
 * @param {{ page?: number, page_size?: number, search?: string, status?: string }} [params]
 * @returns {Promise}
 */
export const fetchFeedbackList = (params) => request.get('/api/feedback/admin', { params })

/**
 * 采纳 / 取消采纳（管理员）
 * @param {number|string} id
 * @param {boolean} adopted
 * @returns {Promise}
 */
export const adoptFeedback = (id, adopted = true) =>
  request.patch(`/api/feedback/admin/${id}`, { adopted })

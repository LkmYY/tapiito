/**
 * @fileoverview 硬件订单 API：展厅公开下单 + 管理后台订单管理。
 */
import request from './request'

/**
 * 提交订单（无需登录，展厅使用）
 * @param {{
 *   template_type: string,
 *   custom_template?: string,
 *   template_payload?: object,
 *   tag_name?: string,
 *   contact_name: string,
 *   contact_phone: string,
 *   contact_address: string,
 *   contact_zipcode?: string,
 *   contact_note?: string
 * }} data
 * @returns {Promise}
 */
export const createOrder = (data) => request.post('/api/orders', data)

/**
 * 订单列表（管理员）
 * @param {{ page?: number, page_size?: number, search?: string, status?: string }} [params]
 * @returns {Promise}
 */
export const fetchOrderList = (params) => request.get('/api/orders', { params })

/**
 * 订单详情（管理员）
 * @param {number|string} id
 * @returns {Promise}
 */
export const fetchOrderDetail = (id) => request.get(`/api/orders/${id}`)

/**
 * 订单转换成标签（管理员）
 * @param {number|string} id
 * @param {{ nfc_id?: string, name?: string, phone?: string }} [data]
 * @returns {Promise}
 */
export const convertOrder = (id, data) => request.post(`/api/orders/${id}/convert`, data || {})

/**
 * 删除订单（管理员）
 * @param {number|string} id
 * @returns {Promise}
 */
export const deleteOrder = (id) => request.delete(`/api/orders/${id}`)

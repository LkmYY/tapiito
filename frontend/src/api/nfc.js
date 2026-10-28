/**
 * @fileoverview NFC 标签 CRUD API（管理员）。
 */
import request from './request'

/**
 * 获取 NFC 标签列表（支持服务端分页）
 * @param {{ page?: number, page_size?: number, search?: string, status?: string, creator_id?: number }} [params]
 * @returns {Promise}
 */
export const fetchNfcList = (params) => request.get('/api/nfc', { params })

/**
 * 获取创建人筛选选项
 * @returns {Promise}
 */
export const fetchCreators = () => request.get('/api/nfc/creators')

/**
 * 获取单个 NFC 标签详情
 * @param {number|string} id
 * @returns {Promise}
 */
export const fetchNfcDetail = (id) => request.get(`/api/nfc/${id}`)

/**
 * 创建 NFC 标签
 * @param {object} data
 * @returns {Promise}
 */
export const createNfc = (data) => request.post('/api/nfc', data)

/**
 * 更新 NFC 标签
 * @param {number|string} id
 * @param {object} data
 * @returns {Promise}
 */
export const updateNfc = (id, data) => request.put(`/api/nfc/${id}`, data)

/**
 * 删除 NFC 标签
 * @param {number|string} id
 * @returns {Promise}
 */
export const deleteNfc = (id) => request.delete(`/api/nfc/${id}`)

/**
 * 批量操作 NFC 标签
 * @param {{ ids: number[], action: 'enable'|'disable'|'delete' }} data
 * @returns {Promise}
 */
export const batchNfc = (data) => request.post('/api/nfc/batch', data)

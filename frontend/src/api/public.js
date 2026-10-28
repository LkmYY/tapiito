/**
 * @fileoverview 公开页面 API：NFC 落地页、照片、备注。
 */
import request from './request'

/**
 * 获取 NFC 标签页面数据
 * 页面需自行处理「标签已停用 / 不存在」的兜底展示，因此关闭全局错误提示
 * @param {string} nfcId
 * @returns {Promise}
 */
export const fetchNfcPage = (nfcId, access = {}) =>
  request.get(`/api/public/nfc/${nfcId}/page`, {
    skipErrorToast: true,
    params: access.ticket ? { ticket: access.ticket } : undefined,
    headers: access.session ? { 'X-Access-Session': access.session } : undefined,
  })

/**
 * 提交暗号换取一次性 ticket
 * @param {string} nfcId
 * @param {string} passcode
 */
export const unlockNfc = (nfcId, passcode) =>
  request.post(`/api/public/nfc/${nfcId}/unlock`, { passcode }, { skipErrorToast: true })

/**
 * 获取首页入驻展厅数据（已启用标签概览）
 * @param {{ limit?: number }} [params]
 * @returns {Promise}
 */
export const fetchGallery = (params) => request.get('/api/public/gallery', { params })

/**
 * 获取标签照片列表
 * @param {string} nfcId
 * @returns {Promise}
 */
export const fetchPhotos = (nfcId) =>
  request.get(`/api/public/nfc/${nfcId}/photos`)

/**
 * 上传照片
 * @param {string} nfcId
 * @param {FormData} formData
 * @returns {Promise}
 */
export const uploadPhoto = (nfcId, formData) =>
  request.post(`/api/public/nfc/${nfcId}/photo`, formData)

/**
 * 提交备注和短信通知
 * @param {string} nfcId
 * @param {{ content: string, photo_ids?: number[] }} data
 * @returns {Promise}
 */
export const submitNote = (nfcId, data) =>
  request.post(`/api/public/nfc/${nfcId}/submit`, data)

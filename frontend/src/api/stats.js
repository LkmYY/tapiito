/**
 * @fileoverview 访问统计 API。
 */
import request from './request'

/**
 * 获取统计概览
 * @param {{ start_date?: string, end_date?: string }} [params]
 * @returns {Promise}
 */
export const fetchStatsOverview = (params) =>
  request.get('/api/stats/overview', { params })

/**
 * 获取单个标签的统计详情
 * @param {string} nfcId
 * @param {{ start_date?: string, end_date?: string }} [params]
 * @returns {Promise}
 */
export const fetchTagStats = (nfcId, params) =>
  request.get(`/api/stats/nfc/${nfcId}`, { params })

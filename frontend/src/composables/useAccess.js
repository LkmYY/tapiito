/**
 * 访问凭证：浏览器会话 ID + 一次性 ticket
 *
 * 会话 ID 存在 localStorage（同一浏览器稳定），ticket 按标签分别保存。
 * 后端把 ticket 绑定到首次使用的会话 —— 同会话可刷新，换浏览器/转发即失效。
 */
const SESSION_KEY = 'tapiito:access-session'
const ticketKey = (nfcId) => `tapiito:ticket:${nfcId}`

/** 取（或创建）本浏览器的会话标识 */
export function getSessionId() {
  try {
    let id = localStorage.getItem(SESSION_KEY)
    if (!id) {
      id = (crypto.randomUUID && crypto.randomUUID()) ||
        `s-${Date.now()}-${Math.random().toString(36).slice(2, 10)}`
      localStorage.setItem(SESSION_KEY, id)
    }
    return id
  } catch {
    return ''
  }
}

/** 取该标签当前有效的 ticket */
export function getTicket(nfcId) {
  try {
    const raw = localStorage.getItem(ticketKey(nfcId))
    if (!raw) return ''
    const { ticket, expiresAt } = JSON.parse(raw)
    if (!ticket || !expiresAt) return ''
    // 过期就丢掉，避免带着废票反复 403
    if (new Date(expiresAt).getTime() < Date.now()) {
      localStorage.removeItem(ticketKey(nfcId))
      return ''
    }
    return ticket
  } catch {
    return ''
  }
}

export function saveTicket(nfcId, ticket, expiresAt) {
  if (!ticket) return
  try {
    localStorage.setItem(ticketKey(nfcId), JSON.stringify({ ticket, expiresAt }))
  } catch {
    // 隐私模式下写不进去，忽略（下次仍需重新输入暗号）
  }
}

export function clearTicket(nfcId) {
  try {
    localStorage.removeItem(ticketKey(nfcId))
  } catch {
    // ignore
  }
}

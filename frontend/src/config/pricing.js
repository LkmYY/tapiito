/**
 * @fileoverview 标签硬件定价（暂定规则，支付未接入，仅用于下单页展示与订单记录）
 *
 * 规则：
 *   1 个   9.9 元
 *   5 个   40 元
 *   10 个  69.9 元
 *   50 个以上   6 元 / 个
 *   100 个以上  5.5 元 / 个
 *   自定义数量：按上面阶梯自动折算金额，也可手动改金额
 */

export const CUSTOM_MODE = 'custom'

/** 固定数量选项（amount 为该档总价） */
export const QUANTITY_PRESETS = [
  { qty: 1, amount: 9.9, label: '1 个', hint: '先试试' },
  { qty: 5, amount: 40, label: '5 个', hint: '¥8 / 个' },
  { qty: 10, amount: 69.9, label: '10 个', hint: '¥6.99 / 个' },
  { qty: 50, amount: 300, label: '50 个', hint: '¥6 / 个' },
  { qty: 100, amount: 550, label: '100 个', hint: '¥5.5 / 个' },
]

const round2 = (v) => Math.round(Number(v || 0) * 100) / 100

/**
 * 按数量推算金额：100 以上 5.5 元/个，50 以上 6 元/个，
 * 其余按「不超过该数量的最大套餐」单价折算。
 * @param {number} qty
 * @returns {number}
 */
export function amountForQuantity(qty) {
  const n = Math.max(1, Math.floor(Number(qty) || 1))
  if (n >= 100) return round2(n * 5.5)
  if (n >= 50) return round2(n * 6)
  const base = [...QUANTITY_PRESETS].reverse().find((p) => p.qty <= n) || QUANTITY_PRESETS[0]
  return round2((base.amount / base.qty) * n)
}

/** 金额展示：9.9 → ¥9.9，40 → ¥40 */
export const money = (v) => `¥${Number(round2(v))}`

/** 单价展示：保留一位小数以内 */
export const unitMoney = (amount, qty) => {
  const n = Math.max(1, Number(qty) || 1)
  return `¥${Number((Number(amount || 0) / n).toFixed(2))}`
}

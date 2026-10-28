/**
 * 微信号点击增强：复制微信号 + 尝试唤起微信
 *
 * 微信号散落在 30+ 个模板里（36 处），逐个改模板成本高且易漏。
 * 这里在展示容器上做一次 DOM 扫描 + 事件委托：
 *   1. 渲染后找出文本恰好等于微信号的节点，给其父元素打 data-wechat 标记
 *   2. 容器统一监听 click，命中标记就复制并尝试唤起微信
 *
 * 这样任何模板（现有或以后新增）只要把 payload.wechat 渲染出来就自动生效。
 */
import { nextTick, onBeforeUnmount, watch } from 'vue'

const WECHAT_SCHEME = 'weixin://'

/** 复制文本：优先 Clipboard API（需 HTTPS），失败回退 execCommand */
async function copyText(text) {
  try {
    if (navigator.clipboard && window.isSecureContext) {
      await navigator.clipboard.writeText(text)
      return true
    }
  } catch {
    // 继续走兜底
  }
  try {
    const ta = document.createElement('textarea')
    ta.value = text
    ta.setAttribute('readonly', '')
    ta.style.cssText = 'position:fixed;left:-9999px;top:0;opacity:0'
    document.body.appendChild(ta)
    ta.select()
    const ok = document.execCommand('copy')
    ta.remove()
    return ok
  } catch {
    return false
  }
}

/** 是否在微信内置浏览器里（在微信里就不需要再唤起） */
function isInWechat() {
  return /micromessenger/i.test(navigator.userAgent)
}

/** 用隐藏 iframe 触发 scheme，失败时页面不受影响 */
function openWechat() {
  if (isInWechat()) return false
  try {
    const iframe = document.createElement('iframe')
    iframe.style.cssText = 'display:none;width:0;height:0;border:0'
    iframe.src = WECHAT_SCHEME
    document.body.appendChild(iframe)
    setTimeout(() => iframe.remove(), 1500)
    return true
  } catch {
    return false
  }
}

/** 给 root 内所有「文本等于微信号」的元素打标记 */
function markWechatNodes(root, wechat) {
  if (!root || !wechat) return
  const target = String(wechat).trim()
  if (!target) return

  const walker = document.createTreeWalker(root, NodeFilter.SHOW_TEXT)
  const hits = []
  let node = walker.nextNode()
  while (node) {
    if (node.nodeValue && node.nodeValue.trim() === target) {
      const parent = node.parentElement
      // 已经标过、或本身是按钮/链接的跳过
      if (parent && !parent.dataset.wechat && !parent.closest('a,button')) {
        hits.push(parent)
      }
    }
    node = walker.nextNode()
  }
  hits.forEach((el) => {
    el.dataset.wechat = target
    el.classList.add('wechat-tappable')
  })
}

/**
 * @param {import('vue').Ref<HTMLElement|null>} rootRef 展示容器
 * @param {() => string} getWechat 取当前微信号
 * @param {(msg: string) => void} notify 提示回调
 */
export function useWechatTap(rootRef, getWechat, notify) {
  let cleanup = null

  const onClick = async (e) => {
    const el = e.target?.closest?.('[data-wechat]')
    if (!el) return
    const value = el.dataset.wechat
    e.preventDefault()
    e.stopPropagation()

    const copied = await copyText(value)
    if (copied) {
      notify(`微信号 ${value} 已复制`)
    } else {
      notify(`微信号：${value}（复制失败，请手动记录）`)
    }
    // 复制完成后再唤起，避免部分机型跳走后复制失败
    openWechat()
  }

  const scan = () => {
    const root = rootRef.value
    if (!root) return
    // 先清掉上一轮标记，避免内容切换后残留
    root.querySelectorAll('[data-wechat]').forEach((el) => {
      delete el.dataset.wechat
      el.classList.remove('wechat-tappable')
    })
    markWechatNodes(root, getWechat())
  }

  const bind = () => {
    const root = rootRef.value
    if (!root) return
    root.addEventListener('click', onClick)
    cleanup = () => root.removeEventListener('click', onClick)
  }

  watch(
    () => getWechat(),
    () => nextTick(() => { scan(); if (!cleanup) bind() }),
    { immediate: true },
  )

  nextTick(() => { scan(); bind() })

  onBeforeUnmount(() => {
    cleanup?.()
  })

  return { rescan: scan }
}

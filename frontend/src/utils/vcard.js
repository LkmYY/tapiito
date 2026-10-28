/**
 * 联系人保存：优先走系统分享面板（可直接「添加到通讯录」），否则回退下载 .vcf
 *
 * 背景：浏览器出于安全限制，网页脚本无法直接写入手机通讯录。
 * 能做到的最优体验是：生成 vCard → 走 Web Share API 交给系统处理，
 * 系统面板里选「通讯录」即可直接入库，不用先落一个文件再手动导入。
 */
export function buildVCard(p = {}) {
  const esc = (v) =>
    String(v || '')
      .replace(/\\/g, '\\\\')
      .replace(/\n/g, '\\n')
      .replace(/;/g, '\\;')
      .replace(/,/g, '\\,')

  const lines = ['BEGIN:VCARD', 'VERSION:3.0', `FN:${esc(p.name)}`]
  if (p.company) lines.push(`ORG:${esc(p.company)}`)
  if (p.title) lines.push(`TITLE:${esc(p.title)}`)
  if (p.phone) lines.push(`TEL;TYPE=CELL:${esc(p.phone)}`)
  if (p.email) lines.push(`EMAIL;TYPE=WORK:${esc(p.email)}`)
  if (p.website) lines.push(`URL:${esc(p.website)}`)
  if (p.wechat) lines.push(`NOTE:微信 ${esc(p.wechat)}`)
  lines.push('END:VCARD')
  return lines.join('\r\n') + '\r\n'
}

const safeFileName = (name) => `${String(name || 'contact').replace(/[\\/:*?"<>|]+/g, '_')}.vcf`

/**
 * @param {object} payload 名片数据
 * @returns {Promise<{ok: boolean, via?: 'share'|'download', reason?: 'wechat'|'cancel'}>}
 */
export async function saveContact(payload = {}) {
  const text = buildVCard(payload)
  const fileName = safeFileName(payload.name)

  // 微信内置浏览器：既不能下载也不能 share，只能引导用户去外部浏览器
  if (/micromessenger/i.test(navigator.userAgent)) {
    return { ok: false, reason: 'wechat' }
  }

  // Web Share API：系统面板里可直接选「添加到通讯录」，这是最接近"直接保存"的路径
  try {
    const file = new File([text], fileName, { type: 'text/vcard' })
    if (navigator.canShare && navigator.canShare({ files: [file] })) {
      await navigator.share({ files: [file], title: payload.name || '联系人' })
      return { ok: true, via: 'share' }
    }
  } catch (err) {
    if (err && err.name === 'AbortError') {
      return { ok: false, reason: 'cancel' }
    }
    // 其它异常继续走下载兜底
  }

  // 兜底：下载 .vcf（iOS/Android 打开该文件都会弹出「添加联系人」）
  const blob = new Blob([text], { type: 'text/vcard;charset=utf-8' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = fileName
  document.body.appendChild(a)
  a.click()
  a.remove()
  setTimeout(() => URL.revokeObjectURL(url), 3000)
  return { ok: true, via: 'download' }
}

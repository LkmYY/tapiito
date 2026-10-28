/**
 * 「奇怪的想法」绘制器集合
 *
 * 每个绘制器签名：(ctx, w, h, t, p) => void
 *   - w/h: 逻辑像素尺寸（已按 DPR 缩放，直接用 CSS 像素坐标）
 *   - t:   秒（随 speed 缩放由调用方保证）
 *   - p:   { title, subtitle, speed, intensity, accent, pointer:{x,y}, nfcId }
 *
 * 约定：绘制器必须自己清屏之外的一切（舞台已 clearRect），并自行 save/restore。
 */

/* ---------------- 通用工具 ---------------- */

const TAU = Math.PI * 2

function accentOf(p, fallback) {
  return p.accent || fallback
}

/** 确定性伪随机（同一 seed 每帧稳定，避免噪点乱跳） */
function seeded(seed) {
  let s = seed % 2147483647
  if (s <= 0) s += 2147483646
  return () => {
    s = (s * 16807) % 2147483647
    return (s - 1) / 2147483646
  }
}

function fitFont(ctx, text, maxWidth, baseSize, family) {
  let size = baseSize
  ctx.font = `700 ${size}px ${family}`
  while (ctx.measureText(text).width > maxWidth && size > 10) {
    size -= 2
    ctx.font = `700 ${size}px ${family}`
  }
  return size
}

const SANS = '"Helvetica Neue", Arial, "PingFang SC", "Microsoft YaHei", sans-serif'
const MONO = 'ui-monospace, SFMono-Regular, Menlo, monospace'

/* ---------------- A. 眼睛与凝视 ---------------- */

/** 1. 血轮眼：红色眼白 + 虹膜 + 三勾玉 + 血丝 */
function sharingan(ctx, w, h, t, p) {
  const cx = w / 2
  const cy = h / 2
  const base = Math.min(w, h) * 0.34
  const speed = 0.4 + p.speed * 0.22
  const accent = accentOf(p, '#e11d2f')

  // 背景暗红
  const bg = ctx.createRadialGradient(cx, cy, base * 0.2, cx, cy, Math.max(w, h) * 0.75)
  bg.addColorStop(0, '#1a0407')
  bg.addColorStop(1, '#040102')
  ctx.fillStyle = bg
  ctx.fillRect(0, 0, w, h)

  // 眼白
  ctx.beginPath()
  ctx.ellipse(cx, cy, base * 1.5, base, 0, 0, TAU)
  ctx.fillStyle = '#f4e6e6'
  ctx.fill()

  // 血丝
  const rnd = seeded(7)
  ctx.strokeStyle = 'rgba(190,20,40,0.55)'
  ctx.lineWidth = Math.max(1, base * 0.012)
  for (let i = 0; i < 26; i++) {
    const a = rnd() * TAU
    const r0 = base * (0.55 + rnd() * 0.3)
    const r1 = base * (0.95 + rnd() * 0.5)
    ctx.beginPath()
    ctx.moveTo(cx + Math.cos(a) * r0, cy + Math.sin(a) * r0 * 0.68)
    ctx.quadraticCurveTo(
      cx + Math.cos(a + 0.12) * r1 * 0.9, cy + Math.sin(a + 0.12) * r1 * 0.6,
      cx + Math.cos(a + 0.2) * r1, cy + Math.sin(a + 0.2) * r1 * 0.7,
    )
    ctx.stroke()
  }

  // 虹膜
  const irisR = base * 0.72
  const iris = ctx.createRadialGradient(cx, cy, irisR * 0.2, cx, cy, irisR)
  iris.addColorStop(0, '#5b0009')
  iris.addColorStop(0.55, accent)
  iris.addColorStop(1, '#7a0010')
  ctx.beginPath()
  ctx.arc(cx, cy, irisR, 0, TAU)
  ctx.fillStyle = iris
  ctx.fill()

  // 勾玉（跟随指针轻微偏移）
  const px = (p.pointer.x - 0.5) * irisR * 0.22
  const py = (p.pointer.y - 0.5) * irisR * 0.22
  const rot = t * speed
  for (let i = 0; i < 3; i++) {
    const a = rot + (i * TAU) / 3
    const r = irisR * 0.62
    ctx.save()
    ctx.translate(cx + px + Math.cos(a) * r, cy + py + Math.sin(a) * r)
    ctx.rotate(a + Math.PI / 2)
    ctx.beginPath()
    ctx.arc(0, 0, irisR * 0.2, 0, TAU)
    ctx.fillStyle = '#0b0002'
    ctx.fill()
    // 勾玉尾巴
    ctx.beginPath()
    ctx.moveTo(irisR * 0.2, 0)
    ctx.quadraticCurveTo(irisR * 0.42, irisR * 0.1, irisR * 0.3, irisR * 0.34)
    ctx.quadraticCurveTo(irisR * 0.12, irisR * 0.2, irisR * 0.2, 0)
    ctx.fill()
    ctx.restore()
  }

  // 瞳孔收缩：强度越高越小，并随呼吸脉动
  const pulse = 1 + Math.sin(t * (1.2 + p.intensity * 0.4)) * 0.06
  ctx.beginPath()
  ctx.arc(cx + px * 1.4, cy + py * 1.4, irisR * (0.3 - p.intensity * 0.022) * pulse, 0, TAU)
  ctx.fillStyle = '#000'
  ctx.fill()

  if (p.title) drawCaption(ctx, w, h, p, '#ff6b6b')
}

/** 2. 万花筒之眼：六重对称几何瞳 */
function kaleidoEye(ctx, w, h, t, p) {
  const cx = w / 2
  const cy = h / 2
  const R = Math.min(w, h) * 0.4
  const accent = accentOf(p, '#a855f7')
  const speed = 0.25 + p.speed * 0.16

  ctx.fillStyle = '#08040f'
  ctx.fillRect(0, 0, w, h)

  for (let ring = 5; ring >= 1; ring--) {
    const rr = (R * ring) / 5
    const rot = t * speed * (ring % 2 ? 1 : -1)
    ctx.save()
    ctx.translate(cx, cy)
    ctx.rotate(rot)
    for (let i = 0; i < 6; i++) {
      ctx.save()
      ctx.rotate((i * TAU) / 6)
      ctx.beginPath()
      ctx.moveTo(0, 0)
      ctx.arc(0, 0, rr, -0.42, 0.42)
      ctx.closePath()
      ctx.fillStyle = `rgba(${168 - ring * 12}, ${85 + ring * 8}, ${247 - ring * 20}, ${0.1 + ring * 0.06})`
      ctx.fill()
      ctx.strokeStyle = accent
      ctx.globalAlpha = 0.18 + ring * 0.06
      ctx.lineWidth = 1.2
      ctx.stroke()
      ctx.restore()
    }
    ctx.restore()
  }

  // 中央瞳孔
  ctx.globalAlpha = 1
  ctx.beginPath()
  ctx.arc(cx, cy, R * 0.16 + Math.sin(t * 2) * R * 0.012, 0, TAU)
  ctx.fillStyle = '#000'
  ctx.fill()
  ctx.strokeStyle = accent
  ctx.lineWidth = 2
  ctx.stroke()

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 3. 无尽瞳孔：瞳孔嵌套，不断向内缩放 */
function infinitePupil(ctx, w, h, t, p) {
  const cx = w / 2
  const cy = h / 2
  const accent = accentOf(p, '#38bdf8')
  const R = Math.min(w, h) * 0.42

  ctx.fillStyle = '#000306'
  ctx.fillRect(0, 0, w, h)

  const phase = (t * (0.15 + p.speed * 0.08)) % 1
  for (let i = 0; i < 9; i++) {
    // 每层按相位向内缩，形成无限坠落
    const k = ((i / 9 + phase) % 1)
    const scale = 1 - k
    const r = R * scale
    if (r < 2) continue
    ctx.beginPath()
    ctx.arc(cx, cy, r, 0, TAU)
    ctx.strokeStyle = accent
    ctx.globalAlpha = 0.12 + scale * 0.5
    ctx.lineWidth = Math.max(1, 3 * scale)
    ctx.stroke()

    // 瞳孔黑心
    ctx.beginPath()
    ctx.arc(cx, cy, r * 0.42, 0, TAU)
    ctx.fillStyle = `rgba(0,0,0,${0.35 + k * 0.5})`
    ctx.fill()
  }
  ctx.globalAlpha = 1
  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 4. 凝视者：一只眼跟着指针转 */
function watcher(ctx, w, h, t, p) {
  const cx = w / 2
  const cy = h / 2
  const R = Math.min(w, h) * 0.36
  const accent = accentOf(p, '#fbbf24')

  ctx.fillStyle = '#0b0a08'
  ctx.fillRect(0, 0, w, h)

  // 眼睑轮廓
  ctx.beginPath()
  ctx.moveTo(cx - R * 1.6, cy)
  ctx.quadraticCurveTo(cx, cy - R * 1.25, cx + R * 1.6, cy)
  ctx.quadraticCurveTo(cx, cy + R * 1.25, cx - R * 1.6, cy)
  ctx.closePath()
  ctx.fillStyle = '#f2efe6'
  ctx.fill()
  ctx.strokeStyle = 'rgba(30,26,20,0.35)'
  ctx.lineWidth = 3
  ctx.stroke()

  // 眼珠朝向指针
  const dx = (p.pointer.x - 0.5) * 2
  const dy = (p.pointer.y - 0.5) * 2
  const maxOff = R * 0.42
  const ex = cx + dx * maxOff
  const ey = cy + dy * maxOff * 0.6

  ctx.save()
  ctx.beginPath()
  ctx.moveTo(cx - R * 1.6, cy)
  ctx.quadraticCurveTo(cx, cy - R * 1.25, cx + R * 1.6, cy)
  ctx.quadraticCurveTo(cx, cy + R * 1.25, cx - R * 1.6, cy)
  ctx.clip()

  const iris = ctx.createRadialGradient(ex, ey, R * 0.1, ex, ey, R * 0.55)
  iris.addColorStop(0, '#2b1a06')
  iris.addColorStop(0.6, accent)
  iris.addColorStop(1, '#5a3a0a')
  ctx.beginPath()
  ctx.arc(ex, ey, R * 0.55, 0, TAU)
  ctx.fillStyle = iris
  ctx.fill()

  ctx.beginPath()
  ctx.arc(ex, ey, R * 0.24, 0, TAU)
  ctx.fillStyle = '#000'
  ctx.fill()

  // 高光
  ctx.beginPath()
  ctx.arc(ex - R * 0.14, ey - R * 0.16, R * 0.1, 0, TAU)
  ctx.fillStyle = 'rgba(255,255,255,0.85)'
  ctx.fill()
  ctx.restore()

  // 眨眼：每 4 秒快速闭一次
  const blink = (t % 4.2) > 4.0
  if (blink) {
    ctx.beginPath()
    ctx.moveTo(cx - R * 1.6, cy)
    ctx.quadraticCurveTo(cx, cy + R * 0.18, cx + R * 1.6, cy)
    ctx.quadraticCurveTo(cx, cy - R * 0.18, cx - R * 1.6, cy)
    ctx.fillStyle = '#0b0a08'
    ctx.fill()
  }

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 5. 复眼阵列：网格小眼，随机眨动 */
function compoundEye(ctx, w, h, t, p) {
  const accent = accentOf(p, '#22d3ee')
  const cols = 12
  const rows = 7
  const cw = w / cols
  const ch = h / rows
  const r = Math.min(cw, ch) * 0.32
  const rnd = seeded(Math.floor(t / 1.6))

  ctx.fillStyle = '#03080c'
  ctx.fillRect(0, 0, w, h)

  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      const cx = cw * (col + 0.5)
      const cy = ch * (row + 0.5)
      const blinkSeed = rnd()
      const closed = blinkSeed > 0.82

      ctx.beginPath()
      ctx.arc(cx, cy, r, 0, TAU)
      ctx.fillStyle = closed ? '#0d1a22' : '#eaf6f8'
      ctx.fill()

      if (!closed) {
        const look = Math.sin(t * 0.8 + col * 0.4 + row * 0.3) * r * 0.28
        ctx.beginPath()
        ctx.arc(cx + look, cy + look * 0.5, r * 0.46, 0, TAU)
        ctx.fillStyle = accent
        ctx.fill()
        ctx.beginPath()
        ctx.arc(cx + look, cy + look * 0.5, r * 0.2, 0, TAU)
        ctx.fillStyle = '#04141a'
        ctx.fill()
      }
      ctx.beginPath()
      ctx.arc(cx, cy, r, 0, TAU)
      ctx.strokeStyle = 'rgba(34,211,238,0.25)'
      ctx.lineWidth = 1
      ctx.stroke()
    }
  }
  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/* ---------------- B. 时间 ---------------- */

/** 6. 巨型翻页钟 */
function flipClock(ctx, w, h, t, p) {
  const accent = accentOf(p, '#e5e7eb')
  ctx.fillStyle = '#0a0a0a'
  ctx.fillRect(0, 0, w, h)

  const now = new Date()
  const hh = String(now.getHours()).padStart(2, '0')
  const mm = String(now.getMinutes()).padStart(2, '0')
  const ss = String(now.getSeconds()).padStart(2, '0')

  const gap = w * 0.012
  // 6 位数字必须落在 86% 宽度内，否则两侧会被裁掉
  const cellW = (w * 0.86 - gap * 5) / 6
  const cellH = cellW * 0.72
  const startX = (w * 0.86 - (cellW * 6 + gap * 5)) / 2 + w * 0.07
  const cy = h * 0.48

  const digits = (hh + mm + ss).split('')
  // 秒数跳动时的轻微抖动
  const shake = Math.sin(t * 12) * (Math.floor(t * 1000) % 1000 < 80 ? 2 : 0)

  digits.forEach((d, i) => {
    const sep = i === 2 || i === 4
    const x = startX + i * (cellW + gap)
    ctx.save()
    ctx.translate(x + cellW / 2, cy + shake)
    ctx.beginPath()
    ctx.roundRect(-cellW / 2, -cellH / 2, cellW, cellH, cellW * 0.08)
    ctx.fillStyle = '#171717'
    ctx.fill()
    ctx.strokeStyle = 'rgba(255,255,255,0.08)'
    ctx.lineWidth = 1
    ctx.stroke()
    // 中缝
    ctx.beginPath()
    ctx.moveTo(-cellW / 2, 0)
    ctx.lineTo(cellW / 2, 0)
    ctx.strokeStyle = 'rgba(0,0,0,0.6)'
    ctx.stroke()

    ctx.fillStyle = sep ? accent : '#fafafa'
    ctx.font = `700 ${cellH * 0.62}px ${MONO}`
    ctx.textAlign = 'center'
    ctx.textBaseline = 'middle'
    ctx.fillText(d, 0, cellH * 0.02)
    ctx.restore()
  })

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 7. 倒计时炸弹 */
function countdownBomb(ctx, w, h, t, p) {
  const accent = accentOf(p, '#ef4444')
  const period = 20
  const remain = period - (t % period)
  const sec = Math.ceil(remain)
  const danger = 1 - remain / period

  // 越接近 0 背景越红、越抖
  const bg = ctx.createRadialGradient(w / 2, h / 2, 10, w / 2, h / 2, Math.max(w, h) * 0.7)
  bg.addColorStop(0, `rgba(60,0,0,${0.55 + danger * 0.4})`)
  bg.addColorStop(1, '#000')
  ctx.fillStyle = bg
  ctx.fillRect(0, 0, w, h)

  const shakeAmp = Math.pow(danger, 2.2) * 10 * (1 + p.intensity * 0.2)
  const ox = Math.sin(t * 47) * shakeAmp
  const oy = Math.cos(t * 39) * shakeAmp

  ctx.save()
  ctx.translate(w / 2 + ox, h / 2 + oy)
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.fillStyle = danger > 0.75 ? '#fff' : accent
  ctx.font = `700 ${Math.min(w, h) * 0.42}px ${MONO}`
  ctx.fillText(remain < 1 ? 'BOOM' : String(sec).padStart(2, '0'), 0, 0)
  ctx.restore()

  // 心跳脉冲
  const pulse = (Math.sin(t * (3 + danger * 9)) + 1) / 2
  ctx.beginPath()
  ctx.arc(w / 2, h / 2, Math.min(w, h) * (0.34 + pulse * 0.03), 0, TAU)
  ctx.strokeStyle = `rgba(239,68,68,${0.15 + danger * 0.5})`
  ctx.lineWidth = 2 + danger * 4
  ctx.stroke()

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 8. 失控指针钟：指针乱转、偶尔卡住 */
function brokenClock(ctx, w, h, t, p) {
  const cx = w / 2
  const cy = h / 2
  const R = Math.min(w, h) * 0.38
  const accent = accentOf(p, '#f8fafc')

  ctx.fillStyle = '#0d0d0f'
  ctx.fillRect(0, 0, w, h)

  ctx.beginPath()
  ctx.arc(cx, cy, R, 0, TAU)
  ctx.strokeStyle = 'rgba(255,255,255,0.14)'
  ctx.lineWidth = 2
  ctx.stroke()

  // 刻度
  for (let i = 0; i < 60; i++) {
    const a = (i / 60) * TAU - Math.PI / 2
    const long = i % 5 === 0
    const r1 = R * (long ? 0.86 : 0.92)
    ctx.beginPath()
    ctx.moveTo(cx + Math.cos(a) * r1, cy + Math.sin(a) * r1)
    ctx.lineTo(cx + Math.cos(a) * R * 0.97, cy + Math.sin(a) * R * 0.97)
    ctx.strokeStyle = long ? 'rgba(255,255,255,0.5)' : 'rgba(255,255,255,0.18)'
    ctx.lineWidth = long ? 2 : 1
    ctx.stroke()
  }

  // 每 2.4 秒卡顿一次
  const stuck = (t % 2.4) > 2.1
  const tt = stuck ? Math.floor(t / 2.4) * 2.4 : t

  const hands = [
    { len: 0.5, speed: 0.7 * p.speed, width: 6, color: '#fff' },
    { len: 0.72, speed: 2.3 * p.speed, width: 4, color: accent },
    { len: 0.86, speed: -5.1 * p.speed, width: 2, color: '#ef4444' },
  ]
  hands.forEach((hd) => {
    const a = tt * hd.speed - Math.PI / 2
    ctx.beginPath()
    ctx.moveTo(cx, cy)
    ctx.lineTo(cx + Math.cos(a) * R * hd.len, cy + Math.sin(a) * R * hd.len)
    ctx.strokeStyle = hd.color
    ctx.lineWidth = hd.width
    ctx.lineCap = 'round'
    ctx.stroke()
  })

  ctx.beginPath()
  ctx.arc(cx, cy, R * 0.04, 0, TAU)
  ctx.fillStyle = '#ef4444'
  ctx.fill()

  if (p.title) drawCaption(ctx, w, h, p, '#ef4444')
}

/** 9. 世界末日钟：永远差一秒到零点 */
function doomsdayClock(ctx, w, h, t, p) {
  const accent = accentOf(p, '#f97316')
  ctx.fillStyle = '#070607'
  ctx.fillRect(0, 0, w, h)

  // 缓慢逼近却永不抵达：指针在 23:59:59 附近来回
  const jitter = Math.sin(t * 1.1) * 0.6
  const secAngle = ((59 + jitter) / 60) * TAU - Math.PI / 2
  const cx = w / 2
  const cy = h * 0.52
  const R = Math.min(w, h) * 0.33

  ctx.beginPath()
  ctx.arc(cx, cy, R, 0, TAU)
  ctx.strokeStyle = 'rgba(249,115,22,0.28)'
  ctx.lineWidth = 2
  ctx.stroke()

  // 秒针无限接近 12 点
  ctx.beginPath()
  ctx.moveTo(cx, cy)
  ctx.lineTo(cx + Math.cos(secAngle) * R * 0.9, cy + Math.sin(secAngle) * R * 0.9)
  ctx.strokeStyle = accent
  ctx.lineWidth = 3
  ctx.lineCap = 'round'
  ctx.stroke()

  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.fillStyle = '#fafafa'
  ctx.font = `700 ${Math.min(w, h) * 0.16}px ${MONO}`
  ctx.fillText('23:59:59', cx, cy + R * 1.28)

  ctx.fillStyle = 'rgba(249,115,22,0.75)'
  ctx.font = `500 ${Math.min(w, h) * 0.032}px ${MONO}`
  ctx.fillText('ONE SECOND TO MIDNIGHT', cx, cy + R * 1.32 + Math.min(w, h) * 0.09)

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 10. 沙漏倒流：粒子向上流 */
function hourglassReverse(ctx, w, h, t, p) {
  const accent = accentOf(p, '#fbbf24')
  ctx.fillStyle = '#0a0806'
  ctx.fillRect(0, 0, w, h)

  const cx = w / 2
  const topY = h * 0.12
  const botY = h * 0.88
  const midY = h / 2
  const maxW = Math.min(w * 0.22, h * 0.34)

  // 沙漏轮廓
  ctx.beginPath()
  ctx.moveTo(cx - maxW, topY)
  ctx.lineTo(cx + maxW, topY)
  ctx.lineTo(cx + maxW * 0.12, midY)
  ctx.lineTo(cx + maxW, botY)
  ctx.lineTo(cx - maxW, botY)
  ctx.lineTo(cx - maxW * 0.12, midY)
  ctx.closePath()
  ctx.strokeStyle = 'rgba(251,191,36,0.35)'
  ctx.lineWidth = 2
  ctx.stroke()

  // 粒子自下而上
  const n = 120
  for (let i = 0; i < n; i++) {
    const seed = i * 0.618
    const prog = ((t * (0.12 + p.speed * 0.05) + seed) % 1)
    const y = botY - (botY - topY) * prog
    const spread = maxW * (1 - Math.abs(y - midY) / (h / 2)) * 0.9
    const x = cx + Math.sin(seed * 47 + t * 0.6) * spread
    const r = 1.2 + (i % 3) * 0.6
    ctx.beginPath()
    ctx.arc(x, y, r, 0, TAU)
    ctx.fillStyle = `rgba(251,191,36,${0.25 + prog * 0.6})`
    ctx.fill()
  }

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/* 通用：底部文案层 */
function drawCaption(ctx, w, h, p, color) {
  const text = p.title
  if (!text) return
  const maxW = w * 0.8
  const size = fitFont(ctx, text, maxW, Math.min(w, h) * 0.06, SANS)
  ctx.save()
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.font = `700 ${size}px ${SANS}`
  ctx.fillStyle = 'rgba(0,0,0,0.45)'
  ctx.fillText(text, w / 2 + 1, h * 0.9 + 1)
  ctx.fillStyle = color || '#fff'
  ctx.fillText(text, w / 2, h * 0.9)
  ctx.restore()
}

const RENDERERS_PART1 = {
  sharingan,
  kaleido_eye: kaleidoEye,
  infinite_pupil: infinitePupil,
  watcher,
  compound_eye: compoundEye,
  flip_clock: flipClock,
  countdown_bomb: countdownBomb,
  broken_clock: brokenClock,
  doomsday_clock: doomsdayClock,
  hourglass_reverse: hourglassReverse,
}

export { TAU, seeded, SANS, MONO, accentOf, fitFont, drawCaption }

/* ---------------- C. 文字 ---------------- */

/** 11. 循环警告跑马灯 */
function marqueeWarning(ctx, w, h, t, p) {
  const accent = accentOf(p, '#facc15')
  ctx.fillStyle = '#0a0a0a'
  ctx.fillRect(0, 0, w, h)

  const text = p.title || 'SYSTEM FAILURE'
  const bandH = h * 0.34
  const y = h / 2 - bandH / 2

  // 斜纹警示带
  ctx.save()
  ctx.beginPath()
  ctx.rect(0, y, w, bandH)
  ctx.clip()
  ctx.fillStyle = accent
  ctx.fillRect(0, y, w, bandH)
  ctx.fillStyle = 'rgba(0,0,0,0.55)'
  const stripe = 34
  const offset = (t * (30 + p.speed * 26)) % (stripe * 2)
  for (let x = -stripe * 4 + offset; x < w + stripe * 4; x += stripe * 2) {
    ctx.save()
    ctx.beginPath()
    ctx.moveTo(x, y)
    ctx.lineTo(x + stripe, y)
    ctx.lineTo(x + stripe - bandH * 0.4, y + bandH)
    ctx.lineTo(x - bandH * 0.4, y + bandH)
    ctx.closePath()
    ctx.fill()
    ctx.restore()
  }
  ctx.restore()

  // 文字跑马灯：白字 + 黑描边，多份循环铺满，保证画面内始终有字
  const size = bandH * 0.52
  ctx.font = `800 ${size}px ${SANS}`
  const tw = ctx.measureText(text).width
  const period = tw + w * 0.16
  const speedX = (w * 0.35 + p.speed * w * 0.12)
  const base = (t * speedX) % period
  ctx.textAlign = 'left'
  ctx.textBaseline = 'middle'
  ctx.lineJoin = 'round'
  ctx.lineWidth = Math.max(3, size * 0.11)
  for (let k = -1; k <= 3; k++) {
    const x = w - base + k * period
    if (x > w || x + tw < 0) continue
    ctx.strokeStyle = '#0a0a0a'
    ctx.strokeText(text, x, h / 2)
    ctx.fillStyle = '#fafafa'
    ctx.fillText(text, x, h / 2)
  }

  if (p.subtitle) {
    ctx.textAlign = 'center'
    ctx.fillStyle = 'rgba(255,255,255,0.4)'
    ctx.font = `500 ${Math.min(w, h) * 0.028}px ${MONO}`
    ctx.fillText(p.subtitle.toUpperCase(), w / 2, h * 0.86)
  }
}

/** 12. 弹幕墙 */
function danmakuWall(ctx, w, h, t, p) {
  const accent = accentOf(p, '#ffffff')
  ctx.fillStyle = '#08080c'
  ctx.fillRect(0, 0, w, h)

  const lines = (p.title || '好奇怪 / 再看一眼 / 这什么 / 有点上头')
    .split(/[/、,，]/).map((s) => s.trim()).filter(Boolean)
  const pool = lines.length ? lines : ['好奇怪']
  const N = 22
  const rnd = seeded(31)

  for (let i = 0; i < N; i++) {
    const lane = i % 9
    const y = h * 0.12 + (lane / 9) * h * 0.76
    const speedX = (60 + rnd() * 150) * (0.5 + p.speed * 0.25)
    const text = pool[i % pool.length]
    const size = Math.min(w, h) * (0.028 + rnd() * 0.02) * (1 + p.intensity * 0.06)
    ctx.font = `600 ${size}px ${SANS}`
    const tw = ctx.measureText(text).width
    const total = w + tw + 80
    const x = w + 40 - ((t * speedX + i * 137) % total)
    ctx.textAlign = 'left'
    ctx.textBaseline = 'middle'
    ctx.fillStyle = i % 5 === 0 ? accent : `rgba(255,255,255,${0.35 + (i % 4) * 0.14})`
    ctx.fillText(text, x, y)
  }
  if (p.subtitle) drawCaption(ctx, w, h, p, accent)
}

/** 13. 竖排霓虹字 */
function neonVertical(ctx, w, h, t, p) {
  const accent = accentOf(p, '#22d3ee')
  ctx.fillStyle = '#05070d'
  ctx.fillRect(0, 0, w, h)

  const text = (p.title || '霓虹').slice(0, 10)
  const chars = [...text]
  const colW = w / Math.max(chars.length, 1)
  const size = Math.min(colW * 0.62, h * 0.2)

  chars.forEach((ch, i) => {
    const x = colW * (i + 0.5)
    const speedY = (h * 0.12) * (0.6 + p.speed * 0.22)
    const y = h - ((t * speedY + i * h * 0.18) % (h + size * 2)) + size
    ctx.save()
    ctx.textAlign = 'center'
    ctx.textBaseline = 'middle'
    ctx.font = `700 ${size}px ${SANS}`
    // 霓虹辉光
    ctx.shadowColor = accent
    ctx.shadowBlur = 22 + p.intensity * 8
    ctx.fillStyle = accent
    ctx.fillText(ch, x, y)
    ctx.shadowBlur = 8
    ctx.fillStyle = 'rgba(255,255,255,0.9)'
    ctx.fillText(ch, x, y)
    ctx.restore()
  })

  // 扫描光带
  const scanY = (t * h * 0.5) % h
  const grad = ctx.createLinearGradient(0, scanY - 40, 0, scanY + 40)
  grad.addColorStop(0, 'rgba(34,211,238,0)')
  grad.addColorStop(0.5, 'rgba(34,211,238,0.12)')
  grad.addColorStop(1, 'rgba(34,211,238,0)')
  ctx.fillStyle = grad
  ctx.fillRect(0, scanY - 40, w, 80)

  if (p.subtitle) drawCaption(ctx, w, h, p, accent)
}

/** 14. 打字机：反复敲出又删除 */
function typewriter(ctx, w, h, t, p) {
  const accent = accentOf(p, '#4ade80')
  ctx.fillStyle = '#07090a'
  ctx.fillRect(0, 0, w, h)

  const full = p.title || '这句话会自己写完再删掉'
  const chars = [...full]
  const typing = 0.11 / (0.6 + p.speed * 0.16)
  const hold = 1.4
  const deleting = 0.045 / (0.6 + p.speed * 0.16)
  const cycle = chars.length * (typing + deleting) + hold * 2
  const phase = t % cycle

  let count = 0
  let cursorOn = true
  const typeEnd = chars.length * typing
  if (phase < typeEnd) {
    count = Math.floor(phase / typing)
    cursorOn = true
  } else if (phase < typeEnd + hold) {
    count = chars.length
    cursorOn = Math.floor(t * 3) % 2 === 0
  } else if (phase < typeEnd + hold + chars.length * deleting) {
    const d = (phase - typeEnd - hold) / deleting
    count = Math.max(0, chars.length - Math.floor(d))
    cursorOn = true
  } else {
    count = 0
    cursorOn = Math.floor(t * 3) % 2 === 0
  }

  const shown = chars.slice(0, count).join('')
  const size = Math.min(w, h) * 0.1
  ctx.font = `600 ${size}px ${MONO}`
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.fillStyle = '#e8f5ee'
  ctx.fillText(shown, w / 2, h / 2)

  if (cursorOn) {
    const tw = ctx.measureText(shown).width
    ctx.fillStyle = accent
    ctx.fillRect(w / 2 + tw / 2 + size * 0.1, h / 2 - size * 0.42, size * 0.5, size * 0.84)
  }

  if (p.subtitle) drawCaption(ctx, w, h, p, accent)
}

/** 15. 乱码雨（Matrix） */
function matrixRain(ctx, w, h, t, p) {
  const accent = accentOf(p, '#22c55e')
  ctx.fillStyle = 'rgba(3,8,5,0.28)'
  ctx.fillRect(0, 0, w, h)

  const cols = Math.max(14, Math.floor(w / 26))
  const cell = w / cols
  const speed = 0.55 + p.speed * 0.35
  const glyphs = '01アイウエオカキクケコサシスセソタチツテトナニヌネノ0123456789'

  ctx.font = `500 ${cell * 0.92}px ${MONO}`
  ctx.textAlign = 'center'
  ctx.textBaseline = 'top'

  for (let i = 0; i < cols; i++) {
    const rnd = seeded(i * 977)
    const offset = rnd() * h
    const y = ((t * speed * h * 0.55 + offset) % (h + h * 0.3)) - h * 0.15
    const len = 12 + Math.floor(rnd() * 14)
    for (let k = 0; k < len; k++) {
      const gy = y - k * cell * 1.05
      if (gy < -cell || gy > h) continue
      const ch = glyphs[Math.floor((rnd() * 1000 + k * 7 + Math.floor(t * 8)) % glyphs.length)]
      const alpha = Math.max(0, 1 - k / len)
      ctx.fillStyle = k === 0 ? '#d9ffe6' : `rgba(34,197,94,${alpha * 0.85})`
      ctx.fillText(ch, cell * (i + 0.5), gy)
    }
  }

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/* ---------------- D. 伪系统 ---------------- */

/** 16. 假终端 */
function fakeTerminal(ctx, w, h, t, p) {
  const accent = accentOf(p, '#4ade80')
  ctx.fillStyle = '#0b0f0c'
  ctx.fillRect(0, 0, w, h)

  const pad = w * 0.06
  const lh = h * 0.062
  const lines = [
    '$ tapiito --scan',
    '[ ok ] nfc module online',
    '[ ok ] mounting /dev/tag0',
    '▶ reading sector 0x00 …',
    `▶ payload: ${p.title || 'HELLO'}`,
    '[warn] checksum mismatch, retrying',
    '▶ decrypting 3 blocks …',
    '✗ permission denied',
    '$ _',
  ]
  const count = Math.min(lines.length, Math.floor((t * (2.2 + p.speed * 0.9)) % (lines.length + 3)))
  const blink = Math.floor(t * 2.5) % 2 === 0

  ctx.font = `500 ${Math.min(h * 0.05, w * 0.026)}px ${MONO}`
  ctx.textAlign = 'left'
  ctx.textBaseline = 'top'

  for (let i = 0; i < count; i++) {
    const line = lines[i]
    ctx.fillStyle = line.startsWith('✗') ? '#f87171'
      : line.startsWith('[warn]') ? '#fbbf24' : '#86efac'
    ctx.fillText(line, pad, pad + i * lh)
  }
  // 光标
  if (blink) {
    const y = pad + Math.min(count, lines.length) * lh
    ctx.fillStyle = accent
    ctx.fillRect(pad, y, Math.min(h * 0.028, w * 0.014), Math.min(h * 0.045, w * 0.023))
  }

  // 边框光晕
  ctx.strokeStyle = 'rgba(74,222,128,0.18)'
  ctx.lineWidth = 1
  ctx.strokeRect(pad * 0.4, pad * 0.4, w - pad * 0.8, h - pad * 0.8)
}

/** 17. 假监控面板 */
function fakeMonitor(ctx, w, h, t, p) {
  const accent = accentOf(p, '#38bdf8')
  ctx.fillStyle = '#070b12'
  ctx.fillRect(0, 0, w, h)

  const items = [
    { label: 'CPU', base: 0.5, color: '#38bdf8' },
    { label: 'MEM', base: 0.66, color: '#a78bfa' },
    { label: 'NET', base: 0.32, color: '#34d399' },
    { label: 'TEMP', base: 0.72, color: '#fb923c' },
  ]
  const padX = w * 0.1
  const barW = w * 0.8
  const top = h * 0.24
  const gap = h * 0.16

  ctx.textAlign = 'left'
  ctx.textBaseline = 'middle'

  items.forEach((it, i) => {
    const y = top + i * gap
    const spike = Math.max(0, Math.sin(t * (1.7 + i * 0.9) + i) * 0.28)
    const val = Math.min(1, it.base + spike * (0.5 + p.intensity * 0.2))
    ctx.fillStyle = 'rgba(255,255,255,0.5)'
    ctx.font = `600 ${h * 0.045}px ${MONO}`
    ctx.fillText(it.label, padX, y)

    // 轨道
    ctx.fillStyle = 'rgba(255,255,255,0.08)'
    ctx.fillRect(padX + w * 0.1, y - h * 0.014, barW * 0.72, h * 0.028)
    // 填充
    ctx.fillStyle = it.color
    ctx.fillRect(padX + w * 0.1, y - h * 0.014, barW * 0.72 * val, h * 0.028)
    // 数值
    ctx.fillStyle = it.color
    ctx.font = `700 ${h * 0.042}px ${MONO}`
    ctx.fillText(`${Math.round(val * 100)}%`, padX + barW * 0.86, y)
  })

  // 顶部标题
  ctx.fillStyle = 'rgba(56,189,248,0.8)'
  ctx.font = `600 ${h * 0.038}px ${MONO}`
  ctx.fillText('SYSTEM MONITOR', padX, h * 0.12)
  if (p.title) {
    ctx.textAlign = 'right'
    ctx.fillStyle = 'rgba(255,255,255,0.55)'
    ctx.fillText(p.title, w - padX, h * 0.12)
  }
}

/** 18. 蓝屏 */
function blueScreen(ctx, w, h, t, p) {
  const accent = accentOf(p, '#9ecbff')
  ctx.fillStyle = '#0b57d0'
  ctx.fillRect(0, 0, w, h)

  const padX = w * 0.1
  let y = h * 0.2
  ctx.textAlign = 'left'
  ctx.textBaseline = 'top'

  ctx.fillStyle = '#fff'
  ctx.font = `300 ${h * 0.16}px ${SANS}`
  ctx.fillText(':(', padX, y)
  y += h * 0.22

  ctx.font = `500 ${Math.min(h * 0.05, w * 0.026)}px ${SANS}`
  ctx.fillText('遇到了一个问题，需要重新启动。', padX, y)
  y += h * 0.09

  ctx.font = `400 ${Math.min(h * 0.035, w * 0.02)}px ${MONO}`
  ctx.fillStyle = 'rgba(255,255,255,0.78)'
  ctx.fillText(`停止代码: ${p.title || 'NFC_UNREACHABLE'}`, padX, y)
  y += h * 0.06
  // 进度永远 0%
  const prog = Math.floor((Math.sin(t * 0.7) + 1) * 2)
  ctx.fillText(`${prog}% 完成`, padX, y)
  y += h * 0.09
  ctx.fillStyle = 'rgba(255,255,255,0.55)'
  ctx.font = `400 ${Math.min(h * 0.03, w * 0.017)}px ${SANS}`
  ctx.fillText(p.subtitle || '如果问题持续出现，请联系设备管理员。', padX, y)

  // 右下角二维码方块（装饰）
  const s = h * 0.16
  const bx = w - padX - s
  const by = h - h * 0.16 - s
  const rnd = seeded(13)
  ctx.fillStyle = 'rgba(255,255,255,0.9)'
  for (let i = 0; i < 6; i++) {
    for (let j = 0; j < 6; j++) {
      if (rnd() > 0.45) {
        ctx.fillRect(bx + (i * s) / 6, by + (j * s) / 6, s / 6 - 1, s / 6 - 1)
      }
    }
  }
  ctx.strokeStyle = 'rgba(158,203,255,0.7)'
  ctx.lineWidth = 2
  ctx.strokeRect(bx, by, s, s)
  void accent
}

/** 19. 雷达扫描 */
function radarScan(ctx, w, h, t, p) {
  const accent = accentOf(p, '#22c55e')
  const cx = w / 2
  const cy = h / 2
  const R = Math.min(w, h) * 0.42

  ctx.fillStyle = '#050d08'
  ctx.fillRect(0, 0, w, h)

  // 同心圆
  for (let i = 1; i <= 4; i++) {
    ctx.beginPath()
    ctx.arc(cx, cy, (R * i) / 4, 0, TAU)
    ctx.strokeStyle = `rgba(34,197,94,${0.1 + i * 0.03})`
    ctx.lineWidth = 1
    ctx.stroke()
  }
  // 十字线
  ctx.beginPath()
  ctx.moveTo(cx - R, cy); ctx.lineTo(cx + R, cy)
  ctx.moveTo(cx, cy - R); ctx.lineTo(cx, cy + R)
  ctx.strokeStyle = 'rgba(34,197,94,0.12)'
  ctx.stroke()

  // 扫描扇形
  const ang = t * (0.9 + p.speed * 0.3)
  const sweep = ctx.createConicGradient
    ? ctx.createConicGradient(ang - 0.9, cx, cy)
    : null
  if (sweep) {
    sweep.addColorStop(0, 'rgba(34,197,94,0)')
    sweep.addColorStop(0.12, 'rgba(34,197,94,0.35)')
    sweep.addColorStop(0.16, 'rgba(34,197,94,0)')
    sweep.addColorStop(1, 'rgba(34,197,94,0)')
    ctx.beginPath()
    ctx.moveTo(cx, cy)
    ctx.arc(cx, cy, R, 0, TAU)
    ctx.fillStyle = sweep
    ctx.fill()
  }

  // 光点：出现后被扫描线"点亮"再熄灭
  const N = 8
  for (let i = 0; i < N; i++) {
    const rnd = seeded(i * 131 + 7)
    const a = rnd() * TAU
    const r = R * (0.25 + rnd() * 0.7)
    const blink = (Math.sin(ang - a + t * 0.4) + 1) / 2
    if (blink < 0.35) continue
    const x = cx + Math.cos(a) * r
    const y = cy + Math.sin(a) * r
    ctx.beginPath()
    ctx.arc(x, y, 3 + blink * 3, 0, TAU)
    ctx.fillStyle = `rgba(190,255,210,${blink})`
    ctx.fill()
  }

  // 扫描线
  ctx.beginPath()
  ctx.moveTo(cx, cy)
  ctx.lineTo(cx + Math.cos(ang) * R, cy + Math.sin(ang) * R)
  ctx.strokeStyle = accent
  ctx.lineWidth = 2
  ctx.stroke()

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 20. 频谱分析 */
function spectrumBars(ctx, w, h, t, p) {
  const accent = accentOf(p, '#f472b6')
  ctx.fillStyle = '#0a0510'
  ctx.fillRect(0, 0, w, h)

  const N = 48
  const barW = (w * 0.9) / N
  const baseY = h * 0.86
  const maxH = h * 0.62

  for (let i = 0; i < N; i++) {
    const rnd = seeded(i * 313 + 5)
    const freq = i / N
    // 低频高、高频低，叠加多个正弦制造"跳动"感
    const wave = Math.abs(
      Math.sin(t * (2.1 + freq * 5) + i * 0.6) * 0.5
      + Math.sin(t * (1.3 + rnd()) + i * 0.2) * 0.35
      + Math.sin(t * 3.7 + i * 1.7) * 0.15,
    )
    const env = Math.pow(1 - freq, 0.6) * (0.55 + p.intensity * 0.12)
    const bh = Math.max(2, Math.min(1, wave * env) * maxH)

    const x = w * 0.05 + i * barW
    const grad = ctx.createLinearGradient(0, baseY - bh, 0, baseY)
    grad.addColorStop(0, accent)
    grad.addColorStop(1, 'rgba(244,114,182,0.15)')
    ctx.fillStyle = grad
    ctx.fillRect(x + barW * 0.16, baseY - bh, barW * 0.68, bh)
  }

  ctx.beginPath()
  ctx.moveTo(w * 0.05, baseY)
  ctx.lineTo(w * 0.95, baseY)
  ctx.strokeStyle = 'rgba(244,114,182,0.4)'
  ctx.lineWidth = 1
  ctx.stroke()

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/* ---------------- E. 故障与复古 ---------------- */

/** 21. CRT 雪花屏 */
function crtSnow(ctx, w, h, t, p) {
  const accent = accentOf(p, '#d4d4d8')
  ctx.fillStyle = '#0b0b0d'
  ctx.fillRect(0, 0, w, h)

  // 噪点：用少量大色块模拟（逐像素太慢）
  const cell = 4
  const rnd = seeded(Math.floor(t * (10 + p.speed * 6)))
  for (let y = 0; y < h; y += cell) {
    for (let x = 0; x < w; x += cell) {
      const v = rnd()
      if (v > 0.62) {
        const g = Math.floor(v * 255)
        ctx.fillStyle = `rgba(${g},${g},${g},${0.35 + v * 0.4})`
        ctx.fillRect(x, y, cell, cell)
      }
    }
  }

  // 横向滚动亮条
  const bandY = (t * h * 0.4) % h
  const grad = ctx.createLinearGradient(0, bandY - h * 0.06, 0, bandY + h * 0.06)
  grad.addColorStop(0, 'rgba(255,255,255,0)')
  grad.addColorStop(0.5, 'rgba(255,255,255,0.12)')
  grad.addColorStop(1, 'rgba(255,255,255,0)')
  ctx.fillStyle = grad
  ctx.fillRect(0, bandY - h * 0.06, w, h * 0.12)

  // 扫描线
  ctx.fillStyle = 'rgba(0,0,0,0.25)'
  for (let y = 0; y < h; y += 3) ctx.fillRect(0, y, w, 1)

  // 偶发亮度跳动
  if ((t % 3.4) > 3.2) {
    ctx.fillStyle = 'rgba(255,255,255,0.18)'
    ctx.fillRect(0, 0, w, h)
  }

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 22. 扫描线撕裂（RGB 错位） */
function scanlineGlitch(ctx, w, h, t, p) {
  const accent = accentOf(p, '#f43f5e')
  ctx.fillStyle = '#060608'
  ctx.fillRect(0, 0, w, h)

  const phrase = p.title || 'SIGNAL LOST'
  const size = Math.min(w, h) * 0.16
  ctx.font = `800 ${size}px ${SANS}`
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'

  // 每 1.2 秒触发一次撕裂
  const burst = (t % 1.2) < 0.16
  const jitter = burst ? (p.intensity + 1) * 8 : 0
  const dx = (Math.random() - 0.5) * jitter
  const dy = (Math.random() - 0.5) * jitter * 0.4

  // RGB 三色错位
  ctx.globalCompositeOperation = 'lighter'
  ctx.fillStyle = 'rgba(255,60,60,0.85)'
  ctx.fillText(phrase, w / 2 + dx - jitter * 0.5, h / 2 + dy)
  ctx.fillStyle = 'rgba(60,255,120,0.8)'
  ctx.fillText(phrase, w / 2 + dx, h / 2 + dy)
  ctx.fillStyle = 'rgba(80,140,255,0.85)'
  ctx.fillText(phrase, w / 2 + dx + jitter * 0.5, h / 2 + dy)
  ctx.globalCompositeOperation = 'source-over'

  // 随机横向切片偏移
  const slices = 14
  const rnd = seeded(Math.floor(t * 6))
  for (let i = 0; i < slices; i++) {
    if (rnd() < 0.55) continue
    const sy = (i / slices) * h
    const sh = h / slices
    const shift = (rnd() - 0.5) * (burst ? 40 : 10)
    const img = ctx.getImageData(0, sy, w, Math.max(1, Math.floor(sh)))
    ctx.putImageData(img, shift, sy)
  }

  // 扫描线
  ctx.fillStyle = 'rgba(0,0,0,0.22)'
  for (let y = 0; y < h; y += 3) ctx.fillRect(0, y, w, 1)

  if (p.subtitle) drawCaption(ctx, w, h, p, accent)
}

/** 23. VHS 录像带 */
function vhsTape(ctx, w, h, t, p) {
  const accent = accentOf(p, '#f87171')
  ctx.fillStyle = '#0a0808'
  ctx.fillRect(0, 0, w, h)

  // 磁带底色噪点
  const cell = 6
  const rnd = seeded(Math.floor(t * 9))
  for (let y = 0; y < h; y += cell) {
    for (let x = 0; x < w; x += cell) {
      if (rnd() > 0.72) {
        ctx.fillStyle = `rgba(120,110,100,${0.08 + rnd() * 0.16})`
        ctx.fillRect(x, y, cell, cell)
      }
    }
  }

  // 中央文案
  const phrase = p.title || 'PLAY ▶'
  ctx.font = `700 ${Math.min(w, h) * 0.13}px ${MONO}`
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.fillStyle = 'rgba(255,255,255,0.86)'
  ctx.fillText(phrase, w / 2 + Math.sin(t * 8) * 1.6, h / 2)

  // REC 红点
  const blink = Math.floor(t * 1.6) % 2 === 0
  ctx.beginPath()
  ctx.arc(w * 0.08, h * 0.12, Math.min(w, h) * 0.016, 0, TAU)
  ctx.fillStyle = blink ? accent : 'rgba(248,113,113,0.2)'
  ctx.fill()
  ctx.textAlign = 'left'
  ctx.fillStyle = 'rgba(255,255,255,0.75)'
  ctx.font = `600 ${Math.min(w, h) * 0.038}px ${MONO}`
  ctx.fillText('REC', w * 0.08 + Math.min(w, h) * 0.028, h * 0.12)

  // 时间码递增
  const total = Math.floor(t)
  const tc = `00:${String(Math.floor(total / 60) % 60).padStart(2, '0')}:${String(total % 60).padStart(2, '0')}`
  ctx.textAlign = 'right'
  ctx.fillStyle = 'rgba(255,255,255,0.8)'
  ctx.fillText(tc, w * 0.92, h * 0.12)
  ctx.textAlign = 'left'
  ctx.fillStyle = 'rgba(255,255,255,0.45)'
  ctx.fillText('SP  0:00:00', w * 0.08, h * 0.88)

  // 横向磁带条纹
  ctx.fillStyle = 'rgba(255,255,255,0.05)'
  ctx.fillRect(0, h * 0.7, w, 2)
  ctx.fillRect(0, h * 0.78, w, 1)
}

/** 24. 老式示波器 */
function oscilloscope(ctx, w, h, t, p) {
  const accent = accentOf(p, '#4ade80')
  ctx.fillStyle = '#04120a'
  ctx.fillRect(0, 0, w, h)

  // 网格
  ctx.strokeStyle = 'rgba(74,222,128,0.12)'
  ctx.lineWidth = 1
  const grid = 8
  for (let i = 1; i < grid; i++) {
    ctx.beginPath()
    ctx.moveTo(0, (h * i) / grid); ctx.lineTo(w, (h * i) / grid)
    ctx.moveTo((w * i) / grid, 0); ctx.lineTo((w * i) / grid, h)
    ctx.stroke()
  }

  // 波形
  const amp = h * 0.18 * (0.6 + p.intensity * 0.12)
  ctx.beginPath()
  for (let x = 0; x <= w; x += 3) {
    const u = x / w
    const y = h / 2
      + Math.sin(u * 14 + t * (2.4 + p.speed * 0.6)) * amp
      + Math.sin(u * 37 - t * 3.1) * amp * 0.35
      + Math.sin(u * 91 + t * 5.3) * amp * 0.12
    if (x === 0) ctx.moveTo(x, y)
    else ctx.lineTo(x, y)
  }
  ctx.strokeStyle = accent
  ctx.lineWidth = 2
  ctx.shadowColor = accent
  ctx.shadowBlur = 12
  ctx.stroke()
  ctx.shadowBlur = 0

  // 读数
  ctx.fillStyle = 'rgba(74,222,128,0.8)'
  ctx.font = `600 ${Math.min(w, h) * 0.034}px ${MONO}`
  ctx.textAlign = 'left'
  ctx.textBaseline = 'top'
  ctx.fillText(`CH1  5.0V/DIV   20ms/DIV   ${(1.2 + Math.sin(t) * 0.4).toFixed(2)}kHz`, w * 0.04, h * 0.06)

  if (p.title) {
    ctx.textAlign = 'right'
    ctx.fillText(p.title, w * 0.96, h * 0.06)
  }
}

/** 25. 跳帧故障（随机色块错位） */
function frameGlitch(ctx, w, h, t, p) {
  const accent = accentOf(p, '#a855f7')
  ctx.fillStyle = '#07060b'
  ctx.fillRect(0, 0, w, h)

  // 每 2 秒重排一次色块
  const epoch = Math.floor(t / 2)
  const rnd = seeded(epoch * 7919 + 13)
  const blocks = 26 + p.intensity * 4

  for (let i = 0; i < blocks; i++) {
    const bw = w * (0.06 + rnd() * 0.26)
    const bh = h * (0.04 + rnd() * 0.22)
    const x = rnd() * (w - bw)
    const y = rnd() * (h - bh)
    const hue = Math.floor(rnd() * 360)
    ctx.fillStyle = `hsla(${hue}, 70%, 55%, ${0.12 + rnd() * 0.22})`
    ctx.fillRect(x, y, bw, bh)
    // 细边框强化"错位感"
    ctx.strokeStyle = `hsla(${hue}, 90%, 70%, 0.35)`
    ctx.lineWidth = 1
    ctx.strokeRect(x, y, bw, bh)
  }

  // 中央文字被切成两半
  if (p.title) {
    const text = p.title
    const size = Math.min(w, h) * 0.14
    ctx.font = `800 ${size}px ${SANS}`
    ctx.textAlign = 'center'
    ctx.textBaseline = 'middle'
    const half = Math.floor(text.length / 2)
    ctx.fillStyle = '#f4f4f5'
    ctx.fillText(text.slice(0, half), w / 2 - (t * 30) % 40, h / 2 - size * 0.5)
    ctx.fillStyle = accent
    ctx.fillText(text.slice(half), w / 2 + (t * 26) % 34, h / 2 + size * 0.5)
  }
}

/* ---------------- F. 奇趣 ---------------- */

/** 26. 无限镜像 */
function infiniteMirror(ctx, w, h, t, p) {
  const accent = accentOf(p, '#e2e8f0')
  ctx.fillStyle = '#04050a'
  ctx.fillRect(0, 0, w, h)

  const layers = 9
  const phrase = p.title || '∞'
  for (let i = 0; i < layers; i++) {
    // 每一层按时间缓慢缩放，形成"往里掉"的错觉
    const base = 0.92 - (i / layers) * 0.86
    const breathe = Math.sin(t * 0.8 - i * 0.5) * 0.012
    const kw = w * (base + breathe)
    const kh = h * (base + breathe)
    const x = (w - kw) / 2
    const y = (h - kh) / 2
    const alpha = 0.85 - (i / layers) * 0.7

    ctx.strokeStyle = `rgba(226,232,240,${alpha})`
    ctx.lineWidth = Math.max(0.6, 2 - i * 0.18)
    ctx.strokeRect(x, y, kw, kh)

    // 每层内框加一点内发光
    if (i < 4) {
      ctx.strokeStyle = `rgba(120,180,255,${alpha * 0.35})`
      ctx.strokeRect(x + kw * 0.03, y + kh * 0.03, kw * 0.94, kh * 0.94)
    }

    if (i % 2 === 0) {
      ctx.save()
      ctx.textAlign = 'center'
      ctx.textBaseline = 'middle'
      ctx.font = `700 ${Math.max(8, kh * 0.09)}px ${MONO}`
      ctx.fillStyle = `rgba(226,232,240,${alpha * 0.75})`
      ctx.fillText(phrase, w / 2, y + kh * 0.5)
      ctx.restore()
    }
  }

  // 中心亮点
  const gl = ctx.createRadialGradient(w / 2, h / 2, 0, w / 2, h / 2, Math.min(w, h) * 0.08)
  gl.addColorStop(0, `rgba(200,225,255,${0.5 + Math.sin(t * 2) * 0.2})`)
  gl.addColorStop(1, 'rgba(200,225,255,0)')
  ctx.fillStyle = gl
  ctx.fillRect(0, 0, w, h)

  if (p.subtitle) drawCaption(ctx, w, h, p, accent)
}

/** 27. 万花筒（三角镜面） */
function kaleidoscope(ctx, w, h, t, p) {
  const accent = accentOf(p, '#818cf8')
  ctx.fillStyle = '#05040a'
  ctx.fillRect(0, 0, w, h)

  const cx = w / 2
  const cy = h / 2
  const R = Math.min(w, h) * 0.52
  const petals = 8
  const rot = t * (0.12 + p.speed * 0.06)

  ctx.save()
  ctx.translate(cx, cy)
  for (let i = 0; i < petals; i++) {
    ctx.save()
    ctx.rotate(rot + (i * TAU) / petals)
    // 镜像另一半
    for (const dir of [1, -1]) {
      ctx.save()
      ctx.scale(dir, 1)
      ctx.beginPath()
      ctx.moveTo(0, 0)
      ctx.quadraticCurveTo(R * 0.42, -R * 0.16, R * 0.92, -R * 0.03)
      ctx.quadraticCurveTo(R * 0.5, R * 0.2, 0, 0)
      ctx.closePath()
      const g = ctx.createLinearGradient(0, 0, R, 0)
      g.addColorStop(0, 'rgba(129,140,248,0.05)')
      g.addColorStop(0.6, `rgba(${140 + i * 8}, ${120 + i * 6}, 248, ${0.16 + p.intensity * 0.03})`)
      g.addColorStop(1, 'rgba(56,189,248,0.05)')
      ctx.fillStyle = g
      ctx.fill()
      ctx.strokeStyle = accent
      ctx.globalAlpha = 0.22
      ctx.lineWidth = 1
      ctx.stroke()
      ctx.globalAlpha = 1
      ctx.restore()
    }
    ctx.restore()
  }
  ctx.restore()

  // 中心圆
  ctx.beginPath()
  ctx.arc(cx, cy, Math.min(w, h) * 0.05, 0, TAU)
  ctx.fillStyle = 'rgba(15,10,30,0.9)'
  ctx.fill()
  ctx.strokeStyle = accent
  ctx.lineWidth = 1.5
  ctx.stroke()

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 28. 心跳失控 */
function heartbeat(ctx, w, h, t, p) {
  const accent = accentOf(p, '#22d3ee')
  ctx.fillStyle = '#040a0c'
  ctx.fillRect(0, 0, w, h)

  // BPM 从 60 涨到 300 再归零
  const cycle = 14
  const phase = (t % cycle) / cycle
  const bpm = Math.round(60 + Math.pow(phase, 2) * 260)

  // 网格
  ctx.strokeStyle = 'rgba(34,211,238,0.1)'
  ctx.lineWidth = 1
  for (let i = 1; i < 8; i++) {
    ctx.beginPath()
    ctx.moveTo(0, (h * i) / 8); ctx.lineTo(w, (h * i) / 8)
    ctx.stroke()
  }

  // 心电波形：一个 QRS 复合波在横向滚动
  const beat = 1 / (bpm / 60)
  const scroll = (t % beat) / beat
  ctx.beginPath()
  for (let x = 0; x <= w; x += 2) {
    const u = x / w
    const local = (u + scroll) % 1
    let v = 0
    // 平直段 + 尖峰
    if (local > 0.42 && local < 0.52) v = Math.sin((local - 0.42) / 0.1 * Math.PI) * 0.22
    if (local > 0.5 && local < 0.56) v = -Math.sin((local - 0.5) / 0.06 * Math.PI) * 1.0
    if (local > 0.56 && local < 0.68) v = Math.sin((local - 0.56) / 0.12 * Math.PI) * 0.3
    const y = h * 0.56 - v * h * 0.26
    if (x === 0) ctx.moveTo(x, y)
    else ctx.lineTo(x, y)
  }
  ctx.strokeStyle = accent
  ctx.lineWidth = 2.4
  ctx.shadowColor = accent
  ctx.shadowBlur = 14
  ctx.stroke()
  ctx.shadowBlur = 0

  // BPM 大字
  ctx.textAlign = 'left'
  ctx.textBaseline = 'top'
  ctx.fillStyle = accent
  ctx.font = `700 ${Math.min(w, h) * 0.17}px ${MONO}`
  ctx.fillText(String(bpm), w * 0.06, h * 0.1)
  ctx.fillStyle = 'rgba(34,211,238,0.6)'
  ctx.font = `600 ${Math.min(w, h) * 0.04}px ${MONO}`
  ctx.fillText('BPM', w * 0.06, h * 0.1 + Math.min(w, h) * 0.19)
  ctx.fillText(bpm > 200 ? 'ARRHYTHMIA' : 'SINUS RHYTHM', w * 0.06, h * 0.1 + Math.min(w, h) * 0.25)

  if (p.title) {
    ctx.textAlign = 'right'
    ctx.fillStyle = 'rgba(34,211,238,0.75)'
    ctx.font = `600 ${Math.min(w, h) * 0.04}px ${MONO}`
    ctx.fillText(p.title, w * 0.94, h * 0.12)
  }
}

/** 29. 深海生物（发光水母） */
function jellyfish(ctx, w, h, t, p) {
  const accent = accentOf(p, '#67e8f9')
  const bg = ctx.createLinearGradient(0, 0, 0, h)
  bg.addColorStop(0, '#02060f')
  bg.addColorStop(0.6, '#04121c')
  bg.addColorStop(1, '#010409')
  ctx.fillStyle = bg
  ctx.fillRect(0, 0, w, h)

  const N = 3 + p.intensity
  for (let i = 0; i < N; i++) {
    const rnd = seeded(i * 577 + 3)
    const baseX = w * (0.15 + rnd() * 0.7)
    const drift = Math.sin(t * (0.35 + rnd() * 0.3) + i) * w * 0.05
    const cx = baseX + drift
    const cy = h * (0.25 + rnd() * 0.45) + Math.sin(t * 0.5 + i * 1.4) * h * 0.05
    const R = Math.min(w, h) * (0.05 + rnd() * 0.05)
    const pulse = 1 + Math.sin(t * (1.2 + i * 0.3)) * 0.08

    // 伞盖
    ctx.save()
    ctx.beginPath()
    ctx.ellipse(cx, cy, R * 1.25 * pulse, R * 0.85 * pulse, 0, Math.PI, TAU)
    ctx.closePath()
    const head = ctx.createRadialGradient(cx, cy - R * 0.3, R * 0.1, cx, cy, R * 1.4)
    head.addColorStop(0, 'rgba(255,255,255,0.75)')
    head.addColorStop(0.35, accent)
    head.addColorStop(1, 'rgba(8,47,73,0.05)')
    ctx.fillStyle = head
    ctx.shadowColor = accent
    ctx.shadowBlur = 26
    ctx.fill()
    ctx.shadowBlur = 0

    // 触须
    ctx.beginPath()
    for (let k = -2; k <= 2; k++) {
      const sx = cx + k * R * 0.42
      const len = R * (2.4 + Math.abs(k) * 0.4)
      ctx.moveTo(sx, cy)
      for (let s = 0; s <= 12; s++) {
        const u = s / 12
        const px = sx + Math.sin(t * 1.6 + u * 3.4 + k + i) * R * 0.45 * u
        const py = cy + len * u
        ctx.lineTo(px, py)
      }
    }
    ctx.strokeStyle = `rgba(103,232,249,${0.28 + p.intensity * 0.04})`
    ctx.lineWidth = 1.4
    ctx.stroke()
    ctx.restore()
  }

  // 气泡
  const rnd2 = seeded(91)
  for (let i = 0; i < 26; i++) {
    const bx = rnd2() * w
    const speed = 0.04 + rnd2() * 0.09
    const by = h - ((t * speed * h + rnd2() * h) % (h + 40)) + 20
    ctx.beginPath()
    ctx.arc(bx, by, 1 + rnd2() * 2.4, 0, TAU)
    ctx.fillStyle = `rgba(186,230,253,${0.12 + rnd2() * 0.25})`
    ctx.fill()
  }

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/** 30. 宇宙坐标 */
function cosmos(ctx, w, h, t, p) {
  const accent = accentOf(p, '#c4b5fd')
  ctx.fillStyle = '#02030a'
  ctx.fillRect(0, 0, w, h)

  const cx = w / 2
  const cy = h / 2

  // 星图：绕中心缓慢旋转
  const rot = t * 0.035
  const rnd = seeded(2024)
  for (let i = 0; i < 160; i++) {
    const a = rnd() * TAU + rot * (0.5 + rnd() * 0.8)
    const r = Math.pow(rnd(), 0.6) * Math.min(w, h) * 0.5
    const x = cx + Math.cos(a) * r
    const y = cy + Math.sin(a) * r * 0.62
    const tw = (Math.sin(t * (1 + rnd() * 2) + i) + 1) / 2
    const size = 0.6 + rnd() * 1.8
    ctx.beginPath()
    ctx.arc(x, y, size, 0, TAU)
    ctx.fillStyle = `rgba(226,232,240,${0.25 + tw * 0.7})`
    ctx.fill()
  }

  // 星座连线
  ctx.strokeStyle = 'rgba(196,181,253,0.16)'
  ctx.lineWidth = 1
  const pts = []
  const rnd2 = seeded(77)
  for (let i = 0; i < 12; i++) {
    const a = rnd2() * TAU + rot
    const r = Math.min(w, h) * (0.15 + rnd2() * 0.32)
    pts.push([cx + Math.cos(a) * r, cy + Math.sin(a) * r * 0.62])
  }
  ctx.beginPath()
  pts.forEach(([x, y], i) => (i ? ctx.lineTo(x, y) : ctx.moveTo(x, y)))
  ctx.stroke()

  // 坐标读数
  ctx.textAlign = 'left'
  ctx.textBaseline = 'top'
  ctx.fillStyle = 'rgba(196,181,253,0.7)'
  ctx.font = `600 ${Math.min(w, h) * 0.032}px ${MONO}`
  const ra = ((t * 0.6) % 24).toFixed(3)
  const dec = (Math.sin(t * 0.4) * 62).toFixed(3)
  ctx.fillText(`RA  ${ra}h   DEC ${dec > 0 ? '+' : ''}${dec}°`, w * 0.05, h * 0.08)
  ctx.fillStyle = 'rgba(196,181,253,0.4)'
  ctx.fillText(`EPOCH J2000   SCAN ${(t * 12).toFixed(0).padStart(6, '0')}`, w * 0.05, h * 0.08 + Math.min(w, h) * 0.05)

  if (p.title) drawCaption(ctx, w, h, p, accent)
}

/* ---------------- 汇总导出 ---------------- */

const RENDERERS_PART2 = {
  marquee_warning: marqueeWarning,
  danmaku_wall: danmakuWall,
  neon_vertical: neonVertical,
  typewriter,
  matrix_rain: matrixRain,
  fake_terminal: fakeTerminal,
  fake_monitor: fakeMonitor,
  blue_screen: blueScreen,
  radar_scan: radarScan,
  spectrum_bars: spectrumBars,
  crt_snow: crtSnow,
  scanline_glitch: scanlineGlitch,
  vhs_tape: vhsTape,
  oscilloscope,
  frame_glitch: frameGlitch,
  infinite_mirror: infiniteMirror,
  kaleidoscope,
  heartbeat,
  jellyfish,
  cosmos,
}

export const RENDERERS = { ...RENDERERS_PART1, ...RENDERERS_PART2 }

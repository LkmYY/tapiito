<template>
  <div ref="wrap" class="weird-stage">
    <canvas ref="cv" class="weird-canvas" />
    <div v-if="hint" class="weird-hint">{{ hint }}</div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount, watch } from 'vue'
import { RENDERERS } from './renderers.js'

const props = defineProps({
  payload: { type: Object, default: () => ({}) },
  styleKey: { type: String, default: 'sharingan' },
  nfcId: { type: String, default: '' },
})

const wrap = ref(null)
const cv = ref(null)
let ctx = null
let raf = 0
let startTs = 0
let ro = null
const pointer = { x: 0.5, y: 0.5 }
const size = { w: 0, h: 0, dpr: 1 }

/** 表单参数：文案、速度、强度、主题色 */
const p = computed(() => {
  const d = props.payload || {}
  return {
    title: String(d.title || '').trim(),
    subtitle: String(d.subtitle || '').trim(),
    speed: clampNum(d.speed, 1, 5, 3),
    intensity: clampNum(d.intensity, 1, 5, 3),
    accent: /^#[0-9a-fA-F]{3,8}$/.test(d.accent || '') ? d.accent : '',
    nfcId: props.nfcId,
    pointer,
  }
})

const hint = computed(() => p.value.subtitle || '')

function clampNum(v, min, max, def) {
  const n = Number(v)
  if (!Number.isFinite(n)) return def
  return Math.min(max, Math.max(min, n))
}

function resize() {
  const el = wrap.value
  const canvas = cv.value
  if (!el || !canvas) return
  const rect = el.getBoundingClientRect()
  const dpr = Math.min(window.devicePixelRatio || 1, 2)
  size.w = Math.max(320, rect.width)
  // 强制 16:9，超出视口则按高度反推宽度
  let w = size.w
  let h = (w * 9) / 16
  const maxH = window.innerHeight
  if (h > maxH) {
    h = maxH
    w = (h * 16) / 9
  }
  size.h = h
  size.dpr = dpr
  canvas.style.width = `${w}px`
  canvas.style.height = `${h}px`
  canvas.width = Math.round(w * dpr)
  canvas.height = Math.round(h * dpr)
  ctx = canvas.getContext('2d')
  ctx.setTransform(dpr, 0, 0, dpr, 0, 0)
}

function onPointer(e) {
  const canvas = cv.value
  if (!canvas) return
  const rect = canvas.getBoundingClientRect()
  const cx = e.touches?.[0]?.clientX ?? e.clientX
  const cy = e.touches?.[0]?.clientY ?? e.clientY
  pointer.x = Math.min(1, Math.max(0, (cx - rect.left) / rect.width))
  pointer.y = Math.min(1, Math.max(0, (cy - rect.top) / rect.height))
}

function loop(ts) {
  if (!startTs) startTs = ts
  const t = (ts - startTs) / 1000
  const r = RENDERERS[props.styleKey] || RENDERERS.sharingan
  if (ctx) {
    ctx.clearRect(0, 0, size.w, size.h)
    try {
      r(ctx, size.w, size.h, t, p.value)
    } catch (err) {
      // 单个绘制器异常不应中断整个页面
      // eslint-disable-next-line no-console
      console.warn('[weird] renderer failed', props.styleKey, err)
    }
  }
  raf = requestAnimationFrame(loop)
}

onMounted(() => {
  resize()
  if (typeof ResizeObserver !== 'undefined') {
    ro = new ResizeObserver(() => resize())
    ro.observe(wrap.value)
  } else {
    window.addEventListener('resize', resize)
  }
  window.addEventListener('pointermove', onPointer, { passive: true })
  window.addEventListener('touchmove', onPointer, { passive: true })
  raf = requestAnimationFrame(loop)
})

onBeforeUnmount(() => {
  cancelAnimationFrame(raf)
  ro?.disconnect()
  window.removeEventListener('resize', resize)
  window.removeEventListener('pointermove', onPointer)
  window.removeEventListener('touchmove', onPointer)
})

// 切换风格时重置时间基准，避免动画从奇怪的相位开始
watch(() => props.styleKey, () => { startTs = 0 })
</script>

<style scoped>
.weird-stage {
  position: relative;
  width: 100%;
  min-height: 100vh;
  display: grid;
  place-items: center;
  background: #05060a;
  overflow: hidden;
}
.weird-canvas {
  display: block;
  max-width: 100%;
  max-height: 100vh;
}
.weird-hint {
  position: absolute;
  left: 50%;
  bottom: 18px;
  transform: translateX(-50%);
  font-size: 11px;
  letter-spacing: 0.28em;
  text-transform: uppercase;
  color: rgba(255, 255, 255, 0.32);
  pointer-events: none;
  white-space: nowrap;
}
</style>

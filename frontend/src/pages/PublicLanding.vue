<template>
  <div class="min-h-screen transition-colors duration-500" :class="ambientClass">
    <div
      class="min-h-screen flex flex-col"
      :class="isLandscape
        ? 'w-full bg-black'
        : 'mx-auto max-w-md shadow-[0_0_0_1px_rgba(148,163,184,0.14),0_24px_64px_-20px_rgba(2,6,23,0.38)] bg-white dark:bg-black'"
    >
      <div v-if="loading" class="p-6 space-y-4">
        <div class="h-8 bg-gray-200 rounded w-3/4 animate-pulse"></div>
        <div class="h-32 bg-gray-200 rounded animate-pulse"></div>
        <div class="h-10 bg-gray-200 rounded animate-pulse"></div>
      </div>

      <!-- 需要暗号 / 链接失效 -->
      <LockScreen
        v-else-if="lockMode"
        :nfc-id="nfcId"
        :mode="lockMode"
        :hint="lockHint"
        @unlocked="onUnlocked"
      />

      <!-- 标签停用 / 不存在 / 加载失败：不报错，展示品牌宣传 -->
      <BrandFallback v-else-if="fallbackReason" :reason="fallbackReason" />

      <div v-else ref="contentRef" class="flex-1 flex flex-col">
        <component
          :is="comp"
          :payload="payload"
          :nfc-id="nfcId"
          v-motion
          :initial="{ opacity: 0, y: 20 }"
          :enter="{ opacity: 1, y: 0, transition: { duration: 500 } }"
        />
        <!-- 横屏沉浸式版面不挂平台页脚，避免打断画面 -->
        <PlatformFooter v-if="!isLandscape" :tone="footerTone" />
      </div>

      <!-- 防截图水印：标签级开关控制 -->
      <div v-if="watermark && !loading && !fallbackReason && !lockMode" class="wm-layer" aria-hidden="true">
        <span v-for="i in 12" :key="i" class="wm-item">{{ watermarkText }}</span>
      </div>
    </div>

    <!-- 轻量提示：复制微信号后反馈 -->
    <transition name="wb-toast">
      <div v-if="toast" class="wb-toast">{{ toast }}</div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute } from 'vue-router'
import { fetchNfcPage } from '../api/public'
import { componentFor, ambientFor, layoutFor } from '../templates/registry.js'
import PlatformFooter from '../components/PlatformFooter.vue'
import BrandFallback from '../components/BrandFallback.vue'
import { useWechatTap } from '../composables/useWechatTap.js'
import { getSessionId, getTicket, saveTicket } from '../composables/useAccess'
import LockScreen from '../components/LockScreen.vue'

const route = useRoute()
const nfcId = route.params.nfcId
const payload = ref({})
const type = ref('')
const customTemplate = ref('')
const loading = ref(true)
/** 非空时渲染品牌宣传页：disabled / notfound / error */
const fallbackReason = ref('')
/** 访问控制：passcode 需输暗号，ticket 表示链接已失效 */
const lockMode = ref('')
const lockHint = ref('')
/** 标签级水印与自动跳转 */
const watermark = ref(false)
const autoCloseSeconds = ref(0)
let redirectTimer = null
const watermarkText = computed(() => {
  const t = new Date()
  const hh = String(t.getHours()).padStart(2, '0')
  const mm = String(t.getMinutes()).padStart(2, '0')
  return `TAPIITO · ${nfcId} · ${hh}:${mm}`
})

const comp = computed(() => componentFor(type.value, customTemplate.value) || 'div')
/** 横屏类型（如「奇怪的想法」）需要全宽容器，不能用竖屏手机宽度 */
const isLandscape = computed(() => layoutFor(type.value) === 'landscape')

// —— 微信号点击复制 + 唤起微信 ——
const contentRef = ref(null)
const toast = ref('')
let toastTimer = null
const notify = (msg) => {
  toast.value = msg
  clearTimeout(toastTimer)
  toastTimer = setTimeout(() => { toast.value = '' }, 2400)
}
useWechatTap(contentRef, () => String(payload.value?.wechat || ''), notify)

const ambientClass = computed(() => (fallbackReason.value ? 'amb-plain' : ambientFor(type.value, customTemplate.value)))
const footerTone = computed(() => (ambientClass.value || '').startsWith('amb-dark') ? 'dark' : 'light')

/** 浏览器页签标题：优先用标签名，未命名时回退到平台名 */
const setTitle = (name) => {
  document.title = name ? `${name} · Tapiito` : 'Tapiito'
}

const loadPage = async () => {
  loading.value = true
  try {
    const { data } = await fetchNfcPage(nfcId, {
      ticket: getTicket(nfcId),
      session: getSessionId(),
    })
    lockMode.value = ''
    // 扫码即发的场景后端会返回新 ticket，存下来供刷新使用
    if (data.ticket) saveTicket(nfcId, data.ticket, data.expires_at || '')

    watermark.value = !!data.watermark
    autoCloseSeconds.value = Number(data.auto_close_seconds) || 0

    if (data.status && data.status !== 'enabled') {
      fallbackReason.value = 'disabled'
      setTitle('Tapiito')
    } else if (!data.template_type) {
      fallbackReason.value = 'notfound'
      setTitle('Tapiito')
    } else {
      type.value = data.template_type
      customTemplate.value = data.custom_template || ''
      payload.value = data.template_payload || {}
      setTitle(data.name)
      startAutoClose()
    }
  } catch (err) {
    const status = err?.response?.status
    const body = err?.response?.data || {}
    if (status === 403) {
      // 需要暗号，或 ticket 已失效（被用掉 / 过期 / 转发到别的浏览器）
      lockMode.value = body.required === 'passcode' ? 'passcode' : 'ticket'
      lockHint.value = body.hint || ''
    } else {
      fallbackReason.value = status === 404 ? 'notfound' : 'error'
      setTitle('Tapiito')
    }
  } finally {
    loading.value = false
  }
}

/** 自动跳转：到点离开页面回品牌页（浏览器不允许脚本直接关页签） */
const startAutoClose = () => {
  clearTimeout(redirectTimer)
  if (!autoCloseSeconds.value || autoCloseSeconds.value <= 0) return
  redirectTimer = setTimeout(() => {
    window.location.replace('https://tapiito.com/')
  }, autoCloseSeconds.value * 1000)
}

const onUnlocked = async ({ ticket, expiresAt }) => {
  saveTicket(nfcId, ticket, expiresAt || '')
  fallbackReason.value = ''
  await loadPage()
}

/**
 * 浏览器往返缓存会让「返回」直接恢复上一帧快照。
 * 解锁后按返回会看到锁屏那一帧（看着像又进了 403），所以恢复时重新校验一次：
 * 此时 ticket 已在本地，校验通过会直接渲染真实内容。
 */
const onPageShow = (e) => {
  if (e.persisted) loadPage()
}

onMounted(() => {
  loadPage()
  window.addEventListener('pageshow', onPageShow)
})

onBeforeUnmount(() => {
  window.removeEventListener('pageshow', onPageShow)
  clearTimeout(redirectTimer)
  clearTimeout(toastTimer)
})
</script>

<style scoped>
/* 微信号可点击：虚线下划线 + 手型，子组件内部需要 :deep */
:deep(.wechat-tappable) {
  cursor: pointer;
  text-decoration: underline;
  text-decoration-style: dotted;
  text-underline-offset: 3px;
  text-decoration-color: currentColor;
}
:deep(.wechat-tappable:active) {
  opacity: 0.6;
}

/* 复制成功后的轻提示 */
.wb-toast {
  position: fixed;
  left: 50%;
  bottom: 32px;
  z-index: 60;
  max-width: 86vw;
  padding: 10px 18px;
  border-radius: 999px;
  background: rgba(20, 24, 29, 0.92);
  color: #fff;
  font-size: 13px;
  line-height: 1.4;
  text-align: center;
  transform: translateX(-50%);
  backdrop-filter: blur(8px);
  box-shadow: 0 8px 24px -8px rgba(0, 0, 0, 0.4);
}
.wb-toast-enter-active,
.wb-toast-leave-active {
  transition: opacity 0.25s ease, transform 0.25s ease;
}
.wb-toast-enter-from,
.wb-toast-leave-to {
  opacity: 0;
  transform: translateX(-50%) translateY(8px);
}

/* 防截图水印：斜向平铺，半透明，不拦截点击 */
.wm-layer {
  position: fixed;
  inset: 0;
  z-index: 45;
  display: flex;
  flex-wrap: wrap;
  align-content: center;
  justify-content: center;
  gap: 42px 34px;
  padding: 24px;
  pointer-events: none;
  overflow: hidden;
  opacity: 0.5;
}
.wm-item {
  transform: rotate(-24deg);
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.14em;
  white-space: nowrap;
  color: rgba(20, 24, 29, 0.16);
  user-select: none;
}

/* 防截图水印：斜向平铺，半透明，不拦截点击 */
.wm-layer {
  position: fixed;
  inset: 0;
  z-index: 45;
  display: flex;
  flex-wrap: wrap;
  align-content: center;
  justify-content: center;
  gap: 42px 34px;
  padding: 24px;
  pointer-events: none;
  overflow: hidden;
  opacity: 0.5;
}
.wm-item {
  transform: rotate(-24deg);
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.14em;
  white-space: nowrap;
  color: rgba(20, 24, 29, 0.16);
  user-select: none;
}
</style>

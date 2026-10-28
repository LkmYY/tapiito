<template>
  <div class="min-h-screen" :class="ambientClass">
    <div
      class="flex min-h-screen flex-col"
      :class="isLandscape ? 'w-full bg-black' : 'mx-auto max-w-md bg-white shadow-[0_0_0_1px_rgba(148,163,184,0.14),0_24px_64px_-20px_rgba(2,6,23,0.38)]'"
    >
      <component :is="comp" v-if="comp" :payload="payload" :nfc-id="nfcId" />
      <div v-else class="flex flex-1 items-center justify-center p-8 text-center">
        <p class="text-sm text-muted-foreground">选择模板类型后开始预览</p>
      </div>
      <PlatformFooter v-if="!isLandscape" :tone="footerTone" />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute } from 'vue-router'
import { componentFor, ambientFor, sampleFor, layoutFor } from '../templates/registry.js'
import PlatformFooter from '../components/PlatformFooter.vue'

const route = useRoute()
const type = ref(route.query.type || '')
const style = ref(route.query.style || '')
const payload = ref({})
const nfcId = ref(route.query.nfcId || 'PREVIEW')

const comp = computed(() => componentFor(type.value, style.value))
const ambientClass = computed(() => ambientFor(type.value, style.value))
/** 横屏类型（奇怪的想法）需要全宽容器 */
const isLandscape = computed(() => layoutFor(type.value) === 'landscape')
const footerTone = computed(() => (ambientClass.value || '').startsWith('amb-dark') ? 'dark' : 'light')

// 首次打开（独立窗口）时用示例数据填充，iframe 里随后由父页面 postMessage 覆盖
if (route.query.type && Object.keys(payload.value).length === 0) {
  payload.value = sampleFor(type.value, style.value)
}

const onMessage = (e) => {
  const data = e.data
  if (!data || data.__nfcPreview !== true) return
  type.value = data.type || ''
  style.value = data.style || ''
  payload.value = data.payload || {}
  if (data.nfcId) nfcId.value = data.nfcId
}

onMounted(() => {
  window.addEventListener('message', onMessage)
  // 通知父页面：已就绪
  window.parent?.postMessage({ __nfcPreviewReady: true }, '*')
})

onBeforeUnmount(() => window.removeEventListener('message', onMessage))
</script>

<style scoped>
.amb-plain { background-color: #f4f4f3; }
.amb-cool { background-color: #eef1f6; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(37, 99, 235, 0.07), transparent 62%); }
.amb-warm { background-color: #f5efe6; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(217, 119, 6, 0.07), transparent 62%); }
.amb-warm-pink { background-color: #f7eef1; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(219, 39, 119, 0.06), transparent 62%); }
.amb-warm-green { background-color: #eef2e9; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(22, 163, 74, 0.07), transparent 62%); }
.amb-dark { background-color: #0a0d14; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(96, 165, 250, 0.09), transparent 62%); }
.amb-dark-red { background-color: #0f0a0b; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(248, 113, 113, 0.1), transparent 62%); }
.amb-dark-gold { background-color: #0e0c09; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(251, 191, 36, 0.09), transparent 62%); }
.amb-dark-green { background-color: #080d0a; background-image: radial-gradient(1100px 520px at 50% -8%, rgba(74, 222, 128, 0.08), transparent 62%); }
</style>

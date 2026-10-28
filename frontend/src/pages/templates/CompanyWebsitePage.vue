<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f5f6f8] text-slate-900">
    <!-- Ambient tint -->
    <div class="pointer-events-none absolute inset-x-0 -top-24 h-72 bg-[radial-gradient(60%_100%_at_50%_0%,rgba(29,78,216,0.07),transparent_72%)]"></div>

    <div class="relative z-10 flex min-h-screen flex-col px-7 pb-8 pt-10">
      <!-- Top bar -->
      <header class="rise flex items-center gap-3 border-b border-slate-900/10 pb-4">
        <div class="flex h-9 w-9 flex-shrink-0 items-center justify-center rounded-lg bg-[#1d4ed8] text-sm font-bold text-white">
          {{ initial }}
        </div>
        <div class="min-w-0 flex-1">
          <p class="truncate text-[13px] font-semibold tracking-wide">{{ payload?.name || '企业官网' }}</p>
          <p class="font-mono text-[9px] uppercase tracking-[0.22em] text-slate-400">Official Site</p>
        </div>
        <span v-if="nfcId" class="font-mono text-[10px] uppercase tracking-[0.18em] text-slate-400">NFC · {{ nfcId }}</span>
      </header>

      <!-- Hero -->
      <section class="rise rise-1 mt-10">
        <div class="flex items-center gap-3">
          <span class="h-px w-8 bg-[#1d4ed8]"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.28em] text-[#1d4ed8]">Company Profile</span>
        </div>
        <h1 class="mt-4 text-balance text-[34px] font-semibold leading-[1.1] tracking-tight">{{ payload?.name || '企业官网' }}</h1>
        <p v-if="payload?.description" class="mt-4 text-[15px] leading-relaxed text-slate-500">{{ payload.description }}</p>
      </section>

      <!-- Features -->
      <section v-if="features.length" class="rise rise-2 mt-10">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[13px] font-semibold tracking-wide">核心业务</h2>
          <span class="font-mono text-[10px] tracking-[0.18em] text-slate-400 tabular-nums">{{ pad2(features.length) }} Items</span>
        </div>
        <ul class="mt-2 border-t border-slate-900/10">
          <li v-for="(f, i) in features" :key="i" class="cp-row group">
            <span class="w-8 flex-shrink-0 font-mono text-[11px] tracking-wide text-[#1d4ed8] tabular-nums">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[15px] font-medium text-slate-700 transition-transform duration-200 group-hover:translate-x-0.5">{{ f }}</span>
            <span class="h-1.5 w-1.5 flex-shrink-0 rounded-full bg-slate-300 transition-colors duration-200 group-hover:bg-[#1d4ed8]"></span>
          </li>
        </ul>
      </section>

      <!-- Contact -->
      <section v-if="hasContact" class="rise rise-3 mt-10">
        <div class="rounded-2xl bg-white p-6 shadow-[0_18px_40px_-26px_rgba(29,78,216,0.5)] ring-1 ring-slate-900/[0.06]">
          <div class="flex items-baseline justify-between">
            <h2 class="text-[13px] font-semibold tracking-wide">联系方式</h2>
            <span class="font-mono text-[10px] uppercase tracking-[0.18em] text-slate-400">Contact</span>
          </div>
          <div class="mt-3">
            <a v-if="payload?.contactInfo?.phone" :href="'tel:' + payload.contactInfo.phone" class="cp-contact">
              <span class="cp-ic"><Phone class="h-4 w-4" /></span>
              <span class="cp-label">电话</span>
              <span class="cp-value tabular-nums">{{ payload.contactInfo.phone }}</span>
            </a>
            <a v-if="payload?.contactInfo?.email" :href="'mailto:' + payload.contactInfo.email" class="cp-contact">
              <span class="cp-ic"><Mail class="h-4 w-4" /></span>
              <span class="cp-label">邮箱</span>
              <span class="cp-value">{{ payload.contactInfo.email }}</span>
            </a>
            <div v-if="payload?.contactInfo?.address" class="cp-contact">
              <span class="cp-ic"><MapPin class="h-4 w-4" /></span>
              <span class="cp-label">地址</span>
              <span class="cp-value">{{ payload.contactInfo.address }}</span>
            </div>
          </div>
        </div>
      </section>

      <!-- CTA + footer -->
      <div class="rise rise-3 mt-auto pt-10">
        <a
          v-if="payload?.url && isSafeUrl"
          :href="payload.url"
          target="_blank"
          rel="noopener noreferrer"
          class="flex w-full items-center justify-center gap-2 rounded-xl bg-[#1d4ed8] py-3.5 text-sm font-semibold text-white shadow-[0_16px_32px_-16px_rgba(29,78,216,0.7)] transition-all hover:bg-[#1a43bd] active:scale-[0.98]"
        >
          访问官方网站
          <ArrowUpRight class="h-4 w-4" />
        </a>
        <p class="mt-3 text-center font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Tap to open · New tab</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, MapPin, ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '企').toString().trim().charAt(0) || '企')
const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.url || ''))
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const hasContact = computed(() => {
  const c = props.payload?.contactInfo || {}
  return Boolean(c.phone || c.email || c.address)
})
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.cp-row {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 15px 0;
  border-bottom: 1px solid rgba(15, 23, 42, 0.08);
}
.cp-contact {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid rgba(15, 23, 42, 0.06);
}
.cp-contact:last-child {
  border-bottom: 0;
}
.cp-ic {
  display: flex;
  height: 32px;
  width: 32px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: rgba(29, 78, 216, 0.08);
  color: #1d4ed8;
}
.cp-label {
  width: 36px;
  flex-shrink: 0;
  font-size: 10px;
  letter-spacing: 0.16em;
  color: #94a3b8;
}
.cp-value {
  margin-left: auto;
  max-width: 64%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 14px;
  font-weight: 500;
  color: #1e293b;
}
a.cp-contact:hover .cp-value {
  color: #1d4ed8;
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.rise {
  animation: rise 0.65s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.rise-1 {
  animation-delay: 0.08s;
}
.rise-2 {
  animation-delay: 0.16s;
}
.rise-3 {
  animation-delay: 0.24s;
}
</style>

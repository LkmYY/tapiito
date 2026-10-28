<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0b0d12] text-white">
    <!-- Ambient light -->
    <div class="pointer-events-none absolute inset-x-0 -top-40 h-80 bg-[radial-gradient(60%_100%_at_50%_0%,rgba(59,130,246,0.20),transparent_72%)]"></div>
    <!-- Film grain -->
    <div class="grain pointer-events-none absolute inset-0 opacity-[0.30]"></div>

    <div class="relative z-10 flex min-h-screen flex-col px-7 pb-8 pt-12">
      <!-- Header meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-white/35">
        <span>Digital Card</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Name block -->
      <div class="rise rise-1 mt-14">
        <h1 class="text-balance text-[34px] font-semibold leading-[1.08] tracking-tight">{{ payload.name || '未命名' }}</h1>
        <p v-if="payload.title" class="mt-3 text-sm font-medium tracking-wide text-blue-400/90">{{ payload.title }}</p>
        <p v-if="payload.company" class="mt-1.5 text-sm text-white/45">{{ payload.company }}</p>
      </div>

      <div class="rise rise-2 mt-10 h-px w-full bg-gradient-to-r from-white/25 via-white/10 to-transparent"></div>

      <!-- Contact rows -->
      <div class="rise rise-2 flex-1">
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-row group">
          <Phone class="h-4 w-4 text-white/35 transition-colors group-hover:text-blue-400" />
          <span class="flex-1 text-[15px] tracking-wide tabular-nums">{{ payload.phone }}</span>
          <ArrowUpRight class="h-3.5 w-3.5 text-white/20 opacity-0 transition-opacity group-hover:opacity-100" />
        </a>

        <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-row group">
          <Mail class="h-4 w-4 text-white/35 transition-colors group-hover:text-blue-400" />
          <span class="flex-1 truncate text-[15px] tracking-wide">{{ payload.email }}</span>
          <ArrowUpRight class="h-3.5 w-3.5 text-white/20 opacity-0 transition-opacity group-hover:opacity-100" />
        </a>

        <a
          v-if="payload.website && isSafeUrl"
          :href="payload.website"
          target="_blank"
          rel="noopener noreferrer"
          class="bc-row group"
        >
          <Globe class="h-4 w-4 text-white/35 transition-colors group-hover:text-blue-400" />
          <span class="flex-1 truncate text-[15px] tracking-wide">{{ payload.website }}</span>
          <ArrowUpRight class="h-3.5 w-3.5 text-white/20 opacity-0 transition-opacity group-hover:opacity-100" />
        </a>

        <div v-if="payload.wechat" class="bc-row group">
          <MessageCircle class="h-4 w-4 text-white/35 transition-colors group-hover:text-blue-400" />
          <span class="flex-1 text-[15px] tracking-wide">{{ payload.wechat }}</span>
          <span class="font-mono text-[10px] uppercase tracking-[0.18em] text-white/20">WeChat</span>
        </div>
      </div>

      <!-- Save -->
      <div class="rise rise-3 mt-10 space-y-3">
        <button
          @click="saveToContacts"
          class="w-full rounded-xl bg-white py-3.5 text-sm font-semibold text-[#0b0d12] transition-all hover:bg-blue-50 active:scale-[0.98]"
        >
          保存到通讯录
        </button>
        <p class="text-center font-mono text-[10px] uppercase tracking-[0.24em] text-white/25">Tap to save · vCard</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, Globe, MessageCircle, ArrowUpRight } from 'lucide-vue-next'
import { saveContact } from '../../utils/vcard'

const props = defineProps(['payload', 'nfcId'])

const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.website || ''))

const escapeVCardText = (value) => {
  if (value === undefined || value === null) return ''
  return String(value)
    .replace(/\\/g, '\\\\')
    .replace(/\n/g, '\\n')
    .replace(/\r/g, '')
    .replace(/;/g, '\\;')
    .replace(/,/g, '\\,')
}

const safeFilename = (value) => {
  const base = (value || 'contact').toString().trim() || 'contact'
  return base.replace(/[\\/:*?"<>|]+/g, '_')
}

const buildVCard = () => {
  const p = props.payload || {}
  const name = escapeVCardText(p.name || '')
  const title = escapeVCardText(p.title || '')
  const company = escapeVCardText(p.company || '')
  const phone = escapeVCardText((p.phone || '').toString().trim())
  const email = escapeVCardText(p.email || '')
  const website = escapeVCardText(p.website || '')
  const wechat = escapeVCardText(p.wechat || '')

  const n = name ? `N:${name};;;;` : 'N:;;;;'
  const fn = name ? `FN:${name}` : 'FN:'
  const lines = [
    'BEGIN:VCARD',
    'VERSION:3.0',
    n,
    fn,
    company ? `ORG:${company}` : '',
    title ? `TITLE:${title}` : '',
    phone ? `TEL;TYPE=CELL:${phone}` : '',
    email ? `EMAIL;TYPE=INTERNET:${email}` : '',
    website ? `URL:${website}` : '',
    wechat ? `NOTE:WeChat: ${wechat}` : '',
    'END:VCARD'
  ]

  return lines.filter(Boolean).join('\r\n') + '\r\n'
}

const saveToContacts = async () => {
  // 统一实现：优先调起系统分享面板直接存进通讯录，不支持时才下载 .vcf
  const res = await saveContact(props.payload || {})
  if (res.reason === 'wechat') {
    window.alert('微信内无法直接保存，请点右上角「在浏览器中打开」后再试')
  }
}
</script>

<style scoped>
.grain {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='140' height='140'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='2' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.45'/%3E%3C/svg%3E");
}

.bc-row {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 16px 2px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.07);
  transition: background-color 0.2s ease;
}
.bc-row:hover {
  background-color: rgba(255, 255, 255, 0.03);
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

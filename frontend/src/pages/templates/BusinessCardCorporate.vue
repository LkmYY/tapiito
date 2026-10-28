<template>
  <div class="min-h-screen bg-[#f4f6fa]">
    <!-- Brand header -->
    <div class="bg-[#123a8f] px-7 pb-20 pt-12 text-white">
      <div class="rise flex items-center gap-4">
        <div class="flex h-12 w-12 items-center justify-center rounded-lg bg-white/10 text-lg font-bold ring-1 ring-white/20">
          {{ initial }}
        </div>
        <div class="min-w-0">
          <p class="truncate text-[15px] font-semibold tracking-wide">{{ payload.company || '企业名片' }}</p>
          <p class="mt-0.5 font-mono text-[10px] uppercase tracking-[0.22em] text-blue-200/70">Corporate Identity</p>
        </div>
      </div>
    </div>

    <!-- Main card -->
    <div class="relative z-10 -mt-12 px-6">
      <div class="rise rise-1 rounded-2xl bg-white p-7 shadow-[0_18px_40px_-18px_rgba(18,58,143,0.35)] ring-1 ring-[#123a8f]/5">
        <h1 class="text-balance text-[26px] font-bold tracking-tight text-slate-900">{{ payload.name || '未命名' }}</h1>
        <p v-if="payload.title" class="mt-1 text-sm font-medium text-[#123a8f]">{{ payload.title }}</p>

        <div class="mt-7 border-t border-slate-100">
          <div v-if="payload.phone" class="c-row">
            <Phone class="h-4 w-4 flex-shrink-0 text-slate-400" />
            <span class="c-label">电话</span>
            <a :href="'tel:' + payload.phone" class="c-value tabular-nums">{{ payload.phone }}</a>
          </div>

          <div v-if="payload.email" class="c-row">
            <Mail class="h-4 w-4 flex-shrink-0 text-slate-400" />
            <span class="c-label">邮箱</span>
            <a :href="'mailto:' + payload.email" class="c-value">{{ payload.email }}</a>
          </div>

          <div v-if="payload.website && isSafeUrl" class="c-row">
            <Globe class="h-4 w-4 flex-shrink-0 text-slate-400" />
            <span class="c-label">网站</span>
            <a :href="payload.website" target="_blank" rel="noopener noreferrer" class="c-value">{{ payload.website }}</a>
          </div>

          <div v-if="payload.wechat" class="c-row">
            <MessageCircle class="h-4 w-4 flex-shrink-0 text-slate-400" />
            <span class="c-label">微信</span>
            <span class="c-value">{{ payload.wechat }}</span>
          </div>
        </div>
      </div>

      <button
        @click="saveToContacts"
        class="rise rise-2 mt-6 w-full rounded-xl bg-[#123a8f] py-3.5 text-sm font-semibold text-white shadow-lg shadow-[#123a8f]/20 transition-all hover:bg-[#0f2f74] active:scale-[0.98]"
      >
        保存到通讯录
      </button>
    </div>

    <!-- Footer -->
    <div class="mt-10 border-t border-slate-200 bg-white py-5 text-center">
      <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">{{ payload.company }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, Globe, MessageCircle } from 'lucide-vue-next'
import { saveContact } from '../../utils/vcard'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.company || props.payload?.name || '?').toString().trim()[0] || '?')
const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.website || ''))

const saveToContacts = async () => {
  // 统一实现：优先调起系统分享面板直接存进通讯录，不支持时才下载 .vcf
  const res = await saveContact(props.payload || {})
  if (res.reason === 'wechat') {
    window.alert('微信内无法直接保存，请点右上角「在浏览器中打开」后再试')
  }
}
</script>

<style scoped>
.c-row {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 0;
  border-bottom: 1px solid #f1f5f9;
}
.c-label {
  width: 56px;
  flex-shrink: 0;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 0.18em;
  color: #94a3b8;
}
.c-value {
  margin-left: auto;
  max-width: 62%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 14px;
  font-weight: 500;
  color: #1e293b;
}
a.c-value:hover {
  color: #123a8f;
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
  animation-delay: 0.1s;
}
.rise-2 {
  animation-delay: 0.2s;
}
</style>

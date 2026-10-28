<template>
  <div class="min-h-screen bg-white px-8 py-14 text-stone-900">
    <div class="mx-auto flex min-h-[80vh] w-full max-w-sm flex-col">
      <!-- Meta -->
      <p class="rise font-mono text-[10px] uppercase tracking-[0.3em] text-stone-400">Business Card</p>

      <!-- Name -->
      <div class="rise rise-1 mt-16">
        <h1 class="text-balance font-serif text-[44px] leading-[1.05] tracking-tight">{{ payload.name || '未命名' }}</h1>
        <p v-if="payload.title" class="mt-4 text-[15px] text-stone-500">{{ payload.title }}</p>
        <p v-if="payload.company" class="mt-1.5 text-[12px] uppercase tracking-[0.18em] text-stone-400">{{ payload.company }}</p>
      </div>

      <!-- Fields -->
      <div class="rise rise-2 mt-16 flex-1 border-t border-stone-200">
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bm-row group">
          <span class="bm-label">Tel</span>
          <span class="bm-value tabular-nums transition-transform group-hover:translate-x-0.5">{{ payload.phone }}</span>
        </a>

        <a v-if="payload.email" :href="'mailto:' + payload.email" class="bm-row group">
          <span class="bm-label">Mail</span>
          <span class="bm-value truncate transition-transform group-hover:translate-x-0.5">{{ payload.email }}</span>
        </a>

        <a
          v-if="payload.website && isSafeUrl"
          :href="payload.website"
          target="_blank"
          rel="noopener noreferrer"
          class="bm-row group"
        >
          <span class="bm-label">Web</span>
          <span class="bm-value truncate transition-transform group-hover:translate-x-0.5">{{ payload.website }}</span>
        </a>

        <div v-if="payload.wechat" class="bm-row group">
          <span class="bm-label">WX</span>
          <span class="bm-value">{{ payload.wechat }}</span>
        </div>
      </div>

      <!-- Save -->
      <button
        @click="saveToContacts"
        class="rise rise-3 mt-12 w-full border border-stone-900 py-3.5 text-[11px] font-semibold uppercase tracking-[0.22em] transition-colors hover:bg-stone-900 hover:text-white active:scale-[0.99]"
      >
        Save Contact
      </button>
      <p class="mt-4 text-center font-mono text-[10px] tracking-[0.2em] text-stone-300">VCARD · .VCF</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { saveContact } from '../../utils/vcard'

const props = defineProps(['payload', 'nfcId'])
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
.bm-row {
  display: flex;
  align-items: baseline;
  gap: 24px;
  padding: 17px 0;
  border-bottom: 1px solid #e7e5e4;
}
.bm-label {
  width: 40px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  color: #a8a29e;
}
.bm-value {
  font-size: 14px;
  color: #1c1917;
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

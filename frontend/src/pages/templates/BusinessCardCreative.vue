<template>
  <div class="relative min-h-screen overflow-hidden bg-[#17130f] text-white">
    <!-- Orange block with angled bottom -->
    <div class="absolute inset-x-0 top-0 h-[54%] bg-[#ff5a1f]" style="clip-path: polygon(0 0, 100% 0, 100% 86%, 0 100%)"></div>
    <!-- Dot grid texture on orange -->
    <div
      class="absolute inset-x-0 top-0 h-[54%] opacity-[0.10]"
      style="
        clip-path: polygon(0 0, 100% 0, 100% 86%, 0 100%);
        background-image: radial-gradient(#17130f 1px, transparent 1px);
        background-size: 12px 12px;
      "
    ></div>

    <div class="relative z-10 flex min-h-screen flex-col px-7 pb-8 pt-12">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.26em] text-[#17130f]/55">
        <span>Creative Card</span>
        <span class="max-w-[45%] truncate text-right">{{ payload.company }}</span>
      </div>

      <!-- Big name -->
      <div class="rise rise-1 mt-12">
        <h1 class="text-balance text-[46px] font-black uppercase leading-[0.98] tracking-tight text-[#1a120c]">
          {{ payload.name || '未命名' }}
        </h1>
        <p v-if="payload.title" class="mt-3 text-sm font-bold text-[#1a120c]/60">{{ payload.title }}</p>
      </div>

      <!-- Vertical accent -->
      <div
        class="pointer-events-none absolute right-6 top-[34%] font-mono text-[11px] uppercase tracking-[0.5em] text-white/25 [writing-mode:vertical-rl]"
      >
        Contact
      </div>

      <!-- Contact list -->
      <div class="rise rise-2 mt-auto pt-28">
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="cc-row group">
          <span class="cc-icon"><Phone class="h-4 w-4" /></span>
          <span class="flex-1 text-[15px] font-medium tabular-nums">{{ payload.phone }}</span>
          <ArrowUpRight class="h-4 w-4 text-white/20 transition-opacity group-hover:text-[#ff5a1f]" />
        </a>

        <a v-if="payload.email" :href="'mailto:' + payload.email" class="cc-row group">
          <span class="cc-icon"><Mail class="h-4 w-4" /></span>
          <span class="flex-1 truncate text-[15px] font-medium">{{ payload.email }}</span>
          <ArrowUpRight class="h-4 w-4 text-white/20 transition-opacity group-hover:text-[#ff5a1f]" />
        </a>

        <a
          v-if="payload.website && isSafeUrl"
          :href="payload.website"
          target="_blank"
          rel="noopener noreferrer"
          class="cc-row group"
        >
          <span class="cc-icon"><Globe class="h-4 w-4" /></span>
          <span class="flex-1 truncate text-[15px] font-medium">{{ payload.website }}</span>
          <ArrowUpRight class="h-4 w-4 text-white/20 transition-opacity group-hover:text-[#ff5a1f]" />
        </a>

        <div v-if="payload.wechat" class="cc-row group">
          <span class="cc-icon"><MessageCircle class="h-4 w-4" /></span>
          <span class="flex-1 text-[15px] font-medium">{{ payload.wechat }}</span>
          <span class="font-mono text-[10px] uppercase tracking-[0.18em] text-white/25">WeChat</span>
        </div>
      </div>

      <!-- Save -->
      <button
        @click="saveToContacts"
        class="rise rise-3 mt-8 w-full rounded-lg bg-[#ff5a1f] py-3.5 text-sm font-bold uppercase tracking-wider text-[#1a120c] transition-all hover:brightness-110 active:scale-[0.98]"
      >
        保存联系人
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, Globe, MessageCircle, ArrowUpRight } from 'lucide-vue-next'
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
.cc-row {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  transition: border-color 0.2s ease;
}
.cc-row:hover {
  border-color: rgba(255, 90, 31, 0.35);
}
.cc-icon {
  display: flex;
  height: 32px;
  width: 32px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: rgba(255, 90, 31, 0.15);
  color: #ff5a1f;
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

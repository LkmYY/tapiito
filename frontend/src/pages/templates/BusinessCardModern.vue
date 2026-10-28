<template>
  <div class="relative min-h-screen overflow-hidden bg-[#fafaf9] px-6 py-10 text-stone-900">
    <!-- Soft light -->
    <div class="pointer-events-none absolute -right-24 -top-24 h-72 w-72 rounded-full bg-indigo-200/45 blur-3xl"></div>
    <div class="pointer-events-none absolute -left-28 top-1/3 h-64 w-64 rounded-full bg-amber-100/60 blur-3xl"></div>

    <div class="relative z-10 mx-auto flex min-h-[80vh] w-full max-w-sm flex-col">
      <!-- Monogram -->
      <div class="rise flex items-start justify-between">
        <div class="flex h-16 w-16 items-center justify-center rounded-2xl bg-stone-900 text-2xl font-semibold text-white shadow-lg shadow-stone-900/10">
          {{ initial }}
        </div>
        <p class="text-right text-[10px] font-medium uppercase leading-relaxed tracking-[0.18em] text-stone-400">
          Digital<br />Business Card
        </p>
      </div>

      <!-- Name -->
      <div class="rise rise-1 mt-9">
        <h1 class="text-balance text-[30px] font-semibold leading-tight tracking-tight">{{ payload.name || '未命名' }}</h1>
        <p class="mt-2.5 text-[15px] text-stone-500">
          {{ payload.title }}<template v-if="payload.title && payload.company"> · </template>{{ payload.company }}
        </p>
      </div>

      <!-- Contact grid -->
      <div class="rise rise-2 mt-10 grid flex-1 grid-cols-2 content-start gap-3">
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="m-card group">
          <Phone class="h-4 w-4 text-indigo-500" />
          <p class="m-label">电话</p>
          <p class="m-value tabular-nums">{{ payload.phone }}</p>
        </a>

        <a v-if="payload.email" :href="'mailto:' + payload.email" class="m-card group">
          <Mail class="h-4 w-4 text-indigo-500" />
          <p class="m-label">邮箱</p>
          <p class="m-value">{{ payload.email }}</p>
        </a>

        <a
          v-if="payload.website && isSafeUrl"
          :href="payload.website"
          target="_blank"
          rel="noopener noreferrer"
          class="m-card group"
        >
          <Globe class="h-4 w-4 text-indigo-500" />
          <p class="m-label">网站</p>
          <p class="m-value">{{ payload.website }}</p>
        </a>

        <div v-if="payload.wechat" class="m-card group">
          <MessageCircle class="h-4 w-4 text-indigo-500" />
          <p class="m-label">微信</p>
          <p class="m-value">{{ payload.wechat }}</p>
        </div>
      </div>

      <!-- Save -->
      <button
        @click="saveToContacts"
        class="rise rise-3 mt-8 w-full rounded-2xl bg-stone-900 py-4 text-sm font-semibold text-white transition-all hover:bg-stone-800 active:scale-[0.98]"
      >
        保存到通讯录
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, Globe, MessageCircle } from 'lucide-vue-next'
import { saveContact } from '../../utils/vcard'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
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
.m-card {
  display: flex;
  flex-direction: column;
  gap: 2px;
  border-radius: 16px;
  border: 1px solid rgba(28, 25, 23, 0.07);
  background: #ffffff;
  padding: 16px;
  transition: border-color 0.2s ease, box-shadow 0.2s ease, transform 0.2s ease;
}
.m-card:hover {
  border-color: rgba(99, 102, 241, 0.35);
  box-shadow: 0 8px 20px -12px rgba(28, 25, 23, 0.18);
  transform: translateY(-1px);
}
.m-label {
  margin-top: 12px;
  font-size: 11px;
  color: #a8a29e;
}
.m-value {
  font-size: 13px;
  font-weight: 500;
  color: #1c1917;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
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

<template>
  <div class="min-h-screen bg-[#f0ece4] px-5 py-10 text-stone-900">
    <div class="mx-auto flex min-h-[85vh] w-full max-w-sm flex-col border border-stone-900/[0.12] bg-[#faf8f3] px-7 py-12">
      <!-- Masthead -->
      <div class="rise flex items-center gap-3">
        <span class="h-px flex-1 bg-stone-900/15"></span>
        <span class="font-mono text-[9px] uppercase tracking-[0.42em] text-stone-500">Company</span>
        <span class="h-px flex-1 bg-stone-900/15"></span>
      </div>

      <!-- Name + ornament -->
      <div class="rise rise-1 mt-14 text-center">
        <h1 class="text-balance font-serif text-[40px] leading-[1.12] tracking-tight">{{ payload?.name || '企业官网' }}</h1>
        <div class="mt-7 flex items-center justify-center gap-3">
          <span class="h-px w-12 bg-[#a16207]/50"></span>
          <span class="h-1.5 w-1.5 rotate-45 bg-[#a16207]"></span>
          <span class="h-px w-12 bg-[#a16207]/50"></span>
        </div>
        <p v-if="payload?.description" class="mt-7 text-[14.5px] leading-[1.9] tracking-[0.01em] text-stone-500">{{ payload.description }}</p>
      </div>

      <!-- Features -->
      <div v-if="features.length" class="rise rise-2 mt-14">
        <p class="text-center font-mono text-[9px] uppercase tracking-[0.42em] text-[#a16207]">Core Services</p>
        <ul class="mt-6 border-t border-stone-900/10">
          <li v-for="(f, i) in features" :key="i" class="el-item">
            <span class="font-serif text-[12px] italic tracking-[0.2em] text-[#a16207]">{{ roman(i + 1) }}</span>
            <span class="mt-2 font-serif text-[18px] leading-snug tracking-wide text-stone-800">{{ f }}</span>
          </li>
        </ul>
      </div>

      <!-- CTA -->
      <div class="rise rise-3 mt-12 text-center">
        <a
          v-if="payload?.url && isSafeUrl"
          :href="payload.url"
          target="_blank"
          rel="noopener noreferrer"
          class="inline-flex items-center gap-2.5 border border-stone-900/80 px-9 py-3.5 text-[10.5px] font-medium uppercase tracking-[0.3em] text-stone-900 transition-colors duration-300 hover:bg-stone-900 hover:text-[#faf8f3] active:scale-[0.99]"
        >
          访问官网
          <ArrowUpRight class="h-3.5 w-3.5" />
        </a>
      </div>

      <!-- Contact -->
      <div v-if="hasContact" class="rise rise-3 mt-12">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-stone-900/10"></span>
          <span class="font-mono text-[9px] uppercase tracking-[0.42em] text-stone-500">Contact</span>
          <span class="h-px flex-1 bg-stone-900/10"></span>
        </div>
        <div class="mt-4">
          <a v-if="payload?.contactInfo?.phone" :href="'tel:' + payload.contactInfo.phone" class="el-contact">
            <span class="el-label">Telephone</span>
            <span class="el-value tabular-nums">{{ payload.contactInfo.phone }}</span>
          </a>
          <a v-if="payload?.contactInfo?.email" :href="'mailto:' + payload.contactInfo.email" class="el-contact">
            <span class="el-label">Email</span>
            <span class="el-value">{{ payload.contactInfo.email }}</span>
          </a>
          <div v-if="payload?.contactInfo?.address" class="el-contact">
            <span class="el-label">Address</span>
            <span class="el-value">{{ payload.contactInfo.address }}</span>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <p v-if="nfcId" class="mt-auto pt-12 text-center font-mono text-[9px] uppercase tracking-[0.42em] text-stone-400">NFC · {{ nfcId }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.url || ''))
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const hasContact = computed(() => {
  const c = props.payload?.contactInfo || {}
  return Boolean(c.phone || c.email || c.address)
})

const ROMAN = ['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX', 'X', 'XI', 'XII']
const roman = (n) => ROMAN[n - 1] || String(n)
</script>

<style scoped>
.el-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 22px 0;
  border-bottom: 1px solid rgba(28, 25, 23, 0.1);
}
.el-contact {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 16px;
  padding: 13px 0;
  border-bottom: 1px solid rgba(28, 25, 23, 0.08);
}
.el-contact:last-child {
  border-bottom: 0;
}
.el-label {
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  text-transform: uppercase;
  letter-spacing: 0.3em;
  color: #a8a29e;
}
.el-value {
  max-width: 62%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: right;
  font-size: 13.5px;
  letter-spacing: 0.02em;
  color: #57534e;
}
a.el-contact:hover .el-value {
  color: #a16207;
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

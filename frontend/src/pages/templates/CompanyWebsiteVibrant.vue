<template>
  <div class="min-h-screen bg-[#fffdf6] text-[#111111]">
    <!-- Top bar -->
    <header class="rise flex items-center justify-between gap-3 border-b-2 border-[#111111] px-6 py-3.5">
      <div class="flex min-w-0 items-center gap-3">
        <span class="flex h-8 w-8 flex-shrink-0 items-center justify-center bg-[#111111] text-[15px] font-black text-[#ffd60a]">{{ initial }}</span>
        <span class="truncate text-[14px] font-black uppercase tracking-tight">{{ payload?.name || '企业官网' }}</span>
      </div>
      <span v-if="nfcId" class="flex-shrink-0 font-mono text-[9px] uppercase tracking-[0.2em] text-[#111111]/40">NFC·{{ nfcId }}</span>
    </header>

    <!-- Hero -->
    <section class="rise rise-1 relative overflow-hidden border-b-2 border-[#111111] bg-[#ffd60a] px-6 pb-11 pt-8">
      <div class="dots pointer-events-none absolute inset-0 opacity-[0.07]"></div>
      <div class="relative">
        <span class="inline-block -rotate-2 rounded-full border-2 border-[#111111] bg-[#fffdf6] px-3.5 py-1 font-mono text-[9px] font-bold uppercase tracking-[0.24em]">Official Site</span>
        <h1 class="mt-5 text-balance text-[40px] font-black uppercase leading-[0.96] tracking-tight">{{ payload?.name || '企业官网' }}</h1>
        <p v-if="payload?.description" class="mt-4 max-w-[34ch] text-[14.5px] font-medium leading-relaxed text-[#111111]/70">{{ payload.description }}</p>
      </div>
    </section>

    <!-- Features -->
    <section v-if="features.length" class="rise rise-2 px-6 pt-8">
      <div class="flex items-baseline justify-between">
        <h2 class="text-[12px] font-black uppercase tracking-[0.2em]">核心优势</h2>
        <span class="font-mono text-[10px] tabular-nums text-[#111111]/40">{{ pad2(features.length) }}</span>
      </div>
      <ul class="mt-4">
        <li v-for="(f, i) in features" :key="i" class="vb-row group">
          <span class="flex h-9 w-9 flex-shrink-0 items-center justify-center bg-[#111111] font-mono text-[12px] font-bold text-[#ffd60a] tabular-nums">{{ pad2(i + 1) }}</span>
          <span class="flex-1 text-[16px] font-bold leading-snug transition-transform duration-200 group-hover:translate-x-1">{{ f }}</span>
        </li>
      </ul>
    </section>

    <!-- CTA -->
    <section class="rise rise-3 px-6 pt-7">
      <a
        v-if="payload?.url && isSafeUrl"
        :href="payload.url"
        target="_blank"
        rel="noopener noreferrer"
        class="flex w-full items-center justify-center gap-2.5 border-2 border-[#111111] bg-[#111111] py-4 text-[13px] font-black uppercase tracking-[0.18em] text-[#ffd60a] transition-colors hover:bg-[#ffd60a] hover:text-[#111111] active:scale-[0.99]"
      >
        访问官方网站
        <ArrowUpRight class="h-4 w-4" />
      </a>
      <p class="mt-3 text-center font-mono text-[9px] uppercase tracking-[0.26em] text-[#111111]/35">External Link · New Tab</p>
    </section>

    <!-- Contact -->
    <section v-if="hasContact" class="mt-9 border-y-2 border-[#111111] bg-[#ffd60a] px-6 py-7">
      <h2 class="text-[12px] font-black uppercase tracking-[0.2em]">联系我们</h2>
      <div class="mt-3">
        <a v-if="payload?.contactInfo?.phone" :href="'tel:' + payload.contactInfo.phone" class="vb-contact group">
          <span class="vb-ic"><Phone class="h-3.5 w-3.5" /></span>
          <span class="flex-1 truncate text-[14.5px] font-semibold tabular-nums">{{ payload.contactInfo.phone }}</span>
          <ArrowUpRight class="h-3.5 w-3.5 opacity-30 transition-opacity group-hover:opacity-100" />
        </a>
        <a v-if="payload?.contactInfo?.email" :href="'mailto:' + payload.contactInfo.email" class="vb-contact group">
          <span class="vb-ic"><Mail class="h-3.5 w-3.5" /></span>
          <span class="flex-1 truncate text-[14.5px] font-semibold">{{ payload.contactInfo.email }}</span>
          <ArrowUpRight class="h-3.5 w-3.5 opacity-30 transition-opacity group-hover:opacity-100" />
        </a>
        <div v-if="payload?.contactInfo?.address" class="vb-contact">
          <span class="vb-ic"><MapPin class="h-3.5 w-3.5" /></span>
          <span class="flex-1 text-[14.5px] font-semibold leading-snug">{{ payload.contactInfo.address }}</span>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="bg-[#111111] px-6 py-5 text-center">
      <p class="font-mono text-[9px] uppercase tracking-[0.3em] text-[#ffd60a]/70">
        {{ payload?.name || '企业官网' }}<template v-if="nfcId"> · NFC {{ nfcId }}</template>
      </p>
    </footer>
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
.dots {
  background-image: radial-gradient(#111111 1px, transparent 1px);
  background-size: 12px 12px;
}

.vb-row {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 13px 0;
  border-bottom: 2px solid #111111;
}

.vb-contact {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid rgba(17, 17, 17, 0.18);
}
.vb-contact:last-child {
  border-bottom: 0;
}
.vb-ic {
  display: flex;
  height: 28px;
  width: 28px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  background: #111111;
  color: #ffd60a;
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

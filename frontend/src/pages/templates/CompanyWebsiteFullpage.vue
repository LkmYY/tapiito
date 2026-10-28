<template>
  <div class="fp-scroll relative bg-[#0e0f13]">
    <!-- 01 · Hero -->
    <section class="fp-section relative flex flex-col justify-between overflow-hidden bg-[#0e0f13] px-7 pb-9 pt-10 text-white">
      <div class="pointer-events-none absolute inset-x-0 -top-24 h-72 bg-[radial-gradient(55%_100%_at_50%_0%,rgba(255,255,255,0.07),transparent_70%)]"></div>

      <header class="rise relative flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-white/35">
        <span>Company Profile</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </header>

      <div class="rise rise-1 relative">
        <span class="block h-[3px] w-10 bg-[#ea4a2b]"></span>
        <p class="mt-6 font-mono text-[10px] uppercase tracking-[0.32em] text-white/40">Official Site</p>
        <h1 class="mt-4 text-balance text-[42px] font-semibold leading-[1.04] tracking-tight">{{ payload?.name || '企业官网' }}</h1>
      </div>

      <div class="rise rise-2 relative flex items-center gap-3">
        <span class="flex h-10 w-10 items-center justify-center rounded-full border border-white/15">
          <ChevronDown class="h-4 w-4 animate-bounce text-white/60" />
        </span>
        <span class="font-mono text-[10px] uppercase tracking-[0.28em] text-white/35">向下滚动</span>
      </div>
    </section>

    <!-- 02 · About -->
    <section class="fp-section relative flex flex-col justify-center bg-[#f6f4f0] px-7 py-20 text-[#14110e]">
      <span class="pointer-events-none absolute right-7 top-14 font-mono text-[64px] leading-none text-[#14110e]/[0.05] tabular-nums">01</span>
      <div class="rise relative">
        <div class="flex items-center gap-3 font-mono text-[10px] uppercase tracking-[0.28em]">
          <span class="text-[#ea4a2b] tabular-nums">( 01 )</span>
          <span class="h-px flex-1 bg-[#14110e]/10"></span>
          <span class="text-[#14110e]/45">About</span>
        </div>
        <h2 class="mt-6 text-[30px] font-semibold tracking-tight">关于我们</h2>
        <p v-if="payload?.description" class="mt-5 text-[16px] leading-[1.85] text-[#14110e]/65">{{ payload.description }}</p>
        <p v-else class="mt-5 text-[15px] leading-relaxed text-[#14110e]/40">暂无公司介绍</p>
        <p class="mt-7 font-mono text-[10px] uppercase tracking-[0.28em] text-[#14110e]/40">— {{ payload?.name || '企业官网' }}</p>
      </div>
    </section>

    <!-- 03 · Features -->
    <section v-if="features.length" class="fp-section relative flex flex-col justify-center bg-[#0e0f13] px-7 py-20 text-white">
      <span class="pointer-events-none absolute right-7 top-14 font-mono text-[64px] leading-none text-white/[0.04] tabular-nums">02</span>
      <div class="rise relative">
        <div class="flex items-center gap-3 font-mono text-[10px] uppercase tracking-[0.28em]">
          <span class="text-[#ea4a2b] tabular-nums">( 02 )</span>
          <span class="h-px flex-1 bg-white/10"></span>
          <span class="text-white/40">Core</span>
        </div>
        <h2 class="mt-6 text-[30px] font-semibold tracking-tight">核心优势</h2>
        <ul class="mt-8 border-t border-white/10">
          <li v-for="(f, i) in features" :key="i" class="fp-row group">
            <span class="w-9 flex-shrink-0 font-mono text-[11px] text-white/30 tabular-nums">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[17px] font-medium leading-snug text-white/90">{{ f }}</span>
            <span class="h-1 w-1 flex-shrink-0 rounded-full bg-white/20 transition-colors group-hover:bg-[#ea4a2b]"></span>
          </li>
        </ul>
      </div>
    </section>

    <!-- 04 · Contact -->
    <section class="fp-section relative flex flex-col justify-center bg-[#f6f4f0] px-7 py-20 text-[#14110e]">
      <span class="pointer-events-none absolute right-7 top-14 font-mono text-[64px] leading-none text-[#14110e]/[0.05] tabular-nums">03</span>
      <div class="rise relative">
        <div class="flex items-center gap-3 font-mono text-[10px] uppercase tracking-[0.28em]">
          <span class="text-[#ea4a2b] tabular-nums">( 03 )</span>
          <span class="h-px flex-1 bg-[#14110e]/10"></span>
          <span class="text-[#14110e]/45">Contact</span>
        </div>
        <h2 class="mt-6 text-[30px] font-semibold tracking-tight">联系我们</h2>

        <div v-if="hasContact" class="mt-6">
          <a v-if="payload?.contactInfo?.phone" :href="'tel:' + payload.contactInfo.phone" class="fp-contact">
            <span class="fp-ic"><Phone class="h-4 w-4" /></span>
            <span class="fp-label">电话</span>
            <span class="fp-value tabular-nums">{{ payload.contactInfo.phone }}</span>
          </a>
          <a v-if="payload?.contactInfo?.email" :href="'mailto:' + payload.contactInfo.email" class="fp-contact">
            <span class="fp-ic"><Mail class="h-4 w-4" /></span>
            <span class="fp-label">邮箱</span>
            <span class="fp-value">{{ payload.contactInfo.email }}</span>
          </a>
          <div v-if="payload?.contactInfo?.address" class="fp-contact">
            <span class="fp-ic"><MapPin class="h-4 w-4" /></span>
            <span class="fp-label">地址</span>
            <span class="fp-value">{{ payload.contactInfo.address }}</span>
          </div>
        </div>

        <a
          v-if="payload?.url && isSafeUrl"
          :href="payload.url"
          target="_blank"
          rel="noopener noreferrer"
          class="group mt-9 flex w-full items-center justify-between rounded-xl bg-[#14110e] px-5 py-4 text-white transition-colors hover:bg-[#ea4a2b] active:scale-[0.99]"
        >
          <span class="text-[14px] font-semibold tracking-wide">访问官方网站</span>
          <ArrowUpRight class="h-5 w-5 transition-transform duration-200 group-hover:-translate-y-0.5 group-hover:translate-x-0.5" />
        </a>

        <p v-if="payload?.url && isSafeUrl" class="mt-3 truncate text-center font-mono text-[10px] tracking-[0.14em] text-[#14110e]/40">{{ payload.url }}</p>

        <p class="mt-12 text-center font-mono text-[9px] uppercase tracking-[0.32em] text-[#14110e]/35">
          © {{ payload?.name || 'Company' }}<template v-if="nfcId"> · NFC {{ nfcId }}</template>
        </p>
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ChevronDown, ArrowUpRight, Phone, Mail, MapPin } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.url || ''))
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const hasContact = computed(() => {
  const c = props.payload?.contactInfo || {}
  return Boolean(c.phone || c.email || c.address)
})
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.fp-scroll {
  height: 100vh;
  overflow-y: auto;
  scroll-snap-type: y proximity;
  scroll-behavior: smooth;
  overscroll-behavior: contain;
  -webkit-overflow-scrolling: touch;
}
.fp-section {
  min-height: 100vh;
  scroll-snap-align: start;
}

.fp-row {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 18px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.fp-contact {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 0;
  border-bottom: 1px solid rgba(20, 17, 14, 0.08);
}
.fp-contact:last-child {
  border-bottom: 0;
}
.fp-ic {
  display: flex;
  height: 32px;
  width: 32px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: rgba(20, 17, 14, 0.05);
  color: rgba(20, 17, 14, 0.45);
}
.fp-label {
  width: 36px;
  flex-shrink: 0;
  font-size: 10px;
  letter-spacing: 0.16em;
  color: rgba(20, 17, 14, 0.4);
}
.fp-value {
  margin-left: auto;
  max-width: 62%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 14px;
  font-weight: 500;
  color: #14110e;
}
a.fp-contact:hover .fp-value {
  color: #ea4a2b;
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

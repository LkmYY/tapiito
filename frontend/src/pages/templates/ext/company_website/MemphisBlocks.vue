<template>
  <div class="cw-mem relative min-h-screen overflow-hidden bg-[#fff7e6] text-[#191919]">
    <div class="cw-mem-dot pointer-events-none absolute right-5 top-24 h-20 w-20 rounded-full"></div>
    <svg class="pointer-events-none absolute -left-4 top-7 h-16 w-24 text-[#2f6fdb]" viewBox="0 0 100 60" fill="none" aria-hidden="true">
      <path d="M3 30c10-24 24 24 34 0s24 24 34 0 18 8 24 0" stroke="currentColor" stroke-width="6" stroke-linecap="round" />
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-9 pt-7">
      <header class="cw-mem-in flex items-center justify-between">
        <span class="cw-mem-tag bg-[#191919] text-[#fff7e6]">企业官网</span>
        <span class="font-mono text-[9px] font-bold uppercase tracking-[0.14em] text-[#191919]/55">NFC {{ nfcId || '000' }}</span>
      </header>

      <section class="cw-mem-in cw-mem-d1 mt-7">
        <div class="relative">
          <span class="cw-mem-star absolute -right-1 -top-5 text-[30px] leading-none text-[#f25c3d]">✳</span>
          <h1 class="text-[48px] font-black leading-[0.9] tracking-[-0.055em]">{{ name }}</h1>
        </div>
        <div class="mt-4 flex items-start gap-2.5">
          <span class="mt-1 inline-block h-3 w-3 flex-shrink-0 rounded-full bg-[#2f6fdb]"></span>
          <p class="text-[13.5px] leading-[1.75] text-[#191919]/75">{{ description }}</p>
        </div>
      </section>

      <section class="cw-mem-in cw-mem-d2 mt-7">
        <div class="flex items-center justify-between">
          <h2 class="cw-mem-tag bg-[#f6c945]">服务清单</h2>
          <span class="font-mono text-[9px] font-bold text-[#191919]/50">{{ pad2(features.length) }} 项</span>
        </div>
        <div v-if="features.length" class="mt-3 grid grid-cols-2 gap-2.5">
          <article v-for="(f, i) in features" :key="i" :class="['cw-mem-card', 'cw-mem-c' + (i % 4)]">
            <span class="font-mono text-[9px] font-bold text-[#191919]/45">{{ pad2(i + 1) }}</span>
            <p class="mt-5 text-[13px] font-bold leading-snug">{{ f }}</p>
          </article>
        </div>
      </section>

      <section class="cw-mem-in cw-mem-d2 mt-5">
        <div class="cw-mem-band">
          <div class="text-center">
            <p class="text-[23px] font-black leading-none">{{ pad2(features.length) }}</p>
            <p class="mt-1 font-mono text-[8px] font-bold uppercase tracking-[0.1em]">Services</p>
          </div>
          <span class="h-8 w-0.5 bg-[#191919]"></span>
          <div class="text-center">
            <p class="text-[23px] font-black leading-none">24H</p>
            <p class="mt-1 font-mono text-[8px] font-bold uppercase tracking-[0.1em]">Reply</p>
          </div>
          <span class="h-8 w-0.5 bg-[#191919]"></span>
          <div class="text-center">
            <p class="text-[23px] font-black leading-none">1:1</p>
            <p class="mt-1 font-mono text-[8px] font-bold uppercase tracking-[0.1em]">Service</p>
          </div>
        </div>
      </section>

      <section class="cw-mem-in cw-mem-d3 mt-8">
        <div class="flex items-center gap-2">
          <h2 class="cw-mem-tag bg-[#f25c3d] text-white">联系窗口</h2>
          <span class="h-1 flex-1 bg-[#191919]"></span>
        </div>
        <div class="cw-mem-contactbox mt-3">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-mem-contact">
            <span class="cw-mem-ic bg-[#f6c945]"><Phone class="h-4 w-4" /></span>
            <span class="flex-1">
              <span class="block font-mono text-[7.5px] font-bold uppercase tracking-[0.14em] text-white/45">Call us</span>
              <span class="block text-[13px] font-bold tabular-nums text-white">{{ contact.phone }}</span>
            </span>
            <span class="text-[16px] text-white/50">↗</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-mem-contact">
            <span class="cw-mem-ic bg-[#7ad0b0]"><Mail class="h-4 w-4" /></span>
            <span class="flex-1 min-w-0">
              <span class="block font-mono text-[7.5px] font-bold uppercase tracking-[0.14em] text-white/45">Write to us</span>
              <span class="block truncate text-[13px] font-bold text-white">{{ contact.email }}</span>
            </span>
            <span class="text-[16px] text-white/50">↗</span>
          </a>
          <div v-if="contact.address" class="cw-mem-contact">
            <span class="cw-mem-ic bg-[#a9c7ff]"><MapPin class="h-4 w-4" /></span>
            <span class="flex-1">
              <span class="block font-mono text-[7.5px] font-bold uppercase tracking-[0.14em] text-white/45">Find us</span>
              <span class="block text-[12.5px] font-medium leading-snug text-white/85">{{ contact.address }}</span>
            </span>
          </div>
          <p v-if="!hasContact" class="py-6 text-center text-[11px] font-bold text-white/40">联系方式未配置</p>
        </div>
      </section>

      <div class="cw-mem-in cw-mem-d3 mt-6">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-mem-cta">
          <span class="flex-1">访问官方网站</span>
          <span class="flex h-7 w-7 items-center justify-center rounded-full bg-[#191919] text-[#fff7e6]">↗</span>
        </a>
      </div>

      <footer class="cw-mem-in cw-mem-d3 mt-7 flex items-center justify-between border-t-2 border-[#191919] pt-2.5 font-mono text-[8px] font-bold uppercase tracking-[0.12em] text-[#191919]/50">
        <span>{{ name }} · Official</span>
        <span>Made with color</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, MapPin } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const name = computed(() => props.payload?.name || '企业官网')
const description = computed(() => props.payload?.description)
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const contact = computed(() => props.payload?.contactInfo || {})
const hasContact = computed(() => Boolean(contact.value.phone || contact.value.email || contact.value.address))
const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.url || ''))
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.cw-mem-dot {
  background-image: radial-gradient(#191919 1.6px, transparent 1.6px);
  background-size: 8px 8px;
  opacity: 0.35;
}
.cw-mem-tag {
  display: inline-block;
  border: 2px solid #191919;
  padding: 4px 8px;
  font-size: 9px;
  font-weight: 800;
  letter-spacing: 0.12em;
  line-height: 1;
  text-transform: uppercase;
}
.cw-mem-star {
  animation: cw-mem-spin 14s linear infinite;
  transform-origin: center;
}
@keyframes cw-mem-spin {
  to { transform: rotate(360deg); }
}
.cw-mem-card {
  position: relative;
  display: flex;
  min-height: 106px;
  flex-direction: column;
  border: 2px solid #191919;
  border-radius: 3px;
  padding: 10px;
  box-shadow: 3px 3px 0 #191919;
  transition: transform 0.18s ease, box-shadow 0.18s ease;
}
.cw-mem-card:hover {
  transform: translate(-2px, -2px) rotate(-0.6deg);
  box-shadow: 6px 6px 0 #191919;
}
.cw-mem-card:active {
  transform: translate(1px, 1px);
  box-shadow: 2px 2px 0 #191919;
}
.cw-mem-c0 { background: #f6c945; }
.cw-mem-c1 { background: #7ad0b0; }
.cw-mem-c2 { background: #a9c7ff; }
.cw-mem-c3 { background: #f7a8c8; }
.cw-mem-band {
  display: flex;
  align-items: center;
  justify-content: space-around;
  border: 2px solid #191919;
  background: #f6c945;
  padding: 12px 8px;
  box-shadow: 3px 3px 0 #191919;
}
.cw-mem-contactbox {
  border: 2px solid #191919;
  background: #191919;
  padding: 4px 14px;
  box-shadow: 4px 4px 0 #a9c7ff;
}
.cw-mem-contact {
  display: flex;
  align-items: center;
  gap: 11px;
  border-bottom: 1px dashed rgba(255, 255, 255, 0.18);
  padding: 10px 0;
}
.cw-mem-contact:last-child {
  border-bottom: 0;
}
a.cw-mem-contact {
  transition: padding-left 0.2s ease;
}
a.cw-mem-contact:hover {
  padding-left: 3px;
}
.cw-mem-ic {
  display: flex;
  height: 31px;
  width: 31px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  border: 1.5px solid #191919;
  color: #191919;
}
.cw-mem-cta {
  display: flex;
  align-items: center;
  gap: 10px;
  border: 2px solid #191919;
  background: #7ad0b0;
  padding: 11px 12px;
  font-size: 14px;
  font-weight: 800;
  box-shadow: 4px 4px 0 #191919;
  transition: transform 0.18s ease, box-shadow 0.18s ease, background-color 0.18s ease;
}
.cw-mem-cta:hover {
  background: #f6c945;
  transform: translate(-1px, -1px);
  box-shadow: 6px 6px 0 #191919;
}
.cw-mem-cta:active {
  transform: translate(2px, 2px);
  box-shadow: 1px 1px 0 #191919;
}
.cw-mem-in {
  animation: cw-mem-in 0.68s cubic-bezier(0.2, 0.9, 0.2, 1) both;
}
.cw-mem-d1 { animation-delay: 0.07s; }
.cw-mem-d2 { animation-delay: 0.16s; }
.cw-mem-d3 { animation-delay: 0.25s; }
@keyframes cw-mem-in {
  from { opacity: 0; transform: translateY(20px) rotate(-0.6deg); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-mem-in,
  .cw-mem-star {
    animation: none !important;
  }
  .cw-mem-card,
  .cw-mem-cta,
  .cw-mem-contact {
    transition: none !important;
  }
}
</style>

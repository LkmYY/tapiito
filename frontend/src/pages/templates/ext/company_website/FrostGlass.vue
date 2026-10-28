<template>
  <div class="cw-frost relative min-h-screen overflow-hidden bg-[#160b2e] text-white">
    <div class="cw-frost-orb cw-frost-orb-a pointer-events-none absolute -left-24 top-10 h-72 w-72 rounded-full"></div>
    <div class="cw-frost-orb cw-frost-orb-b pointer-events-none absolute -right-24 top-64 h-80 w-80 rounded-full"></div>
    <div class="cw-frost-orb cw-frost-orb-c pointer-events-none absolute bottom-10 left-1/3 h-64 w-64 rounded-full"></div>
    <div class="cw-frost-noise pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-10 pt-6">
      <header class="cw-frost-in flex items-center justify-between rounded-full border border-white/15 bg-white/[0.06] px-4 py-2.5 backdrop-blur-xl">
        <span class="flex items-center gap-2 text-[11px] font-medium tracking-wide">
          <span class="h-1.5 w-1.5 rounded-full bg-[#f0abfc]"></span>
          {{ name }}
        </span>
        <span class="font-mono text-[9px] uppercase tracking-[0.18em] text-white/40">NFC · {{ nfcId || 'ID' }}</span>
      </header>

      <section class="cw-frost-in cw-frost-d1 mt-5 overflow-hidden rounded-[26px] border border-white/15 bg-white/[0.075] p-7 backdrop-blur-2xl">
        <div class="flex items-center gap-2">
          <span class="rounded-full border border-[#f0abfc]/35 bg-[#f0abfc]/10 px-2.5 py-1 font-mono text-[8.5px] uppercase tracking-[0.18em] text-[#f0abfc]">Official profile</span>
          <span class="font-mono text-[8.5px] uppercase tracking-[0.16em] text-white/30">/ 2026</span>
        </div>
        <h1 class="mt-6 text-balance text-[40px] font-semibold leading-[1.04] tracking-[-0.04em]">{{ name }}</h1>
        <p class="mt-4 text-[13.5px] leading-[1.85] text-white/60">{{ description }}</p>
        <div class="mt-6 flex items-center gap-4 border-t border-white/10 pt-4">
          <div>
            <p class="text-[19px] font-semibold leading-none text-[#f0abfc]">{{ pad2(features.length) }}</p>
            <p class="mt-1 font-mono text-[8px] uppercase tracking-[0.16em] text-white/35">Services</p>
          </div>
          <span class="h-7 w-px bg-white/10"></span>
          <div>
            <p class="text-[19px] font-semibold leading-none text-[#f0abfc]">24H</p>
            <p class="mt-1 font-mono text-[8px] uppercase tracking-[0.16em] text-white/35">Reply</p>
          </div>
          <span class="ml-auto font-mono text-[8px] uppercase tracking-[0.18em] text-white/25">Glass / 01</span>
        </div>
      </section>

      <section class="cw-frost-in cw-frost-d2 mt-5 rounded-[24px] border border-white/15 bg-white/[0.055] p-5 backdrop-blur-xl">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[13px] font-medium tracking-wide">核心服务</h2>
          <span class="font-mono text-[8.5px] uppercase tracking-[0.2em] text-white/30">Capabilities</span>
        </div>
        <div v-if="features.length" class="mt-3 grid gap-2">
          <div v-for="(f, i) in features" :key="i" class="cw-frost-item group">
            <span class="flex h-7 w-7 flex-shrink-0 items-center justify-center rounded-full border border-white/15 bg-white/[0.06] font-mono text-[9px] text-[#f0abfc]">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[13.5px] text-white/80 transition-colors group-hover:text-white">{{ f }}</span>
            <span class="h-1.5 w-1.5 rounded-full bg-white/20 transition-all duration-300 group-hover:bg-[#f0abfc] group-hover:shadow-[0_0_12px_rgba(240,171,252,0.8)]"></span>
          </div>
        </div>
      </section>

      <section class="cw-frost-in cw-frost-d3 mt-5 rounded-[24px] border border-white/15 bg-gradient-to-br from-white/[0.11] to-white/[0.045] p-5 backdrop-blur-2xl">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[13px] font-medium tracking-wide">联系我们</h2>
          <span class="font-mono text-[8.5px] uppercase tracking-[0.2em] text-white/30">Contact</span>
        </div>
        <div class="mt-2">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-frost-contact">
            <span class="cw-frost-ic"><Phone class="h-3.5 w-3.5" /></span>
            <span class="flex-1">
              <span class="block font-mono text-[7.5px] uppercase tracking-[0.18em] text-white/30">Phone</span>
              <span class="block text-[12.5px] tabular-nums text-white/80">{{ contact.phone }}</span>
            </span>
            <ArrowUpRight class="h-3.5 w-3.5 text-white/25" />
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-frost-contact">
            <span class="cw-frost-ic"><Mail class="h-3.5 w-3.5" /></span>
            <span class="flex-1">
              <span class="block font-mono text-[7.5px] uppercase tracking-[0.18em] text-white/30">Email</span>
              <span class="block truncate text-[12.5px] text-white/80">{{ contact.email }}</span>
            </span>
            <ArrowUpRight class="h-3.5 w-3.5 text-white/25" />
          </a>
          <div v-if="contact.address" class="cw-frost-contact">
            <span class="cw-frost-ic"><MapPin class="h-3.5 w-3.5" /></span>
            <span class="flex-1">
              <span class="block font-mono text-[7.5px] uppercase tracking-[0.18em] text-white/30">Address</span>
              <span class="block text-[12.5px] leading-relaxed text-white/65">{{ contact.address }}</span>
            </span>
          </div>
          <p v-if="!hasContact" class="py-6 text-center text-[11px] text-white/35">联系方式未配置</p>
        </div>
      </section>

      <div class="cw-frost-in cw-frost-d3 mt-5">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-frost-cta">
          <span>访问官方网站</span>
          <ArrowUpRight class="h-4 w-4" />
        </a>
        <p class="mt-4 text-center font-mono text-[8px] uppercase tracking-[0.26em] text-white/20">Layered in glass · Built for touch</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, MapPin, ArrowUpRight } from 'lucide-vue-next'

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
.cw-frost-orb {
  filter: blur(48px);
  opacity: 0.55;
}
.cw-frost-orb-a {
  background: #7c3aed;
}
.cw-frost-orb-b {
  background: #db2777;
  opacity: 0.35;
}
.cw-frost-orb-c {
  background: #0ea5e9;
  opacity: 0.3;
}
.cw-frost-noise {
  opacity: 0.5;
  background-image: radial-gradient(rgba(255, 255, 255, 0.12) 0.5px, transparent 0.5px);
  background-size: 5px 5px;
}
.cw-frost-item {
  display: flex;
  align-items: center;
  gap: 10px;
  border: 1px solid rgba(255, 255, 255, 0.09);
  border-radius: 13px;
  background: rgba(255, 255, 255, 0.035);
  padding: 10px 11px;
  transition: border-color 0.25s ease, background-color 0.25s ease, transform 0.25s ease;
}
.cw-frost-item:hover {
  border-color: rgba(240, 171, 252, 0.4);
  background: rgba(255, 255, 255, 0.07);
  transform: translateX(2px);
}
.cw-frost-contact {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.09);
  padding: 10px 0;
  transition: padding-left 0.2s ease;
}
.cw-frost-contact:last-child {
  border-bottom: 0;
}
a.cw-frost-contact:hover {
  padding-left: 3px;
}
.cw-frost-ic {
  display: flex;
  height: 30px;
  width: 30px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 999px;
  border: 1px solid rgba(240, 171, 252, 0.25);
  background: rgba(240, 171, 252, 0.08);
  color: #f0abfc;
}
.cw-frost-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border: 1px solid rgba(255, 255, 255, 0.35);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.12);
  padding: 14px;
  font-size: 13px;
  font-weight: 600;
  backdrop-filter: blur(14px);
  transition: background-color 0.25s ease, transform 0.2s ease;
}
.cw-frost-cta:hover {
  background: #f0abfc;
  color: #2a0b3d;
}
.cw-frost-cta:active {
  transform: scale(0.985);
}
.cw-frost-in {
  animation: cw-frost-in 0.75s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cw-frost-d1 { animation-delay: 0.07s; }
.cw-frost-d2 { animation-delay: 0.16s; }
.cw-frost-d3 { animation-delay: 0.25s; }
@keyframes cw-frost-in {
  from { opacity: 0; transform: translateY(18px) scale(0.99); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-frost-in {
    animation: none !important;
  }
  .cw-frost-item,
  .cw-frost-contact,
  .cw-frost-cta {
    transition: none !important;
  }
}
</style>

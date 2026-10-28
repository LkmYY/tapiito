<template>
  <div class="cw-pastel relative min-h-screen overflow-hidden bg-[#f4eefb] text-[#2b2340]">
    <div class="cw-pastel-blob cw-pastel-b1 pointer-events-none absolute -left-16 top-16 h-56 w-56 rounded-full"></div>
    <div class="cw-pastel-blob cw-pastel-b2 pointer-events-none absolute -right-16 top-48 h-64 w-64 rounded-full"></div>
    <div class="cw-pastel-blob cw-pastel-b3 pointer-events-none absolute bottom-10 left-1/3 h-48 w-48 rounded-full"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-4 pb-8 pt-5">
      <header class="cw-pastel-in flex items-center justify-between px-1">
        <span class="flex items-center gap-2">
          <span class="flex h-7 w-7 items-center justify-center rounded-xl bg-[#7c6cf0] text-[11px] font-bold text-white">{{ initial }}</span>
          <span class="text-[11px] font-semibold tracking-wide">{{ name }}</span>
        </span>
        <span class="rounded-full bg-white/70 px-2.5 py-1 font-mono text-[8px] uppercase tracking-[0.12em] text-[#7c6cf0]">NFC {{ nfcId || '01' }}</span>
      </header>

      <section class="cw-pastel-in cw-pastel-d1 cw-pastel-card mt-4 bg-[#ffffff]/80 p-6">
        <div class="flex items-center gap-2">
          <span class="h-2 w-2 rounded-full bg-[#ff9ec7]"></span>
          <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#7c6cf0]">Company profile</span>
        </div>
        <h1 class="mt-4 text-[37px] font-bold leading-[1.04] tracking-[-0.04em] text-[#241c38]">{{ name }}</h1>
        <p class="mt-3 text-[13px] leading-[1.8] text-[#2b2340]/60">{{ description }}</p>
        <div class="mt-5 flex items-center gap-2">
          <span class="rounded-full bg-[#ede9ff] px-2.5 py-1 font-mono text-[8px] text-[#7c6cf0]">{{ pad2(features.length) }} SERVICES</span>
          <span class="rounded-full bg-[#ffe4ef] px-2.5 py-1 font-mono text-[8px] text-[#c2517c]">24H REPLY</span>
        </div>
      </section>

      <section class="cw-pastel-in cw-pastel-d2 mt-3 grid grid-cols-3 gap-2.5">
        <div class="cw-pastel-tile bg-[#e8e3ff]">
          <p class="text-[20px] font-bold leading-none text-[#5a49d6]">{{ pad2(features.length) }}</p>
          <p class="mt-1.5 text-[9px] font-medium text-[#5a49d6]/60">服务模块</p>
        </div>
        <div class="cw-pastel-tile bg-[#ffe1ec]">
          <p class="text-[20px] font-bold leading-none text-[#c2517c]">24H</p>
          <p class="mt-1.5 text-[9px] font-medium text-[#c2517c]/60">响应时效</p>
        </div>
        <div class="cw-pastel-tile bg-[#d9f4ee]">
          <p class="text-[20px] font-bold leading-none text-[#2f8a78]">1:1</p>
          <p class="mt-1.5 text-[9px] font-medium text-[#2f8a78]/60">专属对接</p>
        </div>
      </section>

      <section class="cw-pastel-in cw-pastel-d2 cw-pastel-card mt-3 bg-white/75 p-5">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[12px] font-semibold">我们提供</h2>
          <span class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#2b2340]/35">Services</span>
        </div>
        <div v-if="features.length" class="mt-3 grid grid-cols-2 gap-2">
          <div v-for="(f, i) in features" :key="i" class="cw-pastel-svc group">
            <span class="cw-pastel-dot" :class="'cw-pastel-dot-' + (i % 3)"></span>
            <span class="text-[12px] font-medium leading-snug text-[#2b2340]/80 transition-colors group-hover:text-[#5a49d6]">{{ f }}</span>
          </div>
        </div>
      </section>

      <section class="cw-pastel-in cw-pastel-d3 cw-pastel-card mt-3 bg-[#2b2340] p-5 text-white">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[12px] font-semibold">打个招呼</h2>
          <span class="font-mono text-[8px] uppercase tracking-[0.14em] text-white/35">Contact</span>
        </div>
        <div class="mt-2">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-pastel-contact">
            <span class="cw-pastel-ic bg-[#ffe1ec] text-[#c2517c]"><Phone class="h-3.5 w-3.5" /></span>
            <span class="flex-1 text-[12px] tabular-nums text-white/80">{{ contact.phone }}</span>
            <span class="text-white/30">↗</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-pastel-contact">
            <span class="cw-pastel-ic bg-[#e8e3ff] text-[#5a49d6]"><Mail class="h-3.5 w-3.5" /></span>
            <span class="flex-1 truncate text-[12px] text-white/80">{{ contact.email }}</span>
            <span class="text-white/30">↗</span>
          </a>
          <div v-if="contact.address" class="cw-pastel-contact items-start">
            <span class="cw-pastel-ic bg-[#d9f4ee] text-[#2f8a78]"><MapPin class="h-3.5 w-3.5" /></span>
            <span class="flex-1 text-[11.5px] leading-relaxed text-white/65">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-4 text-center text-[10px] text-white/35">联系方式未配置</p>
        </div>
      </section>

      <div class="cw-pastel-in cw-pastel-d3 mt-3">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-pastel-cta">
          <span>访问官方网站</span>
          <span class="flex h-7 w-7 items-center justify-center rounded-full bg-white/20"><ArrowUpRight class="h-3.5 w-3.5" /></span>
        </a>
      </div>

      <footer class="cw-pastel-in cw-pastel-d3 mt-5 flex items-center justify-between px-1 font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#2b2340]/30">
        <span>{{ name }}</span>
        <span>Softly, seriously</span>
      </footer>
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
const initial = computed(() => String(name.value).trim().charAt(0) || '企')
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.cw-pastel-blob {
  filter: blur(46px);
  opacity: 0.5;
}
.cw-pastel-b1 { background: #c9c0ff; }
.cw-pastel-b2 { background: #ffc3da; }
.cw-pastel-b3 { background: #b7ebdf; }
.cw-pastel-card {
  border: 1px solid rgba(124, 108, 240, 0.12);
  border-radius: 22px;
  box-shadow: 0 18px 35px -28px rgba(70, 50, 130, 0.55);
  backdrop-filter: blur(6px);
}
.cw-pastel-tile {
  border-radius: 16px;
  padding: 12px;
  transition: transform 0.22s ease;
}
.cw-pastel-tile:hover {
  transform: translateY(-3px);
}
.cw-pastel-svc {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  border-radius: 12px;
  background: #f7f4ff;
  padding: 9px;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-pastel-svc:hover {
  background: #ede9ff;
  transform: translateY(-2px);
}
.cw-pastel-dot {
  margin-top: 4px;
  height: 6px;
  width: 6px;
  flex-shrink: 0;
  border-radius: 99px;
}
.cw-pastel-dot-0 { background: #7c6cf0; }
.cw-pastel-dot-1 { background: #ff9ec7; }
.cw-pastel-dot-2 { background: #6ac7b4; }
.cw-pastel-contact {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.09);
  padding: 9px 0;
  transition: padding-left 0.2s ease;
}
.cw-pastel-contact:last-child {
  border-bottom: 0;
}
a.cw-pastel-contact:hover {
  padding-left: 3px;
}
.cw-pastel-ic {
  display: flex;
  height: 28px;
  width: 28px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 9px;
}
.cw-pastel-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  border-radius: 18px;
  background: #7c6cf0;
  padding: 9px 9px 9px 18px;
  font-size: 13px;
  font-weight: 600;
  color: white;
  box-shadow: 0 15px 26px -18px rgba(90, 73, 214, 0.9);
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-pastel-cta:hover {
  background: #5a49d6;
}
.cw-pastel-cta:active {
  transform: scale(0.985);
}
.cw-pastel-in {
  animation: cw-pastel-in 0.72s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cw-pastel-d1 { animation-delay: 0.07s; }
.cw-pastel-d2 { animation-delay: 0.16s; }
.cw-pastel-d3 { animation-delay: 0.25s; }
@keyframes cw-pastel-in {
  from { opacity: 0; transform: translateY(18px) scale(0.99); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-pastel-in {
    animation: none !important;
  }
  .cw-pastel-svc,
  .cw-pastel-tile,
  .cw-pastel-contact,
  .cw-pastel-cta {
    transition: none !important;
  }
}
</style>

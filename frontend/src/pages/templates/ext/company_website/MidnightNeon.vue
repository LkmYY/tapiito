<template>
  <div class="cw-neon relative min-h-screen overflow-hidden bg-[#0a0617] text-[#f4eaff]">
    <div class="cw-neon-halo pointer-events-none absolute -top-28 left-1/2 h-72 w-72 -translate-x-1/2 rounded-full"></div>
    <div class="cw-neon-scan pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute bottom-24 right-4 top-40 w-px bg-gradient-to-b from-transparent via-[#ff4fa3]/40 to-transparent"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-10 pt-7">
      <header class="cw-neon-in flex items-center justify-between border-b border-[#ff4fa3]/20 pb-3">
        <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.24em] text-[#ff8bc4]">
          <span class="h-1.5 w-1.5 rounded-full bg-[#ff4fa3] shadow-[0_0_10px_#ff4fa3]"></span>
          After-hours office
        </span>
        <span class="font-mono text-[9px] uppercase tracking-[0.16em] text-[#b9a7d8]/45">N{{ nfcId || '01' }}</span>
      </header>

      <section class="cw-neon-in cw-neon-d1 mt-10">
        <p class="font-mono text-[9px] uppercase tracking-[0.34em] text-[#8f7ab5]">Company / Night edition</p>
        <h1 class="cw-neon-name mt-4 break-words text-[50px] font-black leading-[0.88] tracking-[-0.055em] text-[#fff3fb]">{{ name }}</h1>
        <div class="mt-5 flex items-start gap-3">
          <span class="cw-neon-bar mt-1 flex-shrink-0"></span>
          <p class="text-[13.5px] leading-[1.85] text-[#c9badf]/70">{{ description }}</p>
        </div>
      </section>

      <section class="cw-neon-in cw-neon-d2 mt-8 flex items-center gap-3 rounded-lg border border-[#ff4fa3]/25 bg-[#ff4fa3]/[0.05] px-3 py-2">
        <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#ff8bc4]">Now serving</span>
        <span class="h-px flex-1 bg-[#ff4fa3]/20"></span>
        <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#c9badf]/45">24 / 7 · Global</span>
      </section>

      <section class="cw-neon-in cw-neon-d2 mt-8">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[13px] font-semibold tracking-wide">服务清单</h2>
          <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-[#b9a7d8]/40">{{ pad2(features.length) }} Items</span>
        </div>
        <div v-if="features.length" class="mt-3 space-y-2">
          <div v-for="(f, i) in features" :key="i" class="cw-neon-row group">
            <span class="font-mono text-[10px] text-[#ff8bc4]/60">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[14px] font-medium leading-snug text-[#f4eaff]/85 transition-colors group-hover:text-white">{{ f }}</span>
            <span class="cw-neon-dot"></span>
          </div>
        </div>
      </section>

      <section class="cw-neon-in cw-neon-d2 mt-7 grid grid-cols-3 border-y border-white/10">
        <div class="py-4">
          <p class="text-[19px] font-bold leading-none text-[#ff8bc4]">{{ pad2(features.length) }}</p>
          <p class="mt-1.5 font-mono text-[7.5px] uppercase tracking-[0.14em] text-[#b9a7d8]/40">Services</p>
        </div>
        <div class="border-x border-white/10 px-4 py-4">
          <p class="text-[19px] font-bold leading-none text-[#8ff5e1]">24H</p>
          <p class="mt-1.5 font-mono text-[7.5px] uppercase tracking-[0.14em] text-[#b9a7d8]/40">Reply</p>
        </div>
        <div class="py-4 pl-4">
          <p class="text-[19px] font-bold leading-none text-[#ff8bc4]">VIP</p>
          <p class="mt-1.5 font-mono text-[7.5px] uppercase tracking-[0.14em] text-[#b9a7d8]/40">Service</p>
        </div>
      </section>

      <section class="cw-neon-in cw-neon-d3 mt-8">
        <div class="flex items-baseline justify-between border-b border-[#ff4fa3]/20 pb-2">
          <h2 class="text-[13px] font-semibold tracking-wide">夜间联络</h2>
          <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-[#b9a7d8]/40">Contact</span>
        </div>
        <div class="mt-1">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-neon-contact">
            <span class="cw-neon-ic"><Phone class="h-3.5 w-3.5" /></span>
            <span class="cw-neon-k">电话</span>
            <span class="cw-neon-v tabular-nums">{{ contact.phone }}</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-neon-contact">
            <span class="cw-neon-ic"><Mail class="h-3.5 w-3.5" /></span>
            <span class="cw-neon-k">邮箱</span>
            <span class="cw-neon-v">{{ contact.email }}</span>
          </a>
          <div v-if="contact.address" class="cw-neon-contact items-start">
            <span class="cw-neon-ic"><MapPin class="h-3.5 w-3.5" /></span>
            <span class="cw-neon-k">地址</span>
            <span class="cw-neon-v whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-6 text-center text-[11px] text-[#b9a7d8]/40">联系方式未配置</p>
        </div>
      </section>

      <div class="cw-neon-in cw-neon-d3 mt-7">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-neon-cta">
          <span class="flex-1">访问官方网站</span>
          <span class="cw-neon-cta-ic"><ArrowUpRight class="h-3.5 w-3.5" /></span>
        </a>
      </div>

      <footer class="cw-neon-in cw-neon-d3 mt-7 flex items-center justify-between border-t border-white/10 pt-3 font-mono text-[8px] uppercase tracking-[0.2em] text-[#b9a7d8]/30">
        <span>{{ name }} · After dark</span>
        <span>Tokyo / Shanghai</span>
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
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.cw-neon-halo {
  background: radial-gradient(circle, rgba(255, 79, 163, 0.22), rgba(90, 40, 160, 0.12) 45%, transparent 70%);
  filter: blur(10px);
}
.cw-neon-scan {
  background:
    linear-gradient(rgba(255, 79, 163, 0.03) 1px, transparent 1px);
  background-size: 100% 5px;
  opacity: 0.6;
}
.cw-neon-name {
  text-shadow: 0 0 24px rgba(255, 79, 163, 0.35), 0 0 2px rgba(255, 255, 255, 0.35);
}
.cw-neon-bar {
  height: 42px;
  width: 3px;
  background: linear-gradient(to bottom, #ff4fa3, #8ff5e1);
  box-shadow: 0 0 14px rgba(255, 79, 163, 0.55);
  border-radius: 99px;
}
.cw-neon-row {
  display: flex;
  align-items: center;
  gap: 11px;
  border: 1px solid rgba(255, 79, 163, 0.16);
  border-radius: 9px;
  background: linear-gradient(100deg, rgba(255, 79, 163, 0.05), rgba(143, 245, 225, 0.025));
  padding: 12px;
  transition: border-color 0.25s ease, box-shadow 0.25s ease, transform 0.25s ease;
}
.cw-neon-row:hover {
  border-color: rgba(255, 79, 163, 0.55);
  box-shadow: 0 0 22px -6px rgba(255, 79, 163, 0.5);
  transform: translateX(2px);
}
.cw-neon-dot {
  height: 5px;
  width: 5px;
  flex-shrink: 0;
  border-radius: 99px;
  background: #8ff5e1;
  box-shadow: 0 0 10px rgba(143, 245, 225, 0.8);
  transition: transform 0.25s ease;
}
.cw-neon-row:hover .cw-neon-dot {
  transform: scale(1.45);
}
.cw-neon-contact {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  padding: 11px 0;
  transition: padding-left 0.2s ease;
}
.cw-neon-contact:last-child {
  border-bottom: 0;
}
a.cw-neon-contact:hover {
  padding-left: 3px;
}
.cw-neon-ic {
  display: flex;
  height: 29px;
  width: 29px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  border: 1px solid rgba(255, 79, 163, 0.3);
  background: rgba(255, 79, 163, 0.07);
  color: #ff8bc4;
}
.cw-neon-k {
  flex-shrink: 0;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: rgba(185, 167, 216, 0.4);
}
.cw-neon-v {
  margin-left: auto;
  max-width: 62%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  color: rgba(244, 234, 255, 0.8);
  transition: color 0.2s ease;
}
a.cw-neon-contact:hover .cw-neon-v {
  color: #ff8bc4;
}
.cw-neon-cta {
  display: flex;
  align-items: center;
  gap: 10px;
  border-radius: 10px;
  border: 1px solid rgba(255, 79, 163, 0.55);
  background: linear-gradient(100deg, rgba(255, 79, 163, 0.2), rgba(143, 245, 225, 0.12));
  padding: 12px 12px 12px 16px;
  font-size: 13px;
  font-weight: 600;
  box-shadow: 0 0 30px -12px rgba(255, 79, 163, 0.65);
  transition: background-color 0.25s ease, transform 0.2s ease;
}
.cw-neon-cta:hover {
  background: rgba(255, 79, 163, 0.32);
}
.cw-neon-cta:active {
  transform: scale(0.985);
}
.cw-neon-cta-ic {
  display: flex;
  height: 27px;
  width: 27px;
  align-items: center;
  justify-content: center;
  border-radius: 99px;
  background: #ff4fa3;
  color: #200816;
}
.cw-neon-in {
  animation: cw-neon-in 0.7s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cw-neon-d1 { animation-delay: 0.07s; }
.cw-neon-d2 { animation-delay: 0.16s; }
.cw-neon-d3 { animation-delay: 0.25s; }
@keyframes cw-neon-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-neon-in {
    animation: none !important;
  }
  .cw-neon-row,
  .cw-neon-contact,
  .cw-neon-dot,
  .cw-neon-cta {
    transition: none !important;
  }
}
</style>

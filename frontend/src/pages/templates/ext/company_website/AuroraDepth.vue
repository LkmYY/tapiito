<template>
  <div class="cw-aur relative min-h-screen overflow-hidden bg-[#041a1c] text-[#e6fff8]">
    <div class="cw-aur-ribbon cw-aur-r1 pointer-events-none absolute -left-24 top-0 h-80 w-[34rem]"></div>
    <div class="cw-aur-ribbon cw-aur-r2 pointer-events-none absolute -right-32 top-56 h-80 w-[36rem]"></div>
    <div class="cw-aur-ribbon cw-aur-r3 pointer-events-none absolute -left-20 bottom-0 h-72 w-[30rem]"></div>
    <div class="cw-aur-stars pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto flex min-h-screen w-full max-w-md flex-col px-7 pb-9 pt-7">
      <header class="cw-aur-in flex items-center justify-between">
        <span class="flex items-center gap-2 font-mono text-[8.5px] uppercase tracking-[0.24em] text-[#9ff2df]/70">
          <span class="h-1.5 w-1.5 rounded-full bg-[#7bf0d4] shadow-[0_0_10px_rgba(123,240,212,0.9)]"></span>
          Aurora Company
        </span>
        <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#9ff2df]/35">{{ nfcId || '01' }}</span>
      </header>

      <section class="cw-aur-in cw-aur-d1 mt-20 text-center">
        <div class="cw-aur-ring mx-auto flex h-16 w-16 items-center justify-center rounded-full">
          <span class="h-2 w-2 rounded-full bg-[#c7fff0]"></span>
        </div>
        <p class="mt-6 font-mono text-[8px] uppercase tracking-[0.42em] text-[#9ff2df]/55">Official horizon</p>
        <h1 class="mt-4 text-[44px] font-light leading-[1.02] tracking-[-0.035em] text-[#f0fffa]">{{ name }}</h1>
        <p class="mx-auto mt-5 max-w-[18rem] text-[12.5px] leading-[1.95] text-[#c8ece4]/60">{{ description }}</p>
      </section>

      <div class="cw-aur-in cw-aur-d1 mt-8 flex items-center justify-center gap-2 font-mono text-[8px] uppercase tracking-[0.18em] text-[#9ff2df]/40">
        <span>{{ pad2(features.length) }} capabilities</span>
        <span class="h-1 w-1 rotate-45 bg-[#7bf0d4]/50"></span>
        <span>24H response</span>
        <span class="h-1 w-1 rotate-45 bg-[#7bf0d4]/50"></span>
        <span>1:1 care</span>
      </div>

      <section class="cw-aur-in cw-aur-d2 mt-12">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-gradient-to-r from-transparent to-[#9ff2df]/30"></span>
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.3em] text-[#9ff2df]/70">Service constellation</h2>
          <span class="h-px flex-1 bg-gradient-to-l from-transparent to-[#9ff2df]/30"></span>
        </div>
        <div v-if="features.length" class="mt-4 space-y-2">
          <div v-for="(f, i) in features" :key="i" class="cw-aur-item group">
            <span class="font-mono text-[9px] tabular-nums text-[#7bf0d4]/55">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[13.5px] font-light leading-snug text-[#e6fff8]/80 transition-colors group-hover:text-white">{{ f }}</span>
            <span class="cw-aur-star"></span>
          </div>
        </div>
        <p v-else class="mt-4 border border-dashed border-[#9ff2df]/20 py-8 text-center font-mono text-[10px] text-[#9ff2df]/35">SERVICES NOT CONFIGURED</p>
      </section>

      <section class="cw-aur-in cw-aur-d3 mt-10">
        <div class="text-center">
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.3em] text-[#9ff2df]/70">Signal us</h2>
        </div>
        <div class="mt-3">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-aur-contact">
            <span class="cw-aur-k">电话</span>
            <span class="cw-aur-v tabular-nums">{{ contact.phone }}</span>
            <span class="cw-aur-dot"></span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-aur-contact">
            <span class="cw-aur-k">邮箱</span>
            <span class="cw-aur-v">{{ contact.email }}</span>
            <span class="cw-aur-dot"></span>
          </a>
          <div v-if="contact.address" class="cw-aur-contact items-start">
            <span class="cw-aur-k">地址</span>
            <span class="cw-aur-v whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-5 text-center text-[11px] font-light text-[#9ff2df]/35">联系方式未配置</p>
        </div>
      </section>

      <div class="cw-aur-in cw-aur-d3 mt-8">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-aur-cta">
          <span class="flex-1 text-center">访问官方网站</span>
          <ArrowUpRight class="h-3.5 w-3.5" />
        </a>
      </div>

      <footer class="mt-auto flex items-center justify-between pt-10 font-mono text-[7px] uppercase tracking-[0.22em] text-[#9ff2df]/25">
        <span>{{ name }}</span>
        <span>Under the same sky</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

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
.cw-aur-ribbon {
  filter: blur(46px);
  opacity: 0.55;
  border-radius: 50%;
}
.cw-aur-r1 {
  background: linear-gradient(100deg, rgba(50, 220, 190, 0.55), rgba(80, 120, 255, 0.3), transparent 75%);
  transform: rotate(-18deg);
}
.cw-aur-r2 {
  background: linear-gradient(120deg, rgba(130, 90, 255, 0.45), rgba(40, 220, 200, 0.32), transparent 78%);
  transform: rotate(14deg);
}
.cw-aur-r3 {
  background: linear-gradient(80deg, rgba(30, 160, 255, 0.35), rgba(120, 255, 220, 0.25), transparent 80%);
  transform: rotate(-8deg);
}
.cw-aur-stars {
  background-image:
    radial-gradient(rgba(220, 255, 248, 0.55) 0.5px, transparent 0.5px),
    radial-gradient(rgba(160, 235, 255, 0.4) 0.4px, transparent 0.4px);
  background-size: 37px 37px, 61px 61px;
  background-position: 0 0, 12px 19px;
  -webkit-mask-image: linear-gradient(to bottom, black, transparent 85%);
  mask-image: linear-gradient(to bottom, black, transparent 85%);
  opacity: 0.6;
}
.cw-aur-ring {
  border: 1px solid rgba(159, 242, 223, 0.25);
  box-shadow: 0 0 40px rgba(123, 240, 212, 0.22), inset 0 0 26px rgba(123, 240, 212, 0.12);
  animation: cw-aur-breathe 5s ease-in-out infinite;
}
@keyframes cw-aur-breathe {
  50% { box-shadow: 0 0 60px rgba(123, 240, 212, 0.4), inset 0 0 32px rgba(123, 240, 212, 0.2); }
}
.cw-aur-item {
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid rgba(159, 242, 223, 0.12);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.025);
  padding: 10px 14px;
  backdrop-filter: blur(4px);
  transition: border-color 0.3s ease, background-color 0.3s ease, transform 0.3s ease;
}
.cw-aur-item:hover {
  border-color: rgba(159, 242, 223, 0.4);
  background: rgba(255, 255, 255, 0.05);
  transform: translateX(3px);
}
.cw-aur-star {
  height: 4px;
  width: 4px;
  flex-shrink: 0;
  border-radius: 99px;
  background: #7bf0d4;
  box-shadow: 0 0 8px rgba(123, 240, 212, 0.9);
}
.cw-aur-contact {
  display: flex;
  align-items: baseline;
  gap: 12px;
  border-bottom: 1px solid rgba(159, 242, 223, 0.13);
  padding: 10px 0;
  transition: background-color 0.2s ease;
}
a.cw-aur-contact:hover {
  background: rgba(255, 255, 255, 0.03);
}
.cw-aur-k {
  width: 34px;
  flex-shrink: 0;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: rgba(159, 242, 223, 0.4);
}
.cw-aur-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: center;
  font-size: 12px;
  font-weight: 300;
  color: rgba(230, 255, 248, 0.8);
}
a.cw-aur-contact:hover .cw-aur-v {
  color: #ffffff;
}
.cw-aur-dot {
  height: 4px;
  width: 4px;
  flex-shrink: 0;
  border-radius: 99px;
  background: rgba(159, 242, 223, 0.35);
}
.cw-aur-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border: 1px solid rgba(159, 242, 223, 0.4);
  border-radius: 999px;
  background: rgba(123, 240, 212, 0.08);
  padding: 13px;
  font-size: 12.5px;
  letter-spacing: 0.04em;
  transition: background-color 0.25s ease, color 0.25s ease, transform 0.2s ease;
}
.cw-aur-cta:hover {
  background: #9ff2df;
  color: #072420;
}
.cw-aur-cta:active {
  transform: scale(0.985);
}
.cw-aur-in {
  animation: cw-aur-in 0.9s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-aur-d1 { animation-delay: 0.08s; }
.cw-aur-d2 { animation-delay: 0.18s; }
.cw-aur-d3 { animation-delay: 0.28s; }
@keyframes cw-aur-in {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-aur-in,
  .cw-aur-ring {
    animation: none !important;
  }
  .cw-aur-item,
  .cw-aur-contact,
  .cw-aur-cta {
    transition: none !important;
  }
}
</style>

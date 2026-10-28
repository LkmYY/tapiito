<template>
  <div class="cw-kin relative min-h-screen overflow-hidden bg-[#101010] text-[#f4ef2a]">
    <div class="cw-kin-marquee pointer-events-none absolute inset-x-0 top-[42%] select-none">
      <div class="cw-kin-track">
        <span v-for="n in 4" :key="n" class="cw-kin-ghost">{{ name }}</span>
      </div>
    </div>

    <div class="relative z-10 mx-auto flex min-h-screen w-full max-w-md flex-col px-6 pb-8 pt-6">
      <header class="cw-kin-in flex items-center justify-between border-b border-[#f4ef2a]/25 pb-3">
        <span class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#f4ef2a]/70">Poster No. {{ nfcId || '01' }}</span>
        <span class="h-2 w-2 bg-[#f4ef2a]"></span>
        <span class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#f4ef2a]/70">Company / 官网</span>
      </header>

      <section class="cw-kin-in cw-kin-d1 mt-8">
        <p class="font-mono text-[10px] uppercase tracking-[0.32em] text-[#f4ef2a]/60">Make it loud · Make it clear</p>
        <h1 class="mt-3 text-[64px] font-black leading-[0.82] tracking-[-0.065em] break-words">{{ name }}</h1>
      </section>

      <section class="cw-kin-in cw-kin-d2 mt-6 flex items-stretch gap-3">
        <div class="cw-kin-side">
          <span class="cw-kin-vertical font-mono text-[8px] uppercase tracking-[0.34em] text-[#101010]/60">Est. Corporate</span>
        </div>
        <p class="flex-1 border-t-2 border-[#f4ef2a] pt-3 text-[13px] leading-[1.8] text-[#f4ef2a]/75">{{ description }}</p>
      </section>

      <div class="cw-kin-in cw-kin-d2 mt-5 overflow-hidden border-y border-[#f4ef2a]/30 py-2">
        <div class="cw-kin-ticker">
          <span v-for="n in 2" :key="n" class="flex shrink-0 items-center">
            <span class="px-4 font-mono text-[9px] uppercase tracking-[0.22em] text-[#f4ef2a]/55">Strategy</span>
            <span class="text-[#f4ef2a]/30">◆</span>
            <span class="px-4 font-mono text-[9px] uppercase tracking-[0.22em] text-[#f4ef2a]/55">Design</span>
            <span class="text-[#f4ef2a]/30">◆</span>
            <span class="px-4 font-mono text-[9px] uppercase tracking-[0.22em] text-[#f4ef2a]/55">Delivery</span>
            <span class="text-[#f4ef2a]/30">◆</span>
            <span class="px-4 font-mono text-[9px] uppercase tracking-[0.22em] text-[#f4ef2a]/55">Official site</span>
            <span class="text-[#f4ef2a]/30">◆</span>
          </span>
        </div>
      </div>

      <section class="cw-kin-in cw-kin-d3 mt-6">
        <div class="flex items-baseline justify-between">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.28em] text-[#f4ef2a]/50">What we do</h2>
          <span class="font-mono text-[9px] text-[#f4ef2a]/40">{{ pad2(features.length) }}</span>
        </div>
        <div v-if="features.length" class="mt-1">
          <div v-for="(f, i) in features" :key="i" class="cw-kin-row group">
            <span class="w-8 font-mono text-[10px] text-[#f4ef2a]/45">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[17px] font-bold leading-snug transition-transform duration-200 group-hover:translate-x-1">{{ f }}</span>
            <span class="font-black text-[#f4ef2a]/25 transition-colors group-hover:text-[#f4ef2a]">↗</span>
          </div>
        </div>
      </section>

      <section class="cw-kin-in cw-kin-d3 mt-5 grid grid-cols-3 border-y border-[#f4ef2a]/30 py-3">
        <div>
          <p class="text-[22px] font-black leading-none">{{ pad2(features.length) }}</p>
          <p class="mt-1 font-mono text-[7px] uppercase tracking-[0.14em] text-[#f4ef2a]/40">Services</p>
        </div>
        <div class="border-x border-[#f4ef2a]/20 px-3">
          <p class="text-[22px] font-black leading-none">24H</p>
          <p class="mt-1 font-mono text-[7px] uppercase tracking-[0.14em] text-[#f4ef2a]/40">Reply</p>
        </div>
        <div class="pl-3">
          <p class="text-[22px] font-black leading-none">1:1</p>
          <p class="mt-1 font-mono text-[7px] uppercase tracking-[0.14em] text-[#f4ef2a]/40">Contact</p>
        </div>
      </section>

      <section class="cw-kin-in cw-kin-d3 mt-6 bg-[#f4ef2a] px-4 py-4 text-[#101010]">
        <div class="flex items-center justify-between border-b border-[#101010]/25 pb-2">
          <h2 class="font-mono text-[9px] font-bold uppercase tracking-[0.22em]">Contact</h2>
          <span class="font-mono text-[9px] font-bold">{{ name }}</span>
        </div>
        <div class="mt-1">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-kin-contact">
            <span class="font-mono text-[8px] font-bold uppercase tracking-[0.14em]">电话</span>
            <span class="flex-1 text-right text-[13px] font-bold tabular-nums">{{ contact.phone }}</span>
            <span class="font-black">↗</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-kin-contact">
            <span class="font-mono text-[8px] font-bold uppercase tracking-[0.14em]">邮箱</span>
            <span class="flex-1 truncate text-right text-[12.5px] font-bold">{{ contact.email }}</span>
            <span class="font-black">↗</span>
          </a>
          <div v-if="contact.address" class="cw-kin-contact items-start">
            <span class="font-mono text-[8px] font-bold uppercase tracking-[0.14em]">地址</span>
            <span class="flex-1 text-right text-[12px] font-medium leading-snug">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-3 text-center text-[11px] font-bold">联系方式未配置</p>
        </div>
      </section>

      <div class="cw-kin-in cw-kin-d3 mt-4">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-kin-cta">
          <span class="text-[13px] font-bold uppercase tracking-[0.16em] text-[#101010]">访问官方网站</span>
          <span class="flex h-8 w-8 items-center justify-center bg-[#101010] text-[#f4ef2a]"><ArrowUpRight class="h-4 w-4" /></span>
        </a>
      </div>

      <footer class="mt-auto flex items-center justify-between pt-7 font-mono text-[8px] uppercase tracking-[0.2em] text-[#f4ef2a]/30">
        <span>© {{ name }}</span>
        <span>Printed for screen</span>
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
.cw-kin-marquee {
  opacity: 0.045;
  overflow: hidden;
  white-space: nowrap;
}
.cw-kin-track {
  display: flex;
  animation: cw-kin-slide 26s linear infinite;
}
.cw-kin-ghost {
  font-size: 130px;
  font-weight: 900;
  line-height: 1;
  letter-spacing: -0.06em;
  padding-right: 40px;
  color: #f4ef2a;
}
.cw-kin-ticker {
  display: flex;
  width: max-content;
  animation: cw-kin-ticker 22s linear infinite;
}
.cw-kin-side {
  display: flex;
  width: 22px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  background: #f4ef2a;
}
.cw-kin-vertical {
  writing-mode: vertical-rl;
  transform: rotate(180deg);
}
.cw-kin-row {
  display: flex;
  align-items: center;
  gap: 9px;
  border-bottom: 1px solid rgba(244, 239, 42, 0.18);
  padding: 10px 0;
  transition: background-color 0.2s ease;
}
.cw-kin-row:hover {
  background: rgba(244, 239, 42, 0.06);
}
.cw-kin-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px solid rgba(16, 16, 16, 0.2);
  padding: 9px 0;
  transition: background-color 0.2s ease;
}
.cw-kin-contact:last-child {
  border-bottom: 0;
}
a.cw-kin-contact:hover {
  background: rgba(16, 16, 16, 0.06);
}
.cw-kin-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  background: #f4ef2a;
  padding: 7px 7px 7px 16px;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-kin-cta:hover {
  background: #ffffff;
}
.cw-kin-cta:active {
  transform: scale(0.985);
}
.cw-kin-in {
  animation: cw-kin-in 0.6s cubic-bezier(0.16, 0.9, 0.2, 1) both;
}
.cw-kin-d1 { animation-delay: 0.05s; }
.cw-kin-d2 { animation-delay: 0.14s; }
.cw-kin-d3 { animation-delay: 0.23s; }
@keyframes cw-kin-in {
  from { opacity: 0; transform: translateY(22px); }
  to { opacity: 1; transform: none; }
}
@keyframes cw-kin-slide {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@keyframes cw-kin-ticker {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@media (prefers-reduced-motion: reduce) {
  .cw-kin-in,
  .cw-kin-track,
  .cw-kin-ticker {
    animation: none !important;
  }
  .cw-kin-row,
  .cw-kin-contact,
  .cw-kin-cta {
    transition: none !important;
  }
}
</style>

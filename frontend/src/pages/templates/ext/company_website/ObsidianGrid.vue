<template>
  <div class="cw-obs relative min-h-screen overflow-hidden bg-[#06080a] text-[#e9f0e3]">
    <div class="cw-obs-grid pointer-events-none absolute inset-0"></div>
    <div class="cw-obs-glow pointer-events-none absolute -top-36 left-1/2 h-80 w-80 -translate-x-1/2 rounded-full"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-10 pt-8">
      <header class="cw-obs-rise flex items-center justify-between border-b border-[#c6f24e]/15 pb-4">
        <span class="flex items-center gap-2 font-mono text-[10px] uppercase tracking-[0.24em] text-[#c6f24e]">
          <span class="cw-obs-dot h-1.5 w-1.5 rounded-full bg-[#c6f24e]"></span>
          Corporate Node
        </span>
        <span class="font-mono text-[10px] tracking-[0.16em] text-white/30">NFC {{ nfcId || '—' }}</span>
      </header>

      <section class="cw-obs-rise cw-obs-d1 mt-12">
        <div class="flex items-center gap-3">
          <span class="h-px w-10 bg-[#c6f24e]"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.3em] text-[#c6f24e]/80">Company OS / 001</span>
        </div>
        <h1 class="mt-5 text-balance text-[43px] font-semibold leading-[0.98] tracking-[-0.04em]">{{ name }}</h1>
        <p class="mt-6 max-w-[19rem] text-[14.5px] leading-[1.85] text-white/55">{{ description }}</p>
      </section>

      <section class="cw-obs-rise cw-obs-d2 mt-9 grid grid-cols-3 border-y border-white/10">
        <div class="cw-obs-stat">
          <span class="font-mono text-[24px] font-medium tracking-tight text-[#c6f24e]">{{ pad2(features.length) }}</span>
          <span class="mt-1 font-mono text-[9px] uppercase tracking-[0.18em] text-white/35">Modules</span>
        </div>
        <div class="cw-obs-stat border-x border-white/10">
          <span class="font-mono text-[24px] font-medium tracking-tight text-[#c6f24e]">24H</span>
          <span class="mt-1 font-mono text-[9px] uppercase tracking-[0.18em] text-white/35">Response</span>
        </div>
        <div class="cw-obs-stat">
          <span class="font-mono text-[24px] font-medium tracking-tight text-[#c6f24e]">LIVE</span>
          <span class="mt-1 font-mono text-[9px] uppercase tracking-[0.18em] text-white/35">Status</span>
        </div>
      </section>

      <section class="cw-obs-rise cw-obs-d2 mt-11">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[13px] font-medium tracking-[0.08em]">核心业务矩阵</h2>
          <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-white/30">Capabilities</span>
        </div>
        <div v-if="features.length" class="mt-4 grid grid-cols-2 gap-2.5">
          <article v-for="(f, i) in features" :key="i" class="cw-obs-tile group">
            <span class="font-mono text-[10px] tracking-[0.12em] text-[#c6f24e]/60">{{ pad2(i + 1) }}</span>
            <p class="mt-6 text-[13.5px] font-medium leading-snug text-white/80">{{ f }}</p>
            <span class="mt-4 block h-px w-full bg-white/10 transition-colors duration-300 group-hover:bg-[#c6f24e]/40"></span>
          </article>
        </div>
        <p v-else class="mt-4 rounded-lg border border-dashed border-white/15 px-4 py-8 text-center font-mono text-[11px] text-white/30">// NO_CAPABILITIES_CONFIGURED</p>
      </section>

      <section class="cw-obs-rise cw-obs-d3 mt-11">
        <div class="flex items-baseline justify-between">
          <h2 class="text-[13px] font-medium tracking-[0.08em]">直接联系</h2>
          <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-white/30">Contact</span>
        </div>
        <div class="mt-3 border-t border-white/10">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-obs-contact">
            <span class="cw-obs-ic"><Phone class="h-3.5 w-3.5" /></span>
            <span class="cw-obs-key">电话</span>
            <span class="cw-obs-val tabular-nums">{{ contact.phone }}</span>
            <ArrowUpRight class="h-3.5 w-3.5 text-white/20" />
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-obs-contact">
            <span class="cw-obs-ic"><Mail class="h-3.5 w-3.5" /></span>
            <span class="cw-obs-key">邮箱</span>
            <span class="cw-obs-val">{{ contact.email }}</span>
            <ArrowUpRight class="h-3.5 w-3.5 text-white/20" />
          </a>
          <div v-if="contact.address" class="cw-obs-contact">
            <span class="cw-obs-ic"><MapPin class="h-3.5 w-3.5" /></span>
            <span class="cw-obs-key">地址</span>
            <span class="cw-obs-val whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-7 text-center font-mono text-[11px] text-white/30">// CONTACT_NOT_CONFIGURED</p>
        </div>
      </section>

      <div class="cw-obs-rise cw-obs-d3 mt-9">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-obs-cta">
          <span>进入官方网站</span>
          <ArrowUpRight class="h-4 w-4" />
        </a>
      </div>

      <footer class="mt-8 flex items-center justify-between border-t border-white/[0.07] pt-4 font-mono text-[9px] uppercase tracking-[0.2em] text-white/25">
        <span>{{ name }} · Official</span>
        <span>Est. Digital</span>
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
.cw-obs-grid {
  background-image:
    linear-gradient(rgba(198, 242, 78, 0.045) 1px, transparent 1px),
    linear-gradient(90deg, rgba(198, 242, 78, 0.045) 1px, transparent 1px);
  background-size: 34px 34px;
  -webkit-mask-image: linear-gradient(to bottom, black, transparent 90%);
  mask-image: linear-gradient(to bottom, black, transparent 90%);
}
.cw-obs-glow {
  background: radial-gradient(circle, rgba(198, 242, 78, 0.14), transparent 68%);
  filter: blur(8px);
}
.cw-obs-dot {
  animation: cw-obs-pulse 2.4s ease-out infinite;
}
@keyframes cw-obs-pulse {
  0% { box-shadow: 0 0 0 0 rgba(198, 242, 78, 0.5); }
  70% { box-shadow: 0 0 0 7px rgba(198, 242, 78, 0); }
  100% { box-shadow: 0 0 0 0 rgba(198, 242, 78, 0); }
}
.cw-obs-stat {
  display: flex;
  flex-direction: column;
  padding: 16px 6px;
  transition: background-color 0.25s ease;
}
.cw-obs-stat:hover {
  background: rgba(198, 242, 78, 0.04);
}
.cw-obs-tile {
  position: relative;
  display: flex;
  min-height: 104px;
  flex-direction: column;
  justify-content: space-between;
  border: 1px solid rgba(255, 255, 255, 0.09);
  border-radius: 6px;
  padding: 13px;
  background: linear-gradient(160deg, rgba(255, 255, 255, 0.035), transparent);
  transition: border-color 0.25s ease, transform 0.25s ease, background-color 0.25s ease;
}
.cw-obs-tile:hover {
  border-color: rgba(198, 242, 78, 0.45);
  background: rgba(198, 242, 78, 0.045);
  transform: translateY(-2px);
}
.cw-obs-contact {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.07);
  padding: 12px 0;
  transition: padding-left 0.25s ease;
}
.cw-obs-contact:last-child {
  border-bottom: 0;
}
a.cw-obs-contact:hover {
  padding-left: 4px;
}
.cw-obs-ic {
  display: flex;
  height: 27px;
  width: 27px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(198, 242, 78, 0.22);
  border-radius: 4px;
  color: #c6f24e;
  background: rgba(198, 242, 78, 0.05);
}
.cw-obs-key {
  flex-shrink: 0;
  font-size: 10px;
  letter-spacing: 0.08em;
  color: rgba(255, 255, 255, 0.3);
}
.cw-obs-val {
  margin-left: auto;
  max-width: 62%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  color: rgba(255, 255, 255, 0.72);
}
a.cw-obs-contact:hover .cw-obs-val {
  color: #c6f24e;
}
.cw-obs-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-radius: 5px;
  background: #c6f24e;
  padding: 14px;
  font-size: 13px;
  font-weight: 600;
  letter-spacing: 0.04em;
  color: #0a1004;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-obs-cta:hover {
  background: #d8ff74;
}
.cw-obs-cta:active {
  transform: scale(0.985);
}
.cw-obs-rise {
  animation: cw-obs-rise 0.7s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cw-obs-d1 { animation-delay: 0.07s; }
.cw-obs-d2 { animation-delay: 0.16s; }
.cw-obs-d3 { animation-delay: 0.25s; }
@keyframes cw-obs-rise {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-obs-rise,
  .cw-obs-dot {
    animation: none !important;
  }
  .cw-obs-tile,
  .cw-obs-contact,
  .cw-obs-cta {
    transition: none !important;
  }
}
</style>

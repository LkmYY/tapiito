<template>
  <div class="cw-cons relative min-h-screen overflow-hidden bg-[#e9dec8] text-[#141210]">
    <div class="cw-cons-wedge pointer-events-none absolute -right-24 top-0 h-[28rem] w-64"></div>
    <div class="cw-cons-thin pointer-events-none absolute left-8 top-0 h-full w-px bg-[#141210]/15"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-9 pt-6">
      <header class="cw-cons-in flex items-center justify-between">
        <span class="bg-[#141210] px-2 py-1 font-mono text-[8px] uppercase tracking-[0.18em] text-[#e9dec8]">Company / {{ nfcId || '01' }}</span>
        <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#141210]/50">Construct · 1924</span>
      </header>

      <section class="cw-cons-in cw-cons-d1 mt-10">
        <div class="flex items-start gap-3">
          <span class="cw-cons-vert font-mono text-[9px] uppercase tracking-[0.35em] text-[#b3261e]">Official<br />record</span>
          <div class="min-w-0 flex-1">
            <p class="font-mono text-[9px] uppercase tracking-[0.26em] text-[#141210]/55">The company, forward.</p>
            <h1 class="mt-2 break-words text-[52px] font-black uppercase leading-[0.82] tracking-[-0.055em]">{{ name }}</h1>
          </div>
        </div>
        <div class="mt-6 flex">
          <span class="h-2 w-16 bg-[#b3261e]"></span>
          <span class="h-2 flex-1 bg-[#141210]"></span>
        </div>
        <p class="mt-4 max-w-[20rem] text-[13px] leading-[1.8] text-[#141210]/72">{{ description }}</p>
      </section>

      <section class="cw-cons-in cw-cons-d2 mt-8">
        <div class="flex items-end justify-between border-b-2 border-[#141210] pb-1">
          <h2 class="text-[16px] font-black uppercase tracking-[0.06em]">Services / 服务</h2>
          <span class="font-mono text-[9px] text-[#141210]/45">{{ pad2(features.length) }}</span>
        </div>
        <div v-if="features.length" class="mt-1">
          <div v-for="(f, i) in features" :key="i" class="cw-cons-row group">
            <span class="cw-cons-num">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[14px] font-semibold leading-snug">{{ f }}</span>
            <span class="cw-cons-tri"></span>
          </div>
        </div>
      </section>

      <section class="cw-cons-in cw-cons-d2 mt-6 flex items-stretch">
        <div class="flex-1 bg-[#b3261e] px-4 py-3 text-[#f5ead6]">
          <p class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#f5ead6]/65">Service modules</p>
          <p class="mt-1 text-[25px] font-black leading-none">{{ pad2(features.length) }}</p>
        </div>
        <div class="flex-1 border-y-2 border-r-2 border-[#141210] px-4 py-3">
          <p class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#141210]/45">Reply time</p>
          <p class="mt-1 text-[25px] font-black leading-none">24H</p>
        </div>
      </section>

      <section class="cw-cons-in cw-cons-d3 mt-8">
        <div class="flex items-center gap-2">
          <h2 class="bg-[#141210] px-2 py-1 text-[10px] font-black uppercase tracking-[0.12em] text-[#e9dec8]">Contact</h2>
          <span class="h-0.5 flex-1 bg-[#141210]"></span>
        </div>
        <div class="mt-2">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-cons-contact">
            <span class="cw-cons-k">TEL</span>
            <span class="cw-cons-v tabular-nums">{{ contact.phone }}</span>
            <span class="cw-cons-arrow">→</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-cons-contact">
            <span class="cw-cons-k">MAIL</span>
            <span class="cw-cons-v">{{ contact.email }}</span>
            <span class="cw-cons-arrow">→</span>
          </a>
          <div v-if="contact.address" class="cw-cons-contact items-start">
            <span class="cw-cons-k">ADDR</span>
            <span class="cw-cons-v whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-5 text-center font-mono text-[10px] text-[#141210]/45">CONTACT / 未配置</p>
        </div>
      </section>

      <div class="cw-cons-in cw-cons-d3 mt-6">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-cons-cta">
          <span class="flex-1">访问官方网站</span>
          <span class="font-mono text-[14px]">↗</span>
        </a>
      </div>

      <footer class="cw-cons-in cw-cons-d3 mt-7 flex items-center justify-between border-t-2 border-[#141210] pt-2 font-mono text-[8px] uppercase tracking-[0.14em] text-[#141210]/45">
        <span>{{ name }}</span>
        <span>Forward, together</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

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
.cw-cons-wedge {
  background: linear-gradient(160deg, #b3261e 0 58%, transparent 58%);
  clip-path: polygon(38% 0, 100% 0, 100% 100%, 100% 72%);
  opacity: 0.95;
}
.cw-cons-vert {
  writing-mode: vertical-rl;
  transform: rotate(180deg);
  padding-top: 4px;
}
.cw-cons-row {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(20, 18, 16, 0.22);
  padding: 12px 0;
  transition: background-color 0.2s ease, padding-left 0.2s ease;
}
.cw-cons-row:hover {
  background: rgba(179, 38, 30, 0.06);
  padding-left: 4px;
}
.cw-cons-num {
  display: flex;
  height: 26px;
  width: 26px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  background: #141210;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 9px;
  color: #e9dec8;
  transition: background-color 0.2s ease, color 0.2s ease;
}
.cw-cons-row:hover .cw-cons-num {
  background: #b3261e;
  color: #fff3df;
}
.cw-cons-tri {
  width: 0;
  height: 0;
  flex-shrink: 0;
  border-top: 5px solid transparent;
  border-bottom: 5px solid transparent;
  border-left: 8px solid #b3261e;
  transition: transform 0.2s ease;
}
.cw-cons-row:hover .cw-cons-tri {
  transform: translateX(3px);
}
.cw-cons-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px solid rgba(20, 18, 16, 0.2);
  padding: 10px 0;
  transition: background-color 0.2s ease;
}
a.cw-cons-contact:hover {
  background: rgba(179, 38, 30, 0.06);
}
.cw-cons-k {
  width: 38px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #b3261e;
}
.cw-cons-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  font-weight: 500;
}
.cw-cons-arrow {
  font-weight: 900;
  color: rgba(20, 18, 16, 0.3);
  transition: transform 0.2s ease, color 0.2s ease;
}
a.cw-cons-contact:hover .cw-cons-arrow {
  transform: translateX(3px);
  color: #b3261e;
}
.cw-cons-cta {
  display: flex;
  align-items: center;
  background: #b3261e;
  padding: 13px 15px;
  font-size: 14px;
  font-weight: 900;
  letter-spacing: 0.06em;
  color: #f5ead6;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-cons-cta:hover {
  background: #141210;
}
.cw-cons-cta:active {
  transform: translateY(1px);
}
.cw-cons-in {
  animation: cw-cons-in 0.6s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cw-cons-d1 { animation-delay: 0.05s; }
.cw-cons-d2 { animation-delay: 0.14s; }
.cw-cons-d3 { animation-delay: 0.23s; }
@keyframes cw-cons-in {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-cons-in {
    animation: none !important;
  }
  .cw-cons-row,
  .cw-cons-num,
  .cw-cons-tri,
  .cw-cons-contact,
  .cw-cons-arrow,
  .cw-cons-cta {
    transition: none !important;
  }
}
</style>

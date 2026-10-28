<template>
  <div class="cw-term relative min-h-screen overflow-hidden bg-[#070604] text-[#ffb000]">
    <div class="cw-term-scan pointer-events-none absolute inset-0"></div>
    <div class="cw-term-vignette pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-4 pb-8 pt-6">
      <div class="cw-term-window">
        <div class="flex items-center gap-2 border-b border-[#ffb000]/25 bg-[#ffb000]/[0.06] px-3 py-2">
          <span class="h-2 w-2 rounded-full bg-[#7a5200]"></span>
          <span class="h-2 w-2 rounded-full bg-[#b07800]"></span>
          <span class="h-2 w-2 rounded-full bg-[#ffb000]"></span>
          <span class="ml-2 flex-1 truncate font-mono text-[9px] tracking-[0.08em] text-[#ffb000]/55">company-profile — bash — 80×24</span>
          <span v-if="nfcId" class="font-mono text-[8px] text-[#ffb000]/35">#{{ nfcId }}</span>
        </div>

        <div class="px-4 pb-7 pt-4 font-mono">
          <p class="cw-term-in text-[10px] text-[#ffb000]/40">Last login: today on tty01</p>
          <p class="cw-term-in cw-term-d1 mt-2 text-[11px]"><span class="text-[#ffb000]/50">guest@corp</span><span class="text-[#ffb000]/30">:~$</span> <span class="text-[#ffe7bd]">cat company.profile</span></p>

          <section class="cw-term-in cw-term-d1 mt-4">
            <div class="cw-term-rule text-[9px] text-[#ffb000]/30">┌─ IDENTITY ────────────────────────</div>
            <h1 class="mt-2 break-words text-[30px] font-bold leading-[1.05] tracking-[-0.02em] text-[#ffe0a3]">{{ name }}</h1>
            <p class="mt-2 text-[11.5px] leading-[1.8] text-[#ffb000]/70">{{ description }}</p>
            <div class="cw-term-rule mt-2 text-[9px] text-[#ffb000]/30">└───────────────────────────────────</div>
          </section>

          <section class="cw-term-in cw-term-d2 mt-4">
            <p class="text-[11px]"><span class="text-[#ffb000]/50">guest@corp</span><span class="text-[#ffb000]/30">:~$</span> <span class="text-[#ffe7bd]">ls -la services/</span></p>
            <p class="mt-2 text-[9px] uppercase tracking-[0.08em] text-[#ffb000]/35">total {{ features.length }} · core capabilities</p>
            <div v-if="features.length" class="mt-1">
              <div v-for="(f, i) in features" :key="i" class="cw-term-row group">
                <span class="text-[#ffb000]/45">{{ i === features.length - 1 ? '└──' : '├──' }}</span>
                <span class="w-4 text-[10px] text-[#ffb000]/35">{{ pad2(i + 1) }}</span>
                <span class="flex-1 text-[12px] text-[#ffe7bd]/90 transition-colors group-hover:text-[#fff6e3]">{{ f }}</span>
                <span class="text-[9px] text-[#7a5200] transition-colors group-hover:text-[#ffb000]">rw-r--r--</span>
              </div>
            </div>
            <p v-else class="mt-2 text-[11px] text-[#ffb000]/35">services/: no entries configured</p>
          </section>

          <section class="cw-term-in cw-term-d2 mt-4 grid grid-cols-3 gap-1.5">
            <div class="cw-term-cell">
              <p class="text-[18px] font-bold leading-none text-[#ffe0a3]">{{ pad2(features.length) }}</p>
              <p class="mt-1 text-[7.5px] uppercase tracking-[0.1em] text-[#ffb000]/35">services</p>
            </div>
            <div class="cw-term-cell">
              <p class="text-[18px] font-bold leading-none text-[#ffe0a3]">24h</p>
              <p class="mt-1 text-[7.5px] uppercase tracking-[0.1em] text-[#ffb000]/35">reply</p>
            </div>
            <div class="cw-term-cell">
              <p class="text-[18px] font-bold leading-none text-[#ffe0a3]">1:1</p>
              <p class="mt-1 text-[7.5px] uppercase tracking-[0.1em] text-[#ffb000]/35">support</p>
            </div>
          </section>

          <section class="cw-term-in cw-term-d3 mt-5">
            <p class="text-[11px]"><span class="text-[#ffb000]/50">guest@corp</span><span class="text-[#ffb000]/30">:~$</span> <span class="text-[#ffe7bd]">contact --all</span></p>
            <div class="mt-2 space-y-0.5 text-[11px]">
              <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-term-contact">
                <span class="text-[#ffb000]/40">phone</span>
                <span class="text-[#ffb000]/20">............</span>
                <span class="text-right text-[#ffe7bd] tabular-nums">{{ contact.phone }}</span>
              </a>
              <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-term-contact">
                <span class="text-[#ffb000]/40">email</span>
                <span class="text-[#ffb000]/20">............</span>
                <span class="truncate text-right text-[#ffe7bd]">{{ contact.email }}</span>
              </a>
              <div v-if="contact.address" class="cw-term-contact items-start">
                <span class="text-[#ffb000]/40">addr</span>
                <span class="text-[#ffb000]/20">.............</span>
                <span class="text-right text-[#ffb000]/75">{{ contact.address }}</span>
              </div>
              <p v-if="!hasContact" class="py-2 text-[10px] text-[#ffb000]/35">contact: not configured</p>
            </div>
          </section>

          <div class="cw-term-in cw-term-d3 mt-5">
            <p class="text-[11px]"><span class="text-[#ffb000]/50">guest@corp</span><span class="text-[#ffb000]/30">:~$</span> <span class="text-[#ffe7bd]">open --official</span></p>
            <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-term-cta mt-2">
              <span>[ 访问官方网站 ]</span>
              <span class="flex items-center gap-1">↗<span class="cw-term-cursor">_</span></span>
            </a>
          </div>

          <div class="mt-5 flex items-center justify-between border-t border-[#ffb000]/20 pt-2 text-[9px] text-[#ffb000]/30">
            <span>guest@corp:~$ <span class="cw-term-cursor">█</span></span>
            <span>{{ name }}</span>
          </div>
        </div>
      </div>
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
.cw-term-scan {
  background: repeating-linear-gradient(to bottom, rgba(255, 176, 0, 0.055) 0 1px, transparent 1px 4px);
  opacity: 0.6;
}
.cw-term-vignette {
  background: radial-gradient(120% 80% at 50% 20%, transparent 40%, rgba(0, 0, 0, 0.6) 100%);
}
.cw-term-window {
  border: 1px solid rgba(255, 176, 0, 0.28);
  background: #080705;
  box-shadow: 0 0 60px rgba(255, 176, 0, 0.07), inset 0 0 80px rgba(255, 176, 0, 0.035);
}
.cw-term-rule {
  overflow: hidden;
  white-space: nowrap;
  letter-spacing: 0.02em;
}
.cw-term-row {
  display: flex;
  align-items: center;
  gap: 6px;
  border-bottom: 1px dashed rgba(255, 176, 0, 0.1);
  padding: 5px 0;
  transition: background-color 0.2s ease;
}
.cw-term-row:hover {
  background: rgba(255, 176, 0, 0.055);
}
.cw-term-cell {
  border: 1px solid rgba(255, 176, 0, 0.2);
  padding: 8px;
  transition: background-color 0.2s ease;
}
.cw-term-cell:hover {
  background: rgba(255, 176, 0, 0.06);
}
.cw-term-contact {
  display: flex;
  align-items: baseline;
  gap: 4px;
  padding: 4px 2px;
  transition: background-color 0.2s ease;
}
a.cw-term-contact:hover {
  background: rgba(255, 176, 0, 0.07);
}
.cw-term-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border: 1px solid #ffb000;
  padding: 9px 11px;
  font-size: 11.5px;
  font-weight: 700;
  letter-spacing: 0.06em;
  transition: background-color 0.2s ease, color 0.2s ease;
}
.cw-term-cta:hover {
  background: #ffb000;
  color: #120d03;
}
.cw-term-cta:active {
  transform: translateY(1px);
}
.cw-term-cursor {
  animation: cw-term-blink 1.1s steps(1) infinite;
}
@keyframes cw-term-blink {
  50% { opacity: 0; }
}
.cw-term-in {
  animation: cw-term-in 0.5s cubic-bezier(0.2, 0.9, 0.2, 1) both;
}
.cw-term-d1 { animation-delay: 0.08s; }
.cw-term-d2 { animation-delay: 0.18s; }
.cw-term-d3 { animation-delay: 0.28s; }
@keyframes cw-term-in {
  from { opacity: 0; transform: translateX(-8px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-term-in,
  .cw-term-cursor {
    animation: none !important;
  }
  .cw-term-row,
  .cw-term-cell,
  .cw-term-contact,
  .cw-term-cta {
    transition: none !important;
  }
}
</style>

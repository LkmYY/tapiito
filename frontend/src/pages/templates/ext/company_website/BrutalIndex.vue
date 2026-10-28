<template>
  <div class="cw-br min-h-screen bg-white text-[#0a0a0a]">
    <div class="mx-auto w-full max-w-md pb-8">
      <header class="cw-br-in border-b-4 border-[#0a0a0a] px-5 py-3">
        <div class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.14em]">
          <span>Company Index / {{ nfcId || '001' }}</span>
          <span class="border border-[#0a0a0a] px-1.5 py-0.5">Open</span>
        </div>
      </header>

      <section class="cw-br-in cw-br-d1 px-5 pt-7">
        <p class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#0a0a0a]/45">A company, indexed.</p>
        <h1 class="mt-3 break-words text-[58px] font-black leading-[0.84] tracking-[-0.07em]">{{ name }}</h1>
        <p class="mt-5 max-w-[20rem] text-[13px] leading-[1.75] text-[#0a0a0a]/65">{{ description }}</p>
      </section>

      <section class="cw-br-in cw-br-d2 mt-7 border-y-2 border-[#0a0a0a] px-5 py-2">
        <div class="grid grid-cols-3 divide-x divide-[#0a0a0a]/25">
          <div class="pr-3">
            <p class="font-mono text-[20px] font-bold leading-none">{{ pad2(features.length) }}</p>
            <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#0a0a0a]/45">Services</p>
          </div>
          <div class="px-3">
            <p class="font-mono text-[20px] font-bold leading-none">24H</p>
            <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#0a0a0a]/45">Reply</p>
          </div>
          <div class="pl-3">
            <p class="font-mono text-[20px] font-bold leading-none">1:1</p>
            <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#0a0a0a]/45">Contact</p>
          </div>
        </div>
      </section>

      <section class="cw-br-in cw-br-d2 mt-7 px-5">
        <div class="flex items-baseline justify-between">
          <h2 class="font-mono text-[10px] font-bold uppercase tracking-[0.18em]">Index / Services</h2>
          <span class="font-mono text-[9px] text-[#0a0a0a]/40">01—{{ pad2(features.length) }}</span>
        </div>
        <div v-if="features.length" class="mt-2 border-t-2 border-[#0a0a0a]">
          <div v-for="(f, i) in features" :key="i" class="cw-br-row group">
            <span class="w-12 flex-shrink-0 font-mono text-[22px] font-black leading-none tracking-[-0.05em] text-[#0a0a0a]/22 transition-colors group-hover:text-[#0a0a0a]">{{ pad2(i + 1) }}</span>
            <span class="flex-1 border-l border-[#0a0a0a]/20 pl-3 text-[14px] font-semibold leading-snug">{{ f }}</span>
            <span class="font-mono text-[12px] text-[#0a0a0a]/25 transition-transform duration-200 group-hover:translate-x-1 group-hover:text-[#0a0a0a]">→</span>
          </div>
        </div>
      </section>

      <section class="cw-br-in cw-br-d3 mt-7 px-5">
        <div class="flex items-baseline justify-between border-b-2 border-[#0a0a0a] pb-1.5">
          <h2 class="font-mono text-[10px] font-bold uppercase tracking-[0.18em]">Index / Contact</h2>
          <span class="font-mono text-[9px] text-[#0a0a0a]/40">C—01</span>
        </div>
        <div>
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-br-contact">
            <span class="cw-br-key">TEL</span>
            <span class="cw-br-val tabular-nums">{{ contact.phone }}</span>
            <span class="cw-br-idx">01</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-br-contact">
            <span class="cw-br-key">MAIL</span>
            <span class="cw-br-val">{{ contact.email }}</span>
            <span class="cw-br-idx">02</span>
          </a>
          <div v-if="contact.address" class="cw-br-contact items-start">
            <span class="cw-br-key">ADDR</span>
            <span class="cw-br-val whitespace-normal">{{ contact.address }}</span>
            <span class="cw-br-idx">03</span>
          </div>
          <p v-if="!hasContact" class="py-6 text-center font-mono text-[10px] text-[#0a0a0a]/40">CONTACT INDEX NOT CONFIGURED</p>
        </div>
      </section>

      <div class="cw-br-in cw-br-d3 mt-6 px-5">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-br-cta">
          <span class="flex-1 text-[14px] font-black uppercase tracking-[0.08em]">访问官方网站</span>
          <span class="border-l border-white/30 pl-4 font-mono text-[10px]">WWW ↗</span>
        </a>
      </div>

      <footer class="cw-br-in cw-br-d3 mt-6 flex items-center justify-between border-t-4 border-[#0a0a0a] px-5 pt-2 font-mono text-[8px] uppercase tracking-[0.14em] text-[#0a0a0a]/40">
        <span>Index complete</span>
        <span>{{ name }} · Edition 01</span>
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
.cw-br-row {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(10, 10, 10, 0.18);
  padding: 13px 0;
  transition: background-color 0.18s ease;
}
.cw-br-row:hover {
  background: rgba(10, 10, 10, 0.035);
}
.cw-br-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px solid rgba(10, 10, 10, 0.18);
  padding: 10px 0;
  transition: background-color 0.18s ease;
}
a.cw-br-contact:hover {
  background: rgba(10, 10, 10, 0.04);
}
.cw-br-key {
  width: 40px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.12em;
}
.cw-br-val {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  font-weight: 500;
}
a.cw-br-contact:hover .cw-br-val {
  text-decoration: underline;
  text-underline-offset: 3px;
}
.cw-br-idx {
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 9px;
  color: rgba(10, 10, 10, 0.3);
}
.cw-br-cta {
  display: flex;
  align-items: center;
  background: #0a0a0a;
  padding: 13px 15px;
  color: #ffffff;
  transition: background-color 0.18s ease, transform 0.18s ease;
}
.cw-br-cta:hover {
  background: #333333;
}
.cw-br-cta:active {
  transform: translateY(1px);
}
.cw-br-in {
  animation: cw-br-in 0.55s cubic-bezier(0.16, 0.9, 0.2, 1) both;
}
.cw-br-d1 { animation-delay: 0.05s; }
.cw-br-d2 { animation-delay: 0.13s; }
.cw-br-d3 { animation-delay: 0.21s; }
@keyframes cw-br-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-br-in {
    animation: none !important;
  }
  .cw-br-row,
  .cw-br-contact,
  .cw-br-cta {
    transition: none !important;
  }
}
</style>

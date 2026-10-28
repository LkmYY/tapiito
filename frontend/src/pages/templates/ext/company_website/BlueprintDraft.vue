<template>
  <div class="cw-bp relative min-h-screen overflow-hidden bg-[#0a2745] text-[#dcecff]">
    <div class="cw-bp-grid pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-3 border border-[#dcecff]/20"></div>
    <div class="pointer-events-none absolute left-3 top-1/2 h-10 w-px -translate-y-1/2 bg-[#7fc4ff]/40"></div>
    <div class="pointer-events-none absolute right-3 top-1/3 h-10 w-px bg-[#7fc4ff]/40"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-10 pt-9">
      <header class="cw-bp-in flex items-start justify-between border-b border-[#dcecff]/25 pb-2 font-mono text-[8.5px] uppercase tracking-[0.16em] text-[#9fd0f5]/70">
        <div class="leading-relaxed">
          <p>SHEET 01 / 01</p>
          <p class="text-[#dcecff]/35">CORPORATE PROFILE</p>
        </div>
        <div class="text-right leading-relaxed">
          <p>DWG NO. {{ nfcId || 'CW-001' }}</p>
          <p class="text-[#dcecff]/35">SCALE 1:1 · REV. A</p>
        </div>
      </header>

      <section class="cw-bp-in cw-bp-d1 mt-9">
        <div class="flex items-center gap-3 font-mono text-[8.5px] uppercase tracking-[0.2em] text-[#7fc4ff]">
          <span class="h-1.5 w-1.5 border border-[#7fc4ff]"></span>
          <span>Company elevation / front view</span>
          <span class="h-px flex-1 bg-[#7fc4ff]/25"></span>
        </div>
        <h1 class="mt-6 text-[42px] font-medium leading-[0.98] tracking-[-0.035em] text-white">{{ name }}</h1>
        <div class="mt-5 flex items-start gap-3">
          <span class="mt-1 w-10 flex-shrink-0 border-t border-[#7fc4ff]/50"></span>
          <p class="text-[13px] leading-[1.85] text-[#c4ddf3]/75">{{ description }}</p>
        </div>
      </section>

      <section class="cw-bp-in cw-bp-d2 mt-9">
        <div class="flex items-baseline justify-between border-b border-[#dcecff]/25 pb-2">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#9fd0f5]">01 / Service schedule</h2>
          <span class="font-mono text-[8px] text-[#dcecff]/35">{{ pad2(features.length) }} ITEMS</span>
        </div>
        <div v-if="features.length" class="mt-2">
          <div v-for="(f, i) in features" :key="i" class="cw-bp-row group">
            <span class="font-mono text-[9px] text-[#7fc4ff]/70">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[13.5px] leading-snug text-[#e4f1ff]/90 transition-colors group-hover:text-white">{{ f }}</span>
            <span class="font-mono text-[8px] uppercase tracking-[0.12em] text-[#7fc4ff]/40">OK</span>
          </div>
        </div>
        <p v-else class="mt-3 border border-dashed border-[#9fd0f5]/25 py-8 text-center font-mono text-[10px] text-[#dcecff]/35">SERVICE SCHEDULE NOT CONFIGURED</p>
      </section>

      <section class="cw-bp-in cw-bp-d2 mt-8">
        <div class="grid grid-cols-[1.2fr_1fr] gap-3">
          <div class="border border-[#9fd0f5]/25 p-3">
            <p class="font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#7fc4ff]/70">Section A-A</p>
            <p class="mt-3 font-mono text-[28px] leading-none text-white">{{ pad2(features.length) }}<span class="text-[12px] text-[#7fc4ff]/60"> 项</span></p>
            <p class="mt-2 font-mono text-[8px] tracking-[0.08em] text-[#dcecff]/40">核心服务模块</p>
          </div>
          <div class="border border-[#9fd0f5]/25 p-3">
            <p class="font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#7fc4ff]/70">Detail B</p>
            <p class="mt-3 font-mono text-[28px] leading-none text-white">24<span class="text-[12px] text-[#7fc4ff]/60"> H</span></p>
            <p class="mt-2 font-mono text-[8px] tracking-[0.08em] text-[#dcecff]/40">工作日内响应</p>
          </div>
        </div>
      </section>

      <section class="cw-bp-in cw-bp-d3 mt-8">
        <div class="flex items-baseline justify-between border-b border-[#dcecff]/25 pb-2">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#9fd0f5]">02 / Contact schedule</h2>
          <span class="font-mono text-[8px] text-[#dcecff]/35">REV A</span>
        </div>
        <div class="mt-1">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-bp-contact">
            <span class="cw-bp-k">TEL</span>
            <span class="cw-bp-v tabular-nums">{{ contact.phone }}</span>
            <span class="cw-bp-plus">+</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-bp-contact">
            <span class="cw-bp-k">EMAIL</span>
            <span class="cw-bp-v">{{ contact.email }}</span>
            <span class="cw-bp-plus">+</span>
          </a>
          <div v-if="contact.address" class="cw-bp-contact items-start">
            <span class="cw-bp-k">ADDR</span>
            <span class="cw-bp-v whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-6 text-center font-mono text-[10px] text-[#dcecff]/35">CONTACT SCHEDULE NOT CONFIGURED</p>
        </div>
      </section>

      <div class="cw-bp-in cw-bp-d3 mt-8">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-bp-cta">
          <span class="font-mono text-[9px] tracking-[0.16em] text-[#0a2745]/55">LINK</span>
          <span class="flex-1 text-center">访问官方网站</span>
          <ArrowUpRight class="h-3.5 w-3.5" />
        </a>
      </div>

      <footer class="mt-8 grid grid-cols-3 border border-[#9fd0f5]/20 font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#9fd0f5]/50">
        <span class="border-r border-[#9fd0f5]/20 px-2 py-2">Drawn: CW</span>
        <span class="border-r border-[#9fd0f5]/20 px-2 py-2">Checked: OK</span>
        <span class="px-2 py-2">Sheet: 01</span>
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
.cw-bp-grid {
  background-image:
    linear-gradient(rgba(159, 208, 245, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(159, 208, 245, 0.1) 1px, transparent 1px),
    linear-gradient(rgba(159, 208, 245, 0.04) 1px, transparent 1px),
    linear-gradient(90deg, rgba(159, 208, 245, 0.04) 1px, transparent 1px);
  background-size: 40px 40px, 40px 40px, 8px 8px, 8px 8px;
}
.cw-bp-row {
  display: flex;
  align-items: center;
  gap: 11px;
  border-bottom: 1px dashed rgba(220, 236, 255, 0.18);
  padding: 11px 0;
  transition: background-color 0.2s ease, padding-left 0.2s ease;
}
.cw-bp-row:hover {
  padding-left: 3px;
  background: rgba(127, 196, 255, 0.05);
}
.cw-bp-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px dashed rgba(220, 236, 255, 0.18);
  padding: 10px 0;
  transition: background-color 0.2s ease;
}
a.cw-bp-contact:hover {
  background: rgba(127, 196, 255, 0.06);
}
.cw-bp-k {
  width: 42px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  letter-spacing: 0.12em;
  color: #7fc4ff;
}
.cw-bp-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 11.5px;
  color: rgba(228, 241, 255, 0.78);
}
a.cw-bp-contact:hover .cw-bp-v {
  color: #ffffff;
}
.cw-bp-plus {
  color: rgba(127, 196, 255, 0.4);
}
.cw-bp-cta {
  display: flex;
  align-items: center;
  gap: 10px;
  background: #bfe3ff;
  padding: 13px 14px;
  font-size: 13px;
  font-weight: 600;
  color: #08233d;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-bp-cta:hover {
  background: #ffffff;
}
.cw-bp-cta:active {
  transform: scale(0.985);
}
.cw-bp-in {
  animation: cw-bp-in 0.68s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-bp-d1 { animation-delay: 0.07s; }
.cw-bp-d2 { animation-delay: 0.16s; }
.cw-bp-d3 { animation-delay: 0.25s; }
@keyframes cw-bp-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-bp-in {
    animation: none !important;
  }
  .cw-bp-row,
  .cw-bp-contact,
  .cw-bp-cta {
    transition: none !important;
  }
}
</style>

<template>
  <div class="cw-mani min-h-screen bg-[#fbfbf9] text-[#101010]">
    <div class="mx-auto w-full max-w-md px-6 pb-9 pt-7">
      <header class="cw-mani-in flex items-center justify-between border-b border-[#101010]/20 pb-3">
        <span class="font-mono text-[8.5px] uppercase tracking-[0.28em] text-[#101010]/55">Manifesto · No. {{ nfcId || '01' }}</span>
        <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#101010]/35">A company in principles</span>
      </header>

      <section class="cw-mani-in cw-mani-d1 relative mt-12">
        <span class="cw-mani-watermark pointer-events-none absolute -right-2 -top-16 select-none font-black leading-none">01</span>
        <p class="font-mono text-[9px] uppercase tracking-[0.36em] text-[#101010]/45">We make things</p>
        <h1 class="relative mt-4 text-[46px] font-light leading-[1.02] tracking-[-0.045em]">{{ name }}</h1>
        <p class="relative mt-6 max-w-[19rem] text-[13.5px] leading-[1.9] text-[#101010]/62">{{ description }}</p>
      </section>

      <section class="cw-mani-in cw-mani-d2 mt-12">
        <div class="flex items-center justify-between border-b-2 border-[#101010] pb-2">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.26em]">What we stand for</h2>
          <span class="font-mono text-[9px] tabular-nums text-[#101010]/40">{{ pad2(features.length) }} principles</span>
        </div>
        <div v-if="features.length" class="cw-mani-timeline mt-3">
          <div v-for="(f, i) in features" :key="i" class="cw-mani-step group">
            <span class="cw-mani-node"></span>
            <span class="w-8 flex-shrink-0 font-mono text-[9px] tabular-nums text-[#101010]/35">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[14.5px] leading-snug text-[#101010]/82 transition-all duration-200 group-hover:translate-x-1 group-hover:text-[#101010]">{{ f }}</span>
            <span class="font-mono text-[8px] uppercase tracking-[0.12em] text-[#101010]/25">Principle</span>
          </div>
        </div>
        <p v-else class="mt-3 border-y border-dashed border-[#101010]/25 py-8 text-center font-mono text-[10px] text-[#101010]/40">NO PRINCIPLES CONFIGURED</p>
      </section>

      <section class="cw-mani-in cw-mani-d2 mt-9">
        <div class="grid grid-cols-[auto_1fr] items-end gap-4 border-y border-[#101010]/20 py-5">
          <p class="text-[64px] font-extralight leading-[0.8] tracking-[-0.06em]">{{ pad2(features.length) }}</p>
          <div class="pb-1">
            <p class="text-[11px] font-medium">服务模块，持续更新</p>
            <p class="mt-1 font-mono text-[8px] uppercase tracking-[0.16em] text-[#101010]/35">Delivered with intent · 24H reply</p>
          </div>
        </div>
      </section>

      <section class="cw-mani-in cw-mani-d3 mt-9">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.26em] text-[#101010]/45">Open channel</h2>
        <div class="mt-3 space-y-2.5">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-mani-contact">
            <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#101010]/35">Call</span>
            <span class="flex-1 text-[15px] font-light tabular-nums">{{ contact.phone }}</span>
            <span class="text-[13px] text-[#101010]/30">↗</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-mani-contact">
            <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#101010]/35">Write</span>
            <span class="flex-1 truncate text-[13px] font-light">{{ contact.email }}</span>
            <span class="text-[13px] text-[#101010]/30">↗</span>
          </a>
          <div v-if="contact.address" class="cw-mani-contact">
            <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#101010]/35">Visit</span>
            <span class="flex-1 text-[12px] font-light leading-snug">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-4 text-center font-mono text-[10px] text-[#101010]/35">CONTACT NOT CONFIGURED</p>
        </div>
      </section>

      <div class="cw-mani-in cw-mani-d3 mt-8">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-mani-cta">
          <span>访问官方网站</span>
          <ArrowUpRight class="h-4 w-4" />
        </a>
      </div>

      <footer class="cw-mani-in cw-mani-d3 mt-10 flex items-end justify-between border-t border-[#101010]/15 pt-3">
        <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#101010]/35">{{ name }} · Manifesto edition</span>
        <span class="font-mono text-[8px] text-[#101010]/25">01 / 01</span>
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
.cw-mani-watermark {
  font-size: 138px;
  letter-spacing: -0.08em;
  color: transparent;
  -webkit-text-stroke: 1px rgba(16, 16, 16, 0.11);
  line-height: 0.8;
}
.cw-mani-timeline {
  position: relative;
  padding-left: 7px;
}
.cw-mani-timeline::before {
  content: '';
  position: absolute;
  left: 10px;
  top: 8px;
  bottom: 14px;
  width: 1px;
  background: rgba(16, 16, 16, 0.18);
}
.cw-mani-step {
  position: relative;
  display: flex;
  align-items: center;
  gap: 12px;
  border-bottom: 1px solid rgba(16, 16, 16, 0.1);
  padding: 14px 0 14px 11px;
  transition: background-color 0.2s ease;
}
.cw-mani-step:hover {
  background: rgba(16, 16, 16, 0.025);
}
.cw-mani-node {
  position: absolute;
  left: -1px;
  top: 50%;
  height: 7px;
  width: 7px;
  border-radius: 99px;
  border: 1px solid rgba(16, 16, 16, 0.5);
  background: #fbfbf9;
  transform: translateY(-50%);
  transition: background-color 0.2s ease, border-color 0.2s ease;
}
.cw-mani-step:hover .cw-mani-node {
  background: #101010;
  border-color: #101010;
}
.cw-mani-contact {
  display: flex;
  align-items: baseline;
  gap: 12px;
  border-bottom: 1px solid rgba(16, 16, 16, 0.18);
  padding: 9px 0;
  transition: padding-left 0.2s ease;
}
a.cw-mani-contact:hover {
  padding-left: 4px;
}
.cw-mani-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border: 1px solid #101010;
  padding: 13px 15px;
  font-size: 13px;
  letter-spacing: 0.03em;
  transition: background-color 0.2s ease, color 0.2s ease, transform 0.2s ease;
}
.cw-mani-cta:hover {
  background: #101010;
  color: #fbfbf9;
}
.cw-mani-cta:active {
  transform: scale(0.99);
}
.cw-mani-in {
  animation: cw-mani-in 0.72s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-mani-d1 { animation-delay: 0.07s; }
.cw-mani-d2 { animation-delay: 0.16s; }
.cw-mani-d3 { animation-delay: 0.25s; }
@keyframes cw-mani-in {
  from { opacity: 0; transform: translateY(17px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-mani-in {
    animation: none !important;
  }
  .cw-mani-step,
  .cw-mani-node,
  .cw-mani-contact,
  .cw-mani-cta {
    transition: none !important;
  }
}
</style>

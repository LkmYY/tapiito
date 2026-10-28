<template>
  <div class="cw-let relative min-h-screen overflow-hidden bg-[#e8e0cb] px-3 py-5 text-[#1f3429]">
    <div class="cw-let-paper relative mx-auto w-full max-w-md px-7 pb-8 pt-9">
      <span class="cw-let-fiber pointer-events-none absolute inset-0"></span>

      <header class="cw-let-in relative flex items-start justify-between border-b border-[#1f4d3d]/30 pb-3">
        <div class="flex items-center gap-3">
          <span class="cw-let-mono flex h-10 w-10 items-center justify-center border border-[#1f4d3d]/45 font-serif text-[19px] text-[#1f4d3d]">{{ initial }}</span>
          <div>
            <p class="text-[12px] font-semibold tracking-[0.04em]">{{ name }}</p>
            <p class="mt-0.5 font-mono text-[7.5px] uppercase tracking-[0.24em] text-[#1f4d3d]/55">Established office</p>
          </div>
        </div>
        <div class="text-right font-mono text-[7px] uppercase leading-relaxed tracking-[0.16em] text-[#1f4d3d]/45">
          <p>Ref. {{ nfcId || 'CW-01' }}</p>
          <p>1st edition</p>
        </div>
      </header>

      <section class="cw-let-in cw-let-d1 relative mt-9">
        <p class="font-mono text-[8px] uppercase tracking-[0.32em] text-[#1f4d3d]/55">A letter of introduction</p>
        <h1 class="mt-3 font-serif text-[42px] leading-[1.06] tracking-[-0.02em]">{{ name }}</h1>
        <div class="mt-5 flex items-center gap-3">
          <span class="h-px w-12 bg-[#1f4d3d]/40"></span>
          <span class="font-serif text-[11px] italic text-[#1f4d3d]/60">at your service</span>
        </div>
        <p class="mt-5 text-[13.5px] leading-[1.95] text-[#1f3429]/80">{{ description }}</p>
      </section>

      <section class="cw-let-in cw-let-d2 relative mt-8">
        <div class="flex items-baseline justify-between border-b border-[#1f4d3d]/30 pb-1.5">
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.24em] text-[#1f4d3d]">Our practice</h2>
          <span class="font-mono text-[8px] text-[#1f4d3d]/40">{{ pad2(features.length) }} fields</span>
        </div>
        <div v-if="features.length" class="mt-1">
          <div v-for="(f, i) in features" :key="i" class="cw-let-row group">
            <span class="w-6 flex-shrink-0 font-serif text-[13px] italic text-[#1f4d3d]/55">{{ roman(i + 1) }}</span>
            <span class="flex-1 text-[14px] leading-snug transition-colors duration-200 group-hover:text-[#1f4d3d]">{{ f }}</span>
            <span class="font-mono text-[8px] uppercase tracking-[0.12em] text-[#1f4d3d]/30 transition-colors group-hover:text-[#1f4d3d]/70">included</span>
          </div>
        </div>
        <p v-else class="mt-3 border-y border-dashed border-[#1f4d3d]/25 py-8 text-center font-serif text-[13px] italic text-[#1f4d3d]/40">The practice is not yet listed.</p>
      </section>

      <section class="cw-let-in cw-let-d2 relative mt-7 grid grid-cols-2 gap-3">
        <div class="border border-[#1f4d3d]/25 p-3">
          <p class="font-serif text-[26px] leading-none text-[#1f4d3d]">{{ pad2(features.length) }}</p>
          <p class="mt-2 font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#1f4d3d]/55">Service disciplines</p>
        </div>
        <div class="border border-[#1f4d3d]/25 p-3">
          <p class="font-serif text-[26px] leading-none text-[#1f4d3d]">24<span class="text-[12px]">h</span></p>
          <p class="mt-2 font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#1f4d3d]/55">Reply on workdays</p>
        </div>
      </section>

      <section class="cw-let-in cw-let-d3 relative mt-8">
        <p class="font-mono text-[8px] uppercase tracking-[0.24em] text-[#1f4d3d]">Correspondence</p>
        <div class="mt-2 border-t border-[#1f4d3d]/30">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-let-line">
            <span class="cw-let-k">电话</span>
            <span class="cw-let-v tabular-nums">{{ contact.phone }}</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-let-line">
            <span class="cw-let-k">邮箱</span>
            <span class="cw-let-v">{{ contact.email }}</span>
          </a>
          <div v-if="contact.address" class="cw-let-line items-start">
            <span class="cw-let-k">地址</span>
            <span class="cw-let-v whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-4 font-serif text-[12px] italic text-[#1f4d3d]/40">Correspondence details are not configured.</p>
        </div>
      </section>

      <div class="cw-let-in cw-let-d3 relative mt-7 flex items-end justify-between gap-4">
        <div>
          <p class="font-serif text-[12px] italic text-[#1f4d3d]/60">Yours faithfully,</p>
          <p class="mt-1 font-serif text-[16px] text-[#1f3429]">{{ name }}</p>
        </div>
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-let-cta">
          Website <ArrowUpRight class="h-3 w-3" />
        </a>
      </div>

      <div class="cw-let-in cw-let-d3 relative mt-8 border-t border-[#1f4d3d]/25 pt-3 text-center">
        <p class="font-mono text-[7px] uppercase tracking-[0.26em] text-[#1f4d3d]/40">{{ name }} · Company correspondence · Printed on paper stock</p>
      </div>
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
const initial = computed(() => String(name.value).trim().charAt(0) || '企')
const pad2 = (n) => String(n).padStart(2, '0')
const roman = (n) => ['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII'][n - 1] || String(n)
</script>

<style scoped>
.cw-let-paper {
  background:
    linear-gradient(105deg, rgba(255, 255, 255, 0.35), transparent 35%),
    #f7f2e4;
  box-shadow: 0 18px 50px -28px rgba(31, 52, 41, 0.55), inset 0 0 0 1px rgba(31, 77, 61, 0.1);
}
.cw-let-fiber {
  opacity: 0.45;
  background-image:
    radial-gradient(rgba(31, 77, 61, 0.1) 0.4px, transparent 0.4px),
    radial-gradient(rgba(120, 90, 45, 0.09) 0.4px, transparent 0.4px);
  background-size: 5px 5px, 9px 9px;
  background-position: 0 0, 3px 4px;
  pointer-events: none;
}
.cw-let-mono {
  box-shadow: inset 0 0 0 3px #f7f2e4, inset 0 0 0 4px rgba(31, 77, 61, 0.25);
}
.cw-let-row {
  display: flex;
  align-items: baseline;
  gap: 11px;
  border-bottom: 1px solid rgba(31, 77, 61, 0.16);
  padding: 11px 0;
  transition: padding-left 0.22s ease, background-color 0.22s ease;
}
.cw-let-row:hover {
  padding-left: 3px;
  background: rgba(31, 77, 61, 0.035);
}
.cw-let-line {
  display: flex;
  align-items: baseline;
  gap: 11px;
  border-bottom: 1px solid rgba(31, 77, 61, 0.16);
  padding: 9px 0;
  transition: background-color 0.2s ease;
}
a.cw-let-line:hover {
  background: rgba(31, 77, 61, 0.04);
}
.cw-let-k {
  width: 34px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: rgba(31, 77, 61, 0.45);
}
.cw-let-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  color: rgba(31, 52, 41, 0.8);
  transition: color 0.2s ease;
}
a.cw-let-line:hover .cw-let-v {
  color: #1f4d3d;
  text-decoration: underline;
  text-underline-offset: 3px;
}
.cw-let-cta {
  display: inline-flex;
  flex-shrink: 0;
  align-items: center;
  gap: 5px;
  border-bottom: 1px solid #1f4d3d;
  padding-bottom: 2px;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #1f4d3d;
  transition: gap 0.2s ease, opacity 0.2s ease;
}
.cw-let-cta:hover {
  gap: 9px;
  opacity: 0.7;
}
.cw-let-cta:active {
  transform: translateY(1px);
}
.cw-let-in {
  animation: cw-let-in 0.72s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-let-d1 { animation-delay: 0.08s; }
.cw-let-d2 { animation-delay: 0.17s; }
.cw-let-d3 { animation-delay: 0.26s; }
@keyframes cw-let-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-let-in {
    animation: none !important;
  }
  .cw-let-row,
  .cw-let-line,
  .cw-let-v,
  .cw-let-cta {
    transition: none !important;
  }
}
</style>

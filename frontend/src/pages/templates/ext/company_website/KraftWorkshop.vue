<template>
  <div class="cw-kraft relative min-h-screen overflow-hidden bg-[#c9a170] text-[#342317]">
    <div class="cw-kraft-fiber pointer-events-none absolute inset-0"></div>
    <div class="cw-kraft-tape pointer-events-none absolute -top-3 left-1/2 h-7 w-32 -translate-x-1/2 -rotate-2"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-4 pb-8 pt-8">
      <div class="cw-kraft-docket">
        <div class="cw-kraft-perf pointer-events-none absolute inset-x-0 top-0"></div>

        <div class="px-6 pb-7 pt-8">
          <header class="cw-kraft-in flex items-start justify-between border-b border-[#342317]/30 pb-3">
            <div>
              <p class="font-mono text-[8px] uppercase tracking-[0.22em] text-[#5a412c]/60">Workshop docket</p>
              <p class="mt-1 font-serif text-[15px] font-bold tracking-wide">{{ name }}</p>
            </div>
            <div class="text-right">
              <p class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#5a412c]/55">No. {{ nfcId || '001' }}</p>
              <p class="mt-1 font-mono text-[8px] uppercase tracking-[0.14em] text-[#5a412c]/40">Job / Profile</p>
            </div>
          </header>

          <section class="cw-kraft-in cw-kraft-d1 mt-7">
            <div class="flex items-start gap-3">
              <div class="flex-1">
                <p class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#7a4f27]">Company brief</p>
                <h1 class="mt-3 break-words font-serif text-[38px] font-bold leading-[0.98] tracking-[-0.02em]">{{ name }}</h1>
              </div>
              <span class="cw-kraft-stamp flex h-16 w-16 flex-shrink-0 items-center justify-center border-2 border-[#8b3a2a] text-center font-mono text-[8px] font-bold uppercase leading-tight tracking-[0.08em] text-[#8b3a2a]">Approved<br />QC / 01</span>
            </div>
            <p class="mt-4 border-l-2 border-[#8b3a2a]/45 pl-3 text-[12.5px] leading-[1.8] text-[#342317]/75">{{ description }}</p>
          </section>

          <section class="cw-kraft-in cw-kraft-d2 mt-7">
            <div class="flex items-baseline justify-between border-b-2 border-[#342317]/60 pb-1">
              <h2 class="font-mono text-[9px] font-bold uppercase tracking-[0.18em]">Work items / 服务</h2>
              <span class="font-mono text-[9px] text-[#5a412c]/50">{{ pad2(features.length) }} pcs</span>
            </div>
            <div v-if="features.length" class="mt-1">
              <div v-for="(f, i) in features" :key="i" class="cw-kraft-item group">
                <span class="cw-kraft-check"></span>
                <span class="w-6 flex-shrink-0 font-mono text-[9px] text-[#5a412c]/50">{{ pad2(i + 1) }}</span>
                <span class="flex-1 text-[13.5px] font-medium leading-snug">{{ f }}</span>
                <span class="font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#5a412c]/40">Done</span>
              </div>
            </div>
          </section>

          <section class="cw-kraft-in cw-kraft-d2 mt-6 grid grid-cols-3 border border-[#342317]/35">
            <div class="border-r border-[#342317]/25 p-2.5">
              <p class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#5a412c]/50">Items</p>
              <p class="mt-1 font-serif text-[20px] font-bold leading-none">{{ pad2(features.length) }}</p>
            </div>
            <div class="border-r border-[#342317]/25 p-2.5">
              <p class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#5a412c]/50">Reply</p>
              <p class="mt-1 font-serif text-[20px] font-bold leading-none">24H</p>
            </div>
            <div class="p-2.5">
              <p class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#5a412c]/50">Craft</p>
              <p class="mt-1 font-serif text-[20px] font-bold leading-none">1:1</p>
            </div>
          </section>

          <section class="cw-kraft-in cw-kraft-d3 mt-7">
            <div class="flex items-center gap-2">
              <h2 class="font-mono text-[9px] font-bold uppercase tracking-[0.18em]">Delivery details</h2>
              <span class="h-px flex-1 bg-[#342317]/25"></span>
            </div>
            <div class="mt-2">
              <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-kraft-contact">
                <span class="cw-kraft-k">TEL</span>
                <span class="cw-kraft-v tabular-nums">{{ contact.phone }}</span>
                <span class="cw-kraft-go">↗</span>
              </a>
              <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-kraft-contact">
                <span class="cw-kraft-k">MAIL</span>
                <span class="cw-kraft-v">{{ contact.email }}</span>
                <span class="cw-kraft-go">↗</span>
              </a>
              <div v-if="contact.address" class="cw-kraft-contact items-start">
                <span class="cw-kraft-k">ADDR</span>
                <span class="cw-kraft-v whitespace-normal">{{ contact.address }}</span>
              </div>
              <p v-if="!hasContact" class="py-5 text-center font-mono text-[9px] text-[#5a412c]/40">DETAILS / 未配置</p>
            </div>
          </section>

          <div class="cw-kraft-in cw-kraft-d3 mt-6">
            <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-kraft-cta">
              <span class="flex-1">访问官方网站</span>
              <span class="border-l border-[#f1e3cd]/30 pl-3 font-mono text-[9px]">SHIP ↗</span>
            </a>
          </div>

          <footer class="cw-kraft-in cw-kraft-d3 mt-7 flex items-end justify-between border-t border-[#342317]/25 pt-2">
            <p class="font-serif text-[11px] italic text-[#5a412c]/60">Handmade with care.</p>
            <p class="font-mono text-[7.5px] uppercase tracking-[0.12em] text-[#5a412c]/40">{{ name }} · Workshop</p>
          </footer>
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
.cw-kraft {
  background:
    radial-gradient(90% 70% at 20% 0%, rgba(236, 205, 160, 0.55), transparent 60%),
    linear-gradient(145deg, #cfa878, #bd925e 55%, #c69c6a);
}
.cw-kraft-fiber {
  opacity: 0.4;
  background-image:
    repeating-linear-gradient(0deg, rgba(90, 60, 30, 0.1) 0 1px, transparent 1px 4px),
    repeating-linear-gradient(90deg, rgba(255, 240, 210, 0.12) 0 1px, transparent 1px 6px);
  pointer-events: none;
}
.cw-kraft-tape {
  background: rgba(232, 215, 180, 0.65);
  box-shadow: 0 1px 2px rgba(52, 35, 23, 0.18);
  backdrop-filter: blur(1px);
}
.cw-kraft-docket {
  position: relative;
  background: #e8d3ad;
  box-shadow: 0 30px 55px -32px rgba(52, 35, 23, 0.75), inset 0 0 0 1px rgba(52, 35, 23, 0.15);
}
.cw-kraft-perf {
  height: 7px;
  background-image: radial-gradient(circle at 4px 0, transparent 3.5px, #c9a170 3.6px);
  background-size: 12px 7px;
  opacity: 0.85;
}
.cw-kraft-stamp {
  transform: rotate(-8deg);
  border-radius: 3px;
  opacity: 0.78;
  box-shadow: inset 0 0 0 1px rgba(139, 58, 42, 0.45);
}
.cw-kraft-item {
  display: flex;
  align-items: center;
  gap: 8px;
  border-bottom: 1px dashed rgba(52, 35, 23, 0.28);
  padding: 10px 0;
  transition: background-color 0.2s ease;
}
.cw-kraft-item:hover {
  background: rgba(139, 58, 42, 0.05);
}
.cw-kraft-check {
  height: 12px;
  width: 12px;
  flex-shrink: 0;
  border: 1.5px solid rgba(52, 35, 23, 0.55);
  border-radius: 2px;
  background:
    linear-gradient(45deg, transparent 42%, #8b3a2a 42% 54%, transparent 54%),
    linear-gradient(-45deg, transparent 46%, #8b3a2a 46% 56%, transparent 56%);
  background-size: 70% 70%;
  background-position: center;
  background-repeat: no-repeat;
}
.cw-kraft-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px dashed rgba(52, 35, 23, 0.28);
  padding: 9px 0;
}
.cw-kraft-k {
  width: 40px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #8b3a2a;
}
.cw-kraft-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  font-weight: 500;
}
.cw-kraft-go {
  color: rgba(52, 35, 23, 0.35);
  transition: transform 0.2s ease, color 0.2s ease;
}
a.cw-kraft-contact:hover .cw-kraft-go {
  transform: translate(2px, -2px);
  color: #8b3a2a;
}
a.cw-kraft-contact:hover .cw-kraft-v {
  color: #8b3a2a;
}
.cw-kraft-cta {
  display: flex;
  align-items: center;
  gap: 10px;
  background: #342317;
  padding: 12px 14px;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 11px;
  letter-spacing: 0.1em;
  color: #f1e3cd;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-kraft-cta:hover {
  background: #8b3a2a;
}
.cw-kraft-cta:active {
  transform: translateY(1px);
}
.cw-kraft-in {
  animation: cw-kraft-in 0.7s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cw-kraft-d1 { animation-delay: 0.07s; }
.cw-kraft-d2 { animation-delay: 0.16s; }
.cw-kraft-d3 { animation-delay: 0.25s; }
@keyframes cw-kraft-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-kraft-in {
    animation: none !important;
  }
  .cw-kraft-item,
  .cw-kraft-contact,
  .cw-kraft-go,
  .cw-kraft-cta {
    transition: none !important;
  }
}
</style>

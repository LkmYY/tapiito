<template>
  <div class="cw-sand relative min-h-screen overflow-hidden bg-[#ece4d4] text-[#4a3a2a]">
    <div class="cw-sand-sun pointer-events-none absolute -right-16 -top-16 h-56 w-56 rounded-full"></div>
    <svg class="pointer-events-none absolute -left-10 top-1/3 h-40 w-40 text-[#b08968]/35" viewBox="0 0 100 100" fill="none" aria-hidden="true">
      <circle cx="50" cy="50" r="48" stroke="currentColor" stroke-width="0.6" />
      <circle cx="50" cy="50" r="36" stroke="currentColor" stroke-width="0.6" />
      <path d="M2 50h96M50 2v96" stroke="currentColor" stroke-width="0.6" />
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-10 pt-8">
      <header class="cw-sand-in flex items-center justify-between">
        <span class="flex items-center gap-2">
          <span class="flex h-7 w-7 items-center justify-center rounded-full border border-[#4a3a2a]/30 font-serif text-[12px] italic">{{ initial }}</span>
          <span class="text-[10px] font-medium uppercase tracking-[0.2em] text-[#4a3a2a]/55">Official</span>
        </span>
        <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#4a3a2a]/40">N° {{ nfcId || '—' }}</span>
      </header>

      <section class="cw-sand-in cw-sand-d1 mt-16">
        <p class="font-mono text-[8.5px] uppercase tracking-[0.38em] text-[#8a6a4f]">A company profile</p>
        <h1 class="mt-5 text-[46px] font-light leading-[1.02] tracking-[-0.03em] text-[#3c2e20]">{{ name }}</h1>
        <p class="mt-6 max-w-[19rem] text-[13.5px] font-light leading-[1.95] text-[#4a3a2a]/65">{{ description }}</p>
      </section>

      <div class="cw-sand-in cw-sand-d1 mt-10 flex items-center gap-4">
        <span class="h-px flex-1 bg-[#4a3a2a]/20"></span>
        <span class="h-1 w-1 rounded-full bg-[#8a6a4f]"></span>
        <span class="font-mono text-[8px] uppercase tracking-[0.26em] text-[#4a3a2a]/40">Selected services</span>
        <span class="h-1 w-1 rounded-full bg-[#8a6a4f]"></span>
        <span class="h-px flex-1 bg-[#4a3a2a]/20"></span>
      </div>

      <section class="cw-sand-in cw-sand-d2 mt-8">
        <div v-if="features.length" class="space-y-0">
          <div v-for="(f, i) in features" :key="i" class="cw-sand-row group">
            <span class="font-mono text-[9px] tabular-nums text-[#8a6a4f]/60">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[15px] font-light leading-snug tracking-[0.01em] text-[#3c2e20]/85 transition-all duration-300 group-hover:translate-x-1 group-hover:text-[#3c2e20]">{{ f }}</span>
            <span class="h-px w-4 bg-[#4a3a2a]/25 transition-all duration-300 group-hover:w-7 group-hover:bg-[#8a6a4f]"></span>
          </div>
        </div>
      </section>

      <section class="cw-sand-in cw-sand-d2 mt-10 grid grid-cols-3 border-y border-[#4a3a2a]/15 py-5">
        <div>
          <p class="text-[17px] font-light leading-none text-[#3c2e20]">{{ pad2(features.length) }}</p>
          <p class="mt-2 font-mono text-[7px] uppercase tracking-[0.16em] text-[#4a3a2a]/40">Services</p>
        </div>
        <div class="border-l border-[#4a3a2a]/15 pl-4">
          <p class="text-[17px] font-light leading-none text-[#3c2e20]">24H</p>
          <p class="mt-2 font-mono text-[7px] uppercase tracking-[0.16em] text-[#4a3a2a]/40">Response</p>
        </div>
        <div class="border-l border-[#4a3a2a]/15 pl-4">
          <p class="text-[17px] font-light leading-none text-[#3c2e20]">1:1</p>
          <p class="mt-2 font-mono text-[7px] uppercase tracking-[0.16em] text-[#4a3a2a]/40">Care</p>
        </div>
      </section>

      <section class="cw-sand-in cw-sand-d3 mt-10">
        <h2 class="font-mono text-[8px] uppercase tracking-[0.32em] text-[#8a6a4f]">Get in touch</h2>
        <div class="mt-3">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-sand-contact">
            <span class="cw-sand-v tabular-nums">{{ contact.phone }}</span>
            <span class="cw-sand-k">电话</span>
            <ArrowUpRight class="h-3 w-3" />
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-sand-contact">
            <span class="cw-sand-v">{{ contact.email }}</span>
            <span class="cw-sand-k">邮箱</span>
            <ArrowUpRight class="h-3 w-3" />
          </a>
          <div v-if="contact.address" class="cw-sand-contact">
            <span class="cw-sand-v whitespace-normal">{{ contact.address }}</span>
            <span class="cw-sand-k">地址</span>
          </div>
          <p v-if="!hasContact" class="py-5 text-[11px] font-light text-[#4a3a2a]/45">联系方式未配置</p>
        </div>
      </section>

      <div class="cw-sand-in cw-sand-d3 mt-9">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-sand-cta group">
          <span class="flex-1">访问官方网站</span>
          <span class="flex h-8 w-8 items-center justify-center rounded-full border border-[#4a3a2a]/35 transition-colors group-hover:bg-[#3c2e20]">
            <ArrowUpRight class="h-3.5 w-3.5" />
          </span>
        </a>
      </div>

      <footer class="cw-sand-in cw-sand-d3 mt-12 flex items-center justify-between border-t border-[#4a3a2a]/12 pt-3 font-mono text-[7.5px] uppercase tracking-[0.18em] text-[#4a3a2a]/35">
        <span>{{ name }}</span>
        <span>Quietly confident · Since always</span>
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
const initial = computed(() => String(name.value).trim().charAt(0) || '企')
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.cw-sand-sun {
  background: radial-gradient(circle at 35% 35%, #d9b892, #c8a27a 55%, transparent 72%);
  opacity: 0.45;
  filter: blur(2px);
}
.cw-sand-row {
  display: flex;
  align-items: center;
  gap: 14px;
  border-bottom: 1px solid rgba(74, 58, 42, 0.12);
  padding: 14px 0;
}
.cw-sand-row:first-child {
  border-top: 1px solid rgba(74, 58, 42, 0.12);
}
.cw-sand-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px solid rgba(74, 58, 42, 0.12);
  padding: 11px 0;
}
.cw-sand-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 13px;
  font-weight: 400;
  letter-spacing: 0.01em;
  color: rgba(60, 46, 32, 0.85);
  transition: color 0.2s ease;
}
.cw-sand-k {
  flex-shrink: 0;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: rgba(74, 58, 42, 0.35);
}
a.cw-sand-contact:hover .cw-sand-v {
  color: #8a6a4f;
}
a.cw-sand-contact:hover svg {
  color: #8a6a4f;
}
.cw-sand-cta {
  display: flex;
  align-items: center;
  gap: 10px;
  border-top: 1px solid rgba(74, 58, 42, 0.35);
  border-bottom: 1px solid rgba(74, 58, 42, 0.35);
  padding: 13px 0;
  font-size: 12px;
  letter-spacing: 0.06em;
  transition: padding-left 0.25s ease, background-color 0.25s ease;
}
.cw-sand-cta:hover {
  padding-left: 5px;
  background: rgba(176, 137, 104, 0.07);
}
.cw-sand-cta:active {
  transform: translateY(1px);
}
.cw-sand-in {
  animation: cw-sand-in 0.85s cubic-bezier(0.2, 0.7, 0.2, 1) both;
}
.cw-sand-d1 { animation-delay: 0.08s; }
.cw-sand-d2 { animation-delay: 0.18s; }
.cw-sand-d3 { animation-delay: 0.28s; }
@keyframes cw-sand-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-sand-in {
    animation: none !important;
  }
  .cw-sand-row,
  .cw-sand-contact,
  .cw-sand-cta,
  .cw-sand-row span {
    transition: none !important;
  }
}
</style>

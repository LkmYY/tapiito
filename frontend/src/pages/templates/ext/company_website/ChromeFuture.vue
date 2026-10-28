<template>
  <div class="cw-chrome relative min-h-screen overflow-hidden bg-[#d8dce2] text-[#14161a]">
    <div class="cw-chrome-sheen pointer-events-none absolute inset-0"></div>
    <div class="cw-chrome-orb pointer-events-none absolute -right-20 top-16 h-64 w-64 rounded-full"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-9 pt-6">
      <header class="cw-chrome-in flex items-center justify-between border-b border-[#14161a]/25 pb-2.5">
        <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-[#14161a]/60">Material / Company</span>
        <span class="flex items-center gap-2 font-mono text-[8.5px] uppercase tracking-[0.14em] text-[#14161a]/45">
          Spec {{ nfcId || '01' }}
          <span class="h-1.5 w-1.5 rounded-full bg-[#14161a]"></span>
        </span>
      </header>

      <section class="cw-chrome-in cw-chrome-d1 mt-8">
        <div class="flex items-start justify-between gap-4">
          <p class="font-mono text-[8.5px] uppercase tracking-[0.28em] text-[#14161a]/50">Official surface</p>
          <span class="cw-chrome-badge">CN / 001</span>
        </div>
        <h1 class="cw-chrome-name mt-4 break-words text-[46px] font-black leading-[0.9] tracking-[-0.055em]">{{ name }}</h1>
        <p class="mt-5 max-w-[20rem] text-[13.5px] leading-[1.85] text-[#14161a]/65">{{ description }}</p>
      </section>

      <section class="cw-chrome-in cw-chrome-d1 mt-7 grid grid-cols-[1fr_auto] items-center gap-4 border-y border-[#14161a]/25 py-3">
        <div class="flex items-baseline gap-2">
          <span class="text-[30px] font-light leading-none tracking-[-0.04em]">{{ pad2(features.length) }}</span>
          <span class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#14161a]/45">Service<br />capabilities</span>
        </div>
        <div class="flex items-center gap-1.5">
          <span v-for="n in 5" :key="n" class="cw-chrome-bar" :style="{ height: 6 + n * 5 + 'px' }"></span>
        </div>
      </section>

      <section class="cw-chrome-in cw-chrome-d2 mt-8">
        <div class="flex items-baseline justify-between">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.22em]">Service specification</h2>
          <span class="font-mono text-[8px] text-[#14161a]/40">REV. 01</span>
        </div>
        <div v-if="features.length" class="mt-2 border-t border-[#14161a]/30">
          <div v-for="(f, i) in features" :key="i" class="cw-chrome-row group">
            <span class="w-7 font-mono text-[9px] text-[#14161a]/35">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[13.5px] font-medium leading-snug">{{ f }}</span>
            <span class="cw-chrome-tag">SPEC</span>
          </div>
        </div>
        <p v-else class="mt-3 border border-dashed border-[#14161a]/25 py-8 text-center font-mono text-[10px] text-[#14161a]/40">NO SERVICE SPECIFICATION</p>
      </section>

      <section class="cw-chrome-in cw-chrome-d2 mt-8">
        <div class="overflow-hidden rounded-xl border border-[#14161a]/20 bg-[#eef0f3]/70">
          <div class="flex items-center justify-between border-b border-[#14161a]/15 px-4 py-2.5">
            <span class="font-mono text-[8.5px] uppercase tracking-[0.2em] text-[#14161a]/55">Contact terminal</span>
            <span class="flex items-center gap-1.5 font-mono text-[8px] text-[#14161a]/45"><span class="h-1.5 w-1.5 rounded-full bg-[#2f9e6b]"></span>online</span>
          </div>
          <div class="px-4 py-1">
            <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-chrome-contact">
              <span class="cw-chrome-k">PHONE</span>
              <span class="cw-chrome-v tabular-nums">{{ contact.phone }}</span>
              <span class="cw-chrome-go">↗</span>
            </a>
            <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-chrome-contact">
              <span class="cw-chrome-k">EMAIL</span>
              <span class="cw-chrome-v">{{ contact.email }}</span>
              <span class="cw-chrome-go">↗</span>
            </a>
            <div v-if="contact.address" class="cw-chrome-contact items-start">
              <span class="cw-chrome-k">ADDRESS</span>
              <span class="cw-chrome-v whitespace-normal">{{ contact.address }}</span>
            </div>
            <p v-if="!hasContact" class="py-5 text-center font-mono text-[10px] text-[#14161a]/40">CONTACT NOT CONFIGURED</p>
          </div>
        </div>
      </section>

      <div class="cw-chrome-in cw-chrome-d3 mt-6">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-chrome-cta">
          <span class="flex-1 text-center">访问官方网站</span>
          <span class="flex h-7 w-7 items-center justify-center rounded-full bg-white/15"><ArrowUpRight class="h-3.5 w-3.5" /></span>
        </a>
      </div>

      <footer class="cw-chrome-in cw-chrome-d3 mt-7 flex items-center justify-between border-t border-[#14161a]/15 pt-3 font-mono text-[8px] uppercase tracking-[0.16em] text-[#14161a]/40">
        <span>{{ name }} · Material profile</span>
        <span>Edition / 01</span>
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
.cw-chrome {
  background:
    linear-gradient(145deg, #e8ebef 0%, #cfd4db 40%, #dfe3e8 62%, #c3c9d1 100%);
}
.cw-chrome-sheen {
  background:
    linear-gradient(115deg, transparent 20%, rgba(255, 255, 255, 0.55) 33%, transparent 45%),
    linear-gradient(180deg, rgba(255, 255, 255, 0.25), transparent 30%);
  mix-blend-mode: screen;
  opacity: 0.8;
}
.cw-chrome-orb {
  background:
    radial-gradient(circle at 30% 25%, #ffffff, #dfe4ea 32%, #9fa7b3 58%, #eef1f5 75%, #7e8794 100%);
  box-shadow: inset -18px -22px 45px rgba(60, 68, 80, 0.35), 0 30px 60px -25px rgba(20, 24, 30, 0.4);
  opacity: 0.7;
}
.cw-chrome-name {
  background: linear-gradient(180deg, #0e1013 0%, #2a2e35 45%, #6d7681 70%, #16181c 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  filter: drop-shadow(0 1px 0 rgba(255, 255, 255, 0.55));
}
.cw-chrome-badge {
  border: 1px solid rgba(20, 22, 26, 0.35);
  border-radius: 999px;
  padding: 3px 7px;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 7.5px;
  letter-spacing: 0.12em;
}
.cw-chrome-bar {
  width: 4px;
  border-radius: 1px;
  background: linear-gradient(to top, rgba(20, 22, 26, 0.25), rgba(20, 22, 26, 0.75));
  transition: height 0.3s ease, background 0.3s ease;
}
.cw-chrome-bar:hover {
  background: #14161a;
}
.cw-chrome-row {
  display: flex;
  align-items: center;
  gap: 11px;
  border-bottom: 1px solid rgba(20, 22, 26, 0.16);
  padding: 11px 2px;
  transition: padding-left 0.2s ease, background-color 0.2s ease;
}
.cw-chrome-row:hover {
  padding-left: 5px;
  background: rgba(255, 255, 255, 0.4);
}
.cw-chrome-tag {
  flex-shrink: 0;
  border: 1px solid rgba(20, 22, 26, 0.22);
  padding: 1px 4px;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 7px;
  letter-spacing: 0.1em;
  color: rgba(20, 22, 26, 0.45);
}
.cw-chrome-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px solid rgba(20, 22, 26, 0.12);
  padding: 10px 0;
}
.cw-chrome-contact:last-child {
  border-bottom: 0;
}
.cw-chrome-k {
  width: 52px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 7.5px;
  letter-spacing: 0.12em;
  color: rgba(20, 22, 26, 0.4);
}
.cw-chrome-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  font-weight: 500;
}
.cw-chrome-go {
  color: rgba(20, 22, 26, 0.3);
  transition: transform 0.2s ease, color 0.2s ease;
}
a.cw-chrome-contact:hover .cw-chrome-go {
  transform: translate(2px, -2px);
  color: #14161a;
}
a.cw-chrome-contact:hover .cw-chrome-v {
  text-decoration: underline;
  text-underline-offset: 2px;
}
.cw-chrome-cta {
  display: flex;
  align-items: center;
  gap: 10px;
  border-radius: 999px;
  background: #14161a;
  padding: 9px 10px 9px 18px;
  font-size: 13px;
  font-weight: 600;
  color: #eef0f3;
  box-shadow: 0 12px 24px -14px rgba(20, 22, 26, 0.8);
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-chrome-cta:hover {
  background: #343941;
}
.cw-chrome-cta:active {
  transform: scale(0.985);
}
.cw-chrome-in {
  animation: cw-chrome-in 0.72s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-chrome-d1 { animation-delay: 0.07s; }
.cw-chrome-d2 { animation-delay: 0.16s; }
.cw-chrome-d3 { animation-delay: 0.25s; }
@keyframes cw-chrome-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-chrome-in {
    animation: none !important;
  }
  .cw-chrome-row,
  .cw-chrome-contact,
  .cw-chrome-go,
  .cw-chrome-cta,
  .cw-chrome-bar {
    transition: none !important;
  }
}
</style>

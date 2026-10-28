<template>
  <div class="cw-deco relative min-h-screen overflow-hidden bg-[#f0e5cd] text-[#171411]">
    <div class="cw-deco-rays pointer-events-none absolute inset-x-0 top-0 h-72"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-9 pt-7">
      <header class="cw-deco-in flex items-center justify-between border-b border-[#171411]/30 pb-2">
        <span class="font-mono text-[8px] uppercase tracking-[0.3em] text-[#8b6a24]">Maison d'entreprise</span>
        <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#171411]/45">MMXXVI · {{ nfcId || '01' }}</span>
      </header>

      <section class="cw-deco-in cw-deco-d1 mt-8 text-center">
        <svg class="mx-auto h-14 w-32 text-[#a9852f]" viewBox="0 0 160 70" fill="none" aria-hidden="true">
          <g stroke="currentColor" stroke-width="1">
            <path d="M80 66V20" />
            <path d="M80 66 40 20M80 66l40-46" />
            <path d="M80 66 18 34M80 66l62-32" />
            <path d="M80 66 25 58M80 66l55-8" />
            <path d="M80 66 50 66M80 66h30" />
            <circle cx="80" cy="70" r="2.5" fill="currentColor" stroke="none" />
            <path d="M28 20h104" />
          </g>
        </svg>
        <p class="mt-3 font-mono text-[8px] uppercase tracking-[0.4em] text-[#8b6a24]">Company of distinction</p>
        <h1 class="mt-4 font-serif text-[40px] leading-[1.05] tracking-[0.01em]">{{ name }}</h1>
        <div class="mt-5 flex items-center justify-center gap-3">
          <span class="h-px w-10 bg-[#a9852f]/55"></span>
          <span class="h-1.5 w-1.5 rotate-45 border border-[#a9852f]"></span>
          <span class="h-px w-10 bg-[#a9852f]/55"></span>
        </div>
        <p class="mx-auto mt-5 max-w-[19rem] text-[12.5px] leading-[1.9] text-[#171411]/65">{{ description }}</p>
      </section>

      <section class="cw-deco-in cw-deco-d2 mt-8">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-[#171411]/25"></span>
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.3em] text-[#8b6a24]">Services</h2>
          <span class="h-px flex-1 bg-[#171411]/25"></span>
        </div>
        <div v-if="features.length" class="mt-3 grid grid-cols-2 gap-x-4">
          <div v-for="(f, i) in features" :key="i" class="cw-deco-cell group">
            <span class="font-serif text-[12px] italic text-[#a9852f]">{{ pad2(i + 1) }}</span>
            <p class="mt-2 text-[13px] leading-snug tracking-[0.02em] text-[#171411]/85 transition-colors group-hover:text-[#8b6a24]">{{ f }}</p>
          </div>
        </div>
        <p v-else class="mt-3 border border-dashed border-[#171411]/25 py-8 text-center font-serif text-[12px] italic text-[#171411]/40">Services are not yet listed.</p>
      </section>

      <section class="cw-deco-in cw-deco-d2 mt-8 border-y-2 border-[#171411]/70 py-3">
        <div class="grid grid-cols-3 divide-x divide-[#171411]/25 text-center">
          <div>
            <p class="font-serif text-[22px] leading-none text-[#8b6a24]">{{ pad2(features.length) }}</p>
            <p class="mt-1.5 font-mono text-[7px] uppercase tracking-[0.16em] text-[#171411]/45">Services</p>
          </div>
          <div>
            <p class="font-serif text-[22px] leading-none text-[#8b6a24]">24H</p>
            <p class="mt-1.5 font-mono text-[7px] uppercase tracking-[0.16em] text-[#171411]/45">Reply</p>
          </div>
          <div>
            <p class="font-serif text-[22px] leading-none text-[#8b6a24]">1:1</p>
            <p class="mt-1.5 font-mono text-[7px] uppercase tracking-[0.16em] text-[#171411]/45">Attention</p>
          </div>
        </div>
      </section>

      <section class="cw-deco-in cw-deco-d3 mt-8">
        <div class="flex items-center justify-center gap-3">
          <span class="h-px w-8 bg-[#a9852f]/50"></span>
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.3em] text-[#8b6a24]">Correspondence</h2>
          <span class="h-px w-8 bg-[#a9852f]/50"></span>
        </div>
        <div class="mt-3">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-deco-contact">
            <span class="cw-deco-key"><i>T</i></span>
            <span class="cw-deco-val tabular-nums">{{ contact.phone }}</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-deco-contact">
            <span class="cw-deco-key"><i>E</i></span>
            <span class="cw-deco-val">{{ contact.email }}</span>
          </a>
          <div v-if="contact.address" class="cw-deco-contact items-start">
            <span class="cw-deco-key"><i>A</i></span>
            <span class="cw-deco-val whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-5 text-center font-serif text-[12px] italic text-[#171411]/40">Contact details are not configured.</p>
        </div>
      </section>

      <div class="cw-deco-in cw-deco-d3 mt-7">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-deco-cta">
          <span class="text-[10px] tracking-[0.24em]">访问官方网站</span>
          <ArrowUpRight class="h-3.5 w-3.5" />
        </a>
      </div>

      <footer class="cw-deco-in cw-deco-d3 mt-7 text-center">
        <p class="font-mono text-[7px] uppercase tracking-[0.3em] text-[#171411]/35">{{ name }} · Established in good form</p>
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
.cw-deco-rays {
  background:
    repeating-conic-gradient(from 250deg at 50% -12%, rgba(169, 133, 47, 0.16) 0deg 1.2deg, transparent 1.2deg 7deg);
  -webkit-mask-image: radial-gradient(60% 90% at 50% 0%, black, transparent 78%);
  mask-image: radial-gradient(60% 90% at 50% 0%, black, transparent 78%);
  opacity: 0.8;
}
.cw-deco-cell {
  border-top: 1px solid rgba(23, 20, 17, 0.24);
  padding: 9px 0 12px;
  transition: padding-left 0.22s ease, background-color 0.22s ease;
}
.cw-deco-cell:nth-child(odd) {
  border-right: 1px solid rgba(23, 20, 17, 0.12);
  padding-right: 10px;
}
.cw-deco-cell:nth-child(even) {
  padding-left: 10px;
}
.cw-deco-cell:hover {
  background: rgba(169, 133, 47, 0.055);
}
.cw-deco-contact {
  display: flex;
  align-items: center;
  gap: 12px;
  border-bottom: 1px solid rgba(23, 20, 17, 0.16);
  padding: 9px 8px;
  transition: background-color 0.22s ease, padding-left 0.22s ease;
}
a.cw-deco-contact:hover {
  background: rgba(169, 133, 47, 0.08);
  padding-left: 13px;
}
.cw-deco-key {
  display: flex;
  height: 20px;
  width: 20px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(169, 133, 47, 0.5);
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  color: #8b6a24;
  transform: rotate(45deg);
}
.cw-deco-key i {
  display: block;
  font-style: normal;
  transform: rotate(-45deg);
}
.cw-deco-val {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: center;
  font-size: 12px;
  letter-spacing: 0.02em;
  color: rgba(23, 20, 17, 0.78);
  transition: color 0.2s ease;
}
a.cw-deco-contact:hover .cw-deco-val {
  color: #8b6a24;
}
.cw-deco-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 9px;
  border: 1px solid #171411;
  background: #171411;
  padding: 13px;
  color: #f0e5cd;
  transition: background-color 0.22s ease, color 0.22s ease, transform 0.2s ease;
}
.cw-deco-cta:hover {
  background: #a9852f;
  border-color: #a9852f;
  color: #171411;
}
.cw-deco-cta:active {
  transform: scale(0.985);
}
.cw-deco-in {
  animation: cw-deco-in 0.8s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-deco-d1 { animation-delay: 0.08s; }
.cw-deco-d2 { animation-delay: 0.17s; }
.cw-deco-d3 { animation-delay: 0.26s; }
@keyframes cw-deco-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-deco-in {
    animation: none !important;
  }
  .cw-deco-cell,
  .cw-deco-contact,
  .cw-deco-val,
  .cw-deco-cta {
    transition: none !important;
  }
}
</style>

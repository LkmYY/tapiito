<template>
  <div class="cw-noir relative min-h-screen overflow-hidden bg-[#0b0b0a] text-[#e9e4d8]">
    <div class="cw-noir-grain pointer-events-none absolute inset-0"></div>
    <div class="cw-noir-spot pointer-events-none absolute inset-x-0 top-0 h-[30rem]"></div>
    <div class="cw-noir-strip pointer-events-none absolute bottom-0 left-3 top-0 w-3"></div>
    <div class="cw-noir-strip pointer-events-none absolute bottom-0 right-3 top-0 w-3"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-8 pb-8 pt-7">
      <header class="cw-noir-in flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.24em] text-[#e9e4d8]/40">
        <span>A company picture</span>
        <span>Reel {{ nfcId || '01' }}</span>
      </header>

      <section class="cw-noir-in cw-noir-d1 mt-14 text-center">
        <p class="font-mono text-[8px] uppercase tracking-[0.42em] text-[#d9a85f]">Starring</p>
        <h1 class="mt-4 font-serif text-[46px] leading-[0.98] tracking-[-0.01em] text-[#f2ecdf]">{{ name }}</h1>
        <p class="mt-4 font-mono text-[8px] uppercase tracking-[0.3em] text-[#e9e4d8]/40">A portrait in one act</p>
        <div class="mt-5 flex items-center justify-center gap-2">
          <span class="h-px w-8 bg-[#d9a85f]/50"></span>
          <span class="h-1 w-1 rotate-45 bg-[#d9a85f]"></span>
          <span class="h-px w-8 bg-[#d9a85f]/50"></span>
        </div>
        <p class="mx-auto mt-5 max-w-[18rem] font-serif text-[13px] italic leading-[1.85] text-[#e9e4d8]/60">{{ description }}</p>
      </section>

      <section class="cw-noir-in cw-noir-d2 mt-12">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-[#e9e4d8]/15"></span>
          <h2 class="font-mono text-[8px] uppercase tracking-[0.32em] text-[#d9a85f]">Cast / Services</h2>
          <span class="h-px flex-1 bg-[#e9e4d8]/15"></span>
        </div>
        <div v-if="features.length" class="mt-4">
          <div v-for="(f, i) in features" :key="i" class="cw-noir-row group">
            <span class="w-5 font-mono text-[9px] tabular-nums text-[#e9e4d8]/30">{{ pad2(i + 1) }}</span>
            <span class="font-serif text-[16px] leading-snug text-[#e9e4d8]/85 transition-colors group-hover:text-[#f2ecdf]">{{ f }}</span>
            <span class="flex-1 self-end border-b border-dotted border-[#e9e4d8]/20 pb-1"></span>
            <span class="font-mono text-[7.5px] uppercase tracking-[0.14em] text-[#d9a85f]/60">Act {{ i + 1 }}</span>
          </div>
        </div>
        <p v-else class="mt-4 border-y border-dashed border-[#e9e4d8]/15 py-8 text-center font-serif text-[12px] italic text-[#e9e4d8]/35">The cast has not been announced.</p>
      </section>

      <section class="cw-noir-in cw-noir-d2 mt-10">
        <div class="flex items-center justify-between border-y border-[#e9e4d8]/15 py-4">
          <div>
            <p class="font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#e9e4d8]/35">Runtime</p>
            <p class="mt-1 font-serif text-[19px] text-[#f2ecdf]">{{ pad2(features.length) }} Chapters</p>
          </div>
          <div class="text-center">
            <p class="font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#e9e4d8]/35">Format</p>
            <p class="mt-1 font-serif text-[19px] text-[#f2ecdf]">24H</p>
          </div>
          <div class="text-right">
            <p class="font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#e9e4d8]/35">Rating</p>
            <p class="mt-1 font-serif text-[19px] text-[#f2ecdf]">1:1</p>
          </div>
        </div>
      </section>

      <section class="cw-noir-in cw-noir-d3 mt-10">
        <div class="text-center">
          <h2 class="font-mono text-[8px] uppercase tracking-[0.34em] text-[#d9a85f]">Production office</h2>
        </div>
        <div class="mt-3">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-noir-contact">
            <span class="cw-noir-k">TEL</span>
            <span class="cw-noir-v tabular-nums">{{ contact.phone }}</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-noir-contact">
            <span class="cw-noir-k">MAIL</span>
            <span class="cw-noir-v">{{ contact.email }}</span>
          </a>
          <div v-if="contact.address" class="cw-noir-contact items-start">
            <span class="cw-noir-k">ADDR</span>
            <span class="cw-noir-v whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-5 text-center font-serif text-[12px] italic text-[#e9e4d8]/35">The office is not listed.</p>
        </div>
      </section>

      <div class="cw-noir-in cw-noir-d3 mt-8">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-noir-cta">
          <span>访问官方网站</span>
          <span class="font-mono text-[8px] tracking-[0.14em] text-[#d9a85f]">WATCH ↗</span>
        </a>
      </div>

      <footer class="cw-noir-in cw-noir-d3 mt-7 text-center font-mono text-[7px] uppercase tracking-[0.26em] text-[#e9e4d8]/25">
        {{ name }} · A film in one act · The end
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
.cw-noir-grain {
  opacity: 0.38;
  background-image:
    repeating-radial-gradient(circle at 20% 30%, rgba(255, 255, 255, 0.08) 0 0.4px, transparent 0.5px 2.4px),
    repeating-radial-gradient(circle at 70% 65%, rgba(255, 255, 255, 0.06) 0 0.4px, transparent 0.5px 3.1px);
  pointer-events: none;
}
.cw-noir-spot {
  background: radial-gradient(55% 70% at 50% 0%, rgba(217, 168, 95, 0.16), transparent 70%);
}
.cw-noir-strip {
  background-image: repeating-linear-gradient(to bottom, rgba(233, 228, 216, 0.16) 0 7px, transparent 7px 17px);
  opacity: 0.7;
}
.cw-noir-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 0;
  border-bottom: 1px solid rgba(233, 228, 216, 0.1);
  transition: background-color 0.2s ease;
}
.cw-noir-row:hover {
  background: rgba(217, 168, 95, 0.04);
}
.cw-noir-contact {
  display: flex;
  align-items: baseline;
  gap: 12px;
  border-bottom: 1px solid rgba(233, 228, 216, 0.12);
  padding: 9px 0;
  transition: background-color 0.2s ease;
}
a.cw-noir-contact:hover {
  background: rgba(217, 168, 95, 0.05);
}
.cw-noir-k {
  width: 40px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 7.5px;
  letter-spacing: 0.18em;
  color: #d9a85f;
}
.cw-noir-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: center;
  font-family: 'Times New Roman', Times, serif;
  font-size: 13px;
  letter-spacing: 0.04em;
  color: rgba(233, 228, 216, 0.75);
  transition: color 0.2s ease;
}
a.cw-noir-contact:hover .cw-noir-v {
  color: #f2ecdf;
}
.cw-noir-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border: 1px solid rgba(217, 168, 95, 0.55);
  padding: 12px 14px;
  font-family: 'Times New Roman', Times, serif;
  font-size: 13px;
  letter-spacing: 0.06em;
  color: #f2ecdf;
  transition: background-color 0.25s ease, color 0.25s ease;
}
.cw-noir-cta:hover {
  background: #d9a85f;
  color: #14110c;
}
.cw-noir-cta:active {
  transform: translateY(1px);
}
.cw-noir-in {
  animation: cw-noir-in 0.85s cubic-bezier(0.2, 0.7, 0.2, 1) both;
}
.cw-noir-d1 { animation-delay: 0.08s; }
.cw-noir-d2 { animation-delay: 0.18s; }
.cw-noir-d3 { animation-delay: 0.28s; }
@keyframes cw-noir-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-noir-in {
    animation: none !important;
  }
  .cw-noir-row,
  .cw-noir-contact,
  .cw-noir-v,
  .cw-noir-cta {
    transition: none !important;
  }
}
</style>

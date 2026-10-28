<template>
  <div class="cw-news relative min-h-screen bg-[#e9e4d7] text-[#15130f]">
    <div class="cw-news-grain pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-9 pt-5">
      <header class="cw-news-in border-b border-[#15130f]/40 pb-1.5 text-center">
        <p class="font-mono text-[7.5px] uppercase tracking-[0.26em] text-[#15130f]/55">Morning edition · Company affairs</p>
        <div class="mt-1 flex items-center justify-between font-mono text-[7.5px] uppercase tracking-[0.12em] text-[#15130f]/50">
          <span>Vol. {{ nfcId || '001' }}</span>
          <span>Price · One tap</span>
        </div>
      </header>

      <div class="cw-news-in cw-news-d1 mt-2">
        <p class="cw-news-masthead border-y-[3px] border-[#15130f] py-1.5 text-center font-serif text-[13px] font-bold uppercase tracking-[0.3em]">The Daily Company</p>
      </div>

      <section class="cw-news-in cw-news-d1 mt-5">
        <div class="flex items-center gap-2 font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#a51c1c]">
          <span class="h-1.5 w-1.5 bg-[#a51c1c]"></span>
          <span>Exclusive profile</span>
          <span class="h-px flex-1 bg-[#15130f]/20"></span>
          <span class="text-[#15130f]/45">Page A1</span>
        </div>
        <h1 class="mt-3 text-[46px] font-black leading-[0.88] tracking-[-0.055em]">{{ name }}</h1>
        <p class="mt-4 border-l-2 border-[#a51c1c] pl-3 font-serif text-[14px] italic leading-[1.65] text-[#15130f]/75">{{ description }}</p>
      </section>

      <section class="cw-news-in cw-news-d2 mt-5 border-y border-[#15130f]/25 py-3">
        <div class="grid grid-cols-3 gap-0 text-center">
          <div>
            <p class="font-serif text-[23px] font-bold leading-none">{{ pad2(features.length) }}</p>
            <p class="mt-1 font-mono text-[7px] uppercase tracking-[0.12em] text-[#15130f]/50">Service desks</p>
          </div>
          <div class="border-x border-[#15130f]/20">
            <p class="font-serif text-[23px] font-bold leading-none">24H</p>
            <p class="mt-1 font-mono text-[7px] uppercase tracking-[0.12em] text-[#15130f]/50">Newsroom reply</p>
          </div>
          <div>
            <p class="font-serif text-[23px] font-bold leading-none">1:1</p>
            <p class="mt-1 font-mono text-[7px] uppercase tracking-[0.12em] text-[#15130f]/50">Correspondent</p>
          </div>
        </div>
      </section>

      <section class="cw-news-in cw-news-d2 mt-5">
        <div class="flex items-end justify-between border-b-2 border-[#15130f] pb-1">
          <h2 class="font-serif text-[17px] font-bold uppercase tracking-[0.04em]">Inside this company</h2>
          <span class="font-mono text-[7.5px] uppercase tracking-[0.12em] text-[#15130f]/45">Services</span>
        </div>
        <div v-if="features.length" class="mt-1">
          <div v-for="(f, i) in features" :key="i" class="cw-news-row group">
            <span class="w-5 flex-shrink-0 font-serif text-[13px] font-bold text-[#a51c1c]">{{ pad2(i + 1) }}</span>
            <span class="flex-1 font-serif text-[14.5px] leading-snug transition-colors duration-200 group-hover:text-[#a51c1c]">{{ f }}</span>
            <span class="hidden flex-1 self-end border-b border-dotted border-[#15130f]/30 pb-1 sm:block"></span>
            <span class="font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#15130f]/35">A{{ i + 2 }}</span>
          </div>
        </div>
        <p v-else class="mt-3 border-b border-dashed border-[#15130f]/30 pb-6 text-center font-serif text-[12px] italic text-[#15130f]/45">This column awaits its first entry.</p>
      </section>

      <section class="cw-news-in cw-news-d3 mt-5 grid grid-cols-[1.15fr_1fr] gap-3">
        <div class="border border-[#15130f]/40 p-3">
          <p class="border-b border-[#15130f]/30 pb-1 font-mono text-[7px] uppercase tracking-[0.16em] text-[#a51c1c]">Classified · Contact</p>
          <div class="mt-1">
            <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-news-contact">
              <span class="font-mono text-[7px] uppercase text-[#15130f]/45">Tel</span>
              <span class="text-[11.5px] font-semibold tabular-nums">{{ contact.phone }}</span>
            </a>
            <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-news-contact">
              <span class="font-mono text-[7px] uppercase text-[#15130f]/45">Mail</span>
              <span class="break-all text-[10.5px]">{{ contact.email }}</span>
            </a>
            <div v-if="contact.address" class="cw-news-contact">
              <span class="font-mono text-[7px] uppercase text-[#15130f]/45">Add</span>
              <span class="text-[10.5px] leading-snug">{{ contact.address }}</span>
            </div>
            <p v-if="!hasContact" class="py-3 text-center font-serif text-[10px] italic text-[#15130f]/40">Not configured.</p>
          </div>
        </div>
        <div class="flex flex-col justify-between border border-[#15130f]/40 p-3">
          <div>
            <p class="border-b border-[#15130f]/30 pb-1 font-mono text-[7px] uppercase tracking-[0.16em] text-[#15130f]/55">Editor's note</p>
            <p class="mt-2 font-serif text-[11px] leading-[1.6] text-[#15130f]/65">For partnerships, interviews and service enquiries, our desk is open on working days.</p>
          </div>
          <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-news-cta mt-3">
            阅读线上版 <span>↗</span>
          </a>
        </div>
      </section>

      <footer class="cw-news-in cw-news-d3 mt-6 flex items-center justify-between border-t-2 border-[#15130f] pt-1.5 font-mono text-[7px] uppercase tracking-[0.16em] text-[#15130f]/45">
        <span>{{ name }} · All rights reserved</span>
        <span>Printed daily</span>
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
.cw-news-grain {
  opacity: 0.5;
  background-image: radial-gradient(rgba(21, 19, 15, 0.12) 0.45px, transparent 0.45px);
  background-size: 3px 3px;
  pointer-events: none;
}
.cw-news-masthead {
  font-family: 'Times New Roman', Times, serif;
}
.cw-news-row {
  display: flex;
  align-items: center;
  gap: 8px;
  border-bottom: 1px solid rgba(21, 19, 15, 0.18);
  padding: 9px 0;
  transition: background-color 0.2s ease;
}
.cw-news-row:hover {
  background: rgba(165, 28, 28, 0.04);
}
.cw-news-contact {
  display: flex;
  align-items: baseline;
  gap: 6px;
  border-bottom: 1px dotted rgba(21, 19, 15, 0.28);
  padding: 7px 0;
}
.cw-news-contact > span:first-child {
  width: 22px;
  flex-shrink: 0;
}
a.cw-news-contact:hover > span:last-child {
  color: #a51c1c;
  text-decoration: underline;
  text-underline-offset: 2px;
}
.cw-news-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #15130f;
  padding: 8px 9px;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 9px;
  letter-spacing: 0.06em;
  color: #e9e4d7;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-news-cta:hover {
  background: #a51c1c;
}
.cw-news-cta:active {
  transform: translateY(1px);
}
.cw-news-in {
  animation: cw-news-in 0.65s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-news-d1 { animation-delay: 0.06s; }
.cw-news-d2 { animation-delay: 0.15s; }
.cw-news-d3 { animation-delay: 0.24s; }
@keyframes cw-news-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-news-in {
    animation: none !important;
  }
  .cw-news-row,
  .cw-news-contact,
  .cw-news-cta {
    transition: none !important;
  }
}
</style>

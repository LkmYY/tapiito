<template>
  <div class="cw-gold relative min-h-screen overflow-hidden bg-[#0d0a06] text-[#efe3cb]">
    <div class="cw-gold-aura pointer-events-none absolute inset-x-0 top-0 h-[26rem]"></div>
    <div class="cw-gold-noise pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-9 pt-9">
      <div class="cw-gold-frame">
        <span class="cw-gold-corner cw-gold-tl"></span>
        <span class="cw-gold-corner cw-gold-tr"></span>
        <span class="cw-gold-corner cw-gold-bl"></span>
        <span class="cw-gold-corner cw-gold-br"></span>

        <div class="px-6 pb-8 pt-7">
          <header class="cw-gold-in flex items-center justify-between">
            <span class="font-serif text-[12px] italic tracking-[0.18em] text-[#c9a563]">Maison</span>
            <span class="font-mono text-[8.5px] uppercase tracking-[0.3em] text-[#efe3cb]/40">N° {{ nfcId || '001' }}</span>
          </header>

          <section class="cw-gold-in cw-gold-d1 mt-10 text-center">
            <div class="flex items-center justify-center gap-3">
              <span class="h-px w-8 bg-[#c9a563]/45"></span>
              <span class="h-1 w-1 rotate-45 bg-[#c9a563]"></span>
              <span class="h-px w-8 bg-[#c9a563]/45"></span>
            </div>
            <p class="mt-5 font-mono text-[8.5px] uppercase tracking-[0.4em] text-[#c9a563]/80">House of Enterprise</p>
            <h1 class="mt-4 font-serif text-[41px] leading-[1.08] tracking-[-0.01em]">{{ name }}</h1>
            <p class="mx-auto mt-5 max-w-[18rem] text-[13px] leading-[1.9] text-[#efe3cb]/55">{{ description }}</p>
          </section>

          <section class="cw-gold-in cw-gold-d2 mt-10">
            <div class="flex items-center gap-3">
              <span class="h-px flex-1 bg-[#c9a563]/22"></span>
              <h2 class="font-mono text-[8.5px] uppercase tracking-[0.32em] text-[#c9a563]">Savoir-Faire</h2>
              <span class="h-px flex-1 bg-[#c9a563]/22"></span>
            </div>
            <div v-if="features.length" class="mt-2">
              <article v-for="(f, i) in features" :key="i" class="cw-gold-row group">
                <span class="w-7 flex-shrink-0 font-serif text-[13px] italic text-[#c9a563]/70">{{ roman(i + 1) }}</span>
                <span class="flex-1 text-[14px] leading-snug tracking-[0.01em] text-[#efe3cb]/85">{{ f }}</span>
                <span class="font-serif text-[15px] text-[#c9a563]/30 transition-all duration-300 group-hover:translate-x-1 group-hover:text-[#c9a563]">—</span>
              </article>
            </div>
            <p v-else class="mt-3 border-y border-[#c9a563]/15 py-8 text-center font-serif text-[13px] italic text-[#efe3cb]/35">Services are not yet configured.</p>
          </section>

          <section class="cw-gold-in cw-gold-d2 mt-9 grid grid-cols-3 border-y border-[#c9a563]/18">
            <div class="py-4 text-center">
              <p class="font-serif text-[19px] text-[#c9a563]">{{ pad2(features.length) }}</p>
              <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#efe3cb]/35">Services</p>
            </div>
            <div class="border-x border-[#c9a563]/18 py-4 text-center">
              <p class="font-serif text-[19px] text-[#c9a563]">24H</p>
              <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#efe3cb]/35">Concierge</p>
            </div>
            <div class="py-4 text-center">
              <p class="font-serif text-[19px] text-[#c9a563]">1:1</p>
              <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#efe3cb]/35">Service</p>
            </div>
          </section>

          <section class="cw-gold-in cw-gold-d3 mt-9">
            <div class="text-center">
              <h2 class="font-serif text-[21px] italic text-[#efe3cb]">Enquire</h2>
              <p class="mt-1 font-mono text-[8px] uppercase tracking-[0.3em] text-[#efe3cb]/35">Private &amp; corporate clients</p>
            </div>
            <div class="mt-4">
              <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-gold-contact">
                <span class="cw-gold-label">Téléphone</span>
                <span class="cw-gold-value tabular-nums">{{ contact.phone }}</span>
              </a>
              <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-gold-contact">
                <span class="cw-gold-label">Courriel</span>
                <span class="cw-gold-value">{{ contact.email }}</span>
              </a>
              <div v-if="contact.address" class="cw-gold-contact items-start">
                <span class="cw-gold-label">Adresse</span>
                <span class="cw-gold-value whitespace-normal">{{ contact.address }}</span>
              </div>
              <p v-if="!hasContact" class="py-6 text-center font-serif text-[12px] italic text-[#efe3cb]/35">Contact details are not yet configured.</p>
            </div>
          </section>

          <div class="cw-gold-in cw-gold-d3 mt-8">
            <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-gold-cta">
              <span>访问官方网站</span>
              <ArrowUpRight class="h-3.5 w-3.5" />
            </a>
          </div>

          <footer class="mt-7 text-center font-mono text-[7.5px] uppercase tracking-[0.28em] text-[#efe3cb]/25">
            {{ name }} · Official atelier profile
          </footer>
        </div>
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
const roman = (n) => ['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII'][n - 1] || String(n)
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.cw-gold-aura {
  background:
    radial-gradient(50% 55% at 50% 0%, rgba(201, 165, 99, 0.16), transparent 70%),
    radial-gradient(35% 35% at 85% 40%, rgba(201, 165, 99, 0.06), transparent 75%);
}
.cw-gold-noise {
  opacity: 0.4;
  background-image: radial-gradient(rgba(255, 240, 210, 0.05) 0.5px, transparent 0.5px);
  background-size: 4px 4px;
}
.cw-gold-frame {
  position: relative;
  border: 1px solid rgba(201, 165, 99, 0.34);
  background: linear-gradient(165deg, rgba(201, 165, 99, 0.055), rgba(13, 10, 6, 0.2) 38%, rgba(201, 165, 99, 0.03));
  box-shadow: inset 0 0 0 1px rgba(13, 10, 6, 0.8), 0 28px 70px -40px rgba(201, 165, 99, 0.45);
}
.cw-gold-frame::before {
  content: '';
  position: absolute;
  inset: 5px;
  border: 1px solid rgba(201, 165, 99, 0.13);
  pointer-events: none;
}
.cw-gold-corner {
  position: absolute;
  height: 14px;
  width: 14px;
  border-color: #c9a563;
  border-style: solid;
  z-index: 1;
}
.cw-gold-tl { top: -1px; left: -1px; border-width: 2px 0 0 2px; }
.cw-gold-tr { top: -1px; right: -1px; border-width: 2px 2px 0 0; }
.cw-gold-bl { bottom: -1px; left: -1px; border-width: 0 0 2px 2px; }
.cw-gold-br { bottom: -1px; right: -1px; border-width: 0 2px 2px 0; }
.cw-gold-row {
  display: flex;
  align-items: center;
  gap: 11px;
  border-bottom: 1px solid rgba(201, 165, 99, 0.13);
  padding: 13px 2px;
  transition: background-color 0.25s ease;
}
.cw-gold-row:hover {
  background: rgba(201, 165, 99, 0.05);
}
.cw-gold-contact {
  display: flex;
  align-items: baseline;
  gap: 12px;
  border-bottom: 1px solid rgba(201, 165, 99, 0.13);
  padding: 10px 0;
}
.cw-gold-label {
  width: 62px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: rgba(239, 227, 203, 0.32);
}
.cw-gold-value {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: right;
  font-size: 12.5px;
  color: rgba(239, 227, 203, 0.76);
  transition: color 0.2s ease;
}
a.cw-gold-contact:hover .cw-gold-value {
  color: #c9a563;
}
.cw-gold-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border: 1px solid #c9a563;
  padding: 13px;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 10px;
  letter-spacing: 0.18em;
  color: #c9a563;
  transition: background-color 0.25s ease, color 0.25s ease, transform 0.2s ease;
}
.cw-gold-cta:hover {
  background: #c9a563;
  color: #171006;
}
.cw-gold-cta:active {
  transform: scale(0.985);
}
.cw-gold-in {
  animation: cw-gold-in 0.8s cubic-bezier(0.2, 0.75, 0.25, 1) both;
}
.cw-gold-d1 { animation-delay: 0.08s; }
.cw-gold-d2 { animation-delay: 0.17s; }
.cw-gold-d3 { animation-delay: 0.26s; }
@keyframes cw-gold-in {
  from { opacity: 0; transform: translateY(14px) scale(0.995); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-gold-in {
    animation: none !important;
  }
  .cw-gold-row,
  .cw-gold-value,
  .cw-gold-cta {
    transition: none !important;
  }
}
</style>

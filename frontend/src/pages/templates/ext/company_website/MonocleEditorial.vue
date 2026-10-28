<template>
  <div class="cw-edit min-h-screen bg-[#f7f1e3] text-[#241f1a]">
    <div class="mx-auto w-full max-w-md px-5 pb-10 pt-6">
      <header class="cw-edit-in flex items-center justify-between border-b border-[#241f1a]/25 pb-2 font-mono text-[9px] uppercase tracking-[0.22em] text-[#241f1a]/55">
        <span>Business &amp; Culture</span>
        <span>Vol. 28 — {{ nfcId || '01' }}</span>
        <span>¥ 0.00</span>
      </header>

      <div class="cw-edit-in cw-edit-d1 mt-4 flex items-center justify-between">
        <span class="font-serif text-[13px] italic tracking-wide text-[#7b2d35]">The Company Review</span>
        <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-[#241f1a]/45">Established · Independent</span>
      </div>

      <section class="cw-edit-in cw-edit-d1 mt-7">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-[#241f1a]/25"></span>
          <span class="font-mono text-[9px] uppercase tracking-[0.28em] text-[#7b2d35]">Cover Story</span>
          <span class="h-px flex-1 bg-[#241f1a]/25"></span>
        </div>
        <h1 class="mt-5 text-center font-serif text-[44px] leading-[1.02] tracking-[-0.03em]">{{ name }}</h1>
        <p class="mx-auto mt-4 max-w-[18rem] text-center font-serif text-[14px] italic leading-[1.7] text-[#241f1a]/60">A closer look at the work, the people and the practice.</p>
      </section>

      <section class="cw-edit-in cw-edit-d2 mt-8 border-y border-[#241f1a]/25 py-5">
        <div class="flex gap-4">
          <span class="cw-edit-drop flex h-12 w-12 flex-shrink-0 items-center justify-center border border-[#7b2d35]/40 font-serif text-[30px] leading-none text-[#7b2d35]">{{ initial }}</span>
          <p class="text-[13.5px] leading-[1.9] text-[#241f1a]/75">{{ description }}</p>
        </div>
        <div class="mt-5 flex items-center justify-between border-t border-[#241f1a]/15 pt-2 font-mono text-[8.5px] uppercase tracking-[0.18em] text-[#241f1a]/40">
          <span>Profile / 01</span>
          <span>Words by the editorial desk</span>
        </div>
      </section>

      <section class="cw-edit-in cw-edit-d2 mt-8">
        <div class="flex items-baseline justify-between">
          <h2 class="font-serif text-[22px] tracking-tight">The Practice</h2>
          <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-[#7b2d35]">{{ pad2(features.length) }} Disciplines</span>
        </div>
        <div v-if="features.length" class="mt-3">
          <article v-for="(f, i) in features" :key="i" class="cw-edit-item group">
            <span class="font-serif text-[12px] italic text-[#7b2d35]">{{ roman(i + 1) }}</span>
            <div class="flex-1">
              <h3 class="font-serif text-[19px] leading-snug transition-colors duration-200 group-hover:text-[#7b2d35]">{{ f }}</h3>
              <p class="mt-1 font-mono text-[8.5px] uppercase tracking-[0.18em] text-[#241f1a]/35">Service {{ pad2(i + 1) }} · Full scope</p>
            </div>
            <span class="font-serif text-[18px] text-[#241f1a]/20 transition-transform duration-200 group-hover:translate-x-1 group-hover:text-[#7b2d35]">→</span>
          </article>
        </div>
        <p v-else class="mt-3 border-y border-[#241f1a]/20 py-8 text-center font-serif text-[13px] italic text-[#241f1a]/40">Services are not yet listed.</p>
      </section>

      <section class="cw-edit-in cw-edit-d2 mt-8 grid grid-cols-3 border-y border-[#241f1a]/30 py-4 text-center">
        <div class="border-r border-[#241f1a]/15">
          <p class="font-serif text-[24px] leading-none text-[#7b2d35]">{{ pad2(features.length) }}</p>
          <p class="mt-2 font-mono text-[7px] uppercase tracking-[0.18em] text-[#241f1a]/40">Disciplines</p>
        </div>
        <div class="border-r border-[#241f1a]/15">
          <p class="font-serif text-[24px] leading-none text-[#7b2d35]">24H</p>
          <p class="mt-2 font-mono text-[7px] uppercase tracking-[0.18em] text-[#241f1a]/40">Response</p>
        </div>
        <div>
          <p class="font-serif text-[24px] leading-none text-[#7b2d35]">1:1</p>
          <p class="mt-2 font-mono text-[7px] uppercase tracking-[0.18em] text-[#241f1a]/40">Attention</p>
        </div>
      </section>

      <section class="cw-edit-in cw-edit-d3 mt-9 grid grid-cols-[1fr_1.35fr] gap-4">
        <div class="border-t-2 border-[#7b2d35] pt-2">
          <p class="font-mono text-[8.5px] uppercase tracking-[0.22em] text-[#7b2d35]">Inquiry</p>
          <h2 class="mt-2 font-serif text-[21px] leading-tight">Let's start a conversation.</h2>
          <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-edit-cta mt-4">
            访问官网 <span>↗</span>
          </a>
        </div>
        <div class="border-t border-[#241f1a]/25 pt-2">
          <p class="font-mono text-[8.5px] uppercase tracking-[0.22em] text-[#241f1a]/45">Directory</p>
          <div class="mt-1">
            <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-edit-contact">
              <span>Tel</span><span class="tabular-nums">{{ contact.phone }}</span>
            </a>
            <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-edit-contact">
              <span>Mail</span><span>{{ contact.email }}</span>
            </a>
            <div v-if="contact.address" class="cw-edit-contact items-start">
              <span>Add</span><span class="whitespace-normal">{{ contact.address }}</span>
            </div>
            <p v-if="!hasContact" class="py-4 font-serif text-[12px] italic text-[#241f1a]/40">No contact listed.</p>
          </div>
        </div>
      </section>

      <footer class="cw-edit-in cw-edit-d3 mt-9 flex items-center justify-between border-t border-[#241f1a]/30 pt-2 font-mono text-[8px] uppercase tracking-[0.2em] text-[#241f1a]/40">
        <span>Printed for the web</span>
        <span>{{ name }} · Official profile</span>
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
const initial = computed(() => String(name.value).trim().charAt(0) || '企')
const pad2 = (n) => String(n).padStart(2, '0')
const roman = (n) => ['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX', 'X'][n - 1] || String(n)
</script>

<style scoped>
.cw-edit {
  font-feature-settings: 'liga' 1, 'kern' 1;
}
.cw-edit-drop {
  border-radius: 999px;
  font-style: italic;
}
.cw-edit-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  border-top: 1px solid rgba(36, 31, 26, 0.18);
  padding: 14px 0;
  transition: padding-left 0.25s ease, background-color 0.25s ease;
}
.cw-edit-item:hover {
  padding-left: 4px;
  background: rgba(123, 45, 53, 0.035);
}
.cw-edit-cta {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border-bottom: 1px solid #7b2d35;
  padding-bottom: 3px;
  font-size: 12px;
  font-weight: 600;
  color: #7b2d35;
  transition: gap 0.2s ease, opacity 0.2s ease;
}
.cw-edit-cta:hover {
  gap: 13px;
  opacity: 0.75;
}
.cw-edit-cta:active {
  transform: translateY(1px);
}
.cw-edit-contact {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 10px;
  border-bottom: 1px solid rgba(36, 31, 26, 0.14);
  padding: 8px 0;
  font-size: 11px;
  line-height: 1.5;
}
.cw-edit-contact > span:first-child {
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: rgba(36, 31, 26, 0.4);
}
.cw-edit-contact > span:last-child {
  overflow-wrap: anywhere;
  text-align: right;
  color: rgba(36, 31, 26, 0.78);
}
a.cw-edit-contact:hover > span:last-child {
  color: #7b2d35;
  text-decoration: underline;
  text-underline-offset: 3px;
}
.cw-edit-in {
  animation: cw-edit-in 0.7s cubic-bezier(0.22, 0.7, 0.25, 1) both;
}
.cw-edit-d1 { animation-delay: 0.06s; }
.cw-edit-d2 { animation-delay: 0.15s; }
.cw-edit-d3 { animation-delay: 0.24s; }
@keyframes cw-edit-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-edit-in {
    animation: none !important;
  }
  .cw-edit-item,
  .cw-edit-cta,
  .cw-edit-contact > span {
    transition: none !important;
  }
}
</style>

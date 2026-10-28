<template>
  <div class="cw-bot relative min-h-screen overflow-hidden bg-[#e6ede2] text-[#22382c]">
    <div class="cw-bot-wash pointer-events-none absolute inset-x-0 top-0 h-96"></div>
    <svg class="pointer-events-none absolute -right-6 top-10 h-40 w-32 text-[#4e735a]/30" viewBox="0 0 100 140" fill="none" aria-hidden="true">
      <path d="M50 138V34" stroke="currentColor" stroke-width="1.4" />
      <path d="M50 78C30 74 18 60 14 40c22 2 34 14 36 38Z" stroke="currentColor" stroke-width="1.2" />
      <path d="M50 58c18-2 30-16 34-36-22 2-32 14-34 36Z" stroke="currentColor" stroke-width="1.2" />
      <path d="M50 108c-16-3-26-14-30-30 18 2 28 12 30 30Z" stroke="currentColor" stroke-width="1.2" />
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-9 pt-6">
      <header class="cw-bot-in flex items-center justify-between px-1">
        <span class="font-mono text-[8.5px] uppercase tracking-[0.24em] text-[#3d5c49]/60">Field notes / Company</span>
        <span class="flex items-center gap-1.5 font-mono text-[8px] uppercase tracking-[0.16em] text-[#3d5c49]/45">
          <Leaf class="h-3 w-3" /> {{ nfcId || '01' }}
        </span>
      </header>

      <section class="cw-bot-in cw-bot-d1 mt-5">
        <div class="cw-bot-arch relative px-7 pb-9 pt-12 text-center">
          <span class="cw-bot-line pointer-events-none absolute inset-x-5 top-5 bottom-4 rounded-t-full border border-[#3d5c49]/20"></span>
          <p class="font-mono text-[8px] uppercase tracking-[0.32em] text-[#6d8a74]">Rooted in practice</p>
          <h1 class="mt-4 font-serif text-[40px] leading-[1.04] tracking-[-0.02em] text-[#1f3327]">{{ name }}</h1>
          <span class="mx-auto mt-5 flex h-8 w-8 items-center justify-center rounded-full border border-[#3d5c49]/30 text-[#3d5c49]">
            <Sprout class="h-4 w-4" />
          </span>
          <p class="mx-auto mt-4 max-w-[17rem] text-[12.5px] leading-[1.9] text-[#3d5c49]/70">{{ description }}</p>
        </div>
      </section>

      <section class="cw-bot-in cw-bot-d2 mt-6">
        <div class="flex items-center gap-3">
          <h2 class="font-serif text-[21px] text-[#1f3327]">Our services</h2>
          <span class="h-px flex-1 bg-[#3d5c49]/25"></span>
          <span class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#3d5c49]/45">{{ pad2(features.length) }} plots</span>
        </div>
        <div v-if="features.length" class="mt-2">
          <div v-for="(f, i) in features" :key="i" class="cw-bot-row group">
            <span class="flex h-6 w-6 flex-shrink-0 items-center justify-center rounded-full bg-[#d6e2d4] font-mono text-[8px] text-[#3d5c49]">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[14px] leading-snug text-[#22382c]/85 transition-colors group-hover:text-[#1f3327]">{{ f }}</span>
            <span class="font-serif text-[12px] italic text-[#6d8a74]/0 transition-colors group-hover:text-[#6d8a74]/80">grow</span>
          </div>
        </div>
        <p v-else class="mt-3 border-y border-dashed border-[#3d5c49]/25 py-8 text-center font-serif text-[13px] italic text-[#3d5c49]/40">The garden is not yet planted.</p>
      </section>

      <section class="cw-bot-in cw-bot-d2 mt-7">
        <div class="cw-bot-note">
          <div class="flex items-start gap-3">
            <span class="flex h-7 w-7 flex-shrink-0 items-center justify-center rounded-full border border-[#3d5c49]/25"><Compass class="h-3.5 w-3.5" /></span>
            <div class="flex-1">
              <p class="font-mono text-[7.5px] uppercase tracking-[0.18em] text-[#6d8a74]">Our ground</p>
              <p class="mt-1 text-[12px] leading-relaxed text-[#22382c]/75">{{ contact.address }}</p>
            </div>
            <div class="text-right">
              <p class="font-serif text-[24px] leading-none text-[#3d5c49]">{{ pad2(features.length) }}</p>
              <p class="mt-1 font-mono text-[7px] uppercase tracking-[0.1em] text-[#3d5c49]/45">fields</p>
            </div>
          </div>
        </div>
      </section>

      <section class="cw-bot-in cw-bot-d3 mt-7">
        <div class="flex items-center gap-3">
          <h2 class="font-serif text-[21px] text-[#1f3327]">Say hello</h2>
          <span class="h-px flex-1 bg-[#3d5c49]/25"></span>
        </div>
        <div class="mt-1">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-bot-contact">
            <Phone class="h-3.5 w-3.5 text-[#4e735a]" />
            <span class="flex-1 text-[12.5px] tabular-nums">{{ contact.phone }}</span>
            <span class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#4e735a]/55">call</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-bot-contact">
            <Mail class="h-3.5 w-3.5 text-[#4e735a]" />
            <span class="flex-1 truncate text-[12.5px]">{{ contact.email }}</span>
            <span class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#4e735a]/55">write</span>
          </a>
          <p v-if="!hasContact" class="py-5 text-center font-serif text-[12px] italic text-[#3d5c49]/40">Contact is not yet configured.</p>
        </div>
      </section>

      <div class="cw-bot-in cw-bot-d3 mt-7">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-bot-cta">
          <span>Visit our website</span>
          <span class="flex h-7 w-7 items-center justify-center rounded-full bg-[#22382c]/10"><ArrowUpRight class="h-3.5 w-3.5" /></span>
        </a>
      </div>

      <footer class="cw-bot-in cw-bot-d3 mt-7 flex items-center justify-between border-t border-[#3d5c49]/15 pt-3 font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#3d5c49]/40">
        <span>{{ name }}</span>
        <span>Grown, not assembled</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Leaf, Sprout, Compass, Phone, Mail, ArrowUpRight } from 'lucide-vue-next'

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
.cw-bot-wash {
  background:
    radial-gradient(70% 80% at 50% 0%, rgba(122, 158, 128, 0.28), transparent 70%),
    linear-gradient(to bottom, rgba(255, 255, 255, 0.35), transparent);
}
.cw-bot-arch {
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.55), rgba(214, 226, 212, 0.55));
  border-radius: 200px 200px 18px 18px;
  box-shadow: inset 0 0 0 1px rgba(61, 92, 73, 0.14), 0 22px 45px -32px rgba(34, 56, 44, 0.6);
}
.cw-bot-line {
  border-bottom: 0;
}
.cw-bot-row {
  display: flex;
  align-items: center;
  gap: 11px;
  border-bottom: 1px solid rgba(61, 92, 73, 0.16);
  padding: 11px 0;
  transition: padding-left 0.22s ease, background-color 0.22s ease;
}
.cw-bot-row:hover {
  padding-left: 3px;
  background: rgba(61, 92, 73, 0.035);
}
.cw-bot-note {
  border: 1px solid rgba(61, 92, 73, 0.18);
  border-radius: 14px;
  background: rgba(255, 255, 255, 0.4);
  padding: 12px 14px;
  transition: background-color 0.25s ease, border-color 0.25s ease;
}
.cw-bot-note:hover {
  background: rgba(255, 255, 255, 0.62);
  border-color: rgba(61, 92, 73, 0.3);
}
.cw-bot-contact {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(61, 92, 73, 0.16);
  padding: 11px 0;
  transition: padding-left 0.2s ease;
}
a.cw-bot-contact:hover {
  padding-left: 3px;
}
.cw-bot-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  border-radius: 999px;
  background: #2d4a37;
  padding: 8px 8px 8px 18px;
  font-size: 12px;
  letter-spacing: 0.04em;
  color: #eaf1e7;
  transition: background-color 0.22s ease, transform 0.2s ease;
}
.cw-bot-cta:hover {
  background: #3d5c49;
}
.cw-bot-cta:active {
  transform: scale(0.985);
}
.cw-bot-in {
  animation: cw-bot-in 0.75s cubic-bezier(0.2, 0.75, 0.2, 1) both;
}
.cw-bot-d1 { animation-delay: 0.08s; }
.cw-bot-d2 { animation-delay: 0.17s; }
.cw-bot-d3 { animation-delay: 0.26s; }
@keyframes cw-bot-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-bot-in {
    animation: none !important;
  }
  .cw-bot-row,
  .cw-bot-note,
  .cw-bot-contact,
  .cw-bot-cta {
    transition: none !important;
  }
}
</style>

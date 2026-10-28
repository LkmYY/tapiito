<template>
  <div class="cw-riso relative min-h-screen overflow-hidden bg-[#f2e3c8] text-[#20304a]">
    <div class="cw-riso-grain pointer-events-none absolute inset-0"></div>
    <div class="cw-riso-blob pointer-events-none absolute -right-16 top-24 h-64 w-64 rounded-full"></div>
    <div class="cw-riso-blob2 pointer-events-none absolute -left-20 bottom-40 h-52 w-52 rounded-full"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-8 pt-7">
      <header class="cw-riso-in flex items-end justify-between border-b-2 border-[#20304a] pb-2">
        <span class="font-mono text-[9px] font-bold uppercase tracking-[0.2em]">Riso Press / 印刷品</span>
        <span class="font-mono text-[9px] tabular-nums text-[#20304a]/55">01—{{ nfcId || '00' }}</span>
      </header>

      <section class="cw-riso-in cw-riso-d1 mt-7">
        <p class="inline-block -rotate-2 bg-[#d9482f] px-2 py-1 font-mono text-[8.5px] uppercase tracking-[0.18em] text-[#f8edd9]">Company profile · 企业档案</p>
        <h1 class="cw-riso-title mt-4 text-[52px] font-black leading-[0.88] tracking-[-0.055em]">{{ name }}</h1>
        <div class="mt-5 grid grid-cols-[3.25rem_1fr] gap-3">
          <span class="cw-riso-stamp flex h-12 w-12 items-center justify-center border-2 border-[#20304a] text-center font-mono text-[7.5px] font-bold uppercase leading-tight">OFFICIAL<br />PROFILE</span>
          <p class="text-[13px] leading-[1.8] text-[#20304a]/80">{{ description }}</p>
        </div>
      </section>

      <section class="cw-riso-in cw-riso-d2 mt-8">
        <div class="flex items-center gap-3">
          <h2 class="font-mono text-[10px] font-bold uppercase tracking-[0.18em]">我们的服务</h2>
          <span class="h-0.5 flex-1 bg-[#20304a]"></span>
          <span class="font-mono text-[9px] text-[#20304a]/55">{{ pad2(features.length) }} 项</span>
        </div>
        <div v-if="features.length" class="mt-3 space-y-2">
          <article v-for="(f, i) in features" :key="i" class="cw-riso-ticket group">
            <span class="cw-riso-num">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[14.5px] font-semibold leading-snug transition-transform duration-200 group-hover:translate-x-0.5">{{ f }}</span>
            <span class="font-mono text-[9px] font-bold uppercase tracking-[0.1em] text-[#d9482f]">SVC</span>
          </article>
        </div>
      </section>

      <section class="cw-riso-in cw-riso-d2 mt-8 border-2 border-[#20304a] bg-[#f8edd9]/60">
        <div class="grid grid-cols-3 divide-x-2 divide-[#20304a]">
          <div class="px-2 py-3 text-center">
            <p class="font-mono text-[22px] font-black leading-none">{{ pad2(features.length) }}</p>
            <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#20304a]/55">Services</p>
          </div>
          <div class="px-2 py-3 text-center">
            <p class="font-mono text-[22px] font-black leading-none">24H</p>
            <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#20304a]/55">Reply</p>
          </div>
          <div class="px-2 py-3 text-center">
            <p class="font-mono text-[22px] font-black leading-none">01</p>
            <p class="mt-1 font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#20304a]/55">Team</p>
          </div>
        </div>
      </section>

      <section class="cw-riso-in cw-riso-d3 mt-8">
        <div class="flex items-center gap-3">
          <h2 class="font-mono text-[10px] font-bold uppercase tracking-[0.18em]">联系印刷所</h2>
          <span class="h-0.5 flex-1 bg-[#20304a]"></span>
        </div>
        <div class="mt-2">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-riso-contact">
            <span class="cw-riso-key">TEL</span>
            <span class="cw-riso-val tabular-nums">{{ contact.phone }}</span>
            <span class="cw-riso-go">→</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-riso-contact">
            <span class="cw-riso-key">MAIL</span>
            <span class="cw-riso-val">{{ contact.email }}</span>
            <span class="cw-riso-go">→</span>
          </a>
          <div v-if="contact.address" class="cw-riso-contact items-start">
            <span class="cw-riso-key">ADDR</span>
            <span class="cw-riso-val whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-6 text-center font-mono text-[10px] text-[#20304a]/40">CONTACT / 未配置</p>
        </div>
      </section>

      <div class="cw-riso-in cw-riso-d3 mt-7">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-riso-cta">
          <span class="flex-1">访问官方网站</span>
          <span class="border-l-2 border-[#f8edd9]/40 pl-4 font-mono text-[10px]">WWW ↗</span>
        </a>
      </div>

      <footer class="cw-riso-in cw-riso-d3 mt-6 flex items-center justify-between border-t-2 border-[#20304a] pt-2 font-mono text-[8px] uppercase tracking-[0.12em] text-[#20304a]/50">
        <span>Two inks · One story</span>
        <span>Printed in China</span>
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
.cw-riso-grain {
  opacity: 0.5;
  background-image: radial-gradient(rgba(32, 48, 74, 0.16) 0.6px, transparent 0.6px);
  background-size: 4px 4px;
  mix-blend-mode: multiply;
}
.cw-riso-blob {
  background: #d9482f;
  opacity: 0.14;
  mix-blend-mode: multiply;
  transform: rotate(-18deg);
  border-radius: 42% 58% 63% 37% / 47% 39% 61% 53%;
}
.cw-riso-blob2 {
  background: #20304a;
  opacity: 0.1;
  mix-blend-mode: multiply;
  border-radius: 58% 42% 38% 62% / 55% 62% 38% 45%;
}
.cw-riso-title {
  text-shadow: 3px 3px 0 rgba(217, 72, 47, 0.9), -1.5px -1.5px 0 rgba(32, 48, 74, 0.22);
}
.cw-riso-stamp {
  transform: rotate(-4deg);
  background: rgba(217, 72, 47, 0.08);
  color: #d9482f;
}
.cw-riso-ticket {
  display: flex;
  align-items: center;
  gap: 11px;
  border: 2px solid #20304a;
  background: #f8edd9;
  padding: 11px 12px;
  box-shadow: 3px 3px 0 rgba(32, 48, 74, 0.9);
  transition: transform 0.2s ease, box-shadow 0.2s ease, background-color 0.2s ease;
}
.cw-riso-ticket:hover {
  transform: translate(-1px, -1px);
  box-shadow: 5px 5px 0 rgba(217, 72, 47, 0.85);
  background: #fff7e8;
}
.cw-riso-ticket:active {
  transform: translate(2px, 2px);
  box-shadow: 1px 1px 0 rgba(32, 48, 74, 0.9);
}
.cw-riso-num {
  display: flex;
  height: 25px;
  width: 25px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 999px;
  background: #20304a;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 9px;
  color: #f8edd9;
}
.cw-riso-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px dashed rgba(32, 48, 74, 0.35);
  padding: 10px 0;
}
.cw-riso-key {
  width: 36px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #d9482f;
}
.cw-riso-val {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  font-weight: 500;
}
.cw-riso-go {
  font-weight: 700;
  color: rgba(32, 48, 74, 0.35);
  transition: transform 0.2s ease, color 0.2s ease;
}
a.cw-riso-contact:hover .cw-riso-go {
  transform: translateX(3px);
  color: #d9482f;
}
a.cw-riso-contact:hover .cw-riso-val {
  color: #d9482f;
}
.cw-riso-cta {
  display: flex;
  align-items: center;
  background: #20304a;
  padding: 13px 16px;
  font-size: 14px;
  font-weight: 700;
  color: #f8edd9;
  box-shadow: 4px 4px 0 #d9482f;
  transition: transform 0.2s ease, box-shadow 0.2s ease, background-color 0.2s ease;
}
.cw-riso-cta:hover {
  background: #d9482f;
  box-shadow: 4px 4px 0 #20304a;
  transform: translate(-1px, -1px);
}
.cw-riso-cta:active {
  transform: translate(2px, 2px);
  box-shadow: 1px 1px 0 #20304a;
}
.cw-riso-in {
  animation: cw-riso-in 0.65s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cw-riso-d1 { animation-delay: 0.07s; }
.cw-riso-d2 { animation-delay: 0.16s; }
.cw-riso-d3 { animation-delay: 0.25s; }
@keyframes cw-riso-in {
  from { opacity: 0; transform: translateY(15px) rotate(0.4deg); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-riso-in {
    animation: none !important;
  }
  .cw-riso-ticket,
  .cw-riso-cta,
  .cw-riso-go,
  .cw-riso-val {
    transition: none !important;
  }
}
</style>

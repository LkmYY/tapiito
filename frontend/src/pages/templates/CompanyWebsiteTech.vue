<template>
  <div class="relative min-h-screen overflow-hidden bg-[#05070c] text-white">
    <!-- Grid + dot texture -->
    <div class="grid-bg pointer-events-none absolute inset-0"></div>
    <!-- Ambient glow -->
    <div class="pointer-events-none absolute inset-x-0 -top-32 h-72 bg-[radial-gradient(55%_100%_at_50%_0%,rgba(52,211,153,0.10),transparent_72%)]"></div>

    <div class="relative z-10 flex min-h-screen flex-col px-6 pb-8 pt-10">
      <!-- Status bar -->
      <header class="rise flex items-center justify-between border-b border-white/[0.07] pb-4 font-mono text-[10px] uppercase tracking-[0.2em]">
        <span class="flex items-center gap-2 text-emerald-400/90">
          <span class="status-dot h-1.5 w-1.5 rounded-full bg-emerald-400"></span>
          System Online
        </span>
        <span v-if="nfcId" class="text-white/30">NFC::{{ nfcId }}</span>
      </header>

      <!-- Hero -->
      <section class="rise rise-1 mt-12">
        <p class="font-mono text-[11px] tracking-[0.14em] text-emerald-400/80">// COMPANY_PROFILE</p>
        <h1 class="mt-4 text-balance text-[36px] font-semibold leading-[1.06] tracking-tight">{{ payload?.name || '企业官网' }}</h1>
        <p v-if="payload?.description" class="mt-5 border-l-2 border-emerald-400/40 pl-4 text-[14px] leading-relaxed text-slate-400">{{ payload.description }}</p>
      </section>

      <!-- Features -->
      <section v-if="features.length" class="rise rise-2 mt-10">
        <div class="flex items-center gap-3 font-mono text-[10px] uppercase tracking-[0.2em]">
          <span class="text-emerald-400/80">[ Features ]</span>
          <span class="h-px flex-1 bg-white/[0.08]"></span>
          <span class="text-white/30 tabular-nums">{{ pad2(features.length) }}</span>
        </div>
        <ul class="mt-3">
          <li v-for="(f, i) in features" :key="i" class="tech-row group">
            <span class="font-mono text-[11px] text-emerald-400/70 tabular-nums">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[15px] text-slate-200">{{ f }}</span>
            <Check class="h-4 w-4 text-emerald-400/50 transition-colors group-hover:text-emerald-400" />
          </li>
        </ul>
      </section>

      <!-- Bottom cluster -->
      <div class="rise rise-3 mt-auto pt-10">
        <section>
          <p class="font-mono text-[11px] text-white/30">$ open --official-site</p>
          <a
            v-if="payload?.url && isSafeUrl"
            :href="payload.url"
            target="_blank"
            rel="noopener noreferrer"
            class="mt-3 flex w-full items-center justify-center gap-2 rounded-lg bg-emerald-400 py-3.5 font-mono text-[13px] font-semibold uppercase tracking-[0.14em] text-[#04120b] transition-all hover:bg-emerald-300 active:scale-[0.98]"
          >
            访问官网
            <ArrowUpRight class="h-4 w-4" />
          </a>
          <p v-else class="mt-3 rounded-lg border border-dashed border-white/15 py-3.5 text-center font-mono text-[12px] text-white/30">// NO_URL_CONFIGURED</p>
        </section>

        <section v-if="hasContact" class="mt-9">
          <div class="flex items-center gap-3 font-mono text-[10px] uppercase tracking-[0.2em]">
            <span class="text-emerald-400/80">[ Contact ]</span>
            <span class="h-px flex-1 bg-white/[0.08]"></span>
          </div>
          <div class="mt-2">
            <a v-if="payload?.contactInfo?.phone" :href="'tel:' + payload.contactInfo.phone" class="tech-contact">
              <span class="tech-ic"><Phone class="h-3.5 w-3.5" /></span>
              <span class="tech-label">Tel</span>
              <span class="tech-value tabular-nums">{{ payload.contactInfo.phone }}</span>
            </a>
            <a v-if="payload?.contactInfo?.email" :href="'mailto:' + payload.contactInfo.email" class="tech-contact">
              <span class="tech-ic"><Mail class="h-3.5 w-3.5" /></span>
              <span class="tech-label">Mail</span>
              <span class="tech-value">{{ payload.contactInfo.email }}</span>
            </a>
            <div v-if="payload?.contactInfo?.address" class="tech-contact">
              <span class="tech-ic"><MapPin class="h-3.5 w-3.5" /></span>
              <span class="tech-label">Addr</span>
              <span class="tech-value">{{ payload.contactInfo.address }}</span>
            </div>
          </div>
        </section>

        <p class="mt-8 text-center font-mono text-[10px] tracking-[0.3em] text-white/20">EOF</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Check, ArrowUpRight, Phone, Mail, MapPin } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.url || ''))
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const hasContact = computed(() => {
  const c = props.payload?.contactInfo || {}
  return Boolean(c.phone || c.email || c.address)
})
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.grid-bg {
  background-image:
    radial-gradient(rgba(52, 211, 153, 0.08) 1px, transparent 1px),
    linear-gradient(rgba(148, 163, 184, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(148, 163, 184, 0.05) 1px, transparent 1px);
  background-size: 22px 22px, 44px 44px, 44px 44px;
  -webkit-mask-image: linear-gradient(to bottom, black, transparent 92%);
  mask-image: linear-gradient(to bottom, black, transparent 92%);
}

.status-dot {
  animation: pulse-ring 2.2s cubic-bezier(0.22, 1, 0.36, 1) infinite;
}
@keyframes pulse-ring {
  0% {
    box-shadow: 0 0 0 0 rgba(52, 211, 153, 0.45);
  }
  70% {
    box-shadow: 0 0 0 6px rgba(52, 211, 153, 0);
  }
  100% {
    box-shadow: 0 0 0 0 rgba(52, 211, 153, 0);
  }
}

.tech-row {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 2px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  transition: background-color 0.2s ease;
}
.tech-row:hover {
  background-color: rgba(255, 255, 255, 0.03);
}

.tech-contact {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}
.tech-contact:last-child {
  border-bottom: 0;
}
.tech-ic {
  display: flex;
  height: 30px;
  width: 30px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  border: 1px solid rgba(52, 211, 153, 0.25);
  background: rgba(52, 211, 153, 0.06);
  color: #34d399;
}
.tech-label {
  width: 42px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 0.18em;
  color: rgba(255, 255, 255, 0.3);
}
.tech-value {
  margin-left: auto;
  max-width: 60%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 13.5px;
  color: #cbd5e1;
}
a.tech-contact:hover .tech-value {
  color: #6ee7b7;
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.rise {
  animation: rise 0.65s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.rise-1 {
  animation-delay: 0.08s;
}
.rise-2 {
  animation-delay: 0.16s;
}
.rise-3 {
  animation-delay: 0.24s;
}
</style>
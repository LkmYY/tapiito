<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0b0d0b] text-white">
    <!-- Energy glow -->
    <div class="pointer-events-none absolute inset-x-0 -top-24 h-80 bg-[radial-gradient(55%_100%_at_50%_0%,rgba(163,230,53,0.14),transparent_72%)]"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-10">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-white/35">
        <span>Training Club</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Identity -->
      <div class="rise rise-1 mt-10">
        <h1 class="text-balance text-[34px] font-black uppercase leading-[1.02] tracking-tight">
          {{ payload?.gymName || '健身工作室' }}
        </h1>
        <p class="mt-3.5 text-[11px] font-semibold uppercase tracking-[0.32em] text-[#a3e635]">突破极限 · 超越自我</p>
      </div>

      <!-- Schedule -->
      <div v-if="classes.length" class="rise rise-2 mt-10">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-bold tracking-wide text-white">课程安排</h2>
          <span class="h-px flex-1 bg-white/10"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-white/30">Schedule</span>
        </div>
        <div class="relative mt-5">
          <span class="pointer-events-none absolute bottom-4 left-[5px] top-4 w-px bg-gradient-to-b from-[#a3e635]/50 via-white/10 to-transparent"></span>
          <div v-for="(c, i) in classes" :key="i" class="relative border-b border-white/5 py-4 pl-7 last:border-0">
            <span class="absolute left-0 top-[26px] h-[11px] w-[11px] rounded-full bg-[#a3e635] shadow-[0_0_14px_rgba(163,230,53,0.55)]"></span>
            <div class="flex items-center justify-between gap-3">
              <div class="min-w-0">
                <p class="truncate text-[15px] font-semibold">{{ c?.name || '团体课程' }}</p>
                <p v-if="c?.time" class="mt-1 font-mono text-[11px] tracking-wide tabular-nums text-white/45">{{ c.time }}</p>
              </div>
              <span v-if="c?.coach" class="ft-chip">{{ c.coach }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Facilities -->
      <div v-if="facilities.length" class="rise rise-3 mt-10">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-bold tracking-wide text-white">场馆设施</h2>
          <span class="h-px flex-1 bg-white/10"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-white/30">Facilities</span>
        </div>
        <div class="mt-4 flex flex-wrap gap-2">
          <span v-for="(f, i) in facilities" :key="i" class="ft-tag">
            <Check class="h-3.5 w-3.5 flex-shrink-0 text-[#a3e635]" />
            {{ f }}
          </span>
        </div>
      </div>

      <!-- Address -->
      <div v-if="payload?.address" class="rise rise-3 mt-10 flex items-start gap-3 border-t border-white/10 pt-6">
        <MapPin class="mt-0.5 h-4 w-4 flex-shrink-0 text-[#a3e635]" />
        <p class="text-[13px] leading-relaxed text-white/50">{{ payload.address }}</p>
      </div>

      <!-- CTA -->
      <div v-if="payload?.phone" class="rise rise-3 mt-8">
        <a :href="'tel:' + payload.phone" class="ft-cta">
          <Phone class="h-4 w-4" />
          预约体验课
        </a>
        <p class="mt-3 text-center font-mono text-xs tracking-[0.14em] tabular-nums text-white/40">{{ payload.phone }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Check, MapPin, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const classes = computed(() => {
  const list = props.payload?.classes
  return Array.isArray(list) ? list.filter((item) => item && (item.name || item.time || item.coach)) : []
})

const facilities = computed(() => {
  const list = props.payload?.facilities
  return Array.isArray(list) ? list.filter((f) => typeof f === 'string' && f.trim() !== '') : []
})
</script>

<style scoped>
.ft-chip {
  flex-shrink: 0;
  border-radius: 9999px;
  border: 1px solid rgba(163, 230, 53, 0.35);
  background: rgba(163, 230, 53, 0.08);
  padding: 4px 10px;
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.04em;
  color: #bef264;
}

.ft-tag {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.04);
  padding: 9px 14px;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.85);
}

.ft-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  border-radius: 14px;
  background: #a3e635;
  padding: 15px 0;
  font-size: 15px;
  font-weight: 800;
  letter-spacing: 0.08em;
  color: #0b0d0b;
  box-shadow: 0 20px 44px -18px rgba(163, 230, 53, 0.55);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.ft-cta:hover {
  background: #bef264;
}
.ft-cta:active {
  transform: scale(0.98);
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

<template>
  <div class="pf-root min-h-screen text-[#4a3730]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="pf-enter flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.2em] text-[#a8735e]">
        <span class="flex items-center gap-2"><Heart class="h-3.5 w-3.5 fill-current" />Family First</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="pf-enter pf-d1 relative mt-5">
        <div class="pf-sun" aria-hidden="true"></div>
        <p class="pf-hand text-[14px] text-[#c56b52]">家里人在等你回家。</p>
        <h1 class="mt-1 text-[34px] font-black leading-[1.08] tracking-tight text-[#3e2e28]">如果我不舒服，<br /><span class="text-[#c56b52]">请先联系家人</span></h1>
        <p class="mt-3 max-w-[280px] text-[11px] leading-relaxed text-[#8b7165]">他们知道我的情况。请告诉他们我们在哪里，并陪着等待。</p>
      </header>

      <section class="mt-6">
        <div class="pf-enter pf-d2 flex items-center justify-between">
          <p class="text-[9px] font-bold uppercase tracking-[0.2em] text-[#a8735e]">Family wall · 家人联络</p>
          <span class="text-[9px] tabular-nums text-[#b99a8d]">{{ String(contacts.length).padStart(2, '0') }}</span>
        </div>

        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="pf-primary pf-enter pf-d2 mt-3 block px-4 pb-5 pt-4">
            <div class="pf-photo pf-photo-main relative flex items-center justify-center overflow-hidden">
              <span class="pf-hill"></span>
              <span class="pf-sunmini"></span>
              <UserRound class="relative h-12 w-12 text-[#fff4ea]/85" />
              <span class="absolute bottom-2 left-2 rounded-full bg-[#fff8ef]/85 px-2 py-0.5 text-[8px] font-bold uppercase tracking-[0.1em] text-[#a8735e]">First call</span>
            </div>
            <div class="mt-3 flex items-baseline justify-between gap-2">
              <div>
                <p class="text-[13px] font-black">{{ contacts[0].name || '联系人' }}</p>
                <p v-if="contacts[0].relation" class="text-[10px] text-[#a8735e]">{{ contacts[0].relation }}</p>
              </div>
              <Phone class="h-4 w-4 text-[#c56b52]" />
            </div>
            <p class="pf-hand mt-1 break-words text-[31px] font-bold leading-tight tabular-nums text-[#b14f3b]">{{ contacts[0].phone }}</p>
            <p class="mt-1 text-center text-[8px] font-bold uppercase tracking-[0.16em] text-[#b99a8d]">Tap to call · 点按拨号</p>
          </a>

          <div class="mt-4 grid grid-cols-2 gap-3">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="pf-card pf-enter pf-d3 block px-3 pb-3 pt-2.5"
              :class="i % 2 === 0 ? 'rotate-[1.5deg]' : '-rotate-[1deg]'"
            >
              <div class="pf-photo flex h-20 items-center justify-center overflow-hidden" :class="'pf-tone-' + (i % 3)">
                <UserRound class="h-8 w-8 text-[#fff8ef]/80" />
              </div>
              <p class="mt-2 truncate text-[11px] font-black">{{ c.name || '联系人' }}</p>
              <p class="truncate text-[9px] text-[#a8735e]">{{ c.relation || '家人' }}</p>
              <p class="mt-1 break-words text-[17px] font-bold leading-tight tabular-nums text-[#8f3f2e]">{{ c.phone }}</p>
            </a>
          </div>
        </template>
        <div v-else class="pf-primary mt-3 px-4 py-8 text-center text-[12px] text-[#8b7165]">这里还没有家人的号码，请帮我们补上。</div>
      </section>

      <section class="pf-enter pf-d3 relative mt-7">
        <div class="pf-tape absolute -top-2 left-6 z-10 -rotate-2"></div>
        <div class="pf-note px-4 pb-4 pt-5">
          <div class="flex items-center justify-between">
            <p class="text-[10px] font-black uppercase tracking-[0.14em] text-[#9a6a2f]">Health card · 健康小卡</p>
            <HeartPulse class="h-4 w-4 text-[#c56b52]" />
          </div>
          <div class="mt-2 grid grid-cols-2 gap-x-4 gap-y-2">
            <div v-for="m in medical" :key="m.label" class="border-b border-[#dcc79f] pb-1.5">
              <p class="text-[8px] font-bold uppercase tracking-[0.08em] text-[#ad8c58]">{{ m.label }}</p>
              <p class="mt-0.5 text-[10px] font-medium leading-snug">{{ m.value }}</p>
            </div>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="pf-enter pf-d3 mt-4 flex items-center justify-between rounded-xl bg-[#ffe2d6] px-4 py-3 text-[11px] text-[#7d4f40]">
        <span class="flex items-center gap-2"><PawPrint class="h-3.5 w-3.5" />{{ payload.petName }}<span v-if="payload.breed" class="text-[#a8735e]"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-bold text-[#b14f3b] underline decoration-dotted underline-offset-4">宠物医院</a>
      </div>

      <p class="pf-hand pf-enter pf-d3 mt-6 text-center text-[13px] text-[#b99a8d]">谢谢你，愿一切平安。</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Heart, UserRound, Phone, HeartPulse, PawPrint } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)
const medical = computed(() => [
  { label: '血型', value: props.payload?.bloodType || '—' },
  { label: '过敏', value: props.payload?.allergies },
  { label: '用药', value: props.payload?.medications },
  { label: '病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.pf-root {
  background:
    radial-gradient(rgba(197, 107, 82, 0.1) 1px, transparent 1px),
    #fff1e6;
  background-size: 18px 18px, 100% 100%;
}
.pf-hand { font-family: "Segoe Script", "Bradley Hand", "Comic Sans MS", cursive; }
.pf-sun {
  position: absolute;
  right: 12px;
  top: -10px;
  width: 68px;
  height: 68px;
  border-radius: 999px;
  background: radial-gradient(circle, rgba(250, 179, 135, 0.65), transparent 68%);
}
.pf-primary {
  background: #fffaf3;
  border: 1px solid #efd7c8;
  box-shadow: 0 18px 36px -28px rgba(120, 70, 50, 0.5), 3px 4px 0 rgba(197, 107, 82, 0.08);
  transition: transform 0.22s ease, box-shadow 0.22s ease;
}
.pf-primary:hover { transform: translateY(-2px); box-shadow: 0 24px 44px -28px rgba(120, 70, 50, 0.55); }
.pf-primary:active { transform: scale(0.99); }
.pf-photo {
  background: linear-gradient(160deg, #f8c6a9, #e99882);
  border-radius: 3px;
}
.pf-photo-main { height: 86px; }
.pf-tone-0 { background: linear-gradient(160deg, #a9cfe3, #7fa8c4); }
.pf-tone-1 { background: linear-gradient(160deg, #f4c7ae, #e0a086); }
.pf-tone-2 { background: linear-gradient(160deg, #cbd9b0, #9fb47e); }
.pf-hill {
  position: absolute;
  left: -10%;
  right: -10%;
  bottom: -30px;
  height: 60px;
  border-radius: 50%;
  background: rgba(126, 90, 66, 0.25);
}
.pf-sunmini {
  position: absolute;
  right: 18px;
  top: 14px;
  width: 22px;
  height: 22px;
  border-radius: 999px;
  background: rgba(255, 240, 210, 0.8);
}
.pf-card {
  background: #fffaf3;
  border: 1px solid #efd7c8;
  box-shadow: 3px 4px 0 rgba(197, 107, 82, 0.1);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.pf-card:hover { transform: rotate(0) translateY(-3px); box-shadow: 5px 7px 0 rgba(197, 107, 82, 0.14); }
.pf-card:active { transform: scale(0.98); }
.pf-note {
  background: #fff3cd;
  border: 1px solid #e7cf99;
  box-shadow: 3px 4px 0 rgba(160, 120, 50, 0.1);
  transform: rotate(-0.4deg);
}
.pf-tape {
  width: 58px;
  height: 18px;
  background: rgba(255, 255, 255, 0.55);
  border-left: 1px solid rgba(200, 160, 130, 0.25);
  border-right: 1px solid rgba(200, 160, 130, 0.25);
}
.pf-enter { animation: pf-drop 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pf-d1 { animation-delay: 0.08s; }
.pf-d2 { animation-delay: 0.16s; }
.pf-d3 { animation-delay: 0.24s; }
@keyframes pf-drop {
  from { opacity: 0; transform: translateY(-14px) rotate(-1.5deg); }
  to { opacity: 1; transform: translateY(0) rotate(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pf-enter { animation: none; }
  .pf-primary, .pf-card { transition: none; }
}
</style>

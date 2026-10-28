<template>
  <div class="pt-ticket-root min-h-screen px-4 py-9 text-emerald-950">
    <div class="pt-ticket-content mx-auto w-full max-w-md">
      <div class="pt-ticket-head">
        <span>RETURN TICKET</span>
        <span>ADMIT ONE PET</span>
      </div>

      <article class="pt-ticket">
        <div class="pt-ticket-notch pt-ticket-notch-l" aria-hidden="true"></div>
        <div class="pt-ticket-notch pt-ticket-notch-r" aria-hidden="true"></div>

        <div class="pt-ticket-main">
          <div class="pt-ticket-stubhead">
            <span class="pt-ticket-no">NO. {{ nfcId || '0086' }}</span>
            <span class="pt-ticket-class">CLASS · COMPANION</span>
          </div>

          <div class="pt-ticket-title">
            <span class="pt-ticket-paw" aria-hidden="true">
              <svg viewBox="0 0 48 48" fill="currentColor">
                <ellipse cx="24" cy="31" rx="10" ry="8" />
                <ellipse cx="11" cy="21" rx="4" ry="5" />
                <ellipse cx="21" cy="15" rx="4" ry="5.5" />
                <ellipse cx="32" cy="16" rx="4" ry="5" />
                <ellipse cx="40" cy="24" rx="3.4" ry="4.4" />
              </svg>
            </span>
            <div>
              <p>ONE WAY · HOME</p>
              <h1>{{ payload.petName || '未命名' }}</h1>
            </div>
          </div>

          <div class="pt-ticket-meta">
            <p>{{ payload.breed }}<span>/</span>{{ payload.age || '—' }}<span>/</span>{{ payload.ownerName }}</p>
            <p class="pt-ticket-address">{{ payload.address }}</p>
          </div>

          <div class="pt-ticket-perf" aria-hidden="true"><span>✂</span></div>

          <div class="pt-ticket-body">
            <div class="pt-ticket-med">
              <span class="pt-ticket-label">MEDICAL</span>
              <p>{{ payload.medical || '无特殊记录' }}</p>
            </div>
            <div class="pt-ticket-note">
              <span class="pt-ticket-label">NOTE</span>
              <p>{{ payload.note || '怕生，请勿强行抱抱。' }}</p>
            </div>
          </div>

          <div class="pt-ticket-digits" aria-hidden="true">
            <span v-for="(d, i) in serial" :key="i">{{ d }}</span>
          </div>
        </div>

        <div class="pt-ticket-tear">
          <div class="pt-ticket-tear-left">
            <span class="pt-ticket-tear-label">IF LOST, CALL</span>
            <a :href="'tel:' + payload.phone">{{ payload.phone }}</a>
            <span class="pt-ticket-tear-sub">走失请拨打 · 24 小时接听</span>
          </div>
          <div class="pt-ticket-barcode" aria-hidden="true">
            <i v-for="n in 17" :key="n" :style="{ height: 16 + ((n * 13) % 20) + 'px' }"></i>
          </div>
        </div>

        <div class="pt-ticket-bottomline">
          <span>RETURN TO FAMILY</span>
          <span>TAP TO CALL</span>
        </div>
      </article>

      <div class="pt-ticket-holes" aria-hidden="true">
        <span v-for="n in 6" :key="n"></span>
      </div>
      <p class="pt-ticket-foot">KEEP THIS TICKET · 让爱回家</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const serial = computed(() => {
  const raw = String(props.nfcId || '0086').replace(/\D/g, '').padEnd(8, '0').slice(0, 8)
  return raw.split('')
})
</script>

<style scoped>
.pt-ticket-root {
  background:
    radial-gradient(circle at 50% 0, rgba(52, 211, 153, 0.13), transparent 36%),
    repeating-linear-gradient(90deg, rgba(255, 255, 255, 0.018) 0 1px, transparent 1px 6px),
    #062c25;
  font-family: 'Courier New', 'PingFang SC', monospace;
}
.pt-ticket-content { animation: pt-ticket-in 0.7s cubic-bezier(0.2, 0.85, 0.3, 1) both; }
.pt-ticket-head { display: flex; justify-content: space-between; margin: 0 9px 10px; font-size: 8px; letter-spacing: 0.16em; color: #6ee7b7; }
.pt-ticket {
  position: relative;
  color: #14332b;
  background:
    radial-gradient(circle at 8% 4%, rgba(15, 118, 110, 0.07), transparent 28%),
    #f6f1e3;
  border-radius: 9px;
  box-shadow: 0 22px 42px rgba(0, 0, 0, 0.4);
  overflow: hidden;
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.pt-ticket:hover { transform: translateY(-3px) rotate(-0.25deg); box-shadow: 0 30px 52px rgba(0, 0, 0, 0.46); }
.pt-ticket-notch { position: absolute; top: calc(100% - 132px); width: 23px; height: 23px; border-radius: 50%; background: #062c25; z-index: 3; }
.pt-ticket-notch-l { left: -12px; box-shadow: inset -4px 0 5px rgba(0, 0, 0, 0.16); }
.pt-ticket-notch-r { right: -12px; box-shadow: inset 4px 0 5px rgba(0, 0, 0, 0.16); }
.pt-ticket-main { padding: 15px 18px 14px; }
.pt-ticket-stubhead { display: flex; justify-content: space-between; padding-bottom: 8px; border-bottom: 1px solid rgba(20, 51, 43, 0.27); font-size: 7px; font-weight: 700; letter-spacing: 0.1em; color: #4c7a6d; }
.pt-ticket-title { display: flex; align-items: center; gap: 10px; padding: 13px 0 9px; }
.pt-ticket-paw { width: 39px; flex: none; color: #0f766e; transform: rotate(-9deg); }
.pt-ticket-paw svg { display: block; width: 100%; }
.pt-ticket-title p { margin: 0; font-size: 7px; letter-spacing: 0.18em; color: #7b948b; }
.pt-ticket-title h1 { margin: 1px 0 0; font-family: 'Arial Narrow', 'PingFang SC', sans-serif; font-size: clamp(38px, 11.5vw, 50px); line-height: 1; font-weight: 700; letter-spacing: 0.01em; color: #102f27; }
.pt-ticket-meta { padding: 8px 10px; border: 1px dashed rgba(20, 51, 43, 0.35); background: rgba(15, 118, 110, 0.04); }
.pt-ticket-meta p { margin: 0; font-size: 9px; font-weight: 700; letter-spacing: 0.04em; color: #365a50; }
.pt-ticket-meta p span { margin: 0 5px; color: #98ada5; }
.pt-ticket-address { margin-top: 3px !important; font-size: 8px !important; font-weight: 400 !important; color: #6f877e !important; }
.pt-ticket-perf { position: relative; margin: 12px -18px 11px; border-top: 1px dashed rgba(20, 51, 43, 0.45); }
.pt-ticket-perf span { position: absolute; top: -8px; left: 9px; background: #f6f1e3; padding: 0 3px; color: #7d948b; font-size: 9px; transform: rotate(180deg); }
.pt-ticket-body { display: grid; grid-template-columns: 0.95fr 1.05fr; gap: 12px; }
.pt-ticket-label { display: block; margin-bottom: 3px; font-size: 7px; font-weight: 700; letter-spacing: 0.13em; color: #769289; }
.pt-ticket-body p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.6; color: #365a50; }
.pt-ticket-med p { color: #9f4630; font-weight: 700; }
.pt-ticket-digits { display: flex; gap: 5px; margin-top: 12px; }
.pt-ticket-digits span { flex: 1; padding: 3px 0; text-align: center; border: 1px solid rgba(20, 51, 43, 0.25); font-size: 10px; font-weight: 700; color: #45685d; }
.pt-ticket-tear { display: flex; align-items: center; gap: 10px; padding: 11px 18px 10px; background: #0f766e; color: #f0fdfa; border-top: 1px dashed rgba(246, 241, 227, 0.65); }
.pt-ticket-tear-left { flex: 1; min-width: 0; }
.pt-ticket-tear-label { display: block; font-size: 7px; font-weight: 700; letter-spacing: 0.14em; color: #99f6e4; }
.pt-ticket-tear a { display: block; margin: 1px 0; font-family: 'Arial', 'PingFang SC', sans-serif; font-size: clamp(25px, 7.6vw, 32px); font-weight: 700; letter-spacing: 0.015em; color: #fff; text-decoration: none; transition: color 0.2s ease; }
.pt-ticket-tear a:hover { color: #bbf7d0; }
.pt-ticket-tear-sub { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.06em; color: #a7f3d0; }
.pt-ticket-barcode { flex: none; display: flex; align-items: flex-end; gap: 2px; height: 36px; padding: 0 2px; background: rgba(255, 255, 255, 0.1); }
.pt-ticket-barcode i { display: block; width: 2px; background: #d1fae5; }
.pt-ticket-barcode i:nth-child(3n) { width: 1px; opacity: 0.6; }
.pt-ticket-bottomline { display: flex; justify-content: space-between; padding: 6px 18px 7px; background: #0b5f59; color: #8fd9c7; font-size: 6px; letter-spacing: 0.14em; }
.pt-ticket-holes { display: flex; justify-content: space-around; margin: 9px 18px 0; }
.pt-ticket-holes span { width: 7px; height: 7px; border-radius: 50%; background: #08211c; box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.7); }
.pt-ticket-foot { margin: 14px 0 0; text-align: center; font-size: 8px; letter-spacing: 0.17em; color: #6ee7b7; }
@keyframes pt-ticket-in {
  0% { opacity: 0; transform: translateY(28px) rotate(1.5deg); }
  100% { opacity: 1; transform: translateY(0) rotate(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-ticket-content { animation: none; }
  .pt-ticket, .pt-ticket-tear a { transition: none; }
}
</style>

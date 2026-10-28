<template>
  <div class="nc-mf-root relative min-h-screen overflow-hidden px-5 py-8 text-[#f5e6c8]">
    <div class="nc-mf-room"></div>
    <div class="nc-mf-glow"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-mf-head">
        <span class="nc-mf-station">MIDNIGHT / FM</span>
        <span class="nc-mf-onair"><i></i>ON AIR</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-mf-main">
        <p class="nc-mf-band">98.7 — 99.1 MHz</p>

        <div class="nc-mf-dial">
          <div class="nc-mf-ticks">
            <i v-for="n in 31" :key="n" :class="{ 'nc-mf-tick-major': n % 5 === 1 }"></i>
          </div>
          <span class="nc-mf-needle"></span>
          <div class="nc-mf-freq">
            <span class="nc-mf-freq-big">98.7</span>
            <span class="nc-mf-freq-unit">MHz</span>
          </div>
        </div>

        <div class="nc-mf-show">
          <p class="nc-mf-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
          <h1 class="nc-mf-name">{{ payload?.name }}</h1>
          <p class="nc-mf-role">{{ payload?.role }}</p>
        </div>

        <div class="nc-mf-live">
          <span class="nc-mf-live-label">BROADCASTING FROM</span>
          <span class="nc-mf-live-venue">{{ payload?.venue }}</span>
          <span class="nc-mf-live-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-mf-request">
          <div class="nc-mf-request-head">
            <span>REQUEST LINE</span>
            <span class="nc-mf-eq" aria-hidden="true"><i></i><i></i><i></i><i></i><i></i></span>
          </div>

          <div class="nc-mf-row">
            <span class="nc-mf-row-label">WECHAT</span>
            <span class="nc-mf-row-value">{{ payload?.wechat }}</span>
          </div>

          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-mf-row nc-mf-row-link">
            <span class="nc-mf-row-label">CALL IN</span>
            <span class="nc-mf-row-value">{{ payload.phone }}</span>
            <span class="nc-mf-row-icon"><Phone :size="12" /></span>
          </a>
          <div v-else class="nc-mf-row">
            <span class="nc-mf-row-label">CALL IN</span>
            <span class="nc-mf-row-value nc-mf-empty"></span>
          </div>
        </section>

        <p class="nc-mf-dedication">FOR THE ONES STILL DRIVING HOME.</p>
      </main>

      <footer class="nc-mf-foot">
        <span>00:00 — 04:00</span>
        <span>NO SIGNAL AFTER DAWN</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-mf-root {
  background: #081226;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-mf-room {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 50% 26%, rgba(255, 179, 71, 0.14), transparent 44%),
    repeating-linear-gradient(90deg, rgba(255, 255, 255, 0.022) 0 1px, transparent 1px 60px);
  pointer-events: none;
}
.nc-mf-glow {
  position: absolute;
  top: 56px;
  left: 50%;
  width: 300px;
  height: 300px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 179, 71, 0.2), transparent 68%);
  transform: translateX(-50%);
  filter: blur(24px);
  pointer-events: none;
  animation: nc-mf-breathe 7s ease-in-out infinite;
}
.nc-mf-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(245, 230, 200, 0.18);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(245, 230, 200, 0.4);
  animation: nc-mf-in 0.8s both;
}
.nc-mf-station {
  color: #ffb347;
}
.nc-mf-onair {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  color: #ff7a5c;
}
.nc-mf-onair i {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #ff7a5c;
  box-shadow: 0 0 9px #ff7a5c;
  animation: nc-mf-blink 1.6s ease-in-out infinite;
}
.nc-mf-main {
  flex: 1;
  padding-top: 30px;
}
.nc-mf-band {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.3em;
  color: rgba(245, 230, 200, 0.38);
  animation: nc-mf-in 0.8s 0.06s both;
}
.nc-mf-dial {
  position: relative;
  margin-top: 10px;
  padding: 18px 14px 12px;
  border: 1px solid rgba(255, 179, 71, 0.3);
  background:
    linear-gradient(180deg, rgba(255, 179, 71, 0.08), transparent 55%),
    rgba(4, 12, 26, 0.7);
  box-shadow: inset 0 0 30px rgba(0, 0, 0, 0.4);
  animation: nc-mf-in 0.9s 0.1s both;
}
.nc-mf-ticks {
  display: flex;
  height: 26px;
  align-items: flex-end;
  justify-content: space-between;
  gap: 2px;
  border-bottom: 1px solid rgba(255, 179, 71, 0.35);
}
.nc-mf-ticks i {
  width: 1px;
  height: 7px;
  background: rgba(245, 230, 200, 0.28);
}
.nc-mf-ticks i.nc-mf-tick-major {
  height: 15px;
  background: rgba(255, 179, 71, 0.65);
}
.nc-mf-needle {
  position: absolute;
  top: 6px;
  left: 46%;
  width: 1px;
  height: 52px;
  background: #ff7a5c;
  box-shadow: 0 0 8px rgba(255, 122, 92, 0.9);
  animation: nc-mf-slide 1.1s 0.5s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-mf-freq {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  margin-top: 10px;
}
.nc-mf-freq-big {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 46px;
  line-height: 0.9;
  letter-spacing: -0.03em;
  color: #ffb347;
  text-shadow: 0 0 24px rgba(255, 179, 71, 0.35);
}
.nc-mf-freq-unit {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: rgba(245, 230, 200, 0.42);
}
.nc-mf-show {
  margin-top: 28px;
  animation: nc-mf-in 0.9s 0.2s both;
}
.nc-mf-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.44em;
  color: #ffb347;
  text-transform: uppercase;
}
.nc-mf-name {
  margin-top: 5px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(48px, 15.5vw, 66px);
  font-weight: 400;
  line-height: 0.96;
  letter-spacing: -0.035em;
  color: #f8ecd6;
  overflow-wrap: anywhere;
}
.nc-mf-role {
  margin-top: 8px;
  font-size: 9px;
  letter-spacing: 0.24em;
  color: rgba(245, 230, 200, 0.55);
}
.nc-mf-live {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  gap: 8px;
  margin-top: 20px;
  padding: 8px 10px;
  border-left: 2px solid #ffb347;
  background: rgba(255, 179, 71, 0.05);
  animation: nc-mf-in 0.9s 0.3s both;
}
.nc-mf-live-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: rgba(245, 230, 200, 0.35);
}
.nc-mf-live-venue {
  flex: 1;
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.1em;
  color: rgba(248, 236, 214, 0.9);
  text-transform: uppercase;
}
.nc-mf-live-city {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: rgba(255, 179, 71, 0.7);
}
.nc-mf-request {
  margin-top: 24px;
  border: 1px solid rgba(245, 230, 200, 0.14);
  animation: nc-mf-in 0.95s 0.38s both;
}
.nc-mf-request-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 7px 10px;
  border-bottom: 1px solid rgba(245, 230, 200, 0.12);
  background: rgba(255, 179, 71, 0.07);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(255, 179, 71, 0.7);
}
.nc-mf-eq {
  display: inline-flex;
  align-items: flex-end;
  gap: 2px;
  height: 9px;
}
.nc-mf-eq i {
  width: 2px;
  background: #ffb347;
  animation: nc-mf-eq 1s ease-in-out infinite alternate;
}
.nc-mf-eq i:nth-child(1) { height: 3px; animation-delay: 0.1s; }
.nc-mf-eq i:nth-child(2) { height: 7px; animation-delay: 0.3s; }
.nc-mf-eq i:nth-child(3) { height: 5px; animation-delay: 0.2s; }
.nc-mf-eq i:nth-child(4) { height: 9px; animation-delay: 0.4s; }
.nc-mf-eq i:nth-child(5) { height: 4px; animation-delay: 0.15s; }
.nc-mf-row {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 11px 10px;
  border-bottom: 1px solid rgba(245, 230, 200, 0.08);
  text-decoration: none;
  color: inherit;
}
.nc-mf-row:last-child {
  border-bottom: 0;
}
.nc-mf-row-label {
  width: 56px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: rgba(245, 230, 200, 0.34);
}
.nc-mf-row-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11.5px;
  letter-spacing: 0.04em;
  color: rgba(248, 236, 214, 0.9);
}
.nc-mf-row-link {
  transition: background 0.3s ease, padding 0.3s ease;
}
.nc-mf-row-link:hover {
  background: rgba(255, 179, 71, 0.12);
  padding-left: 16px;
}
.nc-mf-row-link:active {
  background: rgba(255, 179, 71, 0.22);
}
.nc-mf-row-icon {
  display: flex;
  flex: none;
  color: #ffb347;
  transition: transform 0.3s ease;
}
.nc-mf-row-link:hover .nc-mf-row-icon {
  transform: rotate(-10deg) scale(1.1);
}
.nc-mf-empty {
  color: rgba(245, 230, 200, 0.26);
}
.nc-mf-dedication {
  margin-top: 20px;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 10px;
  font-style: italic;
  letter-spacing: 0.04em;
  color: rgba(245, 230, 200, 0.38);
  animation: nc-mf-in 1s 0.5s both;
}
.nc-mf-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 14px;
  border-top: 1px solid rgba(245, 230, 200, 0.16);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(245, 230, 200, 0.28);
  animation: nc-mf-in 1s 0.6s both;
}
@keyframes nc-mf-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-mf-slide {
  from { left: 12%; }
  to { left: 46%; }
}
@keyframes nc-mf-breathe {
  0%, 100% { opacity: 0.7; transform: translateX(-50%) scale(1); }
  50% { opacity: 1; transform: translateX(-50%) scale(1.1); }
}
@keyframes nc-mf-blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.2; }
}
@keyframes nc-mf-eq {
  from { transform: scaleY(0.5); }
  to { transform: scaleY(1); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-mf-root *,
  .nc-mf-root *::before,
  .nc-mf-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

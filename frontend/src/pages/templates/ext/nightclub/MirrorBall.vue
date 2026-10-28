<template>
  <div class="nc-mb-root relative min-h-screen overflow-hidden px-5 py-7 text-[#f5edff]">
    <div class="nc-mb-glow"></div>
    <div class="nc-mb-sparkles" aria-hidden="true">
      <i v-for="n in 14" :key="n" :style="{ left: star[n - 1].x + '%', top: star[n - 1].y + '%', '--dd': star[n - 1].d + 's' }"></i>
    </div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3.5rem)] w-full max-w-md flex-col">
      <header class="nc-mb-head">
        <span>MIRROR BALL</span>
        <span>DISCOTHEQUE / {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-mb-main">
        <div class="nc-mb-rig">
          <span class="nc-mb-wire"></span>
          <div class="nc-mb-ball">
            <span class="nc-mb-ball-shade"></span>
            <span class="nc-mb-ball-high"></span>
          </div>
          <span class="nc-mb-floor"></span>
        </div>

        <p class="nc-mb-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-mb-name">{{ payload?.name }}</h1>
        <p class="nc-mb-role">{{ payload?.role }}</p>

        <div class="nc-mb-venue">
          <span class="nc-mb-venue-label">RESIDENT AT</span>
          <span class="nc-mb-venue-name">{{ payload?.venue }}</span>
          <span class="nc-mb-venue-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-mb-links">
          <div class="nc-mb-link">
            <span class="nc-mb-link-label">WECHAT</span>
            <span class="nc-mb-link-value">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-mb-link nc-mb-link-call">
            <span class="nc-mb-link-label">PHONE</span>
            <span class="nc-mb-link-value">{{ payload.phone }}</span>
            <Phone :size="14" />
          </a>
          <div v-else class="nc-mb-link">
            <span class="nc-mb-link-label">PHONE</span>
            <span class="nc-mb-link-value nc-mb-empty"></span>
          </div>
        </section>

        <p class="nc-mb-motto">SEE YOU UNDER THE LIGHT</p>
      </main>

      <footer class="nc-mb-foot">
        <span>SPIN / SHINE / REPEAT</span>
        <span>{{ payload?.city || 'SHANGHAI' }} · LATE</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])

const star = [
  { x: 16, y: 22, d: 0.2 }, { x: 78, y: 16, d: 1.4 }, { x: 30, y: 42, d: 0.7 },
  { x: 88, y: 38, d: 2.1 }, { x: 10, y: 54, d: 1.1 }, { x: 64, y: 30, d: 0.4 },
  { x: 46, y: 12, d: 1.8 }, { x: 92, y: 62, d: 0.9 }, { x: 22, y: 72, d: 2.4 },
  { x: 72, y: 70, d: 1.6 }, { x: 38, y: 58, d: 0.3 }, { x: 84, y: 82, d: 1.9 },
  { x: 8, y: 86, d: 1.2 }, { x: 56, y: 90, d: 2.2 },
]
</script>

<style scoped>
.nc-mb-root {
  background: #120716;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-mb-glow {
  position: absolute;
  top: 30px;
  left: 50%;
  width: 360px;
  height: 360px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 138, 197, 0.22), rgba(149, 112, 255, 0.08) 50%, transparent 70%);
  transform: translateX(-50%);
  filter: blur(20px);
  pointer-events: none;
  animation: nc-mb-breathe 6s ease-in-out infinite;
}
.nc-mb-sparkles {
  position: absolute;
  inset: 0;
  pointer-events: none;
}
.nc-mb-sparkles i {
  position: absolute;
  width: 2px;
  height: 2px;
  border-radius: 50%;
  background: #fff;
  box-shadow: 0 0 7px 1px rgba(255, 220, 245, 0.9);
  animation: nc-mb-twinkle 3.4s var(--dd) ease-in-out infinite;
}
.nc-mb-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(245, 237, 255, 0.15);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: rgba(245, 237, 255, 0.42);
  animation: nc-mb-in 0.9s both;
}
.nc-mb-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 18px;
  text-align: center;
}
.nc-mb-rig {
  position: relative;
  display: flex;
  height: 210px;
  flex-direction: column;
  align-items: center;
  animation: nc-mb-drop 1.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-mb-wire {
  width: 1px;
  height: 34px;
  background: linear-gradient(180deg, transparent, rgba(245, 237, 255, 0.4));
}
.nc-mb-ball {
  position: relative;
  width: 168px;
  height: 168px;
  overflow: hidden;
  border-radius: 50%;
  background:
    repeating-linear-gradient(0deg, rgba(0, 0, 0, 0.38) 0 1px, transparent 1px 14px),
    repeating-linear-gradient(90deg, rgba(0, 0, 0, 0.38) 0 1px, transparent 1px 14px),
    conic-gradient(from 40deg, #f7e7ff, #9d8ab5 18%, #f3d6ec 30%, #7e6d96 48%, #fff 58%, #b28fae 72%, #e6d4f0 86%, #f7e7ff);
  box-shadow:
    0 0 34px rgba(255, 178, 224, 0.4),
    inset -16px -18px 30px rgba(20, 5, 25, 0.75),
    inset 10px 12px 20px rgba(255, 255, 255, 0.3);
  animation: nc-mb-sway 7s ease-in-out infinite;
}
.nc-mb-ball::before {
  content: '';
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: repeating-conic-gradient(from 0deg, transparent 0 10deg, rgba(255, 255, 255, 0.16) 10deg 12deg);
  mix-blend-mode: screen;
}
.nc-mb-ball-shade {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: radial-gradient(circle at 68% 72%, rgba(20, 4, 24, 0.75), transparent 58%);
}
.nc-mb-ball-high {
  position: absolute;
  top: 18px;
  left: 26px;
  width: 46px;
  height: 32px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.72);
  filter: blur(8px);
  transform: rotate(-28deg);
}
.nc-mb-floor {
  width: 190px;
  height: 24px;
  margin-top: -8px;
  border-radius: 50%;
  background: radial-gradient(ellipse, rgba(255, 138, 197, 0.28), transparent 70%);
  filter: blur(4px);
}
.nc-mb-alias {
  margin-top: 6px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.5em;
  color: #ffa8d4;
  text-transform: uppercase;
  animation: nc-mb-in 0.9s 0.24s both;
}
.nc-mb-name {
  margin-top: 7px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(48px, 15vw, 64px);
  font-weight: 400;
  line-height: 0.98;
  letter-spacing: -0.03em;
  color: #fff1fb;
  text-shadow: 0 0 36px rgba(255, 168, 212, 0.34);
  overflow-wrap: anywhere;
  animation: nc-mb-in 1s 0.28s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-mb-role {
  margin-top: 10px;
  font-size: 8.5px;
  letter-spacing: 0.28em;
  color: rgba(245, 237, 255, 0.55);
  text-transform: uppercase;
  animation: nc-mb-in 0.9s 0.34s both;
}
.nc-mb-venue {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  justify-content: center;
  gap: 8px;
  margin-top: 22px;
  animation: nc-mb-in 0.9s 0.4s both;
}
.nc-mb-venue-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: rgba(245, 237, 255, 0.32);
}
.nc-mb-venue-name {
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.16em;
  color: rgba(255, 241, 251, 0.92);
  text-transform: uppercase;
}
.nc-mb-venue-city {
  padding-left: 8px;
  border-left: 1px solid rgba(245, 237, 255, 0.2);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: rgba(245, 237, 255, 0.4);
}
.nc-mb-links {
  width: 100%;
  margin-top: 26px;
  border-top: 1px solid rgba(245, 237, 255, 0.14);
  border-bottom: 1px solid rgba(245, 237, 255, 0.14);
  animation: nc-mb-in 1s 0.46s both;
}
.nc-mb-link {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 6px;
  text-align: left;
  text-decoration: none;
  color: inherit;
}
.nc-mb-link + .nc-mb-link {
  border-top: 1px solid rgba(245, 237, 255, 0.08);
}
.nc-mb-link-label {
  width: 48px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(245, 237, 255, 0.34);
}
.nc-mb-link-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  letter-spacing: 0.04em;
  color: rgba(255, 241, 251, 0.9);
}
.nc-mb-link-call {
  margin: 0 -6px;
  padding: 12px 12px;
  border-radius: 2px;
  transition: background 0.3s ease, padding 0.3s ease;
}
.nc-mb-link-call:hover {
  background: rgba(255, 168, 212, 0.12);
  padding-left: 18px;
}
.nc-mb-link-call:active {
  background: rgba(255, 168, 212, 0.22);
}
.nc-mb-link-call svg {
  flex: none;
  color: #ffa8d4;
  transition: transform 0.3s ease;
}
.nc-mb-link-call:hover svg {
  transform: rotate(-10deg) scale(1.1);
}
.nc-mb-empty {
  color: rgba(245, 237, 255, 0.3);
}
.nc-mb-motto {
  margin-top: 22px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.3em;
  color: rgba(255, 168, 212, 0.46);
  animation: nc-mb-in 1s 0.54s both;
}
.nc-mb-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 14px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(245, 237, 255, 0.28);
  animation: nc-mb-in 1s 0.62s both;
}
@keyframes nc-mb-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-mb-drop {
  from { opacity: 0; transform: translateY(-44px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-mb-sway {
  0%, 100% { transform: rotate(-2deg) translateX(-2px); }
  50% { transform: rotate(2deg) translateX(2px); }
}
@keyframes nc-mb-breathe {
  0%, 100% { opacity: 0.7; transform: translateX(-50%) scale(1); }
  50% { opacity: 1; transform: translateX(-50%) scale(1.08); }
}
@keyframes nc-mb-twinkle {
  0%, 100% { opacity: 0.15; transform: scale(0.7); }
  50% { opacity: 1; transform: scale(1.3); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-mb-root *,
  .nc-mb-root *::before,
  .nc-mb-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

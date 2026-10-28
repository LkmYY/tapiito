<template>
  <div class="nc-gl-root relative min-h-screen overflow-hidden px-5 py-9 text-[#eef6f5]">
    <div class="nc-gl-blob nc-gl-blob-a"></div>
    <div class="nc-gl-blob nc-gl-blob-b"></div>
    <div class="nc-gl-blob nc-gl-blob-c"></div>
    <div class="nc-gl-noise"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4.5rem)] w-full max-w-md flex-col">
      <header class="nc-gl-head">
        <span class="nc-gl-mark">L27</span>
        <span class="nc-gl-name">LOUNGE 27</span>
        <span class="nc-gl-time">OPEN LATE</span>
      </header>

      <main class="nc-gl-wrap">
        <div class="nc-gl-card">
          <div class="nc-gl-card-top">
            <span class="nc-gl-pill">{{ payload?.alias || 'NO ALIAS' }}</span>
            <span class="nc-gl-stars" aria-hidden="true">✦ ✦</span>
          </div>

          <h1 class="nc-gl-title">{{ payload?.name }}</h1>
          <p class="nc-gl-role">{{ payload?.role }}</p>

          <div class="nc-gl-venue">
            <span class="nc-gl-venue-dot"></span>
            <span class="nc-gl-venue-name">{{ payload?.venue }}</span>
            <span class="nc-gl-venue-city">{{ payload?.city || '—' }}</span>
          </div>

          <div class="nc-gl-divider"></div>

          <div class="nc-gl-contacts">
            <div class="nc-gl-row">
              <span class="nc-gl-label">WECHAT</span>
              <span class="nc-gl-value">{{ payload?.wechat }}</span>
            </div>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-gl-row nc-gl-row-link">
              <span class="nc-gl-label">PHONE</span>
              <span class="nc-gl-value">{{ payload.phone }}</span>
              <span class="nc-gl-go"><ArrowUpRight :size="13" /></span>
            </a>
            <div v-else class="nc-gl-row">
              <span class="nc-gl-label">PHONE</span>
              <span class="nc-gl-value nc-gl-empty"></span>
            </div>
          </div>
        </div>

        <div class="nc-gl-chips">
          <span>COCKTAILS</span>
          <span>LATE SETS</span>
          <span>PRIVATE BOOTH</span>
        </div>
      </main>

      <footer class="nc-gl-foot">
        <span>27TH FLOOR · RING THE BELL</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { ArrowUpRight } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-gl-root {
  background: #071318;
  font-family: 'Inter', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-gl-blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(70px);
  pointer-events: none;
}
.nc-gl-blob-a {
  top: -100px;
  left: -70px;
  width: 300px;
  height: 300px;
  background: rgba(24, 168, 160, 0.36);
  animation: nc-gl-float 12s ease-in-out infinite;
}
.nc-gl-blob-b {
  top: 32%;
  right: -110px;
  width: 280px;
  height: 280px;
  background: rgba(255, 122, 102, 0.32);
  animation: nc-gl-float 15s 1s ease-in-out infinite reverse;
}
.nc-gl-blob-c {
  bottom: -120px;
  left: 12%;
  width: 320px;
  height: 280px;
  background: rgba(67, 86, 214, 0.32);
  animation: nc-gl-float 18s 0.5s ease-in-out infinite;
}
.nc-gl-noise {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(255, 255, 255, 0.14) 0.5px, transparent 0.5px);
  background-size: 4px 4px;
  opacity: 0.18;
  pointer-events: none;
}
.nc-gl-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  animation: nc-gl-in 0.8s both;
}
.nc-gl-mark {
  display: flex;
  width: 30px;
  height: 30px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(238, 246, 245, 0.35);
  border-radius: 50%;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.04em;
  color: rgba(238, 246, 245, 0.82);
}
.nc-gl-name {
  flex: 1;
  font-size: 8px;
  letter-spacing: 0.38em;
  color: rgba(238, 246, 245, 0.62);
}
.nc-gl-time {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(238, 246, 245, 0.4);
}
.nc-gl-wrap {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 40px 0 30px;
}
.nc-gl-card {
  position: relative;
  overflow: hidden;
  padding: 22px 20px 18px;
  border: 1px solid rgba(255, 255, 255, 0.22);
  border-radius: 26px;
  background:
    linear-gradient(150deg, rgba(255, 255, 255, 0.2), rgba(255, 255, 255, 0.04) 40%, rgba(255, 255, 255, 0.1)),
    rgba(8, 30, 34, 0.55);
  box-shadow:
    0 30px 70px -24px rgba(0, 0, 0, 0.7),
    inset 0 1px 0 rgba(255, 255, 255, 0.28),
    inset 0 -1px 0 rgba(255, 255, 255, 0.06);
  backdrop-filter: blur(20px) saturate(1.3);
  animation: nc-gl-card 1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.45s ease, border-color 0.45s ease;
}
.nc-gl-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 12%;
  width: 76%;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.72), transparent);
}
.nc-gl-card:hover {
  transform: translateY(-4px);
  border-color: rgba(255, 255, 255, 0.36);
}
.nc-gl-card-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.nc-gl-pill {
  padding: 4px 10px;
  border: 1px solid rgba(255, 255, 255, 0.26);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.08);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.22em;
  color: rgba(238, 246, 245, 0.85);
  text-transform: uppercase;
}
.nc-gl-stars {
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #ffb4a8;
}
.nc-gl-title {
  margin-top: 28px;
  font-size: clamp(52px, 16.5vw, 68px);
  font-weight: 600;
  line-height: 0.92;
  letter-spacing: -0.055em;
  color: #f7fffd;
  text-shadow: 0 2px 30px rgba(0, 0, 0, 0.25);
  overflow-wrap: anywhere;
}
.nc-gl-role {
  margin-top: 10px;
  font-size: 9.5px;
  font-weight: 500;
  letter-spacing: 0.26em;
  color: rgba(238, 246, 245, 0.58);
}
.nc-gl-venue {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 24px;
  font-size: 10px;
  letter-spacing: 0.12em;
  color: rgba(238, 246, 245, 0.82);
}
.nc-gl-venue-dot {
  width: 5px;
  height: 5px;
  flex: none;
  border-radius: 50%;
  background: #7ef0df;
  box-shadow: 0 0 9px rgba(126, 240, 223, 0.8);
}
.nc-gl-venue-name {
  font-weight: 600;
  text-transform: uppercase;
}
.nc-gl-venue-city {
  padding-left: 8px;
  border-left: 1px solid rgba(255, 255, 255, 0.22);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  color: rgba(238, 246, 245, 0.45);
}
.nc-gl-divider {
  height: 1px;
  margin: 18px 0 6px;
  background: linear-gradient(90deg, rgba(255, 255, 255, 0.28), transparent);
}
.nc-gl-contacts {
  display: flex;
  flex-direction: column;
}
.nc-gl-row {
  display: flex;
  align-items: center;
  gap: 12px;
  min-height: 46px;
  padding: 9px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  text-decoration: none;
  color: inherit;
}
.nc-gl-label {
  width: 48px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(238, 246, 245, 0.4);
}
.nc-gl-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  letter-spacing: 0.04em;
  color: rgba(238, 246, 245, 0.92);
}
.nc-gl-row-link {
  margin: 0 -10px;
  padding: 9px 10px;
  border-radius: 10px;
  transition: background 0.3s ease, padding 0.3s ease;
}
.nc-gl-row-link:hover {
  background: rgba(255, 255, 255, 0.1);
  padding-left: 16px;
}
.nc-gl-row-link:active {
  background: rgba(255, 255, 255, 0.17);
}
.nc-gl-go {
  display: flex;
  flex: none;
  color: #7ef0df;
  transition: transform 0.3s ease;
}
.nc-gl-row-link:hover .nc-gl-go {
  transform: translate(2px, -2px);
}
.nc-gl-empty {
  color: rgba(238, 246, 245, 0.3);
}
.nc-gl-chips {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-top: 16px;
  animation: nc-gl-in 1s 0.5s both;
}
.nc-gl-chips span {
  padding: 4px 9px;
  border: 1px solid rgba(255, 255, 255, 0.14);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.04);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: rgba(238, 246, 245, 0.46);
}
.nc-gl-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: rgba(238, 246, 245, 0.32);
  animation: nc-gl-in 1s 0.62s both;
}
@keyframes nc-gl-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-gl-card {
  from { opacity: 0; transform: translateY(30px) scale(0.97); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-gl-float {
  0%, 100% { transform: translate(0, 0) scale(1); }
  50% { transform: translate(20px, 24px) scale(1.08); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-gl-root *,
  .nc-gl-root *::before,
  .nc-gl-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

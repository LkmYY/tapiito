<template>
  <div class="nc-cf-root relative min-h-screen overflow-hidden px-5 py-8 text-white">
    <div class="nc-cf-brushed"></div>
    <div class="nc-cf-edge nc-cf-edge-l"></div>
    <div class="nc-cf-edge nc-cf-edge-r"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-cf-head">
        <span class="nc-cf-logo">CHROME<span>/</span>FEVER</span>
        <span class="nc-cf-code">CF—{{ nfcId || '001' }}</span>
      </header>

      <section class="nc-cf-hero">
        <p class="nc-cf-kicker">AFTER MIDNIGHT PROGRAM</p>
        <p class="nc-cf-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-cf-name" data-text="">{{ payload?.name }}</h1>
        <div class="nc-cf-reflection" aria-hidden="true">{{ payload?.name }}</div>
        <div class="nc-cf-meta">
          <span>{{ payload?.role }}</span>
          <span class="nc-cf-dot"></span>
          <span>{{ payload?.city || '—' }}</span>
        </div>
      </section>

      <section class="nc-cf-plate">
        <div class="nc-cf-plate-top">
          <span class="nc-cf-plate-label">VENUE</span>
          <span class="nc-cf-plate-venue">{{ payload?.venue }}</span>
        </div>
        <div class="nc-cf-grid">
          <div class="nc-cf-cell">
            <span class="nc-cf-cell-key">WECHAT</span>
            <span class="nc-cf-cell-value">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-cf-cell nc-cf-cell-link">
            <span class="nc-cf-cell-key">PHONE</span>
            <span class="nc-cf-cell-value">{{ payload.phone }}</span>
          </a>
          <div v-else class="nc-cf-cell">
            <span class="nc-cf-cell-key">PHONE</span>
            <span class="nc-cf-cell-value nc-cf-empty"></span>
          </div>
          <div class="nc-cf-cell">
            <span class="nc-cf-cell-key">RESIDENCY</span>
            <span class="nc-cf-cell-value">{{ payload?.venue }}</span>
          </div>
          <div class="nc-cf-cell">
            <span class="nc-cf-cell-key">CITY</span>
            <span class="nc-cf-cell-value">{{ payload?.city || '—' }}</span>
          </div>
        </div>
        <div class="nc-cf-plate-foot">
          <span>POLISHED FOR THE NIGHT</span>
          <span class="nc-cf-bars" aria-hidden="true"><i></i><i></i><i></i><i></i><i></i></span>
        </div>
      </section>

      <footer class="nc-cf-foot">
        <span>BOOKING · {{ payload?.role || 'OPEN FORMAT' }}</span>
        <span>TAP THE PLATE</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-cf-root {
  background: #08090c;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-cf-brushed {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 76% 4%, rgba(255, 255, 255, 0.13), transparent 46%),
    repeating-linear-gradient(96deg, rgba(255, 255, 255, 0.022) 0 1px, transparent 1px 4px),
    linear-gradient(160deg, #16181d 0%, #0a0b0e 55%, #101217 100%);
  pointer-events: none;
}
.nc-cf-brushed::after {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(115deg, transparent 20%, rgba(255, 255, 255, 0.05) 42%, transparent 58%);
  animation: nc-cf-sheen 9s ease-in-out infinite;
}
.nc-cf-edge {
  position: absolute;
  top: 50%;
  width: 1px;
  height: 62vh;
  background: linear-gradient(180deg, transparent, rgba(255, 255, 255, 0.22), transparent);
  pointer-events: none;
}
.nc-cf-edge-l { left: 14px; }
.nc-cf-edge-r { right: 14px; }
.nc-cf-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.12);
  animation: nc-cf-drop 0.8s both;
}
.nc-cf-logo {
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.3em;
  color: #e8e8ea;
}
.nc-cf-logo span {
  padding: 0 3px;
  color: #8b8e96;
}
.nc-cf-code {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.4);
}
.nc-cf-hero {
  padding: 58px 0 34px;
  animation: nc-cf-drop 1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-cf-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.36em;
  color: rgba(255, 255, 255, 0.36);
}
.nc-cf-alias {
  margin-top: 18px;
  font-size: 10px;
  letter-spacing: 0.52em;
  color: #a9adb6;
  text-transform: uppercase;
}
.nc-cf-name {
  margin-top: 6px;
  background: linear-gradient(178deg, #ffffff 0%, #d4d7dd 26%, #6f747d 48%, #f3f4f6 58%, #8e939c 78%, #e6e8eb 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  -webkit-text-stroke: 1px rgba(255, 255, 255, 0.14);
  font-size: clamp(60px, 19vw, 88px);
  font-weight: 800;
  line-height: 0.88;
  letter-spacing: -0.055em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
  filter: drop-shadow(0 10px 26px rgba(0, 0, 0, 0.75));
}
.nc-cf-reflection {
  margin-top: 2px;
  transform: scaleY(-1);
  opacity: 0.16;
  filter: blur(1px);
  font-size: clamp(60px, 19vw, 88px);
  font-weight: 800;
  line-height: 0.65;
  letter-spacing: -0.055em;
  text-transform: uppercase;
  overflow: hidden;
  height: 0.5em;
  color: #cfd3da;
  pointer-events: none;
  mask-image: linear-gradient(180deg, #000, transparent 85%);
}
.nc-cf-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: 22px;
  font-size: 11px;
  letter-spacing: 0.2em;
  color: rgba(255, 255, 255, 0.7);
}
.nc-cf-dot {
  width: 4px;
  height: 4px;
  background: #b9bec8;
  transform: rotate(45deg);
}
.nc-cf-plate {
  position: relative;
  border: 1px solid rgba(255, 255, 255, 0.17);
  background:
    linear-gradient(155deg, rgba(255, 255, 255, 0.09), rgba(255, 255, 255, 0.015) 42%, rgba(255, 255, 255, 0.05)),
    #0d0f13;
  box-shadow: 0 26px 60px -28px #000, inset 0 1px 0 rgba(255, 255, 255, 0.2);
  animation: nc-cf-drop 1s 0.24s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.nc-cf-plate:hover {
  transform: translateY(-3px);
  box-shadow: 0 34px 74px -30px #000, inset 0 1px 0 rgba(255, 255, 255, 0.28);
}
.nc-cf-plate::before {
  content: '';
  position: absolute;
  inset: 5px;
  border: 1px solid rgba(255, 255, 255, 0.06);
  pointer-events: none;
}
.nc-cf-plate-top {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 14px;
  padding: 18px 18px 14px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}
.nc-cf-plate-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.3em;
  color: rgba(255, 255, 255, 0.35);
}
.nc-cf-plate-venue {
  font-size: 13px;
  font-weight: 600;
  letter-spacing: 0.14em;
  color: #f0f1f3;
  text-transform: uppercase;
}
.nc-cf-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
.nc-cf-cell {
  display: flex;
  flex-direction: column;
  gap: 7px;
  padding: 16px 18px;
  border-right: 1px solid rgba(255, 255, 255, 0.06);
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  text-decoration: none;
  color: inherit;
}
.nc-cf-cell:nth-child(2n) {
  border-right: 0;
}
.nc-cf-cell-key {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.26em;
  color: rgba(255, 255, 255, 0.32);
}
.nc-cf-cell-value {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  letter-spacing: 0.05em;
  color: #dfe1e5;
}
.nc-cf-cell-link {
  position: relative;
  overflow: hidden;
  transition: background 0.3s ease;
}
.nc-cf-cell-link::after {
  content: '';
  position: absolute;
  top: 0;
  left: -60%;
  width: 40%;
  height: 100%;
  background: linear-gradient(100deg, transparent, rgba(255, 255, 255, 0.18), transparent);
  transform: skewX(-18deg);
  transition: left 0.55s ease;
}
.nc-cf-cell-link:hover {
  background: rgba(255, 255, 255, 0.06);
}
.nc-cf-cell-link:hover::after {
  left: 120%;
}
.nc-cf-cell-link:active {
  background: rgba(255, 255, 255, 0.12);
}
.nc-cf-empty {
  color: rgba(255, 255, 255, 0.28);
}
.nc-cf-plate-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 13px 18px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.3);
}
.nc-cf-bars {
  display: inline-flex;
  align-items: flex-end;
  gap: 2px;
  height: 10px;
}
.nc-cf-bars i {
  display: block;
  width: 2px;
  background: #c9ccd3;
}
.nc-cf-bars i:nth-child(1) { height: 4px; }
.nc-cf-bars i:nth-child(2) { height: 9px; }
.nc-cf-bars i:nth-child(3) { height: 6px; }
.nc-cf-bars i:nth-child(4) { height: 10px; }
.nc-cf-bars i:nth-child(5) { height: 3px; }
.nc-cf-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: auto;
  padding-top: 24px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.28);
  animation: nc-cf-drop 1s 0.5s both;
}
@keyframes nc-cf-drop {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-cf-sheen {
  0%, 100% { transform: translateX(-30%); }
  50% { transform: translateX(35%); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-cf-root *,
  .nc-cf-root *::before,
  .nc-cf-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

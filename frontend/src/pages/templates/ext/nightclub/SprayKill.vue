<template>
  <div class="nc-sk-root relative min-h-screen overflow-hidden px-5 py-8 text-[#f4f2ea]">
    <div class="nc-sk-splatter nc-sk-splat-a"></div>
    <div class="nc-sk-splatter nc-sk-splat-b"></div>
    <div class="nc-sk-splatter nc-sk-splat-c"></div>
    <div class="nc-sk-tape nc-sk-tape-tl"></div>
    <div class="nc-sk-tape nc-sk-tape-br"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-sk-head">
        <span class="nc-sk-stencil">SPRAY / KILL</span>
        <span class="nc-sk-date">FRI 23:00 · {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-sk-main">
        <p class="nc-sk-kicker">NO GUEST LIST. NO PHOTOS. NO EXCUSES.</p>

        <h1 class="nc-sk-name">{{ payload?.name }}</h1>
        <div class="nc-sk-name-shadow" aria-hidden="true">{{ payload?.name }}</div>

        <div class="nc-sk-role-line">
          <span class="nc-sk-role">{{ payload?.role }}</span>
          <span class="nc-sk-slash">/</span>
          <span class="nc-sk-alias">{{ payload?.alias || 'NO ALIAS' }}</span>
        </div>

        <div class="nc-sk-badge">
          <span>RESIDENT<br />AFTER<br />DARK</span>
          <i>✕</i>
        </div>

        <section class="nc-sk-info">
          <div class="nc-sk-info-head">
            <span>FIND ME AT</span>
            <span class="nc-sk-paint-line"></span>
          </div>
          <p class="nc-sk-venue">{{ payload?.venue }}</p>
          <p class="nc-sk-city">{{ payload?.city || '—' }}</p>

          <div class="nc-sk-contacts">
            <div class="nc-sk-row">
              <span class="nc-sk-label">WECHAT</span>
              <span class="nc-sk-value">{{ payload?.wechat }}</span>
            </div>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-sk-row nc-sk-row-link">
              <span class="nc-sk-label">PHONE</span>
              <span class="nc-sk-value">{{ payload.phone }}</span>
              <Phone :size="14" class="nc-sk-phone-icon" />
            </a>
            <div v-else class="nc-sk-row">
              <span class="nc-sk-label">PHONE</span>
              <span class="nc-sk-value nc-sk-muted"></span>
            </div>
          </div>
        </section>

        <p class="nc-sk-sticker">ASK FOR<br />THE USUAL</p>
      </main>

      <footer class="nc-sk-foot">
        <span>ACID ON THE WALLS</span>
        <span>BASS IN THE BASEMENT</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-sk-root {
  background: #0c0c0b;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-sk-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background:
    radial-gradient(circle at 84% 18%, rgba(255, 45, 135, 0.19), transparent 34%),
    radial-gradient(circle at 8% 82%, rgba(214, 255, 47, 0.14), transparent 32%);
  pointer-events: none;
}
.nc-sk-splatter {
  position: absolute;
  border-radius: 50%;
  filter: blur(1px);
  pointer-events: none;
}
.nc-sk-splat-a {
  top: 70px;
  right: -36px;
  width: 150px;
  height: 150px;
  background: radial-gradient(circle at 40% 38%, #ff2d87 0 28%, rgba(255, 45, 135, 0.4) 42%, transparent 66%);
  transform: rotate(18deg);
  animation: nc-sk-pulse 8s ease-in-out infinite;
}
.nc-sk-splat-b {
  bottom: 120px;
  left: -52px;
  width: 180px;
  height: 150px;
  background: radial-gradient(ellipse at 50% 50%, #d6ff2f 0 22%, rgba(214, 255, 47, 0.35) 40%, transparent 68%);
  animation: nc-sk-pulse 10s 1s ease-in-out infinite;
}
.nc-sk-splat-c {
  top: 52%;
  right: 20px;
  width: 34px;
  height: 26px;
  background: #ff2d87;
  opacity: 0.7;
  transform: rotate(-24deg);
}
.nc-sk-tape {
  position: absolute;
  width: 92px;
  height: 24px;
  background: rgba(230, 222, 196, 0.55);
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.4);
  pointer-events: none;
}
.nc-sk-tape-tl {
  top: 62px;
  left: -28px;
  transform: rotate(-38deg);
  animation: nc-sk-tape 0.9s 0.5s both;
}
.nc-sk-tape-br {
  right: -26px;
  bottom: 82px;
  transform: rotate(-22deg);
  animation: nc-sk-tape 0.9s 0.7s both;
}
.nc-sk-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(244, 242, 234, 0.16);
  animation: nc-sk-in 0.7s both;
}
.nc-sk-stencil {
  padding: 3px 7px;
  border: 1px solid #d6ff2f;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.2em;
  color: #d6ff2f;
  transform: rotate(-1deg);
}
.nc-sk-date {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: rgba(244, 242, 234, 0.4);
}
.nc-sk-main {
  position: relative;
  flex: 1;
  padding-top: 42px;
}
.nc-sk-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: #ff2d87;
  animation: nc-sk-in 0.8s 0.08s both;
}
.nc-sk-name {
  position: relative;
  z-index: 2;
  margin-top: 15px;
  margin-left: -3px;
  font-size: clamp(64px, 21vw, 92px);
  font-weight: 900;
  line-height: 0.82;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  color: #f4f2ea;
  transform: rotate(-2.5deg);
  text-shadow: 3px 3px 0 #ff2d87;
  overflow-wrap: anywhere;
  animation: nc-sk-slam 0.8s 0.12s cubic-bezier(0.2, 1.4, 0.4, 1) both;
}
.nc-sk-name-shadow {
  position: absolute;
  top: 84px;
  left: 4px;
  z-index: 1;
  font-size: clamp(64px, 21vw, 92px);
  font-weight: 900;
  line-height: 0.82;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  color: transparent;
  -webkit-text-stroke: 1px rgba(214, 255, 47, 0.35);
  transform: rotate(-2.5deg) translate(7px, 7px);
  pointer-events: none;
}
.nc-sk-role-line {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 9px;
  margin-top: 26px;
  animation: nc-sk-in 0.8s 0.26s both;
}
.nc-sk-role {
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.18em;
  color: #d6ff2f;
  text-transform: uppercase;
}
.nc-sk-slash {
  color: rgba(244, 242, 234, 0.32);
}
.nc-sk-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.3em;
  color: rgba(244, 242, 234, 0.55);
  text-transform: uppercase;
}
.nc-sk-badge {
  position: absolute;
  top: 38px;
  right: -2px;
  display: flex;
  width: 78px;
  height: 78px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 3px;
  border: 1.5px dashed #d6ff2f;
  border-radius: 50%;
  background: rgba(12, 12, 11, 0.72);
  color: #d6ff2f;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  line-height: 1.5;
  letter-spacing: 0.1em;
  text-align: center;
  transform: rotate(11deg);
  animation: nc-sk-badge 0.9s 0.5s both;
  transition: transform 0.35s ease;
}
.nc-sk-badge:hover {
  transform: rotate(-4deg) scale(1.06);
}
.nc-sk-badge i {
  font-size: 9px;
  font-style: normal;
  color: #ff2d87;
}
.nc-sk-info {
  margin-top: 40px;
  padding: 17px 16px;
  border: 1px solid rgba(244, 242, 234, 0.17);
  background: rgba(10, 10, 9, 0.62);
  box-shadow: 6px 6px 0 rgba(255, 45, 135, 0.16);
  animation: nc-sk-in 0.9s 0.38s both;
  transition: box-shadow 0.3s ease, transform 0.3s ease;
}
.nc-sk-info:hover {
  box-shadow: 9px 9px 0 rgba(214, 255, 47, 0.16);
  transform: translate(-2px, -2px);
}
.nc-sk-info-head {
  display: flex;
  align-items: center;
  gap: 9px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.22em;
  color: rgba(244, 242, 234, 0.4);
}
.nc-sk-paint-line {
  flex: 1;
  height: 3px;
  background: #ff2d87;
  transform: rotate(-0.5deg);
}
.nc-sk-venue {
  margin-top: 10px;
  font-size: 22px;
  font-weight: 800;
  line-height: 1.05;
  letter-spacing: -0.02em;
  color: #f4f2ea;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-sk-city {
  margin-top: 2px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.28em;
  color: #d6ff2f;
}
.nc-sk-contacts {
  margin-top: 14px;
  border-top: 1px solid rgba(244, 242, 234, 0.12);
}
.nc-sk-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 9px 0;
  border-bottom: 1px solid rgba(244, 242, 234, 0.09);
  text-decoration: none;
  color: inherit;
}
.nc-sk-row-link {
  margin: 0 -7px;
  padding: 9px 7px;
  transition: background 0.25s ease, padding 0.25s ease;
}
.nc-sk-row-link:hover {
  background: rgba(255, 45, 135, 0.14);
  padding-left: 13px;
}
.nc-sk-row-link:active {
  background: rgba(255, 45, 135, 0.26);
}
.nc-sk-label {
  width: 44px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: rgba(244, 242, 234, 0.38);
}
.nc-sk-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 12px;
  letter-spacing: 0.04em;
  color: rgba(244, 242, 234, 0.9);
}
.nc-sk-phone-icon {
  flex: none;
  color: #d6ff2f;
}
.nc-sk-muted {
  color: rgba(244, 242, 234, 0.3);
}
.nc-sk-sticker {
  position: absolute;
  right: 2px;
  bottom: 4px;
  padding: 8px 12px;
  background: #d6ff2f;
  color: #0c0c0b;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  font-weight: 700;
  line-height: 1.4;
  letter-spacing: 0.12em;
  text-align: center;
  transform: rotate(7deg);
  transition: transform 0.3s ease;
}
.nc-sk-sticker:hover {
  transform: rotate(-3deg) scale(1.05);
}
.nc-sk-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 16px;
  border-top: 1px solid rgba(244, 242, 234, 0.16);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(244, 242, 234, 0.34);
  animation: nc-sk-in 1s 0.6s both;
}
@keyframes nc-sk-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-sk-slam {
  from { opacity: 0; transform: rotate(-2.5deg) scale(1.35); filter: blur(8px); }
  to { opacity: 1; transform: rotate(-2.5deg) scale(1); filter: blur(0); }
}
@keyframes nc-sk-badge {
  from { opacity: 0; transform: rotate(40deg) scale(0.6); }
  to { opacity: 1; transform: rotate(11deg) scale(1); }
}
@keyframes nc-sk-tape {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes nc-sk-pulse {
  0%, 100% { opacity: 0.8; transform: scale(1); }
  50% { opacity: 1; transform: scale(1.05); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-sk-root *,
  .nc-sk-root *::before,
  .nc-sk-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

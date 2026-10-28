<template>
  <div class="nc-bm-root relative min-h-screen overflow-hidden px-6 py-10 text-[#f2f0eb]">
    <div class="nc-bm-vignette"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-5rem)] w-full max-w-md flex-col">
      <header class="nc-bm-top">
        <span>NC / 01</span>
        <span class="nc-bm-top-mid">NOCTURNE</span>
        <span>{{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-bm-main">
        <p class="nc-bm-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-bm-name">{{ payload?.name }}</h1>
        <p class="nc-bm-role">{{ payload?.role }}</p>

        <div class="nc-bm-marker">
          <span class="nc-bm-marker-line"></span>
          <span class="nc-bm-marker-text">FOUND AFTER DARK</span>
        </div>

        <div class="nc-bm-venue-block">
          <p class="nc-bm-venue">{{ payload?.venue }}</p>
          <p class="nc-bm-city">{{ payload?.city || '—' }}</p>
        </div>

        <section class="nc-bm-list">
          <div class="nc-bm-row">
            <span class="nc-bm-index">01</span>
            <span class="nc-bm-label">WECHAT</span>
            <span class="nc-bm-value">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-bm-row nc-bm-row-link">
            <span class="nc-bm-index">02</span>
            <span class="nc-bm-label">PHONE</span>
            <span class="nc-bm-value">{{ payload.phone }}</span>
            <ArrowUpRight :size="14" class="nc-bm-arrow" />
          </a>
          <div v-else class="nc-bm-row">
            <span class="nc-bm-index">02</span>
            <span class="nc-bm-label">PHONE</span>
            <span class="nc-bm-value nc-bm-empty"></span>
          </div>
        </section>
      </main>

      <footer class="nc-bm-foot">
        <span>WALK IN LATE</span>
        <span class="nc-bm-dot"></span>
        <span>LEAVE LATER</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { ArrowUpRight } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-bm-root {
  background: #030303;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-bm-vignette {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 50% 0%, rgba(255, 255, 255, 0.055), transparent 46%),
    radial-gradient(ellipse at 50% 100%, rgba(120, 120, 120, 0.08), transparent 55%);
  pointer-events: none;
}
.nc-bm-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 14px;
  border-bottom: 1px solid rgba(242, 240, 235, 0.15);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: rgba(242, 240, 235, 0.4);
  animation: nc-bm-in 0.9s both;
}
.nc-bm-top-mid {
  font-size: 8px;
  letter-spacing: 0.44em;
  color: rgba(242, 240, 235, 0.72);
}
.nc-bm-main {
  flex: 1;
  padding-top: 68px;
}
.nc-bm-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.48em;
  color: rgba(242, 240, 235, 0.4);
  text-transform: uppercase;
  animation: nc-bm-in 0.9s 0.08s both;
}
.nc-bm-name {
  margin-top: 14px;
  font-size: clamp(62px, 20vw, 88px);
  font-weight: 200;
  line-height: 0.9;
  letter-spacing: -0.055em;
  color: #f4f2ed;
  overflow-wrap: anywhere;
  animation: nc-bm-name 1.1s 0.12s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-bm-role {
  margin-top: 16px;
  font-size: 10px;
  font-weight: 400;
  letter-spacing: 0.32em;
  color: rgba(242, 240, 235, 0.52);
  animation: nc-bm-in 0.9s 0.28s both;
}
.nc-bm-marker {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-top: 56px;
  animation: nc-bm-in 0.9s 0.34s both;
}
.nc-bm-marker-line {
  width: 42px;
  height: 1px;
  background: rgba(242, 240, 235, 0.45);
  transform-origin: left;
  animation: nc-bm-line 1.1s 0.5s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-bm-marker-text {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.28em;
  color: rgba(242, 240, 235, 0.34);
}
.nc-bm-venue-block {
  margin-top: 18px;
  animation: nc-bm-in 0.9s 0.4s both;
}
.nc-bm-venue {
  font-size: 15px;
  font-weight: 300;
  letter-spacing: 0.16em;
  color: rgba(242, 240, 235, 0.9);
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-bm-city {
  margin-top: 5px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.3em;
  color: rgba(242, 240, 235, 0.38);
}
.nc-bm-list {
  margin-top: 46px;
  border-top: 1px solid rgba(242, 240, 235, 0.18);
  animation: nc-bm-in 1s 0.5s both;
}
.nc-bm-row {
  display: flex;
  align-items: center;
  gap: 13px;
  min-height: 52px;
  border-bottom: 1px solid rgba(242, 240, 235, 0.1);
  text-decoration: none;
  color: inherit;
}
.nc-bm-index {
  width: 22px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  color: rgba(242, 240, 235, 0.28);
}
.nc-bm-label {
  width: 56px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(242, 240, 235, 0.35);
}
.nc-bm-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  font-weight: 300;
  letter-spacing: 0.06em;
  color: rgba(242, 240, 235, 0.9);
}
.nc-bm-row-link {
  position: relative;
  margin: 0 -8px;
  padding: 0 8px;
  transition: background 0.35s ease, padding 0.35s ease;
}
.nc-bm-row-link::before {
  content: '';
  position: absolute;
  left: 0;
  bottom: -1px;
  width: 0;
  height: 1px;
  background: #f2f0eb;
  transition: width 0.4s cubic-bezier(0.22, 1, 0.36, 1);
}
.nc-bm-row-link:hover {
  background: rgba(255, 255, 255, 0.04);
  padding-left: 14px;
}
.nc-bm-row-link:hover::before {
  width: 100%;
}
.nc-bm-row-link:active {
  background: rgba(255, 255, 255, 0.08);
}
.nc-bm-arrow {
  flex: none;
  color: rgba(242, 240, 235, 0.4);
  transition: transform 0.3s ease, color 0.3s ease;
}
.nc-bm-row-link:hover .nc-bm-arrow {
  transform: translate(2px, -2px);
  color: #f2f0eb;
}
.nc-bm-empty {
  color: rgba(242, 240, 235, 0.22);
}
.nc-bm-foot {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  padding-top: 28px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: rgba(242, 240, 235, 0.26);
  animation: nc-bm-in 1.1s 0.7s both;
}
.nc-bm-dot {
  width: 3px;
  height: 3px;
  border-radius: 50%;
  background: rgba(242, 240, 235, 0.5);
}
@keyframes nc-bm-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-bm-name {
  from { opacity: 0; transform: translateY(30px); filter: blur(6px); }
  to { opacity: 1; transform: none; filter: blur(0); }
}
@keyframes nc-bm-line {
  from { transform: scaleX(0); }
  to { transform: scaleX(1); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-bm-root *,
  .nc-bm-root *::before,
  .nc-bm-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

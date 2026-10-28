<template>
  <div class="nc-ss-root relative min-h-screen overflow-hidden px-6 py-8 text-[#ece7f7]">
    <div class="nc-ss-smoke nc-ss-smoke-a"></div>
    <div class="nc-ss-smoke nc-ss-smoke-b"></div>
    <div class="nc-ss-smoke nc-ss-smoke-c"></div>
    <span class="nc-ss-ghost" aria-hidden="true">{{ initial }}</span>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-ss-head">
        <span class="nc-ss-head-left">SMOKE / SIGNAL</span>
        <span class="nc-ss-head-right">{{ payload?.city || 'SHANGHAI' }} · 02:00</span>
      </header>

      <main class="nc-ss-main">
        <p class="nc-ss-index">NO. 07 — A QUIET FREQUENCY</p>
        <p class="nc-ss-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-ss-name">{{ payload?.name }}</h1>
        <p class="nc-ss-role">{{ payload?.role }}</p>

        <div class="nc-ss-wisp" aria-hidden="true">
          <span></span><span></span><span></span>
        </div>

        <div class="nc-ss-venue">
          <span class="nc-ss-venue-label">FOUND AT</span>
          <span class="nc-ss-venue-name">{{ payload?.venue }}</span>
          <span class="nc-ss-venue-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-ss-contact">
          <div class="nc-ss-row">
            <span class="nc-ss-key">微信</span>
            <span class="nc-ss-value">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-ss-row nc-ss-row-link">
            <span class="nc-ss-key">电话</span>
            <span class="nc-ss-value">{{ payload.phone }}</span>
            <ArrowUpRight :size="14" class="nc-ss-arrow" />
          </a>
          <div v-else class="nc-ss-row">
            <span class="nc-ss-key">电话</span>
            <span class="nc-ss-value nc-ss-empty"></span>
          </div>
        </section>
      </main>

      <footer class="nc-ss-foot">
        <span>INHALE · EXHALE · DISAPPEAR</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.nc-ss-root {
  background: #0b0910;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-ss-smoke {
  position: absolute;
  border-radius: 50%;
  filter: blur(70px);
  pointer-events: none;
}
.nc-ss-smoke-a {
  top: -90px;
  left: -60px;
  width: 330px;
  height: 330px;
  background: rgba(124, 99, 190, 0.28);
  animation: nc-ss-drift 14s ease-in-out infinite;
}
.nc-ss-smoke-b {
  top: 38%;
  right: -120px;
  width: 320px;
  height: 300px;
  background: rgba(66, 135, 160, 0.2);
  animation: nc-ss-drift 18s 1s ease-in-out infinite reverse;
}
.nc-ss-smoke-c {
  bottom: -100px;
  left: 18%;
  width: 300px;
  height: 260px;
  background: rgba(190, 120, 140, 0.16);
  animation: nc-ss-drift 16s 0.5s ease-in-out infinite;
}
.nc-ss-ghost {
  position: absolute;
  right: -36px;
  bottom: 16%;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 320px;
  font-weight: 400;
  line-height: 0.7;
  color: transparent;
  -webkit-text-stroke: 1px rgba(236, 231, 247, 0.055);
  pointer-events: none;
}
.nc-ss-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 14px;
  border-bottom: 1px solid rgba(236, 231, 247, 0.14);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.22em;
  color: rgba(236, 231, 247, 0.38);
  animation: nc-ss-in 0.9s both;
}
.nc-ss-head-left {
  color: rgba(236, 231, 247, 0.7);
}
.nc-ss-main {
  flex: 1;
  padding-top: 76px;
}
.nc-ss-index {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: rgba(236, 231, 247, 0.34);
  animation: nc-ss-in 0.9s 0.06s both;
}
.nc-ss-alias {
  margin-top: 30px;
  font-size: 8px;
  letter-spacing: 0.46em;
  color: #c4b2f0;
  text-transform: uppercase;
  animation: nc-ss-in 0.9s 0.12s both;
}
.nc-ss-name {
  margin-top: 10px;
  font-size: clamp(54px, 17vw, 72px);
  font-weight: 200;
  line-height: 0.96;
  letter-spacing: -0.045em;
  color: #f3effb;
  text-shadow: 0 0 46px rgba(196, 178, 240, 0.24);
  overflow-wrap: anywhere;
  animation: nc-ss-name 1.2s 0.16s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-ss-role {
  margin-top: 14px;
  font-size: 9px;
  letter-spacing: 0.3em;
  color: rgba(236, 231, 247, 0.52);
  animation: nc-ss-in 0.9s 0.28s both;
}
.nc-ss-wisp {
  display: flex;
  width: 110px;
  height: 30px;
  align-items: flex-end;
  gap: 6px;
  margin-top: 30px;
  animation: nc-ss-in 1s 0.34s both;
}
.nc-ss-wisp span {
  display: block;
  width: 2px;
  background: linear-gradient(180deg, transparent, rgba(196, 178, 240, 0.6));
  animation: nc-ss-wisp 4s ease-in-out infinite;
}
.nc-ss-wisp span:nth-child(1) { height: 14px; }
.nc-ss-wisp span:nth-child(2) { height: 28px; animation-delay: 0.6s; }
.nc-ss-wisp span:nth-child(3) { height: 19px; animation-delay: 1.2s; }
.nc-ss-venue {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  gap: 10px;
  margin-top: 18px;
  animation: nc-ss-in 0.9s 0.4s both;
}
.nc-ss-venue-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(236, 231, 247, 0.32);
}
.nc-ss-venue-name {
  font-size: 15px;
  font-weight: 500;
  letter-spacing: 0.16em;
  color: rgba(243, 239, 251, 0.9);
  text-transform: uppercase;
}
.nc-ss-venue-city {
  padding-left: 9px;
  border-left: 1px solid rgba(236, 231, 247, 0.2);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: rgba(236, 231, 247, 0.4);
}
.nc-ss-contact {
  margin-top: 50px;
  border-top: 1px solid rgba(236, 231, 247, 0.16);
  animation: nc-ss-in 1s 0.48s both;
}
.nc-ss-row {
  display: flex;
  align-items: center;
  gap: 16px;
  min-height: 56px;
  border-bottom: 1px solid rgba(236, 231, 247, 0.09);
  text-decoration: none;
  color: inherit;
}
.nc-ss-key {
  width: 38px;
  flex: none;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(236, 231, 247, 0.34);
}
.nc-ss-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 13px;
  font-weight: 300;
  letter-spacing: 0.06em;
  color: rgba(243, 239, 251, 0.9);
}
.nc-ss-row-link {
  margin: 0 -8px;
  padding: 0 8px;
  transition: background 0.35s ease, padding 0.35s ease;
}
.nc-ss-row-link:hover {
  background: linear-gradient(90deg, rgba(196, 178, 240, 0.1), transparent);
  padding-left: 16px;
}
.nc-ss-row-link:active {
  background: rgba(196, 178, 240, 0.18);
}
.nc-ss-arrow {
  flex: none;
  color: rgba(196, 178, 240, 0.55);
  transition: transform 0.3s ease, color 0.3s ease;
}
.nc-ss-row-link:hover .nc-ss-arrow {
  transform: translate(3px, -3px);
  color: #ece7f7;
}
.nc-ss-empty {
  color: rgba(236, 231, 247, 0.25);
}
.nc-ss-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 18px;
  border-top: 1px solid rgba(236, 231, 247, 0.12);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(236, 231, 247, 0.26);
  animation: nc-ss-in 1.1s 0.6s both;
}
@keyframes nc-ss-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-ss-name {
  from { opacity: 0; transform: translateY(26px); filter: blur(10px); }
  to { opacity: 1; transform: none; filter: blur(0); }
}
@keyframes nc-ss-drift {
  0%, 100% { transform: translate(0, 0) scale(1); }
  50% { transform: translate(26px, -18px) scale(1.16); }
}
@keyframes nc-ss-wisp {
  0%, 100% { transform: scaleY(0.7) translateX(0); opacity: 0.5; }
  50% { transform: scaleY(1.2) translateX(3px); opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-ss-root *,
  .nc-ss-root *::before,
  .nc-ss-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

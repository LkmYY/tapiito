<template>
  <div class="nc-cl-root relative min-h-screen overflow-hidden px-6 py-9 text-[#2b2419]">
    <div class="nc-cl-fiber"></div>
    <div class="nc-cl-frame"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4.5rem)] w-full max-w-md flex-col">
      <header class="nc-cl-head">
        <span>MAISON NOCTURNE</span>
        <span>CARTE DE NUIT</span>
      </header>

      <main class="nc-cl-main">
        <div class="nc-cl-crest">
          <svg viewBox="0 0 64 64" fill="none" aria-hidden="true">
            <path d="M20 18h24l-3 11c-1.4 5-4.7 7.5-9 7.5s-7.6-2.5-9-7.5l-3-11Z" stroke="currentColor" stroke-width="1.2" />
            <path d="M32 36.5V47M25 50h14" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" />
            <circle cx="32" cy="12" r="1.6" fill="currentColor" />
          </svg>
          <span class="nc-cl-seal">{{ initial }}</span>
        </div>

        <p class="nc-cl-eyebrow">CHAMPAGNE · COGNAC · LATE NIGHTS</p>
        <p class="nc-cl-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-cl-name">{{ payload?.name }}</h1>
        <p class="nc-cl-role">{{ payload?.role }}</p>

        <div class="nc-cl-rule">
          <span></span>
          <i>✦</i>
          <span></span>
        </div>

        <p class="nc-cl-intro">Tonight's pour is served by the house. Ask for the name below.</p>

        <section class="nc-cl-menu">
          <div class="nc-cl-item">
            <span class="nc-cl-item-key">LA RÉSIDENCE</span>
            <span class="nc-cl-dots"></span>
            <span class="nc-cl-item-val">{{ payload?.venue }}</span>
          </div>
          <div class="nc-cl-item">
            <span class="nc-cl-item-key">LA VILLE</span>
            <span class="nc-cl-dots"></span>
            <span class="nc-cl-item-val">{{ payload?.city || '—' }}</span>
          </div>
          <div class="nc-cl-item">
            <span class="nc-cl-item-key">WECHAT PRIVÉ</span>
            <span class="nc-cl-dots"></span>
            <span class="nc-cl-item-val">{{ payload?.wechat }}</span>
          </div>
          <div class="nc-cl-item">
            <span class="nc-cl-item-key">LIGNE DIRECTE</span>
            <span class="nc-cl-dots"></span>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-cl-item-val nc-cl-phone">{{ payload.phone }}</a>
          </div>
        </section>

        <p class="nc-cl-note">— served until the last guest leaves —</p>
      </main>

      <footer class="nc-cl-foot">
        <span>23:00 — LATE</span>
        <span class="nc-cl-foot-mid">TABLE SERVICE ONLY</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.nc-cl-root {
  background: #f5eede;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
}
.nc-cl-fiber {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 82% 12%, rgba(255, 255, 255, 0.8), transparent 43%),
    radial-gradient(rgba(90, 68, 35, 0.13) 0.55px, transparent 0.55px),
    linear-gradient(120deg, rgba(176, 138, 62, 0.06), transparent 50%);
  background-size: auto, 4px 4px, auto;
  pointer-events: none;
}
.nc-cl-frame {
  position: absolute;
  inset: 12px;
  border: 1px solid rgba(176, 138, 62, 0.32);
  pointer-events: none;
}
.nc-cl-frame::before {
  content: '';
  position: absolute;
  inset: 5px;
  border: 1px solid rgba(176, 138, 62, 0.18);
}
.nc-cl-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(43, 36, 25, 0.3);
  font-size: 7.5px;
  letter-spacing: 0.26em;
  color: #8a734d;
  animation: nc-cl-in 0.9s both;
}
.nc-cl-main {
  flex: 1;
  padding-top: 32px;
  text-align: center;
}
.nc-cl-crest {
  position: relative;
  width: 62px;
  height: 62px;
  margin: 0 auto;
  color: #a98544;
  animation: nc-cl-in 0.9s 0.06s both;
}
.nc-cl-crest svg {
  width: 100%;
  height: 100%;
}
.nc-cl-seal {
  position: absolute;
  right: -8px;
  bottom: -7px;
  display: flex;
  width: 24px;
  height: 24px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(168, 61, 45, 0.65);
  border-radius: 50%;
  background: #f5eede;
  font-family: Georgia, serif;
  font-size: 11px;
  font-style: italic;
  color: #a83d2d;
  transition: transform 0.3s ease;
}
.nc-cl-crest:hover .nc-cl-seal {
  transform: rotate(-12deg) scale(1.08);
}
.nc-cl-eyebrow {
  margin-top: 26px;
  font-size: 7px;
  letter-spacing: 0.32em;
  color: #a98544;
  animation: nc-cl-in 0.9s 0.12s both;
}
.nc-cl-alias {
  margin-top: 22px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.5em;
  color: #a83d2d;
  text-transform: uppercase;
  animation: nc-cl-in 0.9s 0.16s both;
}
.nc-cl-name {
  margin-top: 6px;
  font-size: clamp(50px, 16vw, 68px);
  font-weight: 400;
  line-height: 0.94;
  letter-spacing: -0.035em;
  color: #2b2419;
  overflow-wrap: anywhere;
  animation: nc-cl-in 1s 0.2s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-cl-role {
  margin-top: 12px;
  font-size: 9.5px;
  letter-spacing: 0.24em;
  color: #6f6048;
  animation: nc-cl-in 0.9s 0.26s both;
}
.nc-cl-rule {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 26px auto 18px;
  width: 70%;
  color: #a98544;
  animation: nc-cl-in 0.9s 0.32s both;
}
.nc-cl-rule span {
  flex: 1;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(169, 133, 68, 0.6), transparent);
}
.nc-cl-rule i {
  font-size: 9px;
  font-style: normal;
}
.nc-cl-intro {
  max-width: 280px;
  margin: 0 auto;
  font-size: 10.5px;
  font-style: italic;
  line-height: 1.65;
  letter-spacing: 0.03em;
  color: #7c6b52;
  animation: nc-cl-in 0.9s 0.36s both;
}
.nc-cl-menu {
  margin-top: 26px;
  padding: 4px 0;
  border-top: 1px solid rgba(43, 36, 25, 0.18);
  border-bottom: 1px solid rgba(43, 36, 25, 0.18);
  animation: nc-cl-in 1s 0.42s both;
}
.nc-cl-item {
  display: flex;
  align-items: baseline;
  gap: 8px;
  padding: 12px 4px;
  text-align: left;
}
.nc-cl-item + .nc-cl-item {
  border-top: 1px dotted rgba(43, 36, 25, 0.18);
}
.nc-cl-item-key {
  flex: none;
  font-size: 7.5px;
  letter-spacing: 0.16em;
  color: #8a734d;
}
.nc-cl-dots {
  flex: 1;
  border-bottom: 1px dotted rgba(43, 36, 25, 0.35);
  transform: translateY(-3px);
}
.nc-cl-item-val {
  min-width: 0;
  max-width: 58%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  letter-spacing: 0.03em;
  color: #2b2419;
  text-decoration: none;
}
.nc-cl-phone {
  transition: color 0.25s ease, letter-spacing 0.25s ease;
}
.nc-cl-phone:hover {
  color: #a83d2d;
  letter-spacing: 0.08em;
}
.nc-cl-phone:active {
  opacity: 0.6;
}
.nc-cl-empty {
  color: #b3a58e;
}
.nc-cl-note {
  margin-top: 22px;
  font-size: 9px;
  font-style: italic;
  letter-spacing: 0.06em;
  color: #9b8a6e;
  animation: nc-cl-in 1s 0.5s both;
}
.nc-cl-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  padding-top: 16px;
  border-top: 1px solid rgba(169, 133, 68, 0.32);
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #9b8a6e;
  animation: nc-cl-in 1s 0.6s both;
}
.nc-cl-foot-mid {
  font-style: italic;
  letter-spacing: 0.08em;
}
@keyframes nc-cl-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-cl-root *,
  .nc-cl-root *::before,
  .nc-cl-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

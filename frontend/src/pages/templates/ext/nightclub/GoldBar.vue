<template>
  <div class="nc-gb-root relative min-h-screen overflow-hidden px-5 py-8 text-[#f0e3c2]">
    <div class="nc-gb-shaft"></div>
    <div class="nc-gb-grain"></div>
    <span class="nc-gb-side" aria-hidden="true">PRIVATE · 999.9 · LATE</span>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-gb-head">
        <span class="nc-gb-mark">G / B</span>
        <span class="nc-gb-house">THE GOLD BAR · MEMBERS' FLOOR</span>
      </header>

      <main class="nc-gb-main">
        <div class="nc-gb-ingot">
          <div class="nc-gb-ingot-face">
            <span class="nc-gb-ingot-top">999.9 FINE</span>
            <span class="nc-gb-ingot-weight">100 g</span>
            <span class="nc-gb-ingot-line"></span>
            <span class="nc-gb-ingot-label">NOCTURNE RESERVE</span>
          </div>
          <span class="nc-gb-ingot-shadow"></span>
        </div>

        <p class="nc-gb-kicker">THE HOUSE RECOGNISES</p>
        <p class="nc-gb-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-gb-name">{{ payload?.name }}</h1>
        <p class="nc-gb-role">{{ payload?.role }}</p>

        <div class="nc-gb-venue">
          <span class="nc-gb-venue-line"></span>
          <span class="nc-gb-venue-name">{{ payload?.venue }}</span>
          <span class="nc-gb-venue-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-gb-contact">
          <div class="nc-gb-row">
            <span class="nc-gb-key">WECHAT</span>
            <span class="nc-gb-value">{{ payload?.wechat }}</span>
            <span class="nc-gb-row-no">I</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-gb-row nc-gb-row-link">
            <span class="nc-gb-key">PRIVATE LINE</span>
            <span class="nc-gb-value">{{ payload.phone }}</span>
            <span class="nc-gb-row-no">II</span>
          </a>
          <div v-else class="nc-gb-row">
            <span class="nc-gb-key">PRIVATE LINE</span>
            <span class="nc-gb-value nc-gb-empty"></span>
            <span class="nc-gb-row-no">II</span>
          </div>
        </section>

        <p class="nc-gb-note">Discretion is the only currency we accept.</p>
      </main>

      <footer class="nc-gb-foot">
        <span>EST. MMXVIII</span>
        <span class="nc-gb-foot-line"></span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-gb-root {
  background: #0a0805;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
}
.nc-gb-shaft {
  position: absolute;
  top: -160px;
  left: 50%;
  width: 300px;
  height: 620px;
  background: linear-gradient(180deg, rgba(212, 175, 106, 0.22), transparent 72%);
  clip-path: polygon(38% 0, 62% 0, 100% 100%, 0 100%);
  transform: translateX(-50%);
  filter: blur(14px);
  pointer-events: none;
  animation: nc-gb-shaft 1.2s both;
}
.nc-gb-grain {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(rgba(240, 227, 194, 0.13) 0.5px, transparent 0.5px),
    repeating-linear-gradient(90deg, rgba(212, 175, 106, 0.03) 0 1px, transparent 1px 26px);
  background-size: 4px 4px, auto;
  opacity: 0.5;
  pointer-events: none;
}
.nc-gb-side {
  position: absolute;
  top: 50%;
  right: 9px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.32em;
  color: rgba(212, 175, 106, 0.34);
  writing-mode: vertical-rl;
  transform: translateY(-50%);
}
.nc-gb-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 14px;
  border-bottom: 1px solid rgba(212, 175, 106, 0.28);
  animation: nc-gb-in 0.9s both;
}
.nc-gb-mark {
  display: flex;
  width: 28px;
  height: 28px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(212, 175, 106, 0.65);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.06em;
  color: #d4af6a;
}
.nc-gb-house {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.24em;
  color: rgba(240, 227, 194, 0.38);
}
.nc-gb-main {
  flex: 1;
  padding-top: 24px;
}
.nc-gb-ingot {
  position: relative;
  width: 132px;
  margin: 0 auto 32px;
  animation: nc-gb-ingot 1.1s 0.08s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-gb-ingot-face {
  position: relative;
  display: flex;
  height: 72px;
  flex-direction: column;
  justify-content: center;
  gap: 3px;
  padding: 9px 14px;
  border: 1px solid rgba(255, 232, 176, 0.55);
  background:
    linear-gradient(160deg, #f3d99a, #c9a04e 42%, #8a6425 72%, #e3c27d);
  box-shadow: 0 8px 20px -6px rgba(0, 0, 0, 0.8), inset 0 1px 0 rgba(255, 255, 255, 0.5);
  clip-path: polygon(6% 0, 94% 0, 100% 100%, 0 100%);
  color: #241a09;
}
.nc-gb-ingot-face::after {
  content: '';
  position: absolute;
  inset: 3px;
  border: 1px solid rgba(61, 42, 12, 0.28);
  clip-path: polygon(6% 0, 94% 0, 100% 100%, 0 100%);
  pointer-events: none;
}
.nc-gb-ingot-top {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  opacity: 0.7;
}
.nc-gb-ingot-weight {
  font-size: 22px;
  font-weight: 700;
  line-height: 1;
  letter-spacing: 0.02em;
}
.nc-gb-ingot-line {
  width: 100%;
  height: 1px;
  margin: 2px 0;
  background: rgba(36, 26, 9, 0.35);
}
.nc-gb-ingot-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 4.8px;
  letter-spacing: 0.16em;
  opacity: 0.66;
}
.nc-gb-ingot-shadow {
  display: block;
  width: 110px;
  height: 8px;
  margin: 6px auto 0;
  border-radius: 50%;
  background: rgba(212, 175, 106, 0.24);
  filter: blur(5px);
}
.nc-gb-kicker {
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.32em;
  color: rgba(212, 175, 106, 0.65);
  animation: nc-gb-in 0.9s 0.2s both;
}
.nc-gb-alias {
  margin-top: 14px;
  text-align: center;
  font-size: 7.5px;
  letter-spacing: 0.5em;
  color: rgba(240, 227, 194, 0.5);
  text-transform: uppercase;
  animation: nc-gb-in 0.9s 0.24s both;
}
.nc-gb-name {
  margin-top: 6px;
  text-align: center;
  font-size: clamp(48px, 15vw, 64px);
  font-weight: 400;
  line-height: 0.98;
  letter-spacing: -0.03em;
  background: linear-gradient(178deg, #fff4d6 0%, #ecd39a 38%, #b58a43 68%, #f5e3b7 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  overflow-wrap: anywhere;
  animation: nc-gb-in 1s 0.28s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-gb-role {
  margin-top: 10px;
  text-align: center;
  font-size: 8.5px;
  letter-spacing: 0.28em;
  color: rgba(240, 227, 194, 0.55);
  animation: nc-gb-in 0.9s 0.34s both;
}
.nc-gb-venue {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 9px;
  margin-top: 24px;
  animation: nc-gb-in 0.9s 0.4s both;
}
.nc-gb-venue-line {
  width: 24px;
  height: 1px;
  background: rgba(212, 175, 106, 0.55);
}
.nc-gb-venue-name {
  font-size: 11px;
  letter-spacing: 0.2em;
  color: rgba(240, 227, 194, 0.9);
  text-transform: uppercase;
}
.nc-gb-venue-city {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: rgba(240, 227, 194, 0.38);
}
.nc-gb-contact {
  margin-top: 34px;
  border-top: 1px solid rgba(212, 175, 106, 0.28);
  animation: nc-gb-in 1s 0.46s both;
}
.nc-gb-row {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 11px 2px;
  border-bottom: 1px solid rgba(212, 175, 106, 0.14);
  text-decoration: none;
  color: inherit;
}
.nc-gb-key {
  width: 74px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: rgba(212, 175, 106, 0.6);
}
.nc-gb-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11px;
  letter-spacing: 0.05em;
  color: rgba(240, 227, 194, 0.9);
}
.nc-gb-row-no {
  flex: none;
  font-family: Georgia, serif;
  font-size: 8px;
  color: rgba(212, 175, 106, 0.4);
}
.nc-gb-row-link {
  margin: 0 -6px;
  padding: 11px 8px;
  transition: background 0.3s ease, padding 0.3s ease;
}
.nc-gb-row-link:hover {
  background: rgba(212, 175, 106, 0.1);
  padding-left: 15px;
}
.nc-gb-row-link:active {
  background: rgba(212, 175, 106, 0.2);
}
.nc-gb-row-link:hover .nc-gb-value {
  color: #ffe9b8;
}
.nc-gb-empty {
  color: rgba(240, 227, 194, 0.28);
}
.nc-gb-note {
  margin-top: 22px;
  text-align: center;
  font-size: 9.5px;
  font-style: italic;
  letter-spacing: 0.05em;
  color: rgba(240, 227, 194, 0.35);
  animation: nc-gb-in 1s 0.56s both;
}
.nc-gb-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  padding-top: 14px;
  border-top: 1px solid rgba(212, 175, 106, 0.22);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: rgba(240, 227, 194, 0.28);
  animation: nc-gb-in 1s 0.66s both;
}
.nc-gb-foot-line {
  flex: 1;
  height: 1px;
  background: rgba(212, 175, 106, 0.2);
}
@keyframes nc-gb-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-gb-ingot {
  from { opacity: 0; transform: translateY(24px) rotateX(30deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-gb-shaft {
  from { opacity: 0; transform: translateX(-50%) scaleY(0.7); }
  to { opacity: 1; transform: translateX(-50%) scaleY(1); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-gb-root *,
  .nc-gb-root *::before,
  .nc-gb-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

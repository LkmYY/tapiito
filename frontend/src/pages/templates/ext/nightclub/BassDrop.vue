<template>
  <div class="nc-bd-root relative min-h-screen overflow-hidden px-5 py-7 text-[#e7ffd0]">
    <div class="nc-bd-haze"></div>
    <span class="nc-bd-ghost" aria-hidden="true">BASS</span>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3.5rem)] w-full max-w-md flex-col">
      <header class="nc-bd-head">
        <span class="nc-bd-brand"><span class="nc-bd-brand-mark"></span>BASS DROP</span>
        <span class="nc-bd-status">LIVE / {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-bd-main">
        <div class="nc-bd-marquee" aria-hidden="true">
          <span>NEON ROOM · {{ payload?.city || 'SHANGHAI' }} · AFTER HOURS ·</span>
        </div>

        <p class="nc-bd-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-bd-name">{{ payload?.name }}</h1>
        <p class="nc-bd-role">{{ payload?.role }}</p>

        <div class="nc-bd-wave" aria-hidden="true">
          <i v-for="n in 28" :key="n" :style="{ '--h': waveHeights[n - 1] + '%', '--d': n * 0.035 + 's' }"></i>
        </div>

        <div class="nc-bd-decks">
          <div class="nc-bd-deck">
            <div class="nc-bd-deck-head">
              <span>DECK A · VENUE</span>
              <span class="nc-bd-led"></span>
            </div>
            <p class="nc-bd-deck-value">{{ payload?.venue }}</p>
            <p class="nc-bd-deck-sub">{{ payload?.city || '—' }} / RESIDENT NIGHT</p>
            <div class="nc-bd-fader" aria-hidden="true"><span style="width: 72%"></span></div>
          </div>

          <div class="nc-bd-deck">
            <div class="nc-bd-deck-head">
              <span>DECK B · CONTACT</span>
              <span class="nc-bd-led nc-bd-led-alt"></span>
            </div>
            <div class="nc-bd-contact">
              <span class="nc-bd-key">WX</span>
              <span class="nc-bd-value">{{ payload?.wechat }}</span>
            </div>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-bd-contact nc-bd-contact-link">
              <span class="nc-bd-key">TEL</span>
              <span class="nc-bd-value">{{ payload.phone }}</span>
              <Phone :size="12" />
            </a>
            <div v-else class="nc-bd-contact">
              <span class="nc-bd-key">TEL</span>
              <span class="nc-bd-value nc-bd-muted"></span>
            </div>
            <div class="nc-bd-fader" aria-hidden="true"><span style="width: 46%"></span></div>
          </div>
        </div>
      </main>

      <footer class="nc-bd-foot">
        <span>MASTER OUT 128 BPM</span>
        <span class="nc-bd-eq" aria-hidden="true"><i></i><i></i><i></i><i></i></span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])

const waveHeights = [24, 48, 72, 38, 88, 56, 30, 68, 94, 44, 76, 34, 60, 98, 52, 28, 80, 46, 66, 90, 40, 58, 26, 74, 86, 50, 32, 62]
</script>

<style scoped>
.nc-bd-root {
  background: #070a05;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-bd-haze {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 50% 116%, rgba(200, 255, 47, 0.2), transparent 52%),
    radial-gradient(ellipse at 100% 0%, rgba(80, 120, 20, 0.16), transparent 42%);
  pointer-events: none;
}
.nc-bd-ghost {
  position: absolute;
  top: 74px;
  left: -16px;
  font-size: 150px;
  font-weight: 900;
  letter-spacing: -0.06em;
  line-height: 0.8;
  color: transparent;
  -webkit-text-stroke: 1px rgba(200, 255, 47, 0.09);
  pointer-events: none;
}
.nc-bd-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 13px;
  border-bottom: 1px solid rgba(200, 255, 47, 0.22);
  animation: nc-bd-in 0.7s both;
}
.nc-bd-brand {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 10px;
  font-weight: 800;
  letter-spacing: 0.3em;
  color: #c8ff2f;
}
.nc-bd-brand-mark {
  width: 7px;
  height: 14px;
  background: #c8ff2f;
  box-shadow: 0 0 12px rgba(200, 255, 47, 0.7);
}
.nc-bd-status {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(231, 255, 208, 0.4);
}
.nc-bd-main {
  flex: 1;
  padding-top: 30px;
}
.nc-bd-marquee {
  overflow: hidden;
  margin: 0 -5px;
  padding: 5px 0;
  border-top: 1px solid rgba(231, 255, 208, 0.14);
  border-bottom: 1px solid rgba(231, 255, 208, 0.14);
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.26em;
  color: rgba(231, 255, 208, 0.38);
}
.nc-bd-marquee span {
  display: inline-block;
  padding-left: 4px;
  animation: nc-bd-scroll 16s linear infinite;
}
.nc-bd-alias {
  margin-top: 34px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.48em;
  color: #c8ff2f;
  animation: nc-bd-in 0.8s 0.08s both;
}
.nc-bd-name {
  margin-top: 4px;
  font-size: clamp(62px, 20vw, 86px);
  font-weight: 900;
  line-height: 0.84;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  color: #efffdc;
  text-shadow: 0 0 34px rgba(200, 255, 47, 0.24);
  overflow-wrap: anywhere;
  animation: nc-bd-in 0.9s 0.12s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-bd-role {
  margin-top: 12px;
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.3em;
  color: rgba(231, 255, 208, 0.7);
  animation: nc-bd-in 0.9s 0.2s both;
}
.nc-bd-wave {
  display: flex;
  height: 64px;
  align-items: center;
  gap: 3px;
  margin: 26px 0 24px;
  animation: nc-bd-in 0.9s 0.26s both;
}
.nc-bd-wave i {
  flex: 1;
  height: var(--h);
  min-height: 4px;
  background: linear-gradient(180deg, #e8ff9a, #c8ff2f 40%, #5c8a00);
  transform-origin: center;
  animation: nc-bd-bar 1.8s var(--d) ease-in-out infinite alternate;
}
.nc-bd-decks {
  display: flex;
  flex-direction: column;
  gap: 10px;
  animation: nc-bd-in 0.9s 0.34s both;
}
.nc-bd-deck {
  padding: 14px 15px 12px;
  border: 1px solid rgba(231, 255, 208, 0.14);
  background: linear-gradient(145deg, rgba(231, 255, 208, 0.06), rgba(231, 255, 208, 0.015));
  transition: border-color 0.3s ease, transform 0.3s ease, background 0.3s ease;
}
.nc-bd-deck:hover {
  border-color: rgba(200, 255, 47, 0.48);
  background: linear-gradient(145deg, rgba(200, 255, 47, 0.08), rgba(231, 255, 208, 0.02));
  transform: translateY(-2px);
}
.nc-bd-deck-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(231, 255, 208, 0.4);
}
.nc-bd-led {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #c8ff2f;
  box-shadow: 0 0 8px #c8ff2f;
  animation: nc-bd-blink 2.2s ease-in-out infinite;
}
.nc-bd-led-alt {
  background: #ff7a45;
  box-shadow: 0 0 8px #ff7a45;
  animation-delay: 0.8s;
}
.nc-bd-deck-value {
  margin-top: 9px;
  font-size: 17px;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #efffdc;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-bd-deck-sub {
  margin-top: 3px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: rgba(231, 255, 208, 0.38);
}
.nc-bd-contact {
  display: flex;
  align-items: center;
  gap: 9px;
  padding: 7px 0;
  border-bottom: 1px dotted rgba(231, 255, 208, 0.14);
  text-decoration: none;
  color: inherit;
}
.nc-bd-contact-link {
  margin: 0 -6px;
  padding: 7px 6px;
  transition: background 0.25s ease, padding 0.25s ease;
}
.nc-bd-contact-link:hover {
  background: rgba(200, 255, 47, 0.1);
  padding-left: 12px;
  color: #c8ff2f;
}
.nc-bd-contact-link:active {
  background: rgba(200, 255, 47, 0.2);
}
.nc-bd-key {
  width: 26px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.1em;
  color: rgba(200, 255, 47, 0.6);
}
.nc-bd-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  letter-spacing: 0.04em;
  color: rgba(239, 255, 220, 0.86);
}
.nc-bd-muted {
  color: rgba(231, 255, 208, 0.3);
}
.nc-bd-fader {
  height: 2px;
  margin-top: 10px;
  background: rgba(231, 255, 208, 0.1);
}
.nc-bd-fader span {
  display: block;
  height: 100%;
  background: linear-gradient(90deg, #6d9510, #c8ff2f);
  box-shadow: 0 0 8px rgba(200, 255, 47, 0.5);
}
.nc-bd-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 18px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(231, 255, 208, 0.3);
  animation: nc-bd-in 1s 0.55s both;
}
.nc-bd-eq {
  display: inline-flex;
  align-items: flex-end;
  gap: 2px;
  height: 9px;
}
.nc-bd-eq i {
  width: 2px;
  background: #c8ff2f;
  animation: nc-bd-eq 0.9s ease-in-out infinite alternate;
}
.nc-bd-eq i:nth-child(1) { height: 4px; animation-delay: 0.1s; }
.nc-bd-eq i:nth-child(2) { height: 8px; animation-delay: 0.3s; }
.nc-bd-eq i:nth-child(3) { height: 6px; animation-delay: 0.2s; }
.nc-bd-eq i:nth-child(4) { height: 9px; animation-delay: 0.4s; }
@keyframes nc-bd-in {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-bd-scroll {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@keyframes nc-bd-bar {
  from { transform: scaleY(0.55); opacity: 0.7; }
  to { transform: scaleY(1.08); opacity: 1; }
}
@keyframes nc-bd-blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}
@keyframes nc-bd-eq {
  from { transform: scaleY(0.6); }
  to { transform: scaleY(1); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-bd-root *,
  .nc-bd-root *::before,
  .nc-bd-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

<template>
  <div class="nc-fn-root relative min-h-screen overflow-hidden px-4 py-5 text-[#efe0c4]">
    <div class="nc-fn-grain"></div>
    <div class="nc-fn-sprocket nc-fn-sprocket-top"></div>
    <div class="nc-fn-sprocket nc-fn-sprocket-bottom"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-2.5rem)] w-full max-w-md flex-col">
      <header class="nc-fn-head">
        <span>ROLL 01 · 35MM</span>
        <span class="nc-fn-rec"><i></i>REC</span>
        <span>ISO 800</span>
      </header>

      <main class="nc-fn-main">
        <div class="nc-fn-frame nc-fn-frame-hero">
          <span class="nc-fn-frame-no">01A</span>
          <div class="nc-fn-hero-copy">
            <p class="nc-fn-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
            <h1 class="nc-fn-name">{{ payload?.name }}</h1>
            <p class="nc-fn-role">{{ payload?.role }}</p>
          </div>
          <p class="nc-fn-tc">23:47:08 / +08</p>
        </div>

        <div class="nc-fn-venue-frame">
          <div class="nc-fn-venue-left">
            <span class="nc-fn-frame-no">02A</span>
            <p class="nc-fn-venue">{{ payload?.venue }}</p>
            <p class="nc-fn-city">{{ payload?.city || '—' }}</p>
          </div>
          <p class="nc-fn-venue-note">THE ROOM<br />AFTER MIDNIGHT</p>
        </div>

        <section class="nc-fn-sheet">
          <div class="nc-fn-sheet-head">
            <span>CONTACT SHEET</span>
            <span>SELECTED FRAMES</span>
          </div>

          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-fn-shot nc-fn-shot-link">
            <div class="nc-fn-shot-visual">
              <Phone :size="18" />
              <span class="nc-fn-shot-no">03A</span>
            </div>
            <div class="nc-fn-shot-copy">
              <small>PHONE / 电话</small>
              <b>{{ payload.phone }}</b>
            </div>
            <span class="nc-fn-shot-arrow">→</span>
          </a>
          <div v-else class="nc-fn-shot">
            <div class="nc-fn-shot-visual">
              <Phone :size="18" />
              <span class="nc-fn-shot-no">03A</span>
            </div>
            <div class="nc-fn-shot-copy">
              <small>PHONE / 电话</small>
              <b class="nc-fn-empty"></b>
            </div>
          </div>

          <div class="nc-fn-shot">
            <div class="nc-fn-shot-visual">
              <MessageCircle :size="18" />
              <span class="nc-fn-shot-no">03B</span>
            </div>
            <div class="nc-fn-shot-copy">
              <small>WECHAT / 微信</small>
              <b>{{ payload?.wechat }}</b>
            </div>
          </div>
        </section>

        <div class="nc-fn-strip-copy">
          <span>FRAME 04 — {{ payload?.name }}, {{ payload?.city || '—' }}</span>
          <span>NO FLASH</span>
        </div>
      </main>

      <footer class="nc-fn-foot">
        <span>DEVELOPED AT 04:00</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { Phone, MessageCircle } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-fn-root {
  background: #0b0906;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-fn-grain {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 68% 24%, rgba(217, 160, 91, 0.14), transparent 48%),
    radial-gradient(rgba(239, 224, 196, 0.16) 0.5px, transparent 0.5px);
  background-size: auto, 3px 3px;
  opacity: 0.5;
  pointer-events: none;
}
.nc-fn-sprocket {
  position: absolute;
  left: 0;
  right: 0;
  height: 13px;
  background:
    repeating-linear-gradient(90deg, transparent 0 8px, #0b0906 8px 10px, rgba(239, 224, 196, 0.28) 10px 19px, #0b0906 19px 21px);
  border-top: 1px solid rgba(239, 224, 196, 0.16);
  border-bottom: 1px solid rgba(239, 224, 196, 0.16);
  pointer-events: none;
}
.nc-fn-sprocket-top { top: 4px; }
.nc-fn-sprocket-bottom { bottom: 4px; }
.nc-fn-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 14px 2px 10px;
  border-bottom: 1px solid rgba(239, 224, 196, 0.24);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(239, 224, 196, 0.5);
  animation: nc-fn-in 0.8s both;
}
.nc-fn-rec {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  color: #d9663f;
}
.nc-fn-rec i {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #d9663f;
  animation: nc-fn-blink 1.3s ease-in-out infinite;
}
.nc-fn-main {
  flex: 1;
  padding-top: 18px;
}
.nc-fn-frame {
  position: relative;
  border: 1px solid rgba(239, 224, 196, 0.28);
  background:
    linear-gradient(160deg, rgba(217, 160, 91, 0.22), rgba(11, 9, 6, 0.6) 58%),
    rgba(217, 160, 91, 0.06);
}
.nc-fn-frame-hero {
  min-height: 258px;
  padding: 18px;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  overflow: hidden;
  animation: nc-fn-expose 1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-fn-frame-hero::before {
  content: '';
  position: absolute;
  top: -60px;
  right: -40px;
  width: 220px;
  height: 220px;
  border-radius: 50%;
  background: radial-gradient(circle at 40% 35%, rgba(217, 160, 91, 0.55), rgba(120, 55, 20, 0.12) 55%, transparent 70%);
  filter: blur(4px);
  animation: nc-fn-glow 8s ease-in-out infinite;
}
.nc-fn-frame-hero::after {
  content: '';
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 2px, rgba(0, 0, 0, 0.15) 2px 3px);
  pointer-events: none;
}
.nc-fn-frame-no {
  position: absolute;
  top: 9px;
  left: 11px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.12em;
  color: #d9a05b;
}
.nc-fn-hero-copy {
  position: relative;
  z-index: 2;
}
.nc-fn-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.4em;
  color: #d9a05b;
  text-transform: uppercase;
}
.nc-fn-name {
  margin-top: 5px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(48px, 15vw, 64px);
  font-weight: 400;
  line-height: 0.94;
  letter-spacing: -0.045em;
  color: #f6e8ce;
  text-shadow: 0 0 30px rgba(217, 160, 91, 0.4);
  overflow-wrap: anywhere;
}
.nc-fn-role {
  margin-top: 10px;
  font-size: 8.5px;
  letter-spacing: 0.26em;
  color: rgba(239, 224, 196, 0.65);
}
.nc-fn-tc {
  position: absolute;
  right: 12px;
  bottom: 10px;
  z-index: 2;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.1em;
  color: rgba(239, 224, 196, 0.45);
}
.nc-fn-venue-frame {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
  margin-top: 10px;
  padding: 15px 16px;
  border: 1px solid rgba(239, 224, 196, 0.18);
  animation: nc-fn-in 0.9s 0.3s both;
}
.nc-fn-venue-left {
  position: relative;
  min-width: 0;
  padding-top: 13px;
}
.nc-fn-venue {
  font-size: 21px;
  font-weight: 600;
  line-height: 1.05;
  letter-spacing: 0.02em;
  color: #f0dfc1;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-fn-city {
  margin-top: 4px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: #d9a05b;
}
.nc-fn-venue-note {
  flex: none;
  padding-left: 12px;
  border-left: 1px solid rgba(239, 224, 196, 0.2);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  line-height: 1.7;
  letter-spacing: 0.1em;
  color: rgba(239, 224, 196, 0.38);
  text-align: right;
}
.nc-fn-sheet {
  margin-top: 22px;
  border-top: 1px solid rgba(239, 224, 196, 0.24);
  animation: nc-fn-in 0.95s 0.4s both;
}
.nc-fn-sheet-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 0;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(239, 224, 196, 0.42);
}
.nc-fn-shot {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 11px 0;
  border-top: 1px dashed rgba(239, 224, 196, 0.18);
  text-decoration: none;
  color: inherit;
}
.nc-fn-shot-visual {
  position: relative;
  display: flex;
  width: 62px;
  height: 44px;
  flex: none;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(217, 160, 91, 0.35);
  background: linear-gradient(140deg, rgba(217, 160, 91, 0.2), rgba(217, 160, 91, 0.03));
  color: #d9a05b;
  transition: background 0.3s ease, color 0.3s ease;
}
.nc-fn-shot-no {
  position: absolute;
  top: 3px;
  left: 4px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.08em;
  color: rgba(239, 224, 196, 0.5);
}
.nc-fn-shot-copy {
  display: flex;
  min-width: 0;
  flex: 1;
  flex-direction: column;
  gap: 3px;
}
.nc-fn-shot-copy small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(239, 224, 196, 0.38);
}
.nc-fn-shot-copy b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.04em;
  color: rgba(246, 232, 206, 0.9);
}
.nc-fn-shot-arrow {
  flex: none;
  color: rgba(217, 160, 91, 0.6);
  transition: transform 0.3s ease;
}
.nc-fn-shot-link {
  transition: background 0.3s ease;
}
.nc-fn-shot-link:hover {
  background: rgba(217, 160, 91, 0.07);
}
.nc-fn-shot-link:active {
  background: rgba(217, 160, 91, 0.14);
}
.nc-fn-shot-link:hover .nc-fn-shot-visual {
  background: #d9a05b;
  color: #0b0906;
}
.nc-fn-shot-link:hover .nc-fn-shot-arrow {
  transform: translateX(4px);
}
.nc-fn-empty {
  color: rgba(239, 224, 196, 0.28) !important;
}
.nc-fn-strip-copy {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  margin-top: 16px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.12em;
  color: rgba(239, 224, 196, 0.3);
  animation: nc-fn-in 1s 0.55s both;
}
.nc-fn-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 2px 2px;
  border-top: 1px solid rgba(239, 224, 196, 0.18);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(239, 224, 196, 0.3);
  animation: nc-fn-in 1s 0.65s both;
}
@keyframes nc-fn-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-fn-expose {
  from { opacity: 0; filter: brightness(2.4) blur(6px); transform: scale(1.03); }
  to { opacity: 1; filter: none; transform: none; }
}
@keyframes nc-fn-glow {
  0%, 100% { opacity: 0.8; transform: translate(0, 0); }
  50% { opacity: 1; transform: translate(-10px, 8px); }
}
@keyframes nc-fn-blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.2; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-fn-root *,
  .nc-fn-root *::before,
  .nc-fn-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

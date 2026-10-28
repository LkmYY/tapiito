<template>
  <div class="nc-va-root relative min-h-screen overflow-hidden px-5 py-8 text-[#1b1510]">
    <div class="nc-va-paper"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-va-head">
        <span>AFTER HOURS RECORDS</span>
        <span>33⅓ RPM · STEREO</span>
      </header>

      <section class="nc-va-hero">
        <div class="nc-va-disc" aria-hidden="true">
          <div class="nc-va-grooves"></div>
          <div class="nc-va-label">
            <span>AH</span>
            <small>{{ payload?.city || '—' }}</small>
          </div>
          <span class="nc-va-hole"></span>
        </div>

        <div class="nc-va-sleeve">
          <div class="nc-va-sleeve-top">
            <span>AH—024</span>
            <span>12"</span>
          </div>
          <p class="nc-va-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
          <h1 class="nc-va-name">{{ payload?.name }}</h1>
          <p class="nc-va-role">{{ payload?.role }}</p>
          <div class="nc-va-sleeve-line"></div>
          <p class="nc-va-venue">
            {{ payload?.venue }}
            <span>/ {{ payload?.city || '—' }}</span>
          </p>
          <span class="nc-va-sticker">LATE SET</span>
        </div>

        <p class="nc-va-side">SIDE A · NIGHT SERVICE</p>
      </section>

      <section class="nc-va-tracklist">
        <div class="nc-va-track">
          <span class="nc-va-no">A1</span>
          <span class="nc-va-track-copy"><small>身份 / ROLE</small><b>{{ payload?.role }}</b></span>
          <span class="nc-va-time">4:12</span>
        </div>
        <div class="nc-va-track">
          <span class="nc-va-no">A2</span>
          <span class="nc-va-track-copy"><small>场地 / VENUE</small><b>{{ payload?.venue }} · {{ payload?.city || '—' }}</b></span>
          <span class="nc-va-time">5:08</span>
        </div>
        <div class="nc-va-track">
          <span class="nc-va-no">B1</span>
          <span class="nc-va-track-copy"><small>微信 / WECHAT</small><b>{{ payload?.wechat }}</b></span>
          <span class="nc-va-time">3:47</span>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-va-track nc-va-track-link">
          <span class="nc-va-no">B2</span>
          <span class="nc-va-track-copy"><small>电话 / PHONE</small><b>{{ payload.phone }}</b></span>
          <span class="nc-va-play"><Phone :size="13" /></span>
        </a>
        <div v-else class="nc-va-track">
          <span class="nc-va-no">B2</span>
          <span class="nc-va-track-copy"><small>电话 / PHONE</small><b class="nc-va-muted"></b></span>
          <span class="nc-va-time">—</span>
        </div>
      </section>

      <footer class="nc-va-foot">
        <span>PLAY LOUD · LEAVE LATE</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-va-root {
  background: #e9e2d4;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-va-paper {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 88% 8%, rgba(255, 255, 255, 0.7), transparent 45%),
    radial-gradient(rgba(27, 21, 16, 0.14) 0.6px, transparent 0.6px),
    linear-gradient(115deg, rgba(120, 90, 50, 0.06), transparent 40%);
  background-size: auto, 4px 4px, auto;
  opacity: 0.55;
  pointer-events: none;
}
.nc-va-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 11px;
  border-bottom: 1px solid rgba(27, 21, 16, 0.55);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.2em;
  color: #4b3c2d;
  animation: nc-va-in 0.8s both;
}
.nc-va-hero {
  position: relative;
  padding: 30px 0 26px;
}
.nc-va-disc {
  position: absolute;
  top: 18px;
  right: -46px;
  width: 250px;
  height: 250px;
  border-radius: 50%;
  background:
    repeating-radial-gradient(circle at 50% 50%, #17130f 0 1.5px, #29231d 1.5px 3px);
  box-shadow: inset 0 0 0 1px #3b332a, 0 22px 40px -18px rgba(27, 21, 16, 0.7);
  animation: nc-va-spin 22s linear infinite;
  pointer-events: none;
}
.nc-va-grooves {
  position: absolute;
  inset: 12px;
  border-radius: 50%;
  background: repeating-radial-gradient(circle at 50% 50%, transparent 0 6px, rgba(255, 255, 255, 0.06) 6px 7px);
}
.nc-va-label {
  position: absolute;
  top: 50%;
  left: 50%;
  display: flex;
  width: 86px;
  height: 86px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  border-radius: 50%;
  background: #d8572c;
  color: #f4ead9;
  transform: translate(-50%, -50%);
}
.nc-va-label span {
  font-family: Georgia, serif;
  font-size: 22px;
  line-height: 1;
  letter-spacing: 0.04em;
}
.nc-va-label small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  opacity: 0.8;
}
.nc-va-hole {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #e9e2d4;
  transform: translate(-50%, -50%);
  box-shadow: 0 0 0 2px rgba(23, 19, 15, 0.7);
}
.nc-va-sleeve {
  position: relative;
  width: 79%;
  margin-top: 38px;
  padding: 20px 20px 26px;
  border: 1px solid rgba(255, 255, 255, 0.11);
  background:
    linear-gradient(152deg, rgba(255, 255, 255, 0.06), transparent 33%),
    #17130f;
  box-shadow: 0 26px 46px -26px rgba(27, 21, 16, 0.8);
  color: #f0e7d8;
  animation: nc-va-sleeve 1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.4s ease;
}
.nc-va-sleeve:hover {
  transform: translateY(-4px) rotate(-0.4deg);
}
.nc-va-sleeve-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: rgba(240, 231, 216, 0.42);
}
.nc-va-alias {
  margin-top: 30px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.42em;
  color: #d8572c;
  text-transform: uppercase;
}
.nc-va-name {
  margin-top: 7px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(46px, 14vw, 62px);
  font-weight: 400;
  line-height: 0.94;
  letter-spacing: -0.05em;
  overflow-wrap: anywhere;
}
.nc-va-role {
  margin-top: 12px;
  font-size: 10px;
  letter-spacing: 0.18em;
  color: #d9cbb6;
}
.nc-va-sleeve-line {
  width: 34px;
  height: 2px;
  margin: 18px 0 14px;
  background: #d8572c;
}
.nc-va-venue {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  line-height: 1.7;
  letter-spacing: 0.12em;
  color: rgba(240, 231, 216, 0.72);
  text-transform: uppercase;
}
.nc-va-venue span {
  color: rgba(240, 231, 216, 0.38);
}
.nc-va-sticker {
  position: absolute;
  right: -16px;
  bottom: 16px;
  padding: 4px 9px;
  background: #e4b74f;
  color: #17130f;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  transform: rotate(-6deg);
  box-shadow: 0 2px 0 rgba(0, 0, 0, 0.3);
}
.nc-va-side {
  margin-top: 24px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.28em;
  color: #77644e;
}
.nc-va-tracklist {
  border-top: 1px solid rgba(27, 21, 16, 0.35);
  animation: nc-va-in 0.9s 0.36s both;
}
.nc-va-track {
  display: flex;
  align-items: center;
  gap: 13px;
  padding: 13px 0;
  border-bottom: 1px solid rgba(27, 21, 16, 0.16);
  text-decoration: none;
  color: inherit;
}
.nc-va-no {
  width: 24px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  color: #a58e73;
}
.nc-va-track-copy {
  display: flex;
  min-width: 0;
  flex: 1;
  flex-direction: column;
  gap: 3px;
}
.nc-va-track-copy small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: #9a8469;
}
.nc-va-track-copy b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  font-weight: 500;
  letter-spacing: 0.03em;
  color: #241b13;
}
.nc-va-time {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  color: #a58e73;
}
.nc-va-track-link {
  margin: 0 -8px;
  padding: 13px 8px;
  transition: background 0.25s ease, padding 0.25s ease;
}
.nc-va-track-link:hover {
  background: rgba(216, 87, 44, 0.1);
  padding-left: 14px;
}
.nc-va-track-link:active {
  background: rgba(216, 87, 44, 0.2);
}
.nc-va-track-link:hover b {
  color: #b8431f;
}
.nc-va-play {
  display: flex;
  width: 26px;
  height: 26px;
  flex: none;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(27, 21, 16, 0.3);
  border-radius: 50%;
  color: #241b13;
  transition: background 0.25s ease, color 0.25s ease, transform 0.25s ease;
}
.nc-va-track-link:hover .nc-va-play {
  background: #d8572c;
  border-color: #d8572c;
  color: #f4ead9;
  transform: scale(1.08);
}
.nc-va-muted {
  color: #b3a08b !important;
}
.nc-va-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: auto;
  padding-top: 20px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #8b7459;
  animation: nc-va-in 1s 0.6s both;
}
@keyframes nc-va-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-va-sleeve {
  from { opacity: 0; transform: translateY(28px) rotate(2deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-va-spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-va-root *,
  .nc-va-root *::before,
  .nc-va-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

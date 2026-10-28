<template>
  <div class="nc-ah-root relative min-h-screen overflow-hidden px-5 py-7 text-[#0b0b06]">
    <div class="nc-ah-halftone"></div>
    <div class="nc-ah-border"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3.5rem)] w-full max-w-md flex-col">
      <header class="nc-ah-head">
        <span>ACID HOUSE</span>
        <span class="nc-ah-head-mid">ALL NIGHT / NO PHOTOS</span>
        <span>98.7</span>
      </header>

      <main class="nc-ah-main">
        <div class="nc-ah-smiley" aria-hidden="true">
          <span class="nc-ah-eye nc-ah-eye-l"></span>
          <span class="nc-ah-eye nc-ah-eye-r"></span>
          <span class="nc-ah-mouth"></span>
        </div>

        <p class="nc-ah-kicker">A PARTY IN THREE COLOURS</p>

        <div class="nc-ah-title">
          <h1 class="nc-ah-name">{{ payload?.name }}</h1>
          <span class="nc-ah-title-shadow" aria-hidden="true">{{ payload?.name }}</span>
        </div>

        <div class="nc-ah-sub">
          <span class="nc-ah-alias">{{ payload?.alias || 'NO ALIAS' }}</span>
          <span class="nc-ah-slash">//</span>
          <span class="nc-ah-role">{{ payload?.role }}</span>
        </div>

        <div class="nc-ah-venue">
          <span class="nc-ah-venue-label">AT</span>
          <span class="nc-ah-venue-name">{{ payload?.venue }}</span>
          <span class="nc-ah-venue-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-ah-contacts">
          <div class="nc-ah-contact">
            <span class="nc-ah-num">01</span>
            <span class="nc-ah-key">WECHAT</span>
            <span class="nc-ah-value">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-ah-contact nc-ah-contact-link">
            <span class="nc-ah-num">02</span>
            <span class="nc-ah-key">PHONE</span>
            <span class="nc-ah-value">{{ payload.phone }}</span>
            <span class="nc-ah-dial">DIAL</span>
          </a>
          <div v-else class="nc-ah-contact">
            <span class="nc-ah-num">02</span>
            <span class="nc-ah-key">PHONE</span>
            <span class="nc-ah-value nc-ah-empty"></span>
          </div>
        </section>

        <p class="nc-ah-footer-line">SMILE · SWEAT · REPEAT</p>
      </main>

      <footer class="nc-ah-foot">
        <span>FRI 23:00 — SUN 06:00</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-ah-root {
  background: #d9ff00;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-ah-halftone {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(11, 11, 6, 0.18) 0.7px, transparent 0.8px);
  background-size: 4px 4px;
  opacity: 0.5;
  pointer-events: none;
}
.nc-ah-border {
  position: absolute;
  inset: 10px;
  border: 3px solid #0b0b06;
  pointer-events: none;
}
.nc-ah-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  padding: 2px 2px 10px;
  border-bottom: 3px solid #0b0b06;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  font-weight: 700;
  letter-spacing: 0.16em;
  animation: nc-ah-in 0.7s both;
}
.nc-ah-head-mid {
  flex: 1;
  text-align: center;
  font-size: 5.5px;
  color: rgba(11, 11, 6, 0.55);
}
.nc-ah-main {
  position: relative;
  flex: 1;
  padding-top: 34px;
}
.nc-ah-smiley {
  position: absolute;
  top: 10px;
  right: -34px;
  width: 178px;
  height: 178px;
  border: 3px solid #0b0b06;
  border-radius: 50%;
  background: #0b0b06;
  animation: nc-ah-spin-in 1.2s 0.08s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-ah-smiley::before {
  content: '';
  position: absolute;
  inset: 8px;
  border: 1px solid rgba(217, 255, 0, 0.5);
  border-radius: 50%;
}
.nc-ah-eye {
  position: absolute;
  top: 52px;
  width: 18px;
  height: 30px;
  border-radius: 50%;
  background: #d9ff00;
}
.nc-ah-eye-l { left: 48px; transform: rotate(-12deg); }
.nc-ah-eye-r { right: 48px; transform: rotate(12deg); }
.nc-ah-mouth {
  position: absolute;
  bottom: 36px;
  left: 50%;
  width: 88px;
  height: 40px;
  border-bottom: 8px solid #d9ff00;
  border-radius: 0 0 50% 50%;
  transform: translateX(-50%);
}
.nc-ah-kicker {
  position: relative;
  z-index: 2;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  font-weight: 700;
  letter-spacing: 0.28em;
  animation: nc-ah-in 0.8s 0.06s both;
}
.nc-ah-title {
  position: relative;
  z-index: 2;
  margin-top: 20px;
  margin-left: -3px;
  transform: rotate(-3deg);
}
.nc-ah-name {
  position: relative;
  z-index: 2;
  font-size: clamp(60px, 20vw, 88px);
  font-weight: 900;
  line-height: 0.8;
  letter-spacing: -0.07em;
  text-transform: uppercase;
  color: #0b0b06;
  overflow-wrap: anywhere;
  animation: nc-ah-slam 0.7s 0.12s cubic-bezier(0.2, 1.3, 0.4, 1) both;
}
.nc-ah-title-shadow {
  position: absolute;
  top: 6px;
  left: 7px;
  z-index: 1;
  font-size: clamp(60px, 20vw, 88px);
  font-weight: 900;
  line-height: 0.8;
  letter-spacing: -0.07em;
  text-transform: uppercase;
  color: transparent;
  -webkit-text-stroke: 2px rgba(11, 11, 6, 0.35);
  pointer-events: none;
}
.nc-ah-sub {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 8px;
  margin-top: 24px;
  animation: nc-ah-in 0.8s 0.26s both;
}
.nc-ah-alias {
  padding: 3px 7px;
  background: #0b0b06;
  color: #d9ff00;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  font-weight: 700;
  letter-spacing: 0.18em;
  text-transform: uppercase;
}
.nc-ah-slash {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 9px;
  font-weight: 700;
  opacity: 0.4;
}
.nc-ah-role {
  font-size: 9.5px;
  font-weight: 800;
  letter-spacing: 0.18em;
  text-transform: uppercase;
}
.nc-ah-venue {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  gap: 7px;
  margin-top: 20px;
  animation: nc-ah-in 0.8s 0.32s both;
}
.nc-ah-venue-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  font-weight: 700;
  letter-spacing: 0.18em;
}
.nc-ah-venue-name {
  font-size: 15px;
  font-weight: 900;
  letter-spacing: -0.01em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-ah-venue-city {
  padding-left: 7px;
  border-left: 2px solid #0b0b06;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  font-weight: 700;
  letter-spacing: 0.12em;
}
.nc-ah-contacts {
  margin-top: 28px;
  border: 3px solid #0b0b06;
  background: #0b0b06;
  animation: nc-ah-in 0.9s 0.38s both;
  transition: transform 0.3s ease;
}
.nc-ah-contacts:hover {
  transform: rotate(-0.4deg);
}
.nc-ah-contact {
  display: flex;
  align-items: center;
  gap: 9px;
  padding: 11px 10px;
  border-bottom: 1px solid rgba(217, 255, 0, 0.2);
  color: #d9ff00;
  text-decoration: none;
}
.nc-ah-contact:last-child {
  border-bottom: 0;
}
.nc-ah-num {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  opacity: 0.45;
}
.nc-ah-key {
  width: 48px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  font-weight: 700;
  letter-spacing: 0.12em;
  opacity: 0.55;
}
.nc-ah-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.03em;
}
.nc-ah-dial {
  flex: none;
  padding: 2px 5px;
  background: #d9ff00;
  color: #0b0b06;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  font-weight: 700;
  letter-spacing: 0.1em;
}
.nc-ah-contact-link {
  transition: background 0.2s steps(2, end), padding 0.2s steps(2, end);
}
.nc-ah-contact-link:hover {
  background: #d9ff00;
  color: #0b0b06;
  padding-left: 16px;
}
.nc-ah-contact-link:hover .nc-ah-dial {
  background: #0b0b06;
  color: #d9ff00;
}
.nc-ah-contact-link:active {
  background: #b8dd00;
}
.nc-ah-empty {
  opacity: 0.35;
}
.nc-ah-footer-line {
  margin-top: 18px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  font-weight: 700;
  letter-spacing: 0.28em;
  animation: nc-ah-in 1s 0.5s both;
}
.nc-ah-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 2px 2px;
  border-top: 3px solid #0b0b06;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  font-weight: 700;
  letter-spacing: 0.14em;
  animation: nc-ah-in 1s 0.58s both;
}
@keyframes nc-ah-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-ah-slam {
  from { opacity: 0; transform: scale(1.3) rotate(8deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-ah-spin-in {
  from { opacity: 0; transform: rotate(-120deg) scale(0.6); }
  to { opacity: 1; transform: rotate(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-ah-root *,
  .nc-ah-root *::before,
  .nc-ah-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

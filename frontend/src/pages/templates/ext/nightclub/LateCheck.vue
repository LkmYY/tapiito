<template>
  <div class="nc-lc-root relative min-h-screen overflow-hidden px-6 py-10 text-[#152019]">
    <div class="nc-lc-counter"></div>
    <div class="nc-lc-glow"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-5rem)] w-full max-w-md flex-col items-center justify-center">
      <p class="nc-lc-pre">NEON ROOM · NIGHT RECEIPT</p>

      <article class="nc-lc-receipt">
        <div class="nc-lc-zigzag nc-lc-zigzag-top" aria-hidden="true"></div>

        <header class="nc-lc-head">
          <span class="nc-lc-logo">NEON / ROOM</span>
          <p class="nc-lc-address">24H LOUNGE · {{ payload?.city || 'SHANGHAI' }}</p>
          <p class="nc-lc-address">TEL 24H · OPEN UNTIL LATE</p>
        </header>

        <div class="nc-lc-meta">
          <span>ORDER #0417</span>
          <span>03:12 AM</span>
          <span>NFC {{ nfcId || '0000' }}</span>
        </div>

        <div class="nc-lc-rule"></div>

        <h1 class="nc-lc-name">{{ payload?.name }}</h1>
        <p class="nc-lc-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <p class="nc-lc-role">{{ payload?.role }}</p>

        <div class="nc-lc-rule"></div>

        <section class="nc-lc-items">
          <div class="nc-lc-item">
            <span class="nc-lc-qty">1x</span>
            <span class="nc-lc-desc">RESIDENCY<br /><small>{{ payload?.venue }}</small></span>
            <span class="nc-lc-price">--</span>
          </div>
          <div class="nc-lc-item">
            <span class="nc-lc-qty">1x</span>
            <span class="nc-lc-desc">CITY<br /><small>{{ payload?.city || '—' }}</small></span>
            <span class="nc-lc-price">--</span>
          </div>
          <div class="nc-lc-item">
            <span class="nc-lc-qty">1x</span>
            <span class="nc-lc-desc">WECHAT<br /><small>{{ payload?.wechat }}</small></span>
            <span class="nc-lc-price">--</span>
          </div>
        </section>

        <div class="nc-lc-rule nc-lc-rule-double"></div>

        <div class="nc-lc-total">
          <span>TOTAL / 电话</span>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-lc-phone">{{ payload.phone }}</a>
        </div>

        <p class="nc-lc-thanks">THANK YOU · COME BACK LATE</p>
        <p class="nc-lc-tip">TIP: ONE MORE TRACK · ONE MORE ROUND</p>

        <div class="nc-lc-barcode" aria-hidden="true"></div>
        <p class="nc-lc-code">0417 · 8000 · {{ payload?.wechat || 'NIGHT' }}</p>

        <div class="nc-lc-zigzag nc-lc-zigzag-bottom" aria-hidden="true"></div>
      </article>

      <p class="nc-lc-after">KEEP THE RECEIPT · FORGET THE TIME</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-lc-root {
  background: #0f2620;
  font-family: 'SFMono-Regular', Consolas, 'Courier New', monospace;
}
.nc-lc-counter {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 50% 20%, rgba(190, 255, 197, 0.12), transparent 50%),
    repeating-linear-gradient(90deg, rgba(255, 255, 255, 0.025) 0 2px, transparent 2px 40px);
  pointer-events: none;
}
.nc-lc-glow {
  position: absolute;
  bottom: -80px;
  left: 50%;
  width: 360px;
  height: 260px;
  border-radius: 50%;
  background: rgba(120, 220, 180, 0.13);
  filter: blur(50px);
  transform: translateX(-50%);
  pointer-events: none;
}
.nc-lc-pre {
  margin-bottom: 16px;
  font-size: 6.5px;
  letter-spacing: 0.26em;
  color: rgba(221, 235, 219, 0.45);
  animation: nc-lc-in 0.8s both;
}
.nc-lc-receipt {
  position: relative;
  width: 100%;
  padding: 26px 22px 22px;
  background:
    radial-gradient(rgba(21, 32, 25, 0.08) 0.5px, transparent 0.5px),
    #f3f0e4;
  background-size: 4px 4px, auto;
  box-shadow: 0 30px 70px -22px rgba(0, 0, 0, 0.8);
  animation: nc-lc-print 1.2s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.nc-lc-receipt:hover {
  transform: translateY(-4px) rotate(-0.3deg);
  box-shadow: 0 40px 82px -22px rgba(0, 0, 0, 0.85);
}
.nc-lc-zigzag {
  position: absolute;
  left: 0;
  right: 0;
  height: 10px;
  background:
    linear-gradient(135deg, transparent 50%, #f3f0e4 50%) 0 0 / 12px 12px repeat-x,
    linear-gradient(45deg, transparent 50%, #f3f0e4 50%) 0 0 / 12px 12px repeat-x;
}
.nc-lc-zigzag-top {
  top: -9px;
  transform: rotate(180deg);
}
.nc-lc-zigzag-bottom {
  bottom: -9px;
}
.nc-lc-head {
  text-align: center;
}
.nc-lc-logo {
  font-size: 18px;
  font-weight: 700;
  letter-spacing: 0.18em;
  color: #152019;
}
.nc-lc-address {
  margin-top: 4px;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: #7a8177;
}
.nc-lc-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 18px;
  font-size: 6px;
  letter-spacing: 0.06em;
  color: #7a8177;
}
.nc-lc-rule {
  height: 1px;
  margin: 12px 0;
  background: repeating-linear-gradient(90deg, #152019 0 3px, transparent 3px 6px);
  opacity: 0.55;
}
.nc-lc-rule-double {
  height: 4px;
  margin: 16px 0 12px;
  background: none;
  border-top: 1px dashed #152019;
  border-bottom: 1px dashed #152019;
  opacity: 0.5;
}
.nc-lc-name {
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(42px, 14vw, 56px);
  font-weight: 600;
  line-height: 0.96;
  letter-spacing: -0.03em;
  color: #152019;
  overflow-wrap: anywhere;
}
.nc-lc-alias {
  margin-top: 5px;
  font-size: 7px;
  letter-spacing: 0.3em;
  color: #8a6a3a;
  text-transform: uppercase;
}
.nc-lc-role {
  margin-top: 4px;
  font-size: 9px;
  letter-spacing: 0.12em;
  color: #4d6455;
}
.nc-lc-items {
  padding: 2px 0;
}
.nc-lc-item {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  padding: 7px 0;
}
.nc-lc-qty {
  width: 18px;
  flex: none;
  font-size: 8px;
  color: #7a8177;
}
.nc-lc-desc {
  flex: 1;
  font-size: 8px;
  line-height: 1.5;
  letter-spacing: 0.08em;
  color: #152019;
}
.nc-lc-desc small {
  font-size: 8.5px;
  font-weight: 700;
  letter-spacing: 0.04em;
  color: #20332a;
  overflow-wrap: anywhere;
}
.nc-lc-price {
  flex: none;
  font-size: 8px;
  color: #7a8177;
}
.nc-lc-total {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 10px;
  padding: 4px 0;
}
.nc-lc-total > span:first-child {
  font-size: 7.5px;
  letter-spacing: 0.12em;
  color: #4d6455;
}
.nc-lc-phone {
  font-size: clamp(17px, 5.2vw, 22px);
  font-weight: 700;
  letter-spacing: 0.04em;
  color: #152019;
  text-decoration: none;
  transition: color 0.25s ease, letter-spacing 0.25s ease;
}
.nc-lc-phone:hover {
  color: #a33a2b;
  letter-spacing: 0.08em;
}
.nc-lc-phone:active {
  opacity: 0.6;
}
.nc-lc-empty {
  color: #a2a89f;
}
.nc-lc-thanks {
  margin-top: 18px;
  text-align: center;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.18em;
  color: #152019;
}
.nc-lc-tip {
  margin-top: 3px;
  text-align: center;
  font-size: 6px;
  letter-spacing: 0.1em;
  color: #7a8177;
}
.nc-lc-barcode {
  height: 36px;
  margin-top: 14px;
  background: repeating-linear-gradient(90deg, #152019 0 1px, transparent 1px 3px, #152019 3px 5px, transparent 5px 6px, #152019 6px 8px, transparent 8px 12px);
}
.nc-lc-code {
  margin-top: 4px;
  text-align: center;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: #7a8177;
}
.nc-lc-after {
  margin-top: 20px;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(221, 235, 219, 0.34);
  animation: nc-lc-in 1s 0.6s both;
}
@keyframes nc-lc-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-lc-print {
  from { opacity: 0; transform: translateY(-18px); clip-path: inset(0 0 100% 0); }
  to { opacity: 1; transform: none; clip-path: inset(0 0 0 0); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-lc-root *,
  .nc-lc-root *::before,
  .nc-lc-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

<template>
  <div class="nc-rr-root relative min-h-screen overflow-hidden px-5 py-8 text-[#1e2a49]">
    <div class="nc-rr-grain"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-rr-head">
        <span>RISO CLUB PRESS</span>
        <span>ONE COLOUR / TWO INKS</span>
      </header>

      <main class="nc-rr-main">
        <div class="nc-rr-sun" aria-hidden="true">
          <span class="nc-rr-sun-dots"></span>
        </div>

        <p class="nc-rr-kicker">A NIGHT FOR PEOPLE WHO STAY</p>

        <div class="nc-rr-title-wrap">
          <h1 class="nc-rr-name nc-rr-name-red">{{ payload?.name }}</h1>
          <h1 class="nc-rr-name nc-rr-name-blue" aria-hidden="true">{{ payload?.name }}</h1>
        </div>

        <div class="nc-rr-subline">
          <span class="nc-rr-alias">{{ payload?.alias || 'NO ALIAS' }}</span>
          <span class="nc-rr-role">{{ payload?.role }}</span>
        </div>

        <div class="nc-rr-poster">
          <div class="nc-rr-poster-head">
            <span>LIVE / RESIDENT</span>
            <span>DOORS 23:00</span>
          </div>

          <div class="nc-rr-venue">
            <p class="nc-rr-venue-name">{{ payload?.venue }}</p>
            <p class="nc-rr-venue-city">{{ payload?.city || '—' }}</p>
          </div>

          <div class="nc-rr-grid">
            <div class="nc-rr-cell">
              <span class="nc-rr-cell-key">WECHAT</span>
              <span class="nc-rr-cell-val">{{ payload?.wechat }}</span>
            </div>
            <div class="nc-rr-cell">
              <span class="nc-rr-cell-key">ENTRY</span>
              <span class="nc-rr-cell-val">BY NAME</span>
            </div>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-rr-cell nc-rr-cell-link">
              <span class="nc-rr-cell-key">PHONE</span>
              <span class="nc-rr-cell-val">{{ payload.phone }}</span>
            </a>
            <div v-else class="nc-rr-cell">
              <span class="nc-rr-cell-key">PHONE</span>
              <span class="nc-rr-cell-val nc-rr-muted"></span>
            </div>
          </div>
        </div>

        <p class="nc-rr-bottom-line">BRING NOTHING · LEAVE NOTHING · TELL NO ONE</p>
      </main>

      <footer class="nc-rr-foot">
        <span>PRINTED IN SHANGHAI</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-rr-root {
  background: #e9e0cb;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-rr-grain {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(rgba(30, 42, 73, 0.2) 0.6px, transparent 0.6px),
    radial-gradient(rgba(226, 59, 46, 0.12) 0.5px, transparent 0.5px);
  background-size: 5px 5px, 7px 7px;
  background-position: 0 0, 2px 3px;
  opacity: 0.5;
  pointer-events: none;
}
.nc-rr-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 10px;
  border-bottom: 2px solid #1e2a49;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: #4b5877;
  animation: nc-rr-in 0.8s both;
}
.nc-rr-main {
  position: relative;
  flex: 1;
  padding-top: 30px;
}
.nc-rr-sun {
  position: absolute;
  top: 6px;
  right: -50px;
  width: 210px;
  height: 210px;
  border-radius: 50%;
  background: #e23b2e;
  overflow: hidden;
  opacity: 0.95;
  animation: nc-rr-sun 1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-rr-sun-dots {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(#e9e0cb 1.2px, transparent 1.3px);
  background-size: 7px 7px;
  mask-image: linear-gradient(180deg, transparent 20%, #000 78%);
  opacity: 0.55;
}
.nc-rr-kicker {
  position: relative;
  z-index: 2;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.26em;
  color: #e23b2e;
  animation: nc-rr-in 0.8s 0.1s both;
}
.nc-rr-title-wrap {
  position: relative;
  z-index: 2;
  margin-top: 20px;
  margin-left: -3px;
}
.nc-rr-name {
  font-size: clamp(58px, 19vw, 82px);
  font-weight: 900;
  line-height: 0.82;
  letter-spacing: -0.065em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-rr-name-red {
  position: relative;
  z-index: 2;
  color: #e23b2e;
  animation: nc-rr-slam 0.8s 0.16s cubic-bezier(0.2, 1.2, 0.4, 1) both;
}
.nc-rr-name-blue {
  position: absolute;
  top: 4px;
  left: 5px;
  z-index: 1;
  color: #1e2a49;
  mix-blend-mode: multiply;
  opacity: 0.85;
  animation: nc-rr-offset 1s 0.4s both;
}
.nc-rr-subline {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 11px;
  margin-top: 20px;
  animation: nc-rr-in 0.9s 0.3s both;
}
.nc-rr-alias {
  padding: 3px 8px;
  border: 1px solid #1e2a49;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.2em;
  color: #1e2a49;
  text-transform: uppercase;
}
.nc-rr-role {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.2em;
  color: #1e2a49;
}
.nc-rr-poster {
  position: relative;
  margin-top: 30px;
  border: 2px solid #1e2a49;
  background: rgba(233, 224, 203, 0.42);
  animation: nc-rr-in 0.95s 0.4s both;
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.nc-rr-poster:hover {
  transform: translate(-2px, -2px);
  box-shadow: 5px 5px 0 #e23b2e;
}
.nc-rr-poster-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 7px 12px;
  border-bottom: 2px solid #1e2a49;
  background: #1e2a49;
  color: #e9e0cb;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
}
.nc-rr-venue {
  padding: 16px 14px 12px;
  border-bottom: 2px solid #1e2a49;
}
.nc-rr-venue-name {
  font-size: clamp(28px, 9vw, 40px);
  font-weight: 900;
  line-height: 0.95;
  letter-spacing: -0.03em;
  color: #1e2a49;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-rr-venue-city {
  margin-top: 5px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.28em;
  color: #e23b2e;
  text-transform: uppercase;
}
.nc-rr-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
.nc-rr-cell {
  display: flex;
  min-width: 0;
  flex-direction: column;
  gap: 4px;
  padding: 12px 14px;
  border-right: 1px solid rgba(30, 42, 73, 0.45);
  border-bottom: 1px solid rgba(30, 42, 73, 0.45);
  text-decoration: none;
  color: inherit;
}
.nc-rr-cell:nth-child(2n) {
  border-right: 0;
}
.nc-rr-cell:nth-last-child(-n + 2) {
  border-bottom: 0;
}
.nc-rr-cell-key {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: #6d7892;
}
.nc-rr-cell-val {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.03em;
  color: #1e2a49;
}
.nc-rr-cell-link {
  transition: background 0.25s ease, color 0.25s ease;
}
.nc-rr-cell-link:hover {
  background: #e23b2e;
}
.nc-rr-cell-link:hover .nc-rr-cell-key,
.nc-rr-cell-link:hover .nc-rr-cell-val {
  color: #f3ebd9;
}
.nc-rr-cell-link:active {
  background: #c22b20;
}
.nc-rr-muted {
  color: #a09a86;
}
.nc-rr-bottom-line {
  margin-top: 20px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #6d7892;
  animation: nc-rr-in 1s 0.55s both;
}
.nc-rr-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 16px;
  border-top: 1px solid rgba(30, 42, 73, 0.35);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #6d7892;
  animation: nc-rr-in 1s 0.65s both;
}
@keyframes nc-rr-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-rr-slam {
  from { opacity: 0; transform: scale(1.2) translateY(-10px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-rr-offset {
  from { opacity: 0; transform: translate(-9px, 9px); }
  to { opacity: 0.85; transform: none; }
}
@keyframes nc-rr-sun {
  from { opacity: 0; transform: scale(0.6) rotate(-20deg); }
  to { opacity: 0.95; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-rr-root *,
  .nc-rr-root *::before,
  .nc-rr-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

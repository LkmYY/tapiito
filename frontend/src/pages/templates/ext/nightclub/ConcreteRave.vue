<template>
  <div class="nc-cr-root relative min-h-screen overflow-hidden px-4 py-6 text-[#171614]">
    <div class="nc-cr-texture"></div>
    <div class="nc-cr-block nc-cr-block-a"></div>
    <div class="nc-cr-block nc-cr-block-b"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3rem)] w-full max-w-md flex-col">
      <header class="nc-cr-head">
        <span>CONCRETE<br />RAVE</span>
        <span class="nc-cr-head-mid">NO. 05 / STRUCTURE</span>
        <span>{{ payload?.city || 'SHANGHAI' }}<br />{{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-cr-main">
        <p class="nc-cr-kicker">BUILT FOR THE FLOOR</p>

        <h1 class="nc-cr-name">{{ payload?.name }}</h1>
        <div class="nc-cr-subrow">
          <span class="nc-cr-alias">{{ payload?.alias || 'NO ALIAS' }}</span>
          <span class="nc-cr-role">{{ payload?.role }}</span>
        </div>

        <section class="nc-cr-grid">
          <div class="nc-cr-tile nc-cr-tile-dark">
            <span class="nc-cr-tile-no">01</span>
            <span class="nc-cr-tile-key">VENUE</span>
            <strong>{{ payload?.venue }}</strong>
            <small>{{ payload?.city || '—' }}</small>
          </div>

          <div class="nc-cr-tile nc-cr-tile-orange">
            <span class="nc-cr-tile-no">02</span>
            <span class="nc-cr-tile-key">WECHAT</span>
            <strong>{{ payload?.wechat }}</strong>
            <small>SCAN / ADD / ASK</small>
          </div>

          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-cr-tile nc-cr-tile-black nc-cr-tile-link">
            <span class="nc-cr-tile-no">03</span>
            <span class="nc-cr-tile-key">PHONE</span>
            <strong>{{ payload.phone }}</strong>
            <small>CALL THE FLOOR</small>
            <span class="nc-cr-tile-arrow">↗</span>
          </a>
          <div v-else class="nc-cr-tile nc-cr-tile-black">
            <span class="nc-cr-tile-no">03</span>
            <span class="nc-cr-tile-key">PHONE</span>
            <strong class="nc-cr-muted"></strong>
            <small>CALL THE FLOOR</small>
          </div>

          <div class="nc-cr-tile nc-cr-tile-outline">
            <span class="nc-cr-tile-no">04</span>
            <span class="nc-cr-tile-key">SET TIME</span>
            <strong>23:00 — 05:00</strong>
            <small>NO SUPPORT ACT</small>
          </div>
        </section>

        <div class="nc-cr-ticker">
          <span>HARDER · LATER · LOUDER · HARDER · LATER · LOUDER ·</span>
        </div>
      </main>

      <footer class="nc-cr-foot">
        <span>RAW SPACE / RAW SOUND</span>
        <span>ENTER THROUGH THE SIDE DOOR</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-cr-root {
  background: #b7b1a5;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-cr-texture {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(rgba(0, 0, 0, 0.16) 0.6px, transparent 0.6px),
    radial-gradient(rgba(255, 255, 255, 0.16) 0.5px, transparent 0.5px),
    linear-gradient(120deg, rgba(255, 255, 255, 0.18), transparent 40%);
  background-size: 3px 3px, 5px 5px, auto;
  opacity: 0.7;
  pointer-events: none;
}
.nc-cr-block {
  position: absolute;
  background: #171614;
  pointer-events: none;
}
.nc-cr-block-a {
  top: 0;
  right: 0;
  width: 26px;
  height: 34%;
}
.nc-cr-block-b {
  bottom: 0;
  left: 0;
  width: 38%;
  height: 16px;
}
.nc-cr-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 10px;
  padding-bottom: 12px;
  border-bottom: 4px solid #171614;
  font-size: 9px;
  font-weight: 900;
  line-height: 1.05;
  letter-spacing: 0.04em;
  text-transform: uppercase;
  animation: nc-cr-in 0.7s steps(4, end) both;
}
.nc-cr-head > span:first-child {
  padding: 3px 5px;
  background: #171614;
  color: #ff4d00;
}
.nc-cr-head-mid {
  flex: 1;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  font-weight: 400;
  letter-spacing: 0.12em;
  text-align: center;
  color: #4c4942;
}
.nc-cr-head > span:last-child {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  font-weight: 400;
  letter-spacing: 0.12em;
  text-align: right;
  color: #4c4942;
}
.nc-cr-main {
  flex: 1;
  padding-top: 30px;
}
.nc-cr-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.3em;
  color: #4c4942;
  animation: nc-cr-in 0.8s 0.06s both;
}
.nc-cr-name {
  margin-top: 6px;
  margin-left: -4px;
  font-size: clamp(58px, 19.5vw, 86px);
  font-weight: 900;
  line-height: 0.78;
  letter-spacing: -0.07em;
  text-transform: uppercase;
  color: #171614;
  overflow-wrap: anywhere;
  animation: nc-cr-slam 0.75s 0.1s cubic-bezier(0.2, 1.2, 0.4, 1) both;
}
.nc-cr-subrow {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 10px;
  margin-top: 18px;
  animation: nc-cr-in 0.8s 0.22s both;
}
.nc-cr-alias {
  padding: 3px 7px;
  background: #ff4d00;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #171614;
  text-transform: uppercase;
}
.nc-cr-role {
  font-size: 9px;
  font-weight: 900;
  letter-spacing: 0.18em;
  color: #171614;
  text-transform: uppercase;
}
.nc-cr-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3px;
  margin-top: 28px;
  border: 3px solid #171614;
  background: #171614;
  animation: nc-cr-in 0.9s 0.3s both;
}
.nc-cr-tile {
  position: relative;
  display: flex;
  min-height: 122px;
  flex-direction: column;
  gap: 5px;
  padding: 12px 11px;
  background: #c8c2b6;
  text-decoration: none;
  color: inherit;
}
.nc-cr-tile-no {
  position: absolute;
  top: 8px;
  right: 9px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  color: rgba(23, 22, 20, 0.4);
}
.nc-cr-tile-key {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  opacity: 0.55;
}
.nc-cr-tile strong {
  margin-top: auto;
  font-size: 15px;
  font-weight: 900;
  line-height: 1.05;
  letter-spacing: -0.02em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-cr-tile small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.1em;
  opacity: 0.55;
}
.nc-cr-tile-dark {
  background: #d3cdc1;
}
.nc-cr-tile-orange {
  background: #ff4d00;
}
.nc-cr-tile-black {
  background: #171614;
  color: #f2ede4;
}
.nc-cr-tile-black .nc-cr-tile-no,
.nc-cr-tile-black .nc-cr-tile-key,
.nc-cr-tile-black .nc-cr-tile small {
  color: rgba(242, 237, 228, 0.5);
  opacity: 1;
}
.nc-cr-tile-outline {
  background: #b7b1a5;
  border: 2px solid #171614;
}
.nc-cr-tile-link {
  transition: background 0.3s ease, color 0.3s ease;
}
.nc-cr-tile-link:hover {
  background: #ff4d00;
  color: #171614;
}
.nc-cr-tile-link:hover .nc-cr-tile-no,
.nc-cr-tile-link:hover .nc-cr-tile-key,
.nc-cr-tile-link:hover small {
  color: rgba(23, 22, 20, 0.6);
}
.nc-cr-tile-link:active {
  background: #e04400;
}
.nc-cr-tile-arrow {
  position: absolute;
  right: 10px;
  bottom: 9px;
  font-size: 14px;
  opacity: 0;
  transform: translate(-4px, 4px);
  transition: opacity 0.25s ease, transform 0.25s ease;
}
.nc-cr-tile-link:hover .nc-cr-tile-arrow {
  opacity: 1;
  transform: none;
}
.nc-cr-muted {
  color: rgba(242, 237, 228, 0.3);
}
.nc-cr-ticker {
  overflow: hidden;
  margin-top: 16px;
  padding: 5px 0;
  border-top: 3px solid #171614;
  border-bottom: 3px solid #171614;
  white-space: nowrap;
  font-size: 8px;
  font-weight: 900;
  letter-spacing: 0.14em;
  color: #171614;
  animation: nc-cr-in 1s 0.45s both;
}
.nc-cr-ticker span {
  display: inline-block;
  animation: nc-cr-scroll 14s linear infinite;
}
.nc-cr-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 14px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #4c4942;
  animation: nc-cr-in 1s 0.56s both;
}
@keyframes nc-cr-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-cr-slam {
  from { opacity: 0; transform: translateY(-18px) scale(1.15); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-cr-scroll {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-cr-root *,
  .nc-cr-root *::before,
  .nc-cr-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

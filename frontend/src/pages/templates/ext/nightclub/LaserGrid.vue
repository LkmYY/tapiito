<template>
  <div class="nc-lg-root relative min-h-screen overflow-hidden px-5 py-8 text-[#f4f1ec]">
    <div class="nc-lg-beams" aria-hidden="true">
      <i v-for="n in 7" :key="n" :style="{ '--bd': n * 0.42 + 's', '--bx': (n - 4) * 26 + 'px' }"></i>
    </div>
    <div class="nc-lg-floor" aria-hidden="true"></div>
    <div class="nc-lg-split" aria-hidden="true"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-lg-head">
        <span class="nc-lg-brand">LASER / GRID</span>
        <span class="nc-lg-code">SYS.09 · {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-lg-main">
        <p class="nc-lg-kicker">TARGET LOCKED · FLOOR 02</p>
        <p class="nc-lg-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-lg-name">{{ payload?.name }}</h1>
        <p class="nc-lg-role">{{ payload?.role }}</p>

        <div class="nc-lg-readout">
          <span class="nc-lg-readout-label">VENUE</span>
          <span class="nc-lg-readout-value">{{ payload?.venue }}</span>
          <span class="nc-lg-readout-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-lg-console">
          <div class="nc-lg-console-head">
            <span>COMMS CHANNEL</span>
            <span class="nc-lg-signal"><i></i><i></i><i></i><i></i></span>
          </div>

          <div class="nc-lg-row">
            <span class="nc-lg-key">WECHAT</span>
            <span class="nc-lg-value">{{ payload?.wechat }}</span>
            <span class="nc-lg-bar"><i style="width: 82%"></i></span>
          </div>

          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-lg-row nc-lg-row-link">
            <span class="nc-lg-key">PHONE</span>
            <span class="nc-lg-value">{{ payload.phone }}</span>
            <span class="nc-lg-call">CONNECT ↗</span>
          </a>
          <div v-else class="nc-lg-row">
            <span class="nc-lg-key">PHONE</span>
            <span class="nc-lg-value nc-lg-empty"></span>
            <span class="nc-lg-bar"><i style="width: 12%"></i></span>
          </div>

          <div class="nc-lg-row nc-lg-row-last">
            <span class="nc-lg-key">STATUS</span>
            <span class="nc-lg-value nc-lg-live">ACTIVE AFTER 23:00</span>
            <span class="nc-lg-bar"><i style="width: 64%"></i></span>
          </div>
        </section>

        <div class="nc-lg-marquee" aria-hidden="true">
          <span>DO NOT CROSS THE BEAM · DO NOT CROSS THE BEAM ·</span>
        </div>
      </main>

      <footer class="nc-lg-foot">
        <span>GRID SECTOR 09</span>
        <span>RED LINE / NO RETURN</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-lg-root {
  background: #060507;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-lg-beams {
  position: absolute;
  inset: -10% -30%;
  overflow: hidden;
  pointer-events: none;
}
.nc-lg-beams i {
  position: absolute;
  top: 38%;
  left: 50%;
  width: 2px;
  height: 150%;
  background: linear-gradient(180deg, transparent, rgba(255, 36, 54, 0.75) 24%, rgba(255, 36, 54, 0.06) 80%, transparent);
  transform-origin: top center;
  transform: rotate(var(--bx));
  filter: blur(0.5px);
  animation: nc-lg-sweep 6s var(--bd) ease-in-out infinite alternate;
}
.nc-lg-beams i:nth-child(odd) {
  background: linear-gradient(180deg, transparent, rgba(255, 90, 60, 0.5) 30%, transparent);
}
.nc-lg-floor {
  position: absolute;
  right: -20%;
  bottom: -10%;
  left: -20%;
  height: 46%;
  background:
    linear-gradient(90deg, rgba(255, 36, 54, 0.22) 1px, transparent 1px),
    linear-gradient(rgba(255, 36, 54, 0.22) 1px, transparent 1px);
  background-size: 42px 42px;
  transform: perspective(320px) rotateX(62deg);
  transform-origin: bottom center;
  mask-image: linear-gradient(0deg, #000 10%, transparent 86%);
  pointer-events: none;
}
.nc-lg-split {
  position: absolute;
  top: 0;
  right: 0;
  width: 38%;
  height: 100%;
  background: linear-gradient(160deg, rgba(255, 36, 54, 0.1), transparent 60%);
  clip-path: polygon(58% 0, 100% 0, 100% 100%, 0 100%);
  pointer-events: none;
}
.nc-lg-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(255, 36, 54, 0.4);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(244, 241, 236, 0.42);
  animation: nc-lg-in 0.8s both;
}
.nc-lg-brand {
  color: #ff2436;
  text-shadow: 0 0 12px rgba(255, 36, 54, 0.7);
}
.nc-lg-main {
  flex: 1;
  padding-top: 54px;
}
.nc-lg-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.28em;
  color: rgba(255, 36, 54, 0.8);
  animation: nc-lg-in 0.8s 0.06s both;
}
.nc-lg-alias {
  margin-top: 24px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.48em;
  color: rgba(244, 241, 236, 0.5);
  text-transform: uppercase;
  animation: nc-lg-in 0.8s 0.12s both;
}
.nc-lg-name {
  position: relative;
  margin-top: 8px;
  font-size: clamp(58px, 19vw, 84px);
  font-weight: 800;
  line-height: 0.84;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  color: #f6f2ec;
  text-shadow: 0 0 30px rgba(255, 36, 54, 0.24);
  overflow-wrap: anywhere;
  animation: nc-lg-name 1s 0.16s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-lg-name::after {
  content: '';
  position: absolute;
  left: 0;
  bottom: -7px;
  width: 42%;
  height: 2px;
  background: #ff2436;
  box-shadow: 0 0 12px rgba(255, 36, 54, 0.9);
  animation: nc-lg-line 1s 0.7s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-lg-role {
  margin-top: 18px;
  font-size: 9.5px;
  font-weight: 600;
  letter-spacing: 0.28em;
  color: rgba(244, 241, 236, 0.62);
  animation: nc-lg-in 0.8s 0.28s both;
}
.nc-lg-readout {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  gap: 9px;
  margin-top: 30px;
  padding: 9px 0;
  border-top: 1px solid rgba(244, 241, 236, 0.16);
  border-bottom: 1px solid rgba(244, 241, 236, 0.16);
  animation: nc-lg-in 0.9s 0.34s both;
}
.nc-lg-readout-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(255, 36, 54, 0.8);
}
.nc-lg-readout-value {
  flex: 1;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: rgba(244, 241, 236, 0.9);
  text-transform: uppercase;
}
.nc-lg-readout-city {
  padding: 1px 5px;
  border: 1px solid rgba(244, 241, 236, 0.24);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(244, 241, 236, 0.5);
}
.nc-lg-console {
  margin-top: 26px;
  border: 1px solid rgba(255, 36, 54, 0.4);
  background: rgba(10, 6, 9, 0.72);
  box-shadow: 0 0 30px -8px rgba(255, 36, 54, 0.3);
  animation: nc-lg-in 0.95s 0.4s both;
}
.nc-lg-console-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 10px;
  border-bottom: 1px solid rgba(255, 36, 54, 0.35);
  background: rgba(255, 36, 54, 0.08);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(244, 241, 236, 0.5);
}
.nc-lg-signal {
  display: inline-flex;
  align-items: flex-end;
  gap: 2px;
  height: 8px;
}
.nc-lg-signal i {
  width: 2px;
  background: #ff2436;
  animation: nc-lg-signal 1.2s ease-in-out infinite alternate;
}
.nc-lg-signal i:nth-child(1) { height: 3px; animation-delay: 0.1s; }
.nc-lg-signal i:nth-child(2) { height: 5px; animation-delay: 0.3s; }
.nc-lg-signal i:nth-child(3) { height: 7px; animation-delay: 0.2s; }
.nc-lg-signal i:nth-child(4) { height: 4px; animation-delay: 0.5s; }
.nc-lg-row {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 10px;
  border-bottom: 1px solid rgba(244, 241, 236, 0.08);
  text-decoration: none;
  color: inherit;
}
.nc-lg-row-last {
  border-bottom: 0;
}
.nc-lg-key {
  width: 46px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(244, 241, 236, 0.34);
}
.nc-lg-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11.5px;
  letter-spacing: 0.04em;
  color: rgba(244, 241, 236, 0.9);
}
.nc-lg-bar {
  display: block;
  width: 38px;
  height: 2px;
  flex: none;
  background: rgba(244, 241, 236, 0.12);
}
.nc-lg-bar i {
  display: block;
  height: 100%;
  background: #ff2436;
  box-shadow: 0 0 6px rgba(255, 36, 54, 0.7);
}
.nc-lg-call {
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.1em;
  color: #ff2436;
}
.nc-lg-row-link {
  transition: background 0.25s ease, padding 0.25s ease;
}
.nc-lg-row-link:hover {
  background: rgba(255, 36, 54, 0.16);
  padding-left: 16px;
}
.nc-lg-row-link:active {
  background: rgba(255, 36, 54, 0.28);
}
.nc-lg-empty {
  color: rgba(244, 241, 236, 0.26);
}
.nc-lg-live {
  color: #ff5a4a;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.12em;
}
.nc-lg-marquee {
  overflow: hidden;
  margin-top: 20px;
  padding: 4px 0;
  border-top: 1px solid rgba(255, 36, 54, 0.3);
  border-bottom: 1px solid rgba(255, 36, 54, 0.3);
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.24em;
  color: rgba(255, 36, 54, 0.55);
  animation: nc-lg-in 1s 0.52s both;
}
.nc-lg-marquee span {
  display: inline-block;
  animation: nc-lg-scroll 12s linear infinite;
}
.nc-lg-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 14px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(244, 241, 236, 0.28);
  animation: nc-lg-in 1s 0.64s both;
}
@keyframes nc-lg-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-lg-name {
  from { opacity: 0; transform: translateX(-26px); filter: blur(4px); }
  to { opacity: 1; transform: none; filter: blur(0); }
}
@keyframes nc-lg-line {
  from { transform: scaleX(0); transform-origin: left; }
  to { transform: scaleX(1); transform-origin: left; }
}
@keyframes nc-lg-sweep {
  from { transform: rotate(calc(var(--bx) - 4deg)); opacity: 0.55; }
  to { transform: rotate(calc(var(--bx) + 4deg)); opacity: 1; }
}
@keyframes nc-lg-signal {
  from { transform: scaleY(0.5); opacity: 0.5; }
  to { transform: scaleY(1); opacity: 1; }
}
@keyframes nc-lg-scroll {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-lg-root *,
  .nc-lg-root *::before,
  .nc-lg-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

<template>
  <div class="nc-be-root relative min-h-screen overflow-hidden px-5 py-8 text-[#111110]">
    <div class="nc-be-hazard"></div>
    <div class="nc-be-grid"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-be-head">
        <span class="nc-be-agency">NIGHT DIVISION / ENTRY CONTROL</span>
        <span class="nc-be-serial">FORM 09—{{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-be-main">
        <div class="nc-be-permit">
          <span class="nc-be-permit-stamp">CITY ENTRY</span>
          <span class="nc-be-permit-code">PERMIT / 24H</span>
        </div>

        <p class="nc-be-kicker">AUTHORIZED AFTER 23:00</p>
        <h1 class="nc-be-name">{{ payload?.name }}</h1>
        <p class="nc-be-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <p class="nc-be-role">{{ payload?.role }}</p>

        <div class="nc-be-barcode" aria-hidden="true"></div>
        <div class="nc-be-barcode-num" aria-hidden="true">8 4 0 2 1 &nbsp; 0 0 1 3 &nbsp; 8 0 0 0</div>

        <section class="nc-be-table">
          <div class="nc-be-tr">
            <span class="nc-be-th">VENUE / 场地</span>
            <span class="nc-be-td">{{ payload?.venue }}</span>
          </div>
          <div class="nc-be-tr">
            <span class="nc-be-th">CITY / 城市</span>
            <span class="nc-be-td">{{ payload?.city || '—' }}</span>
          </div>
          <div class="nc-be-tr">
            <span class="nc-be-th">WECHAT / 微信</span>
            <span class="nc-be-td">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-be-tr nc-be-tr-link">
            <span class="nc-be-th">PHONE / 电话</span>
            <span class="nc-be-td">{{ payload.phone }}</span>
            <span class="nc-be-call">DIAL</span>
          </a>
          <div v-else class="nc-be-tr">
            <span class="nc-be-th">PHONE / 电话</span>
            <span class="nc-be-td nc-be-empty"></span>
          </div>
        </section>

        <div class="nc-be-stamp">
          <span>APPROVED</span>
          <small>23:00 · NO CURFEW</small>
        </div>
      </main>

      <footer class="nc-be-foot">
        <span>KEEP THIS PASS VISIBLE</span>
        <span>VIOLATION = ONE MORE DRINK</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-be-root {
  background: #f0ede4;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-be-hazard {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 10px;
  background: repeating-linear-gradient(135deg, #ffd400 0 12px, #111110 12px 24px);
  pointer-events: none;
}
.nc-be-grid {
  position: absolute;
  inset: 0;
  background:
    linear-gradient(rgba(17, 17, 16, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(17, 17, 16, 0.05) 1px, transparent 1px);
  background-size: 18px 18px;
  mask-image: linear-gradient(180deg, transparent, #000 30%, #000 90%);
  pointer-events: none;
}
.nc-be-root::after {
  content: '';
  position: absolute;
  right: -60px;
  bottom: 10%;
  width: 220px;
  height: 220px;
  background: #ffd400;
  transform: rotate(18deg);
  opacity: 0.22;
  pointer-events: none;
}
.nc-be-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 8px;
  padding-bottom: 10px;
  border-bottom: 3px solid #111110;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: #111110;
  animation: nc-be-in 0.8s both;
}
.nc-be-serial {
  padding: 2px 5px;
  background: #111110;
  color: #ffd400;
}
.nc-be-main {
  flex: 1;
  padding-top: 26px;
}
.nc-be-permit {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 9px;
  background: #ffd400;
  border: 2px solid #111110;
  animation: nc-be-in 0.8s 0.05s both;
}
.nc-be-permit-stamp {
  font-size: 8px;
  font-weight: 900;
  letter-spacing: 0.3em;
}
.nc-be-permit-code {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
}
.nc-be-kicker {
  margin-top: 32px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: #5a5850;
  animation: nc-be-in 0.8s 0.1s both;
}
.nc-be-name {
  margin-top: 6px;
  font-size: clamp(58px, 19vw, 84px);
  font-weight: 900;
  line-height: 0.82;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  color: #111110;
  overflow-wrap: anywhere;
  animation: nc-be-in 0.9s 0.12s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-be-alias {
  display: inline-block;
  margin-top: 12px;
  padding: 2px 6px;
  border: 1px solid #111110;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: #111110;
  text-transform: uppercase;
  animation: nc-be-in 0.8s 0.2s both;
}
.nc-be-role {
  margin-top: 10px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.22em;
  color: #3b3a35;
  text-transform: uppercase;
  animation: nc-be-in 0.8s 0.24s both;
}
.nc-be-barcode {
  height: 58px;
  margin-top: 26px;
  border: 2px solid #111110;
  background:
    repeating-linear-gradient(90deg, #111110 0 2px, transparent 2px 5px, #111110 5px 6px, transparent 6px 9px, #111110 9px 13px, transparent 13px 15px, #111110 15px 17px, transparent 17px 22px);
  animation: nc-be-in 0.9s 0.3s both;
}
.nc-be-barcode-num {
  padding: 3px 2px 0;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.3em;
  color: #5a5850;
  animation: nc-be-in 0.9s 0.34s both;
}
.nc-be-table {
  margin-top: 22px;
  border-top: 2px solid #111110;
  animation: nc-be-in 0.95s 0.36s both;
}
.nc-be-tr {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 9px 4px;
  border-bottom: 1px solid rgba(17, 17, 16, 0.25);
  text-decoration: none;
  color: inherit;
}
.nc-be-th {
  width: 106px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.1em;
  color: #5a5850;
}
.nc-be-td {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.03em;
  color: #111110;
}
.nc-be-tr-link {
  background: rgba(255, 212, 0, 0.5);
  transition: background 0.25s ease, padding 0.25s ease;
}
.nc-be-tr-link:hover {
  background: #ffd400;
  padding-left: 10px;
}
.nc-be-tr-link:active {
  background: #e6bf00;
}
.nc-be-call {
  flex: none;
  padding: 2px 5px;
  background: #111110;
  color: #ffd400;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.1em;
}
.nc-be-empty {
  color: #a09c8f;
}
.nc-be-stamp {
  display: inline-flex;
  flex-direction: column;
  gap: 3px;
  margin-top: 26px;
  padding: 8px 12px;
  border: 3px double #c22b20;
  color: #c22b20;
  transform: rotate(-4deg);
  animation: nc-be-stamp 0.6s 0.6s cubic-bezier(0.2, 1.4, 0.4, 1) both;
  transition: transform 0.3s ease;
}
.nc-be-stamp:hover {
  transform: rotate(2deg) scale(1.03);
}
.nc-be-stamp span {
  font-size: 15px;
  font-weight: 900;
  letter-spacing: 0.2em;
}
.nc-be-stamp small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
}
.nc-be-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 14px;
  border-top: 3px solid #111110;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: #5a5850;
  animation: nc-be-in 1s 0.7s both;
}
@keyframes nc-be-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-be-stamp {
  from { opacity: 0; transform: rotate(-4deg) scale(1.8); }
  to { opacity: 1; transform: rotate(-4deg) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-be-root *,
  .nc-be-root *::before,
  .nc-be-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

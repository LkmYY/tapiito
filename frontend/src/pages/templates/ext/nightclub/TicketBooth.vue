<template>
  <div class="nc-tb-root relative min-h-screen overflow-hidden px-5 py-10 text-[#12213d]">
    <div class="nc-tb-spot"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-5rem)] w-full max-w-md flex-col justify-center">
      <p class="nc-tb-pre">NEON ROOM · DOOR 03 · 23:00</p>

      <article class="nc-tb-ticket">
        <div class="nc-tb-stub">
          <span class="nc-tb-stub-text">ADMIT ONE</span>
          <span class="nc-tb-stub-star">✦</span>
          <span class="nc-tb-stub-no">No. 0417</span>
        </div>

        <div class="nc-tb-body">
          <div class="nc-tb-head">
            <span class="nc-tb-brand">NEON ROOM</span>
            <span class="nc-tb-serial">NFC / {{ nfcId || '0417' }}</span>
          </div>

          <p class="nc-tb-kicker">LATE NIGHT PRESENTS</p>
          <h1 class="nc-tb-name">{{ payload?.name }}</h1>
          <p class="nc-tb-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
          <p class="nc-tb-role">{{ payload?.role }}</p>

          <div class="nc-tb-rule"></div>

          <div class="nc-tb-info">
            <div class="nc-tb-info-row">
              <span class="nc-tb-key">VENUE</span>
              <span class="nc-tb-val">{{ payload?.venue }}</span>
              <span class="nc-tb-city">{{ payload?.city || '—' }}</span>
            </div>
            <div class="nc-tb-info-row">
              <span class="nc-tb-key">WECHAT</span>
              <span class="nc-tb-val">{{ payload?.wechat }}</span>
            </div>
            <div class="nc-tb-info-row">
              <span class="nc-tb-key">PHONE</span>
              <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-tb-phone">{{ payload.phone }}</a>
            </div>
          </div>

          <div class="nc-tb-bottom">
            <div class="nc-tb-barcode" aria-hidden="true"></div>
            <div class="nc-tb-bottom-copy">
              <span>ONE NIGHT ONLY</span>
              <small>NO RE-ENTRY · KEEP THIS STUB</small>
            </div>
          </div>
        </div>
      </article>

      <div class="nc-tb-after">
        <span>DRESS CODE: BLACK</span>
        <span class="nc-tb-after-dot"></span>
        <span>KNOCK TWICE</span>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-tb-root {
  background: #0b1930;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-tb-spot {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 50% 34%, rgba(233, 196, 106, 0.16), transparent 52%),
    repeating-linear-gradient(115deg, transparent 0 42px, rgba(255, 255, 255, 0.018) 42px 43px);
  pointer-events: none;
}
.nc-tb-pre {
  margin-bottom: 18px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.32em;
  color: rgba(240, 230, 205, 0.48);
  animation: nc-tb-fade 0.8s both;
}
.nc-tb-ticket {
  position: relative;
  display: flex;
  overflow: hidden;
  border-radius: 3px;
  background: #f3e8d2;
  box-shadow: 0 30px 70px -26px rgba(0, 0, 0, 0.8), 0 0 0 1px rgba(255, 255, 255, 0.1);
  animation: nc-tb-drop 0.95s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.nc-tb-ticket::before,
.nc-tb-ticket::after {
  content: '';
  position: absolute;
  left: 62px;
  width: 18px;
  height: 18px;
  border-radius: 50%;
  background: #0b1930;
  z-index: 2;
}
.nc-tb-ticket::before { top: -9px; }
.nc-tb-ticket::after { bottom: -9px; }
.nc-tb-ticket:hover {
  transform: translateY(-5px) rotate(-0.5deg);
  box-shadow: 0 40px 80px -26px rgba(0, 0, 0, 0.85), 0 0 0 1px rgba(255, 255, 255, 0.14);
}
.nc-tb-stub {
  position: relative;
  display: flex;
  width: 63px;
  flex: none;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding: 17px 0;
  border-right: 1px dashed rgba(18, 33, 61, 0.4);
  background:
    radial-gradient(rgba(18, 33, 61, 0.16) 0.6px, transparent 0.6px),
    #e5d3b2;
  background-size: 4px 4px, auto;
}
.nc-tb-stub-text {
  writing-mode: vertical-rl;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.28em;
  color: #a33a2b;
}
.nc-tb-stub-star {
  font-size: 16px;
  color: #a33a2b;
  animation: nc-tb-spin 12s linear infinite;
}
.nc-tb-stub-no {
  writing-mode: vertical-rl;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: rgba(18, 33, 61, 0.5);
}
.nc-tb-body {
  flex: 1;
  min-width: 0;
  padding: 20px 20px 18px;
}
.nc-tb-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 9px;
  border-bottom: 3px solid #12213d;
}
.nc-tb-brand {
  font-size: 11px;
  font-weight: 800;
  letter-spacing: 0.24em;
  color: #a33a2b;
}
.nc-tb-serial {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: rgba(18, 33, 61, 0.6);
}
.nc-tb-kicker {
  margin-top: 24px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.3em;
  color: #8b7658;
}
.nc-tb-name {
  margin-top: 6px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(44px, 14vw, 58px);
  font-weight: 600;
  line-height: 0.96;
  letter-spacing: -0.04em;
  color: #12213d;
  overflow-wrap: anywhere;
}
.nc-tb-alias {
  margin-top: 8px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.32em;
  color: #a33a2b;
  text-transform: uppercase;
}
.nc-tb-role {
  margin-top: 5px;
  font-size: 10.5px;
  font-weight: 600;
  letter-spacing: 0.14em;
  color: #4a5a76;
}
.nc-tb-rule {
  height: 1px;
  margin: 20px 0 4px;
  background: rgba(18, 33, 61, 0.25);
}
.nc-tb-info {
  display: flex;
  flex-direction: column;
}
.nc-tb-info-row {
  display: flex;
  align-items: baseline;
  gap: 10px;
  padding: 9px 0;
  border-bottom: 1px dotted rgba(18, 33, 61, 0.25);
}
.nc-tb-key {
  width: 46px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: #8b7658;
}
.nc-tb-val {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11.5px;
  letter-spacing: 0.02em;
  color: #12213d;
}
.nc-tb-city {
  flex: none;
  padding: 1px 5px;
  border: 1px solid rgba(163, 58, 43, 0.55);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.12em;
  color: #a33a2b;
}
.nc-tb-phone {
  min-width: 0;
  flex: 1;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 12.5px;
  font-weight: 700;
  letter-spacing: 0.06em;
  color: #12213d;
  text-decoration: none;
  transition: color 0.25s ease, letter-spacing 0.25s ease;
}
.nc-tb-phone:hover {
  color: #a33a2b;
  letter-spacing: 0.12em;
}
.nc-tb-phone:active {
  opacity: 0.6;
}
.nc-tb-empty {
  color: #a99a83;
}
.nc-tb-bottom {
  display: flex;
  align-items: flex-end;
  gap: 12px;
  margin-top: 16px;
}
.nc-tb-barcode {
  width: 108px;
  height: 25px;
  flex: none;
  background: repeating-linear-gradient(90deg, #12213d 0 1px, transparent 1px 3px, #12213d 3px 5px, transparent 5px 6px, #12213d 6px 8px, transparent 8px 11px);
  opacity: 0.85;
}
.nc-tb-bottom-copy {
  display: flex;
  flex-direction: column;
  gap: 3px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: #8b7658;
}
.nc-tb-bottom-copy span {
  font-size: 7.5px;
  letter-spacing: 0.24em;
  color: #a33a2b;
}
.nc-tb-after {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding-top: 24px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: rgba(240, 230, 205, 0.4);
  animation: nc-tb-fade 1s 0.7s both;
}
.nc-tb-after-dot {
  width: 3px;
  height: 3px;
  border-radius: 50%;
  background: #e9c46a;
}
@keyframes nc-tb-drop {
  from { opacity: 0; transform: translateY(-26px) rotate(1.6deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-tb-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes nc-tb-spin {
  from { transform: rotate(0); }
  to { transform: rotate(360deg); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-tb-root *,
  .nc-tb-root *::before,
  .nc-tb-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

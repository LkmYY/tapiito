<template>
  <div class="nc-bp-root relative min-h-screen overflow-hidden px-4 py-6 text-[#bfe9ff]">
    <div class="nc-bp-grid"></div>
    <div class="nc-bp-corner nc-bp-corner-tl"></div>
    <div class="nc-bp-corner nc-bp-corner-tr"></div>
    <div class="nc-bp-corner nc-bp-corner-bl"></div>
    <div class="nc-bp-corner nc-bp-corner-br"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3rem)] w-full max-w-md flex-col">
      <header class="nc-bp-head">
        <div>
          <p class="nc-bp-doc">DWG NO. NC-303</p>
          <p class="nc-bp-title">NIGHTCLUB / IDENTITY PLAN</p>
        </div>
        <div class="nc-bp-scale">
          <span>SCALE 1:1</span>
          <span>SHEET 01 / 01</span>
        </div>
      </header>

      <main class="nc-bp-main">
        <div class="nc-bp-dimension">
          <span class="nc-bp-dim-line"></span>
          <span class="nc-bp-dim-label">IDENTITY FIELD · 420 MM</span>
        </div>

        <p class="nc-bp-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-bp-name">{{ payload?.name }}</h1>
        <p class="nc-bp-role">{{ payload?.role }}</p>

        <section class="nc-bp-plan">
          <div class="nc-bp-plan-head">
            <span>PLAN VIEW — RESIDENCY</span>
            <span>REV. 03</span>
          </div>
          <div class="nc-bp-plan-body">
            <div class="nc-bp-plan-box">
              <span class="nc-bp-room">{{ payload?.venue }}</span>
              <span class="nc-bp-room-sub">MAIN FLOOR</span>
              <span class="nc-bp-marker nc-bp-marker-a"></span>
              <span class="nc-bp-marker nc-bp-marker-b"></span>
            </div>
            <div class="nc-bp-annotation">
              <span class="nc-bp-arrow-line"></span>
              <span>CITY / {{ payload?.city || '—' }}</span>
            </div>
            <div class="nc-bp-annotation nc-bp-annotation-low">
              <span class="nc-bp-arrow-line"></span>
              <span>CAPACITY / AFTER DARK</span>
            </div>
          </div>
        </section>

        <section class="nc-bp-schedule">
          <div class="nc-bp-schedule-head">
            <span>CONTACT SCHEDULE</span>
            <span>ALL DIMENSIONS IN MM</span>
          </div>
          <div class="nc-bp-schedule-row">
            <span class="nc-bp-code">WX-01</span>
            <span class="nc-bp-desc">WECHAT / 微信</span>
            <span class="nc-bp-data">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-bp-schedule-row nc-bp-schedule-link">
            <span class="nc-bp-code">TEL-02</span>
            <span class="nc-bp-desc">PHONE / 电话</span>
            <span class="nc-bp-data">{{ payload.phone }}</span>
            <span class="nc-bp-call">↗</span>
          </a>
          <div v-else class="nc-bp-schedule-row">
            <span class="nc-bp-code">TEL-02</span>
            <span class="nc-bp-desc">PHONE / 电话</span>
            <span class="nc-bp-data nc-bp-empty"></span>
          </div>
        </section>

        <div class="nc-bp-title-block">
          <div class="nc-bp-tb-cell">
            <span>PROJECT</span>
            <b>NIGHTCLUB IDENTITY</b>
          </div>
          <div class="nc-bp-tb-cell">
            <span>DRAWN BY</span>
            <b>{{ payload?.name }}</b>
          </div>
          <div class="nc-bp-tb-cell">
            <span>DATE</span>
            <b>23:00 — LATE</b>
          </div>
          <div class="nc-bp-tb-cell nc-bp-tb-nfc">
            <span>NFC TAG</span>
            <b>{{ nfcId || '0000' }}</b>
          </div>
        </div>
      </main>

      <footer class="nc-bp-foot">
        <span>APPROVED FOR CONSTRUCTION</span>
        <span>DO NOT SCALE DRAWING</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-bp-root {
  background: #041f33;
  font-family: 'SFMono-Regular', Consolas, 'Courier New', monospace;
}
.nc-bp-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(88, 196, 245, 0.12) 1px, transparent 1px),
    linear-gradient(90deg, rgba(88, 196, 245, 0.12) 1px, transparent 1px),
    linear-gradient(rgba(88, 196, 245, 0.045) 1px, transparent 1px),
    linear-gradient(90deg, rgba(88, 196, 245, 0.045) 1px, transparent 1px);
  background-size: 52px 52px, 52px 52px, 10px 10px, 10px 10px;
  pointer-events: none;
}
.nc-bp-corner {
  position: absolute;
  width: 26px;
  height: 26px;
  border-color: rgba(127, 220, 255, 0.6);
  border-style: solid;
  pointer-events: none;
}
.nc-bp-corner-tl { top: 14px; left: 14px; border-width: 1px 0 0 1px; }
.nc-bp-corner-tr { top: 14px; right: 14px; border-width: 1px 1px 0 0; }
.nc-bp-corner-bl { bottom: 14px; left: 14px; border-width: 0 0 1px 1px; }
.nc-bp-corner-br { bottom: 14px; right: 14px; border-width: 0 1px 1px 0; }
.nc-bp-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
  padding-bottom: 13px;
  border-bottom: 1px solid rgba(127, 220, 255, 0.4);
  animation: nc-bp-in 0.8s both;
}
.nc-bp-doc {
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(191, 233, 255, 0.5);
}
.nc-bp-title {
  margin-top: 4px;
  font-size: 8px;
  letter-spacing: 0.14em;
  color: #d9f4ff;
}
.nc-bp-scale {
  display: flex;
  flex-direction: column;
  gap: 4px;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(191, 233, 255, 0.42);
  text-align: right;
}
.nc-bp-main {
  flex: 1;
  padding-top: 26px;
}
.nc-bp-dimension {
  display: flex;
  align-items: center;
  gap: 8px;
  animation: nc-bp-fade 1s 0.1s both;
}
.nc-bp-dim-line {
  position: relative;
  flex: 1;
  height: 1px;
  background: rgba(127, 220, 255, 0.5);
}
.nc-bp-dim-line::before,
.nc-bp-dim-line::after {
  content: '';
  position: absolute;
  top: -4px;
  width: 1px;
  height: 9px;
  background: rgba(127, 220, 255, 0.7);
}
.nc-bp-dim-line::before { left: 0; }
.nc-bp-dim-line::after { right: 0; }
.nc-bp-dim-label {
  flex: none;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(191, 233, 255, 0.45);
}
.nc-bp-alias {
  margin-top: 34px;
  font-size: 8px;
  letter-spacing: 0.42em;
  color: #7fdcff;
  text-transform: uppercase;
  animation: nc-bp-in 0.9s 0.16s both;
}
.nc-bp-name {
  margin-top: 6px;
  font-size: clamp(50px, 16vw, 68px);
  font-weight: 400;
  line-height: 0.94;
  letter-spacing: -0.03em;
  color: #e1f6ff;
  text-shadow: 0 0 26px rgba(88, 196, 245, 0.35);
  overflow-wrap: anywhere;
  animation: nc-bp-in 0.95s 0.2s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-bp-role {
  margin-top: 10px;
  font-size: 9px;
  letter-spacing: 0.26em;
  color: rgba(191, 233, 255, 0.62);
  animation: nc-bp-in 0.9s 0.28s both;
}
.nc-bp-plan {
  margin-top: 28px;
  border: 1px solid rgba(127, 220, 255, 0.42);
  animation: nc-bp-in 1s 0.34s both;
}
.nc-bp-plan-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 9px;
  border-bottom: 1px solid rgba(127, 220, 255, 0.3);
  background: rgba(88, 196, 245, 0.09);
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(191, 233, 255, 0.6);
}
.nc-bp-plan-body {
  position: relative;
  padding: 16px;
}
.nc-bp-plan-box {
  position: relative;
  min-height: 104px;
  padding: 13px;
  border: 1px dashed rgba(127, 220, 255, 0.55);
  background:
    linear-gradient(90deg, rgba(88, 196, 245, 0.08) 1px, transparent 1px),
    linear-gradient(rgba(88, 196, 245, 0.08) 1px, transparent 1px);
  background-size: 16px 16px;
}
.nc-bp-room {
  display: block;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
  font-size: 19px;
  font-weight: 700;
  letter-spacing: 0.04em;
  color: #e1f6ff;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.nc-bp-room-sub {
  display: block;
  margin-top: 3px;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(191, 233, 255, 0.45);
}
.nc-bp-marker {
  position: absolute;
  width: 7px;
  height: 7px;
  border: 1px solid #7fdcff;
  border-radius: 50%;
  background: #041f33;
}
.nc-bp-marker::after {
  content: '';
  position: absolute;
  inset: 2px;
  border-radius: 50%;
  background: #7fdcff;
}
.nc-bp-marker-a { top: 14px; right: 14px; }
.nc-bp-marker-b { right: 34px; bottom: 14px; }
.nc-bp-annotation {
  position: absolute;
  top: 24px;
  right: -1px;
  display: flex;
  align-items: center;
  gap: 5px;
  font-size: 5.5px;
  letter-spacing: 0.1em;
  color: rgba(191, 233, 255, 0.5);
  transform: translateX(100%);
  display: none;
}
.nc-bp-annotation-low {
  display: none;
}
.nc-bp-schedule {
  margin-top: 15px;
  border: 1px solid rgba(127, 220, 255, 0.42);
  animation: nc-bp-in 1s 0.42s both;
}
.nc-bp-schedule-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 9px;
  border-bottom: 1px solid rgba(127, 220, 255, 0.3);
  background: rgba(88, 196, 245, 0.09);
  font-size: 6px;
  letter-spacing: 0.14em;
  color: rgba(191, 233, 255, 0.6);
}
.nc-bp-schedule-row {
  display: flex;
  align-items: center;
  gap: 9px;
  padding: 9px;
  border-bottom: 1px solid rgba(127, 220, 255, 0.14);
  text-decoration: none;
  color: inherit;
}
.nc-bp-schedule-row:last-child {
  border-bottom: 0;
}
.nc-bp-code {
  width: 42px;
  flex: none;
  font-size: 6px;
  letter-spacing: 0.08em;
  color: #7fdcff;
}
.nc-bp-desc {
  flex: 1;
  font-size: 6.5px;
  letter-spacing: 0.12em;
  color: rgba(191, 233, 255, 0.44);
}
.nc-bp-data {
  min-width: 0;
  max-width: 52%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11px;
  letter-spacing: 0.03em;
  color: #d9f4ff;
}
.nc-bp-schedule-link {
  margin: 0 -3px;
  padding: 9px 3px;
  transition: background 0.25s ease;
}
.nc-bp-schedule-link:hover {
  background: rgba(127, 220, 255, 0.12);
}
.nc-bp-schedule-link:active {
  background: rgba(127, 220, 255, 0.22);
}
.nc-bp-call {
  color: #7fdcff;
  font-size: 10px;
}
.nc-bp-empty {
  color: rgba(191, 233, 255, 0.28);
}
.nc-bp-title-block {
  display: grid;
  grid-template-columns: 1fr 1fr;
  margin-top: 15px;
  border: 1px solid rgba(127, 220, 255, 0.5);
  animation: nc-bp-in 1s 0.5s both;
}
.nc-bp-tb-cell {
  display: flex;
  flex-direction: column;
  gap: 3px;
  padding: 6px 8px;
  border-right: 1px solid rgba(127, 220, 255, 0.2);
  border-bottom: 1px solid rgba(127, 220, 255, 0.2);
}
.nc-bp-tb-cell:nth-child(2n) {
  border-right: 0;
}
.nc-bp-tb-cell:nth-child(n + 3) {
  border-bottom: 0;
}
.nc-bp-tb-cell span {
  font-size: 5px;
  letter-spacing: 0.14em;
  color: rgba(191, 233, 255, 0.4);
}
.nc-bp-tb-cell b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 7.5px;
  font-weight: 400;
  letter-spacing: 0.06em;
  color: #d9f4ff;
}
.nc-bp-tb-nfc b {
  color: #7fdcff;
}
.nc-bp-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 15px;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: rgba(191, 233, 255, 0.34);
  animation: nc-bp-fade 1s 0.65s both;
}
@keyframes nc-bp-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-bp-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-bp-root *,
  .nc-bp-root *::before,
  .nc-bp-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

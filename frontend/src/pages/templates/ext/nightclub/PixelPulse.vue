<template>
  <div class="nc-pp-root relative min-h-screen overflow-hidden px-4 py-6 text-[#f7ff3c]">
    <div class="nc-pp-grid"></div>
    <div class="nc-pp-scan"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3rem)] w-full max-w-md flex-col">
      <header class="nc-pp-hud">
        <span>PLAYER 1</span>
        <span class="nc-pp-credit">CREDIT 01</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-pp-main">
        <div class="nc-pp-lives" aria-hidden="true">
          <i></i><i></i><i></i>
        </div>

        <p class="nc-pp-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-pp-name">{{ payload?.name }}</h1>
        <p class="nc-pp-role">{{ payload?.role }}</p>

        <div class="nc-pp-stage">
          <span class="nc-pp-stage-label">STAGE</span>
          <span class="nc-pp-stage-name">{{ payload?.venue }}</span>
          <span class="nc-pp-stage-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-pp-menu">
          <div class="nc-pp-menu-head">
            <span>SELECT CONTACT</span>
            <span>▲▼</span>
          </div>
          <div class="nc-pp-item">
            <span class="nc-pp-cursor">▶</span>
            <span class="nc-pp-item-key">WECHAT</span>
            <span class="nc-pp-item-val">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-pp-item nc-pp-item-link">
            <span class="nc-pp-cursor">▶</span>
            <span class="nc-pp-item-key">CALL</span>
            <span class="nc-pp-item-val">{{ payload.phone }}</span>
            <span class="nc-pp-item-arrow">+</span>
          </a>
          <div v-else class="nc-pp-item">
            <span class="nc-pp-cursor">▶</span>
            <span class="nc-pp-item-key">CALL</span>
            <span class="nc-pp-item-val nc-pp-empty"></span>
          </div>
        </section>

        <p class="nc-pp-press"><span>PRESS START</span><i></i></p>
      </main>

      <footer class="nc-pp-foot">
        <span>HI-SCORE 999999</span>
        <span>CONTINUE? 09</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-pp-root {
  background: #05050c;
  font-family: 'Courier New', 'SFMono-Regular', monospace;
}
.nc-pp-grid {
  position: absolute;
  inset: 0;
  background:
    linear-gradient(rgba(247, 255, 60, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(247, 255, 60, 0.05) 1px, transparent 1px);
  background-size: 20px 20px;
  mask-image: linear-gradient(180deg, #000, transparent 80%);
  pointer-events: none;
}
.nc-pp-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 3px, rgba(0, 0, 0, 0.35) 3px 4px);
  pointer-events: none;
}
.nc-pp-hud {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 6px;
  padding: 7px 9px;
  border: 2px solid #f7ff3c;
  box-shadow: 3px 3px 0 #ff2e88;
  font-size: 7px;
  letter-spacing: 0.1em;
  color: #f7ff3c;
  animation: nc-pp-hud 0.6s steps(5, end) both;
}
.nc-pp-credit {
  padding: 2px 5px;
  background: #f7ff3c;
  color: #05050c;
}
.nc-pp-main {
  flex: 1;
  padding-top: 50px;
}
.nc-pp-lives {
  display: flex;
  gap: 6px;
  margin-bottom: 28px;
}
.nc-pp-lives i {
  position: relative;
  width: 13px;
  height: 11px;
  background: #ff2e88;
  clip-path: polygon(0 20%, 20% 0, 50% 25%, 80% 0, 100% 20%, 100% 60%, 50% 100%, 0 60%);
  animation: nc-pp-blink 1.6s steps(2, end) infinite;
}
.nc-pp-lives i:nth-child(2) { animation-delay: 0.3s; }
.nc-pp-lives i:nth-child(3) { animation-delay: 0.6s; opacity: 0.35; }
.nc-pp-alias {
  font-size: 8px;
  letter-spacing: 0.3em;
  color: #5ff6ff;
  animation: nc-pp-in 0.7s 0.08s both;
}
.nc-pp-name {
  margin-top: 8px;
  font-size: clamp(58px, 19vw, 82px);
  font-weight: 900;
  line-height: 0.85;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  color: #f7ff3c;
  text-shadow:
    3px 0 0 #ff2e88,
    -3px 0 0 #5ff6ff,
    0 4px 0 rgba(0, 0, 0, 0.7);
  overflow-wrap: anywhere;
  animation: nc-pp-name 0.65s 0.12s steps(4, end) both;
}
.nc-pp-role {
  margin-top: 14px;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.24em;
  color: #f7ff3c;
  animation: nc-pp-in 0.7s 0.24s both;
}
.nc-pp-stage {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 7px;
  margin-top: 26px;
  padding: 9px 10px;
  border: 1px solid rgba(95, 246, 255, 0.5);
  background: rgba(95, 246, 255, 0.06);
  animation: nc-pp-in 0.7s 0.3s both;
}
.nc-pp-stage-label {
  padding: 1px 4px;
  background: #5ff6ff;
  color: #05050c;
  font-size: 6px;
  font-weight: 700;
  letter-spacing: 0.08em;
}
.nc-pp-stage-name {
  flex: 1;
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.06em;
  color: #d9ffff;
  text-transform: uppercase;
}
.nc-pp-stage-city {
  font-size: 6.5px;
  letter-spacing: 0.1em;
  color: rgba(95, 246, 255, 0.6);
}
.nc-pp-menu {
  margin-top: 30px;
  border: 2px solid #f7ff3c;
  box-shadow: 5px 5px 0 rgba(255, 46, 136, 0.65);
  animation: nc-pp-in 0.8s 0.38s both;
  transition: box-shadow 0.25s ease, transform 0.25s ease;
}
.nc-pp-menu:hover {
  box-shadow: 5px 5px 0 #5ff6ff;
  transform: translate(-1px, -1px);
}
.nc-pp-menu-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 9px;
  border-bottom: 2px solid #f7ff3c;
  background: #f7ff3c;
  color: #05050c;
  font-size: 6.5px;
  font-weight: 700;
  letter-spacing: 0.1em;
}
.nc-pp-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 11px 9px;
  border-bottom: 1px dotted rgba(247, 255, 60, 0.28);
  text-decoration: none;
  color: inherit;
}
.nc-pp-item:last-child {
  border-bottom: 0;
}
.nc-pp-cursor {
  flex: none;
  font-size: 8px;
  color: #ff2e88;
  animation: nc-pp-cursor 1s steps(2, end) infinite;
}
.nc-pp-item-key {
  width: 50px;
  flex: none;
  font-size: 6.5px;
  letter-spacing: 0.08em;
  color: rgba(247, 255, 60, 0.55);
}
.nc-pp-item-val {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11.5px;
  font-weight: 700;
  letter-spacing: 0.03em;
  color: #f4ffd0;
}
.nc-pp-item-arrow {
  flex: none;
  font-size: 14px;
  color: #5ff6ff;
}
.nc-pp-item-link {
  transition: background 0.2s steps(2, end), padding 0.2s steps(2, end);
}
.nc-pp-item-link:hover {
  background: rgba(247, 255, 60, 0.14);
  padding-left: 15px;
}
.nc-pp-item-link:hover .nc-pp-item-key {
  color: #ff2e88;
}
.nc-pp-item-link:active {
  background: rgba(247, 255, 60, 0.3);
}
.nc-pp-empty {
  color: rgba(247, 255, 60, 0.3);
}
.nc-pp-press {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 9px;
  margin-top: 30px;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.3em;
  color: #f7ff3c;
  animation: nc-pp-blink 1.4s steps(2, end) infinite;
}
.nc-pp-press i {
  width: 7px;
  height: 7px;
  background: #ff2e88;
  clip-path: polygon(50% 0, 100% 50%, 50% 100%, 0 50%);
}
.nc-pp-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 12px;
  border-top: 2px solid rgba(247, 255, 60, 0.5);
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(247, 255, 60, 0.45);
  animation: nc-pp-in 1s 0.55s both;
}
@keyframes nc-pp-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-pp-name {
  from { opacity: 0; transform: translateY(-14px) scaleY(1.4); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-pp-hud {
  from { opacity: 0; clip-path: inset(0 100% 0 0); }
  to { opacity: 1; clip-path: inset(0 0 0 0); }
}
@keyframes nc-pp-blink {
  0%, 55% { opacity: 1; }
  56%, 100% { opacity: 0.25; }
}
@keyframes nc-pp-cursor {
  0%, 60% { opacity: 1; }
  61%, 100% { opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-pp-root *,
  .nc-pp-root *::before,
  .nc-pp-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

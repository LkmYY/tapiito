<template>
  <div class="nc-vm-root relative min-h-screen overflow-hidden px-5 py-9 text-[#efe3cc]">
    <div class="nc-vm-glow"></div>
    <div class="nc-vm-lines"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4.5rem)] w-full max-w-md flex-col">
      <header class="nc-vm-head">
        <span class="nc-vm-monogram">V</span>
        <span class="nc-vm-society">THE VELVET SOCIETY</span>
        <span class="nc-vm-est">EST. 1998</span>
      </header>

      <main class="nc-vm-main">
        <p class="nc-vm-eyebrow">PRIVATE MEMBER · RESIDENT</p>

        <section class="nc-vm-card">
          <div class="nc-vm-card-shine"></div>
          <div class="nc-vm-card-top">
            <span class="nc-vm-chip" aria-hidden="true"></span>
            <span class="nc-vm-tier">BLACK TIER</span>
          </div>

          <div class="nc-vm-identity">
            <p class="nc-vm-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
            <h1 class="nc-vm-name">{{ payload?.name }}</h1>
            <p class="nc-vm-role">{{ payload?.role }}</p>
          </div>

          <div class="nc-vm-card-mid">
            <span class="nc-vm-mid-label">RESIDENCY</span>
            <span class="nc-vm-mid-value">{{ payload?.venue }}</span>
            <span class="nc-vm-mid-city">{{ payload?.city || '—' }}</span>
          </div>

          <div class="nc-vm-card-bottom">
            <div class="nc-vm-field">
              <span>WECHAT</span>
              <b>{{ payload?.wechat }}</b>
            </div>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-vm-field nc-vm-field-link">
              <span>PRIVATE LINE</span>
              <b>{{ payload.phone }}</b>
            </a>
            <div v-else class="nc-vm-field">
              <span>PRIVATE LINE</span>
              <b class="nc-vm-empty"></b>
            </div>
          </div>

          <div class="nc-vm-card-foot">
            <span>MEMBER NO. {{ memberNo }}</span>
            <span class="nc-vm-script">velvet</span>
          </div>
        </section>

        <p class="nc-vm-note">This card is non-transferable. The night, however, is yours.</p>
      </main>

      <footer class="nc-vm-foot">
        <span>AFTER 23:00 · MEMBERS ONLY</span>
        <span>{{ payload?.city || 'SHANGHAI' }} / {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const memberNo = computed(() => {
  const seed = (props.payload?.name || 'NOVA').toString()
  let total = 0
  for (let i = 0; i < seed.length; i += 1) total += seed.charCodeAt(i)
  return String(1000 + (total % 8999)).padStart(4, '0')
})
</script>

<style scoped>
.nc-vm-root {
  background: #0a0806;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-vm-glow {
  position: absolute;
  top: -120px;
  left: 50%;
  width: 420px;
  height: 420px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(212, 175, 106, 0.2), transparent 65%);
  transform: translateX(-50%);
  filter: blur(30px);
  pointer-events: none;
  animation: nc-vm-breathe 9s ease-in-out infinite;
}
.nc-vm-lines {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(90deg, transparent 0 58px, rgba(212, 175, 106, 0.035) 58px 59px);
  pointer-events: none;
}
.nc-vm-head {
  display: flex;
  align-items: center;
  gap: 10px;
  padding-bottom: 14px;
  border-bottom: 1px solid rgba(212, 175, 106, 0.24);
  animation: nc-vm-rise 0.8s both;
}
.nc-vm-monogram {
  display: flex;
  width: 26px;
  height: 26px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(212, 175, 106, 0.65);
  font-family: Georgia, 'Songti SC', serif;
  font-size: 14px;
  color: #d4af6a;
}
.nc-vm-society {
  flex: 1;
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.3em;
  color: rgba(239, 227, 204, 0.82);
}
.nc-vm-est {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(212, 175, 106, 0.55);
}
.nc-vm-main {
  display: flex;
  flex: 1;
  flex-direction: column;
  justify-content: center;
  padding: 44px 0 32px;
}
.nc-vm-eyebrow {
  margin-bottom: 13px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.3em;
  color: rgba(212, 175, 106, 0.7);
  animation: nc-vm-rise 0.8s 0.08s both;
}
.nc-vm-card {
  position: relative;
  overflow: hidden;
  border: 1px solid rgba(212, 175, 106, 0.55);
  border-radius: 8px;
  background:
    repeating-radial-gradient(ellipse at 14% 0%, transparent 0 14px, rgba(212, 175, 106, 0.045) 14px 15px),
    linear-gradient(148deg, #1a1610, #0d0b08 46%, #161109);
  box-shadow:
    0 0 0 4px #0a0806,
    0 0 0 5px rgba(212, 175, 106, 0.16),
    0 36px 70px -30px rgba(0, 0, 0, 0.95),
    inset 0 1px 0 rgba(255, 231, 186, 0.16);
  padding: 22px;
  animation: nc-vm-card 1s 0.12s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.45s ease, box-shadow 0.45s ease;
}
.nc-vm-card:hover {
  transform: translateY(-5px) rotateX(1.5deg);
  box-shadow:
    0 0 0 4px #0a0806,
    0 0 0 5px rgba(212, 175, 106, 0.3),
    0 46px 84px -30px rgba(0, 0, 0, 0.95),
    inset 0 1px 0 rgba(255, 231, 186, 0.24);
}
.nc-vm-card-shine {
  position: absolute;
  top: -40%;
  left: -70%;
  width: 45%;
  height: 180%;
  background: linear-gradient(100deg, transparent, rgba(255, 238, 200, 0.1), transparent);
  transform: rotate(12deg);
  animation: nc-vm-sweep 7s 1.4s ease-in-out infinite;
  pointer-events: none;
}
.nc-vm-card-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.nc-vm-chip {
  position: relative;
  width: 33px;
  height: 25px;
  border: 1px solid rgba(212, 175, 106, 0.6);
  border-radius: 4px;
  background: linear-gradient(135deg, rgba(212, 175, 106, 0.48), rgba(212, 175, 106, 0.08));
}
.nc-vm-chip::before,
.nc-vm-chip::after {
  content: '';
  position: absolute;
  background: rgba(10, 8, 6, 0.55);
}
.nc-vm-chip::before {
  top: 7px;
  left: 0;
  width: 100%;
  height: 1px;
}
.nc-vm-chip::after {
  top: 16px;
  left: 0;
  width: 100%;
  height: 1px;
}
.nc-vm-tier {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: rgba(212, 175, 106, 0.8);
}
.nc-vm-identity {
  margin-top: 30px;
}
.nc-vm-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.4em;
  color: #a88d5f;
  text-transform: uppercase;
}
.nc-vm-name {
  margin-top: 5px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(42px, 13vw, 56px);
  font-weight: 400;
  line-height: 0.98;
  letter-spacing: -0.03em;
  background: linear-gradient(178deg, #fff3d8, #d4af6a 46%, #9a7739);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  overflow-wrap: anywhere;
}
.nc-vm-role {
  margin-top: 7px;
  font-size: 9px;
  letter-spacing: 0.22em;
  color: rgba(239, 227, 204, 0.6);
}
.nc-vm-card-mid {
  display: flex;
  align-items: baseline;
  gap: 9px;
  margin-top: 26px;
  padding: 10px 0;
  border-top: 1px solid rgba(212, 175, 106, 0.18);
  border-bottom: 1px solid rgba(212, 175, 106, 0.18);
}
.nc-vm-mid-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(212, 175, 106, 0.6);
}
.nc-vm-mid-value {
  flex: 1;
  font-size: 10.5px;
  font-weight: 600;
  letter-spacing: 0.16em;
  color: #ecdcbc;
  text-transform: uppercase;
}
.nc-vm-mid-city {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: rgba(239, 227, 204, 0.4);
}
.nc-vm-card-bottom {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
  margin-top: 18px;
}
.nc-vm-field {
  display: flex;
  min-width: 0;
  flex-direction: column;
  gap: 4px;
  padding: 9px 10px;
  border: 1px solid rgba(212, 175, 106, 0.16);
  background: rgba(212, 175, 106, 0.035);
  text-decoration: none;
  color: inherit;
}
.nc-vm-field span {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(212, 175, 106, 0.55);
}
.nc-vm-field b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 10.5px;
  font-weight: 500;
  letter-spacing: 0.04em;
  color: #e7d7b8;
}
.nc-vm-field-link {
  transition: background 0.3s ease, border-color 0.3s ease, transform 0.3s ease;
}
.nc-vm-field-link:hover {
  background: rgba(212, 175, 106, 0.13);
  border-color: rgba(212, 175, 106, 0.55);
  transform: translateY(-2px);
}
.nc-vm-field-link:active {
  transform: translateY(0) scale(0.99);
}
.nc-vm-field-link:hover b {
  color: #ffe9bd;
}
.nc-vm-empty {
  color: rgba(239, 227, 204, 0.3) !important;
}
.nc-vm-card-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 20px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(239, 227, 204, 0.28);
}
.nc-vm-script {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 15px;
  font-style: italic;
  letter-spacing: 0.02em;
  color: rgba(212, 175, 106, 0.62);
}
.nc-vm-note {
  margin-top: 22px;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 10.5px;
  font-style: italic;
  line-height: 1.6;
  letter-spacing: 0.04em;
  color: rgba(239, 227, 204, 0.34);
  animation: nc-vm-rise 1s 0.5s both;
}
.nc-vm-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 16px;
  border-top: 1px solid rgba(212, 175, 106, 0.14);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(239, 227, 204, 0.28);
  animation: nc-vm-rise 1s 0.65s both;
}
@keyframes nc-vm-rise {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-vm-card {
  from { opacity: 0; transform: translateY(26px) rotate(-1.2deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-vm-breathe {
  0%, 100% { opacity: 0.7; transform: translateX(-50%) scale(1); }
  50% { opacity: 1; transform: translateX(-50%) scale(1.08); }
}
@keyframes nc-vm-sweep {
  0%, 55% { left: -70%; }
  85%, 100% { left: 130%; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-vm-root *,
  .nc-vm-root *::before,
  .nc-vm-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

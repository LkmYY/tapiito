<template>
  <div class="nc-bp-root relative min-h-screen overflow-hidden px-5 py-8 text-[#f7e8d8]">
    <div class="nc-bp-brick"></div>
    <div class="nc-bp-lamp nc-bp-lamp-l"></div>
    <div class="nc-bp-lamp nc-bp-lamp-r"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-bp-head">
        <span>BACKSTAGE / SERVICE CORRIDOR</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-bp-main">
        <div class="nc-bp-lanyard">
          <span class="nc-bp-strap"></span>
          <span class="nc-bp-clip"></span>
        </div>

        <article class="nc-bp-pass">
          <div class="nc-bp-pass-band">
            <span>ALL ACCESS</span>
            <span class="nc-bp-pass-year">2026</span>
          </div>

          <div class="nc-bp-pass-body">
            <div class="nc-bp-photo">
              <span class="nc-bp-photo-initial">{{ initial }}</span>
              <span class="nc-bp-photo-lines"></span>
              <small>ARTIST</small>
            </div>

            <div class="nc-bp-info">
              <p class="nc-bp-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
              <h1 class="nc-bp-name">{{ payload?.name }}</h1>
              <p class="nc-bp-role">{{ payload?.role }}</p>

              <div class="nc-bp-venue">
                <span>{{ payload?.venue }}</span>
                <span class="nc-bp-city">{{ payload?.city || '—' }}</span>
              </div>
            </div>
          </div>

          <div class="nc-bp-pass-foot">
            <div class="nc-bp-field">
              <span>WECHAT</span>
              <b>{{ payload?.wechat }}</b>
            </div>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-bp-field nc-bp-field-link">
              <span>PHONE</span>
              <b>{{ payload.phone }}</b>
              <Phone :size="12" />
            </a>
            <div v-else class="nc-bp-field">
              <span>PHONE</span>
              <b class="nc-bp-empty"></b>
            </div>
            <div class="nc-bp-field nc-bp-field-code">
              <span>PASS NO.</span>
              <b>{{ passNo }}</b>
            </div>
          </div>

          <div class="nc-bp-pass-stripes" aria-hidden="true"></div>
        </article>

        <p class="nc-bp-note">DO NOT LOSE. DO NOT LEND. DO NOT ASK AGAIN.</p>
      </main>

      <footer class="nc-bp-foot">
        <span>STAGE DOOR · 23:00</span>
        <span>VALID ONE NIGHT ONLY</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
const passNo = computed(() => {
  const raw = (props.payload?.name || 'NOVA').toString()
  let total = 0
  for (let i = 0; i < raw.length; i += 1) total += raw.charCodeAt(i)
  return 'A-' + String(100 + (total % 899))
})
</script>

<style scoped>
.nc-bp-root {
  background: #24070f;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-bp-brick {
  position: absolute;
  inset: 0;
  background:
    repeating-linear-gradient(0deg, transparent 0 28px, rgba(0, 0, 0, 0.3) 28px 30px),
    repeating-linear-gradient(90deg, transparent 0 56px, rgba(0, 0, 0, 0.3) 56px 58px),
    radial-gradient(ellipse at 50% 12%, rgba(255, 111, 97, 0.2), transparent 46%);
  opacity: 0.65;
  pointer-events: none;
}
.nc-bp-lamp {
  position: absolute;
  top: 0;
  width: 2px;
  height: 100px;
  background: rgba(255, 196, 120, 0.3);
  pointer-events: none;
}
.nc-bp-lamp::after {
  content: '';
  position: absolute;
  bottom: -7px;
  left: -6px;
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: #ffc478;
  box-shadow: 0 0 24px 10px rgba(255, 196, 120, 0.35);
  animation: nc-bp-glow 3.4s ease-in-out infinite;
}
.nc-bp-lamp-l { left: 16px; }
.nc-bp-lamp-r { right: 16px; }
.nc-bp-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(247, 232, 216, 0.18);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(247, 232, 216, 0.42);
  animation: nc-bp-in 0.8s both;
}
.nc-bp-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 22px 0;
}
.nc-bp-lanyard {
  position: relative;
  z-index: 2;
  display: flex;
  height: 92px;
  flex-direction: column;
  align-items: center;
  animation: nc-bp-drop 1.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-bp-strap {
  width: 66px;
  height: 52px;
  border: 9px solid #131013;
  border-bottom: 0;
  border-radius: 40px 40px 0 0;
  background: transparent;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.08);
}
.nc-bp-clip {
  width: 18px;
  height: 24px;
  margin-top: -2px;
  border: 2px solid #9a938c;
  border-radius: 3px;
  background: linear-gradient(180deg, #d8d2c8, #7c756d);
  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.5);
}
.nc-bp-pass {
  position: relative;
  overflow: hidden;
  border-radius: 5px;
  background: #f4e7d3;
  box-shadow: 0 34px 70px -26px rgba(0, 0, 0, 0.9), 0 0 0 1px rgba(255, 255, 255, 0.12);
  color: #211816;
  animation: nc-bp-pass 1s 0.16s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.nc-bp-pass:hover {
  transform: rotate(-0.8deg) translateY(-4px);
  box-shadow: 0 44px 84px -26px rgba(0, 0, 0, 0.92), 0 0 0 1px rgba(255, 255, 255, 0.18);
}
.nc-bp-pass-band {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 9px 14px;
  background: #d23a2e;
  color: #fff6e8;
  font-size: 11px;
  font-weight: 900;
  letter-spacing: 0.26em;
}
.nc-bp-pass-year {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  font-weight: 400;
  letter-spacing: 0.18em;
  opacity: 0.8;
}
.nc-bp-pass-body {
  display: flex;
  gap: 14px;
  padding: 15px 14px 13px;
}
.nc-bp-photo {
  position: relative;
  display: flex;
  width: 96px;
  height: 124px;
  flex: none;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  background:
    radial-gradient(circle at 50% 38%, rgba(210, 58, 46, 0.6), transparent 52%),
    linear-gradient(160deg, #2c2320, #100d0c);
  color: #f4e7d3;
}
.nc-bp-photo::after {
  content: '';
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(0deg, transparent 0 3px, rgba(0, 0, 0, 0.18) 3px 4px);
}
.nc-bp-photo-initial {
  position: relative;
  z-index: 2;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 54px;
  line-height: 1;
  color: rgba(244, 231, 211, 0.9);
  text-shadow: 0 0 20px rgba(210, 58, 46, 0.8);
}
.nc-bp-photo-lines {
  position: absolute;
  right: 7px;
  bottom: 20px;
  left: 7px;
  height: 1px;
  background: rgba(244, 231, 211, 0.25);
  box-shadow: 0 -4px 0 rgba(244, 231, 211, 0.12);
}
.nc-bp-photo small {
  position: absolute;
  bottom: 6px;
  left: 7px;
  z-index: 2;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.16em;
  color: rgba(244, 231, 211, 0.55);
}
.nc-bp-info {
  min-width: 0;
  flex: 1;
  padding-top: 7px;
}
.nc-bp-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.3em;
  color: #a3412f;
  text-transform: uppercase;
}
.nc-bp-name {
  margin-top: 4px;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
  font-size: clamp(34px, 10.5vw, 44px);
  font-weight: 900;
  line-height: 0.92;
  letter-spacing: -0.045em;
  text-transform: uppercase;
  color: #211816;
  overflow-wrap: anywhere;
}
.nc-bp-role {
  margin-top: 6px;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: #6e5a4c;
  text-transform: uppercase;
}
.nc-bp-venue {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 6px;
  margin-top: 16px;
  padding-top: 8px;
  border-top: 1px solid rgba(33, 24, 22, 0.2);
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #211816;
  text-transform: uppercase;
}
.nc-bp-city {
  padding: 1px 4px;
  background: #211816;
  color: #f4e7d3;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.08em;
}
.nc-bp-pass-foot {
  display: grid;
  grid-template-columns: 1fr 1fr;
  border-top: 1px dashed rgba(33, 24, 22, 0.35);
}
.nc-bp-field {
  display: flex;
  min-width: 0;
  flex-direction: column;
  gap: 3px;
  padding: 9px 14px;
  border-right: 1px solid rgba(33, 24, 22, 0.12);
  text-decoration: none;
  color: inherit;
}
.nc-bp-field:nth-child(2n) {
  border-right: 0;
}
.nc-bp-field span {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #97826f;
}
.nc-bp-field b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 9.5px;
  font-weight: 700;
  letter-spacing: 0.03em;
  color: #211816;
}
.nc-bp-field-link {
  position: relative;
  transition: background 0.25s ease;
}
.nc-bp-field-link:hover {
  background: rgba(210, 58, 46, 0.12);
}
.nc-bp-field-link:active {
  background: rgba(210, 58, 46, 0.22);
}
.nc-bp-field-link svg {
  position: absolute;
  right: 10px;
  bottom: 9px;
  color: #d23a2e;
  transition: transform 0.25s ease;
}
.nc-bp-field-link:hover svg {
  transform: rotate(-10deg) scale(1.12);
}
.nc-bp-field-code b {
  color: #a3412f;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.08em;
}
.nc-bp-empty {
  color: #b7a594 !important;
}
.nc-bp-pass-stripes {
  position: absolute;
  right: 0;
  bottom: 0;
  width: 44px;
  height: 44px;
  background: repeating-linear-gradient(135deg, #d23a2e 0 5px, transparent 5px 10px);
  opacity: 0.75;
  clip-path: polygon(100% 0, 100% 100%, 0 100%);
}
.nc-bp-note {
  margin-top: 20px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.22em;
  color: rgba(247, 232, 216, 0.34);
  animation: nc-bp-in 1s 0.55s both;
}
.nc-bp-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 14px;
  border-top: 1px solid rgba(247, 232, 216, 0.18);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(247, 232, 216, 0.3);
  animation: nc-bp-in 1s 0.65s both;
}
@keyframes nc-bp-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-bp-drop {
  from { opacity: 0; transform: translateY(-60px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-bp-pass {
  from { opacity: 0; transform: translateY(26px) rotate(1.2deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-bp-glow {
  0%, 100% { opacity: 0.8; }
  50% { opacity: 1; }
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

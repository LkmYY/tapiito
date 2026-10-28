<template>
  <div class="nc-nm-root relative min-h-screen overflow-hidden px-5 py-8 text-[#173b39]">
    <div class="nc-nm-sky"></div>
    <div class="nc-nm-rays"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col justify-center">
      <p class="nc-nm-topline">MOTEL NEON · CHECK-IN 23:00 · NO VACANCY AFTER 02:00</p>

      <article class="nc-nm-postcard">
        <div class="nc-nm-postmark" aria-hidden="true">
          <span>NEON</span>
          <small>JUL 24</small>
        </div>
        <div class="nc-nm-cancel" aria-hidden="true">
          <i></i><i></i><i></i>
        </div>

        <header class="nc-nm-head">
          <span class="nc-nm-vacancy"><i></i>VACANCY</span>
          <span class="nc-nm-room">ROOM {{ roomNo }}</span>
        </header>

        <p class="nc-nm-script">Welcome to the</p>
        <div class="nc-nm-sign">
          <span class="nc-nm-sign-back">MOTEL</span>
          <span class="nc-nm-sign-front">NEON</span>
        </div>

        <div class="nc-nm-divider">
          <span></span>
          <i>✦</i>
          <span></span>
        </div>

        <p class="nc-nm-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-nm-name">{{ payload?.name }}</h1>
        <p class="nc-nm-role">{{ payload?.role }}</p>

        <div class="nc-nm-venue">
          <span class="nc-nm-venue-name">{{ payload?.venue }}</span>
          <span class="nc-nm-venue-city">{{ payload?.city || '—' }}</span>
        </div>

        <section class="nc-nm-message">
          <div class="nc-nm-msg-row">
            <span class="nc-nm-msg-key">WECHAT</span>
            <span class="nc-nm-msg-dots"></span>
            <span class="nc-nm-msg-val">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-nm-msg-row nc-nm-msg-link">
            <span class="nc-nm-msg-key">PHONE</span>
            <span class="nc-nm-msg-dots"></span>
            <span class="nc-nm-msg-val">{{ payload.phone }}</span>
            <Phone :size="12" />
          </a>
          <div v-else class="nc-nm-msg-row">
            <span class="nc-nm-msg-key">PHONE</span>
            <span class="nc-nm-msg-dots"></span>
            <span class="nc-nm-msg-val nc-nm-empty"></span>
          </div>
        </section>

        <footer class="nc-nm-foot">
          <span class="nc-nm-keytag">
            <i></i>
            <b>KEY {{ roomNo }}</b>
          </span>
          <span>LEAVE THE LIGHT ON</span>
        </footer>
      </article>

      <p class="nc-nm-after">STAY ONE NIGHT · REMEMBER NOTHING</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const roomNo = computed(() => {
  const raw = (props.payload?.name || 'NOVA').toString()
  let total = 0
  for (let i = 0; i < raw.length; i += 1) total += raw.charCodeAt(i)
  return String(101 + (total % 88))
})
</script>

<style scoped>
.nc-nm-root {
  background: #0d2f30;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
}
.nc-nm-sky {
  position: absolute;
  inset: 0;
  background:
    linear-gradient(180deg, #0c2a31 0%, #144a4b 48%, #1b5f58 100%),
    radial-gradient(ellipse at 50% 100%, rgba(255, 139, 120, 0.25), transparent 60%);
  pointer-events: none;
}
.nc-nm-rays {
  position: absolute;
  inset: 0;
  background: repeating-conic-gradient(from 180deg at 50% 72%, rgba(255, 139, 120, 0.1) 0 3deg, transparent 3deg 14deg);
  mask-image: radial-gradient(ellipse at 50% 72%, #000 8%, transparent 70%);
  animation: nc-nm-rays 1.2s 0.2s both;
  pointer-events: none;
}
.nc-nm-topline {
  margin-bottom: 14px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(255, 236, 214, 0.55);
  animation: nc-nm-in 0.8s both;
}
.nc-nm-postcard {
  position: relative;
  overflow: hidden;
  padding: 20px 20px 16px;
  border: 2px solid #173b39;
  border-radius: 4px;
  background:
    radial-gradient(rgba(23, 59, 57, 0.1) 0.55px, transparent 0.55px),
    linear-gradient(150deg, #f9ecd4, #f1dfc2);
  background-size: 4px 4px, auto;
  box-shadow: 10px 12px 0 rgba(6, 32, 33, 0.55), 0 30px 60px -28px rgba(0, 0, 0, 0.8);
  animation: nc-nm-card 1s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.nc-nm-postcard:hover {
  transform: translate(-3px, -3px) rotate(-0.3deg);
  box-shadow: 13px 15px 0 rgba(6, 32, 33, 0.6), 0 36px 66px -28px rgba(0, 0, 0, 0.85);
}
.nc-nm-postmark {
  position: absolute;
  top: 14px;
  right: 14px;
  display: flex;
  width: 62px;
  height: 72px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2px;
  border: 2px dashed #c64e3f;
  background: rgba(198, 78, 63, 0.07);
  color: #c64e3f;
  transform: rotate(6deg);
  z-index: 2;
}
.nc-nm-postmark span {
  font-family: Georgia, serif;
  font-size: 17px;
  letter-spacing: 0.04em;
}
.nc-nm-postmark small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.14em;
}
.nc-nm-cancel {
  position: absolute;
  top: 44px;
  right: 76px;
  z-index: 2;
  display: flex;
  flex-direction: column;
  gap: 3px;
  transform: rotate(-10deg);
}
.nc-nm-cancel i {
  display: block;
  width: 58px;
  height: 1px;
  background: rgba(198, 78, 63, 0.5);
}
.nc-nm-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 9px;
  border-bottom: 2px solid #173b39;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: #173b39;
}
.nc-nm-vacancy {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 2px 6px;
  background: #173b39;
  color: #ffd8b8;
}
.nc-nm-vacancy i {
  width: 4px;
  height: 4px;
  border-radius: 50%;
  background: #ff8b78;
  box-shadow: 0 0 6px #ff8b78;
  animation: nc-nm-blink 1.8s ease-in-out infinite;
}
.nc-nm-room {
  color: #a45c47;
}
.nc-nm-script {
  margin-top: 24px;
  font-size: 11px;
  font-style: italic;
  letter-spacing: 0.08em;
  color: #8f6d52;
  animation: nc-nm-in 0.9s 0.1s both;
}
.nc-nm-sign {
  position: relative;
  margin-top: -2px;
  height: 78px;
  animation: nc-nm-in 0.9s 0.14s both;
}
.nc-nm-sign-back,
.nc-nm-sign-front {
  position: absolute;
  top: 0;
  left: 0;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
  font-size: clamp(58px, 18vw, 76px);
  font-weight: 900;
  line-height: 1;
  letter-spacing: -0.04em;
}
.nc-nm-sign-back {
  color: transparent;
  -webkit-text-stroke: 2px #173b39;
  transform: translate(4px, 4px);
}
.nc-nm-sign-front {
  color: #d9573f;
  text-shadow: 0 4px 0 #173b39;
}
.nc-nm-divider {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 12px 0 16px;
  color: #a45c47;
}
.nc-nm-divider span {
  flex: 1;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(23, 59, 57, 0.4));
}
.nc-nm-divider span:last-child {
  background: linear-gradient(90deg, rgba(23, 59, 57, 0.4), transparent);
}
.nc-nm-divider i {
  font-size: 7px;
  font-style: normal;
}
.nc-nm-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.4em;
  color: #a45c47;
  text-transform: uppercase;
  animation: nc-nm-in 0.9s 0.22s both;
}
.nc-nm-name {
  margin-top: 4px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(46px, 15vw, 62px);
  font-weight: 600;
  line-height: 0.95;
  letter-spacing: -0.04em;
  color: #173b39;
  overflow-wrap: anywhere;
  animation: nc-nm-in 1s 0.26s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-nm-role {
  margin-top: 7px;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.24em;
  color: #8f6d52;
  text-transform: uppercase;
  animation: nc-nm-in 0.9s 0.32s both;
}
.nc-nm-venue {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  gap: 8px;
  margin-top: 20px;
  animation: nc-nm-in 0.9s 0.38s both;
}
.nc-nm-venue-name {
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.12em;
  color: #173b39;
  text-transform: uppercase;
}
.nc-nm-venue-city {
  padding: 1px 5px;
  border: 1px solid #c64e3f;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #c64e3f;
}
.nc-nm-message {
  margin-top: 18px;
  border-top: 1px dashed rgba(23, 59, 57, 0.35);
  animation: nc-nm-in 1s 0.44s both;
}
.nc-nm-msg-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 9px 0;
  border-bottom: 1px dashed rgba(23, 59, 57, 0.25);
  text-decoration: none;
  color: inherit;
}
.nc-nm-msg-key {
  width: 48px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #8f6d52;
}
.nc-nm-msg-dots {
  flex: 1;
  border-bottom: 1px dotted rgba(23, 59, 57, 0.4);
  transform: translateY(-3px);
}
.nc-nm-msg-val {
  min-width: 0;
  max-width: 55%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10.5px;
  font-weight: 700;
  letter-spacing: 0.02em;
  color: #173b39;
}
.nc-nm-msg-link {
  margin: 0 -6px;
  padding: 9px 6px;
  transition: background 0.25s ease, padding 0.25s ease;
}
.nc-nm-msg-link:hover {
  background: rgba(198, 78, 63, 0.12);
  padding-left: 12px;
}
.nc-nm-msg-link:active {
  background: rgba(198, 78, 63, 0.24);
}
.nc-nm-msg-link svg {
  flex: none;
  color: #c64e3f;
}
.nc-nm-empty {
  color: #b5a28d;
}
.nc-nm-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  margin-top: 16px;
}
.nc-nm-keytag {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 4px 8px 4px 5px;
  border: 1px solid #173b39;
  border-radius: 2px;
  background: #e7d2b0;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.1em;
  color: #173b39;
  transform: rotate(-2deg);
}
.nc-nm-keytag i {
  width: 9px;
  height: 9px;
  border: 1px solid #173b39;
  border-radius: 50%;
}
.nc-nm-keytag b {
  font-weight: 700;
}
.nc-nm-foot > span:last-child {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #a45c47;
}
.nc-nm-after {
  margin-top: 20px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(255, 236, 214, 0.4);
  animation: nc-nm-in 1s 0.6s both;
}
@keyframes nc-nm-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-nm-card {
  from { opacity: 0; transform: translateY(30px) rotate(1.4deg); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-nm-rays {
  from { opacity: 0; transform: scale(0.9); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-nm-blink {
  0%, 100% { opacity: 1; }
  55% { opacity: 0.3; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-nm-root *,
  .nc-nm-root *::before,
  .nc-nm-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

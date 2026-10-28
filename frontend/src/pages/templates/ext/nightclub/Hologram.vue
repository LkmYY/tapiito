<template>
  <div class="nc-hg-root relative min-h-screen overflow-hidden px-5 py-8 text-[#f4f0ff]">
    <div class="nc-hg-aurora nc-hg-aurora-a"></div>
    <div class="nc-hg-aurora nc-hg-aurora-b"></div>
    <div class="nc-hg-sparkles" aria-hidden="true">
      <i v-for="n in 12" :key="n" :style="{ left: spark[n - 1].x + '%', top: spark[n - 1].y + '%', '--sd': spark[n - 1].d + 's' }"></i>
    </div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col justify-center">
      <header class="nc-hg-head">
        <span>REFRACTION 001</span>
        <span>HOLO / NFC {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-hg-card">
        <div class="nc-hg-sheen"></div>
        <div class="nc-hg-top">
          <span class="nc-hg-tag">{{ payload?.alias || 'NO ALIAS' }}</span>
          <span class="nc-hg-icon" aria-hidden="true">◇</span>
        </div>

        <p class="nc-hg-kicker">FREQUENCY OF THE NIGHT</p>
        <h1 class="nc-hg-name">{{ payload?.name }}</h1>
        <p class="nc-hg-role">{{ payload?.role }}</p>

        <div class="nc-hg-venue">
          <span>{{ payload?.venue }}</span>
          <span class="nc-hg-city">{{ payload?.city || '—' }}</span>
        </div>

        <div class="nc-hg-contacts">
          <div class="nc-hg-contact">
            <span class="nc-hg-contact-label">WECHAT</span>
            <span class="nc-hg-contact-value">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-hg-contact nc-hg-contact-link">
            <span class="nc-hg-contact-label">PHONE</span>
            <span class="nc-hg-contact-value">{{ payload.phone }}</span>
            <Phone :size="13" />
          </a>
          <div v-else class="nc-hg-contact">
            <span class="nc-hg-contact-label">PHONE</span>
            <span class="nc-hg-contact-value nc-hg-empty"></span>
          </div>
        </div>

        <div class="nc-hg-foot">
          <span class="nc-hg-prism"></span>
          <span>PRIVATE FREQUENCY · SHARE WITH CARE</span>
        </div>
      </main>

      <p class="nc-hg-outside">LIGHT BENDS · THE NIGHT DOESN'T</p>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])

const spark = [
  { x: 12, y: 18, d: 0.2 }, { x: 82, y: 12, d: 1.1 }, { x: 28, y: 74, d: 0.6 },
  { x: 92, y: 62, d: 1.7 }, { x: 6, y: 54, d: 1.3 }, { x: 68, y: 84, d: 0.4 },
  { x: 44, y: 8, d: 2.1 }, { x: 76, y: 40, d: 0.9 }, { x: 18, y: 90, d: 1.5 },
  { x: 55, y: 66, d: 0.1 }, { x: 88, y: 88, d: 2.4 }, { x: 36, y: 34, d: 1.9 },
]
</script>

<style scoped>
.nc-hg-root {
  background: #07050e;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-hg-aurora {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.5;
  pointer-events: none;
}
.nc-hg-aurora-a {
  top: -120px;
  left: -80px;
  width: 320px;
  height: 320px;
  background: conic-gradient(from 20deg, #4cc9f0, #b5179e, #f72585, #4cc9f0);
  animation: nc-hg-drift 12s ease-in-out infinite;
}
.nc-hg-aurora-b {
  right: -100px;
  bottom: -80px;
  width: 300px;
  height: 300px;
  background: conic-gradient(from 190deg, #7b2cbf, #06d6a0, #4361ee, #7b2cbf);
  animation: nc-hg-drift 15s 1s ease-in-out infinite reverse;
}
.nc-hg-sparkles {
  position: absolute;
  inset: 0;
  pointer-events: none;
}
.nc-hg-sparkles i {
  position: absolute;
  width: 3px;
  height: 3px;
  background: #fff;
  transform: rotate(45deg);
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.9);
  animation: nc-hg-twinkle 3s var(--sd) ease-in-out infinite;
}
.nc-hg-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 16px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(244, 240, 255, 0.4);
  animation: nc-hg-in 0.9s both;
}
.nc-hg-card {
  position: relative;
  overflow: hidden;
  padding: 24px 22px 18px;
  border-radius: 18px;
  background:
    linear-gradient(150deg, rgba(255, 255, 255, 0.13), rgba(255, 255, 255, 0.035) 42%, rgba(255, 255, 255, 0.09)),
    rgba(13, 10, 24, 0.76);
  box-shadow:
    0 0 0 1px rgba(255, 255, 255, 0.2),
    0 26px 80px -22px rgba(123, 44, 191, 0.6),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(18px);
  animation: nc-hg-card 1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.5s cubic-bezier(0.22, 1, 0.36, 1), box-shadow 0.5s ease;
}
.nc-hg-card:hover {
  transform: perspective(900px) rotateX(1.4deg) rotateY(-1.2deg) translateY(-4px);
  box-shadow:
    0 0 0 1px rgba(255, 255, 255, 0.32),
    0 34px 90px -20px rgba(76, 201, 240, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
}
.nc-hg-card::before {
  content: '';
  position: absolute;
  inset: 0;
  border-radius: 18px;
  padding: 1px;
  background: conic-gradient(from 120deg, #4cc9f0, #f72585, #ffd166, #06d6a0, #7b2cbf, #4cc9f0);
  -webkit-mask: linear-gradient(#000 0 0) content-box, linear-gradient(#000 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  opacity: 0.75;
  pointer-events: none;
}
.nc-hg-sheen {
  position: absolute;
  top: -40%;
  left: -80%;
  width: 55%;
  height: 180%;
  background: linear-gradient(100deg, transparent, rgba(255, 255, 255, 0.18), transparent);
  transform: rotate(14deg);
  animation: nc-hg-sweep 8s 1.2s ease-in-out infinite;
  pointer-events: none;
}
.nc-hg-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: relative;
  z-index: 1;
}
.nc-hg-tag {
  padding: 4px 9px;
  border: 1px solid rgba(255, 255, 255, 0.36);
  border-radius: 999px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: rgba(244, 240, 255, 0.78);
  text-transform: uppercase;
}
.nc-hg-icon {
  font-size: 19px;
  color: #fff;
  text-shadow: 0 0 14px rgba(76, 201, 240, 0.9);
  animation: nc-hg-spin 9s linear infinite;
}
.nc-hg-kicker {
  position: relative;
  z-index: 1;
  margin-top: 30px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.28em;
  color: rgba(244, 240, 255, 0.44);
}
.nc-hg-name {
  position: relative;
  z-index: 1;
  margin-top: 7px;
  background: linear-gradient(92deg, #fff 0%, #a9e8ff 22%, #ffb3e6 48%, #ffe9a8 68%, #b8ffe4 86%, #fff 100%);
  background-size: 220% 100%;
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  font-size: clamp(54px, 17vw, 72px);
  font-weight: 700;
  line-height: 0.92;
  letter-spacing: -0.05em;
  overflow-wrap: anywhere;
  animation: nc-hg-shift 9s linear infinite;
}
.nc-hg-role {
  position: relative;
  z-index: 1;
  margin-top: 10px;
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.24em;
  color: rgba(244, 240, 255, 0.62);
}
.nc-hg-venue {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 22px;
  font-size: 11px;
  letter-spacing: 0.12em;
  color: rgba(244, 240, 255, 0.88);
}
.nc-hg-city {
  padding: 1px 6px;
  border: 1px solid rgba(255, 255, 255, 0.24);
  border-radius: 999px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.1em;
  color: rgba(244, 240, 255, 0.55);
}
.nc-hg-contacts {
  position: relative;
  z-index: 1;
  display: flex;
  flex-direction: column;
  gap: 7px;
  margin-top: 20px;
}
.nc-hg-contact {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 11px 12px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 11px;
  background: rgba(255, 255, 255, 0.045);
  text-decoration: none;
  color: inherit;
}
.nc-hg-contact-label {
  width: 46px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: rgba(244, 240, 255, 0.4);
}
.nc-hg-contact-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  letter-spacing: 0.04em;
  color: rgba(244, 240, 255, 0.92);
}
.nc-hg-contact-link {
  transition: background 0.3s ease, border-color 0.3s ease, transform 0.3s ease;
}
.nc-hg-contact-link svg {
  flex: none;
  color: #a9e8ff;
  transition: transform 0.3s ease;
}
.nc-hg-contact-link:hover {
  background: rgba(255, 255, 255, 0.12);
  border-color: rgba(255, 255, 255, 0.34);
  transform: translateX(3px);
}
.nc-hg-contact-link:active {
  background: rgba(255, 255, 255, 0.18);
}
.nc-hg-contact-link:hover svg {
  transform: rotate(-12deg) scale(1.1);
}
.nc-hg-empty {
  color: rgba(244, 240, 255, 0.3);
}
.nc-hg-foot {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  gap: 9px;
  margin-top: 20px;
  padding-top: 13px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.16em;
  color: rgba(244, 240, 255, 0.34);
}
.nc-hg-prism {
  width: 14px;
  height: 8px;
  background: linear-gradient(90deg, #4cc9f0, #f72585, #ffd166, #06d6a0);
  clip-path: polygon(0 0, 100% 50%, 0 100%);
}
.nc-hg-outside {
  padding-top: 24px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.26em;
  color: rgba(244, 240, 255, 0.26);
  animation: nc-hg-in 1.1s 0.6s both;
}
@keyframes nc-hg-card {
  from { opacity: 0; transform: translateY(26px) scale(0.96); filter: blur(8px); }
  to { opacity: 1; transform: none; filter: blur(0); }
}
@keyframes nc-hg-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-hg-drift {
  0%, 100% { transform: translate(0, 0) scale(1) rotate(0deg); }
  50% { transform: translate(24px, 16px) scale(1.1) rotate(30deg); }
}
@keyframes nc-hg-sweep {
  0%, 50% { left: -80%; }
  80%, 100% { left: 130%; }
}
@keyframes nc-hg-shift {
  from { background-position: 0% 50%; }
  to { background-position: 220% 50%; }
}
@keyframes nc-hg-spin {
  from { transform: rotate(0); }
  to { transform: rotate(360deg); }
}
@keyframes nc-hg-twinkle {
  0%, 100% { opacity: 0.2; transform: rotate(45deg) scale(0.7); }
  50% { opacity: 1; transform: rotate(45deg) scale(1.2); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-hg-root *,
  .nc-hg-root *::before,
  .nc-hg-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

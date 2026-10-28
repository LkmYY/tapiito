<template>
  <div class="nc-nn-root relative min-h-screen overflow-hidden px-5 py-7 text-white">
    <div class="nc-nn-orb nc-nn-orb-a"></div>
    <div class="nc-nn-orb nc-nn-orb-b"></div>
    <div class="nc-nn-grid"></div>
    <div class="nc-nn-scan"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3.5rem)] w-full max-w-md flex-col">
      <header class="nc-nn-top">
        <span class="nc-nn-live"><i></i>DOORS OPEN</span>
        <span>NFC / {{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-nn-stage">
        <div class="nc-nn-halo"></div>
        <div class="nc-nn-badge">
          <span>{{ initial }}</span>
        </div>
        <p class="nc-nn-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-nn-name">{{ payload?.name }}</h1>
        <p class="nc-nn-role">{{ payload?.role }}</p>

        <div class="nc-nn-venue">
          <span class="nc-nn-venue-icon"><MapPin :size="13" /></span>
          <span class="nc-nn-venue-name">{{ payload?.venue }}</span>
          <span class="nc-nn-venue-city">{{ payload?.city || '—' }}</span>
        </div>

        <div class="nc-nn-links">
          <div class="nc-nn-link">
            <span class="nc-nn-link-icon"><MessageCircle :size="16" /></span>
            <span class="nc-nn-link-copy">
              <small>WECHAT</small>
              <b>{{ payload?.wechat }}</b>
            </span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-nn-link nc-nn-link-live">
            <span class="nc-nn-link-icon"><Phone :size="16" /></span>
            <span class="nc-nn-link-copy">
              <small>CALL DIRECT</small>
              <b>{{ payload.phone }}</b>
            </span>
            <ArrowUpRight :size="16" class="nc-nn-arrow" />
          </a>
          <div v-else class="nc-nn-link nc-nn-link-muted">
            <span class="nc-nn-link-icon"><Phone :size="16" /></span>
            <span class="nc-nn-link-copy">
              <small>CALL DIRECT</small>
              <b></b>
            </span>
          </div>
        </div>
      </main>

      <footer class="nc-nn-foot">
        <span>STAY FOR ONE MORE TRACK</span>
        <span class="nc-nn-foot-rule"></span>
        <span>SHANGHAI · CN</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, MessageCircle, MapPin, ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.nc-nn-root {
  background: #06030f;
  font-family: 'Inter', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-nn-orb {
  position: absolute;
  border-radius: 999px;
  filter: blur(80px);
  pointer-events: none;
}
.nc-nn-orb-a {
  top: -100px;
  left: -90px;
  width: 300px;
  height: 300px;
  background: rgba(168, 85, 247, 0.42);
  animation: nc-nn-drift 10s ease-in-out infinite;
}
.nc-nn-orb-b {
  right: -110px;
  bottom: 10px;
  width: 280px;
  height: 280px;
  background: rgba(34, 211, 238, 0.26);
  animation: nc-nn-drift 13s ease-in-out infinite reverse;
}
.nc-nn-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(232, 121, 249, 0.08) 1px, transparent 1px),
    linear-gradient(90deg, rgba(232, 121, 249, 0.08) 1px, transparent 1px);
  background-size: 32px 32px;
  mask-image: radial-gradient(ellipse at 50% 38%, #000 15%, transparent 72%);
  animation: nc-nn-grid-in 1.4s ease both;
  pointer-events: none;
}
.nc-nn-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 3px, rgba(255, 255, 255, 0.02) 3px 4px);
  pointer-events: none;
}
.nc-nn-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.26em;
  color: rgba(255, 255, 255, 0.34);
  animation: nc-nn-fade 0.8s both;
}
.nc-nn-live {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  color: #67e8f9;
}
.nc-nn-live i {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #22d3ee;
  box-shadow: 0 0 10px #22d3ee;
  animation: nc-nn-blink 1.7s ease-in-out infinite;
}
.nc-nn-stage {
  position: relative;
  display: flex;
  flex: 1;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 46px 0 30px;
  text-align: center;
}
.nc-nn-halo {
  position: absolute;
  top: 86px;
  width: 230px;
  height: 230px;
  border: 1px solid rgba(232, 121, 249, 0.16);
  border-radius: 50%;
  box-shadow: 0 0 80px rgba(168, 85, 247, 0.18), inset 0 0 60px rgba(34, 211, 238, 0.06);
  animation: nc-nn-halo 8s ease-in-out infinite;
  pointer-events: none;
}
.nc-nn-badge {
  position: relative;
  display: flex;
  width: 70px;
  height: 70px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(103, 232, 249, 0.5);
  border-radius: 50%;
  background: radial-gradient(circle at 32% 26%, rgba(34, 211, 238, 0.3), rgba(88, 28, 135, 0.6));
  box-shadow: 0 0 30px rgba(34, 211, 238, 0.25), inset 0 0 24px rgba(232, 121, 249, 0.18);
  color: #a5f3fc;
  font-size: 28px;
  font-weight: 200;
  animation: nc-nn-rise 0.9s 0.05s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-nn-alias {
  margin-top: 26px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.5em;
  color: #e879f9;
  text-transform: uppercase;
  animation: nc-nn-rise 0.9s 0.12s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-nn-name {
  margin-top: 8px;
  font-size: clamp(58px, 19vw, 82px);
  font-weight: 200;
  line-height: 0.94;
  letter-spacing: -0.04em;
  text-shadow: 0 0 30px rgba(232, 121, 249, 0.45), 0 0 80px rgba(168, 85, 247, 0.25);
  overflow-wrap: anywhere;
  animation: nc-nn-rise 0.95s 0.18s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-nn-role {
  margin-top: 14px;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.32em;
  color: #67e8f9;
  animation: nc-nn-rise 0.9s 0.24s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-nn-venue {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  margin-top: 22px;
  padding: 7px 12px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 2px;
  background: rgba(255, 255, 255, 0.02);
  font-size: 10.5px;
  letter-spacing: 0.14em;
  color: rgba(255, 255, 255, 0.72);
  animation: nc-nn-rise 0.9s 0.3s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-nn-venue-icon {
  display: flex;
  color: #f0abfc;
}
.nc-nn-venue-name {
  font-weight: 500;
}
.nc-nn-venue-city {
  padding-left: 8px;
  border-left: 1px solid rgba(255, 255, 255, 0.16);
  color: rgba(255, 255, 255, 0.4);
}
.nc-nn-links {
  width: 100%;
  margin-top: 34px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  animation: nc-nn-rise 0.9s 0.38s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-nn-link {
  display: flex;
  align-items: center;
  gap: 13px;
  padding: 14px 4px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.07);
  text-align: left;
  text-decoration: none;
  color: inherit;
}
.nc-nn-link-live {
  transition: background 0.3s ease, padding 0.3s ease;
}
.nc-nn-link-live:hover {
  background: linear-gradient(90deg, rgba(232, 121, 249, 0.1), transparent);
  padding-left: 10px;
}
.nc-nn-link-live:active {
  background: rgba(232, 121, 249, 0.16);
}
.nc-nn-link-icon {
  display: flex;
  width: 34px;
  height: 34px;
  flex: none;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(103, 232, 249, 0.22);
  border-radius: 50%;
  color: #67e8f9;
  background: rgba(34, 211, 238, 0.05);
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
}
.nc-nn-link-live:hover .nc-nn-link-icon {
  border-color: rgba(232, 121, 249, 0.65);
  box-shadow: 0 0 16px rgba(232, 121, 249, 0.3);
}
.nc-nn-link-copy {
  display: flex;
  min-width: 0;
  flex: 1;
  flex-direction: column;
  gap: 2px;
}
.nc-nn-link-copy small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.3em;
  color: rgba(255, 255, 255, 0.3);
}
.nc-nn-link-copy b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 13.5px;
  font-weight: 400;
  letter-spacing: 0.04em;
  color: rgba(255, 255, 255, 0.9);
}
.nc-nn-link-muted b {
  color: rgba(255, 255, 255, 0.32);
}
.nc-nn-arrow {
  flex: none;
  color: rgba(255, 255, 255, 0.25);
  transition: transform 0.3s ease, color 0.3s ease;
}
.nc-nn-link-live:hover .nc-nn-arrow {
  transform: translate(2px, -2px);
  color: #e879f9;
}
.nc-nn-foot {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 10px 0 2px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.22em;
  color: rgba(255, 255, 255, 0.24);
  animation: nc-nn-fade 1.2s 0.7s both;
}
.nc-nn-foot-rule {
  width: 26px;
  height: 1px;
  background: rgba(232, 121, 249, 0.5);
}
@keyframes nc-nn-rise {
  from { opacity: 0; transform: translateY(22px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-nn-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes nc-nn-drift {
  0%, 100% { transform: translate3d(0, 0, 0) scale(1); }
  50% { transform: translate3d(18px, -14px, 0) scale(1.08); }
}
@keyframes nc-nn-blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.25; }
}
@keyframes nc-nn-halo {
  0%, 100% { transform: scale(1); opacity: 0.8; }
  50% { transform: scale(1.06); opacity: 1; }
}
@keyframes nc-nn-grid-in {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-nn-root *,
  .nc-nn-root *::before,
  .nc-nn-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

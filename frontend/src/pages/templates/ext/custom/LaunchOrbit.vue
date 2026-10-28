<template>
  <div class="cu-lo-root min-h-screen">
    <div class="cu-lo-stars" aria-hidden="true"></div>
    <div class="cu-lo-wrap mx-auto w-full max-w-md">
      <header class="cu-lo-top cu-lo-in">
        <span class="cu-lo-brand"><span class="cu-lo-brand-dot"></span>{{ payload?.companyName || '企业主页' }}</span>
        <span class="cu-lo-version">v2.0 / LIVE</span>
      </header>

      <section class="cu-lo-hero">
        <div class="cu-lo-orbit cu-lo-in cu-lo-d1" aria-hidden="true">
          <span class="cu-lo-ring cu-lo-ring--a"></span>
          <span class="cu-lo-ring cu-lo-ring--b"></span>
          <span class="cu-lo-core">{{ initial }}</span>
          <span class="cu-lo-satellite"></span>
        </div>

        <p class="cu-lo-kicker cu-lo-in cu-lo-d1">NEXT-GEN COMPANY PROFILE</p>
        <h1 class="cu-lo-name cu-lo-in cu-lo-d2">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-lo-desc cu-lo-in cu-lo-d2">{{ payload?.description }}</p>
      </section>

      <div class="cu-lo-metrics cu-lo-in cu-lo-d3">
        <div><strong>{{ String(services.length).padStart(2, '0') }}</strong><span>服务模块</span></div>
        <div><strong>24/7</strong><span>在线响应</span></div>
        <div><strong>1:1</strong><span>专属对接</span></div>
      </div>

      <section class="cu-lo-modules cu-lo-in cu-lo-d3">
        <div class="cu-lo-section-head">
          <h2>MODULES</h2>
          <span>服务能力</span>
        </div>
        <div v-if="services.length" class="cu-lo-grid">
          <article v-for="(service, i) in services" :key="i" class="cu-lo-chip">
            <span class="cu-lo-chip-id">M{{ String(i + 1).padStart(2, '0') }}</span>
            <div class="cu-lo-chip-bar"><i :style="{ width: 58 + i * 12 + '%' }"></i></div>
            <h3>{{ service.name }}</h3>
          </article>
        </div>
      </section>

      <section class="cu-lo-contact cu-lo-in cu-lo-d4">
        <div class="cu-lo-contact-copy">
          <span>READY TO LAUNCH?</span>
          <p>接通项目顾问，获取定制方案</p>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-lo-cta">
          <Phone :size="15" />
          <span>{{ payload.phone }}</span>
          <ArrowUpRight :size="15" />
        </a>
      </section>

      <footer class="cu-lo-foot cu-lo-in cu-lo-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-lo-root {
  position: relative;
  overflow: hidden;
  background: #060a1a;
  color: #e7ecff;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.cu-lo-stars {
  position: absolute;
  inset: 0;
  pointer-events: none;
  background-image:
    radial-gradient(circle at 20% 14%, rgba(103, 232, 249, 0.5) 1px, transparent 1.5px),
    radial-gradient(circle at 76% 24%, rgba(167, 139, 250, 0.55) 1px, transparent 1.5px),
    radial-gradient(circle at 42% 68%, rgba(255, 255, 255, 0.35) 1px, transparent 1.4px),
    radial-gradient(circle at 88% 78%, rgba(103, 232, 249, 0.35) 1px, transparent 1.4px);
  background-size: 190px 190px, 230px 230px, 150px 150px, 270px 270px;
}
.cu-lo-root::before {
  content: '';
  position: absolute;
  top: -140px;
  right: -120px;
  height: 340px;
  width: 340px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(37, 99, 235, 0.4), transparent 66%);
  pointer-events: none;
}
.cu-lo-wrap {
  position: relative;
  z-index: 2;
  padding: 22px 18px 30px;
}
.cu-lo-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.cu-lo-brand {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: #b9c4ea;
  max-width: 60%;
}
.cu-lo-brand-dot {
  height: 7px;
  width: 7px;
  border-radius: 50%;
  background: #67e8f9;
  box-shadow: 0 0 10px #67e8f9;
  flex-shrink: 0;
}
.cu-lo-version {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.16em;
  color: #67e8f9;
  border: 1px solid rgba(103, 232, 249, 0.4);
  border-radius: 999px;
  padding: 4px 8px;
}
.cu-lo-hero {
  position: relative;
  padding-top: 40px;
}
.cu-lo-orbit {
  position: absolute;
  right: -30px;
  top: 0;
  height: 190px;
  width: 190px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.cu-lo-ring {
  position: absolute;
  border-radius: 50%;
  border: 1px solid rgba(103, 232, 249, 0.25);
}
.cu-lo-ring--a {
  inset: 0;
  transform: rotate(-18deg) scaleY(0.82);
}
.cu-lo-ring--b {
  inset: 26px;
  border-color: rgba(167, 139, 250, 0.3);
  transform: rotate(24deg) scaleY(0.7);
}
.cu-lo-core {
  position: relative;
  z-index: 2;
  height: 52px;
  width: 52px;
  border-radius: 50%;
  background: radial-gradient(circle at 35% 30%, #dbeafe, #3b82f6 55%, #1e3a8a);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  color: #ffffff;
  box-shadow: 0 0 34px rgba(59, 130, 246, 0.65);
}
.cu-lo-satellite {
  position: absolute;
  top: 24px;
  right: 22px;
  height: 7px;
  width: 7px;
  border-radius: 50%;
  background: #a78bfa;
  box-shadow: 0 0 12px #a78bfa;
}
.cu-lo-kicker {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.26em;
  color: #7dd3fc;
}
.cu-lo-name {
  position: relative;
  margin-top: 11px;
  max-width: 6em;
  font-size: 44px;
  line-height: 1.04;
  font-weight: 700;
  letter-spacing: -0.045em;
  color: #f1f5ff;
  word-break: break-word;
  text-shadow: 0 0 40px rgba(59, 130, 246, 0.3);
}
.cu-lo-desc {
  position: relative;
  margin-top: 13px;
  max-width: 32ch;
  font-size: 12px;
  line-height: 1.85;
  color: #97a3c9;
}
.cu-lo-metrics {
  margin-top: 24px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  border: 1px solid rgba(125, 211, 252, 0.2);
  border-radius: 14px;
  background: rgba(15, 23, 52, 0.6);
  overflow: hidden;
}
.cu-lo-metrics div {
  padding: 11px 8px;
  text-align: center;
}
.cu-lo-metrics div + div {
  border-left: 1px solid rgba(125, 211, 252, 0.14);
}
.cu-lo-metrics strong {
  display: block;
  font-size: 18px;
  font-weight: 700;
  color: #dbeafe;
  letter-spacing: -0.02em;
}
.cu-lo-metrics span {
  display: block;
  margin-top: 2px;
  font-size: 8px;
  letter-spacing: 0.06em;
  color: #7c89b0;
}
.cu-lo-modules {
  margin-top: 26px;
}
.cu-lo-section-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 1px solid rgba(125, 211, 252, 0.2);
  padding-bottom: 7px;
}
.cu-lo-section-head h2 {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.24em;
  color: #a5f3fc;
}
.cu-lo-section-head span {
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #7c89b0;
}
.cu-lo-grid {
  margin-top: 11px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}
.cu-lo-chip {
  border: 1px solid rgba(125, 211, 252, 0.22);
  border-radius: 13px;
  background: linear-gradient(160deg, rgba(30, 58, 138, 0.45), rgba(10, 16, 38, 0.72));
  padding: 12px 12px 13px;
  transition: border-color 0.25s ease, transform 0.25s ease;
}
.cu-lo-chip:hover {
  border-color: rgba(103, 232, 249, 0.65);
  transform: translateY(-2px);
}
.cu-lo-chip-id {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #67e8f9;
}
.cu-lo-chip-bar {
  margin-top: 10px;
  height: 2px;
  width: 100%;
  background: rgba(125, 211, 252, 0.16);
}
.cu-lo-chip-bar i {
  display: block;
  height: 100%;
  background: linear-gradient(90deg, #3b82f6, #67e8f9);
  box-shadow: 0 0 8px rgba(103, 232, 249, 0.7);
}
.cu-lo-chip h3 {
  margin-top: 9px;
  font-size: 13px;
  font-weight: 600;
  line-height: 1.35;
  color: #e2eaff;
}
.cu-lo-empty {
  margin-top: 10px;
  border: 1px dashed rgba(125, 211, 252, 0.3);
  border-radius: 12px;
  padding: 12px;
  font-size: 10px;
  color: #7c89b0;
}
.cu-lo-contact {
  margin-top: 24px;
  border: 1px solid rgba(103, 232, 249, 0.35);
  border-radius: 16px;
  background: linear-gradient(150deg, rgba(30, 64, 175, 0.5), rgba(8, 15, 36, 0.9));
  padding: 16px;
  box-shadow: 0 0 34px rgba(37, 99, 235, 0.22);
}
.cu-lo-contact-copy span {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: #67e8f9;
}
.cu-lo-contact-copy p {
  margin-top: 4px;
  font-size: 11px;
  color: #b9c4ea;
}
.cu-lo-cta {
  margin-top: 12px;
  display: flex;
  align-items: center;
  gap: 8px;
  border-radius: 10px;
  background: linear-gradient(90deg, #2563eb, #0891b2);
  padding: 12px 13px;
  font-size: 17px;
  font-weight: 600;
  color: #ffffff;
  box-shadow: 0 0 24px rgba(37, 99, 235, 0.5);
  transition: filter 0.2s ease;
}
.cu-lo-cta span {
  flex: 1;
}
.cu-lo-cta:hover {
  filter: brightness(1.12);
}
.cu-lo-contact .cu-lo-empty {
  margin-top: 10px;
  color: #b9c4ea;
}
.cu-lo-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #5e6b92;
}
@keyframes cu-lo-in {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-lo-in {
  animation: cu-lo-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-lo-d1 {
  animation-delay: 0.07s;
}
.cu-lo-d2 {
  animation-delay: 0.17s;
}
.cu-lo-d3 {
  animation-delay: 0.26s;
}
.cu-lo-d4 {
  animation-delay: 0.35s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-lo-in {
    animation: none;
  }
  .cu-lo-chip,
  .cu-lo-cta {
    transition: none;
  }
}
</style>

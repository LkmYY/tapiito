<template>
  <div class="cu-bp-root min-h-screen">
    <div class="cu-bp-wrap mx-auto w-full max-w-md">
      <header class="cu-bp-top cu-bp-in">
        <div class="cu-bp-brand">
          <span class="cu-bp-cross"></span>
          <span>DRAWING / {{ nfcId ? 'N' + nfcId : 'N-000' }}</span>
        </div>
        <span class="cu-bp-rev">REV. C</span>
      </header>

      <section class="cu-bp-title cu-bp-in cu-bp-d1">
        <p class="cu-bp-overline">企业服务 / SERVICE ARCHITECTURE</p>
        <h1>{{ payload?.companyName || '企业主页' }}</h1>
        <div class="cu-bp-rule">
          <span></span>
          <i>SCALE 1:1</i>
        </div>
        <p class="cu-bp-desc">{{ payload?.description }}</p>
      </section>

      <section class="cu-bp-plan cu-bp-in cu-bp-d2">
        <div class="cu-bp-plan-head">
          <span>FIG. 01 — SERVICE PLAN</span>
          <span>UNIT / CN</span>
        </div>
        <div class="cu-bp-svg">
          <svg viewBox="0 0 320 150" fill="none" aria-hidden="true">
            <rect x="10" y="10" width="300" height="130" stroke="currentColor" stroke-width="1" stroke-dasharray="5 4" />
            <rect x="30" y="30" width="120" height="90" stroke="currentColor" stroke-width="1.4" />
            <rect x="170" y="30" width="120" height="40" stroke="currentColor" stroke-width="1.4" />
            <rect x="170" y="80" width="120" height="40" stroke="currentColor" stroke-width="1.4" />
            <line x1="30" y1="120" x2="290" y2="120" stroke="currentColor" stroke-width="0.8" />
            <circle cx="90" cy="75" r="18" stroke="currentColor" stroke-width="0.8" />
            <path d="M72 75h36M90 57v36" stroke="currentColor" stroke-width="0.6" />
            <text x="30" y="145" fill="currentColor" font-size="6" font-family="monospace" letter-spacing="1">A</text>
            <text x="150" y="145" fill="currentColor" font-size="6" font-family="monospace" letter-spacing="1">MAIN HALL</text>
            <text x="230" y="145" fill="currentColor" font-size="6" font-family="monospace" letter-spacing="1">B</text>
          </svg>
        </div>
      </section>

      <section class="cu-bp-modules cu-bp-in cu-bp-d3">
        <div class="cu-bp-plan-head">
          <span>FIG. 02 — SCOPE OF WORK</span>
          <span>{{ String(services.length).padStart(2, '0') }} MOD.</span>
        </div>
        <div v-if="services.length" class="cu-bp-module-list">
          <article v-for="(service, i) in services" :key="i" class="cu-bp-module">
            <span class="cu-bp-module-tick"></span>
            <span class="cu-bp-module-no">M-{{ String(i + 1).padStart(2, '0') }}</span>
            <h2>{{ service.name }}</h2>
            <span class="cu-bp-module-dim">{{ 12 + i * 3 }} × {{ 8 + i * 2 }}</span>
          </article>
        </div>
      </section>

      <section class="cu-bp-contact cu-bp-in cu-bp-d4">
        <div class="cu-bp-contact-row">
          <span class="cu-bp-contact-key">CONTACT</span>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-bp-contact-val">
            <Phone :size="14" />
            {{ payload.phone }}
          </a>
        </div>
        <div class="cu-bp-contact-row">
          <span class="cu-bp-contact-key">DRAWN BY</span>
          <span class="cu-bp-contact-val">{{ payload?.companyName || '企业主页' }}</span>
        </div>
      </section>

      <footer class="cu-bp-foot cu-bp-in cu-bp-d4">
        <span>ALL DIMENSIONS IN MILLIMETRES</span>
        <span>{{ nfcId ? 'NFC/' + nfcId : 'NFC/—' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
</script>

<style scoped>
.cu-bp-root {
  overflow: hidden;
  background-color: #0a2a46;
  background-image:
    linear-gradient(rgba(132, 197, 235, 0.13) 1px, transparent 1px),
    linear-gradient(90deg, rgba(132, 197, 235, 0.13) 1px, transparent 1px),
    linear-gradient(rgba(132, 197, 235, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(132, 197, 235, 0.05) 1px, transparent 1px);
  background-size: 64px 64px, 64px 64px, 16px 16px, 16px 16px;
  color: #d7ecf9;
  font-family: 'Courier New', 'PingFang SC', monospace;
}
.cu-bp-wrap {
  padding: 20px 18px 30px;
}
.cu-bp-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border: 1px solid rgba(170, 216, 242, 0.55);
  padding: 8px 10px;
  font-size: 8px;
  letter-spacing: 0.18em;
}
.cu-bp-brand {
  display: flex;
  align-items: center;
  gap: 8px;
}
.cu-bp-cross {
  position: relative;
  display: inline-block;
  height: 10px;
  width: 10px;
}
.cu-bp-cross::before,
.cu-bp-cross::after {
  content: '';
  position: absolute;
  background: #8ed8ff;
}
.cu-bp-cross::before {
  left: 0;
  top: 4px;
  height: 1px;
  width: 10px;
}
.cu-bp-cross::after {
  top: 0;
  left: 4px;
  width: 1px;
  height: 10px;
}
.cu-bp-rev {
  color: #8ed8ff;
}
.cu-bp-title {
  margin-top: 28px;
}
.cu-bp-overline {
  font-size: 8px;
  letter-spacing: 0.26em;
  color: #8ed8ff;
}
.cu-bp-title h1 {
  margin-top: 10px;
  font-family: 'Courier New', 'PingFang SC', monospace;
  font-size: 44px;
  line-height: 1;
  font-weight: 700;
  letter-spacing: -0.06em;
  word-break: break-word;
  color: #ffffff;
}
.cu-bp-rule {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 16px;
}
.cu-bp-rule span {
  height: 1px;
  flex: 1;
  background: rgba(170, 216, 242, 0.6);
}
.cu-bp-rule i {
  font-style: normal;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #8ed8ff;
}
.cu-bp-desc {
  margin-top: 14px;
  max-width: 40ch;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 12px;
  line-height: 1.8;
  color: #a9cde4;
}
.cu-bp-plan {
  margin-top: 26px;
}
.cu-bp-plan-head {
  display: flex;
  justify-content: space-between;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: #8ed8ff;
  border-bottom: 1px solid rgba(170, 216, 242, 0.35);
  padding-bottom: 5px;
}
.cu-bp-svg {
  margin-top: 12px;
  color: #9bd7fb;
  border: 1px solid rgba(170, 216, 242, 0.28);
  background: rgba(5, 26, 44, 0.45);
  padding: 6px;
}
.cu-bp-svg svg {
  display: block;
  width: 100%;
  height: auto;
}
.cu-bp-modules {
  margin-top: 26px;
}
.cu-bp-module-list {
  margin-top: 10px;
  border: 1px solid rgba(170, 216, 242, 0.32);
}
.cu-bp-module {
  position: relative;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 14px 12px;
  border-bottom: 1px solid rgba(170, 216, 242, 0.2);
  transition: background-color 0.2s ease;
}
.cu-bp-module:last-child {
  border-bottom: none;
}
.cu-bp-module:hover {
  background: rgba(142, 216, 255, 0.09);
}
.cu-bp-module-tick {
  height: 8px;
  width: 8px;
  border: 1px solid #8ed8ff;
  flex-shrink: 0;
}
.cu-bp-module-no {
  font-size: 8px;
  letter-spacing: 0.12em;
  color: #8ed8ff;
}
.cu-bp-module h2 {
  flex: 1;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 14px;
  font-weight: 500;
  letter-spacing: 0.02em;
  color: #eaf6ff;
}
.cu-bp-module-dim {
  font-size: 7px;
  letter-spacing: 0.08em;
  color: #7fb3d1;
}
.cu-bp-empty {
  margin-top: 10px;
  border: 1px dashed rgba(170, 216, 242, 0.4);
  padding: 16px;
  font-size: 10px;
  color: #9cc4dd;
}
.cu-bp-contact {
  margin-top: 26px;
  border: 1px solid rgba(170, 216, 242, 0.55);
}
.cu-bp-contact-row {
  display: flex;
  align-items: stretch;
  border-bottom: 1px solid rgba(170, 216, 242, 0.35);
}
.cu-bp-contact-row:last-child {
  border-bottom: none;
}
.cu-bp-contact-key {
  width: 92px;
  flex-shrink: 0;
  padding: 11px 9px;
  font-size: 8px;
  letter-spacing: 0.14em;
  color: #8ed8ff;
  border-right: 1px solid rgba(170, 216, 242, 0.3);
}
.cu-bp-contact-val {
  display: flex;
  align-items: center;
  gap: 7px;
  flex: 1;
  padding: 11px 10px;
  font-size: 12px;
  letter-spacing: 0.03em;
  color: #f2fbff;
  word-break: break-all;
  transition: background-color 0.2s ease;
}
a.cu-bp-contact-val:hover {
  background: rgba(142, 216, 255, 0.12);
}
.cu-bp-foot {
  margin-top: 16px;
  display: flex;
  justify-content: space-between;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #7fb3d1;
}
@keyframes cu-bp-in {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-bp-in {
  animation: cu-bp-in 0.8s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cu-bp-d1 {
  animation-delay: 0.08s;
}
.cu-bp-d2 {
  animation-delay: 0.18s;
}
.cu-bp-d3 {
  animation-delay: 0.28s;
}
.cu-bp-d4 {
  animation-delay: 0.38s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-bp-in {
    animation: none;
  }
  .cu-bp-module,
  .cu-bp-contact-val {
    transition: none;
  }
}
</style>

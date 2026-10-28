<template>
  <div class="cu-cm-root min-h-screen">
    <div class="cu-cm-wrap mx-auto w-full max-w-md">
      <header class="cu-cm-top cu-cm-in">
        <span class="cu-cm-office">STUDIO / 建筑事务所</span>
        <span class="cu-cm-year">2026</span>
      </header>

      <section class="cu-cm-title">
        <span class="cu-cm-figure cu-cm-in cu-cm-d1">01</span>
        <h1 class="cu-cm-name cu-cm-in cu-cm-d1">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-cm-desc cu-cm-in cu-cm-d2">{{ payload?.description }}</p>
      </section>

      <section class="cu-cm-model cu-cm-in cu-cm-d2">
        <img v-if="images.length" :src="images[0]" :alt="payload?.companyName || '项目图片'" class="cu-cm-model-img" />
        <div v-else class="cu-cm-model-art" aria-hidden="true">
          <span class="cu-cm-block cu-cm-block-a"></span>
          <span class="cu-cm-block cu-cm-block-b"></span>
          <span class="cu-cm-block cu-cm-block-c"></span>
          <span class="cu-cm-block cu-cm-block-d"></span>
        </div>
        <div class="cu-cm-model-caption">
          <span>MODEL / 1:100</span>
          <span>{{ payload?.companyName || '企业主页' }}</span>
        </div>
      </section>

      <section class="cu-cm-program cu-cm-in cu-cm-d3">
        <div class="cu-cm-program-head">
          <h2>PROGRAMME</h2>
          <span>服务项目 / {{ String(services.length).padStart(2, '0') }}</span>
        </div>
        <div v-if="services.length" class="cu-cm-grid">
          <article v-for="(service, i) in services" :key="i" class="cu-cm-cell">
            <span class="cu-cm-cell-index">0{{ i + 1 }}</span>
            <h3>{{ service.name }}</h3>
            <span class="cu-cm-cell-dim">{{ 18 + i * 6 }} m²</span>
          </article>
        </div>
      </section>

      <section class="cu-cm-contact cu-cm-in cu-cm-d4">
        <p class="cu-cm-contact-label">PROJECT ENQUIRY</p>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-cm-phone">
          <Phone :size="16" />
          <span>{{ payload.phone }}</span>
        </a>
        <div class="cu-cm-contact-foot">
          <span>{{ payload?.companyName || '企业主页' }}</span>
          <span>{{ nfcId ? 'NFC/' + nfcId : 'NFC/—' }}</span>
        </div>
      </section>
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
const images = computed(() => {
  const list = props.payload?.images
  return Array.isArray(list) ? list.filter((src) => typeof src === 'string' && src.trim()) : []
})
</script>

<style scoped>
.cu-cm-root {
  background-color: #d9d4c9;
  background-image:
    radial-gradient(circle at 18% 22%, rgba(255, 255, 255, 0.35), transparent 48%),
    radial-gradient(circle at 82% 72%, rgba(120, 108, 92, 0.16), transparent 44%);
  color: #2c2822;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
  overflow-x: hidden;
}
.cu-cm-wrap {
  padding: 22px 20px 30px;
}
.cu-cm-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(44, 40, 34, 0.32);
  padding-bottom: 8px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.18em;
  color: #6b6255;
}
.cu-cm-title {
  position: relative;
  padding-top: 34px;
}
.cu-cm-figure {
  position: absolute;
  right: 0;
  top: 18px;
  font-family: 'Times New Roman', serif;
  font-size: 58px;
  line-height: 1;
  color: rgba(168, 90, 61, 0.42);
}
.cu-cm-name {
  max-width: 7em;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  font-size: 42px;
  line-height: 1.06;
  font-weight: 500;
  letter-spacing: -0.02em;
  word-break: break-word;
}
.cu-cm-desc {
  margin-top: 16px;
  max-width: 32ch;
  font-size: 12px;
  line-height: 1.85;
  color: #6b6255;
}
.cu-cm-model {
  margin-top: 24px;
  border: 1px solid rgba(44, 40, 34, 0.35);
  background: #cdc7ba;
  padding: 8px 8px 0;
}
.cu-cm-model-img {
  display: block;
  width: 100%;
  aspect-ratio: 5 / 4;
  object-fit: cover;
  filter: saturate(0.72) contrast(1.05);
}
.cu-cm-model-art {
  position: relative;
  aspect-ratio: 5 / 4;
  overflow: hidden;
  background: linear-gradient(160deg, #c8c1b2, #b3aa99);
}
.cu-cm-block {
  position: absolute;
  background: #e2ddd2;
  border: 1px solid rgba(44, 40, 34, 0.28);
  box-shadow: 6px 8px 14px rgba(44, 40, 34, 0.14);
}
.cu-cm-block-a {
  left: 12%;
  bottom: 18%;
  width: 36%;
  height: 46%;
}
.cu-cm-block-b {
  left: 46%;
  bottom: 18%;
  width: 20%;
  height: 62%;
  background: #a85a3d;
  opacity: 0.85;
}
.cu-cm-block-c {
  left: 66%;
  bottom: 18%;
  width: 22%;
  height: 34%;
}
.cu-cm-block-d {
  left: 12%;
  bottom: 64%;
  width: 54%;
  height: 16%;
  background: #8e8779;
}
.cu-cm-model-art::after {
  content: '';
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(90deg, transparent 0 39px, rgba(255, 255, 255, 0.16) 39px 40px);
}
.cu-cm-model-caption {
  display: flex;
  justify-content: space-between;
  padding: 7px 2px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #6b6255;
}
.cu-cm-program {
  margin-top: 26px;
}
.cu-cm-program-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 1px solid rgba(44, 40, 34, 0.4);
  padding-bottom: 7px;
}
.cu-cm-program-head h2 {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.22em;
}
.cu-cm-program-head span {
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #6b6255;
}
.cu-cm-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  border-left: 1px solid rgba(44, 40, 34, 0.28);
}
.cu-cm-cell {
  position: relative;
  min-height: 108px;
  border-right: 1px solid rgba(44, 40, 34, 0.28);
  border-bottom: 1px solid rgba(44, 40, 34, 0.28);
  padding: 13px;
  transition: background-color 0.25s ease;
}
.cu-cm-cell:hover {
  background: rgba(168, 90, 61, 0.1);
}
.cu-cm-cell-index {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  color: #a85a3d;
}
.cu-cm-cell h3 {
  margin-top: 28px;
  font-size: 15px;
  font-weight: 600;
  line-height: 1.35;
}
.cu-cm-cell-dim {
  position: absolute;
  right: 10px;
  bottom: 8px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  color: #7d7466;
}
.cu-cm-empty {
  margin-top: 10px;
  border: 1px dashed rgba(44, 40, 34, 0.35);
  padding: 15px;
  font-size: 11px;
  color: #6b6255;
}
.cu-cm-contact {
  margin-top: 26px;
  background: #a85a3d;
  color: #f4eee3;
  padding: 17px 16px 13px;
}
.cu-cm-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #e7c7b7;
}
.cu-cm-phone {
  margin-top: 7px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 22px;
  letter-spacing: 0.01em;
  color: #fff8ed;
  transition: opacity 0.2s ease;
}
.cu-cm-phone:hover {
  opacity: 0.78;
}
.cu-cm-contact .cu-cm-empty {
  color: #f4eee3;
  border-color: rgba(244, 238, 227, 0.45);
}
.cu-cm-contact-foot {
  margin-top: 15px;
  border-top: 1px solid rgba(244, 238, 227, 0.32);
  padding-top: 9px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.12em;
  color: #e7c7b7;
}
@keyframes cu-cm-in {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-cm-in {
  animation: cu-cm-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-cm-d1 {
  animation-delay: 0.09s;
}
.cu-cm-d2 {
  animation-delay: 0.18s;
}
.cu-cm-d3 {
  animation-delay: 0.27s;
}
.cu-cm-d4 {
  animation-delay: 0.36s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-cm-in {
    animation: none;
  }
  .cu-cm-cell,
  .cu-cm-phone {
    transition: none;
  }
}
</style>

<template>
  <div class="cu-cl-root min-h-screen">
    <div class="cu-cl-wrap mx-auto w-full max-w-md">
      <header class="cu-cl-top cu-cl-in">
        <span>CLAY &amp; FORM</span>
        <span>HANDMADE / {{ nfcId ? String(nfcId).slice(-3) : '026' }}</span>
      </header>

      <section class="cu-cl-hero">
        <div class="cu-cl-arch cu-cl-in cu-cl-d1">
          <img v-if="images.length" :src="images[0]" :alt="payload?.companyName || '工作室作品'" />
          <div v-else class="cu-cl-arch-art">
            <span class="cu-cl-vase cu-cl-vase--back"></span>
            <span class="cu-cl-vase cu-cl-vase--front"></span>
            <span class="cu-cl-arch-label">{{ initial }}</span>
          </div>
        </div>
        <div class="cu-cl-copy">
          <p class="cu-cl-eyebrow cu-cl-in cu-cl-d1">陶土 · 手作 · 日常</p>
          <h1 class="cu-cl-name cu-cl-in cu-cl-d2">{{ payload?.companyName || '企业主页' }}</h1>
          <p class="cu-cl-desc cu-cl-in cu-cl-d2">{{ payload?.description }}</p>
        </div>
      </section>

      <section class="cu-cl-services cu-cl-in cu-cl-d3">
        <div class="cu-cl-services-head">
          <h2>工作室项目</h2>
          <span>STUDIO MENU</span>
        </div>
        <div v-if="services.length" class="cu-cl-cards">
          <article v-for="(service, i) in services" :key="i" class="cu-cl-card">
            <span class="cu-cl-glaze" :class="'cu-cl-glaze--' + (i % 3)"></span>
            <span class="cu-cl-card-no">0{{ i + 1 }}</span>
            <h3>{{ service.name }}</h3>
            <p>手作体验 · 小班预约</p>
            <span class="cu-cl-card-price">咨询价</span>
          </article>
        </div>
      </section>

      <section class="cu-cl-contact cu-cl-in cu-cl-d4">
        <div class="cu-cl-contact-copy">
          <p class="cu-cl-contact-label">来工作室坐坐</p>
          <p class="cu-cl-contact-note">拉坯、上釉、烧制，皆可预约。</p>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-cl-phone">
          <Phone :size="16" />
          <span>{{ payload.phone }}</span>
        </a>
      </section>

      <footer class="cu-cl-foot cu-cl-in cu-cl-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
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
const images = computed(() => {
  const list = props.payload?.images
  return Array.isArray(list) ? list.filter((src) => typeof src === 'string' && src.trim()) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-cl-root {
  background: #efe7d8;
  color: #3c2f26;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  overflow-x: hidden;
}
.cu-cl-root::before {
  content: '';
  position: fixed;
  inset: 0;
  pointer-events: none;
  background: radial-gradient(circle at 12% 8%, rgba(184, 107, 75, 0.12), transparent 38%);
}
.cu-cl-wrap {
  position: relative;
  padding: 22px 20px 32px;
}
.cu-cl-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(60, 47, 38, 0.22);
  padding-bottom: 9px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #8a6f5b;
}
.cu-cl-hero {
  margin-top: 22px;
}
.cu-cl-arch {
  position: relative;
  width: 78%;
  margin-left: auto;
  aspect-ratio: 3 / 4;
  border-radius: 190px 190px 18px 18px;
  overflow: hidden;
  background: #ddcdb8;
  border: 1px solid rgba(60, 47, 38, 0.16);
}
.cu-cl-arch img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  filter: saturate(0.78) contrast(0.98);
}
.cu-cl-arch-art {
  position: absolute;
  inset: 0;
  background: linear-gradient(165deg, #e7dac6, #cbb59b);
}
.cu-cl-vase {
  position: absolute;
  bottom: 12%;
  border-radius: 46% 46% 18% 18% / 36% 36% 14% 14%;
}
.cu-cl-vase--back {
  left: 15%;
  width: 29%;
  height: 41%;
  background: #a86a4d;
  opacity: 0.72;
}
.cu-cl-vase--front {
  right: 19%;
  width: 38%;
  height: 55%;
  background: #8b6b4f;
}
.cu-cl-vase--front::before {
  content: '';
  position: absolute;
  left: 26%;
  top: -7%;
  width: 48%;
  height: 12%;
  border-radius: 50%;
  background: #765a42;
}
.cu-cl-arch-label {
  position: absolute;
  left: 50%;
  top: 9%;
  transform: translateX(-50%);
  font-family: Georgia, 'Songti SC', serif;
  font-size: 58px;
  color: rgba(239, 231, 216, 0.72);
}
.cu-cl-copy {
  position: relative;
  margin-top: -34px;
  margin-right: 18%;
  background: #f7f0e4;
  border: 1px solid rgba(60, 47, 38, 0.16);
  padding: 18px 16px;
  box-shadow: 0 18px 34px -28px rgba(60, 47, 38, 0.7);
}
.cu-cl-eyebrow {
  font-size: 9px;
  letter-spacing: 0.26em;
  color: #a86a4d;
}
.cu-cl-name {
  margin-top: 9px;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  font-size: 33px;
  line-height: 1.18;
  font-weight: 600;
  letter-spacing: 0.01em;
  word-break: break-word;
}
.cu-cl-desc {
  margin-top: 10px;
  font-size: 11px;
  line-height: 1.8;
  color: #7d6858;
}
.cu-cl-services {
  margin-top: 28px;
}
.cu-cl-services-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 1px solid rgba(60, 47, 38, 0.28);
  padding-bottom: 7px;
}
.cu-cl-services-head h2 {
  font-family: 'Songti SC', serif;
  font-size: 17px;
  font-weight: 600;
}
.cu-cl-services-head span {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #9b8069;
}
.cu-cl-cards {
  margin-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 9px;
}
.cu-cl-card {
  position: relative;
  background: #f7f0e4;
  border: 1px solid rgba(60, 47, 38, 0.16);
  border-radius: 16px;
  padding: 14px 15px;
  display: flex;
  align-items: center;
  gap: 12px;
  transition: transform 0.25s ease;
}
.cu-cl-card:nth-child(2) {
  margin-left: 8px;
}
.cu-cl-card:nth-child(3) {
  margin-right: 8px;
}
.cu-cl-card:hover {
  transform: translateY(-2px) rotate(-0.4deg);
}
.cu-cl-glaze {
  flex-shrink: 0;
  height: 34px;
  width: 34px;
  border-radius: 50%;
}
.cu-cl-glaze--0 {
  background: radial-gradient(circle at 35% 30%, #c88f72, #8b4f36);
}
.cu-cl-glaze--1 {
  background: radial-gradient(circle at 35% 30%, #a9b39a, #62705a);
}
.cu-cl-glaze--2 {
  background: radial-gradient(circle at 35% 30%, #d9c7a5, #a28c63);
}
.cu-cl-card-no {
  position: absolute;
  right: 13px;
  top: 11px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  color: #b29a83;
}
.cu-cl-card h3 {
  flex: 1;
  font-size: 15px;
  font-weight: 600;
  line-height: 1.35;
}
.cu-cl-card p {
  display: none;
}
.cu-cl-card-price {
  flex-shrink: 0;
  border: 1px solid rgba(168, 106, 77, 0.45);
  border-radius: 999px;
  padding: 3px 8px;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #a86a4d;
}
.cu-cl-empty {
  margin-top: 10px;
  padding: 14px;
  border: 1px dashed rgba(60, 47, 38, 0.3);
  border-radius: 14px;
  font-size: 11px;
  color: #8a6f5b;
}
.cu-cl-contact {
  margin-top: 26px;
  background: #8b6b4f;
  color: #f7f0e4;
  border-radius: 20px 20px 6px 6px;
  padding: 18px 16px;
}
.cu-cl-contact-label {
  font-family: 'Songti SC', serif;
  font-size: 17px;
  letter-spacing: 0.1em;
}
.cu-cl-contact-note {
  margin-top: 4px;
  font-size: 10px;
  line-height: 1.7;
  color: #e6d5c2;
}
.cu-cl-phone {
  margin-top: 13px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: #f7f0e4;
  color: #3c2f26;
  border-radius: 999px;
  padding: 9px 15px;
  font-size: 17px;
  transition: transform 0.18s ease;
}
.cu-cl-phone:active {
  transform: scale(0.97);
}
.cu-cl-contact .cu-cl-empty {
  color: #f0e2d1;
  border-color: rgba(247, 240, 228, 0.4);
}
.cu-cl-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #9b8069;
}
@keyframes cu-cl-in {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-cl-in {
  animation: cu-cl-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-cl-d1 {
  animation-delay: 0.08s;
}
.cu-cl-d2 {
  animation-delay: 0.18s;
}
.cu-cl-d3 {
  animation-delay: 0.27s;
}
.cu-cl-d4 {
  animation-delay: 0.36s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-cl-in {
    animation: none;
  }
  .cu-cl-card,
  .cu-cl-phone {
    transition: none;
  }
}
</style>

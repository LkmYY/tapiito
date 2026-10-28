<template>
  <div class="cu-fp-root min-h-screen">
    <svg class="cu-fp-botanical" viewBox="0 0 200 260" fill="none" aria-hidden="true">
      <path d="M100 250C100 190 96 130 78 74" stroke="currentColor" stroke-width="1.1" />
      <path d="M90 186c-26-6-42-24-46-46 25 2 42 18 46 46Z" stroke="currentColor" stroke-width="1" />
      <path d="M96 150c22-8 34-25 36-45-23 4-36 20-36 45Z" stroke="currentColor" stroke-width="1" />
      <path d="M84 112c-20-4-32-18-35-35 19 1 32 14 35 35Z" stroke="currentColor" stroke-width="1" />
      <path d="M80 72c-3-22 7-40 25-51 5 20-4 39-25 51Z" stroke="currentColor" stroke-width="1" />
      <circle cx="104" cy="20" r="7" stroke="currentColor" stroke-width="1" />
      <circle cx="92" cy="31" r="5" stroke="currentColor" stroke-width="1" />
      <circle cx="117" cy="33" r="5" stroke="currentColor" stroke-width="1" />
    </svg>

    <div class="cu-fp-wrap mx-auto w-full max-w-md">
      <header class="cu-fp-top cu-fp-in">
        <span>拾光植物事务所</span>
        <span>FLORA / {{ nfcId ? String(nfcId).slice(-4) : '2026' }}</span>
      </header>

      <section class="cu-fp-hero">
        <div class="cu-fp-hero-main">
          <p class="cu-fp-eyebrow cu-fp-in cu-fp-d1">BOTANICAL ATELIER</p>
          <h1 class="cu-fp-name cu-fp-in cu-fp-d1">{{ payload?.companyName || '企业主页' }}</h1>
          <p class="cu-fp-desc cu-fp-in cu-fp-d2">{{ payload?.description }}</p>
        </div>
        <span class="cu-fp-vertical cu-fp-in cu-fp-d2">花 · 叶 · 枝 · 果</span>
      </section>

      <div class="cu-fp-divider cu-fp-in cu-fp-d2">
        <span></span>
        <i>✿</i>
        <span></span>
      </div>

      <div v-if="images.length" class="cu-fp-photo cu-fp-in cu-fp-d3">
        <img :src="images[0]" :alt="payload?.companyName || '花艺作品'" />
        <span class="cu-fp-photo-cap">SEASONAL WORK / 当季作品</span>
      </div>

      <section class="cu-fp-services cu-fp-in cu-fp-d3">
        <p class="cu-fp-section-title">WHAT WE GROW <span>服务项目</span></p>
        <div v-if="services.length" class="cu-fp-grid">
          <article v-for="(service, i) in services" :key="i" class="cu-fp-pack">
            <div class="cu-fp-pack-top">
              <component :is="petalIcons[i % petalIcons.length]" :size="17" />
              <span>SEED {{ String(i + 1).padStart(2, '0') }}</span>
            </div>
            <h2>{{ service.name }}</h2>
            <p>当季花材 · 手工制作 · 可预约</p>
            <span class="cu-fp-pack-line"></span>
          </article>
        </div>
      </section>

      <section class="cu-fp-contact cu-fp-in cu-fp-d4">
        <div>
          <p class="cu-fp-contact-label">预约一束花 / ORDER A BUNCH</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-fp-phone">
            <Phone :size="15" />
            <span>{{ payload.phone }}</span>
            <ArrowUpRight :size="14" />
          </a>
        </div>
        <span class="cu-fp-seal">{{ initial }}</span>
      </section>

      <footer class="cu-fp-foot cu-fp-in cu-fp-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Flower2, Leaf, Phone, Sprout } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const petalIcons = [Flower2, Leaf, Sprout]

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
.cu-fp-root {
  position: relative;
  overflow: hidden;
  background: #eef1e7;
  color: #2f3a2c;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
}
.cu-fp-botanical {
  position: absolute;
  right: -34px;
  top: 54px;
  width: 190px;
  color: rgba(89, 112, 82, 0.28);
  pointer-events: none;
}
.cu-fp-wrap {
  position: relative;
  padding: 22px 22px 32px;
}
.cu-fp-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(47, 58, 44, 0.22);
  padding-bottom: 9px;
  font-size: 9px;
  letter-spacing: 0.16em;
  color: #6d7a66;
}
.cu-fp-top span:last-child {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
}
.cu-fp-hero {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  margin-top: 28px;
}
.cu-fp-hero-main {
  flex: 1;
  min-width: 0;
}
.cu-fp-eyebrow {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.28em;
  color: #a2605f;
}
.cu-fp-name {
  margin-top: 10px;
  font-size: 42px;
  line-height: 1.08;
  font-weight: 500;
  letter-spacing: 0.02em;
  color: #2f3a2c;
  word-break: break-word;
}
.cu-fp-desc {
  margin-top: 14px;
  max-width: 27ch;
  font-size: 13px;
  line-height: 1.9;
  font-style: italic;
  color: #61705b;
}
.cu-fp-vertical {
  writing-mode: vertical-rl;
  margin-top: 5px;
  font-size: 10px;
  letter-spacing: 0.3em;
  color: #8a987f;
}
.cu-fp-divider {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 24px;
  color: #a2605f;
}
.cu-fp-divider span {
  height: 1px;
  flex: 1;
  background: rgba(162, 96, 95, 0.3);
}
.cu-fp-divider i {
  font-style: normal;
  font-size: 11px;
}
.cu-fp-photo {
  position: relative;
  margin-top: 20px;
  padding: 7px;
  background: #f7f5ee;
  border: 1px solid rgba(47, 58, 44, 0.18);
  box-shadow: 0 16px 30px -24px rgba(47, 58, 44, 0.6);
}
.cu-fp-photo img {
  display: block;
  width: 100%;
  aspect-ratio: 4 / 3;
  object-fit: cover;
  filter: saturate(0.78) sepia(0.08);
}
.cu-fp-photo-cap {
  display: block;
  padding-top: 7px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #7b8873;
}
.cu-fp-services {
  margin-top: 26px;
}
.cu-fp-section-title {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.22em;
  color: #6d7a66;
  border-bottom: 1px solid rgba(47, 58, 44, 0.22);
  padding-bottom: 7px;
}
.cu-fp-section-title span {
  font-family: 'Songti SC', serif;
  font-size: 10px;
  letter-spacing: 0.1em;
  color: #a2605f;
}
.cu-fp-grid {
  margin-top: 14px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.cu-fp-pack {
  position: relative;
  background: #f8f6ef;
  border: 1px solid rgba(47, 58, 44, 0.2);
  padding: 14px 14px 12px;
  transition: transform 0.25s ease, box-shadow 0.25s ease;
}
.cu-fp-pack:hover {
  transform: translateY(-2px) rotate(-0.3deg);
  box-shadow: 0 14px 24px -18px rgba(47, 58, 44, 0.5);
}
.cu-fp-pack::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 4px;
  background: repeating-linear-gradient(180deg, #b48b7f 0 4px, transparent 4px 8px);
}
.cu-fp-pack-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: #a2605f;
}
.cu-fp-pack-top span {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
}
.cu-fp-pack h2 {
  margin-top: 9px;
  font-size: 17px;
  font-weight: 500;
  letter-spacing: 0.04em;
}
.cu-fp-pack p {
  margin-top: 4px;
  font-size: 10px;
  letter-spacing: 0.06em;
  color: #7b8873;
}
.cu-fp-pack-line {
  display: block;
  margin-top: 10px;
  height: 1px;
  width: 32px;
  background: rgba(162, 96, 95, 0.45);
}
.cu-fp-empty {
  margin-top: 10px;
  padding: 14px;
  border: 1px dashed rgba(47, 58, 44, 0.28);
  font-size: 11px;
  color: #7b8873;
}
.cu-fp-contact {
  margin-top: 26px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 14px;
  border-top: 1px solid rgba(47, 58, 44, 0.22);
  border-bottom: 1px solid rgba(47, 58, 44, 0.22);
  padding: 16px 0;
}
.cu-fp-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #7b8873;
}
.cu-fp-phone {
  margin-top: 7px;
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 18px;
  color: #2f3a2c;
  border-bottom: 1px solid rgba(162, 96, 95, 0.5);
  padding-bottom: 3px;
  transition: color 0.2s ease;
}
.cu-fp-phone:hover {
  color: #a2605f;
}
.cu-fp-seal {
  flex-shrink: 0;
  height: 54px;
  width: 54px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(162, 96, 95, 0.55);
  border-radius: 50%;
  color: #a2605f;
  font-size: 24px;
  transform: rotate(-8deg);
}
.cu-fp-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.16em;
  color: #7b8873;
}
@keyframes cu-fp-in {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-fp-in {
  animation: cu-fp-in 0.85s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-fp-d1 {
  animation-delay: 0.08s;
}
.cu-fp-d2 {
  animation-delay: 0.18s;
}
.cu-fp-d3 {
  animation-delay: 0.27s;
}
.cu-fp-d4 {
  animation-delay: 0.36s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-fp-in {
    animation: none;
  }
  .cu-fp-pack,
  .cu-fp-phone {
    transition: none;
  }
}
</style>

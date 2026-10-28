<template>
  <div class="cu-ed-root min-h-screen">
    <div class="cu-ed-rule"></div>
    <div class="cu-ed-wrap mx-auto w-full max-w-md">
      <header class="cu-ed-masthead cu-ed-rise">
        <span>拾光档案 / PROFILE</span>
        <span>ISSUE {{ nfcId ? String(nfcId).slice(-3) : '026' }}</span>
      </header>

      <section class="cu-ed-title">
        <p class="cu-ed-kicker cu-ed-rise cu-ed-d1">独立企业档案 · 二〇二六</p>
        <h1 class="cu-ed-name cu-ed-rise cu-ed-d2">{{ payload?.companyName || '企业主页' }}</h1>
        <div class="cu-ed-standfirst cu-ed-rise cu-ed-d3">
          <p>{{ payload?.description }}</p>
        </div>
      </section>

      <section class="cu-ed-cover cu-ed-rise cu-ed-d3">
        <img v-if="images.length" :src="images[0]" :alt="payload?.companyName || '企业图片'" class="cu-ed-cover-img" />
        <div v-else class="cu-ed-cover-art">
          <span class="cu-ed-cover-initial">{{ initial }}</span>
          <span class="cu-ed-cover-caption">TYPE / IMAGE PLACEHOLDER</span>
        </div>
        <div class="cu-ed-cover-meta">
          <span>PORTRAIT OF A PRACTICE</span>
          <span>NO. {{ nfcId || '—' }}</span>
        </div>
      </section>

      <section class="cu-ed-index cu-ed-rise cu-ed-d4">
        <div class="cu-ed-section-head">
          <h2>服务索引</h2>
          <span>INDEX / {{ String(services.length).padStart(2, '0') }}</span>
        </div>

        <div v-if="services.length" class="cu-ed-list">
          <div v-for="(service, i) in services" :key="i" class="cu-ed-entry">
            <span class="cu-ed-entry-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="cu-ed-entry-name">{{ service.name }}</span>
            <ArrowUpRight class="cu-ed-entry-arrow" :size="16" />
          </div>
        </div>
      </section>

      <section class="cu-ed-contact cu-ed-rise cu-ed-d5">
        <p class="cu-ed-contact-label">预约与咨询 / ENQUIRIES</p>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-ed-phone">
          <Phone :size="17" />
          <span>{{ payload.phone }}</span>
        </a>
        <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-ed-mail">{{ payload.email }}</a>
      </section>

      <footer class="cu-ed-foot cu-ed-rise cu-ed-d5">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC / ' + nfcId : 'NFC / —' }}</span>
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
const images = computed(() => {
  const list = props.payload?.images
  return Array.isArray(list) ? list.filter((src) => typeof src === 'string' && src.trim()) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-ed-root {
  background: #f3eee3;
  color: #241d17;
  font-family: Georgia, 'Songti SC', 'Noto Serif SC', serif;
  overflow-x: hidden;
}
.cu-ed-rule {
  height: 5px;
  background: #7a1f1f;
}
.cu-ed-wrap {
  padding: 26px 24px 40px;
}
.cu-ed-masthead {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(36, 29, 23, 0.28);
  padding-bottom: 10px;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.22em;
  text-transform: uppercase;
  color: #6c5c4b;
}
.cu-ed-title {
  padding-top: 30px;
}
.cu-ed-kicker {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.28em;
  text-transform: uppercase;
  color: #9c2b21;
}
.cu-ed-name {
  margin-top: 12px;
  font-size: 54px;
  line-height: 0.94;
  font-weight: 400;
  letter-spacing: -0.045em;
  color: #1f1710;
  word-break: break-word;
}
.cu-ed-standfirst {
  margin-top: 22px;
  border-left: 2px solid #9c2b21;
  padding-left: 14px;
  max-width: 30ch;
}
.cu-ed-standfirst p {
  font-size: 14px;
  line-height: 1.75;
  font-style: italic;
  color: #4d4035;
}
.cu-ed-cover {
  margin-top: 26px;
  position: relative;
}
.cu-ed-cover-img,
.cu-ed-cover-art {
  width: 100%;
  aspect-ratio: 4 / 5;
  object-fit: cover;
  display: block;
  filter: saturate(0.82) contrast(1.02);
}
.cu-ed-cover-art {
  position: relative;
  overflow: hidden;
  background:
    radial-gradient(circle at 72% 24%, rgba(156, 43, 33, 0.22), transparent 46%),
    linear-gradient(155deg, #ded2bd, #b9a68b 58%, #8f7c62);
  display: flex;
  align-items: flex-end;
  padding: 20px;
}
.cu-ed-cover-art::before {
  content: '';
  position: absolute;
  inset: 14px;
  border: 1px solid rgba(243, 238, 227, 0.42);
}
.cu-ed-cover-initial {
  font-size: 190px;
  line-height: 0.72;
  color: rgba(243, 238, 227, 0.9);
  font-weight: 400;
  letter-spacing: -0.06em;
}
.cu-ed-cover-caption {
  position: absolute;
  top: 22px;
  right: 22px;
  writing-mode: vertical-rl;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(243, 238, 227, 0.78);
}
.cu-ed-cover-meta {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(36, 29, 23, 0.22);
  padding: 9px 0;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #7a6a58;
}
.cu-ed-index {
  margin-top: 34px;
}
.cu-ed-section-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 2px solid #241d17;
  padding-bottom: 8px;
}
.cu-ed-section-head h2 {
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 0.08em;
}
.cu-ed-section-head span {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #8a7864;
}
.cu-ed-entry {
  display: flex;
  align-items: baseline;
  gap: 14px;
  padding: 15px 0;
  border-bottom: 1px solid rgba(36, 29, 23, 0.16);
  transition: padding-left 0.25s ease, color 0.25s ease;
}
.cu-ed-entry:hover {
  padding-left: 6px;
  color: #9c2b21;
}
.cu-ed-entry-no {
  font-family: 'Courier New', monospace;
  font-size: 10px;
  color: #9c2b21;
  width: 22px;
}
.cu-ed-entry-name {
  flex: 1;
  font-size: 19px;
  letter-spacing: 0.01em;
}
.cu-ed-entry-arrow {
  color: #a08d78;
  flex-shrink: 0;
}
.cu-ed-empty,
.cu-ed-muted {
  padding: 18px 0;
  font-size: 13px;
  color: #8a7864;
}
.cu-ed-contact {
  margin-top: 34px;
  background: #241d17;
  color: #f3eee3;
  padding: 20px;
}
.cu-ed-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: #c7a98c;
}
.cu-ed-phone {
  margin-top: 10px;
  display: inline-flex;
  align-items: center;
  gap: 9px;
  font-size: 20px;
  letter-spacing: 0.03em;
  color: #f3eee3;
  border-bottom: 1px solid #9c2b21;
  padding-bottom: 4px;
  transition: color 0.2s ease, border-color 0.2s ease;
}
.cu-ed-phone:hover {
  color: #e7b7a4;
  border-color: #e7b7a4;
}
.cu-ed-mail {
  display: block;
  margin-top: 12px;
  font-family: 'Courier New', monospace;
  font-size: 11px;
  color: #c7a98c;
  word-break: break-all;
}
.cu-ed-foot {
  margin-top: 26px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.18em;
  text-transform: uppercase;
  color: #8a7864;
}
@keyframes cu-ed-rise {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-ed-rise {
  animation: cu-ed-rise 0.75s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-ed-d1 {
  animation-delay: 0.08s;
}
.cu-ed-d2 {
  animation-delay: 0.14s;
}
.cu-ed-d3 {
  animation-delay: 0.22s;
}
.cu-ed-d4 {
  animation-delay: 0.3s;
}
.cu-ed-d5 {
  animation-delay: 0.38s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-ed-rise {
    animation: none;
  }
  .cu-ed-entry,
  .cu-ed-phone {
    transition: none;
  }
}
</style>

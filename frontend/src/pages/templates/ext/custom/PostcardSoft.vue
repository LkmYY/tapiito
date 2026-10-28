<template>
  <div class="cu-ps-root min-h-screen">
    <div class="cu-ps-wrap mx-auto w-full max-w-md">
      <header class="cu-ps-top cu-ps-in">
        <span>POSTCARD FROM</span>
        <span>{{ nfcId ? 'NO. ' + nfcId : 'NO. —' }}</span>
      </header>

      <section class="cu-ps-card cu-ps-in cu-ps-d1">
        <div class="cu-ps-stamp" aria-hidden="true">
          <span class="cu-ps-stamp-inner">{{ initial }}</span>
          <i>26</i>
        </div>

        <div class="cu-ps-postmark" aria-hidden="true">
          <span>SHIGUANG</span>
          <b>2026 · 10 · 20</b>
        </div>

        <p class="cu-ps-greeting cu-ps-in cu-ps-d1">你好，很高兴认识你 ——</p>
        <h1 class="cu-ps-name cu-ps-in cu-ps-d2">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-ps-desc cu-ps-in cu-ps-d2">{{ payload?.description }}</p>

        <div class="cu-ps-note cu-ps-in cu-ps-d3">
          <span class="cu-ps-note-label">A LITTLE NOTE</span>
          <p>我们相信，好的合作从一次轻松的聊天开始。欢迎随时来信或来电。</p>
        </div>

        <div class="cu-ps-services cu-ps-in cu-ps-d3">
          <p class="cu-ps-section-label">WE CAN HELP WITH</p>
          <div v-if="services.length" class="cu-ps-checks">
            <div v-for="(service, i) in services" :key="i" class="cu-ps-check">
              <span class="cu-ps-check-box">✓</span>
              <span>{{ service.name }}</span>
            </div>
          </div>
        </div>

        <div class="cu-ps-divider"></div>

        <section class="cu-ps-contact cu-ps-in cu-ps-d4">
          <div>
            <p class="cu-ps-contact-label">CALL ME BACK</p>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-ps-phone">
              <Phone :size="15" />
              <span>{{ payload.phone }}</span>
            </a>
            <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-ps-mail">{{ payload.email }}</a>
          </div>
          <div class="cu-ps-sign">
            <span>with love,</span>
            <p>{{ payload?.companyName || '企业主页' }}</p>
          </div>
        </section>
      </section>

      <div v-if="images.length" class="cu-ps-photo cu-ps-in cu-ps-d4">
        <img :src="images[0]" :alt="payload?.companyName || '明信片图片'" />
        <span>WISH YOU WERE HERE</span>
      </div>

      <footer class="cu-ps-foot cu-ps-in cu-ps-d4">
        <span>POSTED WITH CARE</span>
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
.cu-ps-root {
  background: #f7e7df;
  color: #4a3b3e;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  overflow-x: hidden;
}
.cu-ps-root::before {
  content: '';
  position: fixed;
  inset: 0;
  pointer-events: none;
  background:
    radial-gradient(circle at 86% 10%, rgba(173, 212, 197, 0.4), transparent 35%),
    radial-gradient(circle at 8% 82%, rgba(247, 202, 189, 0.55), transparent 34%);
}
.cu-ps-wrap {
  position: relative;
  padding: 22px 18px 30px;
}
.cu-ps-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(74, 59, 62, 0.24);
  padding-bottom: 8px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #9a7d78;
}
.cu-ps-card {
  position: relative;
  margin-top: 22px;
  background: #fffaf4;
  border: 1px solid rgba(74, 59, 62, 0.16);
  padding: 22px 18px 18px;
  box-shadow: 0 20px 40px -30px rgba(120, 72, 68, 0.7);
}
.cu-ps-card::before {
  content: '';
  position: absolute;
  inset: 6px;
  border: 1px dashed rgba(154, 125, 120, 0.5);
  pointer-events: none;
}
.cu-ps-stamp {
  position: absolute;
  right: 15px;
  top: 15px;
  z-index: 2;
  height: 52px;
  width: 43px;
  background: #f4cbc0;
  border: 1px dashed #b5736c;
  display: flex;
  align-items: center;
  justify-content: center;
}
.cu-ps-stamp::before {
  content: '';
  position: absolute;
  inset: 3px;
  border: 1px solid rgba(181, 115, 108, 0.5);
}
.cu-ps-stamp-inner {
  font-size: 19px;
  color: #a25b56;
}
.cu-ps-stamp i {
  position: absolute;
  right: 2px;
  bottom: 0;
  font-style: normal;
  font-family: 'Courier New', monospace;
  font-size: 6px;
  color: #a25b56;
}
.cu-ps-postmark {
  position: absolute;
  right: 67px;
  top: 19px;
  z-index: 2;
  height: 42px;
  width: 42px;
  border: 1px solid rgba(154, 125, 120, 0.45);
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  transform: rotate(-13deg);
}
.cu-ps-postmark span {
  font-family: 'Courier New', monospace;
  font-size: 4px;
  letter-spacing: 0.12em;
  color: #a5847e;
}
.cu-ps-postmark b {
  font-family: 'Courier New', monospace;
  font-size: 4px;
  font-weight: 400;
  color: #a5847e;
  margin-top: 1px;
}
.cu-ps-greeting {
  position: relative;
  max-width: 16ch;
  font-size: 12px;
  letter-spacing: 0.06em;
  color: #a25b56;
}
.cu-ps-name {
  position: relative;
  margin-top: 9px;
  max-width: 6em;
  font-size: 38px;
  line-height: 1.14;
  font-weight: 600;
  letter-spacing: 0.01em;
  color: #403236;
  word-break: break-word;
}
.cu-ps-desc {
  position: relative;
  margin-top: 11px;
  max-width: 31ch;
  font-size: 12px;
  line-height: 1.9;
  color: #7e6662;
}
.cu-ps-note {
  position: relative;
  margin-top: 18px;
  background: #f3f7f2;
  border-left: 2px solid #a8c9bb;
  padding: 11px 12px;
}
.cu-ps-note-label {
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: #7d9d8f;
}
.cu-ps-note p {
  margin-top: 5px;
  font-size: 10px;
  line-height: 1.8;
  color: #64786f;
}
.cu-ps-services {
  position: relative;
  margin-top: 19px;
}
.cu-ps-section-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #a5847e;
  border-bottom: 1px solid rgba(154, 125, 120, 0.3);
  padding-bottom: 5px;
}
.cu-ps-checks {
  margin-top: 7px;
}
.cu-ps-check {
  display: flex;
  align-items: center;
  gap: 9px;
  padding: 7px 0;
  border-bottom: 1px dashed rgba(154, 125, 120, 0.28);
  font-size: 13px;
  color: #4a3b3e;
}
.cu-ps-check-box {
  display: inline-flex;
  height: 15px;
  width: 15px;
  align-items: center;
  justify-content: center;
  border: 1px solid #b5736c;
  border-radius: 2px;
  font-family: Arial, sans-serif;
  font-size: 9px;
  color: #a25b56;
  flex-shrink: 0;
}
.cu-ps-empty {
  padding: 11px 0;
  font-size: 10px;
  color: #a5847e;
}
.cu-ps-divider {
  margin-top: 15px;
  border-top: 1px dashed rgba(154, 125, 120, 0.5);
}
.cu-ps-contact {
  position: relative;
  margin-top: 13px;
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 10px;
}
.cu-ps-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: #a5847e;
}
.cu-ps-phone {
  margin-top: 6px;
  display: inline-flex;
  align-items: center;
  gap: 6px;
  color: #403236;
  font-size: 17px;
  letter-spacing: 0.02em;
  border-bottom: 1px solid rgba(165, 91, 86, 0.5);
  padding-bottom: 2px;
  transition: color 0.2s ease;
}
.cu-ps-phone:hover {
  color: #a25b56;
}
.cu-ps-mail {
  display: block;
  margin-top: 6px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  color: #9a7d78;
  word-break: break-all;
}
.cu-ps-sign {
  text-align: right;
  flex-shrink: 0;
}
.cu-ps-sign span {
  font-size: 8px;
  font-style: italic;
  color: #a5847e;
}
.cu-ps-sign p {
  margin-top: 2px;
  font-size: 11px;
  color: #4a3b3e;
  max-width: 10ch;
}
.cu-ps-photo {
  position: relative;
  margin-top: 18px;
  border: 5px solid #fffaf4;
  box-shadow: 0 14px 28px -22px rgba(120, 72, 68, 0.8);
  transform: rotate(-1.2deg);
}
.cu-ps-photo img {
  display: block;
  width: 100%;
  aspect-ratio: 4 / 3;
  object-fit: cover;
  filter: saturate(0.82);
}
.cu-ps-photo span {
  position: absolute;
  left: 8px;
  bottom: 8px;
  background: rgba(255, 250, 244, 0.88);
  padding: 2px 5px;
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: #a25b56;
}
.cu-ps-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #a5847e;
}
@keyframes cu-ps-in {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-ps-in {
  animation: cu-ps-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-ps-d1 {
  animation-delay: 0.08s;
}
.cu-ps-d2 {
  animation-delay: 0.17s;
}
.cu-ps-d3 {
  animation-delay: 0.26s;
}
.cu-ps-d4 {
  animation-delay: 0.35s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-ps-in {
    animation: none;
  }
  .cu-ps-phone {
    transition: none;
  }
}
</style>

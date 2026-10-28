<template>
  <div class="cu-dn-root min-h-screen">
    <div class="cu-dn-wrap mx-auto w-full max-w-md">
      <header class="cu-dn-masthead cu-dn-in">
        <div class="cu-dn-edition">MORNING EDITION</div>
        <div class="cu-dn-title">THE DAILY</div>
        <div class="cu-dn-meta">
          <span>2026 年 10 月 20 日</span>
          <span>第 {{ nfcId ? String(nfcId).slice(-3) : '026' }} 期</span>
          <span>今日企业刊</span>
        </div>
      </header>

      <article class="cu-dn-article">
        <p class="cu-dn-kicker cu-dn-in cu-dn-d1">SPECIAL REPORT · 企业服务</p>
        <h1 class="cu-dn-headline cu-dn-in cu-dn-d1">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-dn-deck cu-dn-in cu-dn-d2">{{ payload?.description }}</p>

        <div class="cu-dn-byline cu-dn-in cu-dn-d2">
          <span>本报讯</span>
          <span>记者 / 编辑部</span>
          <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
        </div>

        <div class="cu-dn-body cu-dn-in cu-dn-d3">
          <div v-if="images.length" class="cu-dn-figure">
            <img :src="images[0]" :alt="payload?.companyName || '企业图片'" />
            <p>▲ 企业资料图 / 本报摄</p>
          </div>
          <div v-else class="cu-dn-figure cu-dn-figure--blank">
            <span>{{ initial }}</span>
            <p>▲ 企业资料图 / 本报资料室</p>
          </div>
        </div>

        <section class="cu-dn-services cu-dn-in cu-dn-d3">
          <h2 class="cu-dn-section-title">服务分类广告</h2>
          <div v-if="services.length" class="cu-dn-classifieds">
            <article v-for="(service, i) in services" :key="i" class="cu-dn-ad">
              <span class="cu-dn-ad-no">{{ String(i + 1).padStart(2, '0') }}</span>
              <h3>{{ service.name }}</h3>
              <p>专业承接 · 欢迎来电</p>
            </article>
          </div>
        </section>
      </article>

      <section class="cu-dn-contact cu-dn-in cu-dn-d4">
        <div class="cu-dn-contact-box">
          <span class="cu-dn-contact-label">ADVERTISING / 咨询热线</span>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-dn-phone">
            <Phone :size="15" />
            {{ payload.phone }}
          </a>
        </div>
        <p class="cu-dn-smallprint">本报所载信息以企业实际提供为准 · 欢迎预约到访</p>
      </section>

      <footer class="cu-dn-foot cu-dn-in cu-dn-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>DAILY / 2026</span>
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
.cu-dn-root {
  background: #eee9dc;
  color: #171512;
  font-family: Georgia, 'Songti SC', 'Noto Serif SC', serif;
  overflow-x: hidden;
}
.cu-dn-wrap {
  padding: 18px 16px 28px;
}
.cu-dn-masthead {
  border-top: 4px solid #171512;
  border-bottom: 1px solid #171512;
  padding: 8px 0 7px;
  text-align: center;
}
.cu-dn-edition {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.3em;
  color: #6f695d;
}
.cu-dn-title {
  margin-top: 2px;
  font-family: 'Times New Roman', Georgia, serif;
  font-size: 46px;
  line-height: 1;
  font-weight: 700;
  letter-spacing: 0.02em;
  transform: scaleX(0.94);
}
.cu-dn-meta {
  margin-top: 5px;
  display: flex;
  justify-content: center;
  gap: 9px;
  font-size: 8px;
  color: #524d43;
}
.cu-dn-meta span + span {
  border-left: 1px solid #a9a294;
  padding-left: 9px;
}
.cu-dn-article {
  margin-top: 18px;
}
.cu-dn-kicker {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #8b2f2f;
}
.cu-dn-headline {
  margin-top: 8px;
  font-size: 45px;
  line-height: 0.98;
  font-weight: 700;
  letter-spacing: -0.035em;
  word-break: break-word;
}
.cu-dn-deck {
  margin-top: 12px;
  border-top: 1px solid #171512;
  border-bottom: 1px solid #171512;
  padding: 9px 0;
  font-size: 12px;
  line-height: 1.75;
  font-style: italic;
  color: #49443a;
}
.cu-dn-byline {
  display: flex;
  justify-content: space-between;
  gap: 6px;
  margin-top: 7px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.1em;
  color: #6f695d;
}
.cu-dn-body {
  margin-top: 15px;
}
.cu-dn-figure img {
  display: block;
  width: 100%;
  aspect-ratio: 16 / 9;
  object-fit: cover;
  filter: grayscale(1) contrast(1.2);
}
.cu-dn-figure--blank {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  aspect-ratio: 16 / 9;
  border: 1px solid #171512;
  background-image: radial-gradient(rgba(23, 21, 18, 0.22) 0.7px, transparent 0.7px);
  background-size: 4px 4px;
}
.cu-dn-figure--blank span {
  font-size: 82px;
  line-height: 1;
  color: rgba(23, 21, 18, 0.18);
}
.cu-dn-figure p {
  margin-top: 4px;
  font-size: 8px;
  letter-spacing: 0.04em;
  color: #6f695d;
}
.cu-dn-services {
  margin-top: 21px;
}
.cu-dn-section-title {
  border-top: 3px solid #171512;
  border-bottom: 1px solid #171512;
  padding: 5px 0;
  font-size: 12px;
  letter-spacing: 0.14em;
}
.cu-dn-classifieds {
  display: grid;
  grid-template-columns: 1fr 1fr;
  border-left: 1px solid #171512;
}
.cu-dn-ad {
  border-right: 1px solid #171512;
  border-bottom: 1px solid #171512;
  padding: 11px 10px;
  transition: background-color 0.2s ease;
}
.cu-dn-ad:hover {
  background: rgba(23, 21, 18, 0.05);
}
.cu-dn-ad-no {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  color: #8b2f2f;
}
.cu-dn-ad h3 {
  margin-top: 12px;
  font-size: 16px;
  font-weight: 700;
  line-height: 1.3;
}
.cu-dn-ad p {
  margin-top: 6px;
  font-size: 8px;
  letter-spacing: 0.04em;
  color: #6f695d;
}
.cu-dn-empty {
  border: 1px solid #171512;
  border-top: none;
  padding: 13px;
  font-size: 11px;
  color: #6f695d;
}
.cu-dn-contact {
  margin-top: 22px;
}
.cu-dn-contact-box {
  border: 2px solid #171512;
  padding: 14px;
  text-align: center;
}
.cu-dn-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #6f695d;
}
.cu-dn-phone {
  margin-top: 8px;
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 24px;
  font-weight: 700;
  letter-spacing: 0.01em;
  color: #171512;
  border-bottom: 2px solid #8b2f2f;
  padding-bottom: 3px;
  transition: color 0.2s ease;
}
.cu-dn-phone:hover {
  color: #8b2f2f;
}
.cu-dn-smallprint {
  margin-top: 7px;
  text-align: center;
  font-size: 8px;
  color: #6f695d;
}
.cu-dn-foot {
  margin-top: 17px;
  border-top: 1px solid #171512;
  padding-top: 6px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.12em;
  color: #6f695d;
}
@keyframes cu-dn-in {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-dn-in {
  animation: cu-dn-in 0.7s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-dn-d1 {
  animation-delay: 0.06s;
}
.cu-dn-d2 {
  animation-delay: 0.15s;
}
.cu-dn-d3 {
  animation-delay: 0.24s;
}
.cu-dn-d4 {
  animation-delay: 0.33s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-dn-in {
    animation: none;
  }
  .cu-dn-ad,
  .cu-dn-phone {
    transition: none;
  }
}
</style>

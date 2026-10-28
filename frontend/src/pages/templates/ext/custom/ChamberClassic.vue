<template>
  <div class="cu-cc-root min-h-screen">
    <div class="cu-cc-wrap mx-auto w-full max-w-md">
      <div class="cu-cc-frame">
        <header class="cu-cc-top cu-cc-in">
          <span>EST. MMVIII</span>
          <span>{{ nfcId ? 'N° ' + nfcId : 'N° —' }}</span>
        </header>

        <section class="cu-cc-crest-wrap cu-cc-in cu-cc-d1">
          <div class="cu-cc-crest">
            <span class="cu-cc-crest-ring"></span>
            <span class="cu-cc-crest-initial">{{ initial }}</span>
          </div>
        </section>

        <section class="cu-cc-title">
          <p class="cu-cc-overline cu-cc-in cu-cc-d2">Chambers &amp; Advisory</p>
          <h1 class="cu-cc-name cu-cc-in cu-cc-d2">{{ payload?.companyName || '企业主页' }}</h1>
          <div class="cu-cc-divider cu-cc-in cu-cc-d3">
            <span></span>
            <i>◆</i>
            <span></span>
          </div>
          <p class="cu-cc-desc cu-cc-in cu-cc-d3">{{ payload?.description }}</p>
        </section>

        <section class="cu-cc-practice cu-cc-in cu-cc-d4">
          <p class="cu-cc-section-label">AREAS OF PRACTICE</p>
          <div v-if="services.length" class="cu-cc-practice-grid">
            <article v-for="(service, i) in services" :key="i" class="cu-cc-practice-item">
              <span class="cu-cc-numeral">{{ numerals[i % numerals.length] }}</span>
              <span class="cu-cc-hair"></span>
              <h2>{{ service.name }}</h2>
            </article>
          </div>
        </section>

        <section class="cu-cc-contact cu-cc-in cu-cc-d5">
          <p class="cu-cc-section-label">ENQUIRIES</p>
          <p class="cu-cc-contact-note">如蒙垂询，敬请致电</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-cc-phone">
            <Phone :size="15" />
            <span>{{ payload.phone }}</span>
          </a>
          <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-cc-mail">{{ payload.email }}</a>
        </section>

        <footer class="cu-cc-foot cu-cc-in cu-cc-d5">
          <span class="cu-cc-foot-line"></span>
          <p>{{ payload?.companyName || '企业主页' }} · 专业服务</p>
          <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
        </footer>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const numerals = ['I', 'II', 'III', 'IV', 'V', 'VI']

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-cc-root {
  background: #0d2b25;
  color: #e9dfc6;
  font-family: 'Times New Roman', Georgia, 'Songti SC', serif;
  overflow-x: hidden;
}
.cu-cc-root::before {
  content: '';
  position: fixed;
  inset: 0;
  pointer-events: none;
  background: radial-gradient(circle at 50% 12%, rgba(198, 163, 92, 0.16), transparent 58%);
}
.cu-cc-wrap {
  position: relative;
  padding: 18px 14px 30px;
}
.cu-cc-frame {
  border: 1px solid rgba(198, 163, 92, 0.55);
  padding: 16px 16px 20px;
  box-shadow: inset 0 0 0 4px #0d2b25, inset 0 0 0 5px rgba(198, 163, 92, 0.28);
}
.cu-cc-top {
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.22em;
  color: #a58c5b;
}
.cu-cc-crest-wrap {
  display: flex;
  justify-content: center;
  margin-top: 24px;
}
.cu-cc-crest {
  position: relative;
  height: 78px;
  width: 78px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(198, 163, 92, 0.7);
  transform: rotate(45deg);
}
.cu-cc-crest-ring {
  position: absolute;
  inset: 6px;
  border: 1px solid rgba(198, 163, 92, 0.35);
}
.cu-cc-crest-initial {
  transform: rotate(-45deg);
  font-size: 34px;
  color: #c6a35c;
  letter-spacing: 0;
}
.cu-cc-title {
  margin-top: 22px;
  text-align: center;
}
.cu-cc-overline {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.32em;
  text-transform: uppercase;
  color: #c6a35c;
}
.cu-cc-name {
  margin-top: 10px;
  font-size: 34px;
  line-height: 1.14;
  font-weight: 400;
  letter-spacing: 0.06em;
  color: #f4ead2;
  word-break: break-word;
}
.cu-cc-divider {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin: 18px 0 14px;
  color: #c6a35c;
}
.cu-cc-divider span {
  height: 1px;
  width: 52px;
  background: rgba(198, 163, 92, 0.45);
}
.cu-cc-divider i {
  font-style: normal;
  font-size: 8px;
}
.cu-cc-desc {
  margin: 0 auto;
  max-width: 32ch;
  font-size: 13px;
  line-height: 1.85;
  font-style: italic;
  color: #bdb195;
}
.cu-cc-practice {
  margin-top: 30px;
  border-top: 1px solid rgba(198, 163, 92, 0.32);
  padding-top: 18px;
}
.cu-cc-section-label {
  text-align: center;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.3em;
  color: #c6a35c;
}
.cu-cc-practice-grid {
  margin-top: 16px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
}
.cu-cc-practice-item {
  text-align: center;
  padding: 4px 2px;
  transition: transform 0.25s ease;
}
.cu-cc-practice-item:hover {
  transform: translateY(-3px);
}
.cu-cc-numeral {
  display: block;
  font-size: 11px;
  color: #c6a35c;
  letter-spacing: 0.16em;
}
.cu-cc-hair {
  display: block;
  margin: 8px auto;
  height: 1px;
  width: 18px;
  background: rgba(198, 163, 92, 0.4);
}
.cu-cc-practice-item h2 {
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  font-size: 13px;
  font-weight: 400;
  line-height: 1.5;
  color: #e9dfc6;
}
.cu-cc-contact {
  margin-top: 26px;
  border-top: 1px solid rgba(198, 163, 92, 0.32);
  border-bottom: 1px solid rgba(198, 163, 92, 0.32);
  padding: 18px 0;
  text-align: center;
}
.cu-cc-contact-note {
  margin-top: 9px;
  font-size: 12px;
  letter-spacing: 0.18em;
  color: #bdb195;
}
.cu-cc-phone {
  margin-top: 12px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 20px;
  letter-spacing: 0.08em;
  color: #e8cd92;
  border-bottom: 1px solid rgba(198, 163, 92, 0.5);
  padding-bottom: 4px;
  transition: color 0.2s ease, border-color 0.2s ease;
}
.cu-cc-phone:hover {
  color: #fff3d4;
  border-color: #fff3d4;
}
.cu-cc-mail {
  display: block;
  margin-top: 10px;
  font-family: 'Courier New', monospace;
  font-size: 10px;
  letter-spacing: 0.1em;
  color: #a58c5b;
  word-break: break-all;
}
.cu-cc-empty {
  margin-top: 10px;
  font-size: 11px;
  color: #96896c;
  font-style: italic;
}
.cu-cc-foot {
  margin-top: 22px;
  text-align: center;
  font-size: 9px;
  letter-spacing: 0.16em;
  color: #8d7c56;
}
.cu-cc-foot-line {
  display: block;
  margin: 0 auto 10px;
  height: 1px;
  width: 34px;
  background: rgba(198, 163, 92, 0.5);
}
.cu-cc-foot p {
  color: #bdb195;
}
.cu-cc-foot span:last-child {
  display: block;
  margin-top: 5px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
}
@keyframes cu-cc-in {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-cc-in {
  animation: cu-cc-in 0.85s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-cc-d1 {
  animation-delay: 0.08s;
}
.cu-cc-d2 {
  animation-delay: 0.16s;
}
.cu-cc-d3 {
  animation-delay: 0.24s;
}
.cu-cc-d4 {
  animation-delay: 0.32s;
}
.cu-cc-d5 {
  animation-delay: 0.4s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-cc-in {
    animation: none;
  }
  .cu-cc-practice-item,
  .cu-cc-phone {
    transition: none;
  }
}
</style>

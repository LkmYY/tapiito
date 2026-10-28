<template>
  <div class="bc-edit-root min-h-screen px-4 py-6 text-[#17130f]">
    <div class="bc-edit-page mx-auto w-full max-w-md">
      <header class="bc-edit-masthead">
        <span class="bc-edit-brand">FORM</span>
        <span class="bc-edit-issue">ISSUE 07<br />PORTRAIT SERIES</span>
        <span class="bc-edit-year">2026</span>
      </header>

      <div class="bc-edit-cover">
        <div class="bc-edit-block bc-edit-block-a"></div>
        <div class="bc-edit-block bc-edit-block-b"></div>
        <p class="bc-edit-vertical">THE PEOPLE ISSUE</p>

        <div class="bc-edit-portrait">
          <span class="bc-edit-initial">{{ initial }}</span>
          <div class="bc-edit-portrait-line"></div>
          <span class="bc-edit-caption">PHOTO / TYPE ONLY</span>
        </div>

        <h1 class="bc-edit-name">{{ payload.name || '未命名' }}</h1>

        <div class="bc-edit-subject">
          <span>PROFILE</span>
          <p>{{ payload.title }}</p>
        </div>
      </div>

      <section class="bc-edit-credits">
        <p class="bc-edit-credit-head">CONTACT &amp; CREDITS</p>
        <div class="bc-edit-columns">
          <div class="bc-edit-col">
            <span class="bc-edit-label">ORG</span>
            <p>{{ payload.company }}</p>
            <span class="bc-edit-label">TEL</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-edit-link">{{ payload.phone }}</a>
          </div>
          <div class="bc-edit-col">
            <span class="bc-edit-label">MAIL</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-edit-link">{{ payload.email }}</a>
            <span class="bc-edit-label">WEB</span>
            <a
              v-if="payload.website"
              :href="payload.website"
              target="_blank"
              rel="noopener noreferrer"
              class="bc-edit-link"
            >
              {{ payload.website }}
            </a>
          </div>
        </div>
        <div class="bc-edit-foot">
          <span>WECHAT / {{ payload.wechat }}</span>
          <span>{{ nfcId ? 'NO. ' + nfcId : 'NO. —' }}</span>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-edit-root {
  overflow-x: hidden;
  background: #f2eee4;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.bc-edit-page {
  border: 1px solid rgba(23, 19, 15, 0.18);
  background: #f8f5ec;
  box-shadow: 0 30px 70px -44px rgba(23, 19, 15, 0.55);
  animation: bc-edit-page-in 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-edit-masthead {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding: 14px 16px 12px;
  border-bottom: 3px solid #17130f;
}
.bc-edit-brand {
  font-family: 'Times New Roman', Georgia, serif;
  font-size: 26px;
  font-weight: 700;
  letter-spacing: -0.06em;
  line-height: 1;
}
.bc-edit-issue {
  padding-top: 2px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  line-height: 1.6;
  letter-spacing: 0.16em;
  text-align: center;
}
.bc-edit-year {
  padding-top: 2px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.12em;
}
.bc-edit-cover {
  position: relative;
  min-height: 480px;
  overflow: hidden;
  padding: 24px 18px 0;
}
.bc-edit-block {
  position: absolute;
  pointer-events: none;
}
.bc-edit-block-a {
  top: 38px;
  right: -48px;
  width: 220px;
  height: 340px;
  background: #c2410c;
  transform: rotate(7deg);
  transition: transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.bc-edit-block-b {
  top: 190px;
  left: -70px;
  width: 250px;
  height: 150px;
  background: #e8d9b8;
  transform: rotate(-5deg);
  transition: transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.bc-edit-cover:hover .bc-edit-block-a {
  transform: rotate(4deg) translateY(-6px);
}
.bc-edit-cover:hover .bc-edit-block-b {
  transform: rotate(-2deg) translateX(6px);
}
.bc-edit-vertical {
  position: absolute;
  top: 30px;
  right: 22px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.3em;
  writing-mode: vertical-rl;
  color: #f8f5ec;
  z-index: 2;
}
.bc-edit-portrait {
  position: relative;
  z-index: 1;
  display: flex;
  width: 218px;
  height: 268px;
  margin: 12px auto 0;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(23, 19, 15, 0.22);
  background: linear-gradient(160deg, #f8f5ec 0%, #e9e1cf 100%);
  box-shadow: 14px 16px 0 rgba(23, 19, 15, 0.12);
}
.bc-edit-portrait::before {
  content: '';
  position: absolute;
  inset: 12px;
  border: 1px solid rgba(23, 19, 15, 0.2);
}
.bc-edit-initial {
  font-family: 'Times New Roman', Georgia, serif;
  font-size: 128px;
  line-height: 1;
  color: rgba(23, 19, 15, 0.86);
  transform: translateY(-8px);
}
.bc-edit-portrait-line {
  position: absolute;
  left: 18px;
  right: 18px;
  bottom: 42px;
  height: 1px;
  background: rgba(23, 19, 15, 0.35);
}
.bc-edit-caption {
  position: absolute;
  left: 18px;
  bottom: 24px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(23, 19, 15, 0.55);
}
.bc-edit-name {
  position: relative;
  z-index: 2;
  margin-top: -18px;
  font-family: 'Times New Roman', Georgia, 'Songti SC', serif;
  font-size: clamp(46px, 15vw, 66px);
  font-weight: 700;
  line-height: 0.9;
  letter-spacing: -0.055em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
  color: #17130f;
  text-shadow: 3px 3px 0 #f8f5ec;
  animation: bc-edit-title 1s 0.2s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-edit-subject {
  position: relative;
  z-index: 2;
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 14px;
  margin-top: 20px;
  padding: 10px 0 22px;
  border-top: 1px solid #17130f;
}
.bc-edit-subject span {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.26em;
}
.bc-edit-subject p {
  max-width: 68%;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 14px;
  font-style: italic;
  text-align: right;
}
.bc-edit-credits {
  border-top: 3px solid #17130f;
  padding: 14px 16px 16px;
  animation: bc-edit-fade 0.9s 0.4s both;
}
.bc-edit-credit-head {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.26em;
  color: #57534e;
}
.bc-edit-columns {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0;
  margin-top: 12px;
}
.bc-edit-col {
  min-width: 0;
  padding-right: 14px;
}
.bc-edit-col + .bc-edit-col {
  border-left: 1px solid rgba(23, 19, 15, 0.2);
  padding-right: 0;
  padding-left: 14px;
}
.bc-edit-label {
  display: block;
  margin-top: 10px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: #a8a29e;
}
.bc-edit-col .bc-edit-label:first-child {
  margin-top: 0;
}
.bc-edit-col p,
.bc-edit-link {
  display: block;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11.5px;
  line-height: 1.5;
  color: #17130f;
}
.bc-edit-link {
  position: relative;
  text-decoration: none;
  transition: color 0.25s ease;
}
.bc-edit-link:hover {
  color: #c2410c;
}
.bc-edit-link:active {
  opacity: 0.6;
}
.bc-edit-muted {
  color: #c7c1b4;
}
.bc-edit-foot {
  display: flex;
  justify-content: space-between;
  gap: 12px;
  margin-top: 16px;
  padding-top: 10px;
  border-top: 1px solid rgba(23, 19, 15, 0.18);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #78716c;
}
@keyframes bc-edit-page-in {
  from { opacity: 0; transform: translateY(24px) rotate(-0.6deg); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-edit-title {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-edit-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-edit-root *,
  .bc-edit-root *::before,
  .bc-edit-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

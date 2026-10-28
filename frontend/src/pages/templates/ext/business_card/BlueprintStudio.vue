<template>
  <div class="bc-blue-root min-h-screen px-4 py-7">
    <div class="bc-blue-grid"></div>
    <div class="bc-blue-wrap mx-auto w-full max-w-md">
      <div class="bc-blue-sheet">
        <span class="bc-blue-cross bc-blue-cross-tl"></span>
        <span class="bc-blue-cross bc-blue-cross-tr"></span>
        <span class="bc-blue-cross bc-blue-cross-bl"></span>
        <span class="bc-blue-cross bc-blue-cross-br"></span>

        <header class="bc-blue-head">
          <div>
            <p class="bc-blue-doc">PROJECT / IDENTITY</p>
            <p class="bc-blue-title">PERSONNEL SPECIFICATION</p>
          </div>
          <div class="bc-blue-rev">
            <span>REV. A</span>
            <span>SHEET 01 / 01</span>
          </div>
        </header>

        <div class="bc-blue-meta">
          <span>DRAWN BY: STUDIO YUNTU</span>
          <span>SCALE: 1:1</span>
          <span>DATE: 2026.10</span>
        </div>

        <section class="bc-blue-subject">
          <div class="bc-blue-axis bc-blue-axis-y">
            <span>0</span><span>50</span><span>100</span>
          </div>
          <div class="bc-blue-figure">
            <span class="bc-blue-figure-ring"></span>
            <span class="bc-blue-figure-initial">{{ initial }}</span>
            <span class="bc-blue-measure bc-blue-measure-top">Ø 68</span>
            <span class="bc-blue-measure bc-blue-measure-side">H 92</span>
          </div>
          <div class="bc-blue-callout">
            <span>01</span>
            <h1>{{ payload.name || '未命名' }}</h1>
            <p>{{ payload.title }}</p>
            <p class="bc-blue-org">{{ payload.company }}</p>
          </div>
        </section>

        <section class="bc-blue-table">
          <div class="bc-blue-table-head">
            <span>REF</span><span>FIELD</span><span>VALUE</span>
          </div>
          <div class="bc-blue-row">
            <span>01</span><span>TEL</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          </div>
          <div class="bc-blue-row">
            <span>02</span><span>EMAIL</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          </div>
          <div class="bc-blue-row">
            <span>03</span><span>URL</span>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
          </div>
          <div class="bc-blue-row">
            <span>04</span><span>WECHAT</span>
            <p :class="{ 'bc-blue-empty': !payload.wechat }">{{ payload.wechat }}</p>
          </div>
        </section>

        <footer class="bc-blue-foot">
          <div class="bc-blue-stamp">APPROVED</div>
          <div class="bc-blue-note">
            <span>NFC REF: {{ nfcId || '—' }}</span>
            <span>ALL DIMENSIONS IN MM</span>
          </div>
        </footer>
      </div>
      <p class="bc-blue-caption">DRAWING NOT TO SCALE · FOR INTRODUCTION PURPOSES ONLY</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-blue-root {
  position: relative;
  overflow: hidden;
  background: #082f49;
  font-family: 'SFMono-Regular', Consolas, 'Courier New', monospace;
  color: #e0f2fe;
}
.bc-blue-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(125, 211, 252, 0.14) 1px, transparent 1px),
    linear-gradient(90deg, rgba(125, 211, 252, 0.14) 1px, transparent 1px),
    linear-gradient(rgba(125, 211, 252, 0.06) 1px, transparent 1px),
    linear-gradient(90deg, rgba(125, 211, 252, 0.06) 1px, transparent 1px);
  background-size: 48px 48px, 48px 48px, 12px 12px, 12px 12px;
}
.bc-blue-root::after {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(ellipse at 50% 38%, transparent 20%, rgba(2, 20, 34, 0.55) 100%);
  pointer-events: none;
}
.bc-blue-wrap {
  position: relative;
  z-index: 1;
}
.bc-blue-sheet {
  position: relative;
  border: 1px solid rgba(186, 230, 253, 0.5);
  background: rgba(3, 35, 56, 0.45);
  padding: 22px 18px 18px;
  box-shadow: 0 30px 60px -30px rgba(2, 6, 23, 0.8);
  animation: bc-blue-draw 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-blue-sheet::before {
  content: '';
  position: absolute;
  inset: 7px;
  border: 1px dashed rgba(186, 230, 253, 0.16);
  pointer-events: none;
}
.bc-blue-cross {
  position: absolute;
  width: 12px;
  height: 12px;
}
.bc-blue-cross::before,
.bc-blue-cross::after {
  content: '';
  position: absolute;
  background: #7dd3fc;
}
.bc-blue-cross::before {
  left: 50%;
  top: 0;
  width: 1px;
  height: 100%;
}
.bc-blue-cross::after {
  top: 50%;
  left: 0;
  width: 100%;
  height: 1px;
}
.bc-blue-cross-tl { top: -6px; left: -6px; }
.bc-blue-cross-tr { top: -6px; right: -6px; }
.bc-blue-cross-bl { bottom: -6px; left: -6px; }
.bc-blue-cross-br { bottom: -6px; right: -6px; }
.bc-blue-head {
  display: flex;
  justify-content: space-between;
  gap: 14px;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(186, 230, 253, 0.45);
}
.bc-blue-doc {
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: #7dd3fc;
}
.bc-blue-title {
  margin-top: 5px;
  font-size: 11px;
  letter-spacing: 0.12em;
  color: #f0f9ff;
}
.bc-blue-rev {
  display: flex;
  flex-direction: column;
  gap: 4px;
  text-align: right;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(186, 230, 253, 0.65);
}
.bc-blue-meta {
  display: flex;
  justify-content: space-between;
  gap: 8px;
  padding: 8px 0;
  border-bottom: 1px dashed rgba(186, 230, 253, 0.25);
  font-size: 5.5px;
  letter-spacing: 0.1em;
  color: rgba(186, 230, 253, 0.55);
}
.bc-blue-subject {
  display: grid;
  grid-template-columns: 18px 118px 1fr;
  gap: 10px;
  align-items: center;
  padding: 24px 0 20px;
  border-bottom: 1px solid rgba(186, 230, 253, 0.3);
}
.bc-blue-axis {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-self: stretch;
  padding: 8px 0;
  border-right: 1px solid rgba(186, 230, 253, 0.25);
  font-size: 5px;
  color: rgba(186, 230, 253, 0.45);
}
.bc-blue-figure {
  position: relative;
  display: flex;
  width: 108px;
  height: 108px;
  align-items: center;
  justify-content: center;
}
.bc-blue-figure-ring {
  position: absolute;
  inset: 6px;
  border: 1px solid #38bdf8;
  border-radius: 50%;
  box-shadow: 0 0 22px rgba(56, 189, 248, 0.25), inset 0 0 22px rgba(56, 189, 248, 0.14);
}
.bc-blue-figure-ring::before {
  content: '';
  position: absolute;
  inset: 13px;
  border: 1px dashed rgba(125, 211, 252, 0.5);
  border-radius: 50%;
}
.bc-blue-figure-initial {
  position: relative;
  font-family: Georgia, serif;
  font-size: 40px;
  color: #bae6fd;
}
.bc-blue-measure {
  position: absolute;
  font-size: 5px;
  letter-spacing: 0.08em;
  color: rgba(125, 211, 252, 0.7);
}
.bc-blue-measure-top {
  top: -9px;
  left: 50%;
  transform: translateX(-50%);
}
.bc-blue-measure-side {
  right: -20px;
  top: 50%;
  transform: rotate(90deg) translateY(-50%);
}
.bc-blue-callout span {
  display: inline-block;
  margin-bottom: 5px;
  border: 1px solid #38bdf8;
  padding: 1px 5px;
  font-size: 6px;
  color: #7dd3fc;
}
.bc-blue-callout h1 {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 21px;
  font-weight: 700;
  line-height: 1.15;
  letter-spacing: -0.02em;
  color: #f0f9ff;
  overflow-wrap: anywhere;
}
.bc-blue-callout p {
  margin-top: 6px;
  font-size: 7.5px;
  letter-spacing: 0.08em;
  color: #7dd3fc;
}
.bc-blue-callout .bc-blue-org {
  margin-top: 3px;
  color: rgba(224, 242, 254, 0.55);
}
.bc-blue-table {
  padding-top: 16px;
}
.bc-blue-table-head {
  display: grid;
  grid-template-columns: 26px 52px 1fr;
  padding-bottom: 5px;
  border-bottom: 1px solid rgba(186, 230, 253, 0.4);
  font-size: 5.5px;
  letter-spacing: 0.16em;
  color: rgba(186, 230, 253, 0.55);
}
.bc-blue-row {
  display: grid;
  grid-template-columns: 26px 52px 1fr;
  align-items: baseline;
  gap: 4px;
  padding: 9px 0;
  border-bottom: 1px dotted rgba(186, 230, 253, 0.2);
  transition: background 0.2s ease;
}
.bc-blue-row:hover {
  background: rgba(56, 189, 248, 0.08);
}
.bc-blue-row > span:first-child {
  font-size: 5.5px;
  color: rgba(125, 211, 252, 0.45);
}
.bc-blue-row > span:nth-child(2) {
  font-size: 6px;
  letter-spacing: 0.1em;
  color: #7dd3fc;
}
.bc-blue-row a,
.bc-blue-row em,
.bc-blue-row p {
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 9px;
  font-style: normal;
  letter-spacing: 0.02em;
  color: #e0f2fe;
  text-decoration: none;
  transition: color 0.2s ease;
}
.bc-blue-row a:hover {
  color: #7dd3fc;
  text-decoration: underline;
}
.bc-blue-row a:active {
  opacity: 0.6;
}
.bc-blue-row em,
.bc-blue-empty {
  color: rgba(186, 230, 253, 0.3);
}
.bc-blue-foot {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 12px;
  margin-top: 18px;
  padding-top: 12px;
  border-top: 1px solid rgba(186, 230, 253, 0.4);
}
.bc-blue-stamp {
  border: 2px solid #38bdf8;
  padding: 4px 8px;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: #38bdf8;
  transform: rotate(-4deg);
  transition: background 0.25s ease, color 0.25s ease, transform 0.25s ease;
}
.bc-blue-stamp:hover {
  background: #38bdf8;
  color: #082f49;
  transform: rotate(0);
}
.bc-blue-note {
  display: flex;
  flex-direction: column;
  gap: 4px;
  text-align: right;
  font-size: 5.5px;
  letter-spacing: 0.1em;
  color: rgba(186, 230, 253, 0.5);
}
.bc-blue-caption {
  margin-top: 12px;
  text-align: center;
  font-size: 5.5px;
  letter-spacing: 0.12em;
  color: rgba(186, 230, 253, 0.35);
}
@keyframes bc-blue-draw {
  from { opacity: 0; transform: translateY(18px); clip-path: inset(0 0 100% 0); }
  to { opacity: 1; transform: none; clip-path: inset(0 0 0 0); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-blue-root *,
  .bc-blue-root *::before,
  .bc-blue-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

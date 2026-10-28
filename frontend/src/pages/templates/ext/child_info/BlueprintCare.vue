<!-- 求助信息卡 · 蓝图标识：工程蓝图网格，信息像规格表一样清楚 -->
<template>
  <div class="ci-blue min-h-screen">
    <div class="mx-auto w-full max-w-md px-5 pb-12 pt-8">
      <div class="ci-sheet">
        <div class="ci-grid" aria-hidden="true"></div>
        <span class="ci-corner ci-corner--tl" aria-hidden="true"></span>
        <span class="ci-corner ci-corner--br" aria-hidden="true"></span>

        <header class="ci-head">
          <div>
            <p class="ci-doc">DOC. CI-2026 / REV. A</p>
            <p class="ci-doc-sub">GUARDIAN ASSISTANCE SPECIFICATION</p>
          </div>
          <span class="ci-mark"><LifeBuoy :size="20" /></span>
        </header>

        <div class="ci-title-row">
          <span class="ci-title-no">01</span>
          <h1>求助信息<br />ASSISTANCE</h1>
        </div>

        <p class="ci-lead">发现有人需要帮助时，请按本图信息联系其家属，并在现场陪伴等待。</p>

        <div class="ci-specs">
          <div class="ci-spec ci-spec--full">
            <span class="ci-key">A · 基本情况 / PROFILE</span>
            <p class="ci-val ci-val--lg">{{ payload.baseInfo }}</p>
          </div>
          <div class="ci-spec">
            <span class="ci-key">B · 学校 / 单位</span>
            <p class="ci-val">{{ payload.school }}</p>
          </div>
          <div class="ci-spec">
            <span class="ci-key">C · 过敏史 / 病史</span>
            <p class="ci-val ci-val--alert">{{ payload.allergies || '无' }}</p>
          </div>
        </div>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-box">
            <span class="ci-call-label">D · 紧急联系电话 / CALL</span>
            <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
          </span>
          <span class="ci-call-side">
            <Phone :size="21" />
            <span>拨打</span>
          </span>
        </a>

        <footer class="ci-foot">
          <span>请陪伴至家属到达</span>
          <span v-if="nfcId">NFC / {{ nfcId }}</span>
          <span>THANK YOU</span>
        </footer>
      </div>
    </div>
  </div>
</template>

<script setup>
import { LifeBuoy, Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-blue {
  background: #071e30;
  color: #d9f3ff;
  font-family: "Bahnschrift", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-sheet {
  position: relative;
  overflow: hidden;
  border: 1px solid #2b6f92;
  background: #0a2a42;
  padding: 21px 19px 17px;
  box-shadow: 0 26px 50px -30px rgba(0, 0, 0, 0.9);
  animation: ci-blue-in 0.6s ease both;
}
.ci-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(109, 197, 232, 0.08) 1px, transparent 1px),
    linear-gradient(90deg, rgba(109, 197, 232, 0.08) 1px, transparent 1px);
  background-size: 22px 22px;
  pointer-events: none;
}
.ci-sheet::after {
  content: "";
  position: absolute;
  inset: 7px;
  border: 1px solid rgba(109, 197, 232, 0.2);
  pointer-events: none;
}
.ci-corner {
  position: absolute;
  z-index: 1;
  height: 13px;
  width: 13px;
  border-color: #73d2f5;
  border-style: solid;
}
.ci-corner--tl { left: 16px; top: 16px; border-width: 1px 0 0 1px; }
.ci-corner--br { right: 16px; bottom: 16px; border-width: 0 1px 1px 0; }
.ci-head {
  position: relative;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 10px;
  border-bottom: 1px solid #2b6f92;
  padding-bottom: 11px;
}
.ci-doc {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #6dc5e8;
}
.ci-doc-sub {
  margin-top: 3px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.08em;
  color: #4b86a4;
}
.ci-mark {
  display: flex;
  height: 34px;
  width: 34px;
  align-items: center;
  justify-content: center;
  border: 1px solid #387d9d;
  color: #7dd8fa;
}
.ci-title-row {
  position: relative;
  display: flex;
  align-items: flex-start;
  gap: 11px;
  margin-top: 17px;
  animation: ci-blue-in 0.6s 0.07s ease both;
}
.ci-title-no {
  border: 1px solid #4f9bba;
  padding: 2px 5px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 9px;
  color: #8fe0ff;
}
.ci-title-row h1 {
  font-size: 30px;
  font-weight: 700;
  line-height: 1.25;
  letter-spacing: 0.01em;
  color: #ecfaff;
}
.ci-lead {
  position: relative;
  margin-top: 10px;
  font-size: 11px;
  line-height: 1.8;
  color: #9ac4d8;
  animation: ci-blue-in 0.6s 0.12s ease both;
}
.ci-specs {
  position: relative;
  margin-top: 16px;
  border-top: 1px solid #2b6f92;
  border-left: 1px solid #2b6f92;
  display: grid;
  grid-template-columns: 1fr 1fr;
  animation: ci-blue-in 0.6s 0.17s ease both;
}
.ci-spec {
  border-right: 1px solid #2b6f92;
  border-bottom: 1px solid #2b6f92;
  padding: 10px;
}
.ci-spec--full { grid-column: 1 / -1; }
.ci-key {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.06em;
  color: #6dc5e8;
}
.ci-val {
  margin-top: 4px;
  font-size: 12px;
  font-weight: 600;
  line-height: 1.55;
  color: #e4f7ff;
}
.ci-val--lg { font-size: 17px; }
.ci-val--alert { color: #ffc26e; }
.ci-call {
  position: relative;
  display: flex;
  align-items: stretch;
  margin-top: 15px;
  border: 1px solid #7dd8fa;
  background: #d6f3ff;
  color: #062434;
  text-decoration: none;
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-blue-in 0.6s 0.23s ease both;
}
.ci-call:hover { background: #eafaff; }
.ci-call:active { transform: scale(0.985); }
.ci-call-box {
  flex: 1;
  min-width: 0;
  padding: 12px 13px;
}
.ci-call-label {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.06em;
  color: #2a6a88;
}
.ci-call-num {
  display: block;
  margin-top: 4px;
  font-size: 28px;
  font-weight: 800;
  line-height: 1.16;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #042030;
}
.ci-call-side {
  display: flex;
  width: 61px;
  flex-shrink: 0;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  border-left: 1px dashed #5ba6c3;
  font-size: 9px;
  font-weight: 700;
  color: #0d4d6a;
}
.ci-empty {
  position: relative;
  margin-top: 15px;
  border: 1px dashed #4b86a4;
  padding: 14px;
  font-size: 12px;
  color: #8fc7de;
}
.ci-foot {
  position: relative;
  display: flex;
  justify-content: space-between;
  gap: 7px;
  border-top: 1px solid #2b6f92;
  margin-top: 15px;
  padding-top: 9px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.06em;
  color: #5d94ae;
  animation: ci-blue-in 0.6s 0.28s ease both;
}
@keyframes ci-blue-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-sheet, .ci-title-row, .ci-lead, .ci-specs, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>

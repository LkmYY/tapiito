<!-- 求助信息卡 · 站台情报：琥珀点阵信息板，像车站显示屏一样直接 -->
<template>
  <div class="ci-board min-h-screen">
    <div class="mx-auto w-full max-w-md px-5 pb-12 pt-8">
      <header class="ci-head">
        <span class="ci-dot" aria-hidden="true"></span>
        <span class="ci-head-title">CARE STATION · 求助站台</span>
        <span v-if="nfcId" class="ci-head-id">PLATFORM {{ nfcId }}</span>
      </header>

      <div class="ci-display">
        <div class="ci-display-top">
          <span class="ci-live">● LIVE</span>
          <span class="ci-time">请留意 · PLEASE NOTICE</span>
        </div>

        <div class="ci-destination">
          <span class="ci-destination-label">本次求助 / DESTINATION</span>
          <h1>送我回家<span class="ci-cursor">_</span></h1>
          <p>发现有人需要帮助。请按下方信息联系家属，并陪伴至家人到达。</p>
        </div>

        <div class="ci-table">
          <div class="ci-table-head">
            <span>信息 / INFO</span>
            <span>内容 / DETAIL</span>
          </div>
          <div class="ci-row">
            <span class="ci-key">基本情况</span>
            <span class="ci-val">{{ payload.baseInfo }}</span>
          </div>
          <div class="ci-row">
            <span class="ci-key">学校 / 单位</span>
            <span class="ci-val">{{ payload.school }}</span>
          </div>
          <div class="ci-row">
            <span class="ci-key">过敏史 / 病史</span>
            <span class="ci-val ci-val--warn">{{ payload.allergies || '无' }}</span>
          </div>
        </div>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-row">
            <span class="ci-call-label"><Phone :size="15" /> 紧急联系电话 / CALL</span>
            <span class="ci-call-cta">点击拨打 →</span>
          </span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        </a>

        <footer class="ci-footer">
          <span>请陪伴至家人到达</span>
          <span class="ci-footer-blink">THANK YOU</span>
        </footer>
      </div>

      <p class="ci-note">请勿让需要帮助的人独自离开</p>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-board {
  background: #070707;
  color: #ffb524;
  font-family: "Cascadia Mono", "Consolas", "Courier New", monospace;
}
.ci-head {
  display: flex;
  align-items: center;
  gap: 7px;
  border-bottom: 1px solid #3a2a08;
  padding-bottom: 10px;
  animation: ci-board-in 0.5s ease both;
}
.ci-dot {
  height: 6px;
  width: 6px;
  border-radius: 50%;
  background: #ffb524;
  box-shadow: 0 0 9px rgba(255, 181, 36, 0.7);
}
.ci-head-title {
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.12em;
}
.ci-head-id {
  margin-left: auto;
  font-size: 8px;
  color: #8a6a22;
}
.ci-display {
  margin-top: 14px;
  border: 1px solid #4a350b;
  background: #0d0b05;
  padding: 16px 15px 14px;
  box-shadow: inset 0 0 50px rgba(255, 181, 36, 0.045);
  animation: ci-board-in 0.6s 0.06s ease both;
}
.ci-display-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border-bottom: 1px dashed #4a350b;
  padding-bottom: 8px;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #a97c1c;
}
.ci-live { color: #ff5a3c; }
.ci-destination {
  padding: 13px 0 12px;
}
.ci-destination-label {
  font-size: 8px;
  letter-spacing: 0.14em;
  color: #8a6a22;
}
.ci-destination h1 {
  margin-top: 5px;
  font-size: 32px;
  font-weight: 700;
  line-height: 1.2;
  letter-spacing: 0.02em;
  color: #ffcf6a;
  text-shadow: 0 0 16px rgba(255, 181, 36, 0.26);
}
.ci-cursor {
  display: inline-block;
  margin-left: 3px;
  color: #ffb524;
  animation: ci-board-blink 1.1s steps(1) infinite;
}
.ci-destination p {
  margin-top: 8px;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 11px;
  line-height: 1.8;
  color: #b39a67;
}
.ci-table {
  border-top: 1px solid #4a350b;
}
.ci-table-head {
  display: grid;
  grid-template-columns: 32% 68%;
  gap: 8px;
  padding: 7px 0;
  font-size: 7px;
  letter-spacing: 0.1em;
  color: #7a5e1e;
  border-bottom: 1px solid #332506;
}
.ci-row {
  display: grid;
  grid-template-columns: 32% 68%;
  gap: 8px;
  border-bottom: 1px solid #2b2007;
  padding: 9px 0;
}
.ci-key {
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 9px;
  line-height: 1.6;
  color: #9b7c34;
}
.ci-val {
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 12px;
  font-weight: 600;
  line-height: 1.55;
  color: #f3e5c3;
}
.ci-val--warn { color: #ff9b6a; }
.ci-call {
  display: block;
  margin-top: 13px;
  border: 1px solid #ffb524;
  background: #ffb524;
  padding: 11px 12px 10px;
  color: #171204;
  text-decoration: none;
  transition: transform 0.16s ease, background 0.16s ease;
  animation: ci-board-in 0.6s 0.16s ease both;
}
.ci-call:hover { background: #ffc957; }
.ci-call:active { transform: scale(0.987); }
.ci-call-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}
.ci-call-label {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.08em;
}
.ci-call-cta {
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 8px;
  font-weight: 700;
}
.ci-call-num {
  display: block;
  margin-top: 6px;
  font-size: 32px;
  font-weight: 700;
  line-height: 1.16;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #0d0b05;
}
.ci-empty {
  margin-top: 13px;
  border: 1px dashed #7a5e1e;
  padding: 13px;
  font-size: 11px;
  color: #b39a67;
}
.ci-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border-top: 1px dashed #4a350b;
  margin-top: 12px;
  padding-top: 8px;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 9px;
  color: #9b7c34;
}
.ci-footer-blink {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.1em;
  animation: ci-board-blink 1.6s steps(1) infinite;
}
.ci-note {
  margin-top: 12px;
  text-align: center;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: #7a5e1e;
  animation: ci-board-in 0.6s 0.22s ease both;
}
@keyframes ci-board-in {
  from { opacity: 0; transform: translateY(11px); }
  to { opacity: 1; transform: none; }
}
@keyframes ci-board-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0.25; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-display, .ci-call, .ci-note { animation: none; }
  .ci-cursor, .ci-footer-blink { animation: none; }
  .ci-call { transition: none; }
}
</style>

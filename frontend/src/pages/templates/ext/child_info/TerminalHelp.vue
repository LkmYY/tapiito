<!-- 求助信息卡 · 终端求助：深绿终端与时序列表，冷静可靠 -->
<template>
  <div class="ci-term min-h-screen">
    <div class="mx-auto w-full max-w-md px-5 pb-12 pt-8">
      <div class="ci-window">
        <div class="ci-titlebar">
          <span class="ci-dots"><i></i><i></i><i></i></span>
          <span class="ci-window-title">care@help: ~/assist</span>
          <span v-if="nfcId" class="ci-window-id">tty{{ nfcId }}</span>
        </div>

        <div class="ci-screen">
          <p class="ci-line ci-line--dim">// 求助信息卡 · SOS PROTOCOL v2.4</p>
          <p class="ci-line ci-line--bright">$ locate guardian --urgent</p>

          <h1 class="ci-title">NEED<br />ASSISTANCE<span class="ci-cursor">_</span></h1>
          <p class="ci-lead">发现有人需要帮助时，请执行下方拨号操作，并留在现场等待家属。</p>

          <div class="ci-block">
            <p class="ci-block-label">[01] 基本情况 / PROFILE</p>
            <p class="ci-name">{{ payload.baseInfo }}</p>
          </div>

          <div class="ci-block">
            <p class="ci-block-label">[02] 学校 / 单位</p>
            <p class="ci-value">{{ payload.school }}</p>
          </div>

          <div class="ci-block">
            <p class="ci-block-label">[03] 过敏史 / 病史</p>
            <p class="ci-value ci-value--warn">{{ payload.allergies || '无' }}</p>
          </div>

          <a
            v-if="payload.emergencyPhone"
            :href="'tel:' + payload.emergencyPhone"
            class="ci-call"
          >
            <span class="ci-call-label">&gt; CALL_GUARDIAN --NOW</span>
            <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
            <span class="ci-call-hint"><Phone :size="14" /> 点击执行拨号 / TAP TO CALL</span>
          </a>

          <p class="ci-end">&gt; 请陪伴至家属到达。谢谢您的善意。</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-term {
  background: #05100b;
  color: #bdf5cf;
  font-family: "Cascadia Mono", "SFMono-Regular", Consolas, monospace;
}
.ci-window {
  overflow: hidden;
  border: 1px solid #1e5b3a;
  border-radius: 8px;
  background: #06150d;
  box-shadow: 0 0 0 1px #03170d, 0 30px 60px -30px rgba(60, 220, 130, 0.22);
  animation: ci-term-in 0.55s ease both;
}
.ci-titlebar {
  display: flex;
  align-items: center;
  gap: 9px;
  border-bottom: 1px solid #16452c;
  background: #0a2115;
  padding: 9px 11px;
}
.ci-dots { display: flex; gap: 5px; }
.ci-dots i {
  height: 7px;
  width: 7px;
  border-radius: 50%;
  background: #2c7049;
}
.ci-dots i:first-child { background: #d97706; }
.ci-dots i:last-child { background: #22c55e; }
.ci-window-title {
  font-size: 9px;
  letter-spacing: 0.04em;
  color: #6fa987;
}
.ci-window-id {
  margin-left: auto;
  font-size: 8px;
  color: #4c7a5d;
}
.ci-screen {
  padding: 20px 17px 22px;
}
.ci-line { font-size: 10px; line-height: 1.7; }
.ci-line--dim { color: #5d8c70; }
.ci-line--bright { margin-top: 4px; color: #7dffb2; }
.ci-title {
  margin-top: 13px;
  font-size: 34px;
  font-weight: 700;
  line-height: 1.12;
  letter-spacing: -0.02em;
  color: #d7ffe6;
  text-shadow: 0 0 18px rgba(90, 255, 160, 0.22);
  animation: ci-term-in 0.55s 0.07s ease both;
}
.ci-cursor {
  display: inline-block;
  margin-left: 3px;
  color: #58ff9d;
  animation: ci-blink 1.1s steps(1) infinite;
}
.ci-lead {
  margin-top: 11px;
  font-size: 11px;
  line-height: 1.85;
  color: #7fb697;
  animation: ci-term-in 0.55s 0.12s ease both;
}
.ci-block {
  margin-top: 14px;
  border-left: 2px solid #1f5d3b;
  padding-left: 11px;
  animation: ci-term-in 0.55s 0.17s ease both;
}
.ci-block-label {
  font-size: 9px;
  letter-spacing: 0.09em;
  color: #58a879;
}
.ci-name {
  margin-top: 4px;
  font-size: 17px;
  font-weight: 700;
  line-height: 1.55;
  color: #ddffea;
}
.ci-value {
  margin-top: 4px;
  font-size: 13px;
  line-height: 1.55;
  color: #b8e8cb;
}
.ci-value--warn { color: #ffc46b; }
.ci-call {
  display: block;
  margin-top: 19px;
  border: 1px solid #52e58c;
  border-radius: 5px;
  background: #0c2c1a;
  padding: 15px;
  color: #eafff1;
  text-decoration: none;
  box-shadow: 0 0 24px -8px rgba(82, 229, 140, 0.3), inset 0 0 28px rgba(82, 229, 140, 0.05);
  transition: background 0.18s ease, box-shadow 0.18s ease, transform 0.18s ease;
  animation: ci-term-in 0.55s 0.22s ease both;
}
.ci-call:hover {
  background: #123d26;
  box-shadow: 0 0 30px -6px rgba(82, 229, 140, 0.44), inset 0 0 28px rgba(82, 229, 140, 0.08);
}
.ci-call:active { transform: scale(0.985); }
.ci-call-label {
  display: block;
  font-size: 9px;
  letter-spacing: 0.06em;
  color: #64e698;
}
.ci-call-num {
  display: block;
  margin-top: 6px;
  font-size: 29px;
  font-weight: 700;
  line-height: 1.2;
  letter-spacing: 0.02em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #ffffff;
}
.ci-call-hint {
  display: flex;
  align-items: center;
  gap: 5px;
  margin-top: 7px;
  font-size: 9px;
  color: #86b89b;
}
.ci-empty {
  margin-top: 19px;
  border: 1px dashed #3e7a59;
  border-radius: 5px;
  padding: 14px;
  font-size: 12px;
  color: #86b89b;
}
.ci-end {
  margin-top: 15px;
  border-top: 1px dashed #1d4d33;
  padding-top: 11px;
  font-size: 9px;
  line-height: 1.8;
  color: #5f9477;
  animation: ci-term-in 0.55s 0.28s ease both;
}
@keyframes ci-term-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: none; }
}
@keyframes ci-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-window, .ci-title, .ci-lead, .ci-block, .ci-call, .ci-end { animation: none; }
  .ci-cursor { animation: none; }
  .ci-call { transition: none; }
}
</style>

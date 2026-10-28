<!-- 求助信息卡 · 时间小径：鼠尾草绿与砂色，用时间轴串起每条信息 -->
<template>
  <div class="ci-trail min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-10">
      <header class="ci-head">
        <span class="ci-brand"><Footprints :size="15" /> 回家的小径</span>
        <span v-if="nfcId" class="ci-id">{{ nfcId }}</span>
      </header>

      <h1 class="ci-title">陪他走完<br />最后一段路</h1>
      <p class="ci-sub">如果您遇见他，请按下面的顺序，帮我们把他送回家。</p>

      <ol class="ci-track">
        <li class="ci-step">
          <span class="ci-node">1</span>
          <div class="ci-step-body">
            <span class="ci-key">基本情况</span>
            <p class="ci-name">{{ payload.baseInfo }}</p>
          </div>
        </li>
        <li class="ci-step">
          <span class="ci-node">2</span>
          <div class="ci-step-body">
            <span class="ci-key">学校 / 单位</span>
            <p class="ci-value">{{ payload.school }}</p>
          </div>
        </li>
        <li class="ci-step">
          <span class="ci-node ci-node--warn">3</span>
          <div class="ci-step-body">
            <span class="ci-key">过敏史 / 病史</span>
            <p class="ci-value ci-value--warn">{{ payload.allergies || '无' }}</p>
          </div>
        </li>
        <li class="ci-step ci-step--last">
          <span class="ci-node ci-node--call"><Phone :size="15" /></span>
          <div class="ci-step-body">
            <span class="ci-key">最后一步 · 拨通电话</span>
            <a
              v-if="payload.emergencyPhone"
              :href="'tel:' + payload.emergencyPhone"
              class="ci-call"
            >
              <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
              <span class="ci-call-tap">点击拨打</span>
            </a>
          </div>
        </li>
      </ol>

      <p class="ci-foot"><HeartHandshake :size="14" /> 请在原地陪伴，直到家人出现</p>
    </div>
  </div>
</template>

<script setup>
import { Footprints, HeartHandshake, Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-trail {
  background: #f2f0e8;
  color: #2c352a;
  font-family: "Avenir Next", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: ci-trail-in 0.55s ease both;
}
.ci-brand {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  letter-spacing: 0.16em;
  color: #71806c;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 9px;
  color: #93a08e;
}
.ci-title {
  margin-top: 34px;
  font-size: 34px;
  font-weight: 600;
  line-height: 1.3;
  letter-spacing: -0.02em;
  color: #283324;
  animation: ci-trail-in 0.6s 0.06s ease both;
}
.ci-sub {
  margin-top: 10px;
  max-width: 19em;
  font-size: 13px;
  line-height: 1.8;
  color: #687563;
  animation: ci-trail-in 0.6s 0.11s ease both;
}
.ci-track {
  position: relative;
  margin-top: 26px;
  animation: ci-trail-in 0.6s 0.16s ease both;
}
.ci-track::before {
  content: "";
  position: absolute;
  left: 15px;
  top: 16px;
  bottom: 33px;
  width: 1px;
  background: repeating-linear-gradient(180deg, #b2bfa9 0 5px, transparent 5px 10px);
}
.ci-step {
  position: relative;
  display: flex;
  gap: 14px;
  padding-bottom: 22px;
}
.ci-node {
  position: relative;
  z-index: 1;
  display: flex;
  height: 31px;
  width: 31px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #f2f0e8;
  box-shadow: inset 0 0 0 1px #9eae95;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 11px;
  color: #56664f;
}
.ci-node--warn {
  box-shadow: inset 0 0 0 1px #c08a5a;
  color: #a05a28;
}
.ci-node--call {
  background: #37503a;
  box-shadow: 0 0 0 4px #dfe7d8;
  color: #f2f0e8;
}
.ci-step-body {
  min-width: 0;
  flex: 1;
  border-bottom: 1px solid #d8dccf;
  padding: 4px 0 13px;
}
.ci-step--last .ci-step-body { border-bottom: none; padding-bottom: 0; }
.ci-key {
  font-size: 9px;
  letter-spacing: 0.12em;
  color: #7d8c76;
}
.ci-name {
  margin-top: 5px;
  font-size: 20px;
  font-weight: 650;
  line-height: 1.55;
  color: #273122;
}
.ci-value {
  margin-top: 5px;
  font-size: 15px;
  font-weight: 550;
  line-height: 1.55;
  color: #394735;
}
.ci-value--warn { color: #9a4c1c; }
.ci-call {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 10px;
  margin-top: 8px;
  border-radius: 12px;
  background: #37503a;
  padding: 14px 16px;
  color: #f6f8f1;
  text-decoration: none;
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-trail-in 0.6s 0.22s ease both;
}
.ci-call:hover { background: #2b4030; }
.ci-call:active { transform: scale(0.985); }
.ci-call-num {
  font-size: 27px;
  font-weight: 700;
  line-height: 1.2;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-tap {
  flex-shrink: 0;
  font-size: 10px;
  color: #c2d3b8;
}
.ci-empty {
  margin-top: 8px;
  border: 1px dashed #a5b49c;
  border-radius: 12px;
  padding: 13px;
  font-size: 12px;
  color: #6d7d67;
}
.ci-foot {
  display: flex;
  align-items: center;
  gap: 7px;
  border-top: 1px solid #d0d7c7;
  margin-top: 2px;
  padding-top: 14px;
  font-size: 11px;
  color: #71806c;
  animation: ci-trail-in 0.6s 0.28s ease both;
}
@keyframes ci-trail-in {
  from { opacity: 0; transform: translateY(13px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-title, .ci-sub, .ci-track, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>

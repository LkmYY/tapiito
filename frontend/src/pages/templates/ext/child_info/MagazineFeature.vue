<!-- 求助信息卡 · 杂志专题：衬线编辑排版，克制、体面，像一篇人物报道 -->
<template>
  <div class="ci-mag min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-12 pt-8">
      <header class="ci-masthead">
        <span class="ci-mast-title">CARE REVIEW</span>
        <span class="ci-mast-meta">VOL. 07 / 2026</span>
      </header>
      <div class="ci-mast-rule"><span>人物 · 求助档案</span><span>ISSUE 07</span></div>

      <p class="ci-kicker">A CARD FOR A SAFE RETURN</p>
      <h1 class="ci-title">如果看见<br /><em>这位家人</em></h1>
      <p class="ci-deck">请替我们停下脚步，拨通一个电话。您的一分钟，是他回家的整段路。</p>

      <div class="ci-rule"></div>

      <div class="ci-columns">
        <div class="ci-lead-col">
          <p class="ci-drop"><span>{{ (payload.baseInfo).charAt(0) }}</span>{{ (payload.baseInfo).slice(1) }}</p>
        </div>
        <div class="ci-side-col">
          <div class="ci-fact">
            <span>学校 / 单位</span>
            <p>{{ payload.school }}</p>
          </div>
          <div class="ci-fact">
            <span>过敏史 / 病史</span>
            <p class="ci-alert">{{ payload.allergies || '无' }}</p>
          </div>
        </div>
      </div>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-copy">
          <span class="ci-call-label">紧急联系电话 / CALL THE FAMILY</span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        </span>
        <span class="ci-call-icon"><Phone :size="21" /></span>
      </a>

      <footer class="ci-foot">
        <span class="ci-foot-no">07</span>
        <p>请陪伴在旁，等家人赶到。感谢您成为这个故事里温柔的一页。</p>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-mag {
  background: #f7f4ee;
  color: #1c1b19;
  font-family: Georgia, "Songti SC", "SimSun", serif;
}
.ci-masthead {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 10px;
  animation: ci-mag-in 0.55s ease both;
}
.ci-mast-title {
  font-size: 21px;
  font-weight: 700;
  letter-spacing: 0.18em;
}
.ci-mast-meta {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #7a756c;
}
.ci-mast-rule {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  margin-top: 6px;
  border-top: 2px solid #1c1b19;
  border-bottom: 1px solid #1c1b19;
  padding: 5px 0;
  font-size: 9px;
  letter-spacing: 0.12em;
  color: #4c4841;
  animation: ci-mag-in 0.55s 0.05s ease both;
}
.ci-kicker {
  margin-top: 23px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: #a1462c;
  animation: ci-mag-in 0.55s 0.09s ease both;
}
.ci-title {
  margin-top: 8px;
  font-size: 43px;
  font-weight: 400;
  line-height: 1.17;
  letter-spacing: -0.02em;
  color: #191816;
  animation: ci-mag-in 0.6s 0.12s ease both;
}
.ci-title em {
  font-style: italic;
  color: #a1462c;
}
.ci-deck {
  margin-top: 12px;
  max-width: 22em;
  font-size: 13px;
  line-height: 1.9;
  color: #5d574e;
  animation: ci-mag-in 0.6s 0.16s ease both;
}
.ci-rule {
  margin-top: 18px;
  border-top: 1px solid #c9c1b3;
}
.ci-columns {
  display: grid;
  grid-template-columns: 1.18fr 0.82fr;
  gap: 15px;
  padding: 14px 0 4px;
  border-bottom: 1px solid #c9c1b3;
  animation: ci-mag-in 0.6s 0.2s ease both;
}
.ci-lead-col { min-width: 0; }
.ci-drop {
  font-size: 13px;
  line-height: 1.95;
  color: #37332c;
  text-align: justify;
}
.ci-drop span {
  float: left;
  margin: 5px 6px 0 0;
  font-size: 42px;
  font-weight: 700;
  line-height: 0.82;
  color: #a1462c;
}
.ci-side-col {
  border-left: 1px solid #d3cbc0;
  padding-left: 12px;
}
.ci-fact {
  border-bottom: 1px dotted #bdb4a6;
  padding: 7px 0;
}
.ci-fact:last-child { border-bottom: none; }
.ci-fact span {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.06em;
  color: #7a756c;
}
.ci-fact p {
  margin-top: 3px;
  font-size: 11px;
  font-weight: 700;
  line-height: 1.5;
}
.ci-fact .ci-alert { color: #a1462c; }
.ci-call {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  margin-top: 17px;
  border-top: 2px solid #a1462c;
  border-bottom: 2px solid #a1462c;
  padding: 13px 1px;
  color: #1c1b19;
  text-decoration: none;
  transition: background 0.2s ease, padding 0.2s ease;
  animation: ci-mag-in 0.6s 0.24s ease both;
}
.ci-call:hover {
  background: #f0e8dd;
  padding-left: 8px;
  padding-right: 8px;
}
.ci-call:active { transform: scale(0.99); }
.ci-call-copy { min-width: 0; }
.ci-call-label {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.12em;
  color: #8b4330;
}
.ci-call-num {
  display: block;
  margin-top: 5px;
  font-size: 31px;
  font-weight: 700;
  line-height: 1.15;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-icon {
  display: flex;
  height: 41px;
  width: 41px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border: 1px solid #1c1b19;
  border-radius: 50%;
}
.ci-empty {
  margin-top: 17px;
  border-top: 2px solid #a1462c;
  border-bottom: 2px solid #a1462c;
  padding: 14px 1px;
  font-size: 13px;
  color: #8b4330;
}
.ci-foot {
  display: flex;
  align-items: flex-start;
  gap: 11px;
  margin-top: 18px;
  animation: ci-mag-in 0.6s 0.28s ease both;
}
.ci-foot-no {
  font-size: 31px;
  font-weight: 700;
  line-height: 1;
  color: #c3b8a6;
}
.ci-foot p {
  font-size: 10.5px;
  line-height: 1.8;
  color: #6b655b;
}
@keyframes ci-mag-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-masthead, .ci-mast-rule, .ci-kicker, .ci-title, .ci-deck, .ci-columns, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>

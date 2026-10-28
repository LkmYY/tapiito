<!-- 求助信息卡 · 水墨家书：宣纸与墨色，一封写给陌生人的托付 -->
<template>
  <div class="ci-ink min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-10">
      <div class="ci-paper">
        <div class="ci-wash ci-wash--one" aria-hidden="true"></div>
        <div class="ci-wash ci-wash--two" aria-hidden="true"></div>

        <div class="ci-side" aria-hidden="true">
          <span>见字如面</span>
          <i></i>
        </div>

        <header class="ci-head">
          <span class="ci-brand">家书 · 托付</span>
          <span v-if="nfcId" class="ci-id">第 {{ nfcId }} 号</span>
        </header>

        <div class="ci-title-wrap">
          <h1>若您遇见<br />我的家人</h1>
          <span class="ci-seal"><span>谢</span><span>谢</span></span>
        </div>

        <p class="ci-lead">他有时会忘记回家的路。若您在街角看见他，请替我们拨一个电话，陪他坐一会儿。</p>

        <div class="ci-brush" aria-hidden="true"></div>

        <section class="ci-info">
          <div class="ci-row">
            <span class="ci-key">基本情况</span>
            <p class="ci-name">{{ payload.baseInfo }}</p>
          </div>
          <div class="ci-row ci-row--split">
            <div>
              <span class="ci-key">学校 / 单位</span>
              <p>{{ payload.school }}</p>
            </div>
            <div>
              <span class="ci-key">过敏史 / 病史</span>
              <p class="ci-alert">{{ payload.allergies || '无' }}</p>
            </div>
          </div>
        </section>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-head">
            <Phone :size="16" />
            <span>紧急联系电话</span>
          </span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
          <span class="ci-call-tap">轻点号码，与我们联系</span>
        </a>

        <p class="ci-sign">愿您平安，也愿善意有回响。</p>
        <p class="ci-date">二〇二六年 · 秋</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-ink {
  background: #d9d2c1;
  color: #252620;
  font-family: Georgia, "Songti SC", "SimSun", serif;
}
.ci-paper {
  position: relative;
  overflow: hidden;
  background-color: #f4efe2;
  background-image:
    repeating-linear-gradient(0deg, rgba(90, 78, 55, 0.035) 0 1px, transparent 1px 5px),
    radial-gradient(circle at 20% 12%, rgba(120, 107, 78, 0.12), transparent 38%);
  padding: 25px 24px 23px 41px;
  box-shadow: 0 28px 50px -32px rgba(52, 45, 30, 0.75);
  animation: ci-ink-in 0.7s ease both;
}
.ci-wash {
  position: absolute;
  border-radius: 50%;
  pointer-events: none;
  filter: blur(1px);
}
.ci-wash--one {
  right: -52px;
  top: 108px;
  height: 190px;
  width: 190px;
  background: radial-gradient(circle at 40% 40%, rgba(47, 54, 46, 0.14), rgba(47, 54, 46, 0) 68%);
}
.ci-wash--two {
  left: -35px;
  bottom: 120px;
  height: 150px;
  width: 150px;
  background: radial-gradient(circle, rgba(47, 54, 46, 0.1), transparent 70%);
}
.ci-side {
  position: absolute;
  left: 14px;
  top: 96px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  color: #8a7c60;
  writing-mode: vertical-rl;
  font-size: 9px;
  letter-spacing: 0.34em;
}
.ci-side i {
  height: 47px;
  width: 1px;
  background: #b9ab8b;
}
.ci-head {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border-bottom: 1px solid #cabfa8;
  padding-bottom: 9px;
}
.ci-brand {
  font-size: 10px;
  letter-spacing: 0.2em;
  color: #7e7055;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #9b8d71;
}
.ci-title-wrap {
  position: relative;
  margin-top: 25px;
  animation: ci-ink-in 0.65s 0.08s ease both;
}
.ci-title-wrap h1 {
  font-size: 40px;
  font-weight: 600;
  line-height: 1.32;
  letter-spacing: 0.04em;
  color: #21231e;
}
.ci-seal {
  position: absolute;
  right: 2px;
  bottom: 3px;
  display: flex;
  height: 55px;
  width: 55px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: 2px solid #a63b2b;
  color: #a63b2b;
  font-size: 12px;
  font-weight: 700;
  line-height: 1.15;
  transform: rotate(7deg);
  opacity: 0.82;
}
.ci-lead {
  position: relative;
  margin-top: 14px;
  max-width: 20em;
  font-size: 13px;
  line-height: 2;
  color: #5f5a4b;
  letter-spacing: 0.02em;
  animation: ci-ink-in 0.65s 0.14s ease both;
}
.ci-brush {
  height: 3px;
  margin: 19px 0 6px;
  background: linear-gradient(90deg, #2a2c26 0%, #2a2c26 38%, rgba(42, 44, 38, 0.35) 64%, transparent 100%);
  transform: rotate(-0.5deg);
}
.ci-info {
  position: relative;
  animation: ci-ink-in 0.65s 0.2s ease both;
}
.ci-row {
  border-bottom: 1px solid rgba(138, 124, 96, 0.28);
  padding: 11px 0;
}
.ci-row--split {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 13px;
}
.ci-key {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.15em;
  color: #94866b;
}
.ci-name {
  margin-top: 5px;
  font-size: 19px;
  font-weight: 600;
  line-height: 1.65;
  color: #20221d;
}
.ci-row p {
  margin-top: 4px;
  font-size: 13px;
  font-weight: 600;
  line-height: 1.55;
  color: #383a32;
}
.ci-row .ci-alert { color: #a63b2b; }
.ci-call {
  display: block;
  margin-top: 17px;
  border: 1px solid #2a2c26;
  background: #2a2c26;
  padding: 14px 16px 12px;
  color: #f1ead9;
  text-decoration: none;
  transition: transform 0.2s ease, background 0.2s ease;
  animation: ci-ink-in 0.65s 0.26s ease both;
}
.ci-call:hover { background: #3a3d34; }
.ci-call:active { transform: scale(0.985); }
.ci-call-head {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  letter-spacing: 0.12em;
  color: #cfc3a8;
}
.ci-call-num {
  display: block;
  margin-top: 5px;
  font-size: 31px;
  font-weight: 600;
  line-height: 1.2;
  letter-spacing: 0.03em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #fff9ec;
}
.ci-call-tap {
  display: block;
  margin-top: 4px;
  font-size: 9px;
  color: #b8ab90;
}
.ci-empty {
  margin-top: 17px;
  border: 1px dashed #8a7c60;
  padding: 14px;
  font-size: 12px;
  color: #7e7055;
}
.ci-sign {
  margin-top: 18px;
  text-align: right;
  font-size: 12px;
  line-height: 1.7;
  color: #5f5a4b;
}
.ci-date {
  margin-top: 4px;
  text-align: right;
  font-size: 9px;
  letter-spacing: 0.16em;
  color: #94866b;
  animation: ci-ink-in 0.65s 0.32s ease both;
}
@keyframes ci-ink-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-paper, .ci-title-wrap, .ci-lead, .ci-info, .ci-call, .ci-date { animation: none; }
  .ci-call { transition: none; }
}
</style>

<!-- 求助信息卡 · 拱门之家：陶土色拱门与奶油底，像一扇等着他回家的门 -->
<template>
  <div class="ci-arch min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-9">
      <header class="ci-head">
        <span class="ci-brand"><DoorOpen :size="16" /> 回家之门</span>
        <span v-if="nfcId" class="ci-id">NO. {{ nfcId }}</span>
      </header>

      <div class="ci-archway">
        <div class="ci-arch-inner">
          <p class="ci-kicker">A WAY HOME</p>
          <h1>请帮我<br />推开这扇门</h1>
          <p class="ci-lead">如果您遇见独自在外的我，请拨打电话，让家人来接我回家。</p>

          <div class="ci-person">
            <span class="ci-person-mark">守</span>
            <div>
              <span class="ci-key">基本情况</span>
              <p class="ci-name">{{ payload.baseInfo }}</p>
            </div>
          </div>
        </div>
      </div>

      <div class="ci-info">
        <div class="ci-item">
          <span class="ci-key">学校 / 单位</span>
          <p>{{ payload.school }}</p>
        </div>
        <div class="ci-item">
          <span class="ci-key">过敏史 / 病史</span>
          <p class="ci-alert">{{ payload.allergies || '无' }}</p>
        </div>
      </div>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-label"><Phone :size="17" /> 紧急联系电话</span>
        <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        <span class="ci-call-tap">点击拨打，让我们接他回家</span>
      </a>

      <p class="ci-foot">谢谢您为他留一扇门。</p>
    </div>
  </div>
</template>

<script setup>
import { DoorOpen, Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-arch {
  background: #f8efe0;
  color: #4a2c1c;
  font-family: Georgia, "Songti SC", "SimSun", serif;
}
.ci-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: ci-arch-in 0.6s ease both;
}
.ci-brand {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  letter-spacing: 0.18em;
  color: #9a5a34;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #b0876c;
}
.ci-archway {
  position: relative;
  margin-top: 24px;
  overflow: hidden;
  border-radius: 155px 155px 12px 12px;
  background: linear-gradient(165deg, #b45309 0%, #9a3f0c 68%, #873508 100%);
  padding: 44px 25px 26px;
  color: #fff4e3;
  box-shadow: 0 28px 48px -30px rgba(120, 53, 15, 0.8);
  animation: ci-arch-in 0.7s 0.06s ease both;
}
.ci-archway::before {
  content: "";
  position: absolute;
  inset: 11px;
  border: 1px solid rgba(255, 244, 227, 0.25);
  border-radius: 145px 145px 7px 7px;
  pointer-events: none;
}
.ci-archway::after {
  content: "";
  position: absolute;
  left: 50%;
  top: 21px;
  height: 7px;
  width: 7px;
  transform: translateX(-50%);
  border-radius: 50%;
  background: #f5cf9a;
  box-shadow: 0 0 0 4px rgba(245, 207, 154, 0.18);
}
.ci-arch-inner { position: relative; text-align: center; }
.ci-kicker {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.26em;
  color: #f1c896;
}
.ci-archway h1 {
  margin-top: 12px;
  font-size: 34px;
  font-weight: 600;
  line-height: 1.34;
  letter-spacing: 0.02em;
}
.ci-lead {
  margin: 11px auto 0;
  max-width: 17em;
  font-size: 12px;
  line-height: 1.9;
  color: #f6ddc2;
}
.ci-person {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: 20px;
  border-top: 1px solid rgba(255, 244, 227, 0.28);
  padding-top: 15px;
  text-align: left;
}
.ci-person-mark {
  display: flex;
  height: 44px;
  width: 44px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255, 244, 227, 0.45);
  border-radius: 50%;
  font-size: 17px;
  color: #ffe6c4;
}
.ci-person .ci-key { color: #e8b98b; }
.ci-name {
  margin-top: 3px;
  font-size: 16px;
  font-weight: 600;
  line-height: 1.55;
  color: #fff7ec;
}
.ci-info {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1px;
  margin-top: 17px;
  border-top: 1px solid #d9c1a4;
  border-bottom: 1px solid #d9c1a4;
  animation: ci-arch-in 0.65s 0.13s ease both;
}
.ci-item {
  padding: 12px 3px;
}
.ci-item + .ci-item {
  border-left: 1px solid #d9c1a4;
  padding-left: 14px;
}
.ci-key {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #a5795a;
}
.ci-item p {
  margin-top: 4px;
  font-size: 13px;
  font-weight: 600;
  line-height: 1.5;
  color: #51301d;
}
.ci-item .ci-alert { color: #a63d1a; }
.ci-call {
  display: block;
  margin-top: 17px;
  border-radius: 5px;
  background: #4a2c1c;
  padding: 16px 18px 15px;
  color: #fff4e3;
  text-decoration: none;
  transition: background 0.2s ease, transform 0.2s ease;
  animation: ci-arch-in 0.65s 0.19s ease both;
}
.ci-call:hover { background: #613a25; }
.ci-call:active { transform: scale(0.985); }
.ci-call-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  letter-spacing: 0.12em;
  color: #e8b98b;
}
.ci-call-num {
  display: block;
  margin-top: 6px;
  font-size: 33px;
  font-weight: 600;
  line-height: 1.16;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #fffdf8;
}
.ci-call-tap {
  display: block;
  margin-top: 4px;
  font-size: 9px;
  color: #c8a384;
}
.ci-empty {
  margin-top: 17px;
  border: 1px dashed #b0876c;
  border-radius: 5px;
  padding: 15px;
  font-size: 12px;
  color: #8f6044;
}
.ci-foot {
  margin-top: 16px;
  text-align: center;
  font-size: 11px;
  letter-spacing: 0.04em;
  color: #9a6a4c;
  animation: ci-arch-in 0.65s 0.25s ease both;
}
@keyframes ci-arch-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-archway, .ci-info, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>

<!-- 求助信息卡 · 牛皮挂牌：行李牌造型与棉绳细节，朴素又可靠 -->
<template>
  <div class="ci-kraft min-h-screen">
    <div class="mx-auto w-full max-w-md px-7 pb-12 pt-10">
      <div class="ci-string" aria-hidden="true">
        <span class="ci-string-knot"></span>
      </div>

      <div class="ci-tag">
        <span class="ci-hole" aria-hidden="true"></span>

        <div class="ci-tag-head">
          <span class="ci-tag-brand"><Luggage :size="15" /> 随身求助牌</span>
          <span v-if="nfcId" class="ci-tag-no">NO. {{ nfcId }}</span>
        </div>

        <div class="ci-tag-body">
          <p class="ci-kicker">IF I AM LOST</p>
          <h1>请帮我<br />联系家人</h1>
          <p class="ci-note">我可能一个人出门了。请您拨打下方的电话，我会在原地等您。</p>

          <div class="ci-slip">
            <div class="ci-slip-row">
              <span>基本情况</span>
              <p>{{ payload.baseInfo }}</p>
            </div>
            <div class="ci-slip-row">
              <span>学校 / 单位</span>
              <p>{{ payload.school }}</p>
            </div>
            <div class="ci-slip-row">
              <span>过敏史 / 病史</span>
              <p class="ci-alert">{{ payload.allergies || '无' }}</p>
            </div>
          </div>

          <a
            v-if="payload.emergencyPhone"
            :href="'tel:' + payload.emergencyPhone"
            class="ci-call"
          >
            <span class="ci-call-label"><Phone :size="16" /> 紧急联系电话</span>
            <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
            <span class="ci-call-tap">点击号码拨打</span>
          </a>
        </div>

        <div class="ci-tag-foot">
          <span>谢谢您 · THANK YOU</span>
          <span class="ci-barcode" aria-hidden="true"></span>
        </div>
      </div>

      <p class="ci-foot">请陪我等到家人出现，再次感谢。</p>
    </div>
  </div>
</template>

<script setup>
import { Luggage, Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-kraft {
  background: #efe6d6;
  color: #332416;
  font-family: "Courier New", "PingFang SC", "Microsoft YaHei", monospace;
}
.ci-string {
  position: relative;
  height: 36px;
  margin-left: 46px;
  width: 2px;
  background: #8a6a45;
  animation: ci-kraft-in 0.5s ease both;
}
.ci-string::before {
  content: "";
  position: absolute;
  left: -27px;
  top: -16px;
  height: 36px;
  width: 58px;
  border: 2px solid #8a6a45;
  border-radius: 50%;
  transform: rotate(-14deg);
}
.ci-string-knot {
  position: absolute;
  left: -5px;
  bottom: -4px;
  height: 10px;
  width: 12px;
  border-radius: 3px;
  background: #7a5b39;
}
.ci-tag {
  position: relative;
  overflow: hidden;
  border-radius: 9px 9px 24px 24px;
  background-color: #c9a678;
  background-image:
    linear-gradient(115deg, rgba(255, 255, 255, 0.1), transparent 40%),
    repeating-linear-gradient(8deg, rgba(90, 60, 25, 0.05) 0 2px, transparent 2px 4px);
  padding: 26px 22px 17px;
  box-shadow: 0 2px 0 #a98456, 0 26px 44px -28px rgba(66, 43, 18, 0.8);
  animation: ci-kraft-in 0.6s 0.06s ease both;
}
.ci-hole {
  position: absolute;
  left: 50%;
  top: 12px;
  height: 16px;
  width: 36px;
  transform: translateX(-50%);
  border-radius: 9px;
  background: #efe6d6;
  box-shadow: inset 0 2px 4px rgba(70, 45, 18, 0.45);
}
.ci-tag-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border-bottom: 1px dashed rgba(51, 36, 22, 0.4);
  padding-bottom: 9px;
}
.ci-tag-brand {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.08em;
}
.ci-tag-no {
  font-size: 9px;
  color: #5f452b;
}
.ci-tag-body {
  padding-top: 15px;
}
.ci-kicker {
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.2em;
  color: #5f452b;
}
.ci-tag h1 {
  margin-top: 7px;
  font-size: 36px;
  font-weight: 800;
  line-height: 1.2;
  letter-spacing: -0.03em;
  color: #2f2012;
}
.ci-note {
  margin-top: 9px;
  font-size: 12px;
  line-height: 1.8;
  color: #4d3720;
}
.ci-slip {
  margin-top: 15px;
  border-radius: 4px;
  background: #f6eddd;
  padding: 5px 13px;
  box-shadow: 0 1px 0 rgba(70, 45, 18, 0.3);
}
.ci-slip-row {
  border-bottom: 1px dashed #cdb795;
  padding: 9px 0;
}
.ci-slip-row:last-child { border-bottom: none; }
.ci-slip-row span {
  display: block;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: #8a6a45;
}
.ci-slip-row p {
  margin-top: 3px;
  font-size: 13px;
  font-weight: 700;
  line-height: 1.5;
  color: #2f2012;
}
.ci-slip-row .ci-alert { color: #9a3412; }
.ci-call {
  display: block;
  margin-top: 14px;
  border-radius: 4px;
  background: #2f2012;
  padding: 13px 14px 12px;
  color: #f6eddd;
  text-decoration: none;
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-kraft-in 0.55s 0.2s ease both;
}
.ci-call:hover { background: #4a3420; }
.ci-call:active { transform: scale(0.985); }
.ci-call-label {
  display: flex;
  align-items: center;
  gap: 5px;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #d8b98c;
}
.ci-call-num {
  display: block;
  margin-top: 5px;
  font-size: 30px;
  font-weight: 800;
  line-height: 1.16;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #ffffff;
}
.ci-call-tap {
  display: block;
  margin-top: 3px;
  font-size: 9px;
  color: #bda17d;
}
.ci-empty {
  margin-top: 14px;
  border: 1px dashed rgba(47, 32, 18, 0.5);
  border-radius: 4px;
  padding: 13px;
  font-size: 12px;
}
.ci-tag-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  margin-top: 16px;
  border-top: 1px dashed rgba(51, 36, 22, 0.4);
  padding-top: 9px;
  font-size: 8px;
  letter-spacing: 0.08em;
  color: #5f452b;
}
.ci-barcode {
  height: 19px;
  width: 58px;
  background: repeating-linear-gradient(90deg, #332416 0 1px, transparent 1px 3px, #332416 3px 5px, transparent 5px 6px);
  opacity: 0.75;
}
.ci-foot {
  margin-top: 15px;
  text-align: center;
  font-size: 10px;
  color: #6f5537;
  animation: ci-kraft-in 0.6s 0.28s ease both;
}
@keyframes ci-kraft-in {
  from { opacity: 0; transform: translateY(13px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-string, .ci-tag, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>

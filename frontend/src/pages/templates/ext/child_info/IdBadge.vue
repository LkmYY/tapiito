<!-- 求助信息卡 · 安心证件：挂绳工牌造型，深青与米白，正式又亲切 -->
<template>
  <div class="ci-badge min-h-screen">
    <div class="mx-auto w-full max-w-md px-7 pb-14 pt-9">
      <div class="ci-lanyard" aria-hidden="true">
        <span class="ci-strap ci-strap--left"></span>
        <span class="ci-strap ci-strap--right"></span>
        <span class="ci-clip"></span>
      </div>

      <div class="ci-card">
        <header class="ci-card-head">
          <span class="ci-card-brand"><BadgeCheck :size="15" /> 平安守护</span>
          <span v-if="nfcId" class="ci-card-no">ID {{ nfcId }}</span>
        </header>

        <div class="ci-photo">
          <span class="ci-photo-ring">守</span>
          <span class="ci-photo-tag">求助信息卡</span>
        </div>

        <div class="ci-body">
          <h1>如果您遇见他<br />请帮他联系家人</h1>
          <p class="ci-lead">他可能需要一点帮助才能回家。请拨打下方电话，并陪伴至家人到达。</p>

          <div class="ci-field">
            <span class="ci-k">基本情况 / PROFILE</span>
            <p class="ci-v ci-v--name">{{ payload.baseInfo }}</p>
          </div>
          <div class="ci-field-row">
            <div class="ci-field">
              <span class="ci-k">学校 / 单位</span>
              <p class="ci-v">{{ payload.school }}</p>
            </div>
            <div class="ci-field">
              <span class="ci-k">过敏史 / 病史</span>
              <p class="ci-v ci-v--alert">{{ payload.allergies || '无' }}</p>
            </div>
          </div>
        </div>

        <div class="ci-call-wrap">
          <a
            v-if="payload.emergencyPhone"
            :href="'tel:' + payload.emergencyPhone"
            class="ci-call"
          >
            <span class="ci-call-head">
              <span class="ci-call-label"><Phone :size="15" /> 紧急联系电话</span>
              <span class="ci-call-tap">点击拨打</span>
            </span>
            <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
          </a>
        </div>

        <footer class="ci-card-foot">
          <div class="ci-barcode" aria-hidden="true"></div>
          <span class="ci-foot-text">请陪伴至家人到达 · 感谢您的善意</span>
        </footer>
      </div>

      <p class="ci-under">此卡仅用于紧急联系 · 信息受保护</p>
    </div>
  </div>
</template>

<script setup>
import { BadgeCheck, Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-badge {
  background: linear-gradient(160deg, #0b3a38 0%, #072a29 100%);
  color: #17312e;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-lanyard {
  position: relative;
  display: flex;
  height: 69px;
  justify-content: center;
  animation: ci-badge-in 0.5s ease both;
}
.ci-strap {
  position: absolute;
  top: -38px;
  height: 96px;
  width: 15px;
  border-radius: 8px;
  background: repeating-linear-gradient(45deg, #d5b26f 0 4px, #f0d9a8 4px 8px);
  box-shadow: inset 0 0 0 1px rgba(70, 48, 10, 0.25);
}
.ci-strap--left { left: 38%; transform: rotate(15deg); }
.ci-strap--right { right: 38%; transform: rotate(-15deg); }
.ci-clip {
  position: absolute;
  bottom: 19px;
  height: 18px;
  width: 43px;
  border-radius: 5px;
  background: #c9c5b6;
  box-shadow: inset 0 -3px 0 rgba(0, 0, 0, 0.18), 0 2px 4px rgba(0, 0, 0, 0.3);
  z-index: 2;
}
.ci-clip::after {
  content: "";
  position: absolute;
  left: 50%;
  top: -8px;
  height: 10px;
  width: 17px;
  transform: translateX(-50%);
  border: 3px solid #a9a495;
  border-radius: 5px 5px 0 0;
}
.ci-card {
  position: relative;
  overflow: hidden;
  border-radius: 18px;
  background: #fbf8ef;
  box-shadow: 0 30px 60px -30px rgba(0, 0, 0, 0.9), inset 0 0 0 1px rgba(255, 255, 255, 0.6);
  animation: ci-badge-in 0.65s 0.06s ease both;
}
.ci-card-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  background: #0d4a46;
  padding: 12px 21px;
  color: #d7efe5;
}
.ci-card-brand {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.16em;
}
.ci-card-no {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #86b8ac;
}
.ci-photo {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 22px 21px 0;
}
.ci-photo-ring {
  display: flex;
  height: 78px;
  width: 78px;
  align-items: center;
  justify-content: center;
  border: 1px solid #b9c7b5;
  border-radius: 50%;
  background: radial-gradient(circle at 34% 28%, #e7efe3, #cbd9c8);
  font-family: Georgia, "Songti SC", serif;
  font-size: 31px;
  color: #2c524b;
  box-shadow: 0 0 0 5px #fbf8ef, 0 0 0 6px #d5ddcf;
}
.ci-photo-tag {
  margin-top: 12px;
  border: 1px solid #c8b98f;
  border-radius: 2px;
  padding: 3px 9px;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #8a7440;
}
.ci-body {
  padding: 17px 21px 0;
}
.ci-body h1 {
  font-size: 26px;
  font-weight: 700;
  line-height: 1.4;
  letter-spacing: -0.02em;
  color: #12332f;
}
.ci-lead {
  margin-top: 7px;
  font-size: 11px;
  line-height: 1.8;
  color: #6e7c73;
}
.ci-field {
  border-top: 1px solid #ded9c9;
  margin-top: 13px;
  padding-top: 10px;
}
.ci-field-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 13px;
}
.ci-field-row .ci-field { margin-top: 10px; }
.ci-k {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.1em;
  color: #8a948c;
}
.ci-v {
  margin-top: 4px;
  font-size: 12px;
  font-weight: 600;
  line-height: 1.5;
  color: #263f39;
}
.ci-v--name { font-size: 16px; }
.ci-v--alert { color: #a14a26; }
.ci-call-wrap {
  margin-top: 16px;
  padding: 0 21px;
}
.ci-call {
  display: block;
  border-radius: 10px;
  background: #0d4a46;
  padding: 13px 15px 12px;
  color: #f3fbf5;
  text-decoration: none;
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-badge-in 0.6s 0.2s ease both;
}
.ci-call:hover { background: #125a54; }
.ci-call:active { transform: scale(0.985); }
.ci-call-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}
.ci-call-label {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #a9d3c6;
}
.ci-call-tap {
  font-size: 8px;
  color: #7fb5a7;
}
.ci-call-num {
  display: block;
  margin-top: 5px;
  font-family: "Bahnschrift", "PingFang SC", sans-serif;
  font-size: 29px;
  font-weight: 700;
  line-height: 1.16;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #ffffff;
}
.ci-empty {
  border: 1px dashed #9bb5a9;
  border-radius: 10px;
  padding: 14px;
  font-size: 12px;
  color: #6b8278;
  text-align: center;
}
.ci-card-foot {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 15px;
  background: #eee9d9;
  padding: 12px 21px;
}
.ci-barcode {
  height: 23px;
  width: 51px;
  flex-shrink: 0;
  background: repeating-linear-gradient(90deg, #3f5148 0 1px, transparent 1px 3px, #3f5148 3px 5px, transparent 5px 6px);
  opacity: 0.6;
}
.ci-foot-text {
  font-size: 8px;
  line-height: 1.6;
  color: #7b816f;
}
.ci-under {
  margin-top: 13px;
  text-align: center;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.05em;
  color: #6f9b8f;
  animation: ci-badge-in 0.6s 0.26s ease both;
}
@keyframes ci-badge-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-lanyard, .ci-card, .ci-call, .ci-under { animation: none; }
  .ci-call { transition: none; }
}
</style>

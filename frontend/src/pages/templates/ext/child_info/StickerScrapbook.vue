<!-- 求助信息卡 · 贴纸手账：胶带与手写便签，像孩子亲手做的一张卡 -->
<template>
  <div class="ci-scrap min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-10">
      <div class="ci-scrap-board">
        <span class="ci-tape ci-tape--tl" aria-hidden="true"></span>
        <span class="ci-tape ci-tape--tr" aria-hidden="true"></span>

        <span class="ci-doodle ci-doodle--star" aria-hidden="true">✶</span>
        <span class="ci-doodle ci-doodle--heart" aria-hidden="true">♥</span>

        <header class="ci-head">
          <span class="ci-tag"><Sticker :size="14" /> 我的求助卡</span>
          <span v-if="nfcId" class="ci-no">#{{ nfcId }}</span>
        </header>

        <h1 class="ci-title">请帮我找<br />到家人呀</h1>
        <p class="ci-lead">如果你看到我一个人，请帮我打个电话。我会乖乖在原地等你。</p>

        <div class="ci-note">
          <span class="ci-pin" aria-hidden="true"></span>
          <span class="ci-note-label">关于我</span>
          <p class="ci-name">{{ payload.baseInfo }}</p>
          <div class="ci-note-lines">
            <p><span>学校 / 单位</span><strong>{{ payload.school }}</strong></p>
            <p><span>过敏史 / 病史</span><strong class="ci-red">{{ payload.allergies || '无' }}</strong></p>
          </div>
        </div>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-sticker"><Phone :size="18" /></span>
          <span class="ci-call-copy">
            <span class="ci-call-label">打这个电话找家人</span>
            <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
          </span>
          <span class="ci-call-arrow" aria-hidden="true">↗</span>
        </a>

        <p class="ci-foot">谢谢你，好不好？<span aria-hidden="true">☺</span></p>
      </div>

      <p class="ci-credit">made with care · 请陪我等一等</p>
    </div>
  </div>
</template>

<script setup>
import { Phone, Sticker } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-scrap {
  background: #f4efe5;
  background-image: radial-gradient(#d8cdb9 0.8px, transparent 0.9px);
  background-size: 17px 17px;
  color: #3a2f3f;
  font-family: "Comic Sans MS", "PingFang SC", "Microsoft YaHei", cursive;
}
.ci-scrap-board {
  position: relative;
  background: #fffdf7;
  padding: 26px 22px 21px;
  box-shadow: 0 2px 0 #e3d9c4, 0 26px 46px -30px rgba(80, 60, 35, 0.6);
  transform: rotate(-0.5deg);
  animation: ci-scrap-in 0.65s ease both;
}
.ci-tape {
  position: absolute;
  z-index: 2;
  height: 24px;
  width: 79px;
  background: rgba(255, 216, 138, 0.72);
  box-shadow: 0 1px 2px rgba(120, 90, 40, 0.12);
}
.ci-tape--tl { left: -16px; top: 12px; transform: rotate(-35deg); }
.ci-tape--tr { right: -17px; top: 18px; transform: rotate(32deg); background: rgba(167, 216, 195, 0.68); }
.ci-doodle {
  position: absolute;
  color: #e7a33e;
}
.ci-doodle--star { right: 21px; top: 60px; font-size: 23px; transform: rotate(12deg); }
.ci-doodle--heart { right: 37px; bottom: 45px; font-size: 15px; color: #d97a86; transform: rotate(-13deg); }
.ci-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}
.ci-tag {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  border: 1.5px solid #7c5b8e;
  border-radius: 3px;
  padding: 4px 7px;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.07em;
  color: #6d4d80;
  transform: rotate(-1.5deg);
}
.ci-no {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #a18a74;
}
.ci-title {
  margin-top: 24px;
  font-size: 37px;
  font-weight: 800;
  line-height: 1.25;
  letter-spacing: -0.01em;
  color: #4e3558;
  animation: ci-scrap-in 0.6s 0.07s ease both;
}
.ci-lead {
  margin-top: 9px;
  max-width: 21em;
  font-size: 13px;
  line-height: 1.9;
  color: #7c6b78;
  animation: ci-scrap-in 0.6s 0.12s ease both;
}
.ci-note {
  position: relative;
  margin-top: 20px;
  background: #fdf3d7;
  padding: 18px 16px 13px;
  box-shadow: 2px 3px 0 rgba(120, 95, 55, 0.14);
  transform: rotate(0.7deg);
  animation: ci-scrap-in 0.6s 0.17s ease both;
}
.ci-pin {
  position: absolute;
  left: 50%;
  top: -7px;
  height: 14px;
  width: 14px;
  transform: translateX(-50%);
  border-radius: 50%;
  background: #d98b4a;
  box-shadow: inset 0 -2px 0 rgba(0, 0, 0, 0.18);
}
.ci-note-label {
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.18em;
  color: #a37c3d;
}
.ci-name {
  margin-top: 4px;
  font-size: 18px;
  font-weight: 800;
  line-height: 1.5;
  color: #4a3925;
}
.ci-note-lines {
  margin-top: 9px;
  border-top: 1px dashed #d8bd8a;
  padding-top: 6px;
}
.ci-note-lines p {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 10px;
  padding: 4px 0;
}
.ci-note-lines span {
  flex-shrink: 0;
  font-size: 10px;
  color: #a37c3d;
}
.ci-note-lines strong {
  max-width: 64%;
  text-align: right;
  font-size: 12px;
  font-weight: 700;
  line-height: 1.5;
  color: #4a3925;
}
.ci-note-lines .ci-red { color: #bb4a3c; }
.ci-call {
  display: flex;
  align-items: center;
  gap: 11px;
  margin-top: 19px;
  border: 2px solid #4e3558;
  border-radius: 15px;
  background: #4e3558;
  padding: 13px 14px;
  color: #fffaf1;
  text-decoration: none;
  transform: rotate(-0.6deg);
  box-shadow: 3px 4px 0 #d9c7e1;
  transition: transform 0.12s ease, box-shadow 0.12s ease, background 0.18s ease;
  animation: ci-scrap-in 0.6s 0.23s ease both;
}
.ci-call:hover { background: #5f426b; }
.ci-call:active {
  transform: rotate(-0.6deg) translate(2px, 3px);
  box-shadow: 0 0 0 #d9c7e1;
}
.ci-call-sticker {
  display: flex;
  height: 39px;
  width: 39px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #ffd98a;
  color: #6a4a1c;
}
.ci-call-copy { min-width: 0; flex: 1; }
.ci-call-label {
  display: block;
  font-size: 10px;
  color: #e3d3ea;
}
.ci-call-num {
  display: block;
  margin-top: 2px;
  font-size: 27px;
  font-weight: 800;
  line-height: 1.2;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-arrow {
  font-size: 18px;
  color: #ffd98a;
}
.ci-empty {
  margin-top: 19px;
  border: 2px dashed #a98db6;
  border-radius: 15px;
  padding: 15px;
  font-size: 12px;
  color: #6d4d80;
  text-align: center;
}
.ci-foot {
  margin-top: 16px;
  font-size: 12px;
  color: #8d7896;
  text-align: center;
  animation: ci-scrap-in 0.6s 0.29s ease both;
}
.ci-credit {
  margin-top: 15px;
  text-align: center;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.08em;
  color: #9a8a77;
  animation: ci-scrap-in 0.6s 0.34s ease both;
}
@keyframes ci-scrap-in {
  from { opacity: 0; transform: translateY(14px) rotate(-0.5deg); }
  to { opacity: 1; transform: translateY(0) rotate(-0.5deg); }
}
@media (prefers-reduced-motion: reduce) {
  .ci-scrap-board, .ci-title, .ci-lead, .ci-note, .ci-call, .ci-foot, .ci-credit { animation: none; }
  .ci-call { transition: none; }
}
</style>

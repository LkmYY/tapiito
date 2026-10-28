<!-- 求助信息卡 · 绘本童趣：薄荷与蜜桃色，像一页温柔的儿童绘本 -->
<template>
  <div class="ci-book min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-9">
      <div class="ci-cover">
        <span class="ci-sun"><Sun :size="18" /></span>
        <span class="ci-cloud ci-cloud--one"></span>
        <span class="ci-cloud ci-cloud--two"></span>

        <p class="ci-chapter">Chapter 01 · 请帮帮我</p>
        <h1 class="ci-title">我找不到<br />家人了</h1>
        <p class="ci-lead">好心人，可以帮我打个电话吗？<br />我的家人马上就来接我。</p>

        <div class="ci-scene" aria-hidden="true">
          <span class="ci-hill ci-hill--back"></span>
          <span class="ci-hill ci-hill--front"></span>
          <span class="ci-tree"></span>
        </div>
      </div>

      <div class="ci-page">
        <div class="ci-field">
          <span class="ci-icon"><UserRound :size="17" /></span>
          <div class="ci-field-body">
            <span class="ci-label">基本情况</span>
            <p class="ci-name">{{ payload.baseInfo }}</p>
          </div>
        </div>
        <div class="ci-field">
          <span class="ci-icon ci-icon--green"><School :size="17" /></span>
          <div class="ci-field-body">
            <span class="ci-label">学校 / 单位</span>
            <p class="ci-text">{{ payload.school }}</p>
          </div>
        </div>
        <div class="ci-field">
          <span class="ci-icon ci-icon--red"><Heart :size="17" /></span>
          <div class="ci-field-body">
            <span class="ci-label">过敏史 / 病史</span>
            <p class="ci-text">{{ payload.allergies || '无' }}</p>
          </div>
        </div>
      </div>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-bubble"><Phone :size="21" /></span>
        <span class="ci-call-copy">
          <span class="ci-call-label">点一下，打给家人</span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        </span>
        <span class="ci-call-heart" aria-hidden="true">♥</span>
      </a>

      <p class="ci-foot"><Sparkles :size="13" /> 谢谢你的帮助，你像小太阳一样温暖</p>
    </div>
  </div>
</template>

<script setup>
import { Heart, Phone, School, Sparkles, Sun, UserRound } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-book {
  background: #eef9f2;
  color: #23453a;
  font-family: "Comic Sans MS", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-cover {
  position: relative;
  overflow: hidden;
  border-radius: 30px 30px 18px 18px;
  background: #ffe6d5;
  padding: 27px 24px 0;
  box-shadow: 0 16px 34px -28px rgba(120, 72, 45, 0.7);
  animation: ci-book-in 0.65s ease both;
}
.ci-sun {
  position: absolute;
  right: 24px;
  top: 23px;
  display: flex;
  height: 39px;
  width: 39px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #ffc95e;
  color: #8a4b13;
}
.ci-cloud {
  position: absolute;
  height: 13px;
  width: 48px;
  border-radius: 13px;
  background: rgba(255, 255, 255, 0.75);
}
.ci-cloud--one { top: 39px; left: 20px; }
.ci-cloud--two { top: 74px; right: 76px; width: 33px; }
.ci-chapter {
  position: relative;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.18em;
  color: #b9673d;
}
.ci-title {
  position: relative;
  margin-top: 49px;
  font-size: 39px;
  font-weight: 800;
  line-height: 1.2;
  letter-spacing: -0.02em;
  color: #7a3f24;
}
.ci-lead {
  position: relative;
  margin-top: 12px;
  font-size: 14px;
  line-height: 1.85;
  color: #9b5d3c;
}
.ci-scene {
  position: relative;
  height: 85px;
  margin-top: 13px;
}
.ci-hill {
  position: absolute;
  bottom: -54px;
  border-radius: 50%;
}
.ci-hill--back {
  left: -30px;
  height: 116px;
  width: 250px;
  background: #bfe6cf;
}
.ci-hill--front {
  right: -60px;
  height: 106px;
  width: 280px;
  background: #a4d9bc;
}
.ci-tree {
  position: absolute;
  bottom: 27px;
  left: 54px;
  height: 10px;
  width: 7px;
  background: #90684b;
  border-radius: 3px;
}
.ci-tree::before {
  content: "";
  position: absolute;
  left: -13px;
  bottom: 7px;
  height: 30px;
  width: 33px;
  border-radius: 50% 50% 44% 44%;
  background: #6fbf92;
}
.ci-page {
  position: relative;
  margin-top: 13px;
  border-radius: 18px;
  background: #ffffff;
  padding: 8px 20px;
  box-shadow: 0 14px 30px -26px rgba(35, 69, 58, 0.55);
  animation: ci-book-in 0.6s 0.1s ease both;
}
.ci-field {
  display: flex;
  align-items: flex-start;
  gap: 11px;
  border-bottom: 2px dashed #dceee4;
  padding: 13px 0;
}
.ci-field:last-child { border-bottom: none; }
.ci-icon {
  display: flex;
  height: 34px;
  width: 34px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 11px;
  background: #ffeedd;
  color: #c26a35;
}
.ci-icon--green { background: #e3f5ea; color: #39845c; }
.ci-icon--red { background: #ffe7e4; color: #c25045; }
.ci-field-body { min-width: 0; }
.ci-label {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #7fa393;
}
.ci-name {
  margin-top: 2px;
  font-size: 17px;
  font-weight: 800;
  line-height: 1.5;
  color: #274b3e;
}
.ci-text {
  margin-top: 2px;
  font-size: 14px;
  font-weight: 600;
  line-height: 1.5;
  color: #3f6355;
}
.ci-call {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: 15px;
  border-radius: 23px;
  background: #2f7d5b;
  padding: 15px 17px;
  color: #ffffff;
  text-decoration: none;
  box-shadow: 0 16px 26px -20px rgba(47, 125, 91, 0.9);
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-book-in 0.6s 0.18s ease both;
}
.ci-call:hover { background: #286b4e; }
.ci-call:active { transform: scale(0.98) rotate(-0.4deg); }
.ci-call-bubble {
  display: flex;
  height: 43px;
  width: 43px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #ffd98a;
  color: #7a4810;
}
.ci-call-copy { min-width: 0; flex: 1; }
.ci-call-label {
  display: block;
  font-size: 11px;
  font-weight: 700;
  color: #c7ecd9;
}
.ci-call-num {
  display: block;
  margin-top: 2px;
  font-size: 27px;
  font-weight: 800;
  line-height: 1.15;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-heart {
  color: #ffd98a;
  font-size: 17px;
}
.ci-empty {
  margin-top: 15px;
  border: 2px dashed #9fc9b4;
  border-radius: 23px;
  padding: 16px;
  font-size: 14px;
  color: #4d7a67;
  text-align: center;
}
.ci-foot {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  margin-top: 16px;
  font-size: 11px;
  color: #6b9382;
  text-align: center;
  animation: ci-book-in 0.6s 0.26s ease both;
}
@keyframes ci-book-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-cover, .ci-page, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>

<template>
  <div class="ac-school-root relative min-h-screen overflow-hidden px-5 py-9 text-[#3a3346]">
    <div class="ac-school-bg"></div>
    <span class="ac-school-doodle ac-school-d1">✿</span>
    <span class="ac-school-doodle ac-school-d2">☆</span>
    <span class="ac-school-doodle ac-school-d3">♡</span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-school-page">
        <span class="ac-school-tape"></span>
        <div class="ac-school-holes">
          <i v-for="n in 7" :key="n"></i>
        </div>

        <header class="ac-school-head">
          <span>星見学園 · 生徒手帳</span>
          <span>NO. {{ nfcId || '2024' }}</span>
        </header>

        <div class="ac-school-id">
          <div class="ac-school-photo">
            <span>{{ initial }}</span>
            <i class="ac-school-photo-tape"></i>
          </div>
          <div class="ac-school-info">
            <p class="ac-school-label">NAME</p>
            <h1 class="ac-school-name">{{ p.name || '未命名角色' }}</h1>
            <p class="ac-school-title">{{ p.title }}</p>
            <p class="ac-school-series">所属 · {{ p.series || '未知作品' }}</p>
          </div>
        </div>

        <div class="ac-school-highlight">
          <span class="ac-school-marker">属性</span>
          <b>{{ p.element }}</b>
          <span class="ac-school-star">★</span>
        </div>

        <div class="ac-school-note">
          <span class="ac-school-pin"></span>
          <p class="ac-school-note-label">放課後のひとこと</p>
          <p class="ac-school-quote">{{ p.quote || '……今天也没有把想说的话说出口。' }}</p>
        </div>

        <div class="ac-school-bottom">
          <div class="ac-school-social">
            <Heart :size="13" />
            <div>
              <small>SNS ACCOUNT</small>
              <b>@{{ p.social }}</b>
            </div>
          </div>
          <div class="ac-school-stamp">出席<br />OK</div>
        </div>

        <footer class="ac-school-foot">
          <span>月曜日 · 放課後 16:40</span>
          <span>一起去天台吧？</span>
        </footer>
      </div>

      <p class="ac-school-caption">—— 青春は、いまこの瞬間に。</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Heart } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const initial = computed(() => (p.value.name || '？').toString().trim().slice(0, 1) || '？')
</script>

<style scoped>
.ac-school-root {
  background: #fbeff2;
  font-family: 'Kaiti SC', 'STKaiti', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-school-bg {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(147, 197, 253, 0.24) 1px, transparent 1px),
    linear-gradient(90deg, transparent 27px, rgba(248, 113, 113, 0.18) 27px 28px, transparent 28px),
    radial-gradient(circle at 88% 8%, rgba(191, 219, 254, 0.6), transparent 30%),
    radial-gradient(circle at 6% 92%, rgba(251, 207, 232, 0.7), transparent 34%);
  background-size: 100% 26px, 100% 100%, 100% 100%, 100% 100%;
}
.ac-school-doodle {
  position: absolute;
  color: rgba(244, 114, 182, 0.45);
  font-size: 22px;
  animation: ac-school-float 6s ease-in-out infinite;
}
.ac-school-d1 { top: 9%; left: 7%; }
.ac-school-d2 { top: 26%; right: 9%; font-size: 16px; animation-delay: 1.5s; }
.ac-school-d3 { bottom: 12%; right: 12%; animation-delay: 2.4s; }
.ac-school-page {
  position: relative;
  border-radius: 10px;
  border: 1px solid #e8d6dd;
  background: #fffdf8;
  box-shadow: 0 20px 46px -24px rgba(190, 24, 93, 0.32);
  padding: 20px 20px 16px 36px;
  animation: ac-school-enter 0.9s cubic-bezier(0.2, 1, 0.3, 1) both;
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.ac-school-page:hover {
  transform: translateY(-4px) rotate(-0.4deg);
  box-shadow: 0 28px 56px -22px rgba(190, 24, 93, 0.4);
}
.ac-school-tape {
  position: absolute;
  top: -10px;
  left: 50%;
  width: 92px;
  height: 24px;
  transform: translateX(-50%) rotate(-2deg);
  background: repeating-linear-gradient(45deg, rgba(253, 186, 116, 0.65) 0 8px, rgba(254, 215, 170, 0.65) 8px 16px);
  opacity: 0.85;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08);
}
.ac-school-holes {
  position: absolute;
  top: 28px;
  bottom: 28px;
  left: 12px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.ac-school-holes i {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: #f3e4e9;
  box-shadow: inset 0 1px 2px rgba(120, 80, 90, 0.18);
}
.ac-school-head {
  display: flex;
  justify-content: space-between;
  border-bottom: 1.5px solid #efd8df;
  padding-bottom: 8px;
  font-size: 8px;
  letter-spacing: 0.14em;
  color: #a98e99;
}
.ac-school-id {
  display: flex;
  gap: 14px;
  margin-top: 18px;
}
.ac-school-photo {
  position: relative;
  display: flex;
  flex: none;
  width: 88px;
  height: 108px;
  align-items: center;
  justify-content: center;
  border: 4px solid #fff;
  background: linear-gradient(160deg, #dbeafe, #fce7f3 58%, #fef3c7);
  box-shadow: 0 3px 0 #efd8df, 0 5px 12px rgba(190, 24, 93, 0.12);
  transform: rotate(-3deg);
}
.ac-school-photo span {
  font-size: 42px;
  font-weight: 700;
  color: #7c5a6b;
  text-shadow: 2px 2px 0 #fff;
}
.ac-school-photo-tape {
  position: absolute;
  top: -9px;
  left: -9px;
  width: 34px;
  height: 13px;
  background: rgba(147, 197, 253, 0.7);
  transform: rotate(-38deg);
}
.ac-school-info { min-width: 0; flex: 1; }
.ac-school-label {
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.3em;
  color: #d8a7b6;
}
.ac-school-name {
  margin-top: 2px;
  font-size: 33px;
  font-weight: 700;
  line-height: 1.12;
  letter-spacing: 0.06em;
  color: #4a3a4c;
  overflow-wrap: anywhere;
}
.ac-school-title {
  display: inline-block;
  margin-top: 6px;
  padding: 1px 7px;
  background: linear-gradient(180deg, transparent 52%, rgba(253, 224, 71, 0.62) 52%);
  font-size: 13px;
  color: #7a5c63;
}
.ac-school-series {
  margin-top: 9px;
  font-size: 10.5px;
  letter-spacing: 0.06em;
  color: #9b8590;
}
.ac-school-highlight {
  display: flex;
  align-items: center;
  gap: 9px;
  margin-top: 18px;
  border: 1.5px dashed #f2c5d3;
  border-radius: 8px;
  background: #fff6f9;
  padding: 9px 12px;
}
.ac-school-marker {
  border-radius: 3px;
  background: #f9a8d4;
  padding: 2px 6px;
  color: #fff;
  font-size: 8px;
  letter-spacing: 0.16em;
}
.ac-school-highlight b {
  flex: 1;
  font-size: 13px;
  letter-spacing: 0.08em;
  color: #be185d;
  overflow-wrap: anywhere;
}
.ac-school-star { color: #fbbf24; font-size: 12px; }
.ac-school-note {
  position: relative;
  margin-top: 18px;
  border: 1px solid #f7e0ae;
  background: #fffbe8;
  padding: 15px 14px 13px;
  box-shadow: 2px 3px 0 rgba(251, 191, 36, 0.18);
  transform: rotate(-0.6deg);
}
.ac-school-pin {
  position: absolute;
  top: -7px;
  left: 50%;
  width: 44px;
  height: 14px;
  transform: translateX(-50%) rotate(1.5deg);
  background: rgba(196, 181, 253, 0.55);
}
.ac-school-note-label {
  font-size: 8.5px;
  letter-spacing: 0.18em;
  color: #c08b3c;
}
.ac-school-quote {
  margin-top: 5px;
  font-size: 15px;
  line-height: 1.85;
  color: #5f4a3f;
  overflow-wrap: anywhere;
}
.ac-school-bottom {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  margin-top: 18px;
  border-top: 1.5px solid #f2e3e8;
  padding-top: 13px;
}
.ac-school-social {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #db2777;
}
.ac-school-social small {
  display: block;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: #c3a4ae;
}
.ac-school-social b {
  display: block;
  margin-top: 1px;
  font-family: Consolas, monospace;
  font-size: 11.5px;
  font-weight: 600;
  color: #8b5a6d;
  overflow-wrap: anywhere;
}
.ac-school-stamp {
  display: flex;
  flex: none;
  width: 48px;
  height: 48px;
  align-items: center;
  justify-content: center;
  border: 2px solid #f472b6;
  border-radius: 50%;
  color: #db2777;
  font-size: 9px;
  font-weight: 700;
  line-height: 1.25;
  letter-spacing: 0.08em;
  text-align: center;
  transform: rotate(9deg);
  opacity: 0.8;
}
.ac-school-foot {
  display: flex;
  justify-content: space-between;
  margin-top: 15px;
  font-size: 8.5px;
  letter-spacing: 0.06em;
  color: #b9a1aa;
}
.ac-school-caption {
  margin-top: 22px;
  text-align: center;
  font-size: 11px;
  letter-spacing: 0.14em;
  color: #b58a9c;
  animation: ac-school-fade 1.4s 0.7s both;
}
@keyframes ac-school-enter {
  from { opacity: 0; transform: translateY(30px) rotate(1.6deg); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-school-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-school-float {
  0%, 100% { transform: translateY(0) rotate(-4deg); opacity: 0.55; }
  50% { transform: translateY(-8px) rotate(4deg); opacity: 0.9; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-school-root *,
  .ac-school-root *::before,
  .ac-school-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

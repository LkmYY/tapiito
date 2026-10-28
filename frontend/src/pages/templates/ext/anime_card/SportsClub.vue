<template>
  <div class="ac-sport-root min-h-screen overflow-hidden px-4 py-9 text-[#103326]">
    <div class="ac-sport-bg"></div>
    <div class="ac-sport-halftone"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-sport-poster">
        <header class="ac-sport-head">
          <span class="ac-sport-crest">
            <Trophy :size="15" aria-hidden="true" />
          </span>
          <div class="ac-sport-head-text">
            <p class="ac-sport-club">疾风运动社 · GALE ATHLETIC CLUB</p>
            <p class="ac-sport-sub">MEMBER CARD / 社团成员证</p>
          </div>
          <span class="ac-sport-season">26<br /><i>S</i></span>
        </header>

        <div class="ac-sport-track">
          <svg class="ac-sport-lanes" viewBox="0 0 400 210" preserveAspectRatio="none" aria-hidden="true">
            <path d="M-40 178 C 90 138, 210 154, 440 96" fill="none" stroke="#ffffff" stroke-width="34" opacity="0.5" />
            <path d="M-40 178 C 90 138, 210 154, 440 96" fill="none" stroke="#0f5a3c" stroke-width="1.5" stroke-dasharray="14 10" opacity="0.5" />
            <path d="M-40 204 C 90 164, 226 182, 440 122" fill="none" stroke="#ffffff" stroke-width="24" opacity="0.42" />
            <path d="M-40 226 C 96 190, 234 208, 440 150" fill="none" stroke="#ff7a2f" stroke-width="3" opacity="0.55" />
          </svg>

          <div class="ac-sport-bib">
            <span class="ac-sport-pin ac-sport-pin-tl"></span>
            <span class="ac-sport-pin ac-sport-pin-tr"></span>
            <span class="ac-sport-pin ac-sport-pin-bl"></span>
            <span class="ac-sport-pin ac-sport-pin-br"></span>
            <p class="ac-sport-bib-label">MEMBER NO.</p>
            <p class="ac-sport-bib-no">{{ bib }}</p>
          </div>

          <div class="ac-sport-nameblock">
            <p class="ac-sport-kicker">ROOKIE OF THE SEASON</p>
            <h1 class="ac-sport-name">{{ p.name || '未报名选手' }}</h1>
            <p class="ac-sport-title">{{ p.title || '称号待定' }}</p>
          </div>
        </div>

        <div class="ac-sport-score">
          <div class="ac-sport-score-head">
            <span>TEAM DATA</span>
            <span class="ac-sport-score-dots">PLAYER PROFILE</span>
            <Zap :size="11" aria-hidden="true" />
          </div>
          <div class="ac-sport-score-row">
            <em>01</em>
            <span>所属作品</span>
            <b>{{ p.series || '未知队伍' }}</b>
          </div>
          <div class="ac-sport-score-row">
            <em>02</em>
            <span>擅长属性</span>
            <b class="ac-sport-accent">{{ p.element || '潜能未测' }}</b>
          </div>
        </div>

        <blockquote class="ac-sport-quote">
          <span class="ac-sport-quote-mark">“</span>
          <p>{{ p.quote || '再跑一圈，终点线就是我的起跑线。' }}</p>
        </blockquote>

        <div class="ac-sport-finish" aria-hidden="true">
          <span v-for="n in 7" :key="n" class="ac-sport-chevron"></span>
        </div>

        <footer class="ac-sport-foot">
          <div class="ac-sport-social">
            <AtSign :size="11" aria-hidden="true" />
            <span>{{ p.social || '未登记账号' }}</span>
          </div>
          <span class="ac-sport-id">CARD ID {{ nfcId || 'PREVIEW' }}</span>
        </footer>
      </div>

      <p class="ac-sport-caption">RUN FAST · 青春不设终点线</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { AtSign, Trophy, Zap } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const bib = computed(() => {
  const digits = String(props.nfcId || '').replace(/\D/g, '')
  if (!digits) return '07'
  return digits.slice(-2).padStart(2, '0')
})
</script>

<style scoped>
.ac-sport-root {
  position: relative;
  background: radial-gradient(circle at 88% 6%, rgba(255, 154, 77, 0.2), transparent 34%), radial-gradient(circle at 6% 94%, rgba(31, 122, 82, 0.14), transparent 42%), linear-gradient(165deg, #f8faf1, #eef5e6 58%, #e2efdb);
  font-family: 'PingFang SC', 'Microsoft YaHei', 'HarmonyOS Sans SC', sans-serif;
}
.ac-sport-bg {
  position: absolute; inset: 0;
  background: repeating-linear-gradient(115deg, rgba(15, 90, 60, 0.05) 0 2px, transparent 2px 34px), linear-gradient(180deg, rgba(255, 255, 255, 0.6), transparent 30%);
}
.ac-sport-halftone {
  position: absolute; right: -60px; top: 120px; width: 260px; height: 260px; background-image: radial-gradient(rgba(255, 122, 47, 0.35) 1.2px, transparent 1.2px);
  background-size: 9px 9px; transform: rotate(12deg); mask-image: radial-gradient(circle, #000, transparent 70%);
  -webkit-mask-image: radial-gradient(circle, #000, transparent 70%);
}
.ac-sport-poster {
  position: relative; overflow: hidden; border-radius: 20px 20px 20px 6px; background: #ffffff;
  box-shadow: 0 30px 60px -26px rgba(11, 68, 45, 0.5), 0 1px 0 rgba(255, 255, 255, 0.9) inset; animation: ac-sport-pop 0.85s cubic-bezier(0.18, 1.1, 0.3, 1) both;
}
.ac-sport-poster::before { content: ''; position: absolute; left: 0; top: 0; width: 9px; height: 100%; background: linear-gradient(180deg, #0f5a3c 0 58%, #ff7a2f 58% 100%); }
.ac-sport-head { display: flex; align-items: center; gap: 10px; padding: 18px 16px 14px 24px; background: linear-gradient(120deg, #0f5a3c, #147a4f 70%, #1c8a56); color: #eaf6ec; }
.ac-sport-crest {
  display: grid; flex: none; place-items: center; width: 36px; height: 36px; border: 1.5px solid rgba(255, 255, 255, 0.65); border-radius: 50%;
  background: rgba(255, 255, 255, 0.08); color: #ffd9a8;
}
.ac-sport-head-text { flex: 1; min-width: 0; }
.ac-sport-club { font-size: 10px; font-weight: 800; letter-spacing: 0.06em; }
.ac-sport-sub { margin-top: 3px; font-size: 6.5px; letter-spacing: 0.24em; color: rgba(226, 245, 228, 0.72); }
.ac-sport-season {
  flex: none; border-radius: 6px; background: #ff7a2f; padding: 5px 8px 4px; text-align: center; color: #ffffff; font-family: Consolas, 'Courier New', monospace; font-size: 15px;
  font-weight: 700; font-style: italic; line-height: 1;
}
.ac-sport-season i { font-size: 7px; font-style: normal; letter-spacing: 0.1em; }
.ac-sport-track { position: relative; padding: 20px 18px 10px 24px; background: linear-gradient(180deg, #f2f8ec, #ffffff 62%); }
.ac-sport-lanes { position: absolute; left: 0; right: 0; bottom: 6px; width: 100%; height: 172px; }
.ac-sport-bib {
  position: relative; width: 116px; margin-left: 2px; border: 2px solid #0f5a3c; border-radius: 10px;
  background: repeating-linear-gradient(45deg, rgba(15, 90, 60, 0.05) 0 7px, transparent 7px 14px), #ffffff; padding: 10px 8px 9px; text-align: center;
  box-shadow: 0 8px 18px -10px rgba(12, 62, 42, 0.55); transform: rotate(-2.4deg); animation: ac-sport-bib 4.5s ease-in-out infinite;
}
.ac-sport-pin {
  position: absolute; width: 7px; height: 7px; border-radius: 50%; background: radial-gradient(circle at 35% 30%, #ffffff, #b9c3bd 45%, #6d7a72);
  box-shadow: 0 1px 2px rgba(20, 40, 30, 0.45);
}
.ac-sport-pin-tl { top: -4px; left: 12px; }
.ac-sport-pin-tr { top: -4px; right: 12px; }
.ac-sport-pin-bl { bottom: -4px; left: 12px; }
.ac-sport-pin-br { bottom: -4px; right: 12px; }
.ac-sport-bib-label { font-size: 6px; font-weight: 700; letter-spacing: 0.24em; color: #4d7a63; }
.ac-sport-bib-no {
  margin-top: 2px; color: #0f5a3c; font-family: 'Arial Black', Impact, Consolas, sans-serif; font-size: 44px; font-weight: 900; font-style: italic; line-height: 0.95;
  letter-spacing: -0.02em;
}
.ac-sport-nameblock { position: relative; margin-top: -40px; padding-left: 128px; }
.ac-sport-kicker { display: inline-block; background: #ff7a2f; padding: 3px 7px 2px; color: #ffffff; font-size: 6px; font-weight: 800; letter-spacing: 0.18em; transform: skewX(-8deg); }
.ac-sport-name { margin-top: 7px; color: #0c4630; font-size: 31px; font-weight: 900; line-height: 1.02; letter-spacing: -0.01em; overflow-wrap: anywhere; }
.ac-sport-title {
  display: inline-block; margin-top: 6px; border-bottom: 3px solid #ffb27f; padding-bottom: 2px; color: #2f6a4e; font-size: 11px; font-weight: 700; letter-spacing: 0.1em;
  overflow-wrap: anywhere;
}
.ac-sport-score {
  margin: 16px 18px 0 24px; border-radius: 10px;
  background: repeating-linear-gradient(0deg, rgba(255, 255, 255, 0.05) 0 1px, transparent 1px 4px), linear-gradient(140deg, #0c3d2a, #12513a); padding: 11px 13px 9px;
  box-shadow: 0 12px 24px -14px rgba(10, 60, 40, 0.8);
}
.ac-sport-score-head {
  display: flex; align-items: center; gap: 8px; border-bottom: 1px solid rgba(190, 235, 205, 0.25); padding-bottom: 7px; color: #ffd9a8;
  font-family: Consolas, 'Courier New', monospace; font-size: 7px; letter-spacing: 0.2em;
}
.ac-sport-score-head svg { margin-left: auto; color: #ff9a4d; }
.ac-sport-score-dots { color: rgba(190, 235, 205, 0.55); }
.ac-sport-score-row { display: flex; align-items: baseline; gap: 9px; padding: 8px 0 6px; }
.ac-sport-score-row + .ac-sport-score-row { border-top: 1px dashed rgba(190, 235, 205, 0.18); }
.ac-sport-score-row em { flex: none; color: #ff9a4d; font-family: Consolas, monospace; font-size: 9px; font-style: normal; }
.ac-sport-score-row span { flex: none; width: 52px; color: rgba(214, 240, 222, 0.62); font-size: 8px; letter-spacing: 0.12em; }
.ac-sport-score-row b {
  min-width: 0; flex: 1; text-align: right; color: #f2fbf4; font-family: Consolas, 'Courier New', monospace; font-size: 12px; letter-spacing: 0.04em; overflow-wrap: anywhere;
}
.ac-sport-score-row .ac-sport-accent { color: #ffb27f; }
.ac-sport-quote { position: relative; margin: 17px 18px 0 24px; border-left: 3px solid #ff7a2f; padding: 4px 8px 4px 13px; }
.ac-sport-quote-mark { position: absolute; left: 6px; top: -9px; color: #ffc9a3; font-family: Georgia, serif; font-size: 28px; line-height: 1; }
.ac-sport-quote p { color: #2c5744; font-size: 12.5px; font-style: italic; line-height: 1.75; letter-spacing: 0.02em; overflow-wrap: anywhere; }
.ac-sport-finish { display: flex; gap: 8px; overflow: hidden; margin-top: 16px; padding-left: 24px; opacity: 0.85; }
.ac-sport-chevron {
  flex: none; width: 20px; height: 20px; border-top: 3px solid #0f5a3c; border-right: 3px solid #0f5a3c; transform: rotate(45deg) skewX(-12deg);
  animation: ac-sport-run 1.6s linear infinite;
}
.ac-sport-chevron:nth-child(2) { animation-delay: 0.16s; border-color: #178a5c; }
.ac-sport-chevron:nth-child(3) { animation-delay: 0.32s; }
.ac-sport-chevron:nth-child(4) { animation-delay: 0.48s; border-color: #178a5c; }
.ac-sport-chevron:nth-child(5) { animation-delay: 0.64s; border-color: #ff7a2f; }
.ac-sport-chevron:nth-child(6) { animation-delay: 0.8s; border-color: #ff9a4d; }
.ac-sport-chevron:nth-child(7) { animation-delay: 0.96s; border-color: #ffb27f; }
.ac-sport-foot {
  display: flex; align-items: center; justify-content: space-between; gap: 10px; margin-top: 14px; background: linear-gradient(120deg, #0f5a3c, #157048);
  padding: 11px 18px 11px 24px; color: #eaf6ec;
}
.ac-sport-social { display: inline-flex; align-items: center; gap: 5px; min-width: 0; font-family: Consolas, 'Courier New', monospace; font-size: 10.5px; font-weight: 700; }
.ac-sport-social span { overflow-wrap: anywhere; }
.ac-sport-social svg { flex: none; color: #ffb27f; }
.ac-sport-id {
  flex: none; border: 1px solid rgba(255, 255, 255, 0.4); border-radius: 4px; padding: 3px 6px 2px; font-family: Consolas, monospace; font-size: 7px; letter-spacing: 0.1em;
  color: rgba(234, 246, 236, 0.9);
}
.ac-sport-caption {
  margin-top: 22px; text-align: center; color: rgba(16, 61, 42, 0.5); font-size: 8px; font-weight: 700; letter-spacing: 0.34em; animation: ac-sport-fade 1.3s 0.6s both;
}
@keyframes ac-sport-pop {
  from { opacity: 0; transform: translateY(26px) scale(0.97); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-sport-bib {
  0%, 100% { transform: rotate(-2.4deg); }
  50% { transform: rotate(-0.6deg) translateY(-2px); }
}
@keyframes ac-sport-run {
  0% { transform: rotate(45deg) skewX(-12deg) translate(6px, 6px); opacity: 0; }
  35% { opacity: 1; }
  100% { transform: rotate(45deg) skewX(-12deg) translate(-16px, -16px); opacity: 0; }
}
@keyframes ac-sport-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-sport-root *, .ac-sport-root *::before, .ac-sport-root *::after { animation: none !important; transition: none !important; }
}
@media (max-width: 390px) {
  .ac-sport-nameblock { padding-left: 0; margin-top: 14px; }
  .ac-sport-name { font-size: 27px; }
}
</style>

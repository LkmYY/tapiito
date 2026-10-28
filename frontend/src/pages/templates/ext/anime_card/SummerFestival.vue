<template>
  <div class="ac-fest-root min-h-screen overflow-hidden px-4 py-8 text-[#f8ecd4]">
    <div class="ac-fest-sky"></div>
    <div class="ac-fest-stars"></div>
    <span class="ac-fest-fw ac-fest-fw-a"></span>
    <span class="ac-fest-fw ac-fest-fw-b"></span>
    <span class="ac-fest-fw ac-fest-fw-c"></span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-fest-lantern-row" aria-hidden="true">
        <span class="ac-fest-lantern ac-fest-lantern-a">祭</span>
        <span class="ac-fest-lantern ac-fest-lantern-b">夏</span>
        <span class="ac-fest-lantern ac-fest-lantern-c">夜</span>
      </div>

      <div class="ac-fest-card">
        <span class="ac-fest-corner ac-fest-corner-tl"></span>
        <span class="ac-fest-corner ac-fest-corner-br"></span>

        <header class="ac-fest-head">
          <span class="ac-fest-seal">祭</span>
          <div class="ac-fest-heading">
            <p class="ac-fest-kicker">SUMMER NIGHT FESTIVAL</p>
            <p class="ac-fest-sub">夏祭り · 浴衣記念</p>
          </div>
          <span class="ac-fest-ticket">NO. {{ nfcId || 'PREVIEW' }}</span>
        </header>

        <div class="ac-fest-art">
          <span class="ac-fest-moon"></span>
          <span class="ac-fest-cloud ac-fest-cloud-a"></span>
          <span class="ac-fest-cloud ac-fest-cloud-b"></span>
          <span class="ac-fest-art-fw ac-fest-art-fw-a"></span>
          <span class="ac-fest-art-fw ac-fest-art-fw-b"></span>

          <svg class="ac-fest-fish ac-fest-fish-a" viewBox="0 0 100 64" aria-hidden="true">
            <defs>
              <linearGradient id="festFishA" x1="0" y1="0" x2="1" y2="1">
                <stop offset="0" stop-color="#ffcf7c" />
                <stop offset="0.45" stop-color="#ff8a3d" />
                <stop offset="1" stop-color="#c9351f" />
              </linearGradient>
            </defs>
            <path d="M25 33 C35 17, 63 15, 75 29 C85 40, 78 52, 62 55 C44 58, 31 47, 25 33 Z" fill="url(#festFishA)" />
            <path d="M29 32 C17 24, 11 18, 5 15 C9 27, 9 37, 5 49 C14 45, 22 41, 30 38 Z" fill="#ff8a3d" opacity="0.9" />
            <path d="M47 52 C50 58, 56 61, 61 61 C57 55, 55 52, 54 48 Z" fill="#ffcf7c" opacity="0.75" />
            <path d="M52 24 C58 18, 64 17, 69 19 C65 23, 62 26, 60 29 Z" fill="#ffe7b5" opacity="0.6" />
            <circle cx="68" cy="33" r="2.2" fill="#32130d" />
            <circle cx="68.8" cy="32.2" r="0.7" fill="#fff6df" />
          </svg>

          <svg class="ac-fest-fish ac-fest-fish-b" viewBox="0 0 100 64" aria-hidden="true">
            <defs>
              <linearGradient id="festFishB" x1="0" y1="0" x2="1" y2="1">
                <stop offset="0" stop-color="#ffe1a1" />
                <stop offset="0.5" stop-color="#ffb24f" />
                <stop offset="1" stop-color="#d7502a" />
              </linearGradient>
            </defs>
            <path d="M25 33 C35 17, 63 15, 75 29 C85 40, 78 52, 62 55 C44 58, 31 47, 25 33 Z" fill="url(#festFishB)" />
            <path d="M29 32 C17 24, 11 18, 5 15 C9 27, 9 37, 5 49 C14 45, 22 41, 30 38 Z" fill="#d7502a" opacity="0.85" />
            <circle cx="68" cy="33" r="2" fill="#32130d" />
          </svg>

          <span class="ac-fest-art-line"></span>
          <span class="ac-fest-art-label">夜空に咲く、一瞬の夏</span>
        </div>

        <div class="ac-fest-nameblock">
          <p class="ac-fest-small">今宵、君と。</p>
          <h1 class="ac-fest-name">{{ p.name || '未署名の浴衣' }}</h1>
          <p class="ac-fest-title">{{ p.title || '称号未定' }}</p>
        </div>

        <div class="ac-fest-meta">
          <div>
            <span>出自作品</span>
            <b>{{ p.series || '夏夜的某处' }}</b>
          </div>
          <div>
            <span>属性</span>
            <b class="ac-fest-accent">{{ p.element || '花火 · 金鱼' }}</b>
          </div>
        </div>

        <blockquote class="ac-fest-quote">
          <span class="ac-fest-quote-mark">“</span>
          <p>{{ p.quote || '焰火落下之前，能不能陪我多走一段路？' }}</p>
        </blockquote>

        <footer class="ac-fest-foot">
          <div class="ac-fest-social">
            <AtSign :size="11" aria-hidden="true" />
            <span>{{ p.social || '未登记账号' }}</span>
          </div>
          <span class="ac-fest-id">提灯番号 · {{ nfcId || 'PREVIEW' }}</span>
        </footer>
      </div>

      <p class="ac-fest-caption">花火は一瞬、夏は永遠。</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { AtSign } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-fest-root {
  position: relative;
  background: radial-gradient(ellipse at 50% -10%, rgba(45, 70, 145, 0.55), transparent 56%), radial-gradient(ellipse at 8% 100%, rgba(126, 31, 48, 0.25), transparent 42%), linear-gradient(180deg, #081128 0%, #0b1738 46%, #070d20 100%);
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-fest-sky {
  position: absolute; inset: 0;
  background: radial-gradient(circle at 24% 16%, rgba(255, 210, 137, 0.1), transparent 7%), radial-gradient(circle at 82% 36%, rgba(255, 120, 92, 0.08), transparent 9%), linear-gradient(180deg, transparent 48%, rgba(3, 7, 18, 0.55));
}
.ac-fest-stars {
  position: absolute; inset: 0;
  background-image: radial-gradient(rgba(255, 246, 218, 0.85) 0.7px, transparent 0.8px), radial-gradient(rgba(255, 218, 160, 0.5) 0.6px, transparent 0.7px);
  background-position: 0 0, 13px 17px; background-size: 41px 47px, 59px 61px; opacity: 0.55; animation: ac-fest-twinkle 5s ease-in-out infinite;
}
.ac-fest-fw {
  position: absolute; width: 250px; height: 250px; border-radius: 50%;
  background: repeating-conic-gradient(from 4deg, rgba(255, 205, 120, 0.85) 0deg 1.1deg, transparent 1.1deg 14deg);
  mask-image: radial-gradient(circle, transparent 0 21%, #000 33% 58%, transparent 66%);
  -webkit-mask-image: radial-gradient(circle, transparent 0 21%, #000 33% 58%, transparent 66%); opacity: 0; animation: ac-fest-burst 7s ease-out infinite;
}
.ac-fest-fw-a { top: 7%; left: -98px; }
.ac-fest-fw-b {
  top: 2%; right: -104px; width: 300px; height: 300px; background: repeating-conic-gradient(from 2deg, rgba(255, 138, 95, 0.8) 0deg 0.9deg, transparent 0.9deg 12deg);
  animation-delay: -2.7s;
}
.ac-fest-fw-c {
  top: 48%; right: -150px; width: 230px; height: 230px; background: repeating-conic-gradient(from 7deg, rgba(255, 226, 156, 0.7) 0deg 0.8deg, transparent 0.8deg 16deg);
  animation-delay: -5s; opacity: 0.35;
}
.ac-fest-lantern-row { position: relative; height: 86px; z-index: 2; }
.ac-fest-lantern {
  position: absolute; top: 0; display: grid; place-items: center; width: 46px; height: 58px; border-top: 4px solid #6f1712; border-bottom: 4px solid #6f1712;
  border-radius: 50% / 40%;
  background: repeating-linear-gradient(90deg, transparent 0 6px, rgba(97, 20, 14, 0.28) 6px 7px), radial-gradient(ellipse at 50% 42%, #ffc46c 0 12%, #f47a36 38%, #c3341d 72%, #7a1b14 100%);
  color: #5c160f; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 19px; font-weight: 700; text-shadow: 0 1px 0 rgba(255, 232, 176, 0.45);
  box-shadow: 0 0 24px rgba(255, 132, 57, 0.34), 0 8px 18px -8px rgba(0, 0, 0, 0.7); transform-origin: top center; animation: ac-fest-sway 4.8s ease-in-out infinite;
}
.ac-fest-lantern::before { content: ''; position: absolute; left: 50%; bottom: 100%; width: 1px; height: 30px; background: rgba(255, 212, 144, 0.48); }
.ac-fest-lantern::after { content: ''; position: absolute; left: 50%; top: 100%; width: 2px; height: 18px; background: linear-gradient(180deg, #ffce7b, rgba(255, 206, 123, 0.1)); }
.ac-fest-lantern-a { left: 8px; margin-top: 22px; }
.ac-fest-lantern-b { left: 50%; margin-left: -23px; margin-top: 4px; width: 50px; height: 64px; animation-delay: -1.6s; }
.ac-fest-lantern-c { right: 7px; margin-top: 30px; width: 40px; height: 52px; animation-delay: -3s; opacity: 0.9; }
.ac-fest-card {
  position: relative; overflow: hidden; margin-top: -2px; border: 1px solid rgba(224, 177, 99, 0.6); border-radius: 3px;
  background: linear-gradient(135deg, rgba(255, 214, 137, 0.06), transparent 34%), linear-gradient(180deg, rgba(17, 32, 74, 0.98), rgba(7, 13, 33, 0.99));
  box-shadow: 0 36px 72px -26px rgba(0, 0, 0, 0.9), 0 0 50px rgba(42, 70, 152, 0.18), inset 0 0 60px rgba(0, 0, 0, 0.35);
  animation: ac-fest-rise 0.95s cubic-bezier(0.19, 1, 0.22, 1) both;
}
.ac-fest-card::before { content: ''; position: absolute; inset: 7px; border: 1px solid rgba(224, 177, 99, 0.18); pointer-events: none; z-index: 2; }
.ac-fest-card::after {
  content: ''; position: absolute; left: 0; top: 0; bottom: 0; width: 5px;
  background: repeating-linear-gradient(180deg, #c99a4d 0 9px, transparent 9px 18px), rgba(255, 213, 138, 0.12); opacity: 0.6;
}
.ac-fest-corner { position: absolute; z-index: 3; width: 22px; height: 22px; border-color: #dcb36e; opacity: 0.85; }
.ac-fest-corner-tl { top: 13px; left: 15px; border-top: 1px solid; border-left: 1px solid; }
.ac-fest-corner-br { right: 15px; bottom: 13px; border-right: 1px solid; border-bottom: 1px solid; }
.ac-fest-head { position: relative; z-index: 3; display: flex; align-items: center; gap: 10px; padding: 22px 20px 13px 23px; }
.ac-fest-seal {
  display: grid; flex: none; place-items: center; width: 33px; height: 33px; border: 1px solid #e4b468; background: rgba(169, 38, 32, 0.22); color: #f0c67d;
  font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 19px; font-weight: 700; transform: rotate(-3deg);
}
.ac-fest-heading { flex: 1; min-width: 0; }
.ac-fest-kicker { color: #e4b468; font-family: Consolas, 'Courier New', monospace; font-size: 7px; letter-spacing: 0.2em; }
.ac-fest-sub { margin-top: 3px; color: #cdd5eb; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 10px; letter-spacing: 0.16em; }
.ac-fest-ticket {
  flex: none; border: 1px solid rgba(228, 180, 104, 0.4); padding: 3px 5px 2px; color: #b9a476; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.06em;
  white-space: nowrap;
}
.ac-fest-art {
  position: relative; overflow: hidden; margin: 0 20px 0 23px; height: 146px; border-top: 1px solid rgba(228, 180, 104, 0.22); border-bottom: 1px solid rgba(228, 180, 104, 0.22);
  background: radial-gradient(circle at 76% 24%, rgba(255, 218, 144, 0.12), transparent 23%), linear-gradient(180deg, rgba(23, 43, 96, 0.8), rgba(7, 14, 36, 0.2));
}
.ac-fest-moon {
  position: absolute; top: 17px; right: 29px; width: 44px; height: 44px; border-radius: 50%; background: radial-gradient(circle at 38% 35%, #fff2ce, #f4cf8d 62%, #d9a559);
  box-shadow: 0 0 26px rgba(255, 216, 143, 0.34); opacity: 0.92;
}
.ac-fest-moon::after { content: ''; position: absolute; inset: 0; border-radius: 50%; background: radial-gradient(circle at 68% 28%, rgba(208, 165, 96, 0.32) 0 17%, transparent 18%); }
.ac-fest-cloud { position: absolute; height: 8px; border-radius: 50%; background: rgba(118, 139, 190, 0.16); filter: blur(3px); }
.ac-fest-cloud-a { top: 72px; right: -10px; width: 112px; }
.ac-fest-cloud-b { top: 102px; right: 78px; width: 78px; opacity: 0.65; }
.ac-fest-art-fw {
  position: absolute; width: 100px; height: 100px; border-radius: 50%;
  background: repeating-conic-gradient(from 0deg, rgba(255, 198, 112, 0.65) 0deg 0.9deg, transparent 0.9deg 18deg);
  mask-image: radial-gradient(circle, transparent 0 18%, #000 32% 60%, transparent 68%);
  -webkit-mask-image: radial-gradient(circle, transparent 0 18%, #000 32% 60%, transparent 68%);
}
.ac-fest-art-fw-a { top: 5px; left: 8px; animation: ac-fest-burst 6s 1s infinite; }
.ac-fest-art-fw-b { top: 34px; left: 52px; width: 66px; height: 66px; opacity: 0.6; animation: ac-fest-burst 6s 3.5s infinite; }
.ac-fest-fish { position: absolute; filter: drop-shadow(0 4px 6px rgba(0, 0, 0, 0.35)); }
.ac-fest-fish-a { left: 31px; bottom: 28px; width: 92px; animation: ac-fest-swim 6s ease-in-out infinite; }
.ac-fest-fish-b { right: 35px; bottom: 47px; width: 56px; opacity: 0.7; animation: ac-fest-swim 7.5s ease-in-out -2.4s infinite reverse; }
.ac-fest-art-line { position: absolute; left: 18px; right: 18px; bottom: 15px; height: 1px; background: linear-gradient(90deg, transparent, rgba(244, 199, 119, 0.4), transparent); }
.ac-fest-art-label {
  position: absolute; left: 19px; bottom: 21px; color: rgba(225, 202, 162, 0.63); font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 6.5px; letter-spacing: 0.24em;
}
.ac-fest-nameblock { position: relative; z-index: 3; padding: 17px 20px 0 23px; }
.ac-fest-small { color: #d7a85e; font-family: 'Songti SC', serif; font-size: 8px; letter-spacing: 0.34em; }
.ac-fest-name {
  margin-top: 5px; color: #fdf2dc; font-family: 'Songti SC', 'Noto Serif SC', 'STSong', serif; font-size: 42px; font-weight: 700; line-height: 1.08; letter-spacing: 0.08em;
  text-shadow: 0 0 22px rgba(255, 198, 111, 0.16), 0 2px 0 rgba(0, 0, 0, 0.35); overflow-wrap: anywhere;
}
.ac-fest-title {
  display: inline-block; margin-top: 8px; border-bottom: 1px solid rgba(221, 166, 89, 0.55); padding-bottom: 2px; color: #ecc789; font-size: 11.5px; letter-spacing: 0.18em;
  overflow-wrap: anywhere;
}
.ac-fest-meta {
  display: grid; grid-template-columns: 1fr 1fr; gap: 12px; margin: 17px 20px 0 23px; border-top: 1px solid rgba(228, 180, 104, 0.2);
  border-bottom: 1px solid rgba(228, 180, 104, 0.2); padding: 10px 0;
}
.ac-fest-meta div { min-width: 0; }
.ac-fest-meta span { display: block; color: #948a78; font-size: 6.5px; letter-spacing: 0.18em; }
.ac-fest-meta b { display: block; margin-top: 4px; color: #e4d9c5; font-size: 10.5px; font-weight: 500; letter-spacing: 0.04em; overflow-wrap: anywhere; }
.ac-fest-meta .ac-fest-accent { color: #f2b96c; }
.ac-fest-quote { position: relative; margin: 15px 20px 0 23px; padding-left: 14px; }
.ac-fest-quote-mark { position: absolute; left: 0; top: -9px; color: rgba(238, 189, 108, 0.6); font-family: Georgia, serif; font-size: 34px; line-height: 1; }
.ac-fest-quote p {
  color: #f2e6d1; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 13.5px; line-height: 1.8; letter-spacing: 0.04em; overflow-wrap: anywhere;
}
.ac-fest-foot {
  position: relative; z-index: 3; display: flex; align-items: center; justify-content: space-between; gap: 10px; margin-top: 17px; border-top: 1px solid rgba(228, 180, 104, 0.24);
  padding: 12px 20px 14px 23px; background: rgba(255, 211, 134, 0.035);
}
.ac-fest-social { display: flex; align-items: center; gap: 5px; min-width: 0; color: #e8c385; font-family: Consolas, 'Courier New', monospace; font-size: 10px; }
.ac-fest-social svg { flex: none; }
.ac-fest-social span { overflow-wrap: anywhere; }
.ac-fest-id { flex: none; color: #9a8c71; font-family: Consolas, monospace; font-size: 6px; letter-spacing: 0.08em; }
.ac-fest-caption {
  margin-top: 22px; text-align: center; color: rgba(229, 195, 139, 0.5); font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 9px; letter-spacing: 0.34em;
  animation: ac-fest-fade 1.6s 0.7s both;
}
@keyframes ac-fest-rise {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-fest-burst {
  0%, 6% { opacity: 0; transform: scale(0.25); }
  12% { opacity: 0.95; }
  46% { opacity: 0.55; transform: scale(1); }
  62%, 100% { opacity: 0; transform: scale(1.16); }
}
@keyframes ac-fest-sway {
  0%, 100% { transform: rotate(-2.2deg); }
  50% { transform: rotate(2.4deg); }
}
@keyframes ac-fest-swim {
  0%, 100% { transform: translate(0, 0) rotate(-2deg); }
  35% { transform: translate(9px, -5px) rotate(3deg); }
  65% { transform: translate(-4px, 4px) rotate(-1deg); }
}
@keyframes ac-fest-twinkle {
  0%, 100% { opacity: 0.38; }
  50% { opacity: 0.68; }
}
@keyframes ac-fest-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-fest-root *, .ac-fest-root *::before, .ac-fest-root *::after { animation: none !important; transition: none !important; }
}
</style>

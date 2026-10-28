<template>
  <div class="ac-forest-root min-h-screen overflow-hidden px-4 py-9 text-[#2f4229]">
    <div class="ac-forest-glow"></div>
    <div class="ac-forest-canopy"></div>

    <svg class="ac-forest-backleaf" viewBox="0 0 240 320" aria-hidden="true">
      <path d="M120 8 C212 58 232 172 168 248 C118 308 44 286 24 208 C4 130 42 52 120 8 Z" fill="none" stroke="#6f8860" stroke-width="1.2" />
      <path d="M121 22 C116 108 102 200 58 290" fill="none" stroke="#6f8860" stroke-width="1" />
      <path d="M112 92 C80 80 56 80 34 92 M116 126 C78 116 54 118 28 138 M112 164 C80 158 58 164 36 182 M102 202 C76 200 58 208 42 226 M88 240 C68 242 54 250 44 264" fill="none" stroke="#6f8860" stroke-width="0.9" />
      <path d="M124 78 C150 62 176 58 198 62 M126 116 C158 98 184 94 208 100 M122 154 C150 140 176 138 198 146 M112 192 C140 180 164 182 184 192 M94 234 C118 226 140 230 158 240" fill="none" stroke="#6f8860" stroke-width="0.9" />
    </svg>

    <span v-for="n in 9" :key="n" class="ac-forest-fly" :class="'ac-forest-fly-' + n"></span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-forest-page">
        <header class="ac-forest-head">
          <span class="ac-forest-mark">
            <Leaf :size="15" aria-hidden="true" />
          </span>
          <div class="ac-forest-heading">
            <p class="ac-forest-kicker">SPIRIT FIELD GUIDE</p>
            <p class="ac-forest-sub">森之精灵图鉴</p>
          </div>
          <span class="ac-forest-no">NO. {{ nfcId || 'PREVIEW' }}</span>
        </header>

        <div class="ac-forest-emblem">
          <span class="ac-forest-emblem-ring"></span>
          <span class="ac-forest-emblem-halo"></span>
          <svg class="ac-forest-emblem-leaf" viewBox="0 0 180 180" aria-hidden="true">
            <defs>
              <linearGradient id="forestLeaf" x1="0" y1="0" x2="1" y2="1">
                <stop offset="0" stop-color="#a9bb8e" stop-opacity="0.5" />
                <stop offset="1" stop-color="#5e7a4d" stop-opacity="0.16" />
              </linearGradient>
            </defs>
            <path d="M90 24 C134 40 150 86 132 124 C116 158 74 164 50 134 C28 106 40 60 90 24 Z" fill="url(#forestLeaf)" stroke="#5e7a4d" stroke-width="1.2" />
            <path d="M90 32 C88 82 82 122 58 146" fill="none" stroke="#5e7a4d" stroke-width="1" opacity="0.9" />
            <path d="M86 62 C66 54 50 54 36 62 M88 88 C60 78 44 80 30 94 M84 114 C64 108 50 112 38 124 M74 138 C60 136 50 140 42 150" fill="none" stroke="#5e7a4d" stroke-width="0.9" opacity="0.8" />
            <path d="M92 56 C112 44 132 44 146 50 M90 82 C116 68 136 68 150 78 M86 108 C108 96 128 98 140 108 M78 132 C96 124 112 128 122 138" fill="none" stroke="#5e7a4d" stroke-width="0.9" opacity="0.8" />
            <circle cx="142" cy="54" r="1.8" fill="#c8a85e" />
            <circle cx="66" cy="66" r="1.2" fill="#c8a85e" opacity="0.7" />
            <circle cx="126" cy="120" r="1.4" fill="#c8a85e" opacity="0.8" />
          </svg>
          <span class="ac-forest-initial">{{ initial }}</span>
          <span class="ac-forest-emblem-caption">PRESSED LEAF · 压叶标本</span>
        </div>

        <div class="ac-forest-titleblock">
          <p class="ac-forest-eyebrow">
            <Sparkles :size="8" aria-hidden="true" />
            晨雾中的第 {{ nfcId || '一' }} 位访客
          </p>
          <h1 class="ac-forest-name">{{ p.name || '无名林精' }}</h1>
          <p class="ac-forest-title">{{ p.title || '尚未被命名的存在' }}</p>
        </div>

        <div class="ac-forest-meta">
          <div>
            <span>出没作品</span>
            <b>{{ p.series || '某片无名森林' }}</b>
          </div>
          <div>
            <span>灵属性</span>
            <b class="ac-forest-element">{{ p.element || '苔 · 微光' }}</b>
          </div>
        </div>

        <blockquote class="ac-forest-quote">
          <span class="ac-forest-quote-mark" aria-hidden="true">“</span>
          <p>{{ p.quote || '别急着赶路，坐下来听一会儿树叶的声音吧。' }}</p>
          <span class="ac-forest-quote-end" aria-hidden="true">”</span>
        </blockquote>

        <div class="ac-forest-note">
          <span class="ac-forest-note-icon">
            <Sprout :size="13" aria-hidden="true" />
          </span>
          <div>
            <p class="ac-forest-note-label">观察记录 / FIELD NOTE</p>
            <p class="ac-forest-note-text">只在雨后初晴的林间出现，脚步轻得像一片落叶。</p>
          </div>
        </div>

        <footer class="ac-forest-foot">
          <div class="ac-forest-social">
            <AtSign :size="11" aria-hidden="true" />
            <div>
              <span>观察者 / SNS</span>
              <b>{{ p.social || '未留下踪迹' }}</b>
            </div>
          </div>
          <span class="ac-forest-serial">FIELD NO. {{ nfcId || 'PREVIEW' }}</span>
        </footer>
      </div>

      <p class="ac-forest-caption">MOSS · MIST · MORNING LIGHT</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { AtSign, Leaf, Sparkles, Sprout } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const initial = computed(() => (p.value.name || '森').toString().trim().slice(0, 1) || '森')
</script>

<style scoped>
.ac-forest-root {
  position: relative;
  background: radial-gradient(circle at 14% 8%, rgba(154, 180, 130, 0.3), transparent 30%), radial-gradient(circle at 88% 78%, rgba(115, 148, 102, 0.2), transparent 38%), radial-gradient(circle at 50% 100%, rgba(224, 215, 171, 0.4), transparent 42%), linear-gradient(165deg, #f6f5e8 0%, #edf2dd 50%, #e5ecd3 100%);
  font-family: 'Songti SC', 'Noto Serif SC', 'STSong', serif;
}
.ac-forest-glow {
  position: absolute; top: -120px; left: 50%; width: 520px; height: 460px; transform: translateX(-50%); border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 252, 222, 0.75), transparent 66%); filter: blur(12px);
}
.ac-forest-canopy {
  position: absolute; inset: 0;
  background-image: radial-gradient(rgba(85, 112, 70, 0.09) 0.7px, transparent 0.8px), radial-gradient(rgba(255, 255, 255, 0.8) 0.6px, transparent 0.7px);
  background-position: 0 0, 9px 13px; background-size: 27px 31px, 23px 29px; opacity: 0.55;
}
.ac-forest-backleaf {
  position: absolute; right: -78px; bottom: -10px; width: 260px; height: 346px; opacity: 0.25; transform: rotate(12deg); animation: ac-forest-sway 9s ease-in-out infinite;
  transform-origin: bottom center;
}
.ac-forest-fly {
  position: absolute; width: 4px; height: 4px; border-radius: 50%; background: #f6e3a5; box-shadow: 0 0 6px 2px rgba(241, 216, 139, 0.7), 0 0 16px 5px rgba(241, 216, 139, 0.22);
  opacity: 0; animation: ac-forest-fly 7s ease-in-out infinite;
}
.ac-forest-fly-1 { left: 7%; top: 18%; animation-delay: 0s; }
.ac-forest-fly-2 { left: 88%; top: 24%; animation-delay: 1.1s; width: 3px; height: 3px; }
.ac-forest-fly-3 { left: 13%; top: 52%; animation-delay: 2.4s; }
.ac-forest-fly-4 { left: 92%; top: 62%; animation-delay: 0.7s; width: 3px; height: 3px; }
.ac-forest-fly-5 { left: 20%; top: 82%; animation-delay: 3.5s; }
.ac-forest-fly-6 { left: 78%; top: 88%; animation-delay: 1.8s; width: 5px; height: 5px; }
.ac-forest-fly-7 { left: 4%; top: 34%; animation-delay: 4.4s; width: 3px; height: 3px; }
.ac-forest-fly-8 { left: 82%; top: 8%; animation-delay: 2.9s; }
.ac-forest-fly-9 { left: 46%; top: 94%; animation-delay: 5.2s; width: 3px; height: 3px; }
.ac-forest-page {
  position: relative; padding: 19px 20px 16px; border: 1px solid rgba(83, 108, 69, 0.26); border-radius: 3px 12px 4px 10px;
  background: radial-gradient(circle at 12% 6%, rgba(166, 186, 143, 0.16), transparent 26%), radial-gradient(circle at 92% 96%, rgba(178, 190, 145, 0.18), transparent 30%), repeating-linear-gradient(93deg, rgba(86, 103, 66, 0.018) 0 1px, transparent 1px 7px), #fbf9ed;
  box-shadow: 0 30px 56px -26px rgba(55, 75, 45, 0.5), 0 1px 0 rgba(255, 255, 255, 0.9) inset, inset 0 0 58px rgba(164, 179, 133, 0.12);
  animation: ac-forest-in 0.95s cubic-bezier(0.2, 0.9, 0.25, 1) both; transition: transform 0.55s ease, box-shadow 0.55s ease;
}
.ac-forest-page:hover {
  transform: translateY(-3px) rotate(0.2deg);
  box-shadow: 0 38px 66px -26px rgba(55, 75, 45, 0.55), 0 1px 0 rgba(255, 255, 255, 0.9) inset, inset 0 0 58px rgba(164, 179, 133, 0.12);
}
.ac-forest-page::before { content: ''; position: absolute; inset: 8px; border: 1px solid rgba(93, 117, 76, 0.16); border-radius: 2px 9px 3px 8px; pointer-events: none; }
.ac-forest-page::after {
  content: ''; position: absolute; top: 0; right: 28px; width: 1px; height: 100%;
  background: linear-gradient(180deg, transparent, rgba(152, 164, 119, 0.14) 16%, rgba(152, 164, 119, 0.14) 84%, transparent); pointer-events: none;
}
.ac-forest-head { position: relative; z-index: 2; display: flex; align-items: center; gap: 9px; border-bottom: 1px solid rgba(87, 111, 70, 0.24); padding-bottom: 12px; }
.ac-forest-mark {
  display: grid; flex: none; place-items: center; width: 32px; height: 32px; border: 1px solid rgba(74, 100, 59, 0.55); border-radius: 50%;
  background: radial-gradient(circle, rgba(201, 215, 178, 0.45), transparent 72%); color: #4e6d40;
}
.ac-forest-heading { flex: 1; min-width: 0; }
.ac-forest-kicker { color: #7c906e; font-family: Consolas, 'Courier New', monospace; font-size: 6px; letter-spacing: 0.24em; }
.ac-forest-sub { margin-top: 3px; color: #405a36; font-size: 11px; font-weight: 600; letter-spacing: 0.2em; }
.ac-forest-no {
  flex: none; border: 1px solid rgba(89, 113, 72, 0.35); border-radius: 2px; padding: 3px 5px 2px; color: #718465; font-family: Consolas, monospace; font-size: 5.5px;
  letter-spacing: 0.08em; white-space: nowrap;
}
.ac-forest-emblem { position: relative; width: 158px; height: 158px; margin: 20px auto 6px; display: grid; place-items: center; }
.ac-forest-emblem-ring { position: absolute; inset: 0; border: 1px dashed rgba(95, 124, 75, 0.48); border-radius: 50%; animation: ac-forest-spin 41s linear infinite; }
.ac-forest-emblem-ring::before { content: ''; position: absolute; inset: 8px; border: 1px solid rgba(120, 145, 97, 0.2); border-radius: 50%; }
.ac-forest-emblem-halo {
  position: absolute; inset: 16px; border-radius: 50%; background: radial-gradient(circle, rgba(179, 199, 147, 0.48), rgba(222, 229, 194, 0.15) 58%, transparent 72%);
  animation: ac-forest-halo 4.5s ease-in-out infinite;
}
.ac-forest-emblem-leaf { position: absolute; width: 118px; height: 118px; transform: rotate(-8deg); filter: drop-shadow(0 5px 7px rgba(71, 93, 54, 0.12)); }
.ac-forest-initial {
  position: relative; z-index: 2; color: #3e5a33; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 47px; font-weight: 700; line-height: 1;
  text-shadow: 0 1px 0 rgba(255, 253, 240, 0.9);
}
.ac-forest-emblem-caption {
  position: absolute; bottom: -7px; left: 50%; transform: translateX(-50%); background: #fbf9ed; padding: 0 7px; color: #839475; font-family: Consolas, monospace; font-size: 5px;
  letter-spacing: 0.18em; white-space: nowrap;
}
.ac-forest-titleblock { position: relative; z-index: 2; margin-top: 8px; text-align: center; }
.ac-forest-eyebrow { display: inline-flex; align-items: center; gap: 4px; color: #80936f; font-size: 7px; letter-spacing: 0.18em; }
.ac-forest-eyebrow svg { color: #b39750; }
.ac-forest-name {
  margin-top: 5px; color: #2f4428; font-family: 'Songti SC', 'Noto Serif SC', 'STSong', serif; font-size: 41px; font-weight: 700; line-height: 1.08; letter-spacing: 0.06em;
  overflow-wrap: anywhere;
}
.ac-forest-title {
  display: inline-flex; align-items: center; gap: 8px; margin-top: 6px; color: #697f5b; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 12px;
  letter-spacing: 0.12em; overflow-wrap: anywhere;
}
.ac-forest-title::before, .ac-forest-title::after { content: ''; width: 16px; height: 1px; background: rgba(106, 130, 88, 0.5); }
.ac-forest-meta {
  position: relative; z-index: 2; display: grid; grid-template-columns: 1fr 1fr; gap: 14px; margin-top: 18px; border-top: 1px solid rgba(96, 119, 78, 0.22);
  border-bottom: 1px solid rgba(96, 119, 78, 0.22); padding: 11px 0;
}
.ac-forest-meta div { min-width: 0; }
.ac-forest-meta span { display: block; color: #8b9b7d; font-size: 6.5px; letter-spacing: 0.18em; }
.ac-forest-meta b {
  display: block; margin-top: 5px; border-bottom: 1px dotted rgba(94, 117, 76, 0.34); padding-bottom: 4px; color: #3b5231; font-size: 10.5px; font-weight: 600;
  letter-spacing: 0.04em; overflow-wrap: anywhere;
}
.ac-forest-meta .ac-forest-element { color: #6e8349; }
.ac-forest-quote {
  position: relative; margin-top: 16px; border-left: 2px solid rgba(114, 139, 87, 0.58);
  background: linear-gradient(90deg, rgba(182, 199, 153, 0.18), transparent 72%), rgba(245, 244, 226, 0.48); padding: 12px 17px 12px 20px;
}
.ac-forest-quote-mark { position: absolute; left: 4px; top: -6px; color: rgba(108, 133, 81, 0.55); font-family: Georgia, serif; font-size: 30px; line-height: 1; }
.ac-forest-quote p { color: #405335; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 15.5px; line-height: 1.8; letter-spacing: 0.03em; overflow-wrap: anywhere; }
.ac-forest-quote-end { position: absolute; right: 5px; bottom: -12px; color: rgba(108, 133, 81, 0.32); font-family: Georgia, serif; font-size: 26px; line-height: 1; }
.ac-forest-note { display: flex; gap: 9px; margin-top: 17px; border: 1px dashed rgba(102, 126, 83, 0.32); border-radius: 3px; padding: 9px 10px; background: rgba(255, 255, 252, 0.32); }
.ac-forest-note-icon { display: grid; flex: none; place-items: center; width: 24px; height: 24px; border-radius: 50%; background: rgba(142, 166, 112, 0.18); color: #5f7b4f; }
.ac-forest-note-label { color: #87987a; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.14em; }
.ac-forest-note-text { margin-top: 4px; color: #5e7050; font-size: 9px; line-height: 1.65; letter-spacing: 0.03em; }
.ac-forest-foot { display: flex; align-items: flex-end; justify-content: space-between; gap: 10px; margin-top: 17px; border-top: 1px solid rgba(96, 119, 78, 0.24); padding-top: 11px; }
.ac-forest-social { display: flex; align-items: center; gap: 6px; min-width: 0; color: #647a55; }
.ac-forest-social svg { flex: none; color: #889d72; }
.ac-forest-social span { display: block; color: #92a184; font-size: 5.5px; letter-spacing: 0.14em; }
.ac-forest-social b {
  display: block; margin-top: 2px; color: #4c653f; font-family: Consolas, 'Courier New', monospace; font-size: 9.5px; font-weight: 600; letter-spacing: 0.02em;
  overflow-wrap: anywhere;
}
.ac-forest-serial { flex: none; color: #91a083; font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.08em; white-space: nowrap; }
.ac-forest-caption {
  margin-top: 23px; text-align: center; color: rgba(67, 88, 54, 0.45); font-family: Consolas, 'Courier New', monospace; font-size: 7.5px; letter-spacing: 0.36em;
  animation: ac-forest-fade 1.5s 0.7s both;
}
@keyframes ac-forest-in {
  from { opacity: 0; transform: translateY(24px) scale(0.985); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-forest-spin {
  to { transform: rotate(360deg); }
}
@keyframes ac-forest-halo {
  0%, 100% { opacity: 0.7; transform: scale(0.97); }
  50% { opacity: 1; transform: scale(1.04); }
}
@keyframes ac-forest-fly {
  0% { opacity: 0; transform: translate(0, 0); }
  16% { opacity: 0.9; }
  52% { opacity: 0.45; transform: translate(9px, -18px); }
  76% { opacity: 0.95; }
  100% { opacity: 0; transform: translate(-6px, -34px); }
}
@keyframes ac-forest-sway {
  0%, 100% { transform: rotate(12deg); }
  50% { transform: rotate(9deg); }
}
@keyframes ac-forest-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-forest-root *, .ac-forest-root *::before, .ac-forest-root *::after { animation: none !important; transition: none !important; }
}
</style>

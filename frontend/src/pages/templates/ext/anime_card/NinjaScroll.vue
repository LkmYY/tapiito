<template>
  <div class="ac-ninja-root relative min-h-screen overflow-hidden px-4 py-7 text-[#eee4cf]">
    <div class="ac-ninja-moon" aria-hidden="true"></div>
    <div class="ac-ninja-mist" aria-hidden="true"></div>
    <div class="ac-ninja-grain" aria-hidden="true"></div>

    <svg class="ac-ninja-shuriken ac-ninja-shuriken-a" viewBox="0 0 100 100" aria-hidden="true">
      <path d="M50 2 61 37 98 50 61 63 50 98 39 63 2 50 39 37Z" fill="none" stroke="currentColor" stroke-width="1.1" />
      <circle cx="50" cy="50" r="10" fill="none" stroke="currentColor" stroke-width="1" />
      <circle cx="50" cy="50" r="2" fill="currentColor" />
    </svg>
    <svg class="ac-ninja-shuriken ac-ninja-shuriken-b" viewBox="0 0 100 100" aria-hidden="true">
      <path d="M50 2 61 37 98 50 61 63 50 98 39 63 2 50 39 37Z" fill="none" stroke="currentColor" stroke-width="1.1" />
      <circle cx="50" cy="50" r="10" fill="none" stroke="currentColor" stroke-width="1" />
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-ninja-top">
        <span class="ac-ninja-classified"><i></i> 機密文書 / EYES ONLY</span>
        <span class="ac-ninja-no">巻之 {{ nfcId || 'PREVIEW' }}</span>
      </header>

      <main class="ac-ninja-scroll">
        <div class="ac-ninja-roller" aria-hidden="true"><span></span></div>

        <article class="ac-ninja-paper">
          <div class="ac-ninja-fibers" aria-hidden="true"></div>
          <div class="ac-ninja-watermark" aria-hidden="true">忍</div>
          <div class="ac-ninja-verse" aria-hidden="true">風林火山</div>

          <div class="ac-ninja-head">
            <span>忍者名鑑</span>
            <span>影組 · 極秘</span>
          </div>

          <div class="ac-ninja-hero">
            <div class="ac-ninja-hero-copy">
              <p class="ac-ninja-role">「{{ p.title }}」</p>
              <h1>{{ p.name || '未命名角色' }}</h1>
              <p class="ac-ninja-series">出典 · {{ p.series || '未知作品' }}</p>
            </div>
            <span class="ac-ninja-seal" aria-hidden="true">秘</span>
          </div>

          <div class="ac-ninja-brush" aria-hidden="true"><i></i></div>

          <blockquote class="ac-ninja-quote">
            <span>口 伝</span>
            <p>「{{ p.quote || '……风过无痕，影落无声，任务完成。' }}」</p>
          </blockquote>

          <div class="ac-ninja-meta">
            <div><span>遁術 / STYLE</span><b>{{ p.element }}</b></div>
            <div><span>連絡 / CONTACT</span><b class="ac-ninja-social">@{{ p.social }}</b></div>
          </div>

          <footer class="ac-ninja-foot">
            <div class="ac-ninja-hanko">忍<small>影組</small></div>
            <div class="ac-ninja-foot-copy">
              <p>風過無痕 · 影落無聲</p>
              <small>SCROLL NO. {{ nfcId || 'PREVIEW' }} · 焼却厳禁</small>
            </div>
          </footer>
        </article>

        <div class="ac-ninja-roller ac-ninja-roller-bottom" aria-hidden="true"><span></span></div>
      </main>

      <p class="ac-ninja-caption"><span>✦</span> 風遁 · 巻物厳封 <span>✦</span></p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-ninja-root { background: radial-gradient(circle at 78% 16%, rgba(127, 29, 29, 0.2), transparent 25%), radial-gradient(ellipse at 20% 100%, rgba(69, 61, 46, 0.2), transparent 44%), linear-gradient(180deg, #070706 0%, #0c0b09 54%, #050504 100%); font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; }
.ac-ninja-moon { position: absolute; top: 68px; right: -28px; width: 162px; height: 162px; border-radius: 50%; background: radial-gradient(circle at 36% 32%, rgba(255, 187, 171, 0.18), transparent 18%), linear-gradient(145deg, #7f1d1d, #3f0c0c 65%, #180404); box-shadow: 0 0 70px rgba(153, 27, 27, 0.23), inset -16px -10px 30px rgba(0, 0, 0, 0.55); opacity: 0.6; }
.ac-ninja-mist { position: absolute; left: -30%; right: -30%; bottom: -40px; height: 340px; background: radial-gradient(ellipse at 30% 60%, rgba(187, 173, 148, 0.1), transparent 42%), radial-gradient(ellipse at 72% 58%, rgba(187, 173, 148, 0.08), transparent 46%), radial-gradient(ellipse at 50% 88%, rgba(122, 116, 104, 0.08), transparent 50%); filter: blur(18px); animation: ac-ninja-drift 14s ease-in-out infinite alternate; pointer-events: none; }
.ac-ninja-grain { position: absolute; inset: 0; background-image: radial-gradient(rgba(242, 233, 213, 0.18) 0.5px, transparent 0.7px); background-size: 7px 7px; opacity: 0.12; pointer-events: none; }
.ac-ninja-shuriken { position: absolute; color: rgba(221, 207, 181, 0.2); pointer-events: none; }
.ac-ninja-shuriken-a { top: 18%; left: -18px; width: 88px; height: 88px; animation: ac-ninja-spin 46s linear infinite; }
.ac-ninja-shuriken-b { right: -17px; bottom: 18%; width: 56px; height: 56px; opacity: 0.65; animation: ac-ninja-spin 34s linear infinite reverse; }
.ac-ninja-top { display: flex; align-items: center; justify-content: space-between; gap: 10px; font-family: Consolas, 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.17em; color: rgba(222, 207, 177, 0.48); animation: ac-ninja-fade 0.9s both; }
.ac-ninja-classified { display: inline-flex; align-items: center; gap: 6px; white-space: nowrap; }
.ac-ninja-classified i { width: 4px; height: 4px; border-radius: 50%; background: #b91c1c; box-shadow: 0 0 8px rgba(185, 28, 28, 0.9); animation: ac-ninja-pulse 2.4s ease-in-out infinite; }
.ac-ninja-no { color: rgba(222, 207, 177, 0.32); white-space: nowrap; }
.ac-ninja-scroll { position: relative; margin-top: 22px; filter: drop-shadow(0 26px 45px rgba(0, 0, 0, 0.65)); animation: ac-ninja-unroll 1s cubic-bezier(0.16, 1, 0.3, 1) both; }
.ac-ninja-roller { position: relative; z-index: 4; height: 15px; border-radius: 4px; background: linear-gradient(180deg, #6b5438 0%, #2e2114 38%, #160f08 72%, #4b3925 100%); box-shadow: 0 4px 10px rgba(0, 0, 0, 0.6); }
.ac-ninja-roller::before, .ac-ninja-roller::after { content: ''; position: absolute; top: 50%; width: 21px; height: 21px; border-radius: 50%; background: radial-gradient(circle at 35% 32%, #b28d57, #5d4529 52%, #2b1e10 100%); box-shadow: 0 2px 5px rgba(0, 0, 0, 0.55); transform: translateY(-50%); }
.ac-ninja-roller::before { left: -7px; }
.ac-ninja-roller::after { right: -7px; }
.ac-ninja-roller span { position: absolute; top: 1px; left: 12%; right: 12%; height: 1px; background: linear-gradient(90deg, transparent, rgba(232, 214, 179, 0.22), transparent); }
.ac-ninja-roller-bottom { margin-top: -1px; }
.ac-ninja-paper { position: relative; overflow: hidden; border: 1px solid rgba(185, 157, 111, 0.32); background: linear-gradient(90deg, rgba(230, 212, 176, 0.035), transparent 18%, transparent 82%, rgba(230, 212, 176, 0.03)), linear-gradient(155deg, #191813 0%, #0e0e0b 52%, #16140f 100%); padding: 23px 19px 18px; box-shadow: inset 0 0 42px rgba(0, 0, 0, 0.55); }
.ac-ninja-fibers { position: absolute; inset: 0; z-index: 0; background-image: repeating-linear-gradient(2deg, transparent 0 7px, rgba(238, 226, 203, 0.021) 7px 8px), repeating-linear-gradient(88deg, transparent 0 19px, rgba(238, 226, 203, 0.017) 19px 20px); mix-blend-mode: screen; pointer-events: none; }
.ac-ninja-watermark { position: absolute; z-index: 0; top: 42px; right: 3px; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 168px; line-height: 1; color: transparent; -webkit-text-stroke: 1px rgba(185, 157, 111, 0.11); transform: rotate(4deg); pointer-events: none; user-select: none; }
.ac-ninja-verse { position: absolute; z-index: 0; top: 132px; right: 8px; writing-mode: vertical-rl; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 26px; letter-spacing: 0.2em; color: rgba(232, 219, 191, 0.045); pointer-events: none; }
.ac-ninja-head { position: relative; z-index: 2; display: flex; align-items: center; justify-content: space-between; gap: 10px; border-bottom: 1px solid rgba(195, 174, 137, 0.2); padding-bottom: 8px; font-size: 7px; letter-spacing: 0.17em; color: rgba(227, 213, 184, 0.42); }
.ac-ninja-head span:first-child { color: #c5a87a; }
.ac-ninja-hero { position: relative; z-index: 2; display: flex; align-items: flex-start; justify-content: space-between; gap: 12px; padding-top: 20px; }
.ac-ninja-hero-copy { min-width: 0; flex: 1; }
.ac-ninja-role { font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 10px; letter-spacing: 0.15em; color: #b89461; overflow-wrap: anywhere; }
.ac-ninja-hero h1 { margin-top: 8px; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: clamp(34px, 9.6vw, 44px); font-weight: 700; line-height: 1.1; letter-spacing: 0.08em; color: #f0e5cf; text-shadow: 0 2px 18px rgba(0, 0, 0, 0.7); overflow-wrap: anywhere; }
.ac-ninja-series { margin-top: 8px; font-size: 8.5px; letter-spacing: 0.08em; color: rgba(221, 206, 177, 0.5); overflow-wrap: anywhere; }
.ac-ninja-seal { flex: none; display: flex; align-items: center; justify-content: center; width: 42px; height: 42px; border: 2px solid rgba(161, 38, 28, 0.82); box-shadow: inset 0 0 0 1px rgba(161, 38, 28, 0.35); font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 25px; line-height: 1; color: #b73a2a; transform: rotate(6deg); opacity: 0.9; }
.ac-ninja-brush { display: flex; align-items: center; margin: 16px 0 0; }
.ac-ninja-brush::before, .ac-ninja-brush::after { content: ''; flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(192, 160, 110, 0.35)); }
.ac-ninja-brush::after { background: linear-gradient(90deg, rgba(192, 160, 110, 0.35), transparent); }
.ac-ninja-brush i { width: 5px; height: 5px; margin: 0 7px; background: #8f2520; transform: rotate(45deg); }
.ac-ninja-quote { position: relative; z-index: 2; margin-top: 15px; border-left: 2px solid rgba(161, 38, 28, 0.72); background: linear-gradient(90deg, rgba(161, 38, 28, 0.075), transparent 80%); padding: 10px 11px; }
.ac-ninja-quote > span { display: block; margin-bottom: 5px; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 8px; letter-spacing: 0.34em; color: #b73a2a; }
.ac-ninja-quote p { font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 14px; line-height: 1.9; color: rgba(241, 230, 208, 0.87); overflow-wrap: anywhere; }
.ac-ninja-meta { position: relative; z-index: 2; display: grid; gap: 0; margin-top: 15px; border-top: 1px solid rgba(195, 174, 137, 0.15); }
.ac-ninja-meta > div { display: flex; align-items: baseline; gap: 10px; border-bottom: 1px dotted rgba(195, 174, 137, 0.14); padding: 7px 1px; }
.ac-ninja-meta > div:last-child { border-bottom: 0; }
.ac-ninja-meta span { flex: none; width: 94px; font-family: Consolas, monospace; font-size: 6px; letter-spacing: 0.1em; color: rgba(197, 168, 122, 0.5); }
.ac-ninja-meta b { min-width: 0; flex: 1; text-align: right; font-size: 10.5px; font-weight: 500; line-height: 1.5; color: rgba(238, 228, 207, 0.8); overflow-wrap: anywhere; }
.ac-ninja-meta b.ac-ninja-social { color: #c88e79; }
.ac-ninja-foot { position: relative; z-index: 2; display: flex; align-items: center; gap: 13px; margin-top: 17px; border-top: 1px solid rgba(195, 174, 137, 0.16); padding-top: 13px; }
.ac-ninja-hanko { flex: none; display: flex; flex-direction: column; align-items: center; justify-content: center; width: 47px; height: 47px; border: 2px solid rgba(169, 42, 30, 0.82); border-radius: 3px; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 25px; line-height: 0.95; color: #b23b2b; transform: rotate(-5deg); opacity: 0.9; }
.ac-ninja-hanko small { margin-top: 2px; font-family: sans-serif; font-size: 5.5px; letter-spacing: 0.1em; }
.ac-ninja-foot-copy { min-width: 0; flex: 1; text-align: right; }
.ac-ninja-foot-copy p { font-family: 'Kaiti SC', serif; font-size: 10px; letter-spacing: 0.12em; color: rgba(226, 213, 187, 0.55); }
.ac-ninja-foot-copy small { display: block; margin-top: 4px; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.05em; color: rgba(226, 213, 187, 0.28); overflow-wrap: anywhere; }
.ac-ninja-caption { margin-top: 20px; text-align: center; font-size: 7.5px; letter-spacing: 0.24em; color: rgba(222, 207, 177, 0.3); animation: ac-ninja-fade 1.3s 0.55s both; }
.ac-ninja-caption span { color: rgba(183, 58, 42, 0.55); margin: 0 6px; }
@keyframes ac-ninja-unroll {
  from { opacity: 0; transform: scaleY(0.86) translateY(-16px); filter: brightness(1.6); }
  to { opacity: 1; transform: none; filter: none; }
}
@keyframes ac-ninja-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes ac-ninja-spin {
  to { transform: rotate(360deg); }
}
@keyframes ac-ninja-drift {
  from { transform: translateX(-18px) translateY(4px); opacity: 0.7; }
  to { transform: translateX(18px) translateY(-8px); opacity: 1; }
}
@keyframes ac-ninja-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-ninja-root *, .ac-ninja-root *::before, .ac-ninja-root *::after { animation: none !important; transition: none !important; }
}
</style>

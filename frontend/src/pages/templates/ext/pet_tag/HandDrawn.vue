<template>
  <div class="pt-draw-root min-h-screen px-5 py-9 text-amber-950">
    <div class="pt-draw-content mx-auto w-full max-w-md">
      <p class="pt-draw-top">— 我的宠物小档案 —</p>

      <article class="pt-draw-paper">
        <span class="pt-draw-tape" aria-hidden="true"></span>

        <div class="pt-draw-doodle pt-draw-doodle-star" aria-hidden="true">✧</div>
        <div class="pt-draw-doodle pt-draw-doodle-heart" aria-hidden="true">♡</div>

        <header class="pt-draw-head">
          <div class="pt-draw-title">
            <span class="pt-draw-hi">你好呀，我是</span>
            <h1>{{ payload.petName || '未命名' }}</h1>
            <p>{{ payload.breed }} · {{ payload.age || '—' }}</p>
          </div>
          <div class="pt-draw-pet" aria-hidden="true">
            <span class="pt-draw-ear pt-draw-ear-l"></span>
            <span class="pt-draw-ear pt-draw-ear-r"></span>
            <span class="pt-draw-head-shape"></span>
            <span class="pt-draw-eye pt-draw-eye-l"></span>
            <span class="pt-draw-eye pt-draw-eye-r"></span>
            <span class="pt-draw-muzzle"></span>
            <span class="pt-draw-nose"></span>
            <span class="pt-draw-tongue"></span>
          </div>
        </header>

        <svg class="pt-draw-divider" viewBox="0 0 340 12" preserveAspectRatio="none" aria-hidden="true">
          <path d="M2 7 Q 28 1, 56 6 T 112 6 T 170 5 T 228 7 T 286 5 T 338 6" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" />
        </svg>

        <div class="pt-draw-owner">
          <span class="pt-draw-key">铲屎官</span>
          <span class="pt-draw-value">{{ payload.ownerName }}</span>
          <span class="pt-draw-key">家</span>
          <span class="pt-draw-value">{{ payload.address }}</span>
        </div>

        <div class="pt-draw-sticky">
          <span class="pt-draw-pin" aria-hidden="true"></span>
          <span class="pt-draw-sticky-title">要记住的身体小秘密</span>
          <p>{{ payload.medical || '暂无特殊记录' }}</p>
        </div>

        <div class="pt-draw-say">
          <span class="pt-draw-say-label">它想说：</span>
          <p>{{ payload.note || '我有点怕生，请先让我闻闻你的手。' }}</p>
        </div>

        <a class="pt-draw-phone" :href="'tel:' + payload.phone">
          <span class="pt-draw-phone-label">走丢时，请打这个电话</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-draw-phone-icon" aria-hidden="true">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
              <path d="M5 4h4l2 5-2.5 1.5a11 11 0 0 0 5 5L15 13l5 2v4a2 2 0 0 1-2 2A16 16 0 0 1 3 6a2 2 0 0 1 2-2z" />
            </svg>
          </span>
        </a>

        <div class="pt-draw-bottom">
          <span class="pt-draw-pawprint" aria-hidden="true">
            <svg viewBox="0 0 48 48" fill="currentColor">
              <ellipse cx="24" cy="31" rx="10" ry="8" />
              <ellipse cx="11" cy="21" rx="4" ry="5" />
              <ellipse cx="21" cy="15" rx="4" ry="5.5" />
              <ellipse cx="32" cy="16" rx="4" ry="5" />
              <ellipse cx="40" cy="24" rx="3.4" ry="4.4" />
            </svg>
          </span>
          <span class="pt-draw-sign">画于一个阳光很好的下午</span>
          <span class="pt-draw-pawprint" aria-hidden="true">
            <svg viewBox="0 0 48 48" fill="currentColor">
              <ellipse cx="24" cy="31" rx="10" ry="8" />
              <ellipse cx="11" cy="21" rx="4" ry="5" />
              <ellipse cx="21" cy="15" rx="4" ry="5.5" />
              <ellipse cx="32" cy="16" rx="4" ry="5" />
              <ellipse cx="40" cy="24" rx="3.4" ry="4.4" />
            </svg>
          </span>
        </div>
      </article>

      <p class="pt-draw-bottomtext">如果捡到我，请帮我拨上面的号码 ♡</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-draw-root {
  background:
    radial-gradient(circle at 18% 14%, rgba(196, 150, 72, 0.1), transparent 30%),
    repeating-linear-gradient(0deg, transparent 0 27px, rgba(151, 110, 45, 0.055) 27px 28px),
    #f8f1df;
  font-family: 'Kaiti SC', 'STKaiti', 'KaiTi', 'Songti SC', serif;
}
.pt-draw-content { animation: pt-draw-wobble 0.8s cubic-bezier(0.2, 0.9, 0.3, 1) both; }
.pt-draw-top { margin: 0 0 15px; text-align: center; font-size: 12px; letter-spacing: 0.28em; color: #9a7644; }
.pt-draw-paper {
  position: relative;
  padding: 26px 22px 20px;
  background: #fffcf1;
  border: 2px solid #4b3926;
  border-radius: 15px 6px 18px 8px / 8px 17px 6px 15px;
  box-shadow: 5px 7px 0 #eadcbf, 0 17px 30px rgba(92, 67, 32, 0.12);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}
.pt-draw-paper:hover { transform: rotate(-0.45deg) translateY(-3px); box-shadow: 7px 10px 0 #eadcbf, 0 22px 36px rgba(92, 67, 32, 0.16); }
.pt-draw-paper::before {
  content: '';
  position: absolute;
  inset: 6px;
  border: 1px dashed rgba(101, 76, 42, 0.25);
  border-radius: 11px 5px 14px 7px / 7px 13px 5px 12px;
  pointer-events: none;
}
.pt-draw-tape { position: absolute; top: -13px; left: 36px; width: 90px; height: 24px; background: rgba(229, 203, 154, 0.7); transform: rotate(-3deg); box-shadow: 0 1px 2px rgba(96, 67, 28, 0.13); }
.pt-draw-doodle { position: absolute; color: #c2985a; line-height: 1; }
.pt-draw-doodle-star { top: 71px; right: 20px; font-size: 23px; transform: rotate(13deg); }
.pt-draw-doodle-heart { top: 120px; right: 35px; font-size: 17px; transform: rotate(-14deg); }
.pt-draw-head { display: flex; gap: 10px; align-items: flex-start; }
.pt-draw-title { flex: 1; padding-top: 6px; }
.pt-draw-hi { display: block; font-size: 13px; letter-spacing: 0.1em; color: #947245; }
.pt-draw-title h1 { margin: 2px 0 0; font-size: clamp(44px, 13vw, 58px); line-height: 1.05; font-weight: 400; letter-spacing: 0.08em; color: #392a1c; }
.pt-draw-title p { margin: 5px 0 0; font-size: 13px; letter-spacing: 0.12em; color: #8f6d41; }
.pt-draw-pet { position: relative; width: 102px; height: 104px; flex: none; margin-top: 4px; transform: rotate(4deg); }
.pt-draw-head-shape { position: absolute; left: 14px; top: 17px; width: 74px; height: 73px; background: #e0bc86; border: 2px solid #4b3926; border-radius: 46% 52% 45% 48% / 50% 47% 53% 46%; }
.pt-draw-ear { position: absolute; top: 2px; width: 30px; height: 43px; background: #d3a76b; border: 2px solid #4b3926; border-radius: 60% 40% 46% 54%; }
.pt-draw-ear-l { left: 9px; transform: rotate(-23deg); }
.pt-draw-ear-r { right: 8px; transform: rotate(24deg); }
.pt-draw-eye { position: absolute; top: 46px; width: 7px; height: 9px; border-radius: 50%; background: #3b2b1c; }
.pt-draw-eye-l { left: 36px; }
.pt-draw-eye-r { right: 34px; }
.pt-draw-muzzle { position: absolute; left: 50%; top: 66px; width: 39px; height: 27px; margin-left: -19.5px; background: #f2ddba; border: 1.5px solid #5b452b; border-radius: 47% 53% 50% 50%; }
.pt-draw-nose { position: absolute; left: 50%; top: 72px; width: 11px; height: 8px; margin-left: -5.5px; background: #3b2b1c; border-radius: 50% 50% 45% 45%; }
.pt-draw-tongue { position: absolute; left: 50%; top: 89px; width: 13px; height: 12px; margin-left: -6.5px; background: #dd8f80; border: 1.5px solid #5b452b; border-radius: 0 0 8px 8px; }
.pt-draw-divider { display: block; width: 100%; height: 12px; margin: 12px 0 10px; color: rgba(90, 67, 39, 0.5); }
.pt-draw-owner { display: grid; grid-template-columns: auto 1fr auto 1fr; gap: 7px; align-items: baseline; padding: 2px 3px 10px; font-size: 12px; }
.pt-draw-key { color: #a18154; letter-spacing: 0.08em; }
.pt-draw-value { font-size: 13px; font-weight: 600; color: #4a3622; border-bottom: 1px solid rgba(127, 96, 55, 0.35); padding-bottom: 1px; }
.pt-draw-sticky {
  position: relative;
  margin: 9px 2px 12px;
  padding: 13px 14px 12px;
  background: #fff0a8;
  border: 1.5px solid #c9a54a;
  border-radius: 3px 6px 4px 5px;
  transform: rotate(-1.1deg);
  box-shadow: 2px 3px 0 rgba(160, 127, 50, 0.2);
}
.pt-draw-pin { position: absolute; top: -7px; left: 50%; width: 12px; height: 12px; margin-left: -6px; border-radius: 50%; background: #cf6d5d; box-shadow: inset -2px -2px 0 rgba(100, 40, 30, 0.25); }
.pt-draw-sticky-title { display: block; margin-bottom: 3px; font-size: 11px; font-weight: 700; letter-spacing: 0.08em; color: #8b5d23; }
.pt-draw-sticky p { margin: 0; font-size: 13px; line-height: 1.6; color: #5e4627; }
.pt-draw-say { padding: 2px 4px 10px; }
.pt-draw-say-label { display: inline-block; font-size: 11px; color: #9a7644; margin-bottom: 3px; }
.pt-draw-say p { margin: 0; font-size: 13px; line-height: 1.75; color: #4e3a25; }
.pt-draw-say p::before { content: '“'; color: #b08a53; }
.pt-draw-say p::after { content: '”'; color: #b08a53; }
.pt-draw-phone {
  position: relative;
  display: block;
  margin-top: 3px;
  padding: 12px 48px 13px 16px;
  border: 2px solid #4b3926;
  border-radius: 19px 8px 16px 9px / 9px 15px 8px 18px;
  background: #fdf4dc;
  color: #3b2b1c;
  text-decoration: none;
  box-shadow: 3px 3px 0 #4b3926;
  transition: transform 0.18s ease, box-shadow 0.18s ease, background 0.18s ease;
}
.pt-draw-phone:hover { background: #fff9e9; transform: translate(-1px, -1px); box-shadow: 5px 5px 0 #4b3926; }
.pt-draw-phone:active { transform: translate(2px, 2px); box-shadow: 1px 1px 0 #4b3926; }
.pt-draw-phone-label { display: block; font-size: 10px; letter-spacing: 0.08em; color: #98764a; }
.pt-draw-phone strong { display: block; margin-top: 1px; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(26px, 8vw, 33px); font-weight: 800; letter-spacing: 0.02em; }
.pt-draw-phone-icon { position: absolute; right: 14px; top: 50%; width: 29px; height: 29px; margin-top: -15px; color: #5d472d; transform: rotate(-8deg); }
.pt-draw-phone-icon svg { width: 100%; height: 100%; }
.pt-draw-bottom { display: flex; align-items: center; gap: 9px; margin-top: 15px; }
.pt-draw-pawprint { width: 17px; color: #b08a53; opacity: 0.75; }
.pt-draw-pawprint svg { display: block; width: 100%; }
.pt-draw-sign { flex: 1; text-align: center; font-size: 10px; letter-spacing: 0.1em; color: #a98859; }
.pt-draw-bottomtext { margin: 18px 0 0; text-align: center; font-size: 11px; letter-spacing: 0.08em; color: #9a7644; }
@keyframes pt-draw-wobble {
  0% { opacity: 0; transform: translateY(20px) rotate(-1.5deg); }
  60% { opacity: 1; transform: translateY(-3px) rotate(0.4deg); }
  100% { opacity: 1; transform: translateY(0) rotate(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-draw-content { animation: none; }
  .pt-draw-paper, .pt-draw-phone { transition: none; }
}
</style>

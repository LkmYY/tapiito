<template>
  <div class="pt-film-root min-h-screen px-4 py-8 text-amber-50">
    <div class="pt-film-content mx-auto w-full max-w-md">
      <div class="pt-film-header">
        <span>KODA · PET FILM</span>
        <span>ROLL {{ nfcId || '0086' }} / 24 EXP</span>
      </div>

      <article class="pt-film-roll">
        <div class="pt-film-holes" aria-hidden="true">
          <span v-for="n in 12" :key="n"></span>
        </div>

        <div class="pt-film-frame">
          <div class="pt-film-photo">
            <span class="pt-film-grain" aria-hidden="true"></span>
            <span class="pt-film-light" aria-hidden="true"></span>
            <div class="pt-film-portrait" aria-hidden="true">
              <span class="pt-film-ear pt-film-ear-l"></span>
              <span class="pt-film-ear pt-film-ear-r"></span>
              <span class="pt-film-face"></span>
              <span class="pt-film-eye pt-film-eye-l"></span>
              <span class="pt-film-eye pt-film-eye-r"></span>
              <span class="pt-film-snout"></span>
              <span class="pt-film-nose"></span>
            </div>
            <span class="pt-film-frame-no">17A</span>
            <p class="pt-film-caption">home, sometime in spring</p>
          </div>

          <div class="pt-film-title">
            <span class="pt-film-role">THE STAR OF THIS ROLL</span>
            <h1>{{ payload.petName || '未命名' }}</h1>
            <p>{{ payload.breed }} · {{ payload.age || '—' }}</p>
          </div>
        </div>

        <div class="pt-film-holes pt-film-holes-b" aria-hidden="true">
          <span v-for="n in 12" :key="n"></span>
        </div>
      </article>

      <div class="pt-film-contact">
        <div class="pt-film-meta">
          <div>
            <span>GUARDIAN</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div>
            <span>ADDRESS</span>
            <p>{{ payload.address }}</p>
          </div>
        </div>

        <div class="pt-film-med">
          <span class="pt-film-med-mark" aria-hidden="true">✚</span>
          <div>
            <span>MEDICAL / 健康</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <p class="pt-film-note">“{{ payload.note || '怕生，请勿强行抱抱。' }}”</p>

        <a class="pt-film-phone" :href="'tel:' + payload.phone">
          <span class="pt-film-phone-leader">FOUND THIS FRAME? CALL</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-film-phone-sub">走失请拨 · 胶片会褪色，爱不会</span>
        </a>

        <footer class="pt-film-foot">
          <span>SHOT ON 35MM · 青禾</span>
          <span>KEEP THIS NEGATIVE SAFE</span>
        </footer>
      </div>

      <p class="pt-film-below">— every frame leads back home —</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-film-root {
  background:
    radial-gradient(circle at 75% 8%, rgba(217, 119, 6, 0.14), transparent 30%),
    repeating-linear-gradient(110deg, rgba(255, 255, 255, 0.018) 0 1px, transparent 1px 7px),
    #15110d;
  font-family: 'Courier New', 'PingFang SC', monospace;
}
.pt-film-content { animation: pt-film-project 0.8s ease both; }
.pt-film-header { display: flex; justify-content: space-between; margin: 0 5px 10px; font-size: 7px; letter-spacing: 0.14em; color: #b08a54; }
.pt-film-roll { background: #1e1912; border: 1px solid #3f3426; box-shadow: 0 22px 40px rgba(0, 0, 0, 0.55); transition: transform 0.35s ease; }
.pt-film-roll:hover { transform: translateY(-3px); }
.pt-film-holes { display: grid; grid-template-columns: repeat(12, 1fr); gap: 4px; padding: 7px 9px; background: #0e0b08; }
.pt-film-holes span { display: block; height: 8px; border-radius: 2px; background: #d9c9a8; opacity: 0.8; box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.5); }
.pt-film-holes-b { padding-top: 6px; }
.pt-film-frame { padding: 9px 11px 12px; }
.pt-film-photo { position: relative; height: 253px; overflow: hidden; background: linear-gradient(165deg, #cfb68e, #a68860 55%, #7f6646); box-shadow: inset 0 0 34px rgba(65, 42, 18, 0.55); }
.pt-film-grain { position: absolute; inset: 0; z-index: 3; background-image: radial-gradient(rgba(255, 255, 255, 0.13) 0.5px, transparent 0.8px); background-size: 3px 3px; opacity: 0.45; mix-blend-mode: overlay; }
.pt-film-light { position: absolute; inset: 0; background: linear-gradient(115deg, transparent 20%, rgba(255, 238, 198, 0.28) 45%, transparent 66%); }
.pt-film-photo::after { content: ''; position: absolute; inset: 0; background: linear-gradient(0deg, rgba(61, 39, 16, 0.55), transparent 45%); }
.pt-film-portrait { position: absolute; left: 50%; bottom: 25px; width: 140px; height: 155px; margin-left: -70px; filter: sepia(0.22) contrast(0.96); }
.pt-film-face { position: absolute; left: 17px; top: 28px; width: 106px; height: 115px; border-radius: 47% 47% 43% 43%; background: #c49a68; box-shadow: inset -14px -9px 0 rgba(70, 43, 20, 0.18); }
.pt-film-ear { position: absolute; top: 2px; width: 42px; height: 68px; background: #af8252; border-radius: 50% 50% 22% 22%; }
.pt-film-ear-l { left: 12px; transform: rotate(-19deg); }
.pt-film-ear-r { right: 12px; transform: rotate(19deg); }
.pt-film-eye { position: absolute; top: 73px; width: 10px; height: 13px; border-radius: 50%; background: #24170c; }
.pt-film-eye-l { left: 47px; }
.pt-film-eye-r { right: 47px; }
.pt-film-snout { position: absolute; left: 50%; top: 97px; width: 50px; height: 35px; margin-left: -25px; border-radius: 50%; background: #e0c29a; }
.pt-film-nose { position: absolute; left: 50%; top: 108px; width: 13px; height: 9px; margin-left: -6.5px; border-radius: 50%; background: #2d1d0f; }
.pt-film-frame-no { position: absolute; top: 8px; right: 9px; z-index: 4; font-size: 9px; font-weight: 700; letter-spacing: 0.08em; color: #f3d9a8; text-shadow: 0 1px 2px rgba(0, 0, 0, 0.6); }
.pt-film-caption { position: absolute; z-index: 4; left: 11px; bottom: 8px; margin: 0; font-family: 'Segoe Script', 'Bradley Hand', cursive; font-size: 11px; color: #f4dfb8; transform: rotate(-1deg); }
.pt-film-title { padding: 13px 2px 2px; }
.pt-film-role { display: block; font-size: 6.5px; letter-spacing: 0.2em; color: #a98653; }
.pt-film-title h1 { margin: 2px 0 0; font-family: 'Arial Narrow', 'PingFang SC', sans-serif; font-size: clamp(38px, 11.5vw, 50px); line-height: 1.04; font-weight: 700; letter-spacing: 0.02em; color: #f6e7c8; }
.pt-film-title p { margin: 5px 0 0; font-size: 8px; letter-spacing: 0.11em; color: #a58a61; }
.pt-film-contact { margin-top: 12px; padding: 15px 17px 13px; background: #f3ead8; color: #3f3322; box-shadow: 0 16px 30px rgba(0, 0, 0, 0.4); border: 1px solid #d8c9ac; }
.pt-film-meta { display: grid; grid-template-columns: 0.8fr 1.2fr; gap: 12px; padding-bottom: 9px; border-bottom: 1px dashed #c5b696; }
.pt-film-meta span { display: block; font-size: 6.5px; font-weight: 700; letter-spacing: 0.13em; color: #a58c64; }
.pt-film-meta p { margin: 2px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.45; color: #564834; }
.pt-film-med { display: flex; align-items: flex-start; gap: 8px; margin-top: 8px; padding: 7px 8px; border: 1px solid #d8b9a2; background: #fbf1e6; }
.pt-film-med-mark { flex: none; color: #a54a32; font-size: 11px; line-height: 1.2; }
.pt-film-med span { font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.09em; color: #a1785b; }
.pt-film-med p { margin: 1px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.5; color: #8c4936; }
.pt-film-note { margin: 9px 2px 11px; font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 12px; line-height: 1.7; color: #6d5c44; }
.pt-film-phone { display: block; padding: 9px 12px 10px; background: #2f261b; color: #f3e6ce; text-decoration: none; box-shadow: 3px 3px 0 #b9a988; transition: transform 0.18s ease, box-shadow 0.18s ease, background 0.18s ease; }
.pt-film-phone:hover { background: #453627; transform: translate(-1px, -1px); box-shadow: 5px 5px 0 #b9a988; }
.pt-film-phone:active { transform: translate(2px, 2px); box-shadow: 1px 1px 0 #b9a988; }
.pt-film-phone-leader { display: block; font-size: 6.5px; font-weight: 700; letter-spacing: 0.16em; color: #c8a878; }
.pt-film-phone strong { display: block; margin: 2px 0; font-family: 'Arial', 'PingFang SC', sans-serif; font-size: clamp(26px, 8vw, 33px); font-weight: 700; letter-spacing: 0.02em; color: #fdf3de; }
.pt-film-phone-sub { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.05em; color: #c2ad8c; }
.pt-film-foot { display: flex; justify-content: space-between; margin-top: 12px; padding-top: 8px; border-top: 1px solid #d8c9ac; font-size: 6px; letter-spacing: 0.09em; color: #ac9c80; }
.pt-film-below { margin: 17px 0 0; text-align: center; font-size: 8px; letter-spacing: 0.14em; color: #9a7f55; }
@keyframes pt-film-project {
  0% { opacity: 0; filter: brightness(2.4) sepia(0.4); transform: translateY(16px); }
  30% { opacity: 0.6; }
  42% { opacity: 0.35; }
  55% { opacity: 1; filter: brightness(1); }
  100% { opacity: 1; transform: translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-film-content { animation: none; }
  .pt-film-roll, .pt-film-phone { transition: none; }
}
</style>

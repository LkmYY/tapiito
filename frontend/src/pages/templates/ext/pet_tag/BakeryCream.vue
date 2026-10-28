<template>
  <div class="pt-bake-root min-h-screen px-5 py-9 text-rose-950">
    <div class="pt-bake-content mx-auto w-full max-w-md">
      <header class="pt-bake-head">
        <span>FRESHLY BAKED · 每日出炉</span>
        <span>NO. {{ nfcId || '086' }}</span>
      </header>

      <article class="pt-bake-card">
        <div class="pt-bake-scallop" aria-hidden="true">
          <span v-for="n in 13" :key="n"></span>
        </div>
        <span class="pt-bake-sprinkle pt-bake-sprinkle-a" aria-hidden="true"></span>
        <span class="pt-bake-sprinkle pt-bake-sprinkle-b" aria-hidden="true"></span>
        <span class="pt-bake-sprinkle pt-bake-sprinkle-c" aria-hidden="true"></span>
        <span class="pt-bake-sprinkle pt-bake-sprinkle-d" aria-hidden="true"></span>

        <div class="pt-bake-badge" aria-hidden="true">
          <span>100%<br /><small>SWEET</small></span>
        </div>

        <div class="pt-bake-title">
          <span class="pt-bake-eyebrow">a little bun from</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }}<span>·</span>{{ payload.age || '—' }}</p>
        </div>

        <div class="pt-bake-icing" aria-hidden="true">
          <span></span><span></span><span></span><span></span><span></span>
        </div>

        <div class="pt-bake-recipe">
          <span class="pt-bake-recipe-head">INGREDIENTS / 家庭成员</span>
          <div class="pt-bake-row">
            <span>铲屎官</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-bake-row">
            <span>烤箱地址</span>
            <p>{{ payload.address }}</p>
          </div>
        </div>

        <div class="pt-bake-med">
          <span class="pt-bake-med-heart" aria-hidden="true">♥</span>
          <div>
            <span>ALLERGY / 健康备忘</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-bake-note">
          <span class="pt-bake-note-label">BAKER'S NOTE</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。给我一点时间。' }}</p>
        </div>

        <a class="pt-bake-phone" :href="'tel:' + payload.phone">
          <span class="pt-bake-phone-top">IF I'M OUT OF THE OVEN, CALL</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-bake-phone-sub">走失请拨打 · 热乎乎的感谢</span>
        </a>

        <footer class="pt-bake-foot">
          <span>BAKED WITH LOVE · 青禾烘焙</span>
          <span>BEST BEFORE: FOREVER</span>
        </footer>
      </article>

      <p class="pt-bake-below">a tiny loaf of joy · 请带我回家</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-bake-root {
  background:
    radial-gradient(circle at 15% 10%, rgba(251, 207, 232, 0.55), transparent 28%),
    radial-gradient(circle at 87% 85%, rgba(253, 230, 138, 0.4), transparent 30%),
    #fff1e8;
  font-family: 'Avenir Next', 'PingFang SC', sans-serif;
}
.pt-bake-content { animation: pt-bake-rise 0.8s cubic-bezier(0.2, 0.9, 0.3, 1.18) both; }
.pt-bake-head { display: flex; justify-content: space-between; margin: 0 6px 13px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.13em; color: #b07a68; }
.pt-bake-card { position: relative; margin-top: 10px; padding: 27px 18px 15px; background: #fffaf4; border: 1px solid #f0d3c4; border-radius: 20px; box-shadow: 0 22px 40px rgba(167, 98, 75, 0.17); transition: transform 0.35s ease, box-shadow 0.35s ease; }
.pt-bake-card:hover { transform: translateY(-4px) rotate(-0.2deg); box-shadow: 0 30px 50px rgba(167, 98, 75, 0.22); }
.pt-bake-scallop { position: absolute; top: -10px; left: 8px; right: 8px; display: flex; justify-content: space-between; }
.pt-bake-scallop span { display: block; width: 27px; height: 27px; border-radius: 50%; background: #f7bfc5; box-shadow: inset 0 -4px 0 rgba(199, 112, 120, 0.18); }
.pt-bake-scallop span:nth-child(even) { background: #fbdfe1; }
.pt-bake-sprinkle { position: absolute; display: block; width: 9px; height: 3px; border-radius: 3px; }
.pt-bake-sprinkle-a { top: 73px; right: 42px; background: #f2a6b4; transform: rotate(31deg); }
.pt-bake-sprinkle-b { top: 105px; right: 23px; background: #f4c45e; transform: rotate(-22deg); }
.pt-bake-sprinkle-c { top: 139px; right: 51px; background: #a8c9b0; transform: rotate(62deg); }
.pt-bake-sprinkle-d { top: 51px; left: 37px; background: #d3b7e0; transform: rotate(13deg); }
.pt-bake-badge { position: absolute; top: 18px; right: 18px; width: 52px; height: 52px; display: grid; place-items: center; border: 1.5px dashed #d9958a; border-radius: 50%; color: #bb7367; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; line-height: 1.1; text-align: center; transform: rotate(8deg); }
.pt-bake-badge small { font-size: 5px; letter-spacing: 0.1em; }
.pt-bake-title { padding: 4px 50px 8px 3px; }
.pt-bake-eyebrow { font-family: 'Segoe Script', 'Bradley Hand', cursive; font-size: 12px; color: #c19484; }
.pt-bake-title h1 { margin: 2px 0 0; font-size: clamp(43px, 12.8vw, 58px); line-height: 1.02; font-weight: 600; letter-spacing: 0.04em; color: #703d35; text-shadow: 0 2px 0 #fbe0d6; word-break: break-word; }
.pt-bake-title p { margin: 7px 0 0; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.1em; color: #ad8679; }
.pt-bake-title p span { margin: 0 5px; color: #d8b6a9; }
.pt-bake-icing { display: flex; gap: 3px; height: 18px; margin: 4px -18px 12px; overflow: hidden; }
.pt-bake-icing span { flex: 1; background: #f9c9ce; border-radius: 0 0 50% 50%; }
.pt-bake-icing span:nth-child(even) { background: #fde3e4; border-radius: 0 0 35% 65%; }
.pt-bake-recipe { padding: 10px 11px; border: 1px dashed #dfb5a6; border-radius: 11px; background: #fff5ed; }
.pt-bake-recipe-head { display: block; margin-bottom: 5px; padding-bottom: 4px; border-bottom: 1px solid #f0d5c8; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.12em; color: #b27e6d; }
.pt-bake-row { display: grid; grid-template-columns: 63px 1fr; gap: 8px; padding: 3px 0; font-size: 10px; line-height: 1.5; }
.pt-bake-row span { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; color: #b79a8f; padding-top: 1px; }
.pt-bake-row p { margin: 0; color: #704b42; }
.pt-bake-med { display: flex; gap: 8px; align-items: flex-start; margin-top: 8px; padding: 9px 10px; border-radius: 11px; background: #fdecec; border: 1px solid #f2c9c4; }
.pt-bake-med-heart { flex: none; color: #d47c72; font-size: 12px; line-height: 1.2; }
.pt-bake-med span { font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.07em; color: #b17168; }
.pt-bake-med p { margin: 1px 0 0; font-size: 10px; line-height: 1.5; color: #8f4b42; }
.pt-bake-note { margin: 9px 2px 12px; }
.pt-bake-note-label { display: inline-block; margin-bottom: 3px; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.1em; color: #bd8a79; }
.pt-bake-note p { margin: 0; font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 12px; line-height: 1.7; color: #75564b; }
.pt-bake-phone { display: block; padding: 11px 13px 12px; border-radius: 14px; background: #b7665e; color: #fff8f3; text-decoration: none; box-shadow: 0 4px 0 #945049; transition: transform 0.18s ease, box-shadow 0.18s ease, background 0.18s ease; }
.pt-bake-phone:hover { background: #c0736a; transform: translateY(-1px); box-shadow: 0 5px 0 #945049; }
.pt-bake-phone:active { transform: translateY(3px); box-shadow: 0 1px 0 #945049; }
.pt-bake-phone-top { display: block; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.12em; color: #fbded6; }
.pt-bake-phone strong { display: block; margin: 2px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(27px, 8.2vw, 34px); font-weight: 700; letter-spacing: 0.02em; color: #fffdfa; }
.pt-bake-phone-sub { font-size: 8px; letter-spacing: 0.05em; color: #f2cabf; }
.pt-bake-foot { display: flex; justify-content: space-between; margin-top: 13px; padding-top: 9px; border-top: 1px dashed #eccfc2; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.07em; color: #b89a8e; }
.pt-bake-below { margin: 18px 0 0; text-align: center; font-family: 'Segoe Script', 'Bradley Hand', cursive; font-size: 12px; color: #b38273; }
@keyframes pt-bake-rise {
  0% { opacity: 0; transform: translateY(26px) scale(0.96); }
  60% { opacity: 1; transform: translateY(-3px) scale(1.008); }
  100% { opacity: 1; transform: translateY(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-bake-content { animation: none; }
  .pt-bake-card, .pt-bake-phone { transition: none; }
}
</style>

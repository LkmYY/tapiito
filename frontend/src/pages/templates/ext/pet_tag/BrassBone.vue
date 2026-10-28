<template>
  <div class="pt-brass-root min-h-screen px-5 py-9 text-amber-50">
    <div class="pt-brass-content mx-auto w-full max-w-md">
      <div class="pt-brass-chain" aria-hidden="true">
        <span v-for="n in 7" :key="n" class="pt-brass-link"></span>
      </div>

      <div class="pt-brass-ring" aria-hidden="true"></div>

      <article class="pt-brass-tag">
        <div class="pt-brass-shine" aria-hidden="true"></div>

        <header class="pt-brass-top">
          <span class="pt-brass-brand">PAW &amp; BRASS</span>
          <span class="pt-brass-no">NO. {{ nfcId || '0086' }}</span>
        </header>

        <p class="pt-brass-if">IF FOUND, PLEASE CALL</p>
        <a class="pt-brass-phone" :href="'tel:' + payload.phone">
          {{ payload.phone }}
        </a>

        <div class="pt-brass-nameplate">
          <span class="pt-brass-paw" aria-hidden="true">
            <svg viewBox="0 0 48 48" fill="none">
              <ellipse cx="24" cy="31" rx="10" ry="8" fill="currentColor" />
              <ellipse cx="11" cy="21" rx="4" ry="5" fill="currentColor" />
              <ellipse cx="21" cy="15" rx="4" ry="5.5" fill="currentColor" />
              <ellipse cx="32" cy="16" rx="4" ry="5" fill="currentColor" />
              <ellipse cx="40" cy="24" rx="3.4" ry="4.4" fill="currentColor" />
            </svg>
          </span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }}<span class="pt-brass-dot">·</span>{{ payload.age || '—' }}</p>
        </div>

        <div class="pt-brass-divider"><span></span></div>

        <dl class="pt-brass-info">
          <div>
            <dt>OWNER</dt>
            <dd>{{ payload.ownerName }}</dd>
          </div>
          <div>
            <dt>HOME</dt>
            <dd>{{ payload.address }}</dd>
          </div>
        </dl>

        <div class="pt-brass-medical">
          <span class="pt-brass-medical-key">MEDICAL</span>
          <span>{{ payload.medical || '无特殊记录' }}</span>
        </div>

        <p class="pt-brass-note">“{{ payload.note || '请温柔对待我，帮我回家。' }}”</p>

        <footer class="pt-brass-foot">
          <span>REWARD · 必有重谢</span>
          <span class="pt-brass-serial">{{ nfcId || 'PT-2026' }}</span>
        </footer>
      </article>

      <p class="pt-brass-tap">轻触铭牌 · 电话立即拨出</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-brass-root {
  position: relative;
  overflow: hidden;
  background:
    radial-gradient(circle at 50% 4%, rgba(255, 214, 140, 0.18), transparent 38%),
    radial-gradient(circle at 12% 88%, rgba(120, 67, 24, 0.34), transparent 46%),
    linear-gradient(160deg, #17100a 0%, #2a1a0e 48%, #120c07 100%);
  font-family: 'Optima', 'Palatino Linotype', 'Songti SC', Georgia, serif;
}
.pt-brass-root::after {
  content: '';
  position: absolute;
  inset: 0;
  pointer-events: none;
  background-image: radial-gradient(rgba(255, 226, 170, 0.13) 0.7px, transparent 0.7px);
  background-size: 7px 7px;
  opacity: 0.4;
}
.pt-brass-content { position: relative; z-index: 1; animation: pt-brass-drop 0.9s cubic-bezier(0.2, 0.8, 0.2, 1) both; }
.pt-brass-chain { display: flex; flex-direction: column; align-items: center; gap: 0; }
.pt-brass-link {
  width: 15px;
  height: 21px;
  border: 2px solid #b98a45;
  border-radius: 8px;
  box-shadow: inset 0 0 0 1px rgba(255, 229, 173, 0.35);
  margin-top: -6px;
  transform: rotate(4deg);
}
.pt-brass-link:nth-child(even) { transform: rotate(-7deg); }
.pt-brass-ring {
  width: 30px;
  height: 30px;
  margin: -2px auto 0;
  border: 6px solid #caa15f;
  border-radius: 50%;
  box-shadow: 0 0 0 2px #6d4a1e, inset 0 0 8px rgba(52, 31, 9, 0.8), 0 6px 14px rgba(0, 0, 0, 0.5);
}
.pt-brass-tag {
  position: relative;
  margin-top: -3px;
  padding: 30px 26px 22px;
  border-radius: 26px 26px 38px 38px;
  color: #3b2408;
  background:
    linear-gradient(115deg, rgba(255, 255, 255, 0.5) 0%, transparent 26%),
    repeating-linear-gradient(100deg, rgba(120, 76, 20, 0.06) 0 2px, transparent 2px 7px),
    linear-gradient(150deg, #f4d693 0%, #c9964c 35%, #e8c27a 56%, #a87432 100%);
  border: 1px solid #f9e6b4;
  box-shadow: 0 26px 50px rgba(0, 0, 0, 0.58), inset 0 0 0 5px rgba(91, 55, 13, 0.17), inset 0 0 34px rgba(96, 57, 14, 0.3);
  overflow: hidden;
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.pt-brass-tag::before {
  content: '';
  position: absolute;
  inset: 12px;
  border: 1px dashed rgba(71, 43, 9, 0.4);
  border-radius: 18px 18px 28px 28px;
  pointer-events: none;
}
.pt-brass-tag:hover { transform: translateY(-4px) rotate(-0.4deg); box-shadow: 0 34px 64px rgba(0, 0, 0, 0.64), inset 0 0 0 5px rgba(91, 55, 13, 0.17); }
.pt-brass-shine {
  position: absolute;
  top: -60%;
  left: -30%;
  width: 55%;
  height: 200%;
  background: linear-gradient(90deg, transparent, rgba(255, 250, 225, 0.5), transparent);
  transform: rotate(18deg);
  animation: pt-brass-shine 5.5s ease-in-out 1.2s infinite;
  pointer-events: none;
}
.pt-brass-top { position: relative; display: flex; justify-content: space-between; font-family: 'Courier New', monospace; font-size: 9px; letter-spacing: 0.16em; color: #5c3c13; }
.pt-brass-if { position: relative; margin: 22px 0 4px; text-align: center; font-size: 10px; letter-spacing: 0.35em; font-weight: 700; color: #6a4518; }
.pt-brass-phone {
  position: relative;
  display: block;
  text-align: center;
  font-family: 'Arial Black', 'Helvetica Neue', sans-serif;
  font-size: clamp(25px, 7.4vw, 32px);
  letter-spacing: 0.02em;
  line-height: 1.15;
  color: #2d1a04;
  text-decoration: none;
  text-shadow: 0 1px 0 rgba(255, 245, 214, 0.85), 0 -1px 0 rgba(73, 43, 8, 0.3);
  transition: transform 0.2s ease;
}
.pt-brass-phone:active { transform: scale(0.97); }
.pt-brass-phone:hover { color: #0f0902; }
.pt-brass-nameplate { position: relative; margin-top: 22px; text-align: center; }
.pt-brass-paw { display: block; width: 38px; margin: 0 auto 6px; color: rgba(74, 45, 10, 0.65); }
.pt-brass-paw svg { display: block; width: 100%; height: auto; }
.pt-brass-nameplate h1 {
  margin: 0;
  font-size: clamp(42px, 12vw, 56px);
  line-height: 1;
  font-weight: 600;
  letter-spacing: 0.12em;
  text-indent: 0.12em;
  text-shadow: 0 1px 0 #fceec6, 0 2px 1px rgba(76, 45, 9, 0.3);
}
.pt-brass-nameplate p { margin: 9px 0 0; font-size: 12px; letter-spacing: 0.2em; color: #5b3b10; }
.pt-brass-dot { margin: 0 7px; color: #9a6c2a; }
.pt-brass-divider { display: flex; align-items: center; gap: 8px; margin: 20px 0 14px; }
.pt-brass-divider::before, .pt-brass-divider::after { content: ''; flex: 1; height: 1px; background: rgba(78, 48, 9, 0.36); }
.pt-brass-divider span { width: 6px; height: 6px; border: 1px solid #6f4714; transform: rotate(45deg); }
.pt-brass-info { display: grid; grid-template-columns: 1fr 1.35fr; gap: 14px; margin: 0; }
.pt-brass-info dt { font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.24em; color: #79501b; margin-bottom: 4px; }
.pt-brass-info dd { margin: 0; font-size: 13px; line-height: 1.5; font-weight: 600; color: #3c2508; }
.pt-brass-medical {
  display: flex;
  gap: 9px;
  align-items: baseline;
  margin-top: 17px;
  padding: 9px 11px;
  border: 1px solid rgba(86, 50, 7, 0.4);
  background: rgba(255, 241, 201, 0.35);
  font-size: 12px;
  line-height: 1.5;
}
.pt-brass-medical-key { flex: none; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.18em; padding: 2px 5px; color: #f7e7bc; background: #5e3c0d; }
.pt-brass-note { margin: 14px 3px 0; font-size: 12px; line-height: 1.65; font-style: italic; color: #5f3e12; text-align: center; }
.pt-brass-foot { display: flex; justify-content: space-between; align-items: center; margin-top: 18px; padding-top: 12px; border-top: 1px solid rgba(83, 49, 10, 0.35); font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.15em; }
.pt-brass-serial { color: #6c4512; }
.pt-brass-tap { margin: 20px 0 0; text-align: center; font-size: 10px; letter-spacing: 0.22em; color: rgba(244, 214, 147, 0.55); }
@keyframes pt-brass-drop {
  from { opacity: 0; transform: translateY(-34px) rotate(-3deg); }
  to { opacity: 1; transform: translateY(0) rotate(0); }
}
@keyframes pt-brass-shine {
  0%, 55% { transform: translateX(0) rotate(18deg); opacity: 0; }
  62% { opacity: 0.8; }
  82%, 100% { transform: translateX(420%) rotate(18deg); opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .pt-brass-content, .pt-brass-shine { animation: none; }
  .pt-brass-tag, .pt-brass-phone { transition: none; }
}
</style>

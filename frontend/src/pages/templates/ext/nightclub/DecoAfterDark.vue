<template>
  <div class="nc-da-root relative min-h-screen overflow-hidden px-6 py-8 text-[#e7d6ab]">
    <div class="nc-da-fan" aria-hidden="true"></div>
    <div class="nc-da-deco nc-da-deco-l"></div>
    <div class="nc-da-deco nc-da-deco-r"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-da-head">
        <span class="nc-da-rule"></span>
        <span>AFTER DARK</span>
        <span class="nc-da-rule"></span>
      </header>

      <main class="nc-da-main">
        <div class="nc-da-crest" aria-hidden="true">
          <span class="nc-da-crest-ray"></span>
          <span class="nc-da-diamond">◆</span>
        </div>

        <p class="nc-da-alias">{{ payload?.alias || 'NO ALIAS' }}</p>
        <h1 class="nc-da-name">{{ payload?.name }}</h1>
        <p class="nc-da-role">{{ payload?.role }}</p>

        <div class="nc-da-ornament">
          <span></span><i>◇</i><span></span>
        </div>

        <p class="nc-da-venue">{{ payload?.venue }}</p>
        <p class="nc-da-city">{{ payload?.city || '—' }} · MEMBERS' SALON</p>

        <section class="nc-da-contacts">
          <div class="nc-da-contact">
            <span class="nc-da-contact-key">WECHAT</span>
            <span class="nc-da-contact-line"></span>
            <span class="nc-da-contact-val">{{ payload?.wechat }}</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-da-contact nc-da-contact-link">
            <span class="nc-da-contact-key">TELEPHONE</span>
            <span class="nc-da-contact-line"></span>
            <span class="nc-da-contact-val">{{ payload.phone }}</span>
          </a>
          <div v-else class="nc-da-contact">
            <span class="nc-da-contact-key">TELEPHONE</span>
            <span class="nc-da-contact-line"></span>
            <span class="nc-da-contact-val nc-da-empty"></span>
          </div>
        </section>

        <p class="nc-da-quote">“The last door is always the most beautiful.”</p>
      </main>

      <footer class="nc-da-foot">
        <span>EST. MCMXCVIII</span>
        <span class="nc-da-foot-mark">✦</span>
        <span>NFC {{ nfcId || '0000' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-da-root {
  background: #06231c;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
}
.nc-da-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 50% 12%, rgba(201, 169, 89, 0.15), transparent 46%),
    linear-gradient(180deg, rgba(6, 35, 28, 0) 40%, rgba(3, 20, 16, 0.75));
  pointer-events: none;
}
.nc-da-fan {
  position: absolute;
  top: -80px;
  left: 50%;
  width: 420px;
  height: 300px;
  border-radius: 50%;
  background: repeating-conic-gradient(from 240deg at 50% 100%, rgba(201, 169, 89, 0.14) 0 2deg, transparent 2deg 10deg);
  mask-image: radial-gradient(ellipse at 50% 100%, #000 10%, transparent 68%);
  transform: translateX(-50%);
  pointer-events: none;
  animation: nc-da-fan 1.2s 0.1s both;
}
.nc-da-deco {
  position: absolute;
  top: 20%;
  bottom: 20%;
  width: 1px;
  background: linear-gradient(180deg, transparent, rgba(201, 169, 89, 0.28), transparent);
  pointer-events: none;
}
.nc-da-deco-l { left: 11px; }
.nc-da-deco-r { right: 11px; }
.nc-da-head {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 8px;
  letter-spacing: 0.4em;
  color: rgba(231, 214, 171, 0.55);
  animation: nc-da-in 0.9s both;
}
.nc-da-rule {
  flex: 1;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(201, 169, 89, 0.5));
}
.nc-da-rule:last-child {
  background: linear-gradient(90deg, rgba(201, 169, 89, 0.5), transparent);
}
.nc-da-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 30px 0 24px;
  text-align: center;
}
.nc-da-crest {
  position: relative;
  display: flex;
  width: 56px;
  height: 56px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(201, 169, 89, 0.55);
  transform: rotate(45deg);
  animation: nc-da-crest 1s 0.08s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-da-crest-ray {
  position: absolute;
  inset: 5px;
  border: 1px solid rgba(201, 169, 89, 0.22);
}
.nc-da-diamond {
  transform: rotate(-45deg);
  font-size: 13px;
  color: #c9a959;
}
.nc-da-alias {
  margin-top: 28px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.5em;
  color: #c9a959;
  text-transform: uppercase;
  animation: nc-da-in 0.9s 0.18s both;
}
.nc-da-name {
  margin-top: 10px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(50px, 16vw, 68px);
  font-weight: 400;
  line-height: 0.98;
  letter-spacing: 0.01em;
  color: #f0e2c0;
  text-shadow: 0 0 40px rgba(201, 169, 89, 0.25);
  overflow-wrap: anywhere;
  animation: nc-da-in 1s 0.22s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-da-role {
  margin-top: 12px;
  font-size: 8.5px;
  letter-spacing: 0.34em;
  color: rgba(231, 214, 171, 0.6);
  text-transform: uppercase;
  animation: nc-da-in 0.9s 0.3s both;
}
.nc-da-ornament {
  display: flex;
  width: 68%;
  align-items: center;
  gap: 10px;
  margin: 28px auto 20px;
  animation: nc-da-in 0.9s 0.36s both;
}
.nc-da-ornament span {
  flex: 1;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(201, 169, 89, 0.55));
}
.nc-da-ornament span:last-child {
  background: linear-gradient(90deg, rgba(201, 169, 89, 0.55), transparent);
}
.nc-da-ornament i {
  font-size: 7px;
  font-style: normal;
  color: #c9a959;
}
.nc-da-venue {
  font-size: 13px;
  letter-spacing: 0.26em;
  color: #e7d6ab;
  text-transform: uppercase;
  overflow-wrap: anywhere;
  animation: nc-da-in 0.9s 0.4s both;
}
.nc-da-city {
  margin-top: 5px;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: rgba(231, 214, 171, 0.42);
  animation: nc-da-in 0.9s 0.44s both;
}
.nc-da-contacts {
  width: 100%;
  margin-top: 34px;
  border-top: 1px solid rgba(201, 169, 89, 0.3);
  border-bottom: 1px solid rgba(201, 169, 89, 0.3);
  padding: 4px 0;
  animation: nc-da-in 1s 0.5s both;
}
.nc-da-contact {
  display: flex;
  align-items: baseline;
  gap: 10px;
  padding: 10px 4px;
  text-decoration: none;
  color: inherit;
}
.nc-da-contact + .nc-da-contact {
  border-top: 1px solid rgba(201, 169, 89, 0.12);
}
.nc-da-contact-key {
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(201, 169, 89, 0.55);
}
.nc-da-contact-line {
  flex: 1;
  height: 1px;
  background: repeating-linear-gradient(90deg, rgba(201, 169, 89, 0.3) 0 2px, transparent 2px 5px);
}
.nc-da-contact-val {
  min-width: 0;
  max-width: 58%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10.5px;
  letter-spacing: 0.05em;
  color: #eddfbc;
  text-align: right;
}
.nc-da-contact-link {
  margin: 0 -6px;
  padding: 10px 10px;
  transition: background 0.3s ease, letter-spacing 0.3s ease;
}
.nc-da-contact-link:hover {
  background: rgba(201, 169, 89, 0.1);
}
.nc-da-contact-link:hover .nc-da-contact-val {
  color: #ffe9ae;
  letter-spacing: 0.1em;
}
.nc-da-contact-link:active {
  background: rgba(201, 169, 89, 0.18);
}
.nc-da-empty {
  color: rgba(231, 214, 171, 0.28);
}
.nc-da-quote {
  margin-top: 24px;
  font-size: 9.5px;
  font-style: italic;
  letter-spacing: 0.04em;
  color: rgba(231, 214, 171, 0.4);
  animation: nc-da-in 1s 0.6s both;
}
.nc-da-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 14px;
  border-top: 1px solid rgba(201, 169, 89, 0.22);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(231, 214, 171, 0.3);
  animation: nc-da-in 1s 0.7s both;
}
.nc-da-foot-mark {
  color: #c9a959;
}
@keyframes nc-da-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-da-crest {
  from { opacity: 0; transform: rotate(45deg) scale(0.7); }
  to { opacity: 1; transform: rotate(45deg) scale(1); }
}
@keyframes nc-da-fan {
  from { opacity: 0; transform: translateX(-50%) scaleY(0.6); }
  to { opacity: 1; transform: translateX(-50%) scaleY(1); }
}
@media (prefers-reduced-motion: reduce) {
  .nc-da-root *,
  .nc-da-root *::before,
  .nc-da-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>

<template>
  <div class="relative min-h-screen overflow-hidden bg-[#fbf7f4] text-[#3a2f2c]">
    <!-- Soft light -->
    <div class="pointer-events-none absolute -left-24 -top-24 h-64 w-64 rounded-full bg-rose-100/70 blur-3xl"></div>
    <div class="pointer-events-none absolute -right-28 top-1/3 h-60 w-60 rounded-full bg-[#f3e4d5]/80 blur-3xl"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-10">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-[#c0a99f]">
        <span>Salon</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Identity -->
      <div class="rise rise-1 mt-10 text-center">
        <h1 class="text-balance font-serif text-[34px] leading-tight tracking-tight text-[#3a2f2c]">
          {{ payload?.salonName || '美发沙龙' }}
        </h1>
        <p class="mt-3 text-[11px] uppercase tracking-[0.3em] text-[#9f1239]/70">专业美业 · 品质服务</p>
        <div class="mt-7 flex items-center gap-3">
          <span class="h-px flex-1 bg-[#ecdcd3]"></span>
          <span class="h-1.5 w-1.5 rotate-45 bg-rose-300"></span>
          <span class="h-px flex-1 bg-[#ecdcd3]"></span>
        </div>
      </div>

      <!-- Price list -->
      <div v-if="services.length" class="rise rise-2 mt-9">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-semibold tracking-wide text-[#3a2f2c]">服务价目</h2>
          <span class="h-px flex-1 bg-[#ecdcd3]"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-[#c0a99f]">Price List</span>
        </div>
        <div class="sl-panel mt-4">
          <div v-for="(s, i) in services" :key="i" class="sl-row">
            <span class="w-6 flex-shrink-0 font-mono text-[10px] tabular-nums text-[#cbb5ab]">{{ String(i + 1).padStart(2, '0') }}</span>
            <div class="min-w-0 flex-1">
              <p class="truncate text-[15px] font-medium text-[#3a2f2c]">{{ s?.name || '服务项目' }}</p>
              <p v-if="s?.duration" class="mt-0.5 text-[11px] tracking-[0.08em] text-[#a98f86]">{{ s.duration }}</p>
            </div>
            <p v-if="s?.price || s?.price === 0" class="flex-shrink-0 text-[15px] font-semibold tabular-nums text-[#9f1239]">
              ¥{{ s.price }}
            </p>
          </div>
        </div>
      </div>

      <!-- Gallery -->
      <div class="rise rise-3 mt-9">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-semibold tracking-wide text-[#3a2f2c]">作品展示</h2>
          <span class="h-px flex-1 bg-[#ecdcd3]"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-[#c0a99f]">Gallery</span>
        </div>
        <div v-if="galleryImages.length" class="mt-4 grid grid-cols-2 gap-2">
          <div v-for="(img, i) in galleryImages" :key="i" class="sl-shot" :class="i === 0 ? 'col-span-2 is-wide' : ''">
            <img :src="img" :alt="'作品 ' + (i + 1)" class="h-full w-full object-cover" />
          </div>
        </div>
        <div v-else class="sl-empty mt-4">
          <Camera class="h-5 w-5 text-[#d4beb4]" />
          <p class="mt-2 text-xs text-[#bda79d]">暂无作品图片</p>
        </div>
      </div>

      <!-- Contact -->
      <div class="rise rise-3 mt-9">
        <div v-if="payload?.address" class="sl-panel flex items-start gap-3.5 p-5">
          <MapPin class="mt-0.5 h-4 w-4 flex-shrink-0 text-[#9f1239]/80" />
          <p class="text-[13px] leading-relaxed text-[#6b5a52]">{{ payload.address }}</p>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="sl-cta mt-4">
          <Phone class="h-4 w-4" />
          立即预约
        </a>
      </div>

      <!-- Footer -->
      <div class="rise rise-3 mt-12 border-t border-[#ecdcd3] pt-5 text-center">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-[#c0a99f]">{{ payload?.salonName || 'Salon' }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Camera, MapPin, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && (item.name || item.price)) : []
})

const galleryImages = computed(() => {
  const images = props.payload?.gallery
  if (!Array.isArray(images)) return []
  return images.filter((src) => typeof src === 'string' && src.trim() !== '')
})
</script>

<style scoped>
.sl-panel {
  border-radius: 16px;
  border: 1px solid #f0e4dc;
  background: #ffffff;
  box-shadow: 0 20px 44px -34px rgba(159, 18, 57, 0.45);
}

.sl-row {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 18px;
  border-bottom: 1px solid #f6ece6;
}
.sl-row:last-child {
  border-bottom: none;
}

.sl-shot {
  aspect-ratio: 1 / 1;
  overflow: hidden;
  border-radius: 14px;
  border: 1px solid #f0e4dc;
  background: #ffffff;
}
.sl-shot.is-wide {
  aspect-ratio: 16 / 10;
}
.sl-shot img {
  transition: transform 0.55s cubic-bezier(0.22, 1, 0.36, 1);
}
.sl-shot:hover img {
  transform: scale(1.03);
}

.sl-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 36px 0;
  border: 1px dashed #e6d5cc;
  border-radius: 16px;
  background: #ffffff;
}

.sl-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  border-radius: 9999px;
  background: #9f1239;
  padding: 14px 0;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.08em;
  color: #fff7f8;
  box-shadow: 0 18px 36px -18px rgba(159, 18, 57, 0.55);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.sl-cta:hover {
  background: #881337;
}
.sl-cta:active {
  transform: scale(0.98);
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.rise {
  animation: rise 0.65s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.rise-1 {
  animation-delay: 0.08s;
}
.rise-2 {
  animation-delay: 0.16s;
}
.rise-3 {
  animation-delay: 0.24s;
}
</style>

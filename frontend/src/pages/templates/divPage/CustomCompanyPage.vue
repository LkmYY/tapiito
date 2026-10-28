<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f5f6f8] text-slate-900">
    <!-- Accent rule -->
    <div class="h-[3px] w-full bg-[#1d4ed8]"></div>
    <!-- Ambient light -->
    <div class="pointer-events-none absolute -right-28 -top-20 h-72 w-72 rounded-full bg-blue-200/35 blur-3xl"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-8">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">
        <span>Company Profile</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Identity -->
      <div class="rise rise-1 mt-9">
        <h1 class="text-balance text-[28px] font-bold leading-tight tracking-tight">{{ payload?.companyName || '企业主页' }}</h1>
        <p v-if="payload?.description" class="mt-3 text-sm leading-relaxed text-slate-500">{{ payload.description }}</p>
      </div>

      <!-- Gallery -->
      <div class="rise rise-2 mt-8">
        <template v-if="galleryImages.length">
          <button type="button" class="cmp-shot group" @click="openImageModal(0)">
            <img :src="galleryImages[0]" alt="公司图片 1" />
          </button>
          <div v-if="restImages.length" class="mt-2 grid grid-cols-2 gap-2">
            <button
              v-for="(img, i) in restImages"
              :key="i"
              type="button"
              class="cmp-tile group"
              @click="openImageModal(i + 1)"
            >
              <img :src="img" :alt="'公司图片 ' + (i + 2)" />
              <span v-if="i === restImages.length - 1 && moreCount > 0" class="cmp-more">+{{ moreCount }}</span>
            </button>
          </div>
        </template>
        <div v-else class="cmp-empty">
          <ImageIcon class="h-5 w-5 text-slate-300" />
          <p class="mt-2 text-xs text-slate-400">暂无图片</p>
        </div>
      </div>

      <!-- Services -->
      <div v-if="services.length" class="rise rise-3 mt-10">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-semibold tracking-wide text-slate-900">服务项目</h2>
          <span class="h-px flex-1 bg-slate-200"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Services</span>
        </div>
        <div class="mt-4 grid grid-cols-2 gap-2.5">
          <div v-for="(service, i) in services" :key="i" class="cmp-service">
            <span class="cmp-service-icon">
              <component :is="resolveServiceIcon(service?.icon)" class="h-[18px] w-[18px]" />
            </span>
            <p class="mt-3 text-[13px] font-semibold leading-snug text-slate-800">{{ service?.name || '服务项目' }}</p>
          </div>
        </div>
      </div>

      <!-- Contact -->
      <div v-if="payload?.phone" class="rise rise-3 mt-10">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-semibold tracking-wide text-slate-900">联系方式</h2>
          <span class="h-px flex-1 bg-slate-200"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Contact</span>
        </div>
        <div class="cmp-card mt-4">
          <a :href="'tel:' + payload.phone" class="cmp-contact group">
            <span class="cmp-service-icon">
              <Phone class="h-[18px] w-[18px]" />
            </span>
            <span class="min-w-0 flex-1">
              <span class="block font-mono text-[10px] uppercase tracking-[0.2em] text-slate-400">电话咨询</span>
              <span class="mt-1 block text-[15px] font-semibold tracking-wide text-slate-800 tabular-nums">{{ payload.phone }}</span>
            </span>
            <ArrowUpRight class="h-4 w-4 flex-shrink-0 text-slate-300 transition-colors group-hover:text-[#1d4ed8]" />
          </a>
          <a :href="'tel:' + payload.phone" class="cmp-cta mt-4">
            <PhoneCall class="h-4 w-4" />
            立即咨询
          </a>
        </div>
      </div>

      <!-- Footer -->
      <div class="rise rise-3 mt-12 border-t border-slate-200 pt-5 text-center">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">{{ payload?.companyName || '企业主页' }}</p>
        <p v-if="nfcId" class="mt-1.5 font-mono text-[10px] tracking-[0.18em] text-slate-300">NFC · {{ nfcId }}</p>
      </div>
    </div>

    <!-- Lightbox -->
    <div v-if="showModal" class="fixed inset-0 z-[200] flex flex-col bg-slate-950/95" @click="closeModal">
      <div class="flex items-center justify-between px-5 py-4">
        <p class="font-mono text-[11px] tracking-[0.2em] text-white/50 tabular-nums">{{ currentImageIndex + 1 }} / {{ galleryImages.length }}</p>
        <button type="button" class="cmp-lb-btn" aria-label="关闭" @click.stop="closeModal">
          <X class="h-5 w-5" />
        </button>
      </div>
      <div class="flex flex-1 items-center justify-center px-4 pb-4" @click.stop>
        <img
          :src="galleryImages[currentImageIndex]"
          :alt="'公司图片 ' + (currentImageIndex + 1)"
          class="max-h-[76vh] w-auto max-w-full rounded-xl object-contain"
        />
      </div>
      <div class="flex items-center justify-center gap-6 pb-8">
        <button
          type="button"
          class="cmp-lb-btn disabled:opacity-30"
          aria-label="上一张"
          :disabled="currentImageIndex === 0"
          @click.stop="prevImage"
        >
          <ChevronLeft class="h-5 w-5" />
        </button>
        <button
          type="button"
          class="cmp-lb-btn disabled:opacity-30"
          aria-label="下一张"
          :disabled="currentImageIndex >= galleryImages.length - 1"
          @click.stop="nextImage"
        >
          <ChevronRight class="h-5 w-5" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import {
  ArrowRight,
  ArrowUpRight,
  Award,
  Briefcase,
  Building2,
  Cake,
  Calendar,
  Camera,
  ChefHat,
  ChevronLeft,
  ChevronRight,
  Clock,
  Coffee,
  Crown,
  Droplet,
  Dumbbell,
  Flame,
  Gem,
  Gift,
  Globe,
  Hammer,
  Heart,
  Image as ImageIcon,
  Leaf,
  Lightbulb,
  Mail,
  MapPin,
  Moon,
  Music,
  Package,
  Paintbrush,
  Palette,
  PartyPopper,
  Percent,
  Phone,
  PhoneCall,
  Rocket,
  Scissors,
  Settings,
  ShieldCheck,
  Sparkles,
  Star,
  Sun,
  Target,
  Truck,
  Users,
  UtensilsCrossed,
  Wrench,
  X,
  Zap
} from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

/**
 * iconMap：把 payload.services[].icon 里的小写词映射到 lucide 图标。
 * 命中即渲染对应图标；未命中（含旧数据里的 emoji）统一回退为 ArrowRight 装饰标，
 * 绝不把原始字符串（emoji 文本）渲染到页面。
 */
const iconMap = {
  wrench: Wrench,
  sparkles: Sparkles,
  heart: Heart,
  truck: Truck,
  phone: Phone,
  mail: Mail,
  globe: Globe,
  star: Star,
  shield: ShieldCheck,
  zap: Zap,
  leaf: Leaf,
  gift: Gift,
  cake: Cake,
  camera: Camera,
  palette: Palette,
  scissors: Scissors,
  music: Music,
  calendar: Calendar,
  clock: Clock,
  users: Users,
  award: Award,
  briefcase: Briefcase,
  lightbulb: Lightbulb,
  flame: Flame,
  droplet: Droplet,
  coffee: Coffee,
  'map-pin': MapPin,
  settings: Settings,
  rocket: Rocket,
  target: Target,
  package: Package,
  percent: Percent,
  party: PartyPopper,
  crown: Crown,
  gem: Gem,
  dumbbell: Dumbbell,
  utensils: UtensilsCrossed,
  chef: ChefHat,
  building: Building2,
  hammer: Hammer,
  paint: Paintbrush,
  sun: Sun,
  moon: Moon
}

// 归一化：'MapPin' / 'map_pin' / 'map pin' / 'sparkles' 都能匹配到同一个键
const normalizeIconKey = (icon) =>
  String(icon || '')
    .trim()
    .replace(/([a-z0-9])([A-Z])/g, '$1-$2')
    .toLowerCase()
    .replace(/[\s_]+/g, '-')

const resolveServiceIcon = (icon) => iconMap[normalizeIconKey(icon)] || ArrowRight

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && (item.name || item.icon)) : []
})

const galleryImages = computed(() => {
  const images = props.payload?.images
  if (!Array.isArray(images)) return []
  return images.filter((src) => typeof src === 'string' && src.trim() !== '')
})
const restImages = computed(() => galleryImages.value.slice(1, 5))
const moreCount = computed(() => Math.max(galleryImages.value.length - 5, 0))

const showModal = ref(false)
const currentImageIndex = ref(0)

const openImageModal = (index) => {
  if (!galleryImages.value.length) return
  currentImageIndex.value = Math.min(Math.max(index, 0), galleryImages.value.length - 1)
  showModal.value = true
  document.body.style.overflow = 'hidden'
}

const closeModal = () => {
  showModal.value = false
  document.body.style.overflow = ''
}

const nextImage = () => {
  if (currentImageIndex.value < galleryImages.value.length - 1) currentImageIndex.value++
}

const prevImage = () => {
  if (currentImageIndex.value > 0) currentImageIndex.value--
}

const onKeydown = (event) => {
  if (!showModal.value) return
  if (event.key === 'Escape') closeModal()
  else if (event.key === 'ArrowRight') nextImage()
  else if (event.key === 'ArrowLeft') prevImage()
}

onMounted(() => {
  window.addEventListener('keydown', onKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', onKeydown)
  document.body.style.overflow = ''
})
</script>

<style scoped>
.cmp-shot {
  display: block;
  width: 100%;
  overflow: hidden;
  border-radius: 16px;
  border: 1px solid rgba(15, 23, 42, 0.08);
  background: #ffffff;
  box-shadow: 0 18px 40px -26px rgba(29, 78, 216, 0.35);
}
.cmp-shot img {
  width: 100%;
  aspect-ratio: 16 / 10;
  object-fit: cover;
  transition: transform 0.55s cubic-bezier(0.22, 1, 0.36, 1);
}
.cmp-shot:hover img {
  transform: scale(1.02);
}

.cmp-tile {
  position: relative;
  display: block;
  overflow: hidden;
  border-radius: 12px;
  border: 1px solid rgba(15, 23, 42, 0.08);
  background: #ffffff;
}
.cmp-tile img {
  width: 100%;
  aspect-ratio: 1 / 1;
  object-fit: cover;
  transition: transform 0.55s cubic-bezier(0.22, 1, 0.36, 1);
}
.cmp-tile:hover img {
  transform: scale(1.03);
}
.cmp-more {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(15, 23, 42, 0.55);
  font-size: 14px;
  font-weight: 600;
  color: #ffffff;
  font-variant-numeric: tabular-nums;
}

.cmp-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 0;
  border: 1px dashed rgba(15, 23, 42, 0.16);
  border-radius: 16px;
  background: #ffffff;
}

.cmp-service {
  border-radius: 14px;
  border: 1px solid rgba(15, 23, 42, 0.07);
  background: #ffffff;
  padding: 16px;
  transition: border-color 0.2s ease, box-shadow 0.2s ease, transform 0.2s ease;
}
.cmp-service:hover {
  border-color: rgba(29, 78, 216, 0.35);
  box-shadow: 0 12px 26px -18px rgba(29, 78, 216, 0.55);
  transform: translateY(-1px);
}
.cmp-service-icon {
  display: inline-flex;
  height: 36px;
  width: 36px;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background: rgba(29, 78, 216, 0.08);
  color: #1d4ed8;
}

.cmp-card {
  border-radius: 16px;
  border: 1px solid rgba(15, 23, 42, 0.07);
  background: #ffffff;
  padding: 4px 16px 16px;
  box-shadow: 0 18px 40px -32px rgba(15, 23, 42, 0.5);
}
.cmp-contact {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 0;
  border-bottom: 1px solid #eef1f5;
}
.cmp-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  border-radius: 12px;
  background: #1d4ed8;
  padding: 13px 0;
  font-size: 14px;
  font-weight: 600;
  color: #ffffff;
  box-shadow: 0 16px 32px -16px rgba(29, 78, 216, 0.55);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.cmp-cta:hover {
  background: #1e40af;
}
.cmp-cta:active {
  transform: scale(0.98);
}

.cmp-lb-btn {
  display: flex;
  height: 40px;
  width: 40px;
  align-items: center;
  justify-content: center;
  border-radius: 9999px;
  background: rgba(255, 255, 255, 0.1);
  color: #ffffff;
  transition: background-color 0.2s ease;
}
.cmp-lb-btn:hover {
  background: rgba(255, 255, 255, 0.2);
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

<template>
  <div class="min-h-screen bg-[#f5f6fa] text-slate-900">
    <!-- Header -->
    <header class="relative overflow-hidden bg-[#172554] text-white">
      <div class="relative mx-auto w-full max-w-md px-6 pb-16 pt-10">
        <GraduationCap class="pointer-events-none absolute -right-3 top-8 h-28 w-28 text-white/[0.06]" />
        <div class="pointer-events-none absolute -left-20 -top-24 h-56 w-56 rounded-full bg-[#1e3a8a]/70 blur-3xl"></div>

        <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-blue-200/60">
          <span>Academy</span>
          <span v-if="nfcId">NFC · {{ nfcId }}</span>
        </div>
        <div class="rise rise-1 mt-9">
          <h1 class="text-balance font-serif text-[30px] leading-tight tracking-tight">{{ payload?.schoolName || '教育机构' }}</h1>
          <p class="mt-3 text-[11px] uppercase tracking-[0.28em] text-blue-200/80">知识改变未来</p>
          <span class="mt-4 block h-0.5 w-10 bg-[#f59e0b]"></span>
        </div>
      </div>
    </header>

    <div class="mx-auto w-full max-w-md px-6 pb-12">
      <!-- Stats -->
      <div v-if="courses.length || teachers.length" class="rise rise-2 relative z-10 -mt-7 flex gap-3">
        <div v-if="courses.length" class="edu-stat">
          <p class="edu-stat-num">{{ courses.length }}</p>
          <p class="edu-stat-label">门课程</p>
        </div>
        <div v-if="teachers.length" class="edu-stat">
          <p class="edu-stat-num">{{ teachers.length }}</p>
          <p class="edu-stat-label">位教师</p>
        </div>
      </div>

      <!-- Courses -->
      <div v-if="courses.length" class="rise rise-3 mt-9">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-semibold tracking-wide">热门课程</h2>
          <span class="h-px flex-1 bg-slate-200"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Courses</span>
        </div>
        <div class="mt-3 space-y-3">
          <div v-for="(c, i) in courses" :key="i" class="edu-card">
            <div class="flex items-start justify-between gap-4">
              <div class="min-w-0 flex-1">
                <div class="flex items-baseline gap-2.5">
                  <span class="font-mono text-[10px] tabular-nums text-slate-300">{{ String(i + 1).padStart(2, '0') }}</span>
                  <p class="text-[15px] font-semibold text-slate-900">{{ c?.name || '课程' }}</p>
                </div>
                <p v-if="c?.desc" class="mt-1.5 text-xs leading-relaxed text-slate-500">{{ c.desc }}</p>
              </div>
              <p v-if="c?.price || c?.price === 0" class="flex-shrink-0 text-[15px] font-bold tabular-nums text-[#b45309]">
                ¥{{ c.price }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Teachers -->
      <div v-if="teachers.length" class="rise rise-3 mt-9">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-semibold tracking-wide">师资力量</h2>
          <span class="h-px flex-1 bg-slate-200"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Faculty</span>
        </div>
        <div class="edu-panel mt-3">
          <div v-for="(t, i) in teachers" :key="i" class="edu-teacher">
            <div class="edu-avatar">
              <img v-if="t?.avatar" :src="t.avatar" :alt="t.name || '教师'" class="h-full w-full object-cover" />
              <span v-else class="font-serif text-xl text-[#1e3a8a]">{{ initialOf(t?.name) }}</span>
            </div>
            <div class="min-w-0 flex-1">
              <p class="truncate text-[14px] font-semibold text-slate-900">{{ t?.name || '教师' }}</p>
              <p v-if="t?.title" class="mt-0.5 truncate text-[11px] tracking-wide text-slate-500">{{ t.title }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Contact -->
      <div v-if="payload?.address || payload?.phone" class="rise rise-3 mt-9">
        <div v-if="payload?.address" class="flex items-start gap-3.5 rounded-2xl bg-white p-5 ring-1 ring-slate-900/5">
          <MapPin class="mt-0.5 h-4 w-4 flex-shrink-0 text-[#1e3a8a]" />
          <p class="text-[13px] leading-relaxed text-slate-600">{{ payload.address }}</p>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="edu-cta mt-4">
          <Phone class="h-4 w-4" />
          咨询报名
        </a>
      </div>

      <!-- Footer -->
      <div class="rise rise-3 mt-12 border-t border-slate-200 pt-5 text-center">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">{{ payload?.schoolName || 'Academy' }}</p>
        <p v-if="nfcId" class="mt-1.5 font-mono text-[10px] tracking-[0.18em] text-slate-300">NFC · {{ nfcId }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { GraduationCap, MapPin, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const courses = computed(() => {
  const list = props.payload?.courses
  return Array.isArray(list) ? list.filter((item) => item && (item.name || item.desc)) : []
})

const teachers = computed(() => {
  const list = props.payload?.teachers
  return Array.isArray(list) ? list.filter((item) => item && (item.name || item.title || item.avatar)) : []
})

const initialOf = (name) => {
  const value = String(name || '').trim()
  return value ? value[0] : '师'
}
</script>

<style scoped>
.edu-stat {
  flex: 1;
  border-radius: 16px;
  border: 1px solid rgba(15, 23, 42, 0.06);
  background: #ffffff;
  padding: 16px 18px;
  box-shadow: 0 18px 40px -28px rgba(23, 37, 84, 0.6);
}
.edu-stat-num {
  font-size: 22px;
  font-weight: 700;
  color: #1e3a8a;
  font-variant-numeric: tabular-nums;
}
.edu-stat-label {
  margin-top: 2px;
  font-size: 11px;
  letter-spacing: 0.08em;
  color: #94a3b8;
}

.edu-card {
  border-radius: 14px;
  border: 1px solid rgba(15, 23, 42, 0.06);
  background: #ffffff;
  padding: 18px;
  box-shadow: 0 14px 30px -26px rgba(23, 37, 84, 0.6);
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.edu-card:hover {
  border-color: rgba(30, 58, 138, 0.25);
  box-shadow: 0 16px 32px -24px rgba(30, 58, 138, 0.5);
}

.edu-panel {
  border-radius: 16px;
  border: 1px solid rgba(15, 23, 42, 0.06);
  background: #ffffff;
  padding: 4px 18px;
  box-shadow: 0 16px 36px -30px rgba(23, 37, 84, 0.6);
}
.edu-teacher {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 0;
  border-bottom: 1px solid #f1f5f9;
}
.edu-teacher:last-child {
  border-bottom: none;
}
.edu-avatar {
  display: flex;
  height: 48px;
  width: 48px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border-radius: 12px;
  border: 1px solid rgba(30, 58, 138, 0.1);
  background: rgba(30, 58, 138, 0.07);
}

.edu-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  border-radius: 14px;
  background: #1e3a8a;
  padding: 14px 0;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.06em;
  color: #ffffff;
  box-shadow: 0 18px 36px -18px rgba(30, 58, 138, 0.55);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.edu-cta:hover {
  background: #172554;
}
.edu-cta:active {
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

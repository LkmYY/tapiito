<!-- 求助信息卡 · 卡通版：圆润温和但克制，全靠图标与形状表达 -->
<template>
  <div class="min-h-screen bg-[#fdfaf1] text-slate-800">
    <!-- 顶部天空色块 -->
    <div class="relative bg-sky-500 px-6 pb-16 pt-10 text-white">
      <div class="pointer-events-none absolute -right-10 -top-12 h-36 w-36 rounded-full bg-white/10"></div>
      <div class="pointer-events-none absolute left-8 top-24 h-2.5 w-2.5 rounded-full bg-amber-300/80"></div>
      <div class="pointer-events-none absolute left-16 top-32 h-1.5 w-1.5 rounded-full bg-white/40"></div>

      <div class="relative flex items-center justify-between">
        <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-sky-100/80">Help Card</span>
        <span class="flex items-center gap-1.5 rounded-full bg-white/15 px-3 py-1 text-[11px] font-semibold text-sky-50">
          <Sparkles class="h-3 w-3" />
          随手相助
        </span>
      </div>

      <div class="relative mt-8 flex items-center gap-4">
        <div class="flex h-16 w-16 flex-shrink-0 items-center justify-center rounded-[1.35rem] bg-white text-sky-500 shadow-[0_6px_0_rgba(3,105,161,0.5)]">
          <LifeBuoy class="h-8 w-8" />
        </div>
        <div class="min-w-0">
          <h1 class="text-[26px] font-extrabold leading-tight tracking-tight">求助信息卡</h1>
          <p class="mt-1 text-[13px] font-medium text-sky-50/90">如果遇到需要帮助的人，请帮忙联系紧急联系人</p>
        </div>
      </div>

      <!-- 波浪分隔 -->
      <svg class="absolute -bottom-px left-0 w-full text-[#fdfaf1]" viewBox="0 0 390 34" preserveAspectRatio="none" aria-hidden="true">
        <path fill="currentColor" d="M0 34V17C26 6 60 2 96 9c36 7 62 19 99 19s63-12 99-19c36-7 70-3 96 8v17H0Z" />
      </svg>
    </div>

    <div class="px-5 pb-12">
      <!-- 基本情况 -->
      <div class="ct-card rise mt-6 border border-sky-100 shadow-[0_5px_0_#e0f2fe]">
        <div class="flex items-center gap-3">
          <span class="ct-icon bg-sky-100 text-sky-600"><ClipboardList class="h-5 w-5" /></span>
          <h2 class="text-[13px] font-bold uppercase tracking-[0.14em] text-sky-600">基本情况</h2>
        </div>
        <p class="mt-3.5 text-[17px] font-bold leading-snug text-slate-800">{{ payload.baseInfo }}</p>
      </div>

      <!-- 学校 / 单位 -->
      <div class="ct-card rise rise-1 mt-3.5 border border-sky-100 shadow-[0_5px_0_#e0f2fe]">
        <div class="flex items-center gap-3">
          <span class="ct-icon bg-sky-100 text-sky-600"><GraduationCap class="h-5 w-5" /></span>
          <h2 class="text-[13px] font-bold uppercase tracking-[0.14em] text-sky-600">学校/单位</h2>
        </div>
        <p class="mt-3.5 text-[15px] font-semibold text-slate-700">{{ payload.school }}</p>
      </div>

      <!-- 过敏史 / 病史：虚线卡片 + 奶油黄强调 -->
      <div class="rise rise-2 mt-3.5 rounded-[1.6rem] border-2 border-dashed border-amber-300 bg-amber-50/70 p-5 shadow-[0_5px_0_#fde68a]">
        <div class="flex items-center gap-3">
          <span class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-2xl bg-amber-200/70 text-amber-700">
            <TriangleAlert class="h-5 w-5" />
          </span>
          <h2 class="text-[13px] font-bold uppercase tracking-[0.14em] text-amber-700">过敏史/病史</h2>
        </div>
        <p class="mt-3.5 text-[15px] font-semibold text-slate-700">{{ payload.allergies || '无' }}</p>
      </div>

      <!-- 紧急联系电话 -->
      <div class="rise rise-3 mt-8">
        <p class="text-center font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">紧急联系电话</p>
        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="mt-3 flex items-center justify-center gap-3 rounded-[1.6rem] bg-sky-500 py-5 text-white shadow-[0_7px_0_#0369a1] transition-all active:translate-y-1 active:shadow-[0_3px_0_#0369a1]"
        >
          <Phone class="h-6 w-6 flex-shrink-0" />
          <span class="break-all text-[27px] font-extrabold tabular-nums tracking-tight">{{ payload.emergencyPhone }}</span>
        </a>
      </div>

      <!-- 求助反馈 -->
      <div class="ct-card rise rise-3 mt-8 border border-red-100 shadow-[0_5px_0_#fee2e2]">
        <div class="flex items-center gap-3">
          <span class="ct-icon bg-red-100 text-red-500"><Siren class="h-5 w-5" /></span>
          <h2 class="text-[13px] font-bold uppercase tracking-[0.14em] text-red-500">求助反馈</h2>
        </div>
        <textarea v-model="emergencyMessage" rows="3" class="ct-input mt-4" placeholder="发现有人需要帮助？请描述当前位置..."></textarea>
        <div class="mt-3">
          <button @click="$refs.fileInput.click()" class="ct-photo">
            <Camera class="h-4 w-4" />
            <span>{{ photoFile ? '已选照片' : '拍照' }}</span>
          </button>
          <input type="file" ref="fileInput" class="hidden" accept="image/*" capture="environment" @change="handleFileChange" />
        </div>
        <div v-if="photoPreview" class="relative mt-3 h-28 overflow-hidden rounded-2xl bg-slate-100">
          <img :src="photoPreview" class="h-full w-full object-cover" />
          <button @click="clearPhoto" class="absolute right-1.5 top-1.5 rounded-full bg-black/55 p-1 text-white">
            <X class="h-3.5 w-3.5" />
          </button>
        </div>
        <button
          @click="submitEmergencyInfo"
          :disabled="submitting"
          class="mt-3 flex w-full items-center justify-center gap-2 rounded-2xl bg-sky-500 py-3.5 text-sm font-bold text-white shadow-[0_5px_0_#0369a1] transition-all active:translate-y-0.5 active:shadow-[0_3px_0_#0369a1] disabled:opacity-60"
        >
          <Loader2 v-if="submitting" class="h-4 w-4 animate-spin" />
          <Send v-else class="h-4 w-4" />
          {{ submitting ? '发送中...' : '发送求助信息' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { LifeBuoy, Sparkles, ClipboardList, GraduationCap, TriangleAlert, Phone, Siren, Camera, Send, Loader2, X } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { uploadPhoto, submitNote } from '../../api/public'

const props = defineProps(['payload', 'nfcId'])

const emergencyMessage = ref('')
const photoFile = ref(null)
const photoPreview = ref(null)
const submitting = ref(false)
const fileInput = ref(null)

const handleFileChange = (e) => {
  const f = e.target.files[0]
  if (!f) return
  photoFile.value = f
  const r = new FileReader()
  r.onload = (ev) => { photoPreview.value = ev.target.result }
  r.readAsDataURL(f)
}

const clearPhoto = () => {
  photoFile.value = null
  photoPreview.value = null
  if (fileInput.value) fileInput.value.value = ''
}

const submitEmergencyInfo = async () => {
  if (!emergencyMessage.value.trim() && !photoFile.value) {
    ElMessage.warning('请填写信息或拍照')
    return
  }
  submitting.value = true
  try {
    let photoId = null
    if (photoFile.value) {
      const fd = new FormData()
      fd.append('file', photoFile.value)
      const { data } = await uploadPhoto(props.nfcId, fd)
      photoId = data.id
    }
    await submitNote(props.nfcId, { content: emergencyMessage.value || '求助信息', photo_ids: photoId ? [photoId] : [] })
    ElMessage.success('已发送')
    emergencyMessage.value = ''
    clearPhoto()
  } catch {
    ElMessage.error('发送失败')
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.ct-card {
  border-radius: 1.6rem;
  background: #ffffff;
  padding: 20px;
}
.ct-icon {
  display: flex;
  height: 40px;
  width: 40px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 14px;
}
.ct-input {
  width: 100%;
  resize: none;
  border-radius: 16px;
  border: 2px solid #e0f2fe;
  background: #f8fbff;
  padding: 12px 14px;
  font-size: 13px;
  line-height: 1.6;
  color: #1e293b;
  outline: none;
  transition: border-color 0.2s ease;
}
.ct-input:focus {
  border-color: #7dd3fc;
}
.ct-input::placeholder {
  color: #94a3b8;
}
.ct-photo {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-radius: 16px;
  border: 2px solid #e0f2fe;
  background: #f8fbff;
  padding: 11px 0;
  font-size: 13px;
  font-weight: 600;
  color: #0284c7;
  transition: border-color 0.2s ease;
}
.ct-photo:hover {
  border-color: #bae6fd;
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

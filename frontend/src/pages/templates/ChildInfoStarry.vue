<!-- 求助信息卡 · 星空版：深蓝夜幕 + 金色星光，柔光克制 -->
<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0a1026] text-slate-100">
    <!-- 星点背景 -->
    <div class="stars-bg pointer-events-none absolute inset-0"></div>
    <!-- 柔光 -->
    <div class="pointer-events-none absolute -top-32 left-1/2 h-72 w-72 -translate-x-1/2 rounded-full bg-[#f5c451]/10 blur-3xl"></div>
    <div class="pointer-events-none absolute -right-24 top-1/2 h-64 w-64 rounded-full bg-indigo-500/10 blur-3xl"></div>

    <!-- 四角星装饰 -->
    <svg class="pointer-events-none absolute right-9 top-24 h-4 w-4 text-amber-300/70" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
      <path d="M8 0c.6 4.4 3.6 7.4 8 8-4.4.6-7.4 3.6-8 8-.6-4.4-3.6-7.4-8-8 4.4-.6 7.4-3.6 8-8Z" />
    </svg>
    <svg class="pointer-events-none absolute left-8 top-[46%] h-2.5 w-2.5 text-slate-300/40" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
      <path d="M8 0c.6 4.4 3.6 7.4 8 8-4.4.6-7.4 3.6-8 8-.6-4.4-3.6-7.4-8-8 4.4-.6 7.4-3.6 8-8Z" />
    </svg>

    <div class="relative z-10 mx-auto flex min-h-screen w-full max-w-md flex-col px-7 pb-12 pt-12">
      <!-- 顶栏 -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.26em] text-slate-400/70">
        <span>Starry Guardian</span>
        <span v-if="nfcId" class="tabular-nums">NFC · {{ nfcId }}</span>
      </div>

      <!-- 标题 -->
      <div class="rise rise-1 mt-11 flex items-center gap-4">
        <div class="flex h-14 w-14 flex-shrink-0 items-center justify-center rounded-full bg-white/[0.06] ring-1 ring-white/15 shadow-[0_0_30px_-8px_rgba(245,196,81,0.55)]">
          <MoonStar class="h-7 w-7 text-amber-300" />
        </div>
        <div class="min-w-0">
          <h1 class="text-[24px] font-bold tracking-wide">求助信息卡</h1>
          <p class="mt-1 text-[12px] leading-relaxed text-slate-400">夜空之下，为需要帮助的人点亮方向</p>
        </div>
      </div>

      <!-- 信息面板 -->
      <div class="rise rise-2 mt-9 rounded-2xl border border-white/[0.08] bg-white/[0.03] p-6">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-amber-300/70">基本情况</p>
        <p class="mt-3 text-balance text-[19px] font-semibold leading-snug tracking-tight text-white">{{ payload.baseInfo }}</p>

        <div class="mt-5 border-t border-white/[0.08]">
          <div class="st-row">
            <GraduationCap class="h-4 w-4 flex-shrink-0 text-slate-400" />
            <span class="st-label">学校/单位</span>
            <span class="st-value">{{ payload.school }}</span>
          </div>
          <div class="st-row">
            <TriangleAlert class="h-4 w-4 flex-shrink-0 text-amber-300/80" />
            <span class="st-label">过敏史/病史</span>
            <span class="st-value">{{ payload.allergies || '无' }}</span>
          </div>
        </div>
      </div>

      <!-- 紧急联系电话 -->
      <div class="rise rise-3 mt-8">
        <p class="text-center font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400/70">紧急联系电话</p>
        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="mt-3 flex items-center justify-center gap-3 rounded-2xl border border-amber-300/35 bg-amber-300/[0.08] py-5 text-amber-200 shadow-[0_0_44px_-14px_rgba(245,196,81,0.6)] transition-all hover:bg-amber-300/[0.14] active:scale-[0.98]"
        >
          <Phone class="h-5 w-5 flex-shrink-0" />
          <span class="break-all text-[26px] font-bold tabular-nums tracking-tight">{{ payload.emergencyPhone }}</span>
        </a>
      </div>

      <!-- 紧急反馈 -->
      <div class="rise rise-3 mt-8 rounded-2xl border border-white/[0.08] bg-white/[0.03] p-6">
        <div class="flex items-center gap-2.5">
          <Siren class="h-4 w-4 text-red-400" />
          <h2 class="text-[15px] font-semibold tracking-wide text-white">紧急情况反馈</h2>
        </div>
        <p class="mt-1.5 text-[12px] leading-relaxed text-slate-400/80">留下发现地点或照片，我们会点亮信号，立即通知紧急联系人。</p>

        <textarea v-model="emergencyMessage" rows="3" class="st-input mt-4" placeholder="如发现有人需要帮助，请描述位置..."></textarea>

        <div class="mt-3">
          <button @click="$refs.fileInput.click()" class="st-photo">
            <Camera class="h-4 w-4" />
            <span>{{ photoFile ? '已选照片' : '拍照' }}</span>
          </button>
          <input type="file" ref="fileInput" class="hidden" accept="image/*" capture="environment" @change="handleFileChange" />
        </div>

        <div v-if="photoPreview" class="relative mt-3 h-28 overflow-hidden rounded-2xl bg-white/5">
          <img :src="photoPreview" class="h-full w-full object-cover" />
          <button @click="clearPhoto" class="absolute right-1.5 top-1.5 rounded-full bg-black/55 p-1 text-white">
            <X class="h-3.5 w-3.5" />
          </button>
        </div>

        <button
          @click="submitEmergencyInfo"
          :disabled="submitting"
          class="mt-4 flex w-full items-center justify-center gap-2 rounded-2xl bg-amber-300 py-3.5 text-sm font-bold text-[#0a1026] transition-all hover:bg-amber-200 active:scale-[0.98] disabled:opacity-60"
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
import { MoonStar, GraduationCap, TriangleAlert, Phone, Siren, Camera, Send, Loader2, X } from 'lucide-vue-next'
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
.stars-bg {
  background-image:
    radial-gradient(1px 1px at 12% 18%, rgba(255, 255, 255, 0.9), transparent),
    radial-gradient(1px 1px at 28% 52%, rgba(255, 255, 255, 0.5), transparent),
    radial-gradient(1.5px 1.5px at 44% 12%, rgba(255, 255, 255, 0.8), transparent),
    radial-gradient(1px 1px at 62% 34%, rgba(255, 255, 255, 0.45), transparent),
    radial-gradient(1.5px 1.5px at 78% 72%, rgba(255, 255, 255, 0.7), transparent),
    radial-gradient(1px 1px at 88% 22%, rgba(255, 255, 255, 0.6), transparent),
    radial-gradient(1px 1px at 18% 82%, rgba(255, 255, 255, 0.4), transparent),
    radial-gradient(1.5px 1.5px at 55% 68%, rgba(255, 255, 255, 0.75), transparent),
    radial-gradient(1px 1px at 36% 92%, rgba(255, 255, 255, 0.5), transparent),
    radial-gradient(1px 1px at 70% 8%, rgba(245, 196, 81, 0.7), transparent);
  animation: twinkle 5s ease-in-out infinite alternate;
}
@keyframes twinkle {
  from {
    opacity: 0.55;
  }
  to {
    opacity: 1;
  }
}

.st-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}
.st-row:last-child {
  border-bottom: none;
}
.st-label {
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  color: rgba(148, 163, 184, 0.65);
}
.st-value {
  margin-left: auto;
  max-width: 60%;
  text-align: right;
  font-size: 14px;
  font-weight: 500;
  color: #e2e8f0;
}
.st-input {
  width: 100%;
  resize: none;
  border-radius: 14px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.04);
  padding: 12px 14px;
  font-size: 13px;
  line-height: 1.6;
  color: #f1f5f9;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.st-input:focus {
  border-color: rgba(245, 196, 81, 0.5);
  box-shadow: 0 0 0 3px rgba(245, 196, 81, 0.1);
}
.st-input::placeholder {
  color: rgba(148, 163, 184, 0.55);
}
.st-photo {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-radius: 14px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  padding: 11px 0;
  font-size: 13px;
  color: rgba(203, 213, 225, 0.75);
  transition: background-color 0.2s ease, border-color 0.2s ease;
}
.st-photo:hover {
  background: rgba(255, 255, 255, 0.05);
  border-color: rgba(255, 255, 255, 0.2);
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

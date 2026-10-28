<!-- 求助信息卡 · 温暖版：蜜橙暖棕，家的感觉 -->
<template>
  <div class="min-h-screen bg-[#fdf7ee] text-stone-800">
    <!-- 屋顶形页头 -->
    <div
      class="relative bg-[#b45309] px-7 pb-20 pt-16 text-amber-50"
      style="clip-path: polygon(0 30px, 50% 0, 100% 30px, 100% 100%, 0 100%)"
    >
      <div
        class="pointer-events-none absolute inset-0 opacity-[0.06]"
        style="background-image: radial-gradient(#fff 1px, transparent 1px); background-size: 14px 14px"
      ></div>
      <div class="relative text-center">
        <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-amber-50/15 ring-1 ring-amber-50/25">
          <Home class="h-7 w-7" />
        </div>
        <h1 class="mt-4 text-[26px] font-bold tracking-tight">求助信息卡</h1>
        <p class="mt-1.5 text-[13px] text-amber-100/90">如果遇到需要帮助的人，请伸出援手 · 感谢您</p>
      </div>
    </div>

    <div class="px-6 pb-12">
      <!-- 信息卡 -->
      <div class="rise -mt-11 rounded-[1.75rem] bg-white p-6 shadow-[0_18px_40px_-22px_rgba(120,53,15,0.45)] ring-1 ring-amber-900/5">
        <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-amber-700/60">基本情况</p>
        <p class="mt-2.5 text-balance text-[20px] font-semibold leading-snug tracking-tight text-stone-900">{{ payload.baseInfo }}</p>

        <div class="mt-5 border-t border-amber-900/10">
          <div class="wm-row">
            <span class="wm-icon bg-amber-100 text-amber-700"><GraduationCap class="h-4 w-4" /></span>
            <span class="wm-label">学校/单位</span>
            <span class="wm-value">{{ payload.school }}</span>
          </div>
          <div class="wm-row">
            <span class="wm-icon bg-red-50 text-red-500"><TriangleAlert class="h-4 w-4" /></span>
            <span class="wm-label">过敏史/病史</span>
            <span class="wm-value">{{ payload.allergies || '无' }}</span>
          </div>
        </div>
      </div>

      <!-- 紧急联系电话 -->
      <div class="rise rise-2 mt-7">
        <p class="text-center font-mono text-[10px] uppercase tracking-[0.24em] text-amber-800/50">紧急联系电话</p>
        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="mt-3 flex items-center justify-between gap-4 rounded-[1.75rem] bg-[#7c2d12] px-6 py-5 text-amber-50 shadow-[0_20px_40px_-20px_rgba(124,45,18,0.6)] transition-transform active:scale-[0.98]"
        >
          <span class="min-w-0">
            <span class="block font-mono text-[10px] uppercase tracking-[0.22em] text-amber-200/60">点击拨打</span>
            <span class="mt-1.5 block break-all text-[25px] font-bold tabular-nums tracking-tight">{{ payload.emergencyPhone }}</span>
          </span>
          <span class="flex h-11 w-11 flex-shrink-0 items-center justify-center rounded-full bg-amber-50/10 ring-1 ring-amber-50/20">
            <Phone class="h-5 w-5" />
          </span>
        </a>
      </div>

      <!-- 紧急反馈 -->
      <div class="rise rise-3 mt-8 rounded-[1.75rem] bg-white p-6 shadow-[0_14px_32px_-24px_rgba(120,53,15,0.5)] ring-1 ring-amber-900/5">
        <div class="flex items-center gap-2.5">
          <Siren class="h-4 w-4 text-red-500" />
          <h2 class="text-[15px] font-semibold tracking-tight text-stone-900">紧急情况反馈</h2>
        </div>
        <p class="mt-1.5 text-[12px] leading-relaxed text-stone-400">留下发现地点或照片，我们会立即通知紧急联系人。</p>

        <textarea v-model="emergencyMessage" rows="3" class="wm-input mt-4" placeholder="如发现有人需要帮助，请描述位置..."></textarea>

        <div class="mt-3">
          <button @click="$refs.fileInput.click()" class="wm-photo">
            <Camera class="h-4 w-4" />
            <span>{{ photoFile ? '已选照片' : '拍照' }}</span>
          </button>
          <input type="file" ref="fileInput" class="hidden" accept="image/*" capture="environment" @change="handleFileChange" />
        </div>

        <div v-if="photoPreview" class="relative mt-3 h-28 overflow-hidden rounded-2xl bg-amber-50">
          <img :src="photoPreview" class="h-full w-full object-cover" />
          <button @click="clearPhoto" class="absolute right-1.5 top-1.5 rounded-full bg-black/55 p-1 text-white">
            <X class="h-3.5 w-3.5" />
          </button>
        </div>

        <button
          @click="submitEmergencyInfo"
          :disabled="submitting"
          class="mt-4 flex w-full items-center justify-center gap-2 rounded-2xl bg-[#b45309] py-3.5 text-sm font-semibold text-amber-50 transition-all hover:bg-[#92400e] active:scale-[0.98] disabled:opacity-60"
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
import { Home, GraduationCap, TriangleAlert, Phone, Siren, Camera, Send, Loader2, X } from 'lucide-vue-next'
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
.wm-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 0;
  border-bottom: 1px solid rgba(120, 53, 15, 0.08);
}
.wm-row:last-child {
  border-bottom: none;
}
.wm-icon {
  display: flex;
  height: 34px;
  width: 34px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
}
.wm-label {
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  color: rgba(120, 53, 15, 0.5);
}
.wm-value {
  margin-left: auto;
  max-width: 58%;
  text-align: right;
  font-size: 14px;
  font-weight: 500;
  color: #292524;
}
.wm-input {
  width: 100%;
  resize: none;
  border-radius: 14px;
  border: 1px solid rgba(120, 53, 15, 0.15);
  background: #fdf7ee;
  padding: 12px 14px;
  font-size: 13px;
  line-height: 1.6;
  color: #292524;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.wm-input:focus {
  border-color: #b45309;
  box-shadow: 0 0 0 3px rgba(180, 83, 9, 0.12);
}
.wm-input::placeholder {
  color: #a8a29e;
}
.wm-photo {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-radius: 14px;
  border: 1px solid rgba(120, 53, 15, 0.15);
  padding: 11px 0;
  font-size: 13px;
  color: rgba(120, 53, 15, 0.7);
  transition: background-color 0.2s ease;
}
.wm-photo:hover {
  background: rgba(120, 53, 15, 0.04);
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

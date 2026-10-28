<!-- 求助信息卡 · 经典版：可信、干净，信息一眼可读 -->
<template>
  <div class="min-h-screen bg-white text-slate-900">
    <div class="mx-auto flex min-h-screen w-full max-w-md flex-col px-7 pb-10 pt-11">
      <!-- 顶部元信息 -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.26em] text-slate-400">
        <span>Help Card</span>
        <span v-if="nfcId" class="tabular-nums">NFC · {{ nfcId }}</span>
      </div>

      <!-- 标题 -->
      <div class="rise rise-1 mt-10 flex items-start gap-4">
        <div class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-xl bg-teal-50 text-teal-700 ring-1 ring-teal-700/10">
          <Baby class="h-6 w-6" />
        </div>
        <div class="min-w-0">
          <h1 class="text-[24px] font-bold leading-tight tracking-tight">求助信息卡</h1>
          <p class="mt-1.5 text-[13px] leading-relaxed text-slate-500">如遇需要帮助的人，请立即联系紧急联系人，感谢您的善意。</p>
        </div>
      </div>

      <!-- 基础信息 -->
      <div class="rise rise-2 mt-9 border-t border-slate-200/70 pt-7">
        <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-slate-400">基础信息</p>
        <p class="mt-2.5 text-balance text-[22px] font-semibold leading-snug tracking-tight">{{ payload.baseInfo }}</p>
      </div>

      <!-- 明细 -->
      <div class="rise rise-2 mt-7">
        <div class="ci-row">
          <GraduationCap class="h-4 w-4 flex-shrink-0 text-slate-400" />
          <span class="ci-label">学校/单位</span>
          <span class="ci-value">{{ payload.school }}</span>
        </div>
        <div class="ci-row">
          <TriangleAlert class="h-4 w-4 flex-shrink-0 text-red-400" />
          <span class="ci-label">过敏史/病史</span>
          <span class="ci-value">{{ payload.allergies || '无' }}</span>
        </div>
      </div>

      <!-- 紧急电话：页面最醒目的元素 -->
      <div class="rise rise-3 mt-9">
        <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-slate-400">紧急联系电话</p>
        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="mt-3 flex items-center justify-between gap-4 rounded-2xl bg-teal-700 px-6 py-5 text-white shadow-[0_18px_36px_-18px_rgba(15,118,110,0.55)] transition-transform active:scale-[0.98]"
        >
          <span class="min-w-0">
            <span class="block font-mono text-[10px] uppercase tracking-[0.22em] text-teal-100/75">点击拨打</span>
            <span class="mt-1.5 block break-all text-[26px] font-bold tabular-nums tracking-tight">{{ payload.emergencyPhone }}</span>
          </span>
          <span class="flex h-11 w-11 flex-shrink-0 items-center justify-center rounded-full bg-white/15">
            <Phone class="h-5 w-5" />
          </span>
        </a>
      </div>

      <!-- 紧急情况反馈 -->
      <div class="rise rise-3 mt-10 rounded-2xl border border-slate-200/80 p-6">
        <div class="flex items-center gap-2.5">
          <Siren class="h-4 w-4 text-red-500" />
          <h2 class="text-[15px] font-semibold tracking-tight">紧急情况反馈</h2>
        </div>
        <p class="mt-1.5 text-[12px] leading-relaxed text-slate-400">留下发现地点或照片，我们会立即通知紧急联系人。</p>

        <textarea
          v-model="emergencyMessage"
          rows="3"
          class="ci-input mt-4"
          placeholder="如发现有人需要帮助，请在此填写所在位置等信息..."
        ></textarea>

        <div class="mt-3">
          <button @click="$refs.fileInput.click()" class="ci-photo">
            <Camera class="h-4 w-4 text-slate-400" />
            <span>{{ photoFile ? '已选择照片' : '拍照 / 上传' }}</span>
          </button>
          <input type="file" ref="fileInput" class="hidden" accept="image/*" capture="environment" @change="handleFileChange" />
        </div>

        <div v-if="photoPreview" class="relative mt-3 h-32 overflow-hidden rounded-xl bg-slate-100">
          <img :src="photoPreview" class="h-full w-full object-cover" />
          <button @click="clearPhoto" class="absolute right-1.5 top-1.5 rounded-full bg-black/55 p-1 text-white">
            <X class="h-3.5 w-3.5" />
          </button>
        </div>

        <button
          @click="submitEmergencyInfo"
          :disabled="submitting"
          class="mt-4 flex w-full items-center justify-center gap-2 rounded-xl bg-red-600 py-3.5 text-sm font-semibold text-white transition-all hover:bg-red-700 active:scale-[0.98] disabled:cursor-not-allowed disabled:opacity-70"
        >
          <Loader2 v-if="submitting" class="h-4 w-4 animate-spin" />
          <Send v-else class="h-4 w-4" />
          {{ submitting ? '正在发送...' : '发送求助信息' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Baby, GraduationCap, TriangleAlert, Phone, Siren, Camera, Send, Loader2, X } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { uploadPhoto, submitNote } from '../../api/public'

const props = defineProps(['payload', 'nfcId'])

const emergencyMessage = ref('')
const photoFile = ref(null)
const photoPreview = ref(null)
const submitting = ref(false)
const fileInput = ref(null)

const handleFileChange = (event) => {
  const file = event.target.files[0]
  if (!file) return
  photoFile.value = file
  const reader = new FileReader()
  reader.onload = (e) => { photoPreview.value = e.target.result }
  reader.readAsDataURL(file)
}

const clearPhoto = () => {
  photoFile.value = null
  photoPreview.value = null
  if (fileInput.value) fileInput.value.value = ''
}

const submitEmergencyInfo = async () => {
  if (!emergencyMessage.value.trim() && !photoFile.value) {
    ElMessage.warning('请填写信息或拍摄照片')
    return
  }
  submitting.value = true
  try {
    let photoId = null
    if (photoFile.value) {
      const formData = new FormData()
      formData.append('file', photoFile.value)
      const { data: photoData } = await uploadPhoto(props.nfcId, formData)
      photoId = photoData.id
    }
    await submitNote(props.nfcId, {
      content: emergencyMessage.value || '紧急情况（无附言）',
      photo_ids: photoId ? [photoId] : []
    })
    ElMessage.success('求助信息已发送')
    emergencyMessage.value = ''
    clearPhoto()
  } catch {
    ElMessage.error('发送失败，请稍后重试')
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.ci-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 15px 0;
  border-bottom: 1px solid rgba(15, 23, 42, 0.07);
}
.ci-row:first-child {
  border-top: 1px solid rgba(15, 23, 42, 0.07);
}
.ci-label {
  flex-shrink: 0;
  width: 56px;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  color: #94a3b8;
}
.ci-value {
  margin-left: auto;
  max-width: 62%;
  text-align: right;
  font-size: 14px;
  font-weight: 500;
  color: #0f172a;
}
.ci-input {
  width: 100%;
  resize: none;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  background: #f8fafc;
  padding: 12px 14px;
  font-size: 13px;
  line-height: 1.6;
  color: #0f172a;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.ci-input:focus {
  border-color: #0f766e;
  box-shadow: 0 0 0 3px rgba(15, 118, 110, 0.12);
}
.ci-input::placeholder {
  color: #94a3b8;
}
.ci-photo {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  padding: 11px 0;
  font-size: 13px;
  color: #64748b;
  transition: background-color 0.2s ease, border-color 0.2s ease;
}
.ci-photo:hover {
  background: #f8fafc;
  border-color: #cbd5e1;
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

<!-- 求助信息卡 · 自然版：森系纸感，路径节点式布局 -->
<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f5f2e9] text-emerald-950">
    <div class="pointer-events-none absolute -right-20 -top-24 h-64 w-64 rounded-full bg-emerald-200/35 blur-3xl"></div>
    <div class="pointer-events-none absolute -left-24 top-1/2 h-56 w-56 rounded-full bg-amber-100/50 blur-3xl"></div>
    <Leaf class="pointer-events-none absolute -right-10 top-28 h-44 w-44 rotate-12 text-emerald-900/[0.05]" />

    <div class="relative z-10 mx-auto flex min-h-screen w-full max-w-md flex-col px-7 pb-12 pt-12">
      <!-- 顶栏 -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.26em] text-emerald-900/40">
        <span>Nature Guardian</span>
        <span v-if="nfcId" class="tabular-nums">NFC · {{ nfcId }}</span>
      </div>

      <!-- 标题 -->
      <div class="rise rise-1 mt-10">
        <div class="flex items-center gap-3">
          <span class="flex h-10 w-10 items-center justify-center rounded-full bg-emerald-900 text-emerald-50">
            <Sprout class="h-5 w-5" />
          </span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-emerald-800/50">求助信息</span>
        </div>
        <h1 class="mt-5 font-serif text-[32px] leading-[1.15] tracking-tight">求助信息卡</h1>
        <p class="mt-2 text-[13px] leading-relaxed text-emerald-900/55">如果遇到需要帮助的人，请联系紧急联系人。</p>
      </div>

      <!-- 路径节点 -->
      <div class="rise rise-2 mt-9 flex-1 border-l border-dashed border-emerald-900/20 pl-7">
        <section class="relative pb-8">
          <span class="nt-dot"></span>
          <p class="nt-label">基本情况</p>
          <p class="mt-2 text-balance text-[20px] font-semibold leading-snug tracking-tight">{{ payload.baseInfo }}</p>
        </section>

        <section class="relative border-t border-emerald-900/10 pb-8 pt-7">
          <span class="nt-dot"></span>
          <p class="nt-label">学校/单位</p>
          <p class="mt-2 text-[15px] font-medium text-emerald-950/85">{{ payload.school }}</p>
        </section>

        <section class="relative border-t border-emerald-900/10 pb-8 pt-7">
          <span class="nt-dot nt-dot-warn"></span>
          <p class="nt-label">过敏史/病史</p>
          <p class="mt-2 flex items-center gap-2 text-[15px] font-medium text-emerald-950/85">
            <TriangleAlert class="h-4 w-4 flex-shrink-0 text-amber-700" />
            <span>{{ payload.allergies || '无' }}</span>
          </p>
        </section>

        <section class="relative border-t border-emerald-900/10 pt-7">
          <span class="nt-dot nt-dot-solid"></span>
          <p class="nt-label">紧急联系电话</p>
          <a
            v-if="payload.emergencyPhone"
            :href="'tel:' + payload.emergencyPhone"
            class="mt-3.5 flex items-center justify-between gap-4 rounded-[1.75rem] bg-emerald-900 px-6 py-5 text-emerald-50 shadow-[0_20px_40px_-20px_rgba(6,78,59,0.6)] transition-transform active:scale-[0.98]"
          >
            <span class="min-w-0">
              <span class="block font-mono text-[10px] uppercase tracking-[0.22em] text-emerald-200/60">点击拨打</span>
              <span class="mt-1.5 block break-all text-[25px] font-bold tabular-nums tracking-tight">{{ payload.emergencyPhone }}</span>
            </span>
            <span class="flex h-11 w-11 flex-shrink-0 items-center justify-center rounded-full bg-emerald-50/10 ring-1 ring-emerald-50/20">
              <Phone class="h-5 w-5" />
            </span>
          </a>
        </section>
      </div>

      <!-- 紧急反馈 -->
      <div class="rise rise-3 mt-9 rounded-[1.75rem] border border-emerald-900/10 bg-white/60 p-6">
        <div class="flex items-center gap-2.5">
          <Siren class="h-4 w-4 text-amber-700" />
          <h2 class="text-[15px] font-semibold tracking-tight">紧急情况反馈</h2>
        </div>
        <p class="mt-1.5 text-[12px] leading-relaxed text-emerald-900/45">留下发现地点或照片，我们会立即通知紧急联系人。</p>

        <textarea v-model="emergencyMessage" rows="3" class="nt-input mt-4" placeholder="如发现有人需要帮助，请描述位置..."></textarea>

        <div class="mt-3">
          <button @click="$refs.fileInput.click()" class="nt-photo">
            <Camera class="h-4 w-4" />
            <span>{{ photoFile ? '已选照片' : '拍照' }}</span>
          </button>
          <input type="file" ref="fileInput" class="hidden" accept="image/*" capture="environment" @change="handleFileChange" />
        </div>

        <div v-if="photoPreview" class="relative mt-3 h-28 overflow-hidden rounded-2xl bg-emerald-900/5">
          <img :src="photoPreview" class="h-full w-full object-cover" />
          <button @click="clearPhoto" class="absolute right-1.5 top-1.5 rounded-full bg-black/55 p-1 text-white">
            <X class="h-3.5 w-3.5" />
          </button>
        </div>

        <button
          @click="submitEmergencyInfo"
          :disabled="submitting"
          class="mt-4 flex w-full items-center justify-center gap-2 rounded-2xl bg-emerald-800 py-3.5 text-sm font-semibold text-emerald-50 shadow-[0_14px_28px_-16px_rgba(6,78,59,0.55)] transition-all hover:bg-emerald-900 active:scale-[0.98] disabled:opacity-60"
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
import { Leaf, Sprout, TriangleAlert, Phone, Siren, Camera, Send, Loader2, X } from 'lucide-vue-next'
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
.nt-dot {
  position: absolute;
  left: -1.75rem;
  top: 5px;
  height: 9px;
  width: 9px;
  transform: translateX(-50%);
  border-radius: 9999px;
  background: #f5f2e9;
  border: 2px solid #065f46;
}
.nt-dot-warn {
  border-color: #b45309;
}
.nt-dot-solid {
  background: #064e3b;
  border-color: #064e3b;
}
.nt-label {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  text-transform: uppercase;
  letter-spacing: 0.22em;
  color: rgba(6, 78, 59, 0.5);
}
.nt-input {
  width: 100%;
  resize: none;
  border-radius: 14px;
  border: 1px solid rgba(6, 78, 59, 0.15);
  background: rgba(255, 255, 255, 0.75);
  padding: 12px 14px;
  font-size: 13px;
  line-height: 1.6;
  color: #064e3b;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.nt-input:focus {
  border-color: #047857;
  box-shadow: 0 0 0 3px rgba(4, 120, 87, 0.12);
}
.nt-input::placeholder {
  color: rgba(6, 78, 59, 0.35);
}
.nt-photo {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-radius: 14px;
  border: 1px solid rgba(6, 78, 59, 0.15);
  padding: 11px 0;
  font-size: 13px;
  color: rgba(6, 78, 59, 0.7);
  transition: background-color 0.2s ease;
}
.nt-photo:hover {
  background: rgba(6, 78, 59, 0.04);
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

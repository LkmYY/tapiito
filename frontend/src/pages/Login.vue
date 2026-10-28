<template>
  <div class="auth-shell">
    <!-- 背景：柔和光晕 + 细网格 -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div class="absolute -left-32 -top-32 h-[420px] w-[420px] rounded-full bg-[#ff6b35]/20 blur-[110px]" />
      <div class="absolute -right-24 top-24 h-[380px] w-[380px] rounded-full bg-[#ffd166]/25 blur-[110px]" />
      <div class="absolute bottom-[-160px] left-1/3 h-[420px] w-[420px] rounded-full bg-[#7bdff2]/20 blur-[120px]" />
      <div class="auth-grid absolute inset-0" />
    </div>

    <div class="relative z-10 w-full max-w-[420px] px-5 py-10">
      <!-- 品牌 -->
      <router-link to="/" class="mb-8 flex items-center justify-center gap-2.5 group">
        <span class="brand-dot">
          <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-5 w-5" draggable="false" />
        </span>
        <span class="text-[22px] font-semibold tracking-tight text-[#14181d]">Tapiito</span>
      </router-link>

      <div class="auth-card">
        <header class="mb-7 text-center">
          <h1 class="text-[26px] font-semibold tracking-tight text-[#14181d]">欢迎回来</h1>
          <p class="mt-2 text-[13px] leading-relaxed text-[#14181d]/55">
            Tap it alive. 碰一下，它就活了
          </p>
        </header>

        <form @submit.prevent="onSubmit" class="space-y-4">
          <div>
            <label class="auth-label">管理员账号</label>
            <div class="auth-field">
              <User class="auth-icon" />
              <input
                v-model="form.username"
                type="text"
                class="auth-input"
                placeholder="输入管理员账号"
                autocomplete="username"
              />
            </div>
          </div>

          <div>
            <label class="auth-label">密码</label>
            <div class="auth-field">
              <Lock class="auth-icon" />
              <input
                v-model="form.password"
                type="password"
                class="auth-input"
                placeholder="输入密码"
                autocomplete="current-password"
              />
            </div>
          </div>

          <button type="submit" class="auth-submit" :disabled="loading">
            <Loader2 v-if="loading" class="h-4 w-4 animate-spin" />
            <template v-else>
              登录
              <ArrowRight class="h-4 w-4 transition-transform duration-300 group-hover:translate-x-0.5" />
            </template>
          </button>
        </form>

      </div>

      <p class="mt-7 text-center text-[11px] tracking-[0.08em] text-[#14181d]/35">
        Tapiito 管理后台
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { User, Lock, Loader2, Nfc, ArrowRight } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()
const loading = ref(false)
const form = reactive({ username: '', password: '' })

const onSubmit = async () => {
  if (!form.username.trim()) {
    ElMessage.warning('请输入用户名')
    return
  }
  if (!form.password.trim()) {
    ElMessage.warning('请输入密码')
    return
  }
  loading.value = true
  try {
    await authStore.loginAdmin(form)
  } catch {
    // 错误已由 request 拦截器处理
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.auth-shell {
  position: relative;
  display: flex;
  min-height: 100vh;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  background: #fbfaf7;
}

.auth-grid {
  background-image:
    linear-gradient(to right, rgba(20, 24, 29, 0.045) 1px, transparent 1px),
    linear-gradient(to bottom, rgba(20, 24, 29, 0.045) 1px, transparent 1px);
  background-size: 56px 56px;
  mask-image: radial-gradient(circle at 50% 40%, black, transparent 72%);
  -webkit-mask-image: radial-gradient(circle at 50% 40%, black, transparent 72%);
}

.brand-dot {
  display: flex;
  height: 30px;
  width: 30px;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  transition: transform 0.3s cubic-bezier(0.22, 1, 0.36, 1);

  background: #15151B;}

.group:hover .brand-dot {
  transform: translateY(-1px) rotate(-4deg);
}

.auth-card {
  border-radius: 22px;
  border: 1px solid rgba(20, 24, 29, 0.07);
  background: rgba(255, 255, 255, 0.82);
  padding: 30px 26px 26px;
  box-shadow:
    0 1px 2px rgba(20, 24, 29, 0.04),
    0 24px 48px -28px rgba(20, 24, 29, 0.28);
  backdrop-filter: blur(18px);
}

.auth-label {
  display: block;
  margin-bottom: 7px;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.02em;
  color: rgba(20, 24, 29, 0.62);
}

.auth-field {
  position: relative;
}

.auth-icon {
  position: absolute;
  left: 13px;
  top: 50%;
  height: 16px;
  width: 16px;
  transform: translateY(-50%);
  color: rgba(20, 24, 29, 0.28);
  transition: color 0.25s ease;
}

.auth-field:focus-within .auth-icon {
  color: #ff6b35;
}

.auth-input {
  width: 100%;
  border-radius: 12px;
  border: 1px solid rgba(20, 24, 29, 0.1);
  background: rgba(255, 255, 255, 0.7);
  padding: 11px 13px 11px 38px;
  font-size: 14px;
  color: #14181d;
  outline: none;
  transition: border-color 0.25s ease, box-shadow 0.25s ease, background 0.25s ease;
}

.auth-input::placeholder {
  color: rgba(20, 24, 29, 0.28);
}

.auth-input:focus {
  border-color: rgba(255, 107, 53, 0.55);
  background: #fff;
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.12);
}

.auth-submit {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  gap: 7px;
  border-radius: 12px;
  background: #14181d;
  padding: 12px;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.01em;
  color: #fff;
  transition: transform 0.2s cubic-bezier(0.22, 1, 0.36, 1), background 0.25s ease, box-shadow 0.25s ease;
}

.auth-submit:hover:not(:disabled) {
  background: #ff6b35;
  transform: translateY(-1px);
  box-shadow: 0 10px 22px -10px rgba(255, 107, 53, 0.65);
}

.auth-submit:active:not(:disabled) {
  transform: translateY(0) scale(0.99);
}

.auth-submit:disabled {
  opacity: 0.55;
  cursor: not-allowed;
}
</style>

<template>
  <CheckCodeLink v-if="!hidden" @click="openModal" class="link"/>
  <Transition name="fade">
    <div v-if="isOpen" class="modal-overlay" @click="closeModal" :ref="isOpen">
      <div class="modal-container" @click.stop>
        <div class="modal-header">
          <h2 class="modal-title">{{ title }}</h2>
          <button class="modal-close" @click="closeModal" aria-label="Закрыть">
            <svg
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
            >
              <line x1="18" y1="6" x2="6" y2="18"></line>
              <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
          </button>
        </div>

        <div class="modal-content">
          <slot>
          </slot>
          <SendButton style="top: -5%" @click="changeState"/>
        </div>
      </div>
      <NeuronAnswer :state-answer="answer" :ref="answer"/>
    </div>
  </Transition>
</template>

<script setup>
import {watch, onMounted, onUnmounted, defineProps, ref} from 'vue'
import CheckCodeLink from "@/components/CheckCodeLink.vue";
import NeuronAnswer from "@/components/NeuronAnswer.vue";
import SendButton from "@/components/SendButton.vue";

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    default: 'Модальное окно'
  },
  hidden : {
    type: Boolean,
    default: true
  }
})
console.log("Окей")
const isOpen = ref(false)

const answer = ref("error")

const changeState = () => {
  if (answer.value === 'ok') {
    answer.value = 'error'
  } else if (answer.value === 'error'){
    answer.value = 'hidden'
  } else {
    answer.value = 'ok'
  }
}

const closeModal = () => {
  isOpen.value = false
}

const openModal = () => {
  isOpen.value = true
}

const handleEscape = (e) => {
  if (e.key === 'Escape') {
    closeModal()
  }
}

// Управление скроллом и слушателем клавиатуры
watch(() => props.modelValue, (newValue) => {
  console.log("Вот тут")
  if (newValue) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
})

onMounted(() => {
  document.addEventListener('keydown', handleEscape)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscape)
  document.body.style.overflow = ''
})
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 1rem;
}

.modal-container {
  background-color: white;
  height: 100%;
  border-radius: 12px;
  width: 100%;
  display: flex;
  flex-direction: column;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1),
  0 10px 10px -5px rgba(0, 0, 0, 0.04);
  overflow: hidden;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  border-bottom: 1px solid #e5e7eb;
}

.modal-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #111827;
  margin: 0;
}

.modal-close {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  padding: 0;
  background: transparent;
  border: none;
  border-radius: 6px;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.modal-close:hover {
  background-color: #f3f4f6;
  color: #111827;
}

.modal-content {
  padding: 1.5rem;
  overflow-y: auto;
}

/* Transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Scrollbar Styling */
.modal-content::-webkit-scrollbar {
  width: 8px;
}

.modal-content::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.modal-content::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 4px;
}

.modal-content::-webkit-scrollbar-thumb:hover {
  background: #a1a1a1;
}

@media (max-width: 640px) {
  .modal-container {
    margin: 0;
    max-height: calc(100vh - 2rem);
  }

  .modal-header {
    padding: 1rem;
  }

  .modal-content {
    padding: 1rem;
  }
}

.link {
  position: relative;
  left: 70%;
  top: -10%;
}
</style>
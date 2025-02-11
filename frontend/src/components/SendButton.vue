<template>
  <button
      class="custom-button"
      :disabled="props.disabled"
      @click="handleClick"
      :class="{ 'button-clicked': isClicked }"
  >
    <span class="button-text">Отправить</span>
    <div class="button-particles">
      <span v-for="n in 4" :key="n" class="particle"></span>
    </div>
    <div class="button-glow"></div>
  </button>
</template>

<script setup>
import { ref, defineProps, defineEmits } from 'vue'

const props = defineProps({
  disabled: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['click'])

const isClicked = ref(false)

const handleClick = () => {
  isClicked.value = true
  emit('click')

  setTimeout(() => {
    isClicked.value = false
  }, 500)
}
</script>

<style scoped>
.custom-button {
  position: relative;
  margin-top: 5%;
  top: 10%;
  padding: 16px 32px;
  font-size: 18px;
  color: #fff;
  background: linear-gradient(135deg, #2d2d2d 0%, #1a1a1a 100%);
  border: none;
  border-radius: 8px;
  cursor: pointer;
  overflow: hidden;
  transition: transform 0.2s ease;
}

.button-text {
  position: relative;
  z-index: 2;
  font-weight: 600;
  letter-spacing: 0.5px;
  background: linear-gradient(to right, #fff, #e0e0e0);
  -webkit-background-clip: text;
  background-clip: text;
  transition: all 0.3s ease;
}

/* Glow Effect */
.button-glow {
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(
      circle,
      rgba(255, 255, 255, 0.1) 0%,
      transparent 70%
  );
  transform: scale(0);
  transition: transform 0.5s ease;
}

/* Particles */
.button-particles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
}

.particle {
  position: absolute;
  width: 4px;
  height: 4px;
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.8);
  opacity: 0;
}

.particle:nth-child(1) { top: 25%; left: 25%; }
.particle:nth-child(2) { top: 25%; right: 25%; }
.particle:nth-child(3) { bottom: 25%; left: 25%; }
.particle:nth-child(4) { bottom: 25%; right: 25%; }

/* Hover Effects */
.custom-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
}

.custom-button:hover .button-text {
  background: linear-gradient(to right, #fff, #fff);
  -webkit-background-clip: text;
  background-clip: text;
}

.custom-button:hover .button-glow {
  transform: scale(1) rotate(45deg);
}

.custom-button:hover .particle {
  animation: particle-burst 0.6s ease-out forwards;
}

.custom-button:hover .particle:nth-child(1) { animation-delay: 0.1s; }
.custom-button:hover .particle:nth-child(2) { animation-delay: 0.2s; }
.custom-button:hover .particle:nth-child(3) { animation-delay: 0.3s; }
.custom-button:hover .particle:nth-child(4) { animation-delay: 0.4s; }

/* Click Effect */
.button-clicked {
  transform: scale(0.95);
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.3);
}

/* Disabled State */
.custom-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

.custom-button:disabled .button-text {
  background: linear-gradient(to right, #999, #777);
  -webkit-background-clip: text;
  background-clip: text;
}

/* Particle Animation */
@keyframes particle-burst {
  0% {
    transform: translate(0, 0) scale(1);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    transform: translate(
        calc(var(--direction-x, 1) * 50px),
        calc(var(--direction-y, 1) * 50px)
    ) scale(0);
    opacity: 0;
  }
}

/* Focus State for Accessibility */
.custom-button:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(255, 255, 255, 0.3);
}

/* Active State */
.custom-button:active {
  transform: scale(0.95);
}

@media (prefers-reduced-motion: reduce) {
  .custom-button,
  .button-text,
  .button-glow,
  .particle {
    transition: none;
    animation: none;
  }

  .custom-button:hover {
    transform: none;
  }
}
</style>